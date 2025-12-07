package main

import (
	"fmt"
	"log"

	"github.com/thunkmetrc/automation/pkg/scraper"
)

func main() {
	fmt.Println("Starting Metrc Spec Generator...")
	if err := scraper.Run(); err != nil {
		log.Fatalf("Error: %v", err)
	}
	fmt.Println("Spec generation complete.")
}
