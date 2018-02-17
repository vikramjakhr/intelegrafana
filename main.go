package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"bufio"
	"encoding/json"
	"bytes"
	"io/ioutil"
	"net/url"
	"io"
	"net/http"
	"os/exec"
	"strconv"
	"math/rand"
)

type CliArgs struct {
	GrafanaURL  *string
	InfluxURL   *string
	InfluxPort  *string
	GrafanaPort *string
	OrgId       *string
	ApiKey      *string
}

type GrafanaArgs struct {
	DashboardName  string
	DatasourceName string
	Hostname       string
	Ports          string
	URLs           string
	Process        string
}

var DashboardInfo *GrafanaArgs = &GrafanaArgs{}
var Args CliArgs = CliArgs{}
var AgentConfig string

func main() {
	fmt.Println("Starting intelegrafana...")
	parseFlags()
	parseInputs()
	installTelegraf(getOSName())
	initAgentConfig()
	createDashboard()
	createAlertDashboard()
	changeTelegrafConfPermission()
	replaceTelegrafConfig()
	restartTelegraf()
	fmt.Println("Finished execution. Exiting intelegrafana now...")
	fmt.Println("Bye!!!")
}

func parseInputs() {
	dashboardName()
	datasourceName()
	hostname()
	ports()
	urls()
	process()
}

