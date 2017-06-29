package main

import (
	"fmt"
	"fog"
)

func main() {
	fmt.Println("STARTING...")
	var app = (&fog.TWebApp{}).Create()
	app.Run()
	fmt.Println("EXITING...")
}
