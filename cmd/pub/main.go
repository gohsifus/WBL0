//Для генерации и отправки данных в канал
package main

import (
	"encoding/json"
	"errors"
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
	file, err := os.Open("data.json")
	if err != nil {
		return []byte{}, err
	}
	defer file.Close()

	return ioutil.ReadAll(file)
}

//getDataForPub Вернет данные для отправки в канал
//Данные могут быть 3 типов: json с моделью из задания, json c такой же структурой и случайными данными и некорректные данные
func getDataForPub() ([]byte, error) {
	n := rand.Intn(3)
	switch n {
	case 0:
		return json.Marshal(models.GetRandOrder())
	case 1:
		return getOriginalData()
	case 2:
		return []byte("i am bad data"), nil
	}

	return []byte{}, errors.New("неожиданное поведение")
}

//Для публикации данных в канал
func main() {
	sc, err := stan.Connect("test-cluster", "client-pub")
	defer sc.Close()
	if err != nil {
		log.Fatalln(err)
	}

	for i := 1; ; i++ {
		data, err := getDataForPub()
		if err != nil{
			log.Fatalln(err)
		}

		sc.Publish("orders", data)
		time.Sleep(time.Second * 3)
	}
}
