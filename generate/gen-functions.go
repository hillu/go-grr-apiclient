package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"text/template"
)

func snake2camel(name string) (result string) {
	for i, f := range strings.Split(name, "_") {
		if i == 0 {
			result = result + f
		} else {
			result = result + strings.Title(f)
		}
	}
	return
}

var paramRegexp = regexp.MustCompile(`(/[a-z/]+/?)((?:<.*?>)?)`)

type param struct{ Prefix, Name string }

func getParams(path string) (params []param) {
	for _, f := range paramRegexp.FindAllStringSubmatch(path, -1) {
		p := param{Prefix: f[1]}
		if len(f[2]) != 0 {
			p.Name = snake2camel(f[2][1 : len(f[2])-1])
		}
		params = append(params, p)
	}
	return
}

var (
	requesttemplate = template.Must(template.New("").Parse(`
func (c APIClient) {{ .Name -}} (rq Api{{ .Name }}Args) (rs *Api{{ .Name }}Result, err error) {
	rs = new(Api {{- .Name -}} Result)
	if err := c.do("{{ .Method }}", "{{ .Path }}", &rq, rs); err != nil {
		return nil, err
	}
	return
}
`,
	))
	gettemplate = template.Must(template.New("").Parse(`
func (c APIClient) {{ .ApiCall.Name -}} (

{{- range $index, $var := .Params -}}
{{- if $index }}, {{ end -}}
{{- with $var.Name }} {{- .}} string {{- end -}}
{{- end -}}

) (rs *Api {{- .ApiCall.Name -}} Result, err error) {
	rs = new(Api{{ .ApiCall.Name }}Result)
	if err := c.get(

{{- range $index, $var := .Params -}}
{{- if $index }} + {{ end -}}
"{{- $var.Prefix -}}"
{{- with $var.Name }} + {{ . }}{{ end -}}
{{- end -}}

, nil, rs); err != nil {
		return nil, err
	}
	return
}
`,
	))
	posttemplate = template.Must(template.New("").Parse(`
func (c APIClient) {{ .ApiCall.Name -}} (rq Api{{ .ApiCall.Name }}Args

{{- range $index, $var := .Params -}}
{{- with $var.Name }}, {{ . }} string {{- end -}}
{{- end -}}

) error {
	return c.post(

{{- range $index, $var := .Params -}}
{{- if $index }} + {{ end -}}
"{{- $var.Prefix -}}"
{{- with $var.Name }} + {{ . }}{{ end -}}
{{- end -}}

, &rq)
}
`,
	))
)

type method int

type queryarg struct {
	Name, Type string
}

/*
All definitions here have been derived from the functions in
grr/gui/api_call_router.go that are annotated with the @Http tag.
*/
type ApiCall struct {
	Name, Method string
	Path         string
	Arg, Result  string
	Queryargs    []queryarg
}

