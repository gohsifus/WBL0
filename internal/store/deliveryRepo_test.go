package store

import (
	_ "github.com/lib/pq"
	"natTest/pkg/models"
	"testing"
)

func TestDeliveryRepoCreate(t *testing.T) {
	s, tearDown := TestStore(t, dataBaseURL)
	defer tearDown("deliveries")

	d, err := s.GetDeliveryRepo().Create(&models.Delivery{
		Name: "aboba",
		Phone: "89892928928",
		Zip: "123123",
		City: "someCity",
		Address: "iAddress",
		Region: "31",
		Email: "some@some.com",
	})

	if d == nil || err != nil{
		t.Error(err)
	}
}

func TestDeliveryRepoGetById(t *testing.T){
	s, tearDown := TestStore(t, dataBaseURL)
	defer tearDown("deliveries")

	d, err := s.GetDeliveryRepo().GetById(12)
	if err != nil{
		t.Error(err)
	}

	if d != nil{
		t.Log(d)
		t.Error("запрос несуществующих данных вернул не нулевые данные")
	}

	d, err = s.GetDeliveryRepo().Create(&models.Delivery{
		Name: "aboba",
		Phone: "89892928928",
		Zip: "123123",
		City: "someCity",
		Address: "iAddress",
		Region: "31",
		Email: "some@some.com",
	})

	d, err = s.GetDeliveryRepo().GetById(d.Id)
	if err != nil{
		t.Error(err)
	}

	if d == nil{
		t.Log(d)
		t.Errorf("запрос записи с id = %v вернул нулевые данные", d.Id)
	}
}
