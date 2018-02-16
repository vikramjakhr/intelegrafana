package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"bufio"
	"encoding/json"
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
	createDashboard()
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
			DashboardInfo.DashboardName = strings.Trim(strings.Trim(input, " "), "\n")
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
			fmt.Println("Please enter a valid datasource name: ")
			input, _ = reader.ReadString('\n')
		} else {
			DashboardInfo.DatasourceName = strings.Trim(strings.Trim(input, " "), "\n")
			break
		}
		count++
		checkExit(count, 3)
	}
}

func ports() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to monitor ports [y/n]: ")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(input)
	isYesNo := string([]byte(input)[0])
	if isYesNo == "Y" {
		fmt.Print("Enter comma(,) separated ports: ")
		input, _ := reader.ReadString('\n')
		count := 1
		for ; count < 3; {
			if input == "\n" || input == "" || strings.Trim(input, " ") == "" {
				fmt.Println("Please enter valid ports: ")
				input, _ = reader.ReadString('\n')
			} else {
				DashboardInfo.Ports = strings.Trim(strings.Trim(input, " "), "\n")
				break
			}
			count++
			checkExit(count, 3)
		}
	}
}

func urls() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to monitor urls [y/n]: ")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(input)
	isYesNo := string([]byte(input)[0])
	if isYesNo == "Y" {
		fmt.Print("Enter comma(,) separated urls: ")
		input, _ := reader.ReadString('\n')
		count := 1
		for ; count < 3; {
			if input == "\n" || input == "" || strings.Trim(input, " ") == "" {
				fmt.Println("Please enter valid urls: ")
				input, _ = reader.ReadString('\n')
			} else {
				DashboardInfo.URLs = strings.Trim(strings.Trim(input, " "), "\n")
				break
			}
			count++
			checkExit(count, 3)
		}
	}
}

func process() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to monitor process [y/n]: ")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(input)
	isYesNo := string([]byte(input)[0])
	if isYesNo == "Y" {
		fmt.Print("Enter comma(,) separated process name (e.g. server.js): ")
		input, _ := reader.ReadString('\n')
		count := 1
		for ; count < 3; {
			if input == "\n" || input == "" || strings.Trim(input, " ") == "" {
				fmt.Println("Please enter valid process name: ")
				input, _ = reader.ReadString('\n')
			} else {
				DashboardInfo.Process = strings.Trim(strings.Trim(input, " "), "\n")
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

func createDashboard() {
	var dashboard map[string]interface{}
	var rows []map[string]interface{}
	var djson string = DashboardJson
	djson = strings.Replace(djson, "$DASHBOARD_NAME$", DashboardInfo.DashboardName, -1)
	djson = strings.Replace(djson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	row := createURLRow()
	if row != nil {
		rows = append(rows, createURLRow())
	}
	row = createPortsRow()
	if row != nil {
		rows = append(rows, createPortsRow())
	}
	row = createProcstatRow()
	if row != nil {
		rows = append(rows, createProcstatRow())
	}
	row = createDiskProcessSwapRow()
	if row != nil {
		rows = append(rows, createDiskProcessSwapRow())
	}
	row = createCPURAMRow()
	if row != nil {
		rows = append(rows, createCPURAMRow())
	}
	row = createIPSystemLoadRow()
	if row != nil {
		rows = append(rows, createIPSystemLoadRow())
	}
	json.Unmarshal([]byte(djson), &dashboard)
	dashboard["rows"] = rows
	byts, _ := json.MarshalIndent(dashboard, "", "   ")
	fmt.Println(string(byts))
}

func createURLRow() map[string]interface{} {
	if DashboardInfo.URLs != "" {
		urls := strings.Split(DashboardInfo.URLs, ",")
		var panels []map[string]interface{}
		for _, p := range urls {
			var paneljson string = URLPanel
			paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DashboardName, -1)
			paneljson = strings.Replace(paneljson, "$URL$", p, -1)
			var panel map[string]interface{}
			json.Unmarshal([]byte(paneljson), &panel)
			panels = append(panels, panel)

			paneljson = PageLoadPanel
			paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DashboardName, -1)
			paneljson = strings.Replace(paneljson, "$URL$", p, -1)
			panel = make(map[string]interface{})
			json.Unmarshal([]byte(paneljson), &panel)
			panels = append(panels, panel)
		}
		row := newRow(URL)
		row["panels"] = panels
		return row
	}
	return nil
}

func createPortsRow() map[string]interface{} {
	if DashboardInfo.Ports != "" {
		ports := strings.Split(DashboardInfo.Ports, ",")
		var panels []map[string]interface{}
		for _, p := range ports {
			var paneljson string = PortPanel
			paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DashboardName, -1)
			paneljson = strings.Replace(paneljson, "$PORT$", p, -1)
			var panel map[string]interface{}
			json.Unmarshal([]byte(paneljson), &panel)
			panels = append(panels, panel)
		}
		row := newRow(PORTS)
		row["panels"] = panels
		return row
	}
	return nil
}

func createProcstatRow() map[string]interface{} {
	if DashboardInfo.Process != "" {
		ports := strings.Split(DashboardInfo.Process, ",")
		var panels []map[string]interface{}
		for _, p := range ports {
			var paneljson string = ProcstatPanel
			paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DashboardName, -1)
			paneljson = strings.Replace(paneljson, "$PROCESS$", p, -1)
			var panel map[string]interface{}
			json.Unmarshal([]byte(paneljson), &panel)
			panels = append(panels, panel)
		}
		row := newRow(PROCSTATS)
		row["panels"] = panels
		return row
	}
	return nil
}

