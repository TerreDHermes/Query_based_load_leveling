package repository

import (
	"backend"
	"backend/internal/service"
	"errors"

	"gorm.io/gorm"
)

var _ service.Repository = (*PostgresDB)(nil)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresRepository(db *gorm.DB) *PostgresDB {
	return &PostgresDB{db: db}
}

func (r *PostgresDB) CreateWallet() (int, float64, error) {
	wallet := backend.Wallets{Balance: 100.00}
	if err := r.db.Create(&wallet).Error; err != nil {
		return 0, 0, err
	}
	return wallet.ID, wallet.Balance, nil
}

func (r *PostgresDB) WalletInfo(id int) (int, float64, error) {
	var wallet backend.Wallets
	if err := r.db.First(&wallet, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, 0, errors.New(" wallet not found")
		}
		return 0, 0, err
	}
	return wallet.ID, wallet.Balance, nil
}
