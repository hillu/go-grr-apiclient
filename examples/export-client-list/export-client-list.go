package main

import (
	"github.com/hillu/go-grr-apiclient"

	"crypto/tls"
	"encoding/csv"
	"flag"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"
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
		Client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		},
	}
	c.Client.Jar, _ = cookiejar.New(nil)
	result, err := c.SearchClients(apiclient.ApiSearchClientsArgs{})
	if err != nil {
		log.Fatal(err)
	}
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()
	w.Write([]string{"Client ID", "Host", "Labels", "Hardware", "First Seen", "Last Booted", "Last Seen",
		"System", "Version", "Manufacturer", "Architecture"})
	for _, client := range result.GetItems() {
		var labels []string
		for _, label := range client.GetLabels() {
			labels = append(labels, label.GetName())
		}

		// Write Line
		w.Write([]string{
			client.GetUrn(),
			client.GetOsInfo().GetNode(),
			strings.Join(labels, " "),
			time.Unix(int64(client.GetFirstSeenAt())/1000000, 0).String(),
			time.Unix(int64(client.GetLastBootedAt())/1000000, 0).String(),
			time.Unix(int64(client.GetLastSeenAt())/1000000, 0).String(),
			client.GetOsInfo().GetSystem(),
			client.GetOsInfo().GetVersion(),
			client.GetHardwareInfo().GetSystemManufacturer(),
			client.GetOsInfo().GetMachine(),
		})
	}
}
