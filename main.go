package main


import (
	"time"
	"github.com/bclouser/gummies/interfaces/device/mqtt"
)

func main() {
	// Read in Config files

	// Parse config files

	// Connect to database

	// Synchronize configs with database (also possibly use Redis)

	// Kick off MQTT handler
	go mqtt.MessageLoop()
	
	// Kick off web server

	// Spin
	for ;; {
		time.Sleep(5* 1000 * 1000 * 1000)
	}
}