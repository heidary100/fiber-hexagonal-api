package crawler

import (
	"github.com/gocolly/colly/v2"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetHttpClient() *http.Client {
	client := &http.Client{Timeout: 10 * time.Second}
	return client
}

func GetResponseBody(endPoint string, client *http.Client) (io.ReadCloser, error) {
	req, err := http.NewRequest("GET", endPoint, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return response.Body, nil
}

//func GetFilesFromWebPage(url string, extensions []string) ([]string, error) {
//	responseBody, err := GetResponseBody(url, GetHttpClient())
//	if err != nil {
//		return nil, err
//	}
//
//	// Load the HTML document
//	doc, err := goquery.NewDocumentFromReader(responseBody)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	var files []string
//	// Find the review items
//	for _, ext := range extensions {
//		doc.Find("a[href*='." + ext + "']").Each(func(i int, s *goquery.Selection) {
//			// For each item found, get the title
//			href, exists := s.Attr("href")
//			if exists {
//				files = append(files, href)
//			}
//		})
//	}
//
//	return files, nil
//}

func GetFilesFromWebPage(url string, extensions []string) ([]string, error) {
	c := colly.NewCollector()
	var files []string
	var selectors []string
	for _, ext := range extensions {
		selectors = append(selectors, "a[href*='."+ext+"']")
	}
	c.OnHTML(strings.Join(selectors, " , "), func(e *colly.HTMLElement) {
		href := e.Attr("href")
		files = append(files, href)
	})

	err := c.Visit(url)
	if err != nil {
		return nil, err
	}

	return files, nil
}
