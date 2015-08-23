package gotest

import (
	"testing"
)

func main() {
	dataA := []string{"A"}
	dataB := []string{"B"}
	for i = 0; i < 1000; i++ {
		append(dataA, "A")
		append(dataB, "B")
	}
	n := &gotification.Notification{Message: "test", AndroidReceivers: dataA, IOSReceivers: dataB}
	c := gotification.Config{"s", "t"}
	c.Set()
	n.Notify
}
