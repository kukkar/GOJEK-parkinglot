package parkingLot

import (
	. "GOJEK-parkinglot/car"
	"errors"
)

// Declaration of Error Constants
const (
	PARKING_LOT_FULL_ERROR        = "Sorry, parking lot is full"
	PARKING_LOT_NOT_CREATED_ERROR = "Parking Lot not created"
	NOT_FOUND_ERROR               = "Not found"
	WRONG_SIZE_PARKING_LOT_ERROR  = "Parking Lot of Size <= 0 cannot be created"
)

// Helper function to add to HashMap, mapping of RegNo to Slot
func (this *parkingLot) mapRegNoToSlot(regNo string, slot int) {
	this = instance
	this.regNoSlotMap[regNo] = slot
}

// Helper function to remove from HashMap, mapping of RegNo to Slot
func (this *parkingLot) unmapRegNo(regNo string) {
	delete(this.regNoSlotMap, regNo)
}

// Helper function to add to HashMap, mapping of slot to Car
func (this *parkingLot) mapSlotToCar(slot int, car Car) {
	this.slotCarMap[slot] = car
}

// Helper function to remove from HashMap, mapping of slot to Car
func (this *parkingLot) unmapSlot(slot int) {
	delete(this.slotCarMap, slot)
}

// Helper function to add to HashSet at given color key in the HashMap
func (this *parkingLot) mapColorToRegNo(color string, regNo string) {
	_, exists := this.colorRegNoMap[color]
	if exists {
		this.colorRegNoMap[color][regNo] = true
	} else {
		this.colorRegNoMap[color] = map[string]bool{regNo: true}
	}
}

// Helper function to remove from HashSet at given color key in the HashMap
func (this *parkingLot) unmapRegNoFromColorMap(color string, regNo string) {
	delete(this.colorRegNoMap[color], regNo)
}

// Verify if parking lot is full
func (this *parkingLot) isParkingLotFull() (bool, error) {
	if this.emptySlots.Len() == 0 {
		err := errors.New(PARKING_LOT_FULL_ERROR)
		return true, err
	}
	return false, nil
}

// Verify if parking lot is created
func (this *parkingLot) isparkingLotCreated() (bool, error) {
	if !this.isParkingLotCreated {
		err := errors.New(PARKING_LOT_NOT_CREATED_ERROR)
		return false, err
	}
	return true, nil
}

// Verify if NumberOfSlots is correct number or not
func (this *parkingLot) verifySlotInitialization(numberOfSlots int) (bool, error) {
	if numberOfSlots <= 0 {
		err := errors.New(WRONG_SIZE_PARKING_LOT_ERROR)
		return false, err
	}
	return true, nil
}

// Method to get car at given slot
func (this *parkingLot) getCarAtSlot(slot int) (Car, bool) {
	car, exists := this.slotCarMap[slot]
	return car, exists
}
