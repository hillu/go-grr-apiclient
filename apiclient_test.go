package apiclient

import (
	"github.com/golang/protobuf/proto"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
	"testing"
)

type examples map[string][]struct {
	Method               string
	URL                  string
	RequestPayload       json.RawMessage
	Response             json.RawMessage
	TestClass            string
	TypeStrippedResponse json.RawMessage `json:"type_stripped_response"`
}

type methodpath struct{ method, path string }

type response struct {
	URL                            *url.URL
	Response, TypeStrippedResponse json.RawMessage
}

type handler struct {
	t        *testing.T
	testdata map[methodpath][]response
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	_, useStripped := query["strip_type_info"]
	for _, s := range h.testdata[methodpath{r.Method, r.URL.Path}] {
		delete(query, "strip_type_info")
		if !reflect.DeepEqual(query, s.URL.Query()) {
			continue
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(")]}'\n"))
		if useStripped {
			h.t.Logf("URL: %s, Output: %s", s.URL, string(s.TypeStrippedResponse))
			w.Write(s.TypeStrippedResponse)
			return
		}
		h.t.Logf("URL: %s, Output: %s", s.URL, string(s.Response))
		w.Write(s.Response)
		return
	}
}

// newHandler reads test_data/api-docs-examples.json (taken from GRR)
// and returns a http.Handler that acts according to the examples.
func newHandler(t *testing.T) http.Handler {
	f, err := os.Open("test_data/api-docs-examples.json")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	e := make(examples)
	h := handler{
		t:        t,
		testdata: make(map[methodpath][]response),
	}
	d := json.NewDecoder(f)
	if err := d.Decode(&e); err != nil {
		t.Fatal(err)
	}
	for _, samples := range e {
		for _, s := range samples {
			u, _ := url.Parse(s.URL)
			mp := methodpath{s.Method, u.Path}
			h.testdata[mp] = append(h.testdata[mp],
				response{u, s.Response, s.TypeStrippedResponse})
		}
	}
	return h
}

func TestApiExamples(t *testing.T) {
	s := httptest.NewServer(newHandler(t))
	defer s.Close()
	u, _ := url.Parse(s.URL)
	c := &APIClient{BaseURL: u}
	for i, test := range []struct {
		f        func() (proto.Message, error)
		expected proto.Message
	}{
		{
			func() (proto.Message, error) {
				return c.SearchClients(ApiSearchClientsArgs{Query: proto.String("C.1000000000000000")})
			},
			&ApiSearchClientsResult{
				Items: []*ApiClient{
					{
						AgentInfo: &ClientInformation{
							ClientName: proto.String("GRR Monitor"),
						},
						HardwareInfo: &HardwareInfo{},
						LastSeenAt:   proto.Uint64(42000000),
						OsInfo: &Uname{
							Fqdn:    proto.String("Host-0.example.com"),
							Node:    proto.String("Host-0"),
							Version: proto.String(""),
						},
						Urn:   proto.String("aff4:/C.1000000000000000"),
						Users: nil,
					},
				},
			}},
		{
			func() (proto.Message, error) {
				return c.ListCronJobs(ApiListCronJobsArgs{})
			},
			&ApiListCronJobsResult{
				Items: []*ApiCronJob{
					{
						AllowOverruns: proto.Bool(false),
						Description:   proto.String("foo"),
						FlowName:      proto.String("GRRVersionBreakDown"),
						FlowRunnerArgs: &FlowRunnerArgs{
							FlowName: proto.String("GRRVersionBreakDown"),
						},
						IsFailing:   proto.Bool(false),
						Lifetime:    proto.Uint64(7200),
						Periodicity: proto.Uint64(86400),
						State:       ApiCronJob_DISABLED.Enum(),
						Urn:         proto.String("aff4:/cron/GRRVersionBreakDown"),
					},
					{
						AllowOverruns: proto.Bool(false),
						Description:   proto.String(""),
						FlowName:      proto.String("LastAccessStats"),
						FlowRunnerArgs: &FlowRunnerArgs{
							FlowName: proto.String("LastAccessStats"),
						},
						IsFailing:   proto.Bool(true),
						LastRunTime: proto.Uint64(230000000),
						Lifetime:    proto.Uint64(86400),
						Periodicity: proto.Uint64(604800),
						State:       ApiCronJob_ENABLED.Enum(),
						Urn:         proto.String("aff4:/cron/LastAccessStats"),
					},
					{
						AllowOverruns: proto.Bool(false),
						Description:   proto.String("bar"),
						FlowName:      proto.String("OSBreakDown"),
						FlowRunnerArgs: &FlowRunnerArgs{
							FlowName: proto.String("OSBreakDown"),
						},
						IsFailing:   proto.Bool(false),
						Lifetime:    proto.Uint64(86400),
						Periodicity: proto.Uint64(604800),
						State:       ApiCronJob_ENABLED.Enum(),
						Urn:         proto.String("aff4:/cron/OSBreakDown"),
					},
				},
				TotalCount: proto.Int64(3),
			},
		},
		{
			func() (proto.Message, error) {
				return c.CreateFlow(ApiCreateFlowArgs{
					ClientId: proto.String("C.1000000000000000"),
					Flow: &ApiFlow{
						Name: proto.String("ListProcesses"),
						Args: MustNewAnyValue(&ListProcessesArgs{
							FetchBinaries: proto.Bool(true),
							FilenameRegex: proto.String("."),
						}),
						RunnerArgs: &FlowRunnerArgs{
							NotifyToUser:  proto.Bool(false),
							OutputPlugins: []*OutputPluginDescriptor{},
							Priority:      GrrMessage_HIGH_PRIORITY.Enum(),
						},
					},
				},
				)
			},

			&ApiFlow{
				Args: MustNewAnyValue(&ListProcessesArgs{
					FetchBinaries: proto.Bool(true),
					FilenameRegex: proto.String("."),
				}),
				Creator:      proto.String("test"),
				LastActiveAt: proto.Uint64(42000000),
				Name:         proto.String("ListProcesses"),
				StartedAt:    proto.Uint64(42000000),
				State:        ApiFlow_RUNNING.Enum(),
				Urn:          proto.String("aff4:/C.1000000000000000/flows/W:ABCDEF"),
			},
		},
	} {
		got, err := test.f()
		if err != nil {
			t.Errorf("test %d: %s", i, err)
		}
		if proto.MarshalTextString(got) != proto.MarshalTextString(test.expected) {
			t.Errorf("test %d:\nUnexpected result\nGot:\n%s\nExpected:\n%s",
				i,
				proto.MarshalTextString(got),
				proto.MarshalTextString(test.expected))
		} else {
			t.Logf("test %d:\nResult: %s", i, proto.MarshalTextString(got))
		}
	}
}
