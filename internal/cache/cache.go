package cache

import "natTest/pkg/models"

type Cache map[int]models.Order

func New() Cache{
	return make(map[int]models.Order)
}

func AddToCache(c Cache, order models.Order) Cache {
	c[order.Id] = order
	return c
}

func (c Cache) GetById(id int) models.Order {
	return c[id]
}
