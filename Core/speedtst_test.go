package speedtst_test

import (
	speedtst "exemple.com/speedtstMirceaD/Core"
)

func TestSpeedtestNetOK() {
	speedtst.RunSpeedTest("SpeedtestNetProvider", nil)
}

func TestSpeedtestNetNOK() {

}

func TestFastcomOK() {
	speedtst.RunSpeedTest("fastcom", nil)
}

func TestFastComNOK() {

}
