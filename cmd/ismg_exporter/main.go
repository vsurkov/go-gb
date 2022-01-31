package main

import (
	ismgExporter "github.com/vsurkov/ismg_exporter/internal/config"
	"log"
)

func main() {
	config := new(ismgExporter.Config)
	config.Load("")
	log.Printf("Config was loaded successfully \n%v", config)
}
