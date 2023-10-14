package main

import (
	"log"
	"os"

	"codeid.northwind/config"
	"codeid.northwind/server"
	_ "github.com/lib/pq"
)

func getConfigFileName() string {
	env := os.Getenv("ENV")

	if env != "" {
		return "northwind" + env
	}
	// == file northwind.toml
	return "northwind"
}

func main() {
	log.Println("Starting nortwind restapi")

	log.Println("Initializing configuration")
	config := config.InitConfig(getConfigFileName())

	log.Println("Initializing database")
	dbHandler := server.InitDatabase(config)
	// log.Println(dbHandler)

	log.Println("Initializing HTTP Server")
	httpServer := server.InitHttpServer(config, dbHandler) 

	httpServer.Start()

	//test insert to category
	/* ctx := context.Background()   // create goroutine
	queries := repositories.New(dbHandler)

	newCategory, err := queries.CreateCategory(ctx, 
		repositories.CreateCategoryParams{
			CategoryID: 101,
			CategoryName: "Test Category",
			Description: "Ini Barang",
			Picture: nil,
		},
	)

	if err != nil{
		log.Fatal("Eror : ", err)
	}
	log.Println(newCategory) */
}