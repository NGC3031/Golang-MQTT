// Golang MQTT routing and forward IoT
package main

import (
    "fmt"
    MQTT "github.com/eclipse/paho.mqtt.golang"
    "os"
    "os/signal"
    "syscall"
)

var knt int
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
    fmt.Printf("MSG: %s\n", msg.Payload())
    text := fmt.Sprintf("this is result msg #%d!", knt)
    knt++
    token := client.Publish("nn/result", 0, false, text)
    token.Wait()
}

func main() {
    knt = 0
    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)

    opts := MQTT.NewClientOptions().AddBroker("tcp://test.mosquitto.org:1883")
    opts.SetClientID("wildcard")
    opts.SetDefaultPublishHandler(f)
    topic := "#"

    opts.OnConnect = func(c MQTT.Client) {
            if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
                    panic(token.Error())
            }
    }
    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
            panic(token.Error())
    } else {
            fmt.Printf("Connected to server\n")
    }
    <-c
}
