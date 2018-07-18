package midware

import (
	"testing"
	"time"
	"business/message"
	"fmt"
)

type AndroidConsumer struct {
}

func (consumer *AndroidConsumer) Consumer(body []byte) {
	fmt.Println(string(body))
}

//TestPublish
func TestPublish(t *testing.T) {
	var msg *message.CommonMessage
	msg = message.NewCommonMessage(12, 1, 1, 0, "test", "test test")
	body, err := msg.GetMessage()

	if err != nil {
		t.Error("转换json出错")
	}

	Publish(AndroidExchangeName, []byte(body))
	Publish(AndroidExchangeName, []byte(body))
	time.Sleep(10 * time.Second)
}

//TestReceive
func TestReceive(t *testing.T) {
	Receive(AndroidExchangeName, new(AndroidConsumer))
	time.Sleep(10 * time.Second)
}
