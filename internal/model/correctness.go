package model

import (
	"math/big"

	"github.com/KenshiTech/unchained/internal/crypto/bls"
	"github.com/KenshiTech/unchained/internal/utils"
	bls12381 "github.com/consensys/gnark-crypto/ecc/bls12-381"

	sia "github.com/pouya-eghbali/go-sia/v2/pkg"
)

type CorrectnessReportPacket struct {
	Correctness
	Signature [48]byte
}

func (c *CorrectnessReportPacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(c.Correctness.Sia().Bytes()).
		AddByteArray8(c.Signature[:])
}

func (c *CorrectnessReportPacket) FromBytes(payload []byte) *CorrectnessReportPacket {
	c.Correctness.FromBytes(payload)
	copy(c.Signature[:], sia.NewFromBytes(payload).ReadByteArray8())

	return c
}

//////

type BroadcastCorrectnessPacket struct {
	Info      Correctness
	Signature [48]byte
	Signer    Signer
}

func (b *BroadcastCorrectnessPacket) Sia() sia.Sia {
	return sia.New().
		EmbedBytes(b.Info.Sia().Bytes()).
		AddByteArray8(b.Signature[:]).
		EmbedBytes(b.Signer.Sia().Bytes())
}

func (b *BroadcastCorrectnessPacket) FromBytes(payload []byte) *BroadcastCorrectnessPacket {
	b.Info.FromBytes(payload)
	copy(b.Signature[:], sia.NewFromBytes(payload).ReadByteArray8())
	b.Signer.FromBytes(payload)

	return b
}

type Correctness struct {
	SignersCount uint64
	Signature    []byte
	Consensus    bool
	Voted        big.Int
	SignerIDs    []int
	Timestamp    uint64
	Hash         []byte
	Topic        [64]byte
	Correct      bool
}

func (c *Correctness) Sia() sia.Sia {
	return sia.New().
		AddUInt64(c.Timestamp).
		AddByteArray8(c.Hash).
		AddByteArray8(c.Topic[:]).
		AddBool(c.Correct)
}

func (c *Correctness) FromBytes(payload []byte) *Correctness {
	siaMessage := sia.NewFromBytes(payload)
	c.Timestamp = siaMessage.ReadUInt64()
	copy(c.Hash, siaMessage.ReadByteArray8())
	copy(c.Topic[:], siaMessage.ReadByteArray8())
	c.Correct = siaMessage.ReadBool()

	return c
}

func (c *Correctness) Bls() (bls12381.G1Affine, error) {
	hash, err := bls.Hash(c.Sia().Bytes())
	if err != nil {
		utils.Logger.Error("Can't hash bls: %v", err)
		return bls12381.G1Affine{}, err
	}

	return hash, err
}
