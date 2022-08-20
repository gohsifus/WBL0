package store

import (
	"database/sql"
	"natTest/pkg/models"
)

//OrderRepo структура для работы с сущностью order из бд
type OrderRepo struct {
	store *Store
}

//Вспомогательная функция возвращающая null для записи в бд если передается nil
func getDeliveryId(d *models.Delivery) interface{}{
	if d == nil{
		return sql.NullInt32{}
	}
	return d.Id
}

func getPaymentId(p *models.Payment) interface{}{
	if p == nil{
		return sql.NullInt32{}
	}
	return p.Id
}

//Create создаст запись в бд
func (o *OrderRepo) Create(m *models.Order) (*models.Order, error) {
	sql := `insert into orders (
				order_uid, 
				track_number, 
				entry, 
				delivery_id, 
				payment_id, 
				locale, 
				internal_signature, 
				customer_id, 
				delivery_service, 
				shardkey, 
				sm_id, 
				date_created, 
				oof_shard
			) values (
				$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
			) RETURNING id`
	err := o.store.db.QueryRow(
		sql,
		m.OrderUid,
		m.TrackNumber,
		m.Entry,
		getDeliveryId(m.Delivery),
		getPaymentId(m.Payment),
		m.Locale,
		m.InternalSignature,
		m.CustomerId,
		m.DeliveryService,
		m.ShardKey,
		m.SmId,
		m.DateCreated,
		m.OofShard,
	).Scan(&m.Id)

	if err != nil {
		return nil, err
	}

	return m, nil
}

//GetList вернет все заявки из бд
func (o *OrderRepo) GetList() ([]models.Order, error) {
	orders := []models.Order{}
	order := models.Order{
		Payment:  &models.Payment{},
		Delivery: &models.Delivery{},
	}

	sql := "select * from orders"

	rows, err := o.store.db.Query(sql)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	count := 0
	for rows.Next() {
		err := rows.Scan(
			&order.Id,
			&order.OrderUid,
			&order.TrackNumber,
			&order.Entry,
			&order.Payment.Id,
			&order.Delivery.Id,
			&order.Locale,
			&order.InternalSignature,
			&order.CustomerId,
			&order.DeliveryService,
			&order.ShardKey,
			&order.SmId,
			&order.DateCreated,
			&order.OofShard,
		)

		count++

		if err != nil {
			return nil, err
		}

		//Получаем платежи
		if order.Payment, err = o.store.GetPaymentRepo().GetById(order.Payment.Id); err != nil {
			return nil, err
		}
		//Получаем доставку
		if order.Delivery, err = o.store.GetDeliveryRepo().GetById(order.Delivery.Id); err != nil {
			return nil, err
		}
		//Получаем позиции(items)
		if order.Items, err = o.store.GetItemRepo().GetByOrderId(order.Id); err != nil {
			return nil, err
		}

		orders = append(orders, order)
	}

	if count == 0{
		return nil, nil
	}

	return orders, nil
}
