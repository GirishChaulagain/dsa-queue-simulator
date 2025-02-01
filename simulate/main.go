// server
package main

import (
	"encoding/json"
	"fmt"
	"github.com/GirishChaulagain/dsa-queue-simulator/shared"
	"net"
)

func main() {

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

}
