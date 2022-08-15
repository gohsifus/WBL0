package store

import "natTest/pkg/models"

//Тут методы для работы с бд сущности order

type OrderRepo struct{
	store *Store
}

func (o *OrderRepo) Create(m *models.Order) (*models.Order, error){
	sql := "insert into orders (order_uid, track_number, entry, delivery_id, payment_id, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) values($1, $2, $3, $4, $5, %6, $7, $8, $9, $10, $11, $12, $13)"
	err := o.store.db.QueryRow(sql,
		m.OrderUid,
		m.TrackNumber,
		m.Entry,
		m.Delivery.Id,
		m.Payment.Id,
		m.Locale,
		m.InternalSignature,
		m.CustomerId,
		m.DeliveryService,
		m.ShardKey,
		m.SmId,
		m.DateCreated,
		m.OofShard,
	).Scan(&m.Id)

	if err != nil{
		return nil, err
	}

	return m, nil
}