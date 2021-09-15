package main

import (
	"fmt"
	
	speedtst "exemple.com/speedtstMirceaD/Core"
)

func main() {
	result:= speedtst.RunSpeedTest("fastcom", nil)
	if result.WasSuccesfull == false {
		fmt.Printf("Fast.com attempt failed with message: %s", result.FriendlyErrorMessage)
	} else {
		fmt.Printf("Fast.com download speed (average) in Mbps: %.2f",result.Speeds.DownloadspeedMbps)
	}
}
