package moviesservice

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/heidary100/fiber-hexagonal-api/internal/core/ports"
	"github.com/heidary100/fiber-hexagonal-api/internal/presenter"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"time"
)

type service struct {
	repository ports.MoviesRepository
}

func NewService(r ports.MoviesRepository) ports.MoviesService {
	return &service{
		repository: r,
	}
}

func (s *service) FetchMovieUrls(name string, extensions []string) ([]presenter.FetchUrlResponse, error) {
	googleResult, err := searchInGoogle(name, 0)
	start := time.Now()
	ch := make(chan presenter.FetchUrlResponse)
	for _, gr := range googleResult.Organic {
		go fetch(gr.Url, extensions, ch) // start a goroutine
	}
	var furArray []presenter.FetchUrlResponse
	for range googleResult.Organic {
		furArray = append(furArray, <-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	return furArray, err
}

func (s *service) Search(name string) (presenter.MovieSearchResponse, error) {
	sr, err := searchInTMDB(name)
	return sr, err
}

func searchInTMDB(name string) (presenter.MovieSearchResponse, error) {
	safeName := url.QueryEscape(name)
	tmdbUrl := fmt.Sprintf("https://api.themoviedb.org/3/search/movie?api_key=a8ac2f9446eab16741b3adf87e14cfe9&language=en-US&page=1&include_adult=false&query=%s", safeName)

	var record presenter.MovieSearchResponse
	// Build the request
	req, err := http.NewRequest("GET", tmdbUrl, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return record, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return record, err
	}

	defer resp.Body.Close()

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return record, nil
}

func searchInGoogle(name string, start int) (presenter.GoogleSearchResponse, error) {
	//safeName := url.QueryEscape(name)
	//rapidApiUrl := fmt.Sprintf("https://google-search1.p.rapidapi.com/google-search?gl=us&hl=en&q=%s", safeName)
	//
	//var record presenter.GoogleSearchResponse
	//// Build the request
	//req, err := http.NewRequest("GET", rapidApiUrl, nil)
	//if err != nil {
	//	log.Fatal("NewRequest: ", err)
	//	return record, err
	//}
	//
	//req.Header.Add("X-RapidAPI-Key", "b087fd0fafmsh336a8c4c9e88212p18c739jsndaa75a368e4d")
	//req.Header.Add("X-RapidAPI-Host", "google-search1.p.rapidapi.com")
	//
	//client := &http.Client{}
	//resp, err := client.Do(req)
	//if err != nil {
	//	log.Fatal("Do: ", err)
	//	return record, err
	//}
	//
	//defer resp.Body.Close()
	//
	//// Use json.Decode for reading streams of JSON data
	//
	//if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
	//	log.Println(err)
	//}

	var gsr presenter.GoogleSearchResponse
	tempBody := `{
    "currentPage": 1,
    "keyword": "nodejs",
    "organic": [
        {
            "domain": "nodejs.org",
            "linkType": "HOME",
            "position": 1,
            "snippet": "Node.js® is a JavaScript runtime built on Chrome's V8 JavaScript engine.",
            "title": "Node.js",
            "url": "https://nodejs.org/"
        },
        {
            "domain": "nodejs.dev",
            "linkType": "HOME",
            "position": 2,
            "snippet": "",
            "title": "Nodejs.dev",
            "url": "https://nodejs.dev/"
        },
        {
            "domain": "www.w3schools.com",
            "linkType": "LANDING",
            "position": 3,
            "snippet": "",
            "title": "Node.js Introduction - W3Schools",
            "url": "https://www.w3schools.com/nodejs/nodejs_intro.asp"
        },
        {
            "domain": "en.wikipedia.org",
            "linkType": "LANDING",
            "position": 4,
            "snippet": "",
            "title": "Node.js - Wikipedia",
            "url": "https://en.wikipedia.org/wiki/Node.js"
        },
        {
            "domain": "www.toptal.com",
            "linkType": "LANDING",
            "position": 5,
            "snippet": "",
            "title": "Why The Hell Would I Use Node.js? A Case-by-Case Tutorial",
            "url": "https://www.toptal.com/nodejs/why-the-hell-would-i-use-node-js"
        },
        {
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
        {
            "domain": "nodejs.org",
            "linkType": "HOME",
            "position": 1,
            "snippet": "Node.js® is a JavaScript runtime built on Chrome's V8 JavaScript engine.",
            "title": "Node.js",
            "url": "https://nodejs.org/"
        },
        {
            "domain": "nodejs.dev",
            "linkType": "HOME",
            "position": 2,
            "snippet": "",
            "title": "Nodejs.dev",
            "url": "https://nodejs.dev/"
        },
        {
            "domain": "www.w3schools.com",
            "linkType": "LANDING",
            "position": 3,
            "snippet": "",
            "title": "Node.js Introduction - W3Schools",
            "url": "https://www.w3schools.com/nodejs/nodejs_intro.asp"
        },
        {
            "domain": "en.wikipedia.org",
            "linkType": "LANDING",
            "position": 4,
            "snippet": "",
            "title": "Node.js - Wikipedia",
            "url": "https://en.wikipedia.org/wiki/Node.js"
        },
        {
            "domain": "www.toptal.com",
            "linkType": "LANDING",
            "position": 5,
            "snippet": "",
            "title": "Why The Hell Would I Use Node.js? A Case-by-Case Tutorial",
            "url": "https://www.toptal.com/nodejs/why-the-hell-would-i-use-node-js"
        },
        {
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
        {
            "domain": "nodejs.org",
            "linkType": "HOME",
            "position": 1,
            "snippet": "Node.js® is a JavaScript runtime built on Chrome's V8 JavaScript engine.",
            "title": "Node.js",
            "url": "https://nodejs.org/"
        },
        {
            "domain": "nodejs.dev",
            "linkType": "HOME",
            "position": 2,
            "snippet": "",
            "title": "Nodejs.dev",
            "url": "https://nodejs.dev/"
        },
        {
            "domain": "www.w3schools.com",
            "linkType": "LANDING",
            "position": 3,
            "snippet": "",
            "title": "Node.js Introduction - W3Schools",
            "url": "https://www.w3schools.com/nodejs/nodejs_intro.asp"
        },
        {
            "domain": "en.wikipedia.org",
            "linkType": "LANDING",
            "position": 4,
            "snippet": "",
            "title": "Node.js - Wikipedia",
            "url": "https://en.wikipedia.org/wiki/Node.js"
        },
        {
            "domain": "www.toptal.com",
            "linkType": "LANDING",
            "position": 5,
            "snippet": "",
            "title": "Why The Hell Would I Use Node.js? A Case-by-Case Tutorial",
            "url": "https://www.toptal.com/nodejs/why-the-hell-would-i-use-node-js"
        },
        {
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
        {
            "domain": "nodejs.org",
            "linkType": "HOME",
            "position": 1,
            "snippet": "Node.js® is a JavaScript runtime built on Chrome's V8 JavaScript engine.",
            "title": "Node.js",
            "url": "https://nodejs.org/"
        },
        {
            "domain": "nodejs.dev",
            "linkType": "HOME",
            "position": 2,
            "snippet": "",
            "title": "Nodejs.dev",
            "url": "https://nodejs.dev/"
        },
        {
            "domain": "www.w3schools.com",
            "linkType": "LANDING",
            "position": 3,
            "snippet": "",
            "title": "Node.js Introduction - W3Schools",
            "url": "https://www.w3schools.com/nodejs/nodejs_intro.asp"
        },
        {
            "domain": "en.wikipedia.org",
            "linkType": "LANDING",
            "position": 4,
            "snippet": "",
            "title": "Node.js - Wikipedia",
            "url": "https://en.wikipedia.org/wiki/Node.js"
        },
        {
            "domain": "www.toptal.com",
            "linkType": "LANDING",
            "position": 5,
            "snippet": "",
            "title": "Why The Hell Would I Use Node.js? A Case-by-Case Tutorial",
            "url": "https://www.toptal.com/nodejs/why-the-hell-would-i-use-node-js"
        },
        {
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
        {
            "domain": "nodejs.org",
            "linkType": "HOME",
            "position": 1,
            "snippet": "Node.js® is a JavaScript runtime built on Chrome's V8 JavaScript engine.",
            "title": "Node.js",
            "url": "https://nodejs.org/"
        },
        {
            "domain": "nodejs.dev",
            "linkType": "HOME",
            "position": 2,
            "snippet": "",
            "title": "Nodejs.dev",
            "url": "https://nodejs.dev/"
        },
        {
            "domain": "www.w3schools.com",
            "linkType": "LANDING",
            "position": 3,
            "snippet": "",
            "title": "Node.js Introduction - W3Schools",
            "url": "https://www.w3schools.com/nodejs/nodejs_intro.asp"
        },
        {
            "domain": "en.wikipedia.org",
            "linkType": "LANDING",
            "position": 4,
            "snippet": "",
            "title": "Node.js - Wikipedia",
            "url": "https://en.wikipedia.org/wiki/Node.js"
        },
        {
            "domain": "www.toptal.com",
            "linkType": "LANDING",
            "position": 5,
            "snippet": "",
            "title": "Why The Hell Would I Use Node.js? A Case-by-Case Tutorial",
            "url": "https://www.toptal.com/nodejs/why-the-hell-would-i-use-node-js"
        },
        {
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        },
{
            "domain": "hub.docker.com",
            "linkType": "LANDING",
            "position": 6,
            "snippet": "",
            "title": "Node - Official Image | Docker Hub",
            "url": "https://hub.docker.com/_/node"
        }
    ],
    "totalResults": 248000000,
    "timeTaken": 0.4
}`

	err := json.Unmarshal([]byte(tempBody), &gsr)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%+v\n", gsr)
	}
	return gsr, nil
}

func fetch(url string, extensions []string, ch chan<- presenter.FetchUrlResponse) {
	start := time.Now()
	resp, err := http.Get(url)
	var fuResult presenter.FetchUrlResponse
	fuResult.WebPageUrl = url
	if err != nil {
		log.Fatal(err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		log.Fatal(err)
		return
	}

	links := doc.Find("a").Map(func(i int, sel *goquery.Selection) string {
		return sel.AttrOr("href", "")
	})

	var properLinks []string

	for _, link := range links {
		if hasExtension(extensions, link) {
			properLinks = append(properLinks, link)
		}
	}

	resp.Body.Close()
	if err != nil {
		log.Fatal(err)
		return
	}

	secs := time.Since(start).Seconds()
	fmt.Sprintf("%.2fs   %s", secs, url)

	fuResult.Urls = properLinks
	ch <- fuResult
}

func hasExtension(extensions []string, url string) bool {
	fExt := filepath.Ext(url)
	for _, eachExt := range extensions {
		if eachExt == fExt {
			return true
		}
	}
	return false
}
