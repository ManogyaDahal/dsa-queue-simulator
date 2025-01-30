package main

/* 
  Vehicle Direction: 
  Initiates left as 0 and increments other by 1 in order
*/
const(
  left = iota
  right 
  down 
  up
)

/* Priority Vehicles Lanes */
const(
  laneA2 = iota 
  laneB2 
  laneC2
  laneD2
)

type vehicle struct {
  direction int 
} 

/* Note: Vehicle queue is circular queue */
type vehicleQueue struct{
  front int 
  rear int 
  size int
  vehicleQueueSize []vehicle
}

type laneData struct{
  lane int 
  numberOfVehicles int
  Priority int
}

type laneQueue struct{
  lane laneData
  size int
}

/* @FUNCTION DEFINATIONS */ 
/* Vehicle Queue */

func (vecQueue *vehicleQueue) initVehicleQueue() {
  vecQueue.front = -1
  vecQueue.rear = -1
  vecQueue.size = -1
}
func (vecQueue *vehicleQueue) isFullVehicleQueue(){}
func (vecQueue *vehicleQueue) isEmptyVehicleQueue() {}
func (vecQueue *vehicleQueue) enqueVehicleQueue() {}
func (vecQueue *vehicleQueue) dequeVehicleQueue() {}

/* Lane Queue */
func (lanQueue *laneQueue) initLaneQueue(){}
func (lanQueue *laneQueue) isFullLaneQueue(){}
func (lanQueue *laneQueue) isEmptyLaneQueue(){}
func (lanQueue *laneQueue) enqueLaneQueue(){}
func (lanQueue *laneQueue) dequeLaneQueue(){}







