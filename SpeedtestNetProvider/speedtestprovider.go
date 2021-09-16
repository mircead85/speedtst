package speedtestnetprovider

import (
	"strings"
	"fmt"
	"os/exec"
	"strconv"

	speedtstapi "exemple.com/speedtstMirceaD/API"
)

const mcCLIcommand = "../External/speedtest.exe";
const mcCLIcommandArg = "**-u****Mbps**";

type SpeedNetProvider struct {
	clicommand	  	 string //"CLICommand" in Config.Fields
	clicommandarg	 string //"CLICommandArg" in Config.Fields
}

func (provider *SpeedNetProvider) DefaultConfig() *speedtstapi.SpeedTestProviderConfig {
	cfg := new(speedtstapi.SpeedTestProviderConfig)
	cfg.Fields = make(map[string]string)
	return cfg
}

func (provider *SpeedNetProvider) Init(config *speedtstapi.SpeedTestProviderConfig) speedtstapi.ErrorRet {
	clicom := mcCLIcommand
	clicomarg := mcCLIcommandArg

	effectiveConfig := config
	if effectiveConfig == nil {
		effectiveConfig = provider.DefaultConfig()
	}

	sCLIComConfig := effectiveConfig.Fields["CLICommand"]
	if sCLIComConfig != "" {
		clicom = sCLIComConfig;
	}

	sCLIComArg := effectiveConfig.Fields["CLICommandArg"]
	if sCLIComArg != "" {
		clicomarg = sCLIComArg;
	}

	provider.clicommand = clicom;
	provider.clicommandarg = clicomarg;

	return speedtstapi.ErrorRet("")
}

func getDataOutOfOutputFileFormat(stringOutput string, speedtype string, attemptResult *speedtstapi.SpeedTestAttemptResult) float32 {
	slicesspeedtype1 := strings.Split(stringOutput, speedtype);
	if(len(slicesspeedtype1) != 2) {
		attemptResult.FriendlyErrorMessage = "Unexpected output format of SpeedTest.net CLI tool";
		return -1.0;
	}

	slicesspeedtype2 := strings.Split(slicesspeedtype1[1], "Mbps");
	if(len(slicesspeedtype2) < 1) {
		attemptResult.FriendlyErrorMessage = "Unexpected output format of SpeedTest.net CLI tool";
		return -1.0;
	}

	speedtypeSpeedinMbpsStr := strings.Trim(slicesspeedtype2[0], " \t\r\n");
	speedtypeSpeedinMbps, err2 := strconv.ParseFloat(speedtypeSpeedinMbpsStr, 32);

	if(err2 != nil) {
		attemptResult.FriendlyErrorMessage = "Unexpected output format of SpeedTest.net CLI tool";
		return -1.0;
	}

	return float32(speedtypeSpeedinMbps);
}

func (provider *SpeedNetProvider) DoSpeedTest(attemptResult *speedtstapi.SpeedTestAttemptResult) {
	attemptResult.WasSuccesfull = false;
	attemptResult.HadLocalError = true;

	cmd := exec.Command(provider.clicommand, provider.clicommandarg);

	cmd.Stdin = strings.NewReader("YES\r\n");
	cmdOutput, err := cmd.Output();

	if(err != nil) {
		attemptResult.FriendlyErrorMessage = fmt.Sprintf("CLI Command for Speedtest.Net failed to execute properly: %s", err);
		return;
	}

	stringOutput := string(cmdOutput);

	attemptResult.HadLocalError = false;
	attemptResult.HadRemoteOrNetworkError = true;
	
	downloadSpeedinMbps := getDataOutOfOutputFileFormat(stringOutput, "Download:", attemptResult);

	if(downloadSpeedinMbps == -1.0) {
		return; //Error had occured in parsing.
	}

	uploadSpeedinMbps := getDataOutOfOutputFileFormat(stringOutput, "Upload:", attemptResult);

	if(uploadSpeedinMbps == -1.0) {
		return; //Error had occured in parsing;
	}

	attemptResult.HadRemoteOrNetworkError=false;
	attemptResult.WasSuccesfull=true;
	attemptResult.Speeds = new(speedtstapi.SpeedTestResults);
	attemptResult.Speeds.DownloadspeedMbps = downloadSpeedinMbps;
	attemptResult.Speeds.UploadspeedMbps = uploadSpeedinMbps;
}
