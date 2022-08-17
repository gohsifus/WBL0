package main

import (
	_ "github.com/lib/pq"
	"natTest/internal/service"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	//Загружаем конфигурации
	config := service.NewConfig()
	config.LoadConfig("./configs/Configs.json")

	//Создаем сервис
	app, err := service.New(config)
	if err != nil {
		panic(err)
	}

	//Запускаем сервис
	wg.Add(1)
	go func() {
		//Во время запуска также восстанавливается кэш
		err := app.Start()
		if err != nil {
			wg.Done()
			panic(err)
		}
	}()

	//Подписываемся на канал и получаем данные из nats
	data, _ := app.Nats.GetDataFromChannel("orders")
	for v := range data {
		//Запись в бд
		err := app.Store.PutOrderToStore(&v)
		if err != nil {
			app.Logger.Error(err)
		}

		//Запись в кэш
		app.Cache = app.Cache.AddToCache(v)
	}

	wg.Wait()
}
