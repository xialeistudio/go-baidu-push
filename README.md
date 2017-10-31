# Baidu-Push Golang SDK
Developed with golang, no other dependencies. [official site](https://push.baidu.com)
## ATTENTION
DO NOT TEST **PushAll* IN PRODUCTION ENVIRONMENT!!!

## Features

### Push
+ PushSingleDevice
+ PushAll 
+ PushTags
+ PushBatchDevice

### Query
+ QueryMsgStatus
+ QueryTimerRecords

#### Tag
+ QueryTags
+ CreateTag
+ AddDevices
+ DelDevices
+ DeviceNum
+ DelTag

### Timer
+ QueryList
+ Cancel

### Topic
+ QueryList

### Statistic
+ Device
+ Topic

## Unit Test
1. Setup environment **CHANNELID**,**APIKEY**,**SECRETKEY**
2. Change following params in **client_test.go** 
    ```
    var msgId = "2683796311873830912"
    var tag = "admin_ios"
    var timerId = "7244807349359244772"
    ```
3. `go test`
