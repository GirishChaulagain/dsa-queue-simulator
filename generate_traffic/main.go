package main

import (
	"fmt"
	"github.com/GirishChaulagain/dsa-queue-simulator/shared"
	"math/rand"
	"time"
)

func main() {
	lanes := []string{"AL1", "AL2", "BL1", "BL2", "CL1", "CL2", "DL1", "DL2"}

	for {
		lane := lanes[rand.Intn(len(lanes))]
		direction := ""
		if lane[len(lane)-1] == '1' {
			direction = "left"
		} else if lane[len(lane)-1] == '2' {
			directions := []string{"straight", "right"}
			direction = directions[rand.Intn(len(directions))]
		}

		vehicle := shared.Vehicle{Lane: lane, VehicleId: rand.Intn(1000), Direction: direction}

		fmt.Printf("Random Vehicle Information: %+v\n", vehicle)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second)

	}
}
