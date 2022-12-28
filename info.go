package i2pcontrol

import "fmt"

// ParticipatingTunnels gets the number of participating tunnels the router has currently
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

// Status queries the status of the router
func Status() (string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.status": nil,
		"Token":             token,
	})
	if err != nil {
		return "", err
	}
	result := retpre["i2p.router.status"].(string)
	return result, nil
}

// NetStatus queries the status of the network connection
func NetStatus() (string, error) {
	retpre, err := Call("RouterInfo", map[string]interface{}{
		"i2p.router.net.status": nil,
		"Token":                 token,
	})
	if err != nil {
		return "", err
	}
	result := int(retpre["i2p.router.net.status"].(float64))
	switch result {
	case 0:
		return "OK", nil
	case 1:
		return "TESTING", nil
	case 2:
		return "FIREWALLED", nil
	case 3:
		return "HIDDEN", nil
	case 4:
		return "WARN_FIREWALLED_AND_FAST", nil
	case 5:
		return "WARN_FIREWALLED_AND_FLOODFILL", nil
	case 6:
		return "WARN_FIREWALLED_WITH_INBOUND_TCP", nil
	case 7:
		return "WARN_FIREWALLED_WITH_UDP_DISABLED", nil
	case 8:
		return "ERROR_I2CP", nil
	case 9:
		return "ERROR_CLOCK_SKEW", nil
	case 10:
		return "ERROR_PRIVATE_TCP_ADDRESS", nil
	case 11:
		return "ERROR_SYMMETRIC_NAT", nil
	case 12:
		return "ERROR_UDP_PORT_IN_USE", nil
	case 13:
		return "ERROR_NO_ACTIVE_PEERS_CHECK_CONNECTION_AND_FIREWALL", nil
	case 14:
		return "ERROR_UDP_DISABLED_AND_TCP_UNSET", nil
	}
	return "Unexpected result", fmt.Errorf("Unexpected result %d", result)
}

// Reseeding checks if the I2P Router is reseeding
func Reseeding() (bool, error) {
	retpre, err := Call("Routerinfo", map[string]interface{}{
		"i2p.router.netdb.isreseeding": nil,
		"Token":                        token,
	})
	if err != nil {
		return false, err
	}
	result := retpre["i2p.router.netdb.isreseeding"].(bool)
	return result, nil
}
