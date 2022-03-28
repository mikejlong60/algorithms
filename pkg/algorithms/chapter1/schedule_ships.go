package chapter1

type ShipState = int //0 is at sea. Non-zero is port number
//This algorithm accomodates both weak and strong instabilities.
type Ship struct {
	Id               int
	ProposedSchedule []ShipState
	ActualSchedule   []ShipState
}

//Algorithm
//1. range loop over array of ships, index i
//2. For current ship range loop over all other ships and make sure there is not a port conflict with element i in any of their proposed schedules
//2.1  If there are no port conflicts, add ProposedSchedule element to end of ActualSchedule. A port conflict is same port at same element. At sea is no conflict
//2.2  Otherwise skip to next element in array of ships
//3 Return array of ships with ActualSchedule
func schedule(ships []*Ship) []*Ship {
	if len(ships) == 0 {
		return []*Ship{}
	}

	//Check the next ship's schedule to see if there are any conflicts
	fEarliestConflict := func(currentShipSchedule, otherShipSchedule []int, previousEarliestConflictDay int) int {
		for i, shipState := range currentShipSchedule {
			if shipState > 0 && shipState == otherShipSchedule[i] {
				if previousEarliestConflictDay > i {
					earliestConflict := i
					return earliestConflict
				}
			}
		}
		return previousEarliestConflictDay
	}

	//An invariant is that all ships have same calendar length reflected in the size of their proposed schedule
	for i, ship := range ships { //Range loop over array of all ships
		currentShipSchedule := ship.ProposedSchedule
		var earliestConflict = 0              //len(currentShipSchedule)
		for j := i + 1; j < len(ships); j++ { //For current ship iterate over all ships later in array and truncate ship's Proposed Schedule at earliest conflict
			otherShipSchedule := ships[j].ProposedSchedule
			earliestConflict = fEarliestConflict(currentShipSchedule, otherShipSchedule, earliestConflict)
		}
		ship.ActualSchedule = ship.ProposedSchedule[0:earliestConflict] //TODO You have bug in this. For two ships with conflict this truncates the next ship too.
	}
	return ships
}
