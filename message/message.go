package message

type Route struct {
	EndPoint string
}

type EndPoint struct {
	Building string
	Room string
	DeviceName string
}

type Message struct {
	RawEndP string
	EndP EndPoint
	Payload []byte

}