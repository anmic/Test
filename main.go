package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	indexFile, err := os.Open("index.html")
	if err != nil {
		fmt.Println(err)
	}

	index, err := ioutil.ReadAll(indexFile)
	if err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/websocket", webSocketHandler)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(index))
	})

	http.ListenAndServe(":3000", nil)
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Client subscribed")

	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		err = conn.WriteMessage(websocket.TextMessage, []byte(strconv.Itoa(i)))
		if err != nil {
			fmt.Println(err)
			break
		}
	}

	conn.Close()
	// fmt.Println("Connection closed")
}
