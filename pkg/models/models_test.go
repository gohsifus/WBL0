package models

import (
	"testing"
	"time"
)

func TestGetSeed(t *testing.T) {
	//проверяем чтобы значение было каждый раз уникальное даже при одновременном вызове
	const numValues = 5
	seed := GetSeed()
	values := [numValues]int64{
		seed(),
		seed(),
		seed(),
		seed(),
		seed(),
	}
	got := make(map[int64]struct{}, numValues)

	for _, v := range values{
		got[v] = struct{}{}
	}

	if len(got) != numValues{
		t.Errorf("сгенерированно значений - %d, получено уникальных - %d: не все значения уникальны", numValues, len(got))
	}
}

func TestGetRandomString(t *testing.T) {
	lenStr := 5
	got := GetRandomString(5, time.Now().UnixNano())
	if len(got) != lenStr{
		t.Errorf("неожиданная длина сгенерированной строки: получено %d, ожидалось %d", len(got), lenStr)
	}
}

func TestGetRandDate(t *testing.T) {
	date1 := GetRandDate()
	time.Sleep(time.Millisecond * 1)
	date2 := GetRandDate()

	if date1 == date2{
		t.Errorf("метод возвращает не случайные значения: %v = %v", date1, date2)
	}
}

func TestGetRandDelivery(t *testing.T) {
	del1 := GetRandDelivery()
	time.Sleep(time.Millisecond * 1)
	del2 := GetRandDelivery()

	if del1 == del2{
		t.Errorf("\nметод возвращает не случайные значения:\n	%v = %v", del1, del2)
	}
}

func TestGetRandPayment(t *testing.T) {
	pay1 := GetRandPayment()
	time.Sleep(time.Millisecond * 1)
	pay2 := GetRandPayment()

	if pay1 == pay2{
		t.Errorf("\nметод возвращает не случайные значения:\n	%v = %v", pay1, pay2)
	}
}

func TestGetRandItem(t *testing.T) {
	item1 := GetRandItem()
	time.Sleep(time.Millisecond * 1)
	item2 := GetRandItem()

	if item1 == item2{
		t.Errorf("\nметод возвращает не случайные значения:\n	%v = %v", item1, item2)
	}
}
