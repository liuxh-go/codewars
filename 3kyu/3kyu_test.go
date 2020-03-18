package _kyu

import (
	"testing"
)

func TestRailFenceCipher(t *testing.T) {
	s := "o ivrDie adofnbn oei!a ad tsimt e  fsA soi ueemc vamrtitqapre pucuonxtau!arere pueiiesns i ntrkroeeatsrrinPoim.cpfa tfglilep uV,deeepronumutou uat teietviiuja m,ti tuecq  eumil sirxe itnei a!d ssomseni iomsiietrdarnstiirsec aiiseeaa"
	n := 2

	//result := Encode(s, n)
	//t.Log(result)

	result := Decode(s, n)
	t.Log(result)
}
