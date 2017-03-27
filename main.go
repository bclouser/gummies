package main


import (
	"fmt"
	"time"
	"github.com/bclouser/gummies/interfaces/mqtt"
	"github.com/bclouser/gummies/interfaces/messageHandler"
    zmq "github.com/pebbe/zmq4"

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

	// Kick off ZMQ message handler
	go zmqServer()
	
	// Kick off web server

	// Spin
	for ;; {
		time.Sleep(5* 1000 * 1000 * 1000)
	}
}

func zmqServer() {
    context, _ := zmq.NewContext()
    socket, _ := context.NewSocket(zmq.PULL)
    //defer context.Close()
    defer socket.Close()
    println("ZMQ SERVER: Binding...")
    err := socket.Connect("tcp://192.168.1.134:5555")
    if err != nil {
    	println("ZMQ SERVER: Failed to bind socket")
    	fmt.Println(err)
    	return
    }
    println("ZMQ SERVER: Bound. Listening for push events")

    // Wait for messages
    for {
    	println("ZMQ SERVER: Inside loop, about to connect socket")
        msg, _ := socket.Recv(0)
        println("ZMQ SERVER: Received ", string(msg))

        // do some fake "work"
        //time.Sleep(time.Second)

        // send reply back to client
        reply := fmt.Sprintf("World")
        socket.Send(reply, 0)
    }
}
