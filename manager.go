package i2pcontrol

func Echo(echo string) (string, error) {
	retpre, err := Call("Echo", map[string]interface{}{
		"Echo":  echo,
		"Token": token,
	})
	if err != nil {
		return "", err
	}
	result := retpre["Result"].(string)
	return result, nil
}

func RestartGraceful() (string, error) {
	_, err := Call("RouterManager", map[string]interface{}{
		"RestartGraceful": nil,
		"Token":           token,
	})
	if err != nil {
		return "", err
	}
	return "Graceful Restart Initiated", nil
}

func Restart() (string, error) {
	_, err := Call("RouterManager", map[string]interface{}{
		"Restart": nil,
		"Token":   token,
	})
	if err != nil {
		return "", err
	}
	return "Restart Initiated", nil
}

func ShutdownGraceful() (string, error) {
	_, err := Call("RouterManager", map[string]interface{}{
		"ShutdownGraceful": nil,
		"Token":            token,
	})
	if err != nil {
		return "", err
	}
	return "Graceful Shutdown Initiated", nil
}

func Shutdown() (string, error) {
	_, err := Call("RouterManager", map[string]interface{}{
		"Shutdown": nil,
		"Token":    token,
	})
	if err != nil {
		return "", err
	}
	return "Shutdown Initiated", nil
}

func FindUpdates() (bool, error) {
	retpre, err := Call("RouterManager", map[string]interface{}{
		"FindUpdates": nil,
		"Token":       token,
	})
	if err != nil {
		return false, err
	}
	result := retpre["FindUpdates"].(bool)
	return result, nil
}

func Update() (string, error) {
	_, err := Call("RouterManager", map[string]interface{}{
		"Update": nil,
		"Token":  token,
	})
	if err != nil {
		return "", err
	}
	return "Update initiated", nil
}
