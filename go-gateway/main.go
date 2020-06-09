package main

import (
	"go-dom-parser/api/routes"
	"go-dom-parser/api/sockets"
	"go-dom-parser/configs"
	"go-dom-parser/core"
	"strconv"
)

var err error

func main() {

	// load configuration
	cfg := configs.SetupConf()

	// Creating a connection to the database
	// configs.DB, err = gorm.Open("mysql", configs.DbURL(configs.BuildDBConfig(cfg)))

	// if err != nil {
	// 	fmt.Println("status: ", err)
	// }

	// defer configs.DB.Close()

	// run the migrations: todo struct
	// configs.DB.AutoMigrate(&models.Todo{})

	// === Processor ===

	p := core.New()
	p.Run()

	// === Processor ===

	// === RMQ configuration ===

	data := []byte("hello world --- xxx")

	ch := sockets.SetupRMQ(cfg)

	ch.AddProcessor("test", p.ProcessorChan)

	ch.Publish(cfg, data)

	ch.Subscribe(cfg)

	// === RMQ configuration ===

	// define routes
	router := routes.SetupRouter()

	// run server
	router.Run(":" + strconv.Itoa(cfg.Host.Port))
}
