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

func (vq *VehicleQueue) Peek() (shared.VehicleInfo, bool) {
	if len(vq.vQueue) == 0 {
		return shared.VehicleInfo{}, false
	}
	return vq.vQueue[0], true
}

type LaneQueue struct {
	lInfo []shared.LaneInfo
}

func InitLaneQueue() *LaneQueue {
	lanes := []string{"AL1", "AL2", "BL1", "BL2", "CL1", "CL2", "DL1", "DL2"}
	lq := &LaneQueue{}

	for _, lane := range lanes {
		lq.lInfo = append(lq.lInfo, shared.LaneInfo{Lane: lane, Count: 0})
	}
	return lq
}
