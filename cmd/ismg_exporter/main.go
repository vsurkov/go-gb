package main

import (
	ismgExporter "github.com/vsurkov/ismg_exporter/internal/ismg_exporter"
	"log"
)

func main() {
	config := new(ismgExporter.Config)
	err := config.LoadFromConfig("..//..//configs//ismg_exporter.json")
	if err != nil {
		log.Printf("Error on loadign from File %v", err)
		config.LoadFromParams()
	}
	log.Printf("Config was loaded:\n%v", config)
}
