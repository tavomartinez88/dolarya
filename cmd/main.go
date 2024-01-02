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
		http.HandleFunc("/", healthCheckHandler) // Asigna el manejador a la ruta "/health"
		err := http.ListenAndServe(":8000", nil) // Escucha en el puerto 8000
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