var apicalls = []ApiCall{

	// Artifacts methods.
	{Name: "ListArtifacts", Method: "GET", Path: "/api/artifacts"},
	{Name: "UploadArtifact", Method: "POST-simple", Path: "/api/artifacts/upload"},
	{Name: "DeleteArtifacts", Method: "POST-simple", Path: "/api/artifacts/delete"},

	// Clients methods.
	{Name: "SearchClients", Method: "GET", Path: "/api/clients"},
	{Name: "GetClient", Method: "GET", Path: "/api/clients/<client_id>", Queryargs: []queryarg{{"timestamp", "int64"}}},
	{Name: "GetClientVersionTimes", Method: "GET-simple", Path: "/api/clients/<client_id>/version-times"},
	// {Name: "InterrogateClient", Method: "POST", Path: "/api/clients/<client_id>/actions/interrogate"},
	// {Name: "GetInterrogateOperationState", Method: "GET", Path: "/api/clients/<client_id>/actions/interrogate/<path:operation_id>"},
	{Name: "GetLastClientIPAddress", Method: "GET-simple", Path: "/api/clients/<client_id>/last-ip"},

	// Virtual file system methods.
	// {Name: "ListFiles", Method: "GET", Path: "/api/clients/<client_id>/vfs-index/"},
	// {Name: "GetFileDetails", Method: "GET", Path: "/api/clients/<client_id>/vfs-details/<path:file_path>"},
	// {Name: "GetFileText", Method: "GET", Path: "/api/clients/<client_id>/vfs-text/<path:file_path>"},
	// {Name: "GetFileBlob", Method: "GET", Path: "/api/clients/<client_id>/vfs-blob/<path:file_path>"},
	// {Name: "GetFileVersionTimes", Method: "GET", Path: "/api/clients/<client_id>/vfs-version-times/<path:file_path>"},
	// {Name: "GetFileDownloadCommand", Method: "GET", Path: "/api/clients/<client_id>/vfs-download-command/<path:file_path>"},
	// {Name: "CreateVfsRefreshOperation", Method: "POST", Path: "/api/clients/<client_id>/vfs-refresh-operations"},
	// {Name: "GetVfsRefreshOperationState", Method: "GET", Path: "/api/clients/<client_id>/vfs-refresh-operations/<path:operation_id>"},
	// {Name: "GetVfsTimeline", Method: "Path": Get, Path: "/api/clients/<client_id>/vfs-timeline/<path:file_path>"},
	// {Name: "GetVfsTimelineAsCsv", Method: "GET", Path: "/api/clients/<client_id>/vfs-timeline-csv/<path:file_path>"},
	// {Name: "UpdateVfsFileContent", Method: "POST", Path: "/api/clients/<client_id>/vfs-update"},
	// {Name: "GetVfsFileContentUpdateState", Method: "GET", Path: "/api/clients/<client_id>/vfs-update/<path:operation_id>"},

	// Clients labels methods.
	{Name: "ListClientsLabels", Method: "GET-simple", Path: "/api/clients/labels"},
	{Name: "AddClientsLabels", Method: "POST-simple", Path: "/api/clients/labels/add"},
	{Name: "RemoveClientsLabels", Method: "POST-simple", Path: "/api/clients/labels/remove"},

	// Clients flows methods.
	{Name: "ListFlows", Method: "GET", Path: "/api/clients/<client_id>/flows"},
	// {Name: "GetFlow", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>"},
	// {Name: "CreateFlow", Method: "POST", Path: "/api/clients/<client_id>/flows"},
	{Name: "CancelFlow", Method: "POST-simple", Path: "/api/clients/<client_id>/flows/<flow_id>/actions/cancel"},
	// {Name: "ListFlowResults", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/results"},
	// {Name: "GetFlowResultsExportCommand", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/results/export-command"},
	// {Name: "GetFlowFilesArchive", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/results/files-archive"},
	// {Name: "ListFlowOutputPlugins", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/output-plugins/<plugin_id>/logs"},
	// {Name: "ListFlowOutputPluginLogs", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/output-plugins/<plugin_id>/errors"},
	// {Name: "ListFlowOutputPluginErrors", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/"},
	// {Name: "ListFlowLogs", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/log"},

	// Global flows methods.
	{Name: "CreateFlow", Method: "POST-simple", Path: "/api/flows"},

	// Cron jobs methods.
	{Name: "ListCronJobs", Method: "GET", Path: "/api/cron-jobs"},
	// {Name: "CreateCronJob", Method: "POST-simple", Path: "/api/cron-jobs"}, Fixme: data type
	// {Name: "DeleteCronJob", Method: "POST-simple", Path: "/api/cron-jobs/<cron_job_id>/actions/delete"},

	// Hunts methods.
	{Name: "ListHunts", Method: "GET", Path: "/api/hunts"},
	// {Name: "GetHunt", Method: "GET-simple", Path: "/api/hunts/<hunt_id>"},
	// {Name: "ListHuntErrors", Method: "GET", Path: "/api/hunts/<hunt_id>/errors"},
	// {Name: "ListHuntLogs", Method: "GET", Path: "/api/hunts/<hunt_id>/log"},
	// {Name: "ListHuntResults", Method: "GET", Path: "/api/hunts/<hunt_id>/results"},
	// {Name: "GetHuntResultsExportCommand", Method: "GET", Path: "/api/hunts/<hunt_id>/results/export-command"},
	// {Name: "ListHuntOutputPlugins", Method: "GET", Path: "/api/hunts/<hunt_id>/output-plugins"},
	// {Name: "ListHuntOutputPluginLogs", Method: "GET", Path: "/api/hunts/<hunt_id>/output-plugins/<plugin_id>/logs"},
	// {Name: "ListHuntOutputPluginErrors", Method: "GET", Path: "/api/hunts/<hunt_id>/output-plugins/<plugin_id>/errors"},
	// {Name: "ListHuntCrashes", Method: "GET", Path: "/api/hunts/<hunt_id>/crashes"},
	// {Name: "GetHuntClientCompletionStats", Method: "GET", Path: "/api/hunts/<hunt_id>/client-completion-stats"},
	// {Name: "GetHuntStats", Method: "GET", Path: "/api/hunts/<hunt_id>/stats"},
	// {Name: "ListHuntClients", Method: "GET", Path: "/api/hunts/<hunt_id>/clients/<client_status>"},
	// {Name: "GetHuntContext", Method: "GET", Path: "/api/hunts/<hunt_id>/context"},
	// {Name: "CreateHunt", Method: "POST", Path: "/api/hunts", Result: "ApiHunt"},
	// {Name: "GetHuntFilesArchive", Method: "GET", Path: "/api/hunts/<hunt_id>/results/clients/<client_id>/fs/<path:vfs_path>"},

	// Stats metrics methods.
	// {Name: "ListStatsStoreMetricsMetadata", Method: "GET", Path: "/api/stats/store/<component>/metadata"},
	// {Name: "GetStatsStoreMetric", Method: "GET", Path: "/api/stats/store/<component>/metrics/<metric_name>"},

	// Approvals methods.
	// {Name: "CreateUserClientApproval", Method: "POST", Path: "/api/users/me/approvals/client/<client_id>"},
	// {Name: "GetUserClientApproval", Method: "GET", Path: "/api/users/me/approvals/client/<client_id>/<reason>"},
	// {Name: "ListUserClientApprovals", Method: "GET", Path: "/api/users/me/approvals/client"},
	// {Name: "ListUserHuntApprovals", Method: "GET", Path: "/api/users/me/approvals/hunt"},
	// {Name: "ListUserCronApprovals", Method: "GET", Path: "/api/users/me/approvals/cron"},

	// User settings methods.
	// {Name: "GetPendingUserNotificationsCount", Method: "GET", Path: "/api/users/me/notifications/pending/count"},
	// {Name: "ListPendingUserNotifications", Method: "GET", Path: "/api/users/me/notifications/pending"},
	// {Name: "DeletePendingUserNotification", Method: "Delete", Path: "/api/users/me/notifications/pending/<timestamp>"},
	// {Name: "ListAndResetUserNotifications", Method: "POST", Path: "/api/users/me/notifications"},
	// {Name: "GetGrrUser", Method: "GET", Path: "/api/users/me"},
	// {Name: "UpdateGrrUser", Method: "POST", Path: "/api/users/me"},
	// {Name: "ListPendingGlobalNotifications", Method: "GET", Path: "/api/users/me/notifications/pending/global"},
	// {Name: "DeletePendingGlobalNotification", Method: "Delete", Path: "/api/users/me/notifications/pending/global/<type>"},

	// Config methods.
	// {Name: "GetConfig", Method: "GET", Path: "/api/config"},
	// {Name: "GetConfigOption", Method: "GET", Path: "/api/config/<name>"},

	// Reflection methods.
	// {Name: "ListKbFields", Method: "GET", Path: "/api/clients/kb-fields"},
	// {Name: "ListFlowDescriptors", Method: "GET", Path: "/api/flows/descriptors"},
	// {Name: "ListAff4AttributeDescriptors", Method: "GET", Path: "/api/reflection/aff4/attributes"},
	// {Name: "GetRDFValueDescriptor", Method: "GET", Path: "/api/reflection/rdfvalue/<type>"},
	// {Name: "ListRDFValuesDescriptors", Method: "GET", Path: "/api/reflection/rdfvalue/all"},
	// {Name: "ListOutputPluginDescriptors", Method: "GET", Path: "/api/output-plugins/all"},
	// {Name: "ListKnownEncodings", Method: "GET", Path: "/api/reflection/file-encodings"},

	// Documentation methods.
	// {Name: "GetDocs", Method: "GET", Path: "/api/docs"},

	// Robot methods.
	// {Name: "StartRobotGetFilesOperation", Method: "POST", Path: "/api/robot-actions/get-files"},
	// {Name: "GetRobotGetFilesOperationState", Method: "GET", Path: "/api/robot-actions/get-files/<path:operation_id>"},
}

func main() {
	f, err := os.Create("apifunctions.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(f, `// This file has been autogenerated by gen-functions.go

package apiclient
`)
	defer f.Close()
	for _, call := range apicalls {
		var err error
		switch call.Method {
		case "GET", "POST":
			err = requesttemplate.Execute(f, call)
		case "GET-simple":
			err = gettemplate.Execute(f, struct {
				ApiCall
				Params []param
			}{call, getParams(call.Path)})
		case "POST-simple":
			err = posttemplate.Execute(f, struct {
				ApiCall
				Params []param
			}{call, getParams(call.Path)})
		default:
			log.Fatalf("Error in struct definition for %s", call.Name)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}
