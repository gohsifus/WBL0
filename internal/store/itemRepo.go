package store

import (
	sql2 "database/sql"
	"natTest/pkg/models"
)

//ItemRepo структура для работы с сущностью item из бд
type ItemRepo struct {
	store *Store
}

func getOrderIdValue(orderId int) interface{}{
	if orderId == 0{
		return sql2.NullInt32{}
	}
	return orderId
}

//Create создаст запись в бд
func (i *ItemRepo) Create(m *models.Item) (*models.Item, error) {
	sql := "insert into items (chrt_id, track_number, price, rid, name, sale , size, total_price, nm_id, brand, status, order_id) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) RETURNING id"

	err := i.store.db.QueryRow(
		sql,
		m.ChrtId,
		m.TrackNumber,
		m.Price,
		m.Rid,
		m.Name,
		m.Sale,
		m.Size,
		m.TotalPrice,
		m.NmId,
		m.Brand,
		m.Status,
		getOrderIdValue(m.OrderId),
	).Scan(&m.Id)

	if err != nil {
		return nil, err
	}

	return m, nil
}

//GetByOrderId вернет все записи по id заявки
func (i *ItemRepo) GetByOrderId(orderId int) ([]models.Item, error) {
	items := []models.Item{}
	item := &models.Item{}

	sql := "select * from items where order_id = $1"
	rows, err := i.store.db.Query(sql, orderId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	count := 0
	for rows.Next() {
		err = rows.Scan(
			&item.Id,
			&item.ChrtId,
			&item.TrackNumber,
			&item.Price,
			&item.Rid,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmId,
			&item.Brand,
			&item.Status,
			&item.OrderId,
		)

		count++

		if err != nil {
			return nil, err
		}
		items = append(items, *item)
	}

	if count == 0 {
		return nil, nil
	}

	return items, nil
}
