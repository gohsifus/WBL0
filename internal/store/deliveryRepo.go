package store

import "natTest/pkg/models"

//DeliveryRepo структура для работы с сущностью delivery из бд
type DeliveryRepo struct {
	store *Store
}

//Create создаст запись в бд
func (d *DeliveryRepo) Create(m *models.Delivery) (*models.Delivery, error) {
	sql := "insert into deliveries (name, phone, zip, city, address, region, email) values($1, $2, $3, $4, $5, $6, $7) RETURNING id"
	err := d.store.db.QueryRow(sql, m.Name, m.Phone, m.Zip, m.City, m.Address, m.Region, m.Email).Scan(&m.Id)

	if err != nil {
		return nil, err
	}

	return m, nil
}

//GetById вернет запись по id
func (d *DeliveryRepo) GetById(id int) (*models.Delivery, error) {
	delivery := &models.Delivery{}

	sql := "select * from deliveries where id = $1"
	rows, err := d.store.db.Query(sql, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&delivery.Id,
			&delivery.Name,
			&delivery.Phone,
			&delivery.Zip,
			&delivery.City,
			&delivery.Address,
			&delivery.Region,
			&delivery.Email,
		)

		if err != nil {
			return nil, err
		}
	}

	return delivery, nil
}
