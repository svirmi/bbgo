// Code generated by "requestgen -method GET -responseType .APIResponse -responseDataField Result -url /api/fills -type GetFillsRequest -responseDataType []Fill"; DO NOT EDIT.

package ftxapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"time"
)

func (g *GetFillsRequest) Market(market string) *GetFillsRequest {
	g.market = &market
	return g
}

func (g *GetFillsRequest) StartTime(startTime time.Time) *GetFillsRequest {
	g.startTime = &startTime
	return g
}

func (g *GetFillsRequest) EndTime(endTime time.Time) *GetFillsRequest {
	g.endTime = &endTime
	return g
}

func (g *GetFillsRequest) OrderID(orderID int) *GetFillsRequest {
	g.orderID = &orderID
	return g
}

func (g *GetFillsRequest) Order(order string) *GetFillsRequest {
	g.order = &order
	return g
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetFillsRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}
	// check market field -> json key market
	if g.market != nil {
		market := *g.market

		// assign parameter of market
		params["market"] = market
	} else {
	}
	// check startTime field -> json key start_time
	if g.startTime != nil {
		startTime := *g.startTime

		// assign parameter of startTime
		// convert time.Time to seconds time stamp
		params["start_time"] = strconv.FormatInt(startTime.Unix(), 10)
	} else {
	}
	// check endTime field -> json key end_time
	if g.endTime != nil {
		endTime := *g.endTime

		// assign parameter of endTime
		// convert time.Time to seconds time stamp
		params["end_time"] = strconv.FormatInt(endTime.Unix(), 10)
	} else {
	}
	// check orderID field -> json key orderId
	if g.orderID != nil {
		orderID := *g.orderID

		// assign parameter of orderID
		params["orderId"] = orderID
	} else {
	}
	// check order field -> json key order
	if g.order != nil {
		order := *g.order

		// assign parameter of order
		params["order"] = order
	} else {
	}

	query := url.Values{}
	for k, v := range params {
		query.Add(k, fmt.Sprintf("%v", v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetFillsRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetFillsRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for k, v := range params {
		query.Add(k, fmt.Sprintf("%v", v))
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetFillsRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (g *GetFillsRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (g *GetFillsRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for k, v := range slugs {
		needleRE := regexp.MustCompile(":" + k + "\\b")
		url = needleRE.ReplaceAllString(url, v)
	}

	return url
}

func (g *GetFillsRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := g.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for k, v := range params {
		slugs[k] = fmt.Sprintf("%v", v)
	}

	return slugs, nil
}

func (g *GetFillsRequest) Do(ctx context.Context) ([]Fill, error) {

	// no body params
	var params interface{}
	query, err := g.GetQueryParameters()
	if err != nil {
		return nil, err
	}

	apiURL := "/api/fills"

	req, err := g.client.NewAuthenticatedRequest(ctx, "GET", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponse
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	var data []Fill
	if err := json.Unmarshal(apiResponse.Result, &data); err != nil {
		return nil, err
	}
	return data, nil
}