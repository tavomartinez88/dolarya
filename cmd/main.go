package main

import (
	"dollar-bot/pkg"
	"fmt"
)

func main() {
	fmt.Println("dolarya...")

	//utils.LoadEnvironment()

	h := pkg.NewHandler()
	h.HandleMessage()
}
