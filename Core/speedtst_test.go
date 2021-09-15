package speedtst_test

import (
	"testing"

	speedtstapi "exemple.com/speedtstMirceaD/API"
	speedtst "exemple.com/speedtstMirceaD/Core"
)

const FastcomDwlBenchmarkMin = 78.14;
const FastcomDwlBenchmarkMax = 486.98;

func TestSpeedtestNetOK(t *testing.T) {
	//result := speedtst.RunSpeedTest("speedtestnet", nil)
	//t.Errorf(result.FriendlyErrorMessage);
}

func TestSpeedtestNetNOK(t *testing.T) {

}

func DoFastcomTestOK(providerCfg *speedtstapi.SpeedTestProviderConfig, t *testing.T) {
	result:= speedtst.RunSpeedTest("fastcom", providerCfg)
	if result.WasSuccesfull == false {
		t.Fatalf("Test failed with friendly error message: %s", result.FriendlyErrorMessage);
		return;
	}

	if(result.Speeds.DownloadspeedMbps < FastcomDwlBenchmarkMin / 3.0 || result.Speeds.DownloadspeedMbps > FastcomDwlBenchmarkMax * 3.0) {
		t.Errorf("Test results straid too far from benchmark. Expected between %f/3 and %f*3, got %f.", FastcomDwlBenchmarkMin, FastcomDwlBenchmarkMax, result.Speeds.DownloadspeedMbps)
		return;
	}
}

func TestFastcomOKDefaultConfig(t *testing.T) {
	DoFastcomTestOK(nil, t);
}

func TestFastcomOKCustomConfing(t *testing.T) {
	fastComOKconfig := new(speedtstapi.SpeedTestProviderConfig);
	fastComOKconfig.Fields = make(map[string]string);
	fastComOKconfig.Fields["Token"]="YXNkZmFzZGxmbnNkYWZoYXNkZmhrYWxm";
	fastComOKconfig.Fields["UrlCount"]="7";
	DoFastcomTestOK(fastComOKconfig, t);
}

func TestFastComNOKInvalidToken(t *testing.T) {
	fastComNOKconfig := new(speedtstapi.SpeedTestProviderConfig);
	fastComNOKconfig.Fields = make(map[string]string);
	fastComNOKconfig.Fields["Token"]="qwerty";
	result:= speedtst.RunSpeedTest("fastcom", fastComNOKconfig);
	if(result.WasSuccesfull) {
		t.Fatalf("Speed test succeded when it should have failed with token \"qwerty\".");
	}
	if(result.HadLocalError == true || result.HadRemoteOrNetworkError == false || 
		result.FriendlyErrorMessage != "The number of response target URLs was less than 1. It seams that Fast.com changed format of response, or maybe an invalid token was used.") {
		t.Errorf("Speed failed but for the wrong reasons: %s", result.FriendlyErrorMessage);
	}
}
