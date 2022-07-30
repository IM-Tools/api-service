/**
  @author:panliang
  @data:2022/7/30
  @note
**/
package client

import (
	"testing"
)

func TestMessage(t *testing.T) {
	var service GrpcMessageService
	service.SendGpcMessage([]byte(`{
  "msg_id":1,
  "msg_client_id":1, 
  "msg_code":200,
  "uid":"1a573f1e-a29a-4d88-8aa7-349f19dc1f4f",
  "to_uid":"ff654923-ad76-49ae-bbdc-70e0297caf44",
  "form_id":34,
  "to_id":35,
  "msg_type":1,
  "channel_type":1,
  "message":"你好！",
  "data":""
}`), "127.0.0.1:8002")
}
