package i2pcontrol

func ParticipatingTunnels() (int, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.tunnels.participating": nil,
		"Token":                                token,
	})
	if err != nil {
		return -1, err
	}
	result := int(retpre["i2p.router.net.tunnels.participating"].(float64))
	return result, nil
}
