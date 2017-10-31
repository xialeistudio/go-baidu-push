package baidu_push

import "encoding/json"

type MsgBuilder interface {
	Build() (string, error)
	AddParam(key string, value interface{})
	RemoveParam(key string)
	GetParam(key string) interface{}
}

type AndroidMsgBuilder struct {
	Title                  string                 `json:"title,omitempty"`
	Description            string                 `json:"description"`
	NotificationBuilderId  int                    `json:"notification_builder_id,omitempty"`
	NotificationBasicStyle int                    `json:"notification_basic_style,omitempty"`
	OpenType               int                    `json:"open_type,omitempty"`
	Url                    string                 `json:"url,omitempty"`
	PkgContent             string                 `json:"pkg_content,omitempty"`
	CustomContent          map[string]interface{} `json:"custom_content,omitempty"`
}

func (builder *AndroidMsgBuilder) Build() (string, error) {
	buf, err := json.Marshal(builder)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (builder *AndroidMsgBuilder) AddParam(key string, value interface{}) {
	if builder.CustomContent == nil {
		builder.CustomContent = make(map[string]interface{})
	}
	builder.CustomContent[key] = value
}

func (builder *AndroidMsgBuilder) RemoveParam(key string) {
	if builder.CustomContent == nil {
		return
	}
	delete(builder.CustomContent, key)
}

func (builder *AndroidMsgBuilder) GetParam(key string) interface{} {
	if builder.CustomContent == nil {
		return nil
	}
	return builder.CustomContent[key]
}
