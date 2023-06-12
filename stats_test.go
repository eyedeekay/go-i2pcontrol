package i2pcontrol

import (
	//"log"
	"testing"
)

func TestStats(t *testing.T) {
	Initialize("localhost", "7657", "jsonrpc")
	result, err := Authenticate("itoopie")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Connected to API version 1:", result)
	t.Log("Token for this testing session:", token)
	sbps, err := SendBps()
	if err != nil {
		t.Log(err)
	}
	t.Log("Sent BPS:", sbps)
	rbps, err := ReceiveBps()
	if err != nil {
		t.Log(err)
	}
	t.Log("Received BPS:", rbps)
	exp, err := ExploratoryBuildExpire()
	if err != nil {
		t.Log(err)
	}
	t.Log("Exploratory Build Expire:", exp)
	rej, err := ExploratoryBuildReject()
	if err != nil {
		t.Log(err)
	}
	t.Log("Exploratory Build Reject:", rej)
	succ, err := ExploratoryBuildSuccess()
	if err != nil {
		t.Log(err)
	}
	t.Log("Exploratory Build Success:", succ)
	expp, err := ExploratoryBuildExpirePercentage()
	if err != nil {
		t.Log(err)
	}
	t.Log("Exploratory Build Expire Percent:", expp)
	rejp, err := ExploratoryBuildRejectPercentage()
	if err != nil {
		t.Log(err)
	}
	t.Log("Exploratory Build Reject Percent:", rejp)
	succp, err := ExploratoryBuildSuccessPercentage()
	if err != nil {
		t.Log(err)
	}
	t.Log("Exploratory Build Success Percent:", succp)

}
