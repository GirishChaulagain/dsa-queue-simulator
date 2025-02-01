// server
package main

import (
	"encoding/json"
	"fmt"
	"github.com/GirishChaulagain/dsa-queue-simulator/shared"
	"net"
)

var vehicleQueues []*VehicleQueue
var lanes = []string{"AL1", "AL2", "BL1", "BL2", "CL1", "CL2", "DL1", "DL2"}

func main() {

	vehicleQueues = make([]*VehicleQueue, len(lanes))

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting simulator", err)
		return
	}
	defer listen.Close()

	fmt.Println("Simulator started, waiting for connection...")

	for {
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(connection)

	}

}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	var vehicle shared.VehicleInfo
	decoder := json.NewDecoder(connection)
	for {
		err := decoder.Decode((&vehicle))
		if err != nil {
			fmt.Println("Error Decoding vehicle:", err)
			return
		}

		fmt.Printf("Received vehicle: %+v\n", vehicle)

	}
}
