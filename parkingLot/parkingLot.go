package parkingLot

import (
	"GOJEK-parkinglot/helpers"
	"container/heap"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	. "GOJEK-parkinglot/car"
)

// Singleton Implementation of ParkingLot. Hence we use this variable to store the Instance of Parking Lot
var instance *parkingLot
var once sync.Once

// To solve multiple use cases, we use embedding to form a combination of data structures for a parking lot
// colorRegNoMap,regNoSlotMap,slotCarMap is implemented in such a way such that find and remove operations are done in O(1)
// emptySlots in implemented as Heap so that complexity of extract minimum and insert to log(n) on average. Since
// there wasn't much information as to if either of operations of leave and park could be more than the other
// hence a call was to taken to average out the optimality on both the operations.
// The respective explanation for each is mentioned below:
type parkingLot struct {
	// heap of empty slots to optimize to log(n) operations of allocating a smaller empty slot and vacating any given slot
	emptySlots helpers.IntHeap
	// max size of the parking lot
	maxSize int
	// check if parking lot has been initialized or not
	isParkingLotCreated bool
	// Map of registration number to the slot for answering queries of "slot_number_for_registration_number" efficiently in O(1)
	regNoSlotMap map[string]int
	// Map of Slots to Cars for maintaining information as to which car is parked at which slot
	slotCarMap map[int]Car
	// Map of Car Color to registration number Hast set(implemented as another map of string to bool)
	// Used for answering queries of "slot_number_for_registration_number" efficiently in O(1)
	colorRegNoMap map[string]map[string]bool
}

// Get Instance of Parking lot
func GetInstance() *parkingLot {
	once.Do(func() {
		instance = &parkingLot{}
	})
	return instance
}

// Initialize all the data structures and fix the value of MaxSlots of Parking Lot
func Initialize(numberOfSlots int) error {
	this := GetInstance()
	if _, err := this.verifySlotInitialization(numberOfSlots); err != nil {
		return err
	}

	this.emptySlots = helpers.IntHeap{}
	i := 1
	for i <= numberOfSlots {
		this.emptySlots = append(this.emptySlots, i)
		i++
	}

	heap.Init(&this.emptySlots)
	this.slotCarMap = map[int]Car{}
	this.colorRegNoMap = map[string]map[string]bool{}
	this.regNoSlotMap = map[string]int{}
	this.maxSize = numberOfSlots
	this.isParkingLotCreated = true

	fmt.Println("Created a parking lot with " + strconv.Itoa(numberOfSlots) + " slots")
	return nil
}

// Given a Car Object, park it at the nearest slot
// The overall complexity of this function is O(log n)
// The complexity arises from the fact that Extract-Min from a min heap of emptySlots is O(log n)
// Rest all operations of HashMaps are O(1)
func Park(car Car) error {
	this := GetInstance()
	if _, err := this.isparkingLotCreated(); err != nil {
		return err
	}

	// Validate if the parking lot is already full or not
	if _, err := this.isParkingLotFull(); err != nil {
		return err
	} else {

		// Validate if the car is already parked somewhere to check mistyped input
		if slot, _ := GetSlotNoForRegNo(car.GetRegNo(), false); slot == 0 {
			// Complexity : O(log n)
			emptySlot := heap.Pop(&this.emptySlots)
			this.mapRegNoToSlot(car.GetRegNo(), emptySlot.(int))
			this.mapSlotToCar(emptySlot.(int), car)
			this.mapColorToRegNo(car.GetColor(), car.GetRegNo())
			fmt.Println("Allocated slot number: " + strconv.Itoa(emptySlot.(int)))
			return nil
		} else {
			err := errors.New("Car with this registration number already parked at slot: " + strconv.Itoa(slot))
			return err
		}

	}
}

