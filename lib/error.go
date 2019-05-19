package lib

import "fmt"

// Wrapper Error struct
type ErrWrapper struct {
	err error
}

// Wrapper function to output the error for a given function
func (ew *ErrWrapper) do(f func() error) error {
	ew.err = f()
	if ew.err != nil {
		fmt.Println(ew.err.Error())
		return ew.err
	}
	return nil
}
