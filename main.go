package main


import (
	"time"
	"github.com/bclouser/gummies/interfaces/mqtt"
	"github.com/bclouser/gummies/interfaces/messageHandler"

	// Eventually, all devices will be initialized via a device init call on probe
	"github.com/bclouser/gummies/devices/philipsHue"
)

var configFile="./"


func main() {
	// Read in Config files

	// Parse config files
	//cfg := GetConfig(configFile)

	// Connect to database

	// Setup our MessageHandler
	messageHandler.Init()
	// Synchronize configs with database (also possibly use Redis)

	philipsHue.Init()

	// Maybe we should have a configuration section?
	// sync configurations with devices over mqtt
	mqtt.SyncDevices()

	// Kick off MQTT handler
	go mqtt.MessageLoop()
	
	// Kick off web server

	// Spin
	for ;; {
		time.Sleep(5* 1000 * 1000 * 1000)
	}
}