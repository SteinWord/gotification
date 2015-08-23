package gotification

import (
	"fmt"
	"github.com/alexjlockwood/gcm"
	"github.com/mattprice/Go-APNs"
	"time"
)

var (
	apnpem string
	gcmkey string
)

type Config struct {
	APNCertFile string
	GCMAPIKey   string
}

type Notification struct {
	Message          string
	AndroidReceivers []string
	IOSReceivers     []string
}

func (c *Config) Set() {
	apnpem = c.APNCertFile
	gcmkey = c.GCMAPIKey
}

func (n *Notification) Notify() (result bool) {
	task1 := make(chan bool)
	task2 := make(chan bool)
	go func() {
		for _, v := range n.AndroidReceivers {
			go test(v, n.Message, gcmkey)
		}
		task1 <- true
	}()
	go func() {
		for _, v := range n.IOSReceivers {
			go test(v, n.Message, apnpem)
		}
		task2 <- true
	}()
	<-task1
	<-task2
	return
}

func test(att string, msg string, cert string) {
	fmt.Println(time.Now(), att, msg, cert)
	time.Sleep(1 * time.Second)
}

func apnNotify(att string, msg string, certpem string) (status bool) {
	status = true
	err := apns.LoadCertificateFile(false, certpem)
	if err != nil {
		status = false
	}
	payload := &apns.Notification{Alert: fmt.Sprintf(msg), Badge: 0, Sandbox: true}
	payload.SetExpiryDuration(24 * time.Hour)
	err = payload.SendTo(att)
	if err != nil {
		status = false
	}
	return
}

func gcmNotify(att string, msg string, apikey string) (status bool) {
	status = true
	d := map[string]interface{}{"message": msg}
	sender := &gcm.Sender{ApiKey: apikey}
	data := gcm.NewMessage(d, att)
	_, err := sender.Send(data, 2)
	if err != nil {
		status = false
	}
	return
}
