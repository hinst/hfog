package main

import "fog"

func main() {
	var app = (&fog.TWebApp{}).Create()
	app.Run()
}
