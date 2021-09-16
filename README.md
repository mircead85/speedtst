# speedtst
GO SpeedTst Demo

This is a class library in GO which tests the download and upload speeds by using Ookla's https://www.speedtest.net/ and Netflix's https://fast.com/.

For speedtest.net it relies on the speedtest.net CLI tools, downloaded from https://www.speedtest.net/apps/cli. Usage of this module implies acceptance of the licencese agreement for those tools, also as shown when running "speedtest.exe" (where it asks you to accept license - slienced in this library).

For Fast.com only download speed is reported.

Certain minor custom configuration options are allowed, as model is generic (allowing more Providers to be easily added). See each provider file's ..Provider struct for details.

Sample Usage1:

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

Output: 
Fast.com download speed (average) in Mbps: 565.86
Speedtest.net download speed (average) in Mbps: 931.31
Speedtest.net upload speed (average) in Mbps: 940.61

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

Hope you enjoy it! I am not by far fully satisified with all OOP decisions (especially failure of generic JSON parsing) and all, but, my first 1 day GO project, I hope you enjoyed it.

To run the tests, do from the project's root directory (where this README.md is located):

go test ./Core

To run the main file which shows some sample usage do from the project's root directory (where this README.md is located):

go run .


I tested the project on Windows with Visual Studio Code terminal.
For me, for tests and run to work properly, the project files needed to be located on the %GOPATH%/src/ path (so C:\GoProg\src\speedtstapi is the project root dir for me).

Cheers!