package store

import (
	"database/sql"
)

type Store struct{
	config *Config
	db *sql.DB
	deliveryRepo *DeliveryRepo
	paymentRepo *PaymentRepo
	orderRepo *OrderRepo
	itemRepo *ItemRepo
}

func New(config *Config) *Store{
	return &Store{
		config: config,
	}
}

//Open произведет подключение к бд и запишет подключение в хранилище
func (s *Store) Open() error{
	db, err := sql.Open("postgres", s.config.DatabaseUrl)
	if err != nil{
		return err
	}

	err = db.Ping()
	if err != nil{
		return err
	}


	s.db = db

	return nil
}

func (s *Store) Close(){
	s.db.Close()
}

func (s *Store) GetDeliveryRepo() *DeliveryRepo{
	if s.deliveryRepo != nil{
		return s.deliveryRepo
	}

	s.deliveryRepo = &DeliveryRepo{s}

	return s.deliveryRepo
}


