package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/l-gu/goproject/internal/app/config"
	"github.com/l-gu/goproject/internal/app/logger"
)

func staticFilesHandler(w http.ResponseWriter, r *http.Request) {
	fileName := config.GetWebDir() + r.URL.Path
	log("Static file : URL path '" + r.URL.Path + "' --> file '" + fileName + "'")
	http.ServeFile(w, r, fileName)
}

func log(v ...interface{}) {
	logger.Log("main", v...) // v... : treat input to function as variadic
}

func main() {

	log("Starting server ... ")

	printEnv()

	log("Init configuration... ")
	err := config.InitConfig("config.json")
	if err != nil {
		log("ERROR", err)
	}

	initPersistence(config.GetDaoType())

	log("Setting static files handler... ")
	// Set handler to serve static files
	http.HandleFunc("/", staticFilesHandler)

	log("Setting application handlers... ")
	initWebMVC()

	// Launch the http server
	webPort := ":" + config.GetWebPort()
	log("Launching http server (port=" + webPort + ") ... ")
	//log.Fatal(http.ListenAndServe(webPort, nil))
	err = http.ListenAndServe(webPort, nil)
	if err != nil {
		log("ERROR", err)
	}
}

func printEnv() {
	// Current dir
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log("ERROR", err)
	}
	log("current dir :", dir)

	dir, err = filepath.Abs("./")
	if err != nil {
		log("ERROR", err)
	}
	log("file path   :", dir)
}
