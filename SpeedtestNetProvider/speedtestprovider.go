package speedtestnetprovider

import (
	"sync"

	speedtstapi "exemple.com/speedtstMirceaD/API"
)

type SpeedNetProvider struct {
	clicommandpath string
	clifileresult  string
	testRunning    *sync.Mutex
}

func (provider *SpeedNetProvider) DefaultConfig() *speedtstapi.SpeedTestProviderConfig {
	return nil
}

func (provider *SpeedNetProvider) Init(config *speedtstapi.SpeedTestProviderConfig) speedtstapi.ErrorRet {
	return speedtstapi.ErrorRet("")
}

func (provider *SpeedNetProvider) DoSpeedTest(attemptResult *speedtstapi.SpeedTestAttemptResult) {

}
