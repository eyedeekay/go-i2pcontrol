package main

import (
    "flag"
)

import ".."

func main(){
    AddrString := flag.String("host", "127.0.0.1",
		"host: of the i2pcontrol interface")
	PortString := flag.String("port", "7650",
		":port of the i2pcontrol interface")
    ApiVersion := flag.Int64("api", 1,
		"version of the i2pcontrol interface")
    Password := flag.String("pass", "default",
		"password for authenticating to the i2pcontrol interface")
    flag.Parse()

    apiversion := *ApiVersion
    i2pcontrolhost := *AddrString
    i2pcontrolport := *PortString
    password := *Password

    i2pcontrol.I2pControlVerboseLogging = true

    auth := i2pcontrol.NewI2pControl(apiversion, password, i2pcontrolhost, i2pcontrolport)

    auth.Echo("i2pcontrol")

}
