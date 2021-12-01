package token

import (
	"github.com/ethereum/go-ethereum/ethclient"
	bscClient "github.com/ethereum/go-ethereum/ethclient"
	"gitlab.com/distributed_lab/kit/pgdb"
	"gitlab.com/distributed_lab/logan/v3"
	"gitlab.com/tokend/subgroup/tokenproject/internal/config"
	"gitlab.com/tokend/subgroup/tokenproject/internal/data"
	"gitlab.com/tokend/subgroup/tokenproject/internal/data/postgres"
)

type Service struct {
	bscClient        bscClient.Client
	tokenQ          data.TokenQ
	contractAddress  config.ContractConfig
	log              *logan.Entry
	db               *pgdb.DB
	eth              *ethclient.Client
}

func New(cfg config.Config) *Service {
	log := cfg.Log().WithField("main_service", "checker")

	return &Service{
		tokenQ:          postgres.NewTokenQ(cfg.DB()),
		contractAddress:  cfg.ContractConfig(),
		log:              log,
		db:               cfg.DB(),
		eth:              cfg.EthClient(),
	}
}
