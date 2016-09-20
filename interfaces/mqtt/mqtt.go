/*
 * Copyright (c) 2013 IBM Corp.
 *
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v10.html
 *
 * Contributors:
 *    Seth Hoenig
 *    Allan Stockdill-Mander
 *    Mike Robertson
 */

package mqtt

import (
	"fmt"
	"log"
	"os"
	"time"
	"strings"
	"github.com/eclipse/paho.mqtt.golang"
	//"github.com/bclouser/gummies/config"
	"github.com/bclouser/gummies/interfaces/phillipsHue"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())

	// validate against configs?

	// add some stuff to database. or queue up those actions

	topicSplit := strings.Split(msg.Topic(), "/")
	if len(topicSplit) > 0 {
		fmt.Println("OK, we split our topic... its this", topicSplit)
		device := topicSplit[len(topicSplit)-1]
		phillipsHue.ToggleLight(device)
	}

}
func SyncDevices(){
	//cfg := config.GetConfig()
	//fmt.Println("OK, we got cfg: %s\n", cfg.Devices[0])
}

func MessageLoop() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://192.168.0.199:1883").SetClientID("home")
	//opts.SetUsername("maurice")
	//opts.SetPassword("saucy")
	opts.SetKeepAlive(20 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(20 * time.Second)

	c := mqtt.NewClient(opts)
	token := c.Connect();

	if token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := c.Subscribe("#", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	//time.Sleep(6 * time.Second)
	fmt.Println("About to token.Wait()")
	token.Wait()
	fmt.Println("Done token.Wait()ing. About to jump into this infinite loop")
	for ;; {
		time.Sleep(1000 * 1000 * 1000)
	}

	if token := c.Unsubscribe("#"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}


	c.Disconnect(250)

	time.Sleep(1 * time.Second)
}