package main 

import (
	"encoding/json"
	"fmt"
	"net"
)

/* Defined the same struct as the server */
type infoSend struct{
	VehicleId int `json:"vehicle_id"`
	RoadInfo int  `json:"road"`
	Direction int `json:"direction"`
}

func ReceiveData(conn net.Conn) {
  /* JSON decoder */
	decoder := json.NewDecoder(conn) 	
  for {
		var data infoSend

    /* Reads the JSON and parse it into struct */
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Println("Connection closed.")
			break
		}
		fmt.Println("Received:", data)
	}
}

func connectingWithServer() {
  conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Connected to server. Receiving data...")
	ReceiveData(conn)
}
