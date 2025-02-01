package main

import "github.com/GirishChaulagain/dsa-queue-simulator/shared"

type VehicleQueue struct {
	vQueue []shared.VehicleInfo
}

func (vq *VehicleQueue) Enqueue(vehicle shared.VehicleInfo) {
	vq.vQueue = append(vq.vQueue, vehicle)
}

func (vq *VehicleQueue) Dequeue() (shared.VehicleInfo, bool) {
	if len(vq.vQueue) == 0 {
		return shared.VehicleInfo{}, false
	}
	vehicle := vq.vQueue[0]
	vq.vQueue = vq.vQueue[1:]
	return vehicle, true
}

func (vq *VehicleQueue) Size() int {
	return len(vq.vQueue)
}
