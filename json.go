package i2pcontrol


import (
    "fmt"
    "log"
    "strconv"
)

type request struct {
    method, jsonrpc, id string
    api int
    params *map[string]interface{}
}

type jsonStructure struct {
    api int64
}

func (j *jsonStructure) Api() string {
    return strconv.FormatInt(j.api, 10)
}

func (j *jsonStructure) Id() string {
    rstring := randomString(10)
    return rstring
}

func (j *jsonStructure) FormatRequest(rq request) string{
    if I2pControlVerboseLogging {
        rstring := "\n{\n"
        rstring += "  \"method\": \""+ rq.method +"\"\n"
        rstring += "  \"jsonrpc\": \""+ rq.jsonrpc +"\"\n"
        rstring += "  \"id\": \""+ rq.id +"\"\n"
        rstring += "  \"params\": {\n"
        for key, val := range *rq.params {
            rstring += "    \""+ key +"\": \"" + fmt.Sprintf("%v",val) +"\"\n"
        }
        rstring += "  }\n"
        rstring += "}\n"
        Log(rstring)
        return rstring
    }
    return ""
}

func (j *jsonStructure) Authenticate(v int64, pw string) (string, request) {
    m := "Authenticate"
    id := j.Id()
    rparams := request{
        id: id,
        method: m,
        params: &map[string]interface{}{
            "API": j.api,
            "Password": pw,
        },
        jsonrpc: "2.0",
    }
    return j.FormatRequest(rparams), rparams
}

func (j *jsonStructure) Echo(t string, s ...string) (string, request) {
    m := "Echo"
    id := j.Id()
    rparams := request{
        id: id,
        method: m,
        params: &map[string]interface{}{
            "Token": t,
            "Echo": s,
        },
        jsonrpc: "2.0",
    }
    return j.FormatRequest(rparams), rparams
}

func (j *jsonStructure) jsonStructure() jsonStructure {
    return *j
}

func NewJsonStructure(api int64) *jsonStructure {
    var j jsonStructure
    j.api = api
    return &j
}

func Log(s ...string) string {
    var str string
    for _, sub := range s {
        str += sub
    }
    if I2pControlVerboseLogging {
        log.Println(s)
    }
    return str
}

func Err(e error, l string, s ...string) (bool, string) {
    var str string
    for _, sub := range s {
        str += sub
    }
    b := false
    Log(l)
    if e != nil {
        b = true
        log.Fatal(s, e)
    }
    return b, str
}
