package config


import (
	"fmt"
	//"io/ioutil"
	//"encoding/json"
)


type Config struct {
	Devices []string
}
var singleConfig *Config = nil;

func (Config) GimmeAllDevices() {

}

func InitConfig(fileName string){
	fmt.Println("Initializing config with %s", fileName)
}

func GetConfig() *Config {
	if singleConfig == nil{
		singleConfig = new(Config)
	}
	return singleConfig
}


func TestFunc() {
	fmt.Println	("some how this happened......not sure how")
}
