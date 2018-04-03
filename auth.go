package i2pcontrol

import "fmt"

type i2pControlStructure struct {
    jsonstructure  *jsonStructure
    i2pcontrolhost string
    i2pcontrolport string
}

func (i *i2pControlStructure) i2pControlHost() string{
    return i.i2pcontrolhost
}

func (i *i2pControlStructure) i2pControlPort() string{
    return i.i2pcontrolport
}

func (i *i2pControlStructure) i2pControlDo(s ...string) string{
    return ""
}

func (i *i2pControlStructure) Authenticate(s string) (string, string) {
    query := i.jsonstructure.Authenticate("API" , "1", "Password", s)
    fmt.Println(query)
    response := i.i2pControlDo(query)
    return query, response
}
func (i *i2pControlStructure) Echo(s string) (string, string) {
    query := i.jsonstructure.Echo("API" , "1", "Token", s)
    fmt.Println(query)
    response := i.i2pControlDo(query)
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
    i.jsonstructure = NewJsonStructure()
    return &i
}
