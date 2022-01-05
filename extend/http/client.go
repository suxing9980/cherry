package cherryHttp

import (
	"bytes"
	"encoding/json"
	cherryLogger "github.com/cherry-game/cherry/logger"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	postContentType = "application/x-www-form-urlencoded"
	jsonContentType = "application/json"

	DefaultTimeout = 5 * time.Second
)

func GET(url string, values ...map[string]string) (string, *http.Response, error) {
	client := http.Client{Timeout: DefaultTimeout}

	if len(values) > 0 {
		rst := ToUrlValues(values[0])
		url = AddParams(url, rst)
	}

	rsp, err := client.Get(url)
	if err != nil {
		return "", rsp, err
	}

	defer func(body io.ReadCloser) {
		e := body.Close()
		if e != nil {
			cherryLogger.Warnf("HTTP GET [url = %s], error = %s", url, e)
		}
	}(rsp.Body)

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", rsp, err
	}

	return string(body), rsp, nil
}

func POST(url string, values map[string]string) (string, *http.Response, error) {
	client := http.Client{Timeout: DefaultTimeout}

	rst := ToUrlValues(values)
	rsp, err := client.Post(url, postContentType, strings.NewReader(rst.Encode()))
	if err != nil {
		return "", rsp, err
	}

	defer func(body io.ReadCloser) {
		e := body.Close()
		if e != nil {
			cherryLogger.Warnf("HTTP POST [url = %s], error = %s", url, e)
		}
	}(rsp.Body)

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", rsp, err
	}

	return string(body), rsp, nil
}

func PostJSON(url string, values interface{}) (string, *http.Response, error) {
	client := http.Client{Timeout: DefaultTimeout}

	jsonBytes, err := json.Marshal(values)
	if err != nil {
		return "", nil, err
	}

	rsp, err := client.Post(url, jsonContentType, bytes.NewBuffer(jsonBytes))
	if err != nil {
		return "", rsp, err
	}

	defer func(Body io.ReadCloser) {
		e := Body.Close()
		if e != nil {
			cherryLogger.Warnf("HTTP PostJSON [url = %s], error = %s", url, e)
		}
	}(rsp.Body)

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return "", rsp, err
	}

	return string(body), rsp, nil
}

func AddParams(url string, params url.Values) string {
	if len(params) == 0 {
		return url
	}

	if !strings.Contains(url, "?") {
		url += "?"
	}

	if strings.HasSuffix(url, "?") || strings.HasSuffix(url, "&") {
		url += params.Encode()
	} else {
		url += "&" + params.Encode()
	}

	return url
}

func ToUrlValues(values map[string]string) url.Values {
	rst := make(url.Values)
	for k, v := range values {
		rst.Add(k, v)
	}
	return rst
}
