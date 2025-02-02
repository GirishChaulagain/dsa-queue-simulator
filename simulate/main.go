// server
package main

import (
	"encoding/json"
	"fmt"
	"github.com/GirishChaulagain/dsa-queue-simulator/shared"
	"net"
)

var vehicleQueues []*VehicleQueue
var laneQueue *LaneQueue
var lanes = []string{"AL1", "AL2", "BL1", "BL2", "CL1", "CL2", "DL1", "DL2"}

func main() {

	vehicleQueues = make([]*VehicleQueue, len(lanes))
	laneQueue = InitLaneQueue()

	for i := range lanes {
		vehicleQueues[i] = &VehicleQueue{}
	}

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

		laneIndex := getLaneIndex(vehicle.Lane)

		if laneIndex == -1 {
			continue
		}

		vehicleQueues[laneIndex].Enqueue(vehicle)
		laneQueue.IncrementLane(laneIndex)

		fmt.Println("lane Index = ", laneIndex)

		fmt.Println("Vehicle Count in lane Index ", laneIndex, " is : ", laneQueue.lInfo[laneIndex].Count)

		fmt.Printf("Received vehicle: %+v\n", vehicle)

		fmt.Printf("Incoming vehicle %v enqueud at lane index %d \n", vehicle, laneIndex)

		if frontVehicle, ok := vehicleQueues[laneIndex].Peek(); ok {
			fmt.Printf("Vehicle at the front of lane index %d is %d\n", laneIndex, frontVehicle.VehicleId)
		} else {
			fmt.Println("Queue is empty for lane index:", laneIndex)
		}

		fmt.Println("******************************************************")
	}
}

func getLaneIndex(lane string) int {
	for i, l := range lanes {
		if l == lane {
			return i
		}
	}
	return -1
}
