package store

import "natTest/pkg/models"

//PaymentRepo структура для работы с сущностью delivery из бд
type PaymentRepo struct {
	store *Store
}

//Create создаст запись в бд
func (p *PaymentRepo) Create(m *models.Payment) (*models.Payment, error) {
	sql := "insert into payments (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
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

	if err != nil {
		return nil, err
	}

	return m, nil
}

//GetById вернет запись по id
func (p *PaymentRepo) GetById(id int) (*models.Payment, error) {
	payment := &models.Payment{}

	sql := "select * from payments where id = $1"
	rows, err := p.store.db.Query(sql, id)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	count := 0
	for rows.Next() {
		err := rows.Scan(
			&payment.Id,
			&payment.Transaction,
			&payment.RequestId,
			&payment.Currency,
			&payment.Provider,
			&payment.Amount,
			&payment.PaymentDt,
			&payment.Bank,
			&payment.DeliveryCost,
			&payment.GoodsTotal,
			&payment.CustomFee,
		)

		count++
		if err != nil {
			return nil, err
		}
	}

	if count == 0{
		return nil, nil
	}

	return payment, nil
}
