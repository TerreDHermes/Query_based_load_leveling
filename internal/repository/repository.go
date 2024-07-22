package repository

import (
	"backend/internal/service"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CashProxyRepository struct {
	service.Repository
	Cache *lru.Cache[int, float64]
}

func NewRepository(db *gorm.DB) service.Repository {
	return &CashProxyRepository{
		Repository: NewPostgresRepository(db),
		Cache:      InitCacheLRU(),
	}
}

func (db *CashProxyRepository) CreateWallet() (int, float64, error) {
	walletID, balance, err := db.Repository.CreateWallet()
	if err == nil {
		db.Cache.Add(walletID, balance)
	}
	return walletID, balance, err
}

func (db *CashProxyRepository) WalletInfo(id int) (int, float64, error) {
	var err error
	balance, ok := db.Cache.Get(id)
	if ok == true {
		logrus.Info("Чтение произошло из кэша")
	} else {
		id, balance, err = db.Repository.WalletInfo(id)
		if err != nil {
			return 0, 0, err
		} else {
			logrus.Info("Чтение произошло из Postgres")
		}
	}
	return id, balance, nil
}
