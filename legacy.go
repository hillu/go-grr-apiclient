package apiclient

import (
	"net/http"
	"net/url"
	"path"
	"strings"
)

// The simple dialog box for "Force Cron", "Run Hunt", "Stop Hunt"
// contain JS snippets similar to the following:
//
// grr.ExecuteRenderer(
//     "ConfirmationDialogRenderer.Layout",
//     {
//         "check_access_subject": "aff4:/cron/CleanCronJobs",
//         "unique": 1045,
//         "cron_urn": "aff4:/cron/CleanCronJobs",
//         "renderer": "ForceRunCronJobConfirmationDialog",
//         "id": "legacy_6"
//     });
func (c *APIClient) executeConfirmationDialogRenderer(renderer string, params url.Values) (*http.Response, error) {
	u := *c.BaseURL
	u.Path = path.Join(u.Path, "/render/RenderAjax", renderer)
	params.Set("renderer", renderer)
	rq, err := http.NewRequest("POST", u.String(), strings.NewReader(params.Encode()))
	if err != nil {
		return nil, err
	}
	rq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	return c.DoRequest(rq)
}

// ForceRunCronLegacy schedules the cron hunt cronId to be run
// immediately.
//
// This function uses a legacy (pre-API) interface.
func (c *APIClient) ForceRunCronLegacy(cronId string) (*http.Response, error) {
	return c.executeConfirmationDialogRenderer(
		"ForceRunCronJobConfirmationDialog",
		url.Values{"cron_urn": []string{cronId}},
	)
}

// EnableCronLegacy enables the cron hunt cronId to be run.
//
// This function uses a legacy (pre-API) interface.
func (c *APIClient) EnableCronLegacy(cronId string) (*http.Response, error) {
	return c.executeConfirmationDialogRenderer(
		"EnableCronJobConfirmationDialog",
		url.Values{"cron_urn": []string{cronId}},
	)
}

// DisableCronLegacy disables the cron hunt cronId to be run.
//
// This function uses a legacy (pre-API) interface.
func (c *APIClient) DisableCronLegacy(cronId string) (*http.Response, error) {
	return c.executeConfirmationDialogRenderer(
		"DisableCronJobConfirmationDialog",
		url.Values{"cron_urn": []string{cronId}},
	)
}

// RunHuntLegacy activates hunt huntId.
//
// This function uses a legacy (pre-API) interface.
func (c *APIClient) RunHuntLegacy(huntId string) (*http.Response, error) {
	return c.executeConfirmationDialogRenderer(
		"RunHuntConfirmationDialog",
		url.Values{"hunt_id": []string{huntId}},
	)
}

// StopHuntLegacy deactivates hunt huntId.
//
// This function uses a legacy (pre-API) interface.
func (c *APIClient) StopHuntLegacy(huntId string) (*http.Response, error) {
	return c.executeConfirmationDialogRenderer(
		"StopHuntConfirmationDialog",
		url.Values{"hunt_id": []string{huntId}},
	)
}
