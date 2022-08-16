package models

import (
	"math/rand"
	"strings"
	"time"
)

//GetSeed Вернет случайное не повторяющееся число int64
func GetSeed() func() int64 {
	//Замыкание чтобы можно было вызывать одновременно и получать разные значения
	seed := time.Now().UnixNano()
	return func() int64 {
		seed += 1
		return seed
	}
}

//GetRandDelivery Вернет структуру типа delivery инициализированную случайными значениями
func GetRandDelivery() *Delivery {
	seed := GetSeed()
	return &Delivery{
		Name:    GetRandomString(6, seed()),
		Phone:   GetRandomString(6, seed()),
		Zip:     GetRandomString(6, seed()),
		City:    GetRandomString(6, seed()),
		Address: GetRandomString(6, seed()),
		Region:  GetRandomString(6, seed()),
		Email:   GetRandomString(6, seed()),
	}
}

//GetRandPayment Вернет структуру типа payment инициализированную случайными значениями
func GetRandPayment() *Payment {
	seed := GetSeed()
	return &Payment{
		Transaction:  GetRandomString(6, seed()),
		RequestId:    GetRandomString(6, seed()),
		Currency:     GetRandomString(6, seed()),
		Provider:     GetRandomString(6, seed()),
		Amount:       rand.Intn(1_000_000),
		PaymentDt:    rand.Intn(1_000_000),
		Bank:         GetRandomString(6, seed()),
		DeliveryCost: rand.Intn(1_000_000),
		GoodsTotal:   rand.Intn(1_000_000),
		CustomFee:    rand.Intn(1_000_000),
	}
}

//GetRandItem Вернет структуру типа item инициализированную случайными значениями
func GetRandItem() Item {
	seed := GetSeed()
	return Item{
		ChrtId:      rand.Intn(1_000_000),
		TrackNumber: GetRandomString(6, seed()),
		Price:       rand.Intn(1_000_000),
		Rid:         GetRandomString(6, seed()),
		Name:        GetRandomString(6, seed()),
		Sale:        rand.Intn(1_000_000),
		Size:        GetRandomString(6, seed()),
		TotalPrice:  rand.Intn(1_000_000),
		NmId:        rand.Intn(1_000_000),
		Brand:       GetRandomString(6, seed()),
		Status:      rand.Intn(1_000_000),
	}
}

//GetRandOrder Вернет структуру типа order инициализированную случайными значениями
func GetRandOrder() Order {
	seed := GetSeed()
	return Order{
		OrderUid:          GetRandomString(6, seed()),
		TrackNumber:       GetRandomString(6, seed()),
		Entry:             GetRandomString(6, seed()),
		Delivery:          GetRandDelivery(),
		Payment:           GetRandPayment(),
		Items:             []Item{GetRandItem()},
		Locale:            GetRandomString(6, seed()),
		InternalSignature: GetRandomString(6, seed()),
		CustomerId:        GetRandomString(6, seed()),
		DeliveryService:   GetRandomString(6, seed()),
		ShardKey:          GetRandomString(6, seed()),
		SmId:              rand.Intn(1_000_000),
		DateCreated:       GetRandDate(),
		OofShard:          GetRandomString(6, seed()),
	}
}

//GetRandomString Вернет случйную строку заданной длины
func GetRandomString(len int, seed int64) string {
	rand.Seed(seed)

	resultStr := strings.Builder{}
	for i := 0; i < len; i++ {
		char := 40 + rand.Intn(79)
		resultStr.WriteRune(rune(char))
	}
	return resultStr.String()
}

//GetRandDate вернет случайную дату
func GetRandDate() time.Time {
	rand.Seed(time.Now().UnixNano())
	m := rand.Intn(120)
	return time.Date(2000, time.Month(m), 0, 0, 0, 0, 0, time.UTC)
}
