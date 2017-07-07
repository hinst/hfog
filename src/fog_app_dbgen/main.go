package main

import "fog"

func main() {
	var app = (&fog.TDBGenApp{})
	app.BugsEnabled = true
	app.AttachmentsEnabled = true
	app.Run()
}
