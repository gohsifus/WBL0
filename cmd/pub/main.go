//Для генерации и отправки данных в канал
package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"math/rand"
	"natTest/pkg/models"
	"os"
	"time"
)

//getOriginalData вернет данные json из задания
func getOriginalData() ([]byte, error) {
	file, err := os.Open("./cmd/pub/store/model.json")
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}

//getDataForPub Вернет данные для отправки в канал
//Данные могут быть 4 типов:
//	json с моделью из задания,
//	json c такой же структурой и случайными данными,
//	некорректные данные,
//  json частично повторяющий структуру
func getDataForPub() ([]byte, error) {
	n := rand.Intn(4)
	switch n {
	case 0:
		return json.Marshal(models.GetRandOrder())
	case 1:
		return getOriginalData()
	case 2:
		return []byte("i am bad data"), nil
	case 3:
		return []byte(`{"track_number": "qweqwe", "entry": "asdasd"}`), nil
	}

	return []byte{}, errors.New("неожиданное поведение")
}

//Публикация данных в канал
func main() {
	sc, err := stan.Connect("test-cluster", "client-pub")
	defer sc.Close()
	if err != nil {
		log.Fatalln(err)
	}

	for i := 1; ; i++ {
		data, err := getDataForPub()
		if err != nil {
			log.Fatalln(err)
		}

		err = sc.Publish("orders", data)
		if err != nil {
			fmt.Println("err " + err.Error())
		}
		fmt.Println(i)
		time.Sleep(time.Second * 3)
	}
}