// Given a slot x, vacate it
// The overall complexity of this function is O(log n)
// The complexity arises rom the fact that insert to a min heap is O(log n)
// All other operations on HashMaps are O(1)
func Leave(slot int) error {
	this := GetInstance()
	if _, err := this.isparkingLotCreated(); err != nil {
		return err
	}

	// Validate if the slot has some car parked there or not
	if car, exists := this.slotCarMap[slot]; exists {
		// Complexity : O(log n)
		heap.Push(&this.emptySlots, slot)
		this.unmapRegNo(car.GetRegNo())
		this.unmapRegNoFromColorMap(car.GetColor(), car.GetRegNo())
		this.unmapSlot(slot)
		fmt.Println("Slot number " + strconv.Itoa(slot) + " is free")
		return nil
	} else {
		err := errors.New(NOT_FOUND_ERROR)
		return err
	}
}

// Output the status of the parking lot
func Status() error {
	this := GetInstance()
	if _, err := this.isparkingLotCreated(); err != nil {
		return err
	}

	fmt.Println("Slot No.\tRegistration No.\tColour")
	i := 1
	for i <= this.maxSize {
		if car, exists := this.slotCarMap[i]; exists {
			fmt.Println(strconv.Itoa(i) + "\t" + strings.ToUpper(car.GetRegNo()) + "\t" + helpers.ToCamelCase(car.GetColor()))

		}
		i++
	}
	return nil
}

// Fetch the Registration Numbers of Cars with Given Color
// Complexity is O(1) since we have stored the data in a HashMap->HashSet
func GetRegNosForCarsWithColor(color string, print bool) ([]string, error) {
	this := GetInstance()
	if _, err := this.isparkingLotCreated(); err != nil {
		return []string{}, err
	}

	regNoSlice := []string{}
	if regNoMap, exists := this.colorRegNoMap[color]; exists {
		for regNo, _ := range regNoMap {
			regNoSlice = append(regNoSlice, strings.ToUpper(regNo))
		}
	}

	if len(regNoSlice) > 0 {
		if print {
			fmt.Println(strings.Join(regNoSlice, ","))
		}
		return regNoSlice, nil
	} else {
		err := errors.New(NOT_FOUND_ERROR)
		if print {
			fmt.Println(err.Error())
		}
		return regNoSlice, err
	}
}

// Fetch Slot Numbers for Cars with Given Color
// Complexity is O(m). We first we fetch registration numbers from HashSet which is O(1)
// and then we fetch slots from registration numbers which again is HashMap. So if there are
// m cars, it will take m iterations to fetch slot for each registration number
func GetSlotNosForCarsWithColor(color string) ([]int, error) {
	this := GetInstance()
	if _, err := this.isparkingLotCreated(); err != nil {
		return []int{}, err
	}

	regNos, _ := GetRegNosForCarsWithColor(color, false)
	slotsString := []string{}
	slots := []int{}
	for _, regNo := range regNos {
		if slot, exists := this.regNoSlotMap[regNo]; exists {
			slotsString = append(slotsString, strconv.Itoa(slot))
			slots = append(slots, slot)
		}
	}

	if len(slots) > 0 {
		fmt.Println(strings.Join(slotsString, ","))
		return slots, nil
	} else {
		err := errors.New(NOT_FOUND_ERROR)
		return slots, err
	}

}

// Fetch slot number for given rgistration number
// Complexity is O(1) since it is derived from a HashSet
func GetSlotNoForRegNo(regNo string, print bool) (int, error) {
	this := GetInstance()
	if _, err := this.isparkingLotCreated(); err != nil {
		return 0, err
	}

	// Validate if there is any car with given reg no
	if slot, exists := this.regNoSlotMap[regNo]; exists {
		if print {
			fmt.Println(slot)
		}
		return slot, nil
	} else {
		err := errors.New(NOT_FOUND_ERROR)
		return 0, err
	}
}

// Initialize the parking lot in package init()
// Singleton based initialization for Parking Lot Instance
func init() {
	instance = GetInstance()
}
