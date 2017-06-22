package main

import (
	"fmt"
	"fog"
)

func main() {
	fmt.Println("STARTING...")
	var app = (&fog.TApp{}).Create()
	app.Run()
}