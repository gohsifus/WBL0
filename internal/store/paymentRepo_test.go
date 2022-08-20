package store

import (
	"natTest/pkg/models"
	"testing"
)

func TestPaymentRepoCreate(t *testing.T) {
	s, tearDown := TestStore(t, dataBaseURL)
	defer tearDown("payments")

	p, err := s.GetPaymentRepo().Create(&models.Payment{
		Transaction:  "qwe",
		RequestId:    "qwe",
		Currency:     "qwe",
		Provider:     "qwe",
		Amount:       0,
		PaymentDt:    0,
		Bank:         "qwe",
		DeliveryCost: 0,
		GoodsTotal:   0,
		CustomFee:    0,
	})

	if p == nil || err != nil {
		t.Error(err)
	}
}

func TestPaymentRepoGetById(t *testing.T){
	s, tearDown := TestStore(t, dataBaseURL)
	defer tearDown("payments")

	p, err := s.GetPaymentRepo().Create(&models.Payment{
		Transaction:  "qwe",
		RequestId:    "qwe",
		Currency:     "qwe",
		Provider:     "qwe",
		Amount:       123,
		PaymentDt:    123,
		Bank:         "qwe",
		DeliveryCost: 234,
		GoodsTotal:   234,
		CustomFee:    0,
	})

	if p == nil || err != nil{
		t.Error(err)
	}

	p, err = s.GetPaymentRepo().GetById(p.Id)

	if p == nil || err != nil{
		t.Error(err)
	}
}
