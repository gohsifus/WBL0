package store

import (
	"natTest/pkg/models"
	"testing"
)

func TestItemRepoCreate(t *testing.T) {
	s, tearDown := TestStore(t, dataBaseURL)
	defer tearDown("items")

	i, err := s.GetItemRepo().Create(&models.Item{
		ChrtId:      123,
		TrackNumber: "qweq",
		Price:       123,
		Rid:         "werwe",
		Name:        "qwe",
		Sale:        343,
		Size:        "qwe",
		TotalPrice:  232,
		NmId:        343,
		Brand:       "qwe",
		Status:      1,
	})

	if err != nil{
		t.Error(err)
	}

	if i == nil{
		t.Error("метод вернул нулевое значение")
	}
}

func TestItemRepoGetById(t *testing.T){
	s, tearDown := TestStore(t, dataBaseURL)
	defer tearDown("items")

	items, err := s.GetItemRepo().GetByOrderId(12)
	if items != nil || err != nil{
		t.Error(err)
	}
}
