package main

import (
	"fmt"
	
	speedtst "exemple.com/speedtstMirceaD/Core"
)

func main() {
	result:= speedtst.RunSpeedTest("fastcom", nil)
	if result.WasSuccesfull == false {
		fmt.Printf("Fast.com attempt failed with message: %s\n", result.FriendlyErrorMessage)
	} else {
		fmt.Printf("Fast.com download speed (average) in Mbps: %.2f\n",result.Speeds.DownloadspeedMbps)
	}

	result = speedtst.RunSpeedTest("speedtestnet", nil);
	if(result.WasSuccesfull==false) {
		fmt.Printf("Speedtest.net CLI attempt failed with message: %s\n", result.FriendlyErrorMessage)
	} else {
		fmt.Printf("Speedtest.net download speed (average) in Mbps: %.2f\n", result.Speeds.DownloadspeedMbps);
		fmt.Printf("Speedtest.net upload speed (average) in Mbps: %.2f\n", result.Speeds.UploadspeedMbps);
	}
}
