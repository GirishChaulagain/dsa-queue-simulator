package main

import (
	"encoding/json"
	"fmt"
	"github.com/GirishChaulagain/dsa-queue-simulator/shared"
	"math/rand"
	"net"
	"time"
)

func main() {
	lanes := []string{"AL1", "AL2", "BL1", "BL2", "CL1", "CL2", "DL1", "DL2"}

	connection, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to simulator:", err)
		return
	}
	defer connection.Close()

	for {
		lane := lanes[rand.Intn(len(lanes))]
		direction := ""
		if lane[len(lane)-1] == '1' {
			direction = "left"
		} else if lane[len(lane)-1] == '2' {
			directions := []string{"straight", "right"}
			direction = directions[rand.Intn(len(directions))]
		}

		vehicle := shared.VehicleInfo{Lane: lane, VehicleId: rand.Intn(1000), Direction: direction}
		data, err := json.Marshal(vehicle)
		if err != nil {
			fmt.Println("Error marshalling vehicle:", err)
			continue
		}

		_, err = connection.Write(data)
		if err != nil {
			fmt.Println("Error sending message", err)
			continue
		}

		fmt.Printf("Random Vehicle Information: %+v\n", vehicle)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	}
}
