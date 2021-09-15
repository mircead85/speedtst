package speedtstapi //We define here common types and interfaces for all Speed Test Providers

//What we are actually interested in
type SpeedTestResults struct {
	uploadspeedMbps   float32 //in MBps
	downloadspeedMbps float32 //in MBps
}

//Info about how an ongoing or completed test (atttempt) went
type SpeedTestAttemptResult struct {
	TestAttemptIsRunning    bool
	WasSuccesfull           bool
	HadLocalError           bool
	HadRemoteOrNetworkError bool
	FriendlyErrorMessage    string
	Speeds                  *SpeedTestResults
}

type ErrorRet string

type SpeedTestProviderConfig struct {
	Fields map[string]string //Each provider will interpret this key/value collection as it sees fit
}

//A generic interface to represent a provider capable of doing an internet speed test
type SpeedTestProvider interface {
	DefaultConfig() *SpeedTestProviderConfig
	Init(config *SpeedTestProviderConfig) ErrorRet
	DoSpeedTest(attemptResult *SpeedTestAttemptResult)
}
