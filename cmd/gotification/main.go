package main

import (
	"github.com/SteinWord/gotification"
)

func main() {
	dataA := []string{"A"}
	dataB := []string{"B"}
	for i := 0; i < 1000; i++ {
		dataA = append(dataA, "A")
		dataB = append(dataB, "B")
	}
	n := gotification.Notification{Message: "test", AndroidReceivers: dataA, IOSReceivers: dataB}
	c := gotification.Config{"s", "t"}
	c.Set()
	n.Notify()
}
