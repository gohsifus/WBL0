package store

import (
	"fmt"
	"strings"
	"testing"
)

//Тестирование работы с бд производится не с помощью mock а с помощью тестовой базы данных
//это нужно для того чтобы можно было отловить ошибки на уровне запроса

//TestStore ...
func TestStore(t *testing.T, databaseUrl string) (*Store, func(...string)) {
	t.Helper()
	config := NewConfig()
	config.DatabaseUrl = databaseUrl
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	//Функция для очищения таблиц после тестирования
	return s, func(tables... string){
		if len(tables) > 0{
			if _, err := s.db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", "))); err != nil{
				t.Fatal(err)
			}
		}
		s.Close()
	}
}
