package store

import "natTest/pkg/models"

type ItemRepo struct {
	store *Store
}

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
		m.OrderId,
	).Scan(&m.Id)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (i *ItemRepo) GetByOrderId(orderId int) ([]models.Item, error) {
	items := []models.Item{}
	item := &models.Item{}

	sql := "select * from items where order_id = $1"
	rows, err := i.store.db.Query(sql, orderId)
	if err != nil {
		return nil, err
	}

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

		if err != nil{
			return nil, err
		}
		items = append(items, *item)
	}

	return items, nil
}
