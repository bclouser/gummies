package devices

import (
	"github.com/bclouser/gummies/message"
	"github.com/bclouser/gummies/devices/philipsHue"
)


func EndPointKnown(endpoint string) bool {
	return true
}

func ActOnMessage(msg message.Message) {
	// currently we only have hue devices... so its hard coded right now
	philipsHue.ProcessMessage(msg)
}