package i2pcontrol

import (
	//"log"
	"testing"
)

func TestAuth(t *testing.T) {
	Initialize("localhost", "7657", "jsonrpc")
	result, err := Authenticate("itoopie")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Connected to API version 1:", result)
	t.Log("Token for this testing session:", token)
	echo, err := Echo("Hello I2PControl")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(echo, "Is there an echo in here?")
	update, err := FindUpdates()
	if err != nil {
		t.Fatal(err)
	}
	if update {
		t.Log("Your I2P router needs an update")
	} else {
		t.Log("Your I2P router doesn't need an update")
	}

}
