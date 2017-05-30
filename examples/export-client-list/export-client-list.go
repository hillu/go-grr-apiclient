package main

import (
	"github.com/hillu/go-grr-apiclient"

	"crypto/tls"
	"encoding/csv"
	"flag"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	var urlstr string
	flag.StringVar(&urlstr, "url", "http://admin:admin@localhost:8000", "Base URL for GRR AdminUI server")
	flag.Parse()
	baseurl, err := url.Parse(urlstr)
	if err != nil {
		log.Fatal(err)
	}
	c := &apiclient.APIClient{
		BaseURL: baseurl,
		Client: &http.Client{Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		},
	}
	result, err := c.SearchClients(apiclient.ApiSearchClientsArgs{})
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()
	w.Write([]string{"Client ID", "Host", "Labels"})
	for _, client := range result.GetItems() {
		var labels []string
		for _, label := range client.GetLabels() {
			labels = append(labels, label.GetName())
		}
		w.Write([]string{
			client.GetUrn(),
			client.GetOsInfo().GetNode(),
			strings.Join(labels, " "),
		})
	}
}
