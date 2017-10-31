package baidu_push

import (
	"testing"
	"os"
)

var client = NewClient(os.Getenv("APIKEY"), os.Getenv("SECRETKEY"), DeviceIOS)

func TestClientPushSingleDevice(t *testing.T) {
	msg := &IOSMsgBuilder{
		Aps: &IOSMsgAps{
			Alert: "天气不错",
		},
	}
	result, err := msg.Build()
	if err != nil {
		t.Error(err)
		return
	}
	channelId := "4651203809675659079"
	req := NewRequest("/rest/3.0/push/single_device")
	req.Params["channel_id"] = channelId
	req.Params["msg"] = result
	req.Params["msg_type"] = MsgNotification
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}
