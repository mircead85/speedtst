package speedtst

import (
	speedtstapi "exemple.com/speedtstMirceaD/API"
	fastcomprovider "exemple.com/speedtstMirceaD/FastComProvider"
	speedtestnetprovider "exemple.com/speedtstMirceaD/SpeedtestNetProvider"
)

func init() {
	//Read Config file if possible itf to initialize magical constants throughout.
}

func getProvider(providerName string) speedtstapi.SpeedTestProvider{
	if providerName == "speedtestnet" {
		return new(speedtestnetprovider.SpeedNetProvider)
	}
	
	if(providerName == "fastcom") {
		return new(fastcomprovider.FastComProvider);
	}

	return nil;
}

func RunSpeedTest(speedTestProviderName string, providerConfig *speedtstapi.SpeedTestProviderConfig) *speedtstapi.SpeedTestAttemptResult {
	result := new(speedtstapi.SpeedTestAttemptResult)

	provider := getProvider(speedTestProviderName)
	if provider == nil {
		result.TestAttemptIsRunning = false
		result.HadLocalError = true
		result.FriendlyErrorMessage = "Specified Provider Name is not recognized."
		return result
	}

	var initResult = provider.Init(providerConfig)
	if initResult != "" {
		result.TestAttemptIsRunning = false
		result.HadLocalError = true
		result.FriendlyErrorMessage = "Configuration specified (including potentially default) is invalid for the chosen Provider or failed to initialize: " + string(initResult) + "."
		return result
	}

	provider.DoSpeedTest(result)
	
	return result
}
