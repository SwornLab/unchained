package app

import (
	"github.com/TimeleapLabs/unchained/internal/consts"
	"github.com/TimeleapLabs/unchained/internal/crypto"
	"github.com/TimeleapLabs/unchained/internal/crypto/ethereum"
	postgresRepo "github.com/TimeleapLabs/unchained/internal/repository/postgres"
	correctnessService "github.com/TimeleapLabs/unchained/internal/service/correctness"
	evmlogService "github.com/TimeleapLabs/unchained/internal/service/evmlog"
	"github.com/TimeleapLabs/unchained/internal/service/pos"
	uniswapService "github.com/TimeleapLabs/unchained/internal/service/uniswap"
	"github.com/TimeleapLabs/unchained/internal/transport/client"
	"github.com/TimeleapLabs/unchained/internal/transport/client/conn"
	"github.com/TimeleapLabs/unchained/internal/transport/client/handler"
	"github.com/TimeleapLabs/unchained/internal/transport/database/postgres"
	"github.com/TimeleapLabs/unchained/internal/utils"
)

// Consumer starts the Unchained consumer and contains its DI.
func Consumer() {
	utils.Logger.
		With("Mode", "Consumer").
		With("Version", consts.Version).
		With("Protocol", consts.ProtocolVersion).
		Info("Running Unchained")

	crypto.InitMachineIdentity(
		crypto.WithEvmSigner(),
		crypto.WithBlsIdentity(),
	)

	ethRPC := ethereum.New()
	pos := pos.New(ethRPC)
	db := postgres.New()

	eventLogRepo := postgresRepo.NewEventLog(db)
	signerRepo := postgresRepo.NewSigner(db)
	assetPrice := postgresRepo.NewAssetPrice(db)
	correctnessRepo := postgresRepo.NewCorrectness(db)

	correctnessService := correctnessService.New(pos, signerRepo, correctnessRepo)
	evmLogService := evmlogService.New(ethRPC, pos, eventLogRepo, signerRepo, nil)
	uniswapService := uniswapService.New(ethRPC, pos, signerRepo, assetPrice)

	conn.Start()

	handler := handler.NewConsumerHandler(correctnessService, uniswapService, evmLogService)
	client.NewRPC(handler)

	select {}
}
