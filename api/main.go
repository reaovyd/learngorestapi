package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/reaovyd/learngorestapi/api/router"
)

type serverProperties struct {
	DOMAIN string `json:"DOMAIN"`
	PORT   string `json:"PORT"`
}

type configFileMap struct {
	SERVER_PROPERTIES serverProperties `json:"SERVER_PROPERTIES"`
}

func readConfigFile(fileName *string) (configFileMap, error) {
	data, err := os.ReadFile(*fileName)
	mappedData := configFileMap{}
	if err != nil {
		return mappedData, err
	}
	err = json.Unmarshal([]byte(data), &mappedData)
	return mappedData, err
}

func main() {
	CONFIG_FILE := os.Getenv("CONFIG")
	mappedData, err := readConfigFile(&CONFIG_FILE)

	if err != nil {
		log.Fatal(err)
		log.Fatal("File was unable to be read")
		os.Exit(1)
	}
	server := mappedData.SERVER_PROPERTIES.DOMAIN + ":" + mappedData.SERVER_PROPERTIES.PORT

	r := gin.Default()
	router.SetupRoutes(r)

	r.Run(server)
}
