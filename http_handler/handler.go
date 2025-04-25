package httphandler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func Ptr[T any](t T) *T {
	return &t
}

type HTTPResponse struct {
	Status    int
	Headers   http.Header
	Body      map[string]interface{}
	Tets_Body interface{} //	TODO: remove later, test for storing arbitrary json or list
}

func tabs(tabs int) string {
	temp := ""
	for range tabs {
		temp += "\t"
	}
	return temp
}
func (resp *HTTPResponse) resolveInterfaceList(temp string, key string, list []interface{}, depth int) string {
	temp = fmt.Sprintf("%s\n%s%s:", temp, tabs(depth-1), key)
	for key, value := range list {
		switch v := value.(type) {
		case string:
			temp = fmt.Sprintf("%s\n%s%d: %s", temp, tabs(depth), key, v)
		case float64:
			temp = fmt.Sprintf("%s\n%s%d: %f", temp, tabs(depth), key, v)
		case bool:
			temp = fmt.Sprintf("%s\n%s%d: %b", temp, tabs(depth), key, v)
		case []interface{}:
			temp = resp.resolveInterfaceList(temp, strconv.Itoa(key), v, depth+1)
		case map[string]interface{}:
			temp = resp.resolveInterfaceMap(temp, strconv.Itoa(key), v, depth+1)
		default:
			temp = fmt.Sprintf("%s\n%s%d: unknown type", tabs(depth), temp, key)
		}
	}
	return temp
}
func (resp *HTTPResponse) resolveInterfaceMap(temp string, key string, Map map[string]interface{}, depth int) string {
	temp = fmt.Sprintf("%s\n%s%s", temp, tabs(depth-1), key)
	for key, value := range Map {
		switch v := value.(type) {
		case string:
			temp = fmt.Sprintf("%s\n%s%s: %s", temp, tabs(depth), key, v)
		case float64:
			temp = fmt.Sprintf("%s\n%s%s: %f", temp, tabs(depth), key, v)
		case bool:
			temp = fmt.Sprintf("%s\n%s%s: %b", temp, tabs(depth), key, v)
		case []interface{}:
			temp = resp.resolveInterfaceList(temp, key, v, depth+1)
		case map[string]interface{}:
			temp = resp.resolveInterfaceMap(temp, key, v, depth+1)
		default:
			temp = fmt.Sprintf("%s\n%s%s: unknown type", tabs(depth), temp, key)
		}
	}
	return temp
}

func (resp *HTTPResponse) ToString() string {
	temp := ""
	temp = fmt.Sprintf("%s\n\tStatus: %d", temp, resp.Status)

	temp = fmt.Sprintf("%s\n", temp)
	for key, value := range resp.Headers {
		temp = fmt.Sprintf("%s\n\t%s: %v", temp, key, value)
	}
	temp = fmt.Sprintf("%s\n", temp)

	temp = fmt.Sprintf("%s\n", temp)
	for key, value := range resp.Body {
		switch v := value.(type) {
		case string:
			temp = fmt.Sprintf("%s\n\t%s: %s", temp, key, v)
		case float64:
			temp = fmt.Sprintf("%s\n\t%s: %f", temp, key, v)
		case bool:
			temp = fmt.Sprintf("%s\n\t%s: %b", temp, key, v)
		case []interface{}:
			temp = resp.resolveInterfaceList(temp, key, v, 1)
		case map[string]interface{}:
			temp = resp.resolveInterfaceMap(temp, key, v, 1)
		default:
			temp = fmt.Sprintf("%s\n\t%s: unknown type", temp, key)
		}
	}
	temp = fmt.Sprintf("%s\n", temp)

	return temp
}

func Handle(client *http.Client, request Request) (*HTTPResponse, error) {

	var err error
	var resp *http.Response

	req, req_err := http.NewRequest(request.Method, request.Url, &request.Body)
	if req_err != nil {
		return nil, fmt.Errorf("ERROR: failed to create request: %+v\n\t%+v", request, req_err)
	}
	for range request.Retries + 1 {
		resp, err = client.Do(req)
		if err == nil {
			break
		}
	}
	if err != nil {
		return nil, fmt.Errorf("ERROR: failed to get response after %d Retries: %+v\n\t%+v", request.Retries, request, err)
	}

	defer resp.Body.Close()

	resp_body_bytes, resp_err := io.ReadAll(resp.Body)

	var body_json map[string]interface{}
	if resp_err != nil {
		return nil, fmt.Errorf("ERROR: failed to read the bytes from the response body %+v", resp_err)
	}
	json.Unmarshal(resp_body_bytes, &body_json)

	var json_interface interface{}
	if resp_err != nil {
		return nil, fmt.Errorf("ERROR: failed to read the bytes from the response body INTERFACE %+v", resp_err)
	}
	json.Unmarshal(resp_body_bytes, &json_interface)

	return &HTTPResponse{
			Status:    resp.StatusCode,
			Headers:   resp.Request.Header,
			Body:      body_json,
			Tets_Body: json_interface,
		},
		nil
}
