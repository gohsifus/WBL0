package store

import "natTest/pkg/models"

//Тут будут все методы для работы с бд с сущностью delivery
type DeliveryRepo struct{
	store *Store
}

func (d *DeliveryRepo) Create(m *models.Delivery) (*models.Delivery, error){
	sql := "insert into deliveries (name, phone, zip, city, address, region, email) values($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	err := d.store.db.QueryRow(sql, m.Name, m.Phone, m.Zip, m.City, m.Address, m.Region, m.Email).Scan(&m.Id)

	if err != nil{
		return nil, err
	}

	return m, nil
}