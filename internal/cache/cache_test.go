package cache

import (
	"natTest/pkg/models"
	"testing"
)

func TestCache_AddToCache(t *testing.T) {
	cache := New()
	dataForAdd := models.Order{
		Id:          1,
		TrackNumber: "qweqwe",
	}
	cache.AddToCache(dataForAdd)

	if len(cache) != 1 {
		t.Errorf("неожиданное количество элементов после добавления: ожидалось %d, получено %d", 1, len(cache))
	}
	if cache[1].TrackNumber != dataForAdd.TrackNumber || cache[1].Id != dataForAdd.Id {
		t.Errorf("неожиданные данные в кеше после добавления: ожидалось: %v, получено: %v", dataForAdd, cache[1])
	}
}

func TestCache_GetById(t *testing.T) {
	cache := New()
	dataForAdd := models.Order{
		Id:          1,
		TrackNumber: "qweqwe",
	}
	cache[1] = dataForAdd
	cache[2] = dataForAdd

	got := cache.GetById(1)
	if got.Id != dataForAdd.Id || got.TrackNumber != dataForAdd.TrackNumber{
		t.Errorf("метод вернул неожиданное значение: ожидалось %v, получено %v", dataForAdd, got)
	}
}

func TestCache_Restore(t *testing.T) {
	cache := New()
	dataForAdd := []models.Order{
		{Id: 1, TrackNumber: "qwe"},
		{Id: 2, TrackNumber: "asdasd"},
		{Id: 3, TrackNumber: "xcsdfs"},
	}

	cache.Restore(dataForAdd)

	if len(dataForAdd) != len(cache){
		t.Errorf("неожиданное количество элементов после восстановления: ожидалось %v, получано %v", len(dataForAdd), len(cache))
	}
}
