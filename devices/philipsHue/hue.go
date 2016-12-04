package philipsHue

import (
    "github.com/collinux/gohue"
    "github.com/bclouser/gummies/message"
    "fmt"
    "strconv"
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
	fmt.Println("Toggling light numer: ", lightNum)
    //someLight, error := bridge.GetLightByName(lightName)

    someLight, error := bridge.GetLightByIndex(lightNum)
    if error == nil {
    	someLight.Toggle()
    }
}

func ProcessMessage(msg message.Message) {
    fmt.Println("HUE: Building ", msg.EndP.Building)
    fmt.Println("HUE: Room ", msg.EndP.Room)
    fmt.Println("HUE: DeviceName ", msg.EndP.DeviceName)

    if "light" == msg.EndP.DeviceName[:len(msg.EndP.DeviceName)-1] {
        // Get the last character which should contain the index of the light
        lightNum, err := strconv.Atoi( msg.EndP.DeviceName[len(msg.EndP.DeviceName)-1:] )
        if err != nil {
            fmt.Println("Failed to parse number from light string")
            return
        }
        toggleLight(lightNum)
    }
    //fmt.Fscanf("%s%d", &lightName, &lightNum)
    //fmt.Println("lightName ", lightName, " lightNum ", lightNum)
}