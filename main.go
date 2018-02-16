package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"bufio"
)

type CliArgs struct {
	GrafanaURL  *string
	GrafanaPort *string
	OrgId       *string
	ApiKey      *string
}

type GrafanaArgs struct {
	DashboardName  string
	DatasourceName string
	Ports          string
	URLs           string
	Process        string
}

var DashboardInfo *GrafanaArgs = &GrafanaArgs{}
var Args CliArgs = CliArgs{}

func main() {
	fmt.Println("Starting intelegrafana.")
	//parseFlags()
	parseInputs()
	fmt.Println(*DashboardInfo)
	fmt.Println("Finished execution. Exiting intelegrafana now...")
	fmt.Println("Bye!!!")
}

func parseInputs() {
	dashboardName()
	datasourceName()
	ports()
	urls()
	process()
}

func parseFlags() {
	fmt.Println("Parsing command line arguments.")
	args := CliArgs{}
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

func dashboardName() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter dashboard name: ")
	input, _ := reader.ReadString('\n')
	count := 1
	for ; count < 3; {
		if input == "\n" || input == "" || strings.Trim(input, " ") == "" {
			fmt.Println("Please enter a valid dashboard name:")
			input, _ = reader.ReadString('\n')
		} else {
			DashboardInfo.DashboardName = strings.Trim(input, " ")
			break
		}
		count++
		checkExit(count, 3)
	}
}

func datasourceName() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter datasource name: ")
	input, _ := reader.ReadString('\n')
	count := 1
	for ; count < 3; {
		if input == "\n" || input == "" || strings.Trim(input, " ") == "" {
			fmt.Println("Please enter a valid datasource name:")
			input, _ = reader.ReadString('\n')
		} else {
			DashboardInfo.DatasourceName = strings.Trim(input, " ")
			break
		}
		count++
		checkExit(count, 3)
	}
}

func ports() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to monitor ports [y/n]:")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(input)
	isYesNo := string([]byte(input)[0])
	if isYesNo == "Y" {
		fmt.Print("Enter comma(,) separated ports:")
		input, _ := reader.ReadString('\n')
		count := 1
		for ; count < 3; {
			if input == "\n" || input == "" || strings.Trim(input, " ") == "" {
				fmt.Println("Please enter valid ports:")
				input, _ = reader.ReadString('\n')
			} else {
				DashboardInfo.Ports = strings.Trim(input, " ")
				break
			}
			count++
			checkExit(count, 3)
		}
	}
}

func urls() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to monitor urls [y/n]:")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(input)
	isYesNo := string([]byte(input)[0])
	if isYesNo == "Y" {
		fmt.Print("Enter comma(,) separated urls:")
		input, _ := reader.ReadString('\n')
		count := 1
		for ; count < 3; {
			if input == "\n" || input == "" || strings.Trim(input, " ") == "" {
				fmt.Println("Please enter valid urls:")
				input, _ = reader.ReadString('\n')
			} else {
				DashboardInfo.URLs = strings.Trim(input, " ")
				break
			}
			count++
			checkExit(count, 3)
		}
	}
}

func process() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to monitor process [y/n]:")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(input)
	isYesNo := string([]byte(input)[0])
	if isYesNo == "Y" {
		fmt.Print("Enter comma(,) separated process name (e.g. server.js):")
		input, _ := reader.ReadString('\n')
		count := 1
		for ; count < 3; {
			if input == "\n" || input == "" || strings.Trim(input, " ") == "" {
				fmt.Println("Please enter valid process name:")
				input, _ = reader.ReadString('\n')
			} else {
				DashboardInfo.Process = strings.Trim(input, " ")
				break
			}
			count++
			checkExit(count, 3)
		}
	}
}

func checkExit(count, max int) {
	if count == max {
		fmt.Println("You've made 3 unsuccessful attempts.")
		fmt.Println("Exiting now...")
		fmt.Println("Bye!!!")
		os.Exit(1)
	}
}
