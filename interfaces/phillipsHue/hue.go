package phillipsHue

import (
    "github.com/collinux/gohue"
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

func ToggleLight(lightName string) {
	fmt.Println("Toggling light with name: ", lightName)
    //someLight, error := bridge.GetLightByName(lightName)

    // Get the last character which should contain the index of the light
    index, _ := strconv.Atoi( string(lightName[len(lightName)-1]) );
    someLight, error := bridge.GetLightByIndex(index)
    if error == nil {
    	someLight.Toggle()
    }
}