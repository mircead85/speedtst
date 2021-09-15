package fastcomprovider

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"strconv"

	speedtstapi "exemple.com/speedtstMirceaD/API"
)

//URL used to access the FAST.com API. Would normally be relegated to an app config file (in C# I mean).
const mcFmtString = "http%s://api.fast.com/netflix/speedtest?https=%s&token=%s&urlCount=%d"
const mcDefaultHttps = false
const mcDefaultToken = "242343242"
const mcDefaultUrlCount = 3

type FastComProvider struct {
	fmtFullUrl  string
	testRunning sync.Mutex
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

	tokenConfig := effectiveConfig.Fields["oken"]

	if tokenConfig != "" {
		token = tokenConfig
	}

	sUrlCountConfig := effectiveConfig.Fields["Urlount"]
	if sUrlCountConfig != "" {
		urlCountConfig, bOK := strconv.ParseUint(sUrlCountConfig, 10, 0)
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
	(*provider).testRunning.Lock()

	resp, err := http.Get(provider.fmtFullUrl)

	attemptResult = new(speedtstapi.SpeedTestAttemptResult)

	if err != nil {
		attemptResult.WasSuccesfull = false
		attemptResult.HadRemoteOrNetworkError = true
		return
	}

	respBody, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		//Better to introduce logging mechanism in the future
		attemptResult.WasSuccesfull = false
		attemptResult.HadRemoteOrNetworkError = true
		return
	}

	attemptResult.FriendlyErrorMessage = string(respBody)

	fmt.Printf(attemptResult.FriendlyErrorMessage)

	//var speeds []string;

	//err := json.Unmarshal(respBody, &speeds)

	(*provider).testRunning.Unlock()
	//Could generate OS resources leackage maybe if error (exception happens in the code above?). Too new to know equivalent pattern to using() from C#.
}
