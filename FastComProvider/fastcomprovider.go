package fastcomprovider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	speedtstapi "exemple.com/speedtstMirceaD/API"
)

//URL used to access the FAST.com API. Would normally be relegated to an app config file (in C# I mean).
const mcFmtString = "http%s://api.fast.com/netflix/speedtest/v2?https=%s&token=%s&urlCount=%d"
const mcDefaultHttps = false
const mcDefaultToken = "YXNkZmFzZGxmbnNkYWZoYXNkZmhrYWxm"
const mcDefaultUrlCount = 3

type FastcomResponse struct
{
	client FastcomClient 
	targets []FastcomTarget
}
type FastcomClient struct
{
	ip string
	asn string
	isp string
	location FastcomLocation
}

type FastcomLocation struct
{
	city string
	country string
}

type FastcomTarget struct
{
 	name string
	url string
	location FastcomLocation
}

type FastComProvider struct {
	fmtFullUrl  string //Inited based on Config fields "UseHttps", "Token" and "UrlCount" which have good default values.
}

func (provider *FastComProvider) DefaultConfig() *speedtstapi.SpeedTestProviderConfig {
	cfg := new(speedtstapi.SpeedTestProviderConfig)
	cfg.Fields = make(map[string]string)
	return cfg
}

func (provider *FastComProvider) Init(config *speedtstapi.SpeedTestProviderConfig) speedtstapi.ErrorRet {
	bHttps := mcDefaultHttps
	token := mcDefaultToken
	urlCount := mcDefaultUrlCount

	effectiveConfig := config
	if effectiveConfig == nil {
		effectiveConfig = provider.DefaultConfig()
	}

	sHttpsConfig := effectiveConfig.Fields["UseHttps"]
	if sHttpsConfig != "" {
		bHttpsConfig, bOK := strconv.ParseBool(sHttpsConfig)
		if bOK != nil {
			return "ERROR: Invalid UseHttps value. Boolean string expected." //Should be extracted to Error messages globably (Global Map/Enum)
		}
		bHttps = bHttpsConfig
	}

	tokenConfig := effectiveConfig.Fields["Token"]

	if tokenConfig != "" {
		token = tokenConfig
	}

	sUrlCountConfig := effectiveConfig.Fields["UrlCount"]
	if sUrlCountConfig != "" {
		urlCountConfig, bOK := strconv.ParseUint(sUrlCountConfig, 10, 32)
		if bOK != nil || urlCountConfig > 10 {
			return "ERROR: Invalid UrlCount value. Unsigned int <= 10 string expected." //Should be extracted to Error messages globably (Global MapEnum)
		}
		urlCount = int(urlCountConfig)
	}

	sHttpsPref := ""
	sHttpsPram := "false"
	if bHttps {
		sHttpsPref = "s"
		sHttpsPram = "true"
	}

	provider.fmtFullUrl = fmt.Sprintf(mcFmtString, sHttpsPref, sHttpsPram, token, urlCount)

	return speedtstapi.ErrorRet("")
}

func (provider *FastComProvider) DoSpeedTest(attemptResult *speedtstapi.SpeedTestAttemptResult) {
	attemptResult.WasSuccesfull = false
	attemptResult.HadRemoteOrNetworkError = true

	resp, err := http.Get(provider.fmtFullUrl)

	if err != nil {
		attemptResult.FriendlyErrorMessage = "Base URL invalid for Fast.com (maybe they changed it, or the token expired)?.."
		return
	}

	respBody, err2 := ioutil.ReadAll(resp.Body)
	resp.Body.Close();

	if err2 != nil {
		//Better to introduce logging mechanism in the future
		attemptResult.FriendlyErrorMessage = "Error reading response from Fast.com. Maybe network error."
		return
	}

	/* //What I tried with JSON parsing and didn't manage to get it working in alloted time.
	var fastcomResponseObj FastcomResponse;

	attemptResult.FriendlyErrorMessage=string(respBody);

	err3 := json.Unmarshal(respBody, &fastcomResponseObj)

	if err3 != nil {
		attemptResult.FriendlyErrorMessage = "Error parsing response body as json - maybe Fast.com changed format."
		return
	}
		
	targets := fastcomResponseObj.targets;

	attemptResult.FriendlyErrorMessage = "afterJsoning" + fmt.Sprintf("%s", targets[0].url);
	/**/

	var totalDwlTimeInMbs float64;
	totalSlices := 0;
	var averageDwlTimeInMbps float64;
	
	respBodyStrSlices := strings.Split(fmt.Sprintf("%s",respBody), "\"url\":\"");
	if len(respBodyStrSlices) < 2 {
		attemptResult.FriendlyErrorMessage = "The number of response target URLs was less than 1. It seams that Fast.com changed format of response, or maybe an invalid token was used."
		return
	}
	
	for _, targetUrlSlice := range respBodyStrSlices {
		indexOfNextQuote := strings.Index(targetUrlSlice, "\"");
		if(indexOfNextQuote < 0) {
			break;
		}
		targetUrl := targetUrlSlice[0:indexOfNextQuote];
		startTimeInNs := time.Now();
		respTarget, err4 := http.Get(targetUrl);
		if(err4 != nil) {
			continue;
		}

		respBodyTarget, err5 := ioutil.ReadAll(respTarget.Body)
		respTarget.Body.Close();
		endTimeInNs := time.Now();

		if err5 != nil {
			continue;
		}

		bodyLenInBytes := len(respBodyTarget);

		totalSlices++;
		sliceDwlTime := float64(endTimeInNs.Sub(startTimeInNs).Seconds());
		curSpeed := (float64(bodyLenInBytes*8) / (1024.0*1024.0)  ) / sliceDwlTime
		totalDwlTimeInMbs+=curSpeed;
	}

	if(totalSlices == 0) {
		attemptResult.FriendlyErrorMessage = "No target URL succeded. Maybe Fast.com changed the policies or the methods in getting results.";
		return;
	}

	averageDwlTimeInMbps = totalDwlTimeInMbs / float64(totalSlices);

	attemptResult.WasSuccesfull = true;
	attemptResult.HadRemoteOrNetworkError = false;
	attemptResult.Speeds = new(speedtstapi.SpeedTestResults);
	attemptResult.Speeds.DownloadspeedMbps = float32(averageDwlTimeInMbps);
	attemptResult.Speeds.UploadspeedMbps = -1.0; //Fast.com does not support Upload speed test.
}
