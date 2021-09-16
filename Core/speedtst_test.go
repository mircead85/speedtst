package speedtst_test

import (
	"testing"

	speedtstapi "exemple.com/speedtstMirceaD/API"
	speedtst "exemple.com/speedtstMirceaD/Core"
)

const FastcomDwlBenchmarkMin = 78.14;
const FastcomDwlBenchmarkMax = 486.98;

const SpeedtestDwlBenchmarkMin = 918.77;
const SpeedtestDwlBenchmarkMax = 933.22;

const SpeedtestUplBenchmarkMin = 410.03;
const SpeedtestUplBenchmarkMax = 940.20;


func DoTestOK(providerName string, providerCfg *speedtstapi.SpeedTestProviderConfig, dwlBenchMin float32, dwlBenchMax float32, uplBenchMin float32, uplBenchMax float32,  t *testing.T) {
	result:= speedtst.RunSpeedTest(providerName, providerCfg)
	if result.WasSuccesfull == false {
		t.Fatalf("Test failed with friendly error message: %s", result.FriendlyErrorMessage);
		return;
	}

	if(dwlBenchMin >= 0 && dwlBenchMax >=0) {
		if(result.Speeds.DownloadspeedMbps < dwlBenchMin / 3.0 || result.Speeds.DownloadspeedMbps > dwlBenchMax * 3.0) {
			t.Errorf("Test results for download speed straid too far from benchmark. Expected between %f/3 and %f*3, got %f.", dwlBenchMin, dwlBenchMax, result.Speeds.DownloadspeedMbps)
			return;
		}
	}

	if(uplBenchMin >= 0 && uplBenchMax >=0) {
		if(result.Speeds.UploadspeedMbps < uplBenchMin / 3.0 || result.Speeds.UploadspeedMbps > uplBenchMax * 3.0) {
			t.Errorf("Test results for upload speed straid too far from benchmark. Expected between %f/3 and %f*3, got %f.", uplBenchMin, uplBenchMax, result.Speeds.UploadspeedMbps)
			return;
		}
	}
}

func TestSpeednetOKDefaultConfig(t *testing.T) {
	DoTestOK("speedtestnet", nil, SpeedtestDwlBenchmarkMin, SpeedtestDwlBenchmarkMax, SpeedtestUplBenchmarkMin, SpeedtestUplBenchmarkMax, t);
}

func TestSpeednetNetOKCustomCofig(t *testing.T) {
	speedtestnetConfig := new(speedtstapi.SpeedTestProviderConfig);
	speedtestnetConfig.Fields = make(map[string]string);
	speedtestnetConfig.Fields["CLICommand"]="../Core/../External/speedtest.exe";
	DoTestOK("speedtestnet", speedtestnetConfig, FastcomDwlBenchmarkMin, FastcomDwlBenchmarkMax, SpeedtestUplBenchmarkMin, SpeedtestUplBenchmarkMax, t);
}

func TestSpeednetNOK(t *testing.T) {
	speedtestnetConfig := new(speedtstapi.SpeedTestProviderConfig);
	speedtestnetConfig.Fields = make(map[string]string);
	speedtestnetConfig.Fields["CLICommand"]="speedtest.exe";
	
	result:= speedtst.RunSpeedTest("speedtestnet", speedtestnetConfig)
	if result.WasSuccesfull == true {
		t.Fatalf("Speed test succeded when it should have failed with reporting data specified in Kbps not Mbps.");
		return;
	}

	if result.HadLocalError == false || result.HadRemoteOrNetworkError == true {
			t.Errorf("Speed failed but for the wrong reasons: %s", result.FriendlyErrorMessage);
	}
}

func TestFastcomOKDefaultConfig(t *testing.T) {
	DoTestOK("fastcom", nil, FastcomDwlBenchmarkMin, FastcomDwlBenchmarkMax, -1.0, -1.0, t);
}

func TestFastcomOKCustomConfing(t *testing.T) {
	fastComOKconfig := new(speedtstapi.SpeedTestProviderConfig);
	fastComOKconfig.Fields = make(map[string]string);
	fastComOKconfig.Fields["Token"]="YXNkZmFzZGxmbnNkYWZoYXNkZmhrYWxm";
	fastComOKconfig.Fields["UrlCount"]="7";
	DoTestOK("fastcom", fastComOKconfig, FastcomDwlBenchmarkMin, FastcomDwlBenchmarkMax, -1.0, -1.0, t);
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
