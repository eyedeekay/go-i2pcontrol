package main

import (
	"flag"
)

import ".."

func main() {
	AddrString := flag.String("host", "127.0.0.1",
		"host: of the i2pcontrol interface")
	PortString := flag.String("port", "7650",
		":port of the i2pcontrol interface")
	ApiVersion := flag.Int("api", 1,
		"version of the i2pcontrol interface")
	Password := flag.String("pass", "default",
		"password for authenticating to the i2pcontrol interface")
    CertPath := flag.String("cert", "/var/lib/i2pd/i2pcontrol.cert.pem",
		"password for authenticating to the i2pcontrol interface")
    KeyPath := flag.String("key", "/var/lib/i2pd/i2pcontrol.key.pem",
		"password for authenticating to the i2pcontrol interface")
    Debug := flag.Bool("debug", true,
		"version of the i2pcontrol interface")
	flag.Parse()

	apiversion := *ApiVersion
	i2pcontrolhost := *AddrString
	i2pcontrolport := *PortString
	password := *Password
    certpath := *CertPath
    keypath := *KeyPath

	i2pcontrol.I2pControlVerboseLogging = *Debug

	auth := i2pcontrol.NewI2pControl(apiversion, certpath, keypath, password, i2pcontrolhost, i2pcontrolport)

	auth.Echo("i2pcontrol")

}
