package main

import (

)

func main(){
	ping := Probe{1, "testping", "209.56.104.2", true}
	ping.IPAddressConstructor()
	ping.Pinger()
}
