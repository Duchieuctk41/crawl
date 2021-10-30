package route

import (
	"fmt"
	"log"

	"net/http"

	"crawl/conf"
	"crawl/pkg/driver"
	"crawl/pkg/handler"
	"crawl/pkg/repo"
	srv "crawl/pkg/service"

	"github.com/gorilla/mux"
)

func Start() {
	config, err := conf.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	db := driver.ConnectMongo()

	newRepo := repo.NewRepo(db.Client.Database(config.DBName))
	crawlService := srv.NewService(newRepo)
	crawlHandler := handler.NewHandler(crawlService)

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/crawl/collect", crawlHandler.CrawlCollection).Methods("GET")

	fmt.Println("listening in localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", r))
}
