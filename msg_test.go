package baidu_push

import (
	"testing"
	"strings"
)

func TestAndroidMsgBuilderBuild(t *testing.T) {
	msg1 := &AndroidMsgBuilder{
		Description: "天气不错",
	}
	_, err := msg1.Build()
	if err != nil {
		t.Error(err)
	}
	params := make(map[string]interface{})
	params["test"] = 1
	msg2 := &AndroidMsgBuilder{
		Title:                  "标题",
		Description:            "内容",
		NotificationBuilderId:  0,
		NotificationBasicStyle: 7,
		OpenType:               0,
		Url:                    "https://www.baidu.com",
		PkgContent:             "com.baidu.push",
		CustomContent:          params,
	}
	_, err = msg2.Build()
	if err != nil {
		t.Error(err)
	}
}
func TestAndroidMsgBuilderParam(t *testing.T) {
	msg := &AndroidMsgBuilder{
		Description: "description",
	}
	msg.AddParam("url", "https://www.baidu.com")
	result, err := msg.Build()
	if err != nil {
		t.Error(err)
		return
	}
	if !strings.Contains(result, "https://www.baidu.com") {
		t.Error("AndroidMsgBuilder AddParam failed")
		return
	}
	value := msg.GetParam("url")
	if value.(string) != "https://www.baidu.com" {
		t.Error("AndroidMsgBuilder GetParam failed")
		return
	}
	msg.RemoveParam("url")
	if msg.GetParam("url") != nil {
		t.Error("AndroidMsgBuilder RemoveParam failed")
	}
}

func TestIOSMsgBuilderBuild(t *testing.T) {
	aps := &IOSMsgAps{
		Alert: "天气不错",
		Sound: "1",
		Badge: 1,
	}
	msg := &IOSMsgBuilder{
		Aps: aps,
	}
	_, err := msg.Build()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestIOSMsgBuilderParam(t *testing.T) {
	aps := &IOSMsgAps{
		Alert: "天气不错",
		Sound: "1",
		Badge: 1,
	}
	msg := &IOSMsgBuilder{
		Aps: aps,
	}
	msg.AddParam("url", "https://www.baidu.com")
	result, err := msg.Build()
	if err != nil {
		t.Error(err)
		return
	}
	if !strings.Contains(result, "https://www.baidu.com") {
		t.Error("IOSMsgBuilder AddParam failed")
		return
	}
	value := msg.GetParam("url")
	if value != "https://www.baidu.com" {
		t.Error("IOSMsgBuilder GetParam failed")
		return
	}
	msg.RemoveParam("url")
	if msg.GetParam("url") != nil {
		t.Error("IOSMsgBuilder RemoveParam failed")
	}
}
