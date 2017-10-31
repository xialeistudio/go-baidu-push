package baidu_push

import (
	"testing"
	"os"
	"time"
)

var client = NewClient(os.Getenv("APIKEY"), os.Getenv("SECRETKEY"), DeviceIOS)
var channelId = os.Getenv("CHANNELID")
var msgId = "2683796311873830912"
var tag = "admin_ios"
var timerId = "7244807349359244772"
/**
 **********************TEST PUSH******************
 */
func TestClientPushSingleDevice(t *testing.T) {
	t.Log("\n\n**********************TEST PUSH******************")
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
	req := NewRequest("/rest/3.0/push/single_device")
	req.Params["channel_id"] = channelId
	req.Params["msg"] = result
	req.Params["msg_type"] = MsgNotification
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientPushSingleDevice", res)
}

func TestClientPushTags(t *testing.T) {
	msg := &IOSMsgBuilder{
		Aps: &IOSMsgAps{
			Alert: "天气不错:批量推送",
		},
	}
	result, err := msg.Build()
	if err != nil {
		t.Error(err)
		return
	}
	req := NewRequest("/rest/3.0/push/tags")
	req.Params["type"] = 1
	req.Params["tag"] = tag
	req.Params["msg_type"] = MsgNotification
	req.Params["msg"] = result
	req.Params["send_time"] = time.Now().Unix() + 120
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientPushTags", res)
}

func TestClientPushBatchDevice(t *testing.T) {
	msg := &IOSMsgBuilder{
		Aps: &IOSMsgAps{
			Alert: "天气不错:批量推送",
		},
	}
	result, err := msg.Build()
	if err != nil {
		t.Error(err)
		return
	}
	req := NewRequest("/rest/3.0/push/batch_device")
	req.Params["channel_ids"] = `["` + channelId + `"]`
	req.Params["msg_type"] = MsgNotification
	req.Params["msg"] = result

	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientPushBatchDevice", res)
}

/**
 **********************TEST QUERY******************
 */
func TestClientReportQueryMsgStatus(t *testing.T) {
	t.Log("\n\n**********************TEST QUERY******************")
	req := NewRequest("/rest/3.0/report/query_msg_status")
	req.Params["msg_id"] = msgId
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientReportQueryMsgStatus", res)
}

func TestClientQueryTimerRecords(t *testing.T) {
	req := NewRequest("/rest/3.0/report/query_timer_records")
	req.Params["timer_id"] = timerId
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientQueryTimerRecords", res)
}

/**
 **********************TEST TAG******************
 */
func TestClientAppQueryTags(t *testing.T) {
	t.Log("\n\n**********************TEST TAG******************")
	req := NewRequest("/rest/3.0/app/query_tags")
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientAppQueryTags", res)
}

func TestClientAppCreateTag(t *testing.T) {
	req := NewRequest("/rest/3.0/app/create_tag")
	req.Params["tag"] = "test20171031"
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientAppCreateTag", res)
}

func TestClientTagAddDevices(t *testing.T) {
	req := NewRequest("/rest/3.0/tag/add_devices")
	req.Params["tag"] = "test20171031"
	req.Params["channel_ids"] = `["` + channelId + `"]`
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientTagAddDevices", res)
}
func TestClientTagDelDevices(t *testing.T) {
	req := NewRequest("/rest/3.0/tag/del_devices")
	req.Params["tag"] = "test20171031"
	req.Params["channel_ids"] = `["` + channelId + `"]`
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientTagDelDevices", res)
}

func TestClientTagDeviceNum(t *testing.T) {
	req := NewRequest("/rest/3.0/tag/device_num")
	req.Params["tag"] = "test20171031"
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientTagDeviceNum", res)
}

func TestClientAppDelTag(t *testing.T) {
	req := NewRequest("/rest/3.0/app/del_tag")
	req.Params["tag"] = "test20171031"
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientAppDelTag", res)
}

/**
 **********************TEST TIMER******************
 */
func TestClientTimerQueryList(t *testing.T) {
	t.Log("\n\n**********************TEST TIMER******************")
	req := NewRequest("/rest/3.0/timer/query_list")
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientTimerQueryList", res)
}

func TestClientTimerCancel(t *testing.T) {
	req := NewRequest("/rest/3.0/timer/cancel")
	req.Params["timer_id"] = timerId
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientTimerQueryList", res)
}

/**
 **********************TEST TOPIC******************
 */
func TestClientTopicQueryList(t *testing.T) {
	t.Log("\n\n**********************TEST TOPIC******************")
	req := NewRequest("/rest/3.0/topic/query_list")
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientTopicQueryList", res)
}

/**
 **********************TEST REPORT******************
 */

func TestClientReportStatisticDevice(t *testing.T) {
	t.Log("\n\n**********************TEST REPORT******************")
	req := NewRequest("/rest/3.0/report/statistic_device")
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientReportStatisticDevice", res)
}

func TestClientStatisticTopic(t *testing.T) {
	req := NewRequest("/rest/3.0/report/statistic_topic")
	res, err := client.Execute(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("TestClientStatisticDevice", res)
}
