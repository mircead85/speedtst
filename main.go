package main

import (
	"fmt"
	
	speedtst "exemple.com/speedtstMirceaD/Core"
	speedtstapi "exemple.com/speedtstMirceaD/API"
)

func main() {
	result:= speedtst.RunSpeedTest("fastcom", nil)
	if result.WasSuccesfull == false {
		fmt.Printf("Fast.com attempt failed with message: %s\n", result.FriendlyErrorMessage)
	} else {
		fmt.Printf("Fast.com download speed (average) in Mbps: %.2f\n",result.Speeds.DownloadspeedMbps)
	}

	speedtestnetConfig := new(speedtstapi.SpeedTestProviderConfig);
	speedtestnetConfig.Fields = make(map[string]string);
	speedtestnetConfig.Fields["CLICommand"]="./External/speedtest.exe";
	result = speedtst.RunSpeedTest("speedtestnet", speedtestnetConfig);
	if(result.WasSuccesfull==false) {
		fmt.Printf("Speedtest.net CLI attempt failed with message: %s\n", result.FriendlyErrorMessage)
	} else {
		fmt.Printf("Speedtest.net download speed (average) in Mbps: %.2f\n", result.Speeds.DownloadspeedMbps);
		fmt.Printf("Speedtest.net upload speed (average) in Mbps: %.2f\n", result.Speeds.UploadspeedMbps);
	}
}
