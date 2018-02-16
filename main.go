package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

type Args struct {
	GrafanaURL  *string
	GrafanaPort *string
	OrgId       *string
	ApiKey      *string
}

func main() {
	fmt.Println("Starting intelegrafana.")
	parseFlags()
	fmt.Println("Finished execution. Exiting intelegrafana now...")
	fmt.Println("Bye!!!")
}

func parseFlags() {
	fmt.Println("Parsing command line arguments.")
	args := Args{}
	args.GrafanaURL = flag.String("grafanaURL", "", "grafana endpoint")
	args.GrafanaPort = flag.String("grafanaPort", "", "grafana port")
	args.OrgId = flag.String("orgId", "", "Dashboard organization ID")
	args.ApiKey = flag.String("apiKey", "", "Organization API Key")
	flag.Parse()
	fmt.Println(fmt.Sprintf("No of command line arguments that have been set: %d", flag.NFlag()))
	if flag.NFlag() < 4 ||
		strings.Trim(*args.GrafanaURL, " ") == "" ||
		strings.Trim(*args.GrafanaPort, " ") == "" ||
		strings.Trim(*args.OrgId, " ") == "" ||
		strings.Trim(*args.ApiKey, " ") == "" {
		fmt.Println("Invalid statup agruments...")
		fmt.Println("Usage: ./intelegrafana -grafanUrl=<URL> -grafanaPort=<port> orgId=<orgid> -apiKey=<apikey>")
		fmt.Println("Exiting now...")
		fmt.Println("Bye!!!")
		os.Exit(1)
	}
}
