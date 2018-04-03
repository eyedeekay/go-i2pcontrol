package i2pcontrol

import (
    "fmt"
    "net/rpc"
)

type i2pControlStructure struct {
    jsonstructure  *jsonStructure
    i2pcontrolhost string
    i2pcontrolport string
    i2pcontroltoken string
    i2pcontrolclient *rpc.Client
    i2pcontrolerr error
}

func (i *i2pControlStructure) i2pControlHost() string{
    return i.i2pcontrolhost
}

func (i *i2pControlStructure) i2pControlPort() string{
    return i.i2pcontrolport
}

func (i *i2pControlStructure) i2pControlDo(method string, params ...string) string{
    var reply *string
    i.i2pcontrolerr = i.i2pcontrolclient.Call(method, params, reply)
    if i.i2pcontrolerr != nil {
        fmt.Println("reply error")
    }
    return *reply
}

func (i *i2pControlStructure) Authenticate(v, pw string) (string, string) {
    query := i.jsonstructure.Authenticate("API" , v, "Password", pw)
    fmt.Println(query)
    response := i.i2pControlDo("Authenticate", pw)
    return query, response
}
func (i *i2pControlStructure) Echo(e string) (string, string) {
    if i.i2pcontroltoken == "" {
        return "", "i2pControl Auth Token not present."
    }
    query := i.jsonstructure.Echo("Token", i.i2pcontroltoken, "Echo", e)
    fmt.Println(query)
    response := i.i2pControlDo("Echo", e)
    return query, response
}

func (i *i2pControlStructure) i2pControl() i2pControlStructure {
    return *i
}

func NewI2pControl(hostport ...string) *i2pControlStructure {
    var i i2pControlStructure
    if hostport != nil {
        if len(hostport) > 0 {
            i.i2pcontrolhost = hostport[0]
        }else if len(hostport) > 1 {
            i.i2pcontrolport = hostport[1]
        }
    }
    i.i2pcontroltoken = ""
    i.i2pcontrolclient, i.i2pcontrolerr = rpc.DialHTTP("tcp", i.i2pcontrolhost +":"+ i.i2pcontrolport )
    i.jsonstructure = NewJsonStructure()
    return &i
}
