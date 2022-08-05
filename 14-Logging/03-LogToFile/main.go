package main

import (
	"logToFile/core"

	"github.com/pkg/errors"
)

type myStruct struct {
	Name string
}

var logger = core.NewFileLogger()

func main() {

	logger.Print("test")
	myStruct := myStruct{Name: "test"}
	logger.Info().Interface("Body", myStruct).Str("Category", "Search").
		Int("DurationTime", 80).
		Msg("Searching for a thing")

	err := errors.New("this is an error")

	logger.Error().Err(err).Msg("Error occurred")

	err = func3()
	if err != nil {

		logger.Error().Stack().Err(err).Msg("Error occurred (func3 error)")
	}
}

func func1() error {
	return errors.New("this is an error (func1)")
}

func func2() error {
	err := func1()
	if err != nil {
		return err
	}
	return nil
}

func func3() error {
	err := func2()
	if err != nil {
		return err
	}
	return nil
}

// trace
// debug
// info
// warn
// error
// fatal(critical)
// panic
