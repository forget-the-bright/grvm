package proxy

import (
	"net/http"
	"net/url"
	"os"
	"strings"
)

func HttpGetByProxy(url string) (*http.Response, error) {
	if cli := getClientProxy(); cli != nil {
		request, _ := http.NewRequest("GET", url, nil)
		return cli.Do(request)
	} else {
		return http.Get(url)
	}

}

func getClientProxy() *http.Client {
	if proxy := getEnvProxy(); proxy == "" {
		return nil
	} else {
		proxyUrl, err := url.Parse(proxy)
		if err != nil {
			panic(err)
		}
		return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
	}
}

func getEnvProxy() string {
	if proxy_url := getEnvAny("HTTP_PROXY", "http_proxy"); proxy_url != "" {
		if !strings.Contains(proxy_url, "http://") {
			return "http://" + proxy_url
		}
		return proxy_url
	}
	return ""
}

func getEnvAny(names ...string) string {
	for _, n := range names {
		if val := os.Getenv(n); val != "" {
			return val
		}
	}
	return ""
}
