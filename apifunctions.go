// This file has been autogenerated by gen-functions.go

package apiclient

func (c APIClient) ListArtifacts(rq ApiListArtifactsArgs) (rs *ApiListArtifactsResult, err error) {
	rs = new(ApiListArtifactsResult)
	if err := c.do("GET", "/api/artifacts", &rq, rs); err != nil {
		return nil, err
	}
	return
}

func (c APIClient) UploadArtifact(rq ApiUploadArtifactArgs) error {
	return c.post("/api/artifacts/upload", &rq)
}

func (c APIClient) DeleteArtifacts(rq ApiDeleteArtifactsArgs) error {
	return c.post("/api/artifacts/delete", &rq)
}

func (c APIClient) SearchClients(rq ApiSearchClientsArgs) (rs *ApiSearchClientsResult, err error) {
	rs = new(ApiSearchClientsResult)
	if err := c.do("GET", "/api/clients", &rq, rs); err != nil {
		return nil, err
	}
	return
}

func (c APIClient) GetClient(rq ApiGetClientArgs) (rs *ApiGetClientResult, err error) {
	rs = new(ApiGetClientResult)
	if err := c.do("GET", "/api/clients/<client_id>", &rq, rs); err != nil {
		return nil, err
	}
	return
}

func (c APIClient) GetClientVersionTimes(clientId string, ) (rs *ApiGetClientVersionTimesResult, err error) {
	rs = new(ApiGetClientVersionTimesResult)
	if err := c.get("/api/clients/" + clientId + "/version", nil, rs); err != nil {
		return nil, err
	}
	return
}

func (c APIClient) GetLastClientIPAddress(clientId string, ) (rs *ApiGetLastClientIPAddressResult, err error) {
	rs = new(ApiGetLastClientIPAddressResult)
	if err := c.get("/api/clients/" + clientId + "/last", nil, rs); err != nil {
		return nil, err
	}
	return
}

func (c APIClient) ListClientsLabels() (rs *ApiListClientsLabelsResult, err error) {
	rs = new(ApiListClientsLabelsResult)
	if err := c.get("/api/clients/labels", nil, rs); err != nil {
		return nil, err
	}
	return
}

func (c APIClient) AddClientsLabels(rq ApiAddClientsLabelsArgs) error {
	return c.post("/api/clients/labels/add", &rq)
}

func (c APIClient) RemoveClientsLabels(rq ApiRemoveClientsLabelsArgs) error {
	return c.post("/api/clients/labels/remove", &rq)
}

func (c APIClient) ListFlows(rq ApiListFlowsArgs) (rs *ApiListFlowsResult, err error) {
	rs = new(ApiListFlowsResult)
	if err := c.do("GET", "/api/clients/<client_id>/flows", &rq, rs); err != nil {
		return nil, err
	}
	return
}

func (c APIClient) CancelFlow(rq ApiCancelFlowArgs, clientId string, flowId string) error {
	return c.post("/api/clients/" + clientId + "/flows/" + flowId + "/actions/cancel", &rq)
}

func (c APIClient) CreateFlow(rq ApiCreateFlowArgs) error {
	return c.post("/api/flows", &rq)
}

func (c APIClient) ListCronJobs(rq ApiListCronJobsArgs) (rs *ApiListCronJobsResult, err error) {
	rs = new(ApiListCronJobsResult)
	if err := c.do("GET", "/api/cron-jobs", &rq, rs); err != nil {
		return nil, err
	}
	return
}

func (c APIClient) ListHunts(rq ApiListHuntsArgs) (rs *ApiListHuntsResult, err error) {
	rs = new(ApiListHuntsResult)
	if err := c.do("GET", "/api/hunts", &rq, rs); err != nil {
		return nil, err
	}
	return
}