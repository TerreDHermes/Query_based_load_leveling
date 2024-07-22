package service

import (
	"sync"
)

type Repository interface {
	CreateWallet() (int, float64, error)
	WalletInfo(id int) (int, float64, error)
}

type walletResult struct {
	ID  int
	Err error
}

type Service struct {
	Repository
	QueryQueue chan func()
	QueueWG    sync.WaitGroup
}

func NewService(repos Repository) *Service {
	s := &Service{
		Repository: repos,
		QueryQueue: make(chan func(), 100), // Размер буфера очереди 100
	}

	// Создаем 10 воркеров
	for i := 0; i < 10; i++ {
		s.QueueWG.Add(1)
		go s.HandlerTask()
	}
	return s
}

func (s *Service) HandlerTask() {
	defer s.QueueWG.Done()
	for task := range s.QueryQueue {
		task()
	}
}

func (s *Service) CreateWallet() (int, error) {
	resultChan := make(chan walletResult)

	s.QueryQueue <- func() {
		id, _, err := s.Repository.CreateWallet()
		resultChan <- walletResult{ID: id, Err: err}
	}

	result := <-resultChan
	return result.ID, result.Err
}

func (s *Service) WalletInfo(id int) (int, float64, error) {
	walletID, balance, err := s.Repository.WalletInfo(id)
	return walletID, balance, err
}
