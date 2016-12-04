package messageHandler

import (
	"fmt"
	"strings"
	"github.com/bclouser/gummies/interfaces/database"
	"github.com/bclouser/gummies/interfaces/devices"
	"github.com/bclouser/gummies/message"
)


/*
The interface for anything involving messages... This includes routing (for now)
*/

//func GetRoute(topic string) (*Route, error) {
//	return new(Route), nil
//}

func IsEndpointValid(endpoint string) bool{
	return true
}

func IsPayloadValid(payload []byte) bool {
	return true
}

func IncomingMessageHandler(topic string, payload string) {

}

// TODO, actually define this Message struct
func NewMessageIn(msg message.Message) {

	// validate against configs?
	if !IsEndpointValid(msg.RawEndP) {
		fmt.Println("Failed to find registered endpoint ", msg.RawEndP)
		return
	}

	if !IsPayloadValid(msg.Payload) {
		fmt.Println("msg payload invalid...")
		return
	}

	// Pass msg along to device (HarborMaster)
		// This module will check if any code has been registered for this endpoint
		// Obviously call it if something has been registered


		// Update database. It would be cool if both of these were goRoutines, but the database
	    	// update might be dependent on the response from the harborMaster/device

	// add some stuff to database. or queue up those actions
	topicSplit := strings.Split(msg.RawEndP, "/")
	if len(topicSplit) > 0 {

		// first item in topic should be the building name?
		//route := router.GetRoute(msg.Topic())

		// second item in topic should be the room
		// get all items for this specific room

		// second item in the topic should be the device name (unique by room)
		// getRoute(building, room, deviceName)

		// Do we need to do anything? Most of the time we will just be
		// updating the database... possibly all the time
		// if database.hasActions(building, room, deviceName) {
		//	messageHandler.sendOutAction()... something like this
		//}
		fmt.Println("OK, we split our topic... its this", topicSplit)
		if len(topicSplit) == 3 {
			// Topics: building, room, deviceName
			err := database.Update(topicSplit[0], topicSplit[1], topicSplit[2])
			if err != nil {
				fmt.Println("Failed to update database. Error: ", err)
			}
			
			msg.EndP.Building = topicSplit[0]
			msg.EndP.Room = topicSplit[1]
			msg.EndP.DeviceName = topicSplit[2]

			// Do we need to send out a msg? we should ask our registered devices
			// if any of them care about the received endpoint
			if devices.EndPointKnown(msg.RawEndP) {
				// Ok, so someone cares about this. Pass along the msg
				devices.ActOnMessage(msg)
			}


		} else {
			fmt.Println("length of topic string is not 3")
		}

	} else {
		fmt.Println("Bad topic string")
	}
}

func Init() error {
	/* 
	The plan:
		- Call into mqtt and register the "messageHandler function"
			-- this function will get called whenever mqtt receives a msg
			-- it will do all validation/routing/queueing

	*/
	return nil
}