package main

import (
	"GOJEK-parkinglot/lib"
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 {
		// take commands from file
		fileName := argsWithoutProg[0]
		lib.ReadAndProcessFromFile(fileName)
	} else {
		fmt.Println("\n INTERACTIVE COMMANDS :- \n CREATE_PARKING_LOT <numberOfslots> \n PARK <carNumber> <color> \n LEAVE <slotNumber> \n STATUS \n REGISTRATION_NUMBERS_FOR_CARS_WITH_COLOUR <colorOfCar>  \n SLOT_NUMBERS_FOR_CARS_WITH_COLOUR <colorOfCar> \n SLOT_NUMBER_FOR_REGISTRATION_NUMBER <regNumber> \n ################################################### \n ENTER YOUR INPUT \n #####################################")
		//We need to make it interactive session now
		lib.ReadAndProcessStdIn()
	}
}
