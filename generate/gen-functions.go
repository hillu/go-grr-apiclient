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

var paramRegexp = regexp.MustCompile(`(/[a-z-/]+/?)((?:<.*?>)?)`)

type param struct{ Prefix, Name string }

func getParams(path string) (params []param) {
	for _, f := range paramRegexp.FindAllStringSubmatch(path, -1) {
		p := param{Prefix: f[1]}
		if len(f[2]) != 0 {
			// get rid of prefix:
			name := f[2][1 : len(f[2])-1]
			n := strings.Split(name, ":")
			name = n[len(n)-1]
			p.Name = snake2camel(name)
		}
		params = append(params, p)
	}
	return
}

var (
	requesttemplate = template.Must(template.New("").
			Funcs(template.FuncMap{"title": strings.Title}).
			Parse(`
{{- $arg := or .Arg ( print "Api" .Name "Args" ) -}}
{{- $result := or .Result ( print "Api" .Name "Result" ) -}}
func (c *APIClient) {{ .Name -}} (rq {{ $arg -}} ) (rs *{{ $result }}, err error) {
	rs = new( {{- $result -}} )
	if err := c.do("{{ .Method }}", {{ range $index, $var := .Params -}}
{{- if $index -}}+{{- end -}}
"{{- $var.Prefix -}}"
{{- with $var.Name }}+path.Base(rq.Get{{ title . }}()){{ end -}}
{{- end -}}

, &rq, rs); err != nil {
		return nil, err
	}
	return
}
`,
		))
	gettemplate = template.Must(template.New("").Parse(`
{{- $result := or .Result ( print "Api" .Name "Result" ) -}}
func (c *APIClient) {{ .ApiCall.Name -}} (

{{- range $index, $var := .Params -}}
{{- if and $index $var.Name }}, {{ end -}}
{{- with $var.Name }} {{- .}} string {{- end -}}
{{- end -}}

) (rs *{{ $result }}, err error) {
	rs = new( {{- $result -}} )
	if err := c.get(

{{- range $index, $var := .Params -}}
{{- if $index -}}+{{- end -}}
"{{- $var.Prefix -}}"
{{- with $var.Name }}+{{ . }}{{ end -}}
{{- end -}}

, nil, rs); err != nil {
		return nil, err
	}
	return
}
`,
	))
	posttemplate = template.Must(template.New("").Parse(`
{{- $arg := or .Arg ( print "Api" .Name "Args" ) -}}
func (c *APIClient) {{ .ApiCall.Name -}} (rq {{ $arg }}

{{- range $index, $var := .Params -}}
{{- with $var.Name }}, {{ . }} string {{- end -}}
{{- end -}}

) error {
	return c.post(

{{- range $index, $var := .Params -}}
{{- if $index }}+{{ end -}}
"{{- $var.Prefix -}}"
{{- with $var.Name }}+{{ . }}{{ end -}}
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
	{Name: "InterrogateClient", Method: "POST", Path: "/api/clients/<client_id>/actions/interrogate"},
	// FIXME: client_id
	// {Name: "GetInterrogateOperationState", Method: "GET", Path: "/api/clients/<client_id>/actions/interrogate/<path:operation_id>"},
	{Name: "GetLastClientIPAddress", Method: "GET-simple", Path: "/api/clients/<client_id>/last-ip"},

	// Virtual file system methods.
	{Name: "ListFiles", Method: "GET", Path: "/api/clients/<client_id>/vfs-index/"},
	{Name: "GetFileDetails", Method: "GET", Path: "/api/clients/<client_id>/vfs-details/<path:file_path>"},
	{Name: "GetFileText", Method: "GET", Path: "/api/clients/<client_id>/vfs-text/<path:file_path>"},
	// FIXME: Binary Stream
	// {Name: "GetFileBlob", Method: "GET", Path: "/api/clients/<client_id>/vfs-blob/<path:file_path>"},
	{Name: "GetFileVersionTimes", Method: "GET", Path: "/api/clients/<client_id>/vfs-version-times/<path:file_path>"},
	{Name: "GetFileDownloadCommand", Method: "GET", Path: "/api/clients/<client_id>/vfs-download-command/<path:file_path>"},
	{Name: "CreateVfsRefreshOperation", Method: "POST", Path: "/api/clients/<client_id>/vfs-refresh-operations"},
	// FIXME: client_id
	// {Name: "GetVfsRefreshOperationState", Method: "GET", Path: "/api/clients/<client_id>/vfs-refresh-operations/<path:operation_id>"},
	{Name: "GetVfsTimeline", Method: "GET", Path: "/api/clients/<client_id>/vfs-timeline/<path:file_path>"},
	// FIXME: Binary Stream
	// {Name: "GetVfsTimelineAsCsv", Method: "GET", Path: "/api/clients/<client_id>/vfs-timeline-csv/<path:file_path>"},
	{Name: "UpdateVfsFileContent", Method: "POST", Path: "/api/clients/<client_id>/vfs-update"},
	// FIXME: client_id
	// {Name: "GetVfsFileContentUpdateState", Method: "GET", Path: "/api/clients/<client_id>/vfs-update/<path:operation_id>"},

	// Clients labels methods.
	{Name: "ListClientsLabels", Method: "GET-simple", Path: "/api/clients/labels"},
	{Name: "AddClientsLabels", Method: "POST-simple", Path: "/api/clients/labels/add"},
	{Name: "RemoveClientsLabels", Method: "POST-simple", Path: "/api/clients/labels/remove"},

	// Clients flows methods.
	{Name: "ListFlows", Method: "GET", Path: "/api/clients/<client_id>/flows"},
	{Name: "GetFlow", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>", Result: "ApiFlow"},
	{Name: "CreateFlow", Method: "POST", Path: "/api/clients/<client_id>/flows", Result: "ApiFlow"},
	{Name: "CancelFlow", Method: "POST-simple", Path: "/api/clients/<client_id>/flows/<flow_id>/actions/cancel"},
	{Name: "ListFlowResults", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/results"},
	{Name: "GetFlowResultsExportCommand", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/results/export-command"},
	// FIXME: Binary Stream
	// {Name: "GetFlowFilesArchive", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/results/files-archive"},
	{Name: "ListFlowOutputPlugins", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/output-plugins"},
	{Name: "ListFlowOutputPluginLogs", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/output-plugins/<plugin_id>/errors"},
	{Name: "ListFlowOutputPluginErrors", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/output-plugins/<plugin_id>/errors"},
	{Name: "ListFlowLogs", Method: "GET", Path: "/api/clients/<client_id>/flows/<flow_id>/log"},

	// Global flows methods.
	{Name: "CreateGlobalFlow", Method: "POST", Path: "/api/flows", Arg: "ApiCreateFlowArgs", Result: "ApiFlow"},

	// Cron jobs methods.
	{Name: "ListCronJobs", Method: "GET", Path: "/api/cron-jobs"},
	{Name: "CreateCronJob", Method: "POST-simple", Path: "/api/cron-jobs", Arg: "ApiCronJob", Result: "ApiCronJob"},
	{Name: "DeleteCronJob", Method: "POST-simple", Path: "/api/cron-jobs/<cron_job_id>/actions/delete"},

	// Hunts methods.
	{Name: "ListHunts", Method: "GET", Path: "/api/hunts"},
	{Name: "GetHunt", Method: "GET-simple", Path: "/api/hunts/<hunt_id>", Result: "ApiHunt"},
	{Name: "ListHuntErrors", Method: "GET", Path: "/api/hunts/<hunt_id>/errors"},
	{Name: "ListHuntLogs", Method: "GET", Path: "/api/hunts/<hunt_id>/log"},
	{Name: "ListHuntResults", Method: "GET", Path: "/api/hunts/<hunt_id>/results"},
	{Name: "GetHuntResultsExportCommand", Method: "GET", Path: "/api/hunts/<hunt_id>/results/export-command"},
	{Name: "ListHuntOutputPlugins", Method: "GET", Path: "/api/hunts/<hunt_id>/output-plugins"},
	{Name: "ListHuntOutputPluginLogs", Method: "GET", Path: "/api/hunts/<hunt_id>/output-plugins/<plugin_id>/logs"},
	{Name: "ListHuntOutputPluginErrors", Method: "GET", Path: "/api/hunts/<hunt_id>/output-plugins/<plugin_id>/errors"},
	{Name: "ListHuntCrashes", Method: "GET", Path: "/api/hunts/<hunt_id>/crashes"},
	{Name: "GetHuntClientCompletionStats", Method: "GET", Path: "/api/hunts/<hunt_id>/client-completion-stats"},
	{Name: "GetHuntStats", Method: "GET", Path: "/api/hunts/<hunt_id>/stats"},
	// FIXME: ApiListHuntClientsArgs.ClientStatus is not a string
	// {Name: "ListHuntClients", Method: "GET", Path: "/api/hunts/<hunt_id>/clients/<client_status>"},
	{Name: "GetHuntContext", Method: "GET", Path: "/api/hunts/<hunt_id>/context"},
	{Name: "CreateHunt", Method: "POST", Path: "/api/hunts", Result: "ApiHunt"},
	// FIXME: Binary Stream
	// {Name: "GetHuntFilesArchive", Method: "GET", Path: "/api/hunts/<hunt_id>/results/clients/<client_id>/fs/<path:vfs_path>"},

	// Stats metrics methods.
	// FIXME: unknown result type
	// {Name: "ListStatsStoreMetricsMetadata", Method: "GET", Path: "/api/stats/store/<component>/metadata"},
	// FIXME: unknown result type
	// {Name: "GetStatsStoreMetric", Method: "GET", Path: "/api/stats/store/<component>/metrics/<metric_name>"},

	// Approvals methods.
	{Name: "CreateUserClientApproval", Method: "POST", Path: "/api/users/me/approvals/client/<client_id>", Result: "ApiUserClientApproval"},
	{Name: "GetUserClientApproval", Method: "GET", Path: "/api/users/me/approvals/client/<client_id>/<reason>", Result: "ApiUserClientApproval"},
	{Name: "ListUserClientApprovals", Method: "GET", Path: "/api/users/me/approvals/client"},
	{Name: "ListUserHuntApprovals", Method: "GET", Path: "/api/users/me/approvals/hunt"},
	{Name: "ListUserCronApprovals", Method: "GET", Path: "/api/users/me/approvals/cron"},

	// User settings methods.
	{Name: "GetPendingUserNotificationsCount", Method: "GET-simple", Path: "/api/users/me/notifications/pending/count"},
	{Name: "ListPendingUserNotifications", Method: "GET", Path: "/api/users/me/notifications/pending"},
	// FIXME
	// {Name: "DeletePendingUserNotification", Method: "Delete", Path: "/api/users/me/notifications/pending/<timestamp>"},
	{Name: "ListAndResetUserNotifications", Method: "POST", Path: "/api/users/me/notifications"},
	{Name: "GetGrrUser", Method: "GET-simple", Path: "/api/users/me", Result: "ApiGrrUser"},
	{Name: "UpdateGrrUser", Method: "POST-simple", Path: "/api/users/me", Arg: "ApiGrrUser"},
	{Name: "ListPendingGlobalNotifications", Method: "GET-simple", Path: "/api/users/me/notifications/pending/global"},
	// FIXME: DELETE-simple?
	// {Name: "DeletePendingGlobalNotification", Method: "DELETE", Path: "/api/users/me/notifications/pending/global/<type>"},

	// Config methods.
	{Name: "GetConfig", Method: "GET-simple", Path: "/api/config"},
	{Name: "GetConfigOption", Method: "GET", Path: "/api/config/<name>", Result: "ApiConfigOption"},

	// Reflection methods.
	{Name: "ListKbFields", Method: "GET-simple", Path: "/api/clients/kb-fields"},
	{Name: "ListFlowDescriptors", Method: "GET", Path: "/api/flows/descriptors"},
	// FIXME: unknown type
	// {Name: "ListAff4AttributeDescriptors", Method: "GET", Path: "/api/reflection/aff4/attributes"},
	// FIXME: type is a reserved word, Type is not.
	// {Name: "GetRDFValueDescriptor", Method: "GET-simple", Path: "/api/reflection/rdfvalue/<type>"},
	// FIXME: unknown type
	// {Name: "ListRDFValuesDescriptors", Method: "GET-simple", Path: "/api/reflection/rdfvalue/all"},
	// FIXME: unknown type
	// {Name: "ListOutputPluginDescriptors", Method: "GET-simple", Path: "/api/output-plugins/all"},
	{Name: "ListKnownEncodings", Method: "GET-simple", Path: "/api/reflection/file-encodings"},

	// Documentation methods.
	// {Name: "GetDocs", Method: "GET", Path: "/api/docs"},

	// Robot methods.
	{Name: "StartRobotGetFilesOperation", Method: "POST", Path: "/api/robot-actions/get-files"},
	{Name: "GetRobotGetFilesOperationState", Method: "GET", Path: "/api/robot-actions/get-files/<path:operation_id>"},
}

func main() {
	f, err := os.Create("apifunctions.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(f, `// This file has been autogenerated by gen-functions.go

package apiclient

import (
	"path"
)

`)
	defer f.Close()
	for _, call := range apicalls {
		var err error
		switch call.Method {
		case "GET", "POST":
			err = requesttemplate.Execute(f, struct {
				ApiCall
				Params []param
			}{call, getParams(call.Path)})
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
