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

func Test_getGradleAllInfo(t *testing.T) {
	resp, _ := httpGetByProxy("https://gradle.org/releases/")
	resp_checksums, _ := httpGetByProxy("https://gradle.org/release-checksums/")
	//resp, _ := http.Get("https://gradle.org/releases/")
	defer resp.Body.Close()
	defer resp_checksums.Body.Close()
	doc_selector, _ := goquery.NewDocumentFromReader(resp.Body)
	doc__checksums_selector, _ := goquery.NewDocumentFromReader(resp_checksums.Body)
	t.Run("", func(t *testing.T) {
		grvms := make([]*GradleItem, 0)
		//获取所有的sha256
		lis := doc__checksums_selector.Find(".layout__main").Find("ul[style]")

		resources_docs := doc_selector.Find(".resources-contents")
		//获取所有的版本号
		a_docs := resources_docs.Find("a[name]")
		//获取所有的发布时间
		version_times := resources_docs.Find("p[class='u-text-with-icon u-no-margin-bottom u-no-margin-top']")

		a_docs.Each(func(j int, a_doc *goquery.Selection) {
			version := a_doc.AttrOr("name", "")
			sha256 := lis.Eq(j).Find("li").Eq(0).Find("code").Text()
			version_time := version_times.Eq(j).Find("span").Eq(1).Text()
			fmt.Printf("version: %v    time: %v     sha256: %v\n", version, "", sha256)
			grvms = append(grvms, &GradleItem{
				Version:     version,
				ReleaseTime: version_time,
				FileName:    "gradle-" + version + "-bin.zip",
				FileType:    "zip",
				Sha256:      sha256,
				Sha256Url:   "https://downloads.gradle-dn.com/distributions/gradle-" + version + "-bin.zip.sha256",
				DownloadUrl: "https://downloads.gradle-dn.com/distributions/gradle-" + version + "-bin.zip",
			})
		})
		fmt.Printf("len(grvms): %v\n", len(grvms))
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
		grvms := make([]*GradleItem, 0)
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
			sha256 := lis.Eq(i).Find("li").Eq(0).Find("code").Text()
			version := *versions[i]
			maps[version] = sha256
			fmt.Printf("version: %v ,  sha256: %v\n", version, sha256)
			grvms = append(grvms, &GradleItem{
				Version:     version,
				FileName:    "gradle-" + version + "-bin.zip",
				FileType:    "zip",
				Sha256:      sha256,
				Sha256Url:   "https://downloads.gradle-dn.com/distributions/gradle-" + version + "-bin.zip.sha256",
				DownloadUrl: "https://downloads.gradle-dn.com/distributions/gradle-" + version + "-bin.zip",
			})
		}
		fmt.Printf("len(grvms): %v\n", len(grvms))
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
	proxyUrl, err := url.Parse("http://127.0.0.1:7890")
	if err != nil {
		panic(err)
	}
	return &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}}
}
