package main

import (
	"github.com/YuriRoberto/go-api/database"
	server "github.com/YuriRoberto/go-api/server"
)

func main() {
	//level := os.Getenv("LOG_LEVEL")
	//if level == "fatal" {
	//	log.SetLevel(log.FatalLevel)
	//}
	//if level == "debug" {
	//	log.SetLevel(log.DebugLevel)
	//}
	//log.Info("Starting API...")
	//
	//http.HandleFunc("/", index)
	//err := http.ListenAndServe(":8000", nil)
	//if err != nil {
	//	log.Fatalf("Problema ao subir o ListenAndServer: %s", err.Error())
	//}

	database.StartDB()

	server := server.NewServer()

	server.Run()
}

//func index(w http.ResponseWriter, r *http.Request) {
//
//}
