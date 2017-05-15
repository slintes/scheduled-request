package main

import (
	"flag"
	"github.com/jasonlvhit/gocron"
	"log"
	"net/http"
	"strings"
)

var interval uint64
var url string
var method string
var contentType string
var body string

func init() {
	flag.Uint64Var(&interval, "interval", 0, "Interval in minutes.")
	flag.StringVar(&url, "url", "", "Url to call.")
	flag.StringVar(&method, "method", "GET", "Http method, GET and POST are supported. Default GET.")
	flag.StringVar(&contentType, "contentType", "application/json", "If method is POST, the content type of the body. Default application/json.")
	flag.StringVar(&body, "body", "{}", "If method is POST, the post body. Default {}.")

	flag.Parse()

	if interval == 0 || len(url) == 0 {
		flag.Usage()
		log.Fatalln("Missing Argument!")
	}

	if !strings.EqualFold(method, "POST") && !strings.EqualFold(method, "GET") {
		flag.Usage()
		log.Fatalln("Unkown method %v, supported methods are GET and POST", method)
	}
}

func main() {

	flag.Visit(printFlag)
	log.Print("Starting...")
	gocron.Every(interval).Minutes().Do(call)
	<-gocron.Start()

}

func printFlag(flag *flag.Flag) {
	log.Printf("%s: %v", flag.Name, flag.Value)
}

func call() {
	var response *http.Response

	var err error

	if strings.EqualFold(method, "POST") {
		response, err = http.Post(url, contentType, strings.NewReader(body))
		defer response.Body.Close()
	} else if strings.EqualFold(method, "GET") {
		response, err = http.Get(url)
		defer response.Body.Close()
	} else {
		log.Printf("Unknown method %v!", method)
	}

	if err != nil {
		log.Printf("Error requesting %v: %v", url, err.Error())
	} else {
		log.Printf("Requested %v, response status: %v", url, response.Status)
	}
}
