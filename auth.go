package i2pcontrol

import (
    "github.com/powerman/rpc-codec/jsonrpc2"
)

var I2pControlVerboseLogging bool

type i2pControlStructure struct {
    jsonstructure  *jsonStructure
    i2pcontrolhost string
    i2pcontrolport string
    i2pcontroltoken string
    i2pcontrolclient *jsonrpc2.Client
    i2pcontrolerr error
}

func (i *i2pControlStructure) i2pControlHost() string{
    return i.i2pcontrolhost
}

func (i *i2pControlStructure) i2pControlPort() string{
    return i.i2pcontrolport
}

func (i *i2pControlStructure) i2pControlAddress() string{
    return i.i2pControlHost() + ":" + i.i2pControlPort()
}

func (i *i2pControlStructure) i2pControlDo(method string, rq request) string{
    var reply *string
    i.i2pcontrolerr = i.i2pcontrolclient.Call(method, rq, reply)
    if c, _ := Err(i.i2pcontrolerr, "Sending request", "Reply error: "); !c {
        Log("Reply:", *reply)
    }
    return *reply
}

func (i *i2pControlStructure) Authenticate(v int64, pw string) (string, string) {
    q, query := i.jsonstructure.Authenticate(v, pw)
    response := i.i2pControlDo("Authenticate", query)
    return q, response
}

func (i *i2pControlStructure) Echo(e string) (string, string) {
    if i.i2pcontroltoken == "" {
        return "", "i2pControl Auth Token not present."
    }
    q, query := i.jsonstructure.Echo("Token", i.i2pcontroltoken, "Echo", e)
    response := i.i2pControlDo("Echo", query)
    return q, response
}

func (i *i2pControlStructure) i2pControl() i2pControlStructure {
    return *i
}

func NewI2pControl(api int64, password, host, port string) *i2pControlStructure {
    var i i2pControlStructure
    i.i2pcontrolhost = host
    i.i2pcontrolport = port
    Log("i2p Control Interface")
    Log("  i2pcontrolhost: ", i.i2pControlHost())
    Log("  i2pcontrolport: ", i.i2pControlPort())
    Log("Connecting:", i.i2pControlAddress())
    i.i2pcontrolclient, i.i2pcontrolerr = jsonrpc2.Dial("tcp", i.i2pControlAddress() )
    if c, _ := Err(i.i2pcontrolerr, "Connecting control interface", "Control interface connection error"); !c {
        Log("Connected Control Interface")
    }
    i.jsonstructure = NewJsonStructure(api)
    _, i.i2pcontroltoken = i.Authenticate(api, password)
    return &i
}
