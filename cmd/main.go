package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/logger"
	"github.com/rs/cors"
	"projects.org/sample/sample-api/config"
	"projects.org/sample/sample-api/web"
)

const (
	//TimeOutSecond ...
	TimeOutSecond = 120
)

var (
	arqConfig string
	verbose   = flag.Bool("verbose", false, "print info level logs to stdout")
)

func init() {
	flag.StringVar(&arqConfig, "conf", "", "Config file in JSON format")
}

const logPath = "log.txt"
const configPath = "conf.json"

// Main function
func main() {
	flag.Parse()

	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()

	config := config.NewConfig(configPath)
	logInFile := config.EnableLogFile && config.LogFile != ""

	hWeb := &web.Handler{}
	if logInFile {
		logger := logger.Init("LoggerAPI", *verbose, false, lf)
		hWeb.Logger = logger
		defer logger.Close()
	}

	if logInFile {
		logger.Info("Using LOG in file")
	} else {
		log.Println("Without LOG in file")
	}

	router := web.Router(hWeb)

	allowedParam := make(map[string][]string)
	if err := json.Unmarshal([]byte(config.AllowedParam), &allowedParam); err != nil {
		log.Println("Error in json Unmarshal from allowedOrigins. Detail:", err)
		os.Exit(1)
	}
	c := cors.New(cors.Options{
		AllowedOrigins: allowedParam["Origins"],
		AllowedHeaders: allowedParam["Headers"],
		AllowedMethods: allowedParam["Methods"],
		Debug:          false,
	})

	s := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.Port),
		Handler:      c.Handler(router),
		ReadTimeout:  TimeOutSecond * time.Second,
		WriteTimeout: TimeOutSecond * time.Second,
	}

	if logInFile {
		logger.Info("Server starting in port: ", config.Port)
	} else {
		log.Println("Server running in port:", config.Port)
	}

	log.Println("... I'm online ...")

	if err := s.ListenAndServe(); err != nil {
		log.Println("Error in start server. Error:", err)
		os.Exit(1)
	}

}
