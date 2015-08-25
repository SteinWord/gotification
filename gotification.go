package gotification

import (
	"fmt"
	"github.com/alexjlockwood/gcm"
	"github.com/mattprice/Go-APNs"
	"io/ioutil"
	"sync"
	"time"
)

var (
	apnauth string
	gcmauth string
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
	apnauth = c.APNCertFile
	gcmauth = c.GCMAPIKey
}

func (n *Notification) Notify() bool {
	task1 := make(chan bool)
	task2 := make(chan bool)
	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	go func() {
		wg1.Add(len(n.AndroidReceivers))
		for i, v := range n.AndroidReceivers {
			go test(v, n.Message, i, &wg1)
		}
		wg1.Wait()
		task1 <- true
	}()
	go func() {
		wg2.Add(len(n.IOSReceivers))
		for i, v := range n.IOSReceivers {
			go test(v, n.Message, i, &wg2)
		}
		wg2.Wait()
		task2 <- true
	}()
	<-task1
	<-task2
	println("finish task")
	return true
}

func test(att string, msg string, i int, wg *sync.WaitGroup) bool {
	content := []byte(fmt.Sprintf("%s %s %s\n", time.Now(), att, msg))
	ioutil.WriteFile("/Users/QsF/tmp/go-testfile.log", content, 0755)
	fmt.Println(time.Now(), att, msg, i)
	time.Sleep(1 * time.Second)
	wg.Done()
	return true
}

func apnNotify(att string, msg string) {
	apns.LoadCertificateFile(false, apnauth)
	payload := &apns.Notification{Alert: fmt.Sprintf(msg), Badge: 0, Sandbox: true}
	payload.SetExpiryDuration(24 * time.Hour)
	payload.SendTo(att)
}

func gcmNotify(att string, msg string) {
	d := map[string]interface{}{"message": msg}
	sender := &gcm.Sender{ApiKey: gcmauth}
	data := gcm.NewMessage(d, att)
	sender.Send(data, 2)
}
