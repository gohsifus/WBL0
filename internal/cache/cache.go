package cache

import "natTest/pkg/models"

type Cache map[int]models.Order

func New() Cache {
	return make(map[int]models.Order)
}

//Restore запишет множество заявок в кеш
func (c *Cache) Restore(orders []models.Order) {
	for _, v := range orders {
		*c = c.AddToCache(v)
	}
}

//AddToCache Добавит заяку в кеш
func (c Cache) AddToCache(order models.Order) Cache {
	c[order.Id] = order
	return c
}

//GetById вернет заявку из кеша по id
func (c Cache) GetById(id int) models.Order {
	return c[id]
}
