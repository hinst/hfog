package main

import "fog"
import "runtime/debug"

func main() {
	debug.SetGCPercent(10)
	var app = (&fog.TDBGenApp{})
	app.BugsEnabled = true
	app.AttachmentsEnabled = true
	app.Run()
}
