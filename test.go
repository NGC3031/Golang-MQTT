package main

import (
	"fmt"
	"os"
//	"strconv"
	"os/signal"
//	"io/ioutil"
//	"log"
	"net/http"
//	"encoding/json"
//	"github.com/gorilla/mux"
	"golang.org/x/net/websocket"
//	"github.com/gorilla/websocket"
	"github.com/yosssi/gmq/mqtt"
	"github.com/yosssi/gmq/mqtt/client"
)
type Msg struct {
	charge string
}
//Assign pkt for struct to send as JSON
var Pkt Msg

func main() {
	
	//router := mux.NewRouter()
	http.Handle("/ws", websocket.Handler(wsSignal))
	//http.Handle("/", rootHandler)
//	log.Fatal(http.ListenAndServe(":8844", router))
	go func() {
		fmt.Println("Start Go function")
//		panic(http.ListenAndServe(":8070", nil))
	}()
//	if err := http.ListenAndServe(":80", nil); err != nil {
  //      log.Fatal("ListenAndServe:", err)
//	}

	fmt.Println("Start MQTT")
	// Set up channel on which to send signal notifications.
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt, os.Kill)

	// Create an MQTT Client.
	cli := client.New(&client.Options{
		// Define the processing of the error handler.
		ErrorHandler: func(err error) {
			fmt.Println(err)
		},
	})

	// Terminate the Client.
	defer cli.Terminate()

	// Connect to the MQTT Server.
//	go func (){
		err := cli.Connect(&client.ConnectOptions{
		Network:  "tcp",
		Address:  "localhost:1883",
		ClientID: []byte(" "),
	})
	if err != nil {
		panic(err)
	}
//	)()

	// Subscribe to topics.
	err = cli.Subscribe(&client.SubscribeOptions{
		SubReqs: []*client.SubReq{
			&client.SubReq{
				TopicFilter: []byte("device"),
				QoS:         mqtt.QoS0,
				// Define the processing of the message handler.
				Handler: func(topicName, message []byte) {
					fmt.Println(string(topicName), string(message))
                                        fmt.Println("Bing")	
				},
			},
			&client.SubReq{
				TopicFilter: []byte("device/#"),
				QoS:         mqtt.QoS1,
				Handler: func(topicName, message []byte) {
					fmt.Println(string(topicName), string(message))
				fmt.Println("Bing-2")
				Pkt.charge=string(message)
                         //       if err = conn.WriteMessage(msgType, message); err != nil {
                        //        return
                       // }

				},
			},
		},
	})
	if err != nil {
		panic(err)
	}

	// Publish a message.
	err = cli.Publish(&client.PublishOptions{
		QoS:       mqtt.QoS0,
		TopicName: []byte("device/batt"),
		Message:   []byte("testMessage"),
	})
	if err != nil {
		panic(err)
	}

	// Unsubscribe from topics.
	err = cli.Unsubscribe(&client.UnsubscribeOptions{
		TopicFilters: [][]byte{
			[]byte("foo"),
		},
	})
	if err != nil {
		panic(err)
	}

	// Wait for receiving a signal.
	<-sigc

	// Disconnect the Network Connection.
	if err := cli.Disconnect(); err != nil {
		panic(err)
	}
}

// SetCharge receives a pointer to the Msg struct to update charge
////func (f *Msg) SetCharge(charge string) {
//	f.charge = charge
//	fmt.Println(f)
//}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Webserver Started")
//	content, err := ioutil.ReadFile("index.html")
//	if err != nil {
//		fmt.Println("Could not open file.", err)
//	}
//	fmt.Fprintf(w, "%s", content)
}

//func wsHandler(w http.ResponseWriter, r *http.Request) {
//	if r.Header.Get("Origin") != "http://"+r.Host {
//		http.Error(w, "Origin not allowed", 403)
//		return
//	}
//	conn, err := websocket.Upgrade(w, r, w.Header(), 1024, 1024)
//	if err != nil {
//		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
//	}
//	fmt.Println("Socket Server Started")
//	go echo(conn)
//}

func wsSignal(ws *websocket.Conn) {
fmt.Println("Socket")
//	for {
		//m := Pkt{}

//		err := conn.ReadJSON(&Pkt)
//		if err != nil {
//			fmt.Println("Error reading json.", err)
//		}

//		fmt.Printf("Got message: %#v\n", Pkt)

		if err := websocket.Message.Send(ws, Pkt); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Sending on socket")
//	}
}

