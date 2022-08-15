package store

import "natTest/pkg/models"

//Тут методы для работы с бд сущности payment

type PaymentRepo struct{
	store *Store
}

func (p PaymentRepo) Create(m *models.Payment) (*models.Payment, error){
	sql := "insert into payments (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	err := p.store.db.QueryRow(
		sql,
		m.Transaction,
		m.RequestId,
		m.Currency,
		m.Provider,
		m.Amount,
		m.PaymentDt,
		m.Bank,
		m.DeliveryCost,
		m.GoodsTotal,
		m.CustomFee,
		).Scan(&m.Id)

	if err != nil{
		return nil, err
	}

	return m, nil
}