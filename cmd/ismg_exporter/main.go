package main

import (
	"fmt"
	ismgExporter "github.com/vsurkov/ismg_exporter/internal/ismg_exporter"
)

func main() {
	config := new(ismgExporter.Config)
	config.Load()

	fmt.Printf("Config was loaded:\n%v", config)
}
