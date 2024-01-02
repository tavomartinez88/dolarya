package main

import (
	"dollar-bot/pkg"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("dolarya...")

	//uncomment to work locally
	//utils.LoadEnvironment()

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