func initAgentConfig() {
	AgentConfig = TelegrafConfig
	AgentConfig = strings.Replace(AgentConfig, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	AgentConfig = strings.Replace(AgentConfig, "$INFLUX_URL$", *Args.InfluxURL, -1)
	AgentConfig = strings.Replace(AgentConfig, "$INFLUX_PORT$", *Args.InfluxPort, -1)
}

func parseFlags() {
	fmt.Println("Parsing command line arguments.")
	args := CliArgs{}
	args.GrafanaURL = flag.String("grafanaUrl", "", "grafana endpoint")
	args.GrafanaPort = flag.String("grafanaPort", "", "grafana port")
	args.InfluxURL = flag.String("influxUrl", "", "grafana endpoint")
	args.InfluxPort = flag.String("influxPort", "8086", "grafana port")
	args.OrgId = flag.String("orgId", "", "Dashboard organization ID")
	args.ApiKey = flag.String("apiKey", "", "Organization API Key")
	flag.Parse()
	fmt.Println(fmt.Sprintf("No of command line arguments that have been set: %d", flag.NFlag()))
	if flag.NFlag() < 5 ||
		strings.Trim(*args.GrafanaURL, " ") == "" ||
		strings.Trim(*args.GrafanaPort, " ") == "" ||
		strings.Trim(*args.InfluxURL, " ") == "" ||
		strings.Trim(*args.OrgId, " ") == "" ||
		strings.Trim(*args.ApiKey, " ") == "" {
		fmt.Println("Invalid statup agruments...")
		fmt.Println("Usage: ./intelegrafana -grafanUrl=<URL> -grafanaPort=<port> -influxUrl=<URL> -influxPort=<port> orgId=<orgid> -apiKey=<apikey>")
		fmt.Println("Exiting now...")
		fmt.Println("Bye!!!")
		os.Exit(1)
	}
	Args = args
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

func hostname() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter hostname: ")
	input, _ := reader.ReadString('\n')
	count := 1
	for ; count < 3; {
		if input == "\n" || input == "" || strings.Trim(input, " ") == "" {
			fmt.Println("Please enter a valid hostname: ")
			input, _ = reader.ReadString('\n')
		} else {
			DashboardInfo.Hostname = strings.Trim(strings.Trim(input, " "), "\n")
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
	var payload map[string]map[string]interface{}
	var rows []map[string]interface{}
	var djson string = DashboardJson
	djson = strings.Replace(djson, "$DASHBOARD_NAME$", DashboardInfo.DashboardName, -1)
	djson = strings.Replace(djson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	row := createURLRow()
	if row != nil {
		rows = append(rows, row)
	}
	row = createPortsRow()
	if row != nil {
		rows = append(rows, row)
	}
	row = createProcstatRow()
	if row != nil {
		rows = append(rows, row)
	}
	row = createDiskProcessSwapRow()
	if row != nil {
		rows = append(rows, row)
	}
	row = createCPURAMRow()
	if row != nil {
		rows = append(rows, row)
	}
	row = createIPSystemLoadRow()
	if row != nil {
		rows = append(rows, row)
	}
	json.Unmarshal([]byte(djson), &payload)
	dashboard := payload["dashboard"]
	dashboard["rows"] = rows
	// creating datasource
	createDatasource()

	// creating dashboard
	byts, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		fmt.Println("Error occurred while parsing dashboard json.")
		fmt.Println(err)
		ExitWithStatus1("Json parsing error")
	}
	//fmt.Println(string(byts))
	statusCode, resp, err := httpPost(fmt.Sprintf("%s/api/dashboards/db", *Args.GrafanaURL),
		authHeader(*Args.ApiKey), nil, bytes.NewReader(byts))
	if err != nil {
		fmt.Println("Some error occurred while crating dashboard:", err)
		ExitWithStatus1("Error creating dashboard.")
	}
	m := make(map[string]interface{})
	json.Unmarshal(resp, &m)
	if statusCode == http.StatusOK && m["status"] == "success" {
		fmt.Printf("Dashboard with name %s created successfully.", DashboardInfo.DashboardName)
	} else if statusCode == http.StatusPreconditionFailed {
		fmt.Printf("Dashboard with name %s already exists.", DashboardInfo.DashboardName)
	} else {
		ExitWithStatus1(fmt.Sprintf("Invalid status code found: %d.", statusCode))
	}
}

func createAlertDashboard() {
	dashboardName := DashboardInfo.DashboardName + " Alerts"
	var payload map[string]map[string]interface{}
	var rows []map[string]interface{}
	var djson string = AlertDashboardJson
	djson = strings.Replace(djson, "$DASHBOARD_NAME$", dashboardName, -1)
	djson = strings.Replace(djson, "$TAG$", DashboardInfo.Hostname, -1)
	djson = strings.Replace(djson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	row := createAllAlertsRow()
	if row != nil {
		rows = append(rows, row)
	}
	row = createURLAlertRow()
	if row != nil {
		rows = append(rows, row)
	}
	row = createPortsAlertRow()
	if row != nil {
		rows = append(rows, row)
	}
	row = createProcstatAlertRow()
	if row != nil {
		rows = append(rows, row)
	}
	row = createDiskInodeAlertRow()
	if row != nil {
		rows = append(rows, row)
	}
	row = createCPURAMAlertRow()
	if row != nil {
		rows = append(rows, row)
	}
	row = createSystemLoadAlertRow()
	if row != nil {
		rows = append(rows, row)
	}
	json.Unmarshal([]byte(djson), &payload)
	dashboard := payload["dashboard"]
	dashboard["rows"] = rows

	// creating dashboard
	byts, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		fmt.Println("Error occurred while parsing dashboard json.")
		fmt.Println(err)
		ExitWithStatus1("Json parsing error")
	}
	//fmt.Println(string(byts))
	statusCode, resp, err := httpPost(fmt.Sprintf("%s/api/dashboards/db", *Args.GrafanaURL),
		authHeader(*Args.ApiKey), nil, bytes.NewReader(byts))
	if err != nil {
		fmt.Println("Some error occurred while crating dashboard:", err)
		ExitWithStatus1("Error creating dashboard.")
	}
	m := make(map[string]interface{})
	json.Unmarshal(resp, &m)
	if statusCode == http.StatusOK && m["status"] == "success" {
		fmt.Printf("Dashboard with name %s created successfully.", dashboardName)
	} else if statusCode == http.StatusPreconditionFailed {
		fmt.Printf("Dashboard with name %s already exists.", dashboardName)
	} else {
		ExitWithStatus1(fmt.Sprintf("Invalid status code found: %d.", statusCode))
	}
}

type GrafanaDatasource struct {
	OrgId    int `json:"orgId"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Access   string `json:"access"`
	Url      string `json:"url"`
	Database string `json:"database"`
}

func createDatasource() {
	orgId, err := strconv.Atoi(*Args.OrgId)
	if err != nil {
		ExitWithStatus1("Error while converting orgid")
	}
	ds := &GrafanaDatasource{
		OrgId:    orgId,
		Name:     DashboardInfo.DatasourceName,
		Type:     "influxdb",
		Access:   "proxy",
		Database: DashboardInfo.DatasourceName,
		Url:      fmt.Sprintf("%s:%s", *Args.InfluxURL, *Args.InfluxPort),

	}
	byts, err := json.Marshal(ds)
	if err != nil {
		ExitWithStatus1(fmt.Sprintf("Some error occured while creating grafana datasource."))
	}
	statusCode, resp, err := httpPost(fmt.Sprintf("%s/api/datasources", *Args.GrafanaURL),
		authHeader(*Args.ApiKey), nil, bytes.NewReader(byts))
	if err != nil {
		ExitWithStatus1(fmt.Sprintf("Some error occured while creating grafana datasource."))
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(resp, &m)
	if err != nil {
		ExitWithStatus1(fmt.Sprintf("Some error occured while creating grafana datasource."))
	}
	if statusCode == http.StatusOK || statusCode == http.StatusConflict {
		fmt.Printf("Datasource with name %s created/exists. StatusCode: %d", DashboardInfo.DatasourceName, statusCode)
	} else {
		msg := fmt.Sprintf("Some error occurred whie creating datasource with name %s.", DashboardInfo.DatasourceName)
		ExitWithStatus1(msg)
	}
}

func ExitWithStatus1(msg string) {
	fmt.Println(msg)
	fmt.Println("Exiting now...")
	fmt.Println("Bye!!!")
	os.Exit(1)
}

func createURLRow() map[string]interface{} {
	if DashboardInfo.URLs != "" {
		urls := strings.Split(DashboardInfo.URLs, ",")
		var panels []map[string]interface{}
		for _, p := range urls {
			var paneljson string = URLPanel
			paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
			paneljson = strings.Replace(paneljson, "$URL$", p, -1)
			var panel map[string]interface{}
			json.Unmarshal([]byte(paneljson), &panel)
			panel["id"] = rand.Intn(1000)
			panels = append(panels, panel)

			paneljson = PageLoadPanel
			paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
			paneljson = strings.Replace(paneljson, "$URL$", p, -1)
			panel = make(map[string]interface{})
			json.Unmarshal([]byte(paneljson), &panel)
			panel["id"] = rand.Intn(1000)
			panels = append(panels, panel)

			var buff bytes.Buffer
			buff.WriteString(AgentConfig)
			buff.WriteString("\n\n")
			buff.WriteString(strings.Replace(TelegrafInputHttpResponse, "$URL$", p, -1))
			AgentConfig = buff.String()
		}
		row := newRow("URL Monitoring", 80)
		row["panels"] = panels
		return row
	}
	return nil
}

func createAllAlertsRow() map[string]interface{} {
	var panels []map[string]interface{}
	var panel map[string]interface{}
	json.Unmarshal([]byte(OKAlertPanel), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	panel = make(map[string]interface{})
	json.Unmarshal([]byte(PausedAlertPanel), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	panel = make(map[string]interface{})
	json.Unmarshal([]byte(AlertingPanel), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	panel = make(map[string]interface{})
	json.Unmarshal([]byte(NoDateAlertPanel), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	panel = make(map[string]interface{})
	json.Unmarshal([]byte(ExecutionErrorAlertPanel), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	row := newRow("Alerts List", 300)
	row["panels"] = panels
	return row
}

func createURLAlertRow() map[string]interface{} {
	if DashboardInfo.URLs != "" {
		urls := strings.Split(DashboardInfo.URLs, ",")
		var panels []map[string]interface{}
		for _, p := range urls {
			var paneljson string = URLAlertPanel
			paneljson = strings.Replace(paneljson, "$HOSTNAME$", DashboardInfo.Hostname, -1)
			paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
			paneljson = strings.Replace(paneljson, "$URL$", p, -1)
			var panel map[string]interface{}
			json.Unmarshal([]byte(paneljson), &panel)
			panel["id"] = rand.Intn(1000)
			panels = append(panels, panel)
		}
		row := newRow("URL Alerts", 230)
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
			paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
			paneljson = strings.Replace(paneljson, "$PORT$", p, -1)
			var panel map[string]interface{}
			json.Unmarshal([]byte(paneljson), &panel)
			panel["id"] = rand.Intn(1000)
			panels = append(panels, panel)

			var buff bytes.Buffer
			buff.WriteString(AgentConfig)
			buff.WriteString("\n\n")
			buff.WriteString(strings.Replace(TelegrafInputNetResponse, "$PORT$", p, -1))
			AgentConfig = buff.String()
		}
		row := newRow("Ports Monitoring", 63)
		row["panels"] = panels
		return row
	}
	return nil
}

func createPortsAlertRow() map[string]interface{} {
	if DashboardInfo.Ports != "" {
		ports := strings.Split(DashboardInfo.Ports, ",")
		var panels []map[string]interface{}
		for _, p := range ports {
			var paneljson string = PortAlertPanel
			paneljson = strings.Replace(paneljson, "$HOSTNAME$", DashboardInfo.Hostname, -1)
			paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
			paneljson = strings.Replace(paneljson, "$PORT$", p, -1)
			var panel map[string]interface{}
			json.Unmarshal([]byte(paneljson), &panel)
			panel["id"] = rand.Intn(1000)
			panels = append(panels, panel)
		}
		row := newRow("Ports Alert", 230)
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
			paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
			paneljson = strings.Replace(paneljson, "$PROCESS$", p, -1)
			var panel map[string]interface{}
			json.Unmarshal([]byte(paneljson), &panel)
			panel["id"] = rand.Intn(1000)
			panels = append(panels, panel)

			var buff bytes.Buffer
			buff.WriteString(AgentConfig)
			buff.WriteString("\n\n")
			buff.WriteString(strings.Replace(TelegrafInputProcstat, "$PROCESS$", p, -1))
			AgentConfig = buff.String()
		}
		row := newRow("Process monitoring", 63)
		row["panels"] = panels
		return row
	}
	return nil
}

func createProcstatAlertRow() map[string]interface{} {
	if DashboardInfo.Process != "" {
		ports := strings.Split(DashboardInfo.Process, ",")
		var panels []map[string]interface{}
		for _, p := range ports {
			var paneljson string = ProcstatPanel
			paneljson = strings.Replace(paneljson, "$HOSTNAME$", DashboardInfo.Hostname, -1)
			paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
			paneljson = strings.Replace(paneljson, "$PROCESS$", p, -1)
			var panel map[string]interface{}
			json.Unmarshal([]byte(paneljson), &panel)
			panel["id"] = rand.Intn(1000)
			panels = append(panels, panel)
		}
		row := newRow("Process Alerts", 230)
		row["panels"] = panels
		return row
	}
	return nil
}

func createDiskProcessSwapRow() map[string]interface{} {
	var panels []map[string]interface{}
	var paneljson string = DiskPanel
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	var panel map[string]interface{}
	json.Unmarshal([]byte(paneljson), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	paneljson = ProcessPanel
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	panel = make(map[string]interface{})
	json.Unmarshal([]byte(paneljson), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	paneljson = Swap
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	panel = make(map[string]interface{})
	json.Unmarshal([]byte(paneljson), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	row := newRow("Disk, Process & Swap Monitoring", 230)
	row["panels"] = panels
	return row
}

func createDiskInodeAlertRow() map[string]interface{} {
	var panels []map[string]interface{}
	var paneljson string = DiskAlertPanel
	paneljson = strings.Replace(paneljson, "$HOSTNAME$", DashboardInfo.Hostname, -1)
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	var panel map[string]interface{}
	json.Unmarshal([]byte(paneljson), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	paneljson = INodeAlertPanel
	paneljson = strings.Replace(paneljson, "$HOSTNAME$", DashboardInfo.Hostname, -1)
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	panel = make(map[string]interface{})
	json.Unmarshal([]byte(paneljson), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	row := newRow("Disk & Inode Alerts", 230)
	row["panels"] = panels
	return row
}

func createCPURAMRow() map[string]interface{} {
	var panels []map[string]interface{}
	var paneljson string = CPUPanel
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	var panel map[string]interface{}
	json.Unmarshal([]byte(paneljson), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	paneljson = RAMPanel
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	panel = make(map[string]interface{})
	json.Unmarshal([]byte(paneljson), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	row := newRow("CPU & RAM Monitoring", 230)
	row["panels"] = panels
	return row
}

func createCPURAMAlertRow() map[string]interface{} {
	var panels []map[string]interface{}
	var paneljson string = CPUAlertPanel
	paneljson = strings.Replace(paneljson, "$HOSTNAME$", DashboardInfo.Hostname, -1)
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	var panel map[string]interface{}
	json.Unmarshal([]byte(paneljson), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	paneljson = RAMAlertPanel
	paneljson = strings.Replace(paneljson, "$HOSTNAME$", DashboardInfo.Hostname, -1)
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	panel = make(map[string]interface{})
	json.Unmarshal([]byte(paneljson), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	row := newRow("CPU & RAM Alerts", 230)
	row["panels"] = panels
	return row
}

func createIPSystemLoadRow() map[string]interface{} {
	var panels []map[string]interface{}
	var paneljson string = IPTrafficPanel
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	var panel map[string]interface{}
	json.Unmarshal([]byte(paneljson), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	paneljson = SystemLoadPanel
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	panel = make(map[string]interface{})
	json.Unmarshal([]byte(paneljson), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	row := newRow("IP Traffic & System load Monitoring", 230)
	row["panels"] = panels
	return row
}

func createSystemLoadAlertRow() map[string]interface{} {
	var panels []map[string]interface{}
	var paneljson string = SystemLoadAlertPanel
	paneljson = strings.Replace(paneljson, "$HOSTNAME$", DashboardInfo.Hostname, -1)
	paneljson = strings.Replace(paneljson, "$DATASOURCE_NAME$", DashboardInfo.DatasourceName, -1)
	var panel map[string]interface{}
	json.Unmarshal([]byte(paneljson), &panel)
	panel["id"] = rand.Intn(1000)
	panels = append(panels, panel)

	row := newRow("System load Alerts", 230)
	row["panels"] = panels
	return row
}

func newRow(name string, height int) map[string]interface{} {
	return map[string]interface{}{
		"collapse":        false,
		"repeat":          nil,
		"repeatIteration": nil,
		"repeatRowId":     nil,
		"showTitle":       false,
		"titleSize":       "h6",
		"title":           name,
		"height":          height,
	}
}

func httpPost(url string, headers map[string]string, params map[string]string, body io.Reader) (int, []byte, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return 0, nil, err
	}
	setHeaders(req, headers)
	query := setParams(req, params)
	req.URL.RawQuery = query.Encode()
	fmt.Println("POST Request :" + req.URL.String())
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, nil, err
	}
	defer resp.Body.Close()
	read, err := ioutil.ReadAll(resp.Body)
	return resp.StatusCode, read, err
}

func authHeader(key string) map[string]string {
	return map[string]string{"Authorization": fmt.Sprintf("Bearer %s", key)}
}

func setHeaders(req *http.Request, headers map[string]string) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}
}

func setParams(req *http.Request, params map[string]string) url.Values {
	query := req.URL.Query()
	if params != nil {
		for k, v := range params {
			query.Set(k, fmt.Sprintf("%v", v))
		}
	}
	return query
}

const (
	UBUNTU = "UBUNTU"
	CENTOS = "CENTOS"
)

func getOSName() string {
	output, err := exec.Command("awk", "-F=", "/^ID=/{print $2}", "/etc/os-release").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		ExitWithStatus1(err.Error())
	}
	platform := strings.Trim(string(output), "\n")
	if platform == "ubuntu" || strings.Contains(platform, "ubuntu"){
		return UBUNTU
	} else if platform == "centos" || strings.Contains(platform, "centos") {
		return CENTOS
	} else {
		ExitWithStatus1("Only ubuntu and centos are supported")
	}
	return ""
}

func getDistributionName() string {
	output, err := exec.Command("awk", "-F=", "/^DISTRIB_CODENAME=/{print $2}", "/etc/lsb-release").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		ExitWithStatus1(err.Error())
	}
	return strings.Trim(string(output), "\n")
}

func installTelegraf(platform string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Do you want to install telegraf [y/n]: ")
	input, _ := reader.ReadString('\n')
	input = strings.ToUpper(input)
	isYesNo := string([]byte(input)[0])
	if isYesNo == "Y" {
		fmt.Println(fmt.Sprintf("Installing telegraf for %s", platform))
		switch platform {
		case UBUNTU:
			_, err := exec.Command("wget", "https://dl.influxdata.com/telegraf/releases/telegraf_1.5.2-1_amd64.deb", "-P", "/tmp").CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(err.Error())
				ExitWithStatus1(err.Error())
			}
			_, err = exec.Command("sudo", "dpkg", "-i", "/tmp/telegraf_1.5.2-1_amd64.deb").CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(err.Error())
				ExitWithStatus1(err.Error())
			}
			_, err = exec.Command("sudo", "rm", "-f", "/tmp/telegraf_1.5.2-1_amd64.deb").CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(err.Error())
				ExitWithStatus1(err.Error())
			}
			replaceTelegrafBinary()
			break
		case CENTOS:
			_, err := exec.Command("wget", "https://dl.influxdata.com/telegraf/releases/telegraf-1.5.2-1.x86_64.rpm", "-P", "/tmp").CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(err.Error())
				ExitWithStatus1(err.Error())
			}
			_, err = exec.Command("sudo", "yum", "localinstall", "/tmp/telegraf-1.5.2-1.x86_64.rpm").CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(err.Error())
				ExitWithStatus1(err.Error())
			}
			_, err = exec.Command("sudo", "rm", "-f", "/tmp/telegraf-1.5.2-1.x86_64.rpm").CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(err.Error())
				ExitWithStatus1(err.Error())
			}
			replaceTelegrafBinary()
			_, err = exec.Command("sudo", "chkconfig", "telegraf", "on").CombinedOutput()
			if err != nil {
				os.Stderr.WriteString(err.Error())
				ExitWithStatus1(err.Error())
			}
			break
		}
	}
}

func replaceTelegrafBinary() {
	_, err := exec.Command("wget", "https://s3.ap-south-1.amazonaws.com/girnar-telegraf/telegraf", "-P", "/tmp").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		ExitWithStatus1(err.Error())
	}
	_, err = exec.Command("sudo", "chmod", "775", "/tmp/telegraf").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		ExitWithStatus1(err.Error())
	}
	_, err = exec.Command("sudo", "mv", "-f", "/tmp/telegraf", "/usr/bin").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		ExitWithStatus1(err.Error())
	}
}

func changeTelegrafConfPermission() {
	_, err := exec.Command("sudo", "chmod", "777", "/etc/telegraf/telegraf.conf").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		ExitWithStatus1(err.Error())
	}
}

func replaceTelegrafConfig() {
	ioutil.WriteFile("/etc/telegraf/telegraf.conf", []byte(AgentConfig), 0775)
}

func restartTelegraf() {
	_, err := exec.Command("sudo", "service", "telegraf", "restart").CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		ExitWithStatus1(err.Error())
	}
}
