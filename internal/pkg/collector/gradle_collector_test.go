package collector

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func Test_GetGradleSha256ByUrl(t *testing.T) {
	resp, _ := httpGetByProxy("https://downloads.gradle-dn.com/distributions/gradle-8.0.2-bin.zip.sha256")
	defer resp.Body.Close()
	bytes, _ := ioutil.ReadAll(resp.Body)
	t.Run("", func(t *testing.T) {
		fmt.Println(string(bytes))
	})

}

func Test_getCollectorVersionShaSum(t *testing.T) {
	resp, _ := httpGetByProxy("https://gradle.org/release-checksums/")

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("false")
	}
	versions := make([]*string, 0)
	maps := make(map[string]string)
	doc_selector, _ := goquery.NewDocumentFromReader(resp.Body)
	t.Run("", func(t *testing.T) {
		docs := doc_selector.Find(".layout__main")
		a_docs := docs.Find("a[name]")
		a_docs.Each(func(j int, a_doc *goquery.Selection) {
			version := a_doc.AttrOr("name", "")
			if version != "" {
				versions = append(versions, &version)

			}
		})
		lis := docs.Find("ul[style]")

		for i := 0; i < len(versions); i++ {
			sha256 := lis.Eq(i).Find("li").Eq(i).Find("code").Text()
			maps[*versions[i]] = sha256
			fmt.Printf("version: %v ,  sha256: %v\n", *versions[i], sha256)
		}
		/* 	fmt.Printf("len(versions): %v\n", len(versions))
		s := lis.Eq(0).Find("li").Eq(0).Find("code")
		fmt.Printf("lis: %v\n", s.Text())
		docs.Each(func(j int, doc *goquery.Selection) {

		}) */
	})
}

func Test_getCollectorVersion(t *testing.T) {
	resp, _ := httpGetByProxy("https://gradle.org/releases/")
	//resp, _ := http.Get("https://gradle.org/releases/")
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("false")
	}
	doc_selector, _ := goquery.NewDocumentFromReader(resp.Body)
	t.Run("", func(t *testing.T) {
		a_docs := doc_selector.Find(".resources-contents").Find("a[name]")
		a_docs.Each(func(j int, a_doc *goquery.Selection) {
			version := a_doc.AttrOr("name", "")
			fmt.Printf("version: %v\n", version)
		})
	})
}

func httpGetByProxy(url string) (*http.Response, error) {
	request, _ := http.NewRequest("GET", url, nil)
	return getClientProxy().Do(request)
}

func getClientProxy() *http.Client {
	proxyUrl, err := url.Parse("http://127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}
