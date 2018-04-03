package i2pcontrol


import (
    "strings"
)

type jsonStructure struct {
    echo string
}

func (j *jsonStructure) Id() string {
    rstring := randomString(10)
    return rstring
}

func (j *jsonStructure) Format(m string, s ...string) string {
    rstring := "{\n"
    rstring += "  \"method\": \"" + m +"\"\n"
    rstring += "  \"jsonrpc\": \"2.0\"\n"
    rstring += "  \"id\": \"" + j.Id() + "\"\n"
    rstring += "  \"params\": { \n"
    var params string
    for index, value := range s {
        if index != 0 && (index + 1)%2 == 0 {
            params += strings.TrimRight("    \"" + s[index-1] + "\": \"" + value + "\"", " ") + "\n"
        }
    }
    rstring += params + "  }"
    rstring += "}\n"
    return rstring
}

func (j *jsonStructure) Authenticate(s ...string) string{
    return j.Format("echo", s...)
}

func (j *jsonStructure) Echo(s ...string) string{
    return j.Format("echo", s...)
}

func (j *jsonStructure) jsonStructure() jsonStructure {
    return *j
}

func NewJsonStructure() *jsonStructure {
    var j jsonStructure
    return &j
}
