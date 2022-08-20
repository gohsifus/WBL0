package store

import (
	"natTest/pkg/models"
	"testing"
	"time"
)

func TestOrderRepoCreate(t *testing.T) {
	s, tearDown := TestStore(t, dataBaseURL)
	defer tearDown("orders")

	o, err := s.GetOrderRepo().Create(
		&models.Order{
			OrderUid:          "qwe",
			TrackNumber:       "qwe",
			Entry:             "qwe",
			Delivery:          nil,
			Payment:           nil,
			Items:             nil,
			Locale:            "qweq",
			InternalSignature: "qweq",
			CustomerId:        "qwe",
			DeliveryService:   "qwe",
			ShardKey:          "qwe",
			SmId:              0,
			DateCreated:       time.Time{},
			OofShard:          "qwe",
		},
	)

	if o == nil || err != nil{
		t.Error(err)
	}
}

func TestOrderRepoGetList(t *testing.T){
	s, tearDown := TestStore(t, dataBaseURL)
	defer tearDown("orders")

	o, err := s.GetOrderRepo().GetList()

	if o != nil || err != nil{
		t.Error(err)
	}
}
