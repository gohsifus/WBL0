package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"natTest/internal/service"
)

func main(){
	config := service.NewConfig()
	config.LoadConfig("./configs/tmpConfig.json")
	fmt.Printf("%#+v\n", config)
}

/*func main() {
	//Загружаем конфиги
	config := apiserver.NewConfig()
	if err := config.LoadConfig("./configs/apiServer.json"); err != nil {
		log.Fatalln(fmt.Errorf("Ошибка инициализации конфигурации:  %s", err.Error()))
	}

	s := apiserver.New(config)

	nats, err := natsSubscriber.New(config.Nats)
	if err != nil {
		s.Logger.Error(fmt.Errorf("Ошибка подключения к NATS:  %s", err.Error()))
	}

	nats.GetDataFromChannel()

	//Запуск сервера

	if err := s.Start(); err != nil {
		s.Logger.Error("Ошибка запуска сервера: " + err.Error())
	}

	//Надо считывание с натс и запуск сервера параллельно потом реализовать запись в кэш и в бд

	/*sc, err := stan.Connect("test-cluster", "client-sub")
	if err != nil{
		panic(err)
	}
	defer sc.Close()

	_, err = sc.Subscribe("orders", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil{
		panic(err)
	}*//*
}*/
