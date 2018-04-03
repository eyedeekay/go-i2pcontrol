package main

import (
    "flag"
)

import "github.com/eyedeekay/go-i2pcontrol"

func main(){
    AddrString := flag.String("host", "127.0.0.1",
		"host: of the i2pcontrol interface")
	PortString := flag.String("port", "7650",
		":port of the i2pcontrol interface")

    flag.Parse()

    i2pcontrolhost := *AddrString
    i2pcontrolport := *PortString

    auth := i2pcontrol.NewI2pControl(i2pcontrolhost, i2pcontrolport)

    auth.Echo("test")

}
