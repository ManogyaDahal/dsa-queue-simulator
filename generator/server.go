/* Necessary Server Codes for running server */
package generator 

import (
  "time"
  "encoding/json"
	"fmt"
	"net"
  "math/rand"
)

type infoSend struct{
	vehicleId int `json:"vehicle_id"`
	roadInfo int  `json:"road"`
	direction int `json:"direction"`
}
func startServer (){
  fmt.Println("Starting server on port 8080")

  ln, err := net.Listen("tcp","localhost:8080")
  if err != nil{
    fmt.Println("Error starting the server:",err)
    return 
  }
  defer ln.Close()

  /* Accepting a client's connection */ 
  conn, err := ln.Accept()
  if err != nil {
    fmt.Println("Error accepting clinet:", err)
    return 
  }
  defer conn.Close()
  fmt.Println("Client connected",conn.RemoteAddr() )

  sendData(conn)
}
/* 
  @NOTE: 
  I said to return the random integers for road info and direction
  Because I have defined constants for each in queue.go 
*/
func generateRandomData() infoSend{
  //In json format
 return infoSend{
    vehicleId: rand.Intn(100),
    roadInfo: rand.Intn(4),
    direction: rand.Intn(4),
  } 
}

func sendData(conn net.Conn){
	rand.Seed(time.Now().UnixNano())
  encoder := json.NewEncoder(conn)
  /* looping to send the data */
  for {
    data := generateRandomData()
    err := encoder.Encode(data)
    if err != nil {
      fmt.Println("Failed to encode data:",err)
    }
    fmt.Println("Data", data)
		time.Sleep(2 * time.Second) 
  }
}






