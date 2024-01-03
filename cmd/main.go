package main

import (
	"dollar-bot/internal/utils"
	"dollar-bot/pkg"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("dolarya...")

	if os.Getenv("ENVIRONMENT") == "dev" {
		utils.LoadEnvironment()
	}

	go func() {
		http.HandleFunc("/", healthCheckHandler)
		err := http.ListenAndServe(":8000", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	h := pkg.NewHandler()
	h.HandleMessage()
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
