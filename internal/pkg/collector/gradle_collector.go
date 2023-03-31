package collector

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/forget-the-bright/grvm/internal/pkg/config"
	"github.com/forget-the-bright/grvm/internal/pkg/proxy"
)

func getGradleAllInfo() []*GradleItem {
	resp, _ := proxy.HttpGetByProxy(Collector_Archive_Url)
	resp_checksums, _ := proxy.HttpGetByProxy(Collector_Release_Checksums)

	defer resp.Body.Close()
	defer resp_checksums.Body.Close()
	doc_selector, _ := goquery.NewDocumentFromReader(resp.Body)
	doc__checksums_selector, _ := goquery.NewDocumentFromReader(resp_checksums.Body)

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
		//fmt.Printf("version: %v    time: %v     sha256: %v\n", version, "", sha256)
		grvms = append(grvms, build_GradleItem(version, version_time, sha256))
	})
	//fmt.Printf("len(grvms): %v\n", len(grvms))
	return config.ReverseArray(grvms)
}
