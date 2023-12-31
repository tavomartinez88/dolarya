package main

import (
	"dollar-bot/pkg"
	"fmt"
)

func main() {
	fmt.Println("dolarya...")

	//uncomment to work locally
	//utils.LoadEnvironment()

	h := pkg.NewHandler()
	h.HandleMessage()
}
