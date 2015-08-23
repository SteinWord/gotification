package gotification

import (
	"github.com/alexjlockwood/gcm"
	"github.com/mattprice/Go-APNs"
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
	go func() {
		for _, v := range n.AndroidReceivers {
			gcm(v, n.Message, gcmkey)
		}
	}()
	go func() {
		for _, v := range n.IOSReceivers {
			apn(v, n.Message, apnpem)
		}
	}()
}

func apn(att string, msg string, certpem string) (status bool) {
	status = true
	err := apns.LoadCertificateFile(false, certpem)
	if err != nil {
		status = false
	}
	payload := &apns.Notification{Alert: fmt.Sprintf(msg), Badge: 0, Sandbox: true}
	payload.SetExpiryDuration(24 * time.Hour)
	err := payload.SendTo(att)
	if err != nil {
		status = false
	}
	return
}

func gcm(att string, msg string, apikey string) (status bool) {
	status = true
	d := map[string]interface{}{"message": msg}
	sender := &gcm.Sender{ApiKey: apikey}
	data := gcm.NewMessage(d, regId)
	_, err := sender.Send(data, 2)
	if err != nil {
		status = false
	}
	return
}
