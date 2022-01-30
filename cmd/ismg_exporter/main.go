package main

import (
	ismgExporter "github.com/vsurkov/ismg_exporter/internal/ismg_exporter"
	"log"
)

func main() {
	config := new(ismgExporter.Config)
	config.Load()
	log.Printf("Config was loaded successfully \n%v", config)
}
