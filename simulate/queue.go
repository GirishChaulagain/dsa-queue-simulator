package main

import (
	"github.com/GirishChaulagain/dsa-queue-simulator/shared"
	"sync"
)

type VehicleQueue struct {
	vQueue []shared.VehicleInfo
	mu     sync.Mutex
}

func (vq *VehicleQueue) Enqueue(vehicle shared.VehicleInfo) {
	vq.mu.Lock()
	defer vq.mu.Unlock()
	vq.vQueue = append(vq.vQueue, vehicle)
}

func (vq *VehicleQueue) Dequeue() (shared.VehicleInfo, bool) {
	vq.mu.Lock()
	defer vq.mu.Unlock()
	if len(vq.vQueue) == 0 {
		return shared.VehicleInfo{}, false
	}
	vehicle := vq.vQueue[0]
	vq.vQueue = vq.vQueue[1:]
	return vehicle, true
}

func (vq *VehicleQueue) Size() int {
	vq.mu.Lock()
	defer vq.mu.Unlock()
	return len(vq.vQueue)
}

func (vq *VehicleQueue) Peek() (shared.VehicleInfo, bool) {
	vq.mu.Lock()
	defer vq.mu.Unlock()
	if len(vq.vQueue) == 0 {
		return shared.VehicleInfo{}, false
	}
	return vq.vQueue[0], true
}

type LaneQueue struct {
	lInfo []shared.LaneInfo
	mu    sync.Mutex
}

func InitLaneQueue() *LaneQueue {
	lanes := []string{"AL1", "AL2", "BL1", "BL2", "CL1", "CL2", "DL1", "DL2"}
	lq := &LaneQueue{}

	for _, lane := range lanes {
		lq.lInfo = append(lq.lInfo, shared.LaneInfo{Lane: lane, Count: 0})
	}
	return lq
}

func (lq *LaneQueue) IncrementLane(index int) {
	lq.mu.Lock()
	defer lq.mu.Unlock()
	if index >= 0 && index < len(lq.lInfo) {
		lq.lInfo[index].Count++
	}
}

func (lq *LaneQueue) DecrementLane(index int) {
	lq.mu.Lock()
	defer lq.mu.Unlock()
	if index >= 0 && index < len(lq.lInfo) && lq.lInfo[index].Count > 0 {
		lq.lInfo[index].Count--
	}
}

func (lq *LaneQueue) GetPriorityLane() (int, bool) {
	lq.mu.Lock()
	defer lq.mu.Unlock()
	for i, info := range lq.lInfo {
		if info.Count > 10 {
			return i, true
		}
	}
	return -1, false
}
