# speedtst
GO SpeedTst Demo

This is a class library in GO which tests the download and upload speeds by using Ookla's https://www.speedtest.net/ and Netflix's https://fast.com/.

Sample Usage1:

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

Output: 
Fast.com download speed (average) in Mbps: 73.14

Sample Usage2 (with custom config):

import (
	"fmt"

	speedtst "exemple.com/speedtstMirceaD/Core"
    speedtstapi "exemple.com/speedtstMirceaD/API"
)

func main() {
    fastComOKconfig := new(speedtstapi.SpeedTestProviderConfig);
	fastComOKconfig.Fields = make(map[string]string);
	fastComOKconfig.Fields["UrlCount"]="7";
	result:= speedtst.RunSpeedTest("fastcom", nil)

	if result.WasSuccesfull == false {
		fmt.Printf("Fast.com attempt failed with message: %s", result.FriendlyErrorMessage)
	} else {
		fmt.Printf("Fast.com download speed (average) in Mbps: %.2f",result.Speeds.DownloadspeedMbps)
	}
}

Output: 
Fast.com download speed (average) in Mbps: 486.98

Hope you enjoy it!