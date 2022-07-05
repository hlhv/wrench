package main

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"os"
	"strconv"
)

func newKey(text *string, cost *int) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(*text), *cost)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(hashed))
}

func validateKeyCost(args []string) (err error) {
	if len(args) < 1 {
		return errors.New("bug")
	}
	num, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}

	if num < bcrypt.MinCost {
		return errors.New(fmt.Sprint(
			"Cost ", num, " is too low, must be at least ",
			bcrypt.MinCost,
		))
	}

	if num > bcrypt.MaxCost {
		return errors.New(fmt.Sprint(
			"Cost ", num, " is too high, must be at most ",
			bcrypt.MaxCost,
		))
	}

	return nil
}