func createDiskProcessSwapRow() map[string]interface{} {
	var panels []map[string]interface{}
	var paneljson string = DiskPanel
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DashboardName, -1)
	var panel map[string]interface{}
	json.Unmarshal([]byte(paneljson), &panel)
	panels = append(panels, panel)

	paneljson = ProcessPanel
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DashboardName, -1)
	panel = make(map[string]interface{})
	json.Unmarshal([]byte(paneljson), &panel)
	panels = append(panels, panel)

	paneljson = Swap
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DashboardName, -1)
	panel = make(map[string]interface{})
	json.Unmarshal([]byte(paneljson), &panel)
	panels = append(panels, panel)

	row := newRow(DISK_PROCESS_SWAP)
	row["panels"] = panels
	return row
}

func createCPURAMRow() map[string]interface{} {
	var panels []map[string]interface{}
	var paneljson string = CPUPanel
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DashboardName, -1)
	var panel map[string]interface{}
	json.Unmarshal([]byte(paneljson), &panel)
	panels = append(panels, panel)

	paneljson = RAMPanel
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DashboardName, -1)
	panel = make(map[string]interface{})
	json.Unmarshal([]byte(paneljson), &panel)
	panels = append(panels, panel)

	row := newRow(CPU_RAM)
	row["panels"] = panels
	return row
}

func createIPSystemLoadRow() map[string]interface{} {
	var panels []map[string]interface{}
	var paneljson string = IPTrafficPanel
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DashboardName, -1)
	var panel map[string]interface{}
	json.Unmarshal([]byte(paneljson), &panel)
	panels = append(panels, panel)

	paneljson = SystemLoadPanel
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DashboardName, -1)
	panel = make(map[string]interface{})
	json.Unmarshal([]byte(paneljson), &panel)
	panels = append(panels, panel)

	row := newRow(IP_SYSTEMLOAD)
	row["panels"] = panels
	return row
}

const (
	URL               = "URL"
	PORTS             = "PORTS"
	PROCSTATS         = "PROCSTATS"
	DISK_PROCESS_SWAP = "DISK_PROCESS_SWAP"
	CPU_RAM           = "CPU_RAM"
	IP_SYSTEMLOAD     = "IP_SYSTEMLOAD"
)

func newRow(rType string) map[string]interface{} {
	row := map[string]interface{}{
		"collapse":        false,
		"repeat":          nil,
		"repeatIteration": nil,
		"repeatRowId":     nil,
		"showTitle":       false,
		"titleSize":       "h6",
	}
	switch rType {
	case URL:
		row["title"] = fmt.Sprintf("%s Row", rType)
		row["height"] = 80
		break
	case PORTS:
		row["title"] = fmt.Sprintf("%s Row", rType)
		row["height"] = 63
		break
	case PROCSTATS:
		row["title"] = fmt.Sprintf("%s Row", rType)
		row["height"] = 63
		break
	case DISK_PROCESS_SWAP:
		row["title"] = fmt.Sprintf("%s Row", rType)
		row["height"] = 230
		break
	case CPU_RAM:
		row["title"] = fmt.Sprintf("%s Row", rType)
		row["height"] = 230
		break
	case IP_SYSTEMLOAD:
		row["title"] = fmt.Sprintf("%s Row", rType)
		row["height"] = 230
		break
	}
	return row
}
