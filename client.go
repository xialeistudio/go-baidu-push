package baidu_push

type Client struct {
	ApiKey       string
	SecretKey    string
	DeployStatus int
	DeviceType   int
}

func NewClient(apiKey, secretKey string, deployStatus, deviceType int) *Client {
	client := &Client{
		ApiKey:       apiKey,
		SecretKey:    secretKey,
		DeployStatus: deployStatus,
	}
	if deviceType > 0 {
		client.DeviceType = deviceType
	}
	return client
}
