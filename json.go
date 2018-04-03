package i2pcontrol

type jsonStructure struct {
    echo string
}

func (j *jsonStructure) Id() string {
    rstring := randomString(10)
    return rstring
}

func (j *jsonStructure) Format(m string, s []string) string {
    rstring := "{\n"
    rstring += "method: \"" + m +"\"\n"
    rstring += "jsonrpc: \"2.0\"\n"
    rstring += "id: \"" + j.Id() + "\"\n"
    rstring += "params: \""
    for _, value := range s {
        rstring += value + " "
    }
    rstring += "\"\n"
    rstring += "}\n"
    return rstring
}

func (j *jsonStructure) Echo(s ...string) string{
    return j.Format("echo", s)
}

func (j *jsonStructure) jsonStructure() jsonStructure {
    return *j
}

func NewJsonStructure() *jsonStructure {
    var j jsonStructure
    return &j
}
