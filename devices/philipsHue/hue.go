package philipsHue

import (
    "github.com/collinux/gohue"
    "github.com/bclouser/gummies/message"
    "fmt"
    "strconv"
    "encoding/json"
)

var bridge hue.Bridge
//var lights hue.Light

func Init() {
    // It is recommended that you save the username from bridge.CreateUser
    // so you don't have to press the link button every time and re-auth.
    // When CreateUser is called it will print the generated user token.
    bridgesOnNetwork, _ := hue.FindBridges()
    bridge = bridgesOnNetwork[0]
    //username, _ := bridge.CreateUser("maurice")
    //fmt.Println("Username is = ", username)
    bridge.Login("B3QEtDAWv7Y8FbOOGwwinntLHd6pDWyTQCmqE5Ua")
   	lights, _ := bridge.GetAllLights()
    for _, light := range lights {
        //light.On()
        //light.SetBrightness(100)
        //light.ColorLoop(true)
    	fmt.Println("light name: ", light.Name)
    }
}

func toggleLight(lightNum int) {
	fmt.Println("Toggling light number: ", lightNum)
    //someLight, error := bridge.GetLightByName(lightName)

    light, error := bridge.GetLightByIndex(lightNum)
    if error == nil {
        light.Toggle()
    }
}

func setLightBrightness(lightNum int, brightness int) {
    fmt.Println("Setting brightness of light: ", lightNum)

    light, error := bridge.GetLightByIndex(lightNum)
    if error == nil {
        if brightness >= 100 { 
            light.SetBrightness(100)
        } else if brightness >= 0 {
            light.SetBrightness(brightness)
        } else {
            fmt.Println("HUE: Bad brightness value specified")
        }
    }
}

func ProcessMessage(msg message.Message) {
    fmt.Println("HUE: Building ", msg.EndP.Building)
    fmt.Println("HUE: Room ", msg.EndP.Room)
    fmt.Println("HUE: DeviceName ", msg.EndP.DeviceName)

    var m message.HueMessage

    if "light" == msg.EndP.DeviceName[:len(msg.EndP.DeviceName)-1] {
        // Get the last character which should contain the index of the light
        lightNum, err := strconv.Atoi( msg.EndP.DeviceName[len(msg.EndP.DeviceName)-1:] )
        if err != nil {
            fmt.Println("Failed to parse number from light string")
            return
        }


        err = json.Unmarshal(msg.Payload, &m)
        if err != nil {
            fmt.Println("Failed to parse json message")
        }

        fmt.Println("Received command ", m.Command)
        switch m.Command {
            case "toggleCurrent":
                // toggle light without modifying brightness
                toggleLight(lightNum)

            case "toggleSet":
                // toggle light and set the brightness
                setLightBrightness(lightNum, m.Brightness)
                toggleLight(lightNum)
            default:
                fmt.Println("Unknown Command")
        }
    }
    //fmt.Fscanf("%s%d", &lightName, &lightNum)
    //fmt.Println("lightName ", lightName, " lightNum ", lightNum)
}