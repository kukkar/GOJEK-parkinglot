package lib

import (
	"GOJEK-parkinglot/car"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"GOJEK-parkinglot/parkingLot"

	"github.com/fatih/color"
)

//map of allowed commands along with the arguments to read
var allowedCommands = map[string]int{
	"create_parking_lot": 1,
	"park":               2,
	"leave":              1,
	"status":             0,
	"registration_numbers_for_cars_with_colour": 1,
	"slot_numbers_for_cars_with_colour":         1,
	"slot_number_for_registration_number":       1,
}

var argumentsErrors = map[string]error{
	"create_parking_lot": fmt.Errorf("Provide Parking slots to fill in"),
	"park":               fmt.Errorf("Needed Two parameter to park RegNumber and Color"),
	"leave":              fmt.Errorf("Kindly provide slot number to leave"),
	"status":             fmt.Errorf("No Input Require"),
	"registration_numbers_for_cars_with_colour": fmt.Errorf("Provide color to get Car Details"),
	"slot_numbers_for_cars_with_colour":         fmt.Errorf("Provide color to get slot number"),
	"slot_number_for_registration_number":       fmt.Errorf("Provide registration Number to get slot Number"),
}
var red = color.New(color.FgRed).PrintfFunc()

const (
	UNSUPPORTED_COMMAND           = "Unsupported Command"
	UNSUPPORTED_COMMAND_ARGUMENTS = "Unsupported Command Arguments"
)

// Process the command taken in from file/stdin
// Separate the command and arguments for command
// Validate the command and then do the necessary action
func processCommand(command string) error {
	commandDelimited := strings.Split(command, " ")
	lengthOfCommand := len(commandDelimited)
	arguments := []string{}
	if lengthOfCommand < 1 {
		err := errors.New(UNSUPPORTED_COMMAND)
		red(err.Error())
		return err
	} else if lengthOfCommand == 1 {
		command = commandDelimited[0]
	} else {
		command = commandDelimited[0]
		arguments = commandDelimited[1:]
	}

	// check if command is one of the allowed commands
	if numberOfArguments, exists := allowedCommands[command]; exists {

		if len(arguments) != numberOfArguments {
			red(argumentsErrors[command].Error())
			return argumentsErrors[command]
		}

		w := &ErrWrapper{}

		// after validation of number of arguments per command, perform the necessary command
		switch command {
		case "create_parking_lot":
			if numberOfSlots, err := strconv.Atoi(arguments[0]); err != nil {
				red(err.Error())
				return err
			} else {
				return parkingLot.Initialize(numberOfSlots)
			}

		case "park":
			regNo := arguments[0]
			color := arguments[1]
			car := car.Create(regNo, color)

			return w.do(func() error {
				return parkingLot.Park(car)
			})

		case "leave":
			if slot, err := strconv.Atoi(arguments[0]); err != nil {
				red(err.Error())
				return err
			} else {
				return w.do(func() error {
					return parkingLot.Leave(slot)
				})
			}

		case "status":
			return w.do(func() error {
				return parkingLot.Status()
			})

		case "registration_numbers_for_cars_with_colour":
			color := arguments[0]
			return w.do(func() error {
				_, err := parkingLot.GetRegNosForCarsWithColor(color, true)
				return err
			})

		case "slot_numbers_for_cars_with_colour":
			color := arguments[0]
			return w.do(func() error {
				_, err := parkingLot.GetSlotNosForCarsWithColor(color)
				return err
			})

		case "slot_number_for_registration_number":
			regNo := arguments[0]
			return w.do(func() error {
				_, err := parkingLot.GetSlotNoForRegNo(regNo, true)
				return err
			})
		}
		return errors.New("Not Reachable Code")
	} else {
		err := errors.New(UNSUPPORTED_COMMAND)
		red(err.Error())

		return err
	}
}
