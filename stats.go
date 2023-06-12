package i2pcontrol

import (
	"fmt"
	"log"
)

// RateStat executes a GetRate call
func RateStat(Stat string, Period int) (int, error) {
	retpre, err := Call("GetRate", map[string]interface{}{
		"Stat":   Stat,
		"Period": Period,
		"Token":  token,
	})
	if err != nil {
		return -1, err
	}
	result := int(retpre["Result"].(float64))
	return result, nil
}

func SendBps() (int, error) {
	return RateStat("bw.sendBps", 300000)
}

func ReceiveBps() (int, error) {
	return RateStat("bw.receiveBps", 300000)
}

func ExploratoryBuildExpire() (int, error) {
	return RateStat("tunnel.buildExploratoryExpire", 600000)
}

func ExploratoryBuildReject() (int, error) {
	return RateStat("tunnel.buildExploratoryReject", 600000)
}

func ExploratoryBuildSuccess() (int, error) {
	return RateStat("tunnel.buildExploratorySuccess", 600000)
}

func ExploratoryBuildRejectPercentage() (int, error) {
	explReject, err := ExploratoryBuildReject()
	if err != nil {
		return -1, fmt.Errorf("unable to calculate exploratory build rejection percent: %s", err)
	}
	explSuccess, err := ExploratoryBuildSuccess()
	if err != nil {
		log.Println(err)
	}
	explExpire, err := ExploratoryBuildExpire()
	if err != nil {
		log.Println(err)
	}
	explTotal := explReject + explSuccess + explExpire
	if explTotal == 0 {
		return 0, nil
	}
	return int(float64(explReject) / float64(explTotal) * 100), nil
}

func ExploratoryBuildSuccessPercentage() (int, error) {
	explSuccess, err := ExploratoryBuildSuccess()
	if err != nil {
		return -1, fmt.Errorf("unable to calculate exploratory build success percent: %s", err)
	}
	explReject, err := ExploratoryBuildReject()
	if err != nil {
		log.Println(err)
	}
	explExpire, err := ExploratoryBuildExpire()
	if err != nil {
		log.Println(err)
	}
	explTotal := explReject + explSuccess + explExpire
	if explTotal == 0 {
		return 0, nil
	}
	return int(float64(explSuccess) / float64(explTotal) * 100), nil
}

func ExploratoryBuildExpirePercentage() (int, error) {
	explExpire, err := ExploratoryBuildExpire()
	if err != nil {
		return -1, fmt.Errorf("unable to calculate exploratory build expire percent: %s", err)
	}
	explSuccess, err := ExploratoryBuildSuccess()
	if err != nil {
		log.Println(err)
	}
	explReject, err := ExploratoryBuildReject()
	if err != nil {
		log.Println(err)
	}
	explTotal := explReject + explSuccess + explExpire
	if explTotal == 0 {
		return 0, nil
	}
	return int(float64(explExpire) / float64(explTotal) * 100), nil
}
