package store

import "natTest/pkg/models"

type ItemRepo struct{
	store *Store
}

func (i ItemRepo) Create(m *models.Item) (*models.Item, error){
	sql := "insert into items (chrt_id, track_number, price, rid, name, sale , size, total_price, nm_id, brand, status, order_id) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
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

	if err != nil{
		return nil, err
	}

	return m, nil
}
