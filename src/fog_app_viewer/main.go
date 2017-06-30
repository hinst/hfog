package main

import (
	"fmt"
	"fog"
)

func main() {
	var app = (&fog.TWebApp{}).Create()
	app.Run()
	fmt.Println("EXITING...")
}
