package pos

import (
	"context"
	"github.com/TimeleapLabs/unchained/internal/config"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum/contracts"
	"github.com/TimeleapLabs/unchained/internal/service/pos/eip712"
	"github.com/TimeleapLabs/unchained/internal/utils"
	"github.com/TimeleapLabs/unchained/internal/utils/address"
	"github.com/puzpuzpuz/xsync/v3"
	"math/big"
)

type Service interface {
	GetTotalVotingPower() (*big.Int, error)
	GetVotingPowerFromContract(address [20]byte, block *big.Int) (*big.Int, error)
	GetVotingPower(address [20]byte, block *big.Int) (*big.Int, error)
	GetVotingPowerOfPublicKey(ctx context.Context, pkBytes [96]byte) (*big.Int, error)
}

type service struct {
	ethRPC       ethereum.RPC
	posContract  *contracts.UnchainedStaking
	votingPowers *xsync.MapOf[[20]byte, *big.Int]
	lastUpdated  *xsync.MapOf[[20]byte, *big.Int]
	base         *big.Int
	eip712Signer *eip712.Signer
}

func (s *service) GetTotalVotingPower() (*big.Int, error) {
	return s.posContract.GetTotalVotingPower(nil)
}

func (s *service) GetVotingPowerFromContract(address [20]byte, block *big.Int) (*big.Int, error) {
	votingPower, err := s.posContract.GetVotingPower(nil, address)
	if err != nil {
		return votingPower, err
	}

	s.votingPowers.Store(address, votingPower)
	s.lastUpdated.Store(address, block)

	return votingPower, nil
}

func (s *service) minBase(power *big.Int) *big.Int {
	if power == nil || power.Cmp(s.base) < 0 {
		return s.base
	}

	return power
}

func (s *service) GetVotingPower(address [20]byte, block *big.Int) (*big.Int, error) {
	powerLastUpdated, ok := s.lastUpdated.Load(address)
	if !ok {
		powerLastUpdated = big.NewInt(0)
	}

	updateAt := new(big.Int).Add(powerLastUpdated, big.NewInt(25000))

	if block.Cmp(updateAt) > 0 {
		votingPower, err := s.GetVotingPowerFromContract(address, block)
		return s.minBase(votingPower), err
	}

	if votingPower, ok := s.votingPowers.Load(address); ok {
		return s.minBase(votingPower), nil
	}

	return s.base, nil
}

func (s *service) GetVotingPowerOfPublicKey(ctx context.Context, pkBytes [96]byte) (*big.Int, error) {
	_, addrHex := address.CalculateHex(pkBytes[:])
	block, err := s.ethRPC.GetBlockNumber(ctx, config.App.ProofOfStake.Chain)
	if err != nil {
		return nil, err
	}
	return s.GetVotingPower(addrHex, big.NewInt(int64(block)))
}

func New(ethRPC ethereum.RPC) Service {
	s := &service{
		ethRPC:       ethRPC,
		base:         big.NewInt(config.App.ProofOfStake.Base),
		votingPowers: xsync.NewMapOf[[20]byte, *big.Int](),
		lastUpdated:  xsync.NewMapOf[[20]byte, *big.Int](),
	}

	pkBytes := crypto.Identity.Bls.PublicKey.Bytes()
	addrHexStr, addrHex := address.CalculateHex(pkBytes[:])

	utils.Logger.
		With("Address", addrHexStr).
		Info("PoS identity initialized")

	var err error

	s.posContract, err = s.ethRPC.GetNewStakingContract(
		config.App.ProofOfStake.Chain,
		config.App.ProofOfStake.Address,
		false,
	)
	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Failed to connect to the staking contract")

		panic(err)
	}

	power, err := s.GetVotingPower(addrHex, big.NewInt(0))
	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Failed to get voting power")

		panic(err)
	}

	total, err := s.GetTotalVotingPower()
	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Failed to get total voting power")

		panic(err)
	}

	utils.Logger.
		With("Power", utils.BigIntToFloat(power)).
		With("Network", utils.BigIntToFloat(total)).
		Info("PoS")

	chainID, err := s.posContract.GetChainId(nil)
	if err != nil {
		utils.Logger.
			With("Error", err).
			Error("Failed to get chain ID")

		panic(err)
	}

	s.eip712Signer = eip712.New(chainID, config.App.ProofOfStake.Address)

	return s
}
