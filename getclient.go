package apiclient

import (
	"net/url"
	"path"
	"strconv"
)

func (c APIClient) GetClient(client string, timestamp uint64) (rs *ApiGetClientResult, err error) {
	rs = new(ApiGetClientResult)
	values := make(url.Values)
	if timestamp != 0 {
		values.Set("timestamp", strconv.FormatUint(timestamp, 10))
	}
	if err := c.get(path.Join("/api/clients", client), values, rs); err != nil {
		return nil, err
	}
	return
}
