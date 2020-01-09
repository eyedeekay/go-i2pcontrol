package i2pcontrol

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/ybbus/jsonrpc"
	"io/ioutil"
	"net/http"
)

var (
	rpcClient jsonrpc.RPCClient
	RPCOpts   *jsonrpc.RPCClientOpts
	token     string
)

//
func Initialize(host, port, path string) {
	RPCOpts = &jsonrpc.RPCClientOpts{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
	}
	rpcClient = jsonrpc.NewClientWithOpts("http://"+host+":"+port+"/"+path+"/", RPCOpts)
}

func InitializeWithSelfSignedCert(host, port, path, cert string) error {
	caCert, err := ioutil.ReadFile(cert)
	if err != nil {
		return err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	RPCOpts = &jsonrpc.RPCClientOpts{
		HTTPClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: caCertPool,
				},
			},
		},
	}
	rpcClient = jsonrpc.NewClientWithOpts("http://"+host+":"+port+"/"+path+"/", RPCOpts)
	return nil
}

func Call(method string, params interface{}) (map[string]interface{}, error) {
	response, err := rpcClient.Call(method, params)
	if err != nil {
		return nil, err
	}
	if response.Error != nil {
		return nil, response.Error
	}
	//var retv string
	var retpre map[string]interface{}
	err = response.GetObject(&retpre)
	return retpre, nil
}

func Authenticate(password string) (int, error) {
	retpre, err := Call("Authenticate", map[string]interface{}{
		"API":      1,
		"Password": "itoopie",
	})
	if err != nil {
		return -1, err
	}
	token = retpre["Token"].(string)
	version := int(retpre["API"].(float64))
	return version, nil
}
