package main

import (
	"dollar-bot/internal/utils"
	"dollar-bot/pkg"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("dolarya...")

	//uncomment to work locally
	utils.LoadEnvironment()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	h := pkg.NewHandler()
	h.HandleMessage()

	http.ListenAndServe(":8000", nil)
}
