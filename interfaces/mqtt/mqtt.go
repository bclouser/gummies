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
	"github.com/eclipse/paho.mqtt.golang"
	"github.com/bclouser/gummies/message"
	"github.com/bclouser/gummies/interfaces/messageHandler"
)

var f mqtt.MessageHandler = func(client mqtt.Client, inMsg mqtt.Message) {
	topic := inMsg.Topic()
	payload := inMsg.Payload()
	fmt.Printf("TOPIC: %s\n", topic)
	fmt.Printf("MSG: %s\n", payload)

	// convert mqtt message to generic message type
	message := message.Message{RawEndP:topic, Payload:payload}

	messageHandler.NewMessageIn(message)

}
func SyncDevices(){
	//cfg := config.GetConfig()
	//fmt.Println("OK, we got cfg: %s\n", cfg.Devices[0])
}

func MessageLoop() {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions()
	opts.AddBroker("tcp://192.168.1.199:1883").SetClientID("home")
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