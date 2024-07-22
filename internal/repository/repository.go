package repository

import (
	// "backend/internal"

	"backend/internal/service"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewRepository(db *gorm.DB) service.Repository {
	return &DB{
		Repository: NewPostgresRepository(db),
		Cache:      InitCacheLRU(),
	}
}

type DB struct { //CashProxyRepository
	service.Repository
	Cache *lru.Cache[int, float64]
}

func (db *DB) CreateWallet() (int, float64, error) {
	walletID, balance, err := db.Repository.CreateWallet()
	if err == nil {
		db.Cache.Add(walletID, balance)
	}
	return walletID, balance, err
}

func (db *DB) WalletInfo(id int) (int, float64, error) {
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
