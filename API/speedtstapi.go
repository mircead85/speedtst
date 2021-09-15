package speedtstapi //We define here common types and interfaces for all Speed Test Providers
//I tried to declared this as a speedtestapi package, but because of beginner status with Go [on Windows], I was unable to reference it

//What we are actually interested in
type SpeedTestResults struct {
	uploadspeedMbps   float32 //in MBps
	downloadspeedMbps float32 //in MBps
}

//Info about how an ongoing or completed test (atttempt) went
type speedTestAttemptResult struct {
	testAttemptIsRunning    bool
	wasSuccesfull           bool
	hadLocalError           bool
	hadRemoteOrNetworkError bool
	friendlyErrorMessage    string
}

//Each provider will interpret this key/value collection as it sees fit
type SpeedTestProviderConfig []string

//A generic interface to represent a provider capable of doing an internet speed test
type speedTestProvider interface {
	init(config SpeedTestProviderConfig)
	doSpeedTest(resetifRunning bool, blockThread bool)
	lastSpeedTestResults() *SpeedTestResults
}
