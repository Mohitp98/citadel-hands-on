// error is an interface that have one function of
// type error interface {
// Error() string
// }

package main

import (
	"errors"
	"fmt"
)

func getError(b bool) (err error) {
	if b == true {
		// err = fmt.Errorf("this is an error :( ")
		err = errors.New("this is an error :(")
	}
	return
}

type invalidInput struct {
	filledName string
}

func (i invalidInput) Error() string {
	return fmt.Sprintf("invalid input on field '%v'", i.filledName)
}

func validateInput(field, value string) (err error) {
	if value == "" {
		err = invalidInput{filledName: field}
	}
	return
}

func main() {
	// {
	// 	err := getError(true)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }

	// {
	// 	// call to validate input func with valid input
	// 	err := validateInput("firstname", "mohit")
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }

	// {
	// 	// call to validate input func with invalid input
	// 	err := validateInput("lastname", "")
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }

	// {
	// 	// case 1
	// 	//check on error value
	// 	err := inValidLength("")
	// 	if errors.Is(err, errInvalidLength) {
	// 		fmt.Println("we know err value using 'erros.IS")
	// 	} else {
	// 		fmt.Println("we do not know about this error!")
	// 	}
	// }

	{
		// case 2
		//check on error value
		err := inValidLength("mohit")
		if errors.Is(err, errInvalidLength) {
			fmt.Println("we know err value using 'erros.IS")
		} else {
			fmt.Println("we do not know about this error!")
		}
	}
}

var errInvalidLength = errors.New("Invalid length")

func inValidLength(s string) error {
	if len(s) < 1 { // case 1
		return errInvalidLength
	} else if s == "error" {
		return errors.New("Invalid input")
	}
	return nil
}
