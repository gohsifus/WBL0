package store

import (
	"database/sql"
	"fmt"
	"natTest/pkg/models"
)

//Store структура для работы с бд
type Store struct {
	config       *Config
	db           *sql.DB
	deliveryRepo *DeliveryRepo
	paymentRepo  *PaymentRepo
	orderRepo    *OrderRepo
	itemRepo     *ItemRepo
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

//Open произведет подключение к бд и запишет подключение в хранилище
func (s *Store) Open() error {
	db, err := sql.Open("postgres", s.config.DatabaseUrl)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

//Close закроет подключение к бд
func (s *Store) Close() {
	s.db.Close()
}

func (s *Store) GetDeliveryRepo() *DeliveryRepo {
	if s.deliveryRepo != nil {
		return s.deliveryRepo
	}

	s.deliveryRepo = &DeliveryRepo{s}

	return s.deliveryRepo
}

func (s *Store) GetOrderRepo() *OrderRepo {
	if s.orderRepo != nil {
		return s.orderRepo
	}

	s.orderRepo = &OrderRepo{s}

	return s.orderRepo
}

func (s *Store) GetPaymentRepo() *PaymentRepo {
	if s.paymentRepo != nil {
		return s.paymentRepo
	}

	s.paymentRepo = &PaymentRepo{s}

	return s.paymentRepo
}

func (s *Store) GetItemRepo() *ItemRepo {
	if s.itemRepo != nil {
		return s.itemRepo
	}

	s.itemRepo = &ItemRepo{s}

	return s.itemRepo
}

//PutOrderToStore запишет структуру order в базу
func (s *Store) PutOrderToStore(m *models.Order) error {
	_, err := s.GetDeliveryRepo().Create(m.Delivery)
	if err != nil {
		return fmt.Errorf("ошибка записи заявки в бд: %s", err)
	}
	_, err = s.GetPaymentRepo().Create(m.Payment)
	if err != nil {
		return fmt.Errorf("ошибка записи заявки в бд: %s", err)
	}

	_, err = s.GetOrderRepo().Create(m)
	if err != nil {
		return fmt.Errorf("ошибка записи заявки в бд: %s", err)
	}

	for _, item := range m.Items {
		item.OrderId = m.Id
		_, err = s.GetItemRepo().Create(&item)
		if err != nil {
			return fmt.Errorf("ошибка записи заявки в бд: %s", err)
		}
	}

	return nil
}
