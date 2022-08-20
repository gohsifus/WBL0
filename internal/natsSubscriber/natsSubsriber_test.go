package natsSubscriber

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"natTest/pkg/models"
	"testing"
	"time"
)

func TestNatsSubscriber_GetDataFromChannel(t *testing.T) {
	nats, err := New(NewConfig(), logrus.New())

	if err != nil{
		t.Skipf("ошибка подключения к nats: %s", err.Error())
	} else {
		defer nats.conn.Close()
	}

	testOrder := models.Order{
		Id:                1,
		OrderUid:          "qwe",
		TrackNumber:       "qwe",
		Entry:             "qwe",
		Delivery:          &models.Delivery{
			Id:      1,
			Name:    "qwe",
			Phone:   "qwe",
			Zip:     "qwe",
			City:    "qwe",
			Address: "qwe",
			Region:  "qwe",
			Email:   "qwe",
		},
		Payment:           &models.Payment{
			Id:           1,
			Transaction:  "qwe",
			RequestId:    "qwe",
			Currency:     "qwe",
			Provider:     "qwe",
			Amount:       123,
			PaymentDt:    123,
			Bank:         "qwe",
			DeliveryCost: 123,
			GoodsTotal:   123,
			CustomFee:    123,
		},
		Items:             []models.Item{
			{
				Id:          1,
				ChrtId:      123,
				TrackNumber: "qwe",
				Price:       123,
				Rid:         "qwe",
				Name:        "qwe",
				Sale:        123,
				Size:        "qwe",
				TotalPrice:  13,
				NmId:        123,
				Brand:       "qwe",
				Status:      2323,
				OrderId:     1,
			},
		},
		Locale:            "qwe",
		InternalSignature: "qwe",
		CustomerId:        "qwe",
		DeliveryService:   "qwe",
		ShardKey:          "qwe",
		SmId:              123,
		DateCreated:       time.Now(),
		OofShard:          "qwe",
	}

	testData, err := json.Marshal(testOrder)
	if err != nil{
		t.Skipf("ошибка преобразования тестовых данных: %s", err.Error())
	}

	err = nats.conn.Publish("testChannel", []byte("bad data"))
	if err != nil{
		t.Skipf("данные не были опубликованы по причине: %s", err.Error())
	}

	err = nats.conn.Publish("testChannel", testData)
	if err != nil{
		t.Skipf("данные не были опубликованы по причине: %s", err.Error())
	}

	got, err := nats.GetDataFromChannel("testChannel")
	if err != nil{
		t.Errorf("метод выполнился с ошибкой: %s", err.Error())
	}

	v := <-got
	if v.Id != testOrder.Id || v.TrackNumber != testOrder.TrackNumber{
		t.Errorf("получены неожиданные данные: ожидалось %v, получено %v", testOrder, v)
	}
}
