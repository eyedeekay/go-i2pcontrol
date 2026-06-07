package i2pcontrol

// Checks our UPnp
func Upnp() (string, error) {
	retpre, err := Call("NetworkSetting", map[string]interface{}{
		"i2p.router.net.upnp": nil,
		"Token":               token,
	})
	if err != nil {
		return "", err
	}
	if retpre["i2p.router.net.upnp"] == nil {
		return "Upnp Disabled", nil
	}
	result := retpre["i2p.router.net.upnp"].(string)

	return result, nil
}
