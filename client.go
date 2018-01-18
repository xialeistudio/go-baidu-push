package baidu_push

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Client struct {
	ApiKey     string
	SecretKey  string
	DeviceType int
}

func NewClient(apiKey, secretKey string, deviceType int) *Client {
	client := &Client{
		ApiKey:    apiKey,
		SecretKey: secretKey,
	}
	if deviceType > 0 {
		client.DeviceType = deviceType
	}
	return client
}

func (c *Client) getUserAgent() string {
	return fmt.Sprintf("BCCS_SDK/3.0 %s go/%s", runtime.GOOS, runtime.Version())
}

func (c *Client) getContentType() string {
	return "application/x-www-form-urlencoded;charset=utf-8"
}

func (c *Client) Execute(req *Request) (map[string]interface{}, error) {
	if req.Params["device_type"] == nil && c.DeviceType > 0 {
		req.Params["device_type"] = c.DeviceType
	}
	req.Params["apikey"] = c.ApiKey
	req.Params["timestamp"] = time.Now().Unix()
	req.Params["sign"] = c.getSignature(req)
	return c.sendRequest(req)
}

func (c *Client) sendRequest(req *Request) (map[string]interface{}, error) {
	values := url.Values{}
	for key, value := range req.Params {
		switch v := value.(type) {
		case int:
			values[key] = []string{strconv.Itoa(v)}
		case string:
			values[key] = []string{v}
		case int64:
			values[key] = []string{strconv.FormatInt(v, 10)}
		}
	}
	httpReq, err := http.NewRequest("POST", req.Url, strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	httpReq.Header.Set("Content-Type", c.getContentType())
	httpReq.Header.Set("User-Agent", c.getUserAgent())
	client := &http.Client{}
	res, err := client.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	result := make(map[string]interface{})
	err = json.Unmarshal(buf, &result)
	return result, err
}

func (c *Client) getSignature(req *Request) string {
	str := "POST" + req.Url
	keys := make([]string, 0, len(req.Params))
	for key := range req.Params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		switch v := req.Params[key].(type) {
		case string:
			str += key + "=" + v
		case int:
			str += key + "=" + strconv.Itoa(v)
		case int64:
			str += key + "=" + strconv.FormatInt(v, 10)
		}
	}
	str += c.SecretKey
	m := md5.New()
	m.Write([]byte(url.QueryEscape(str)))
	return fmt.Sprintf("%x", m.Sum(nil))
}
