package main

import (
	"fmt"
	"github.com/akamensky/argparse"
	"golang.org/x/crypto/bcrypt"
	"os"
)

func main() {
	parser := argparse.NewParser(
		"", "Swiss army knife for configuring HLHV")

	newKeyCommand := parser.NewCommand("newkey", "Generate a new key")
	keyText := newKeyCommand.String("t", "text", &argparse.Options{
		Required: true,
		Help:     "Contents of the key as text",
	})
	keyCost := newKeyCommand.Int("c", "cost", &argparse.Options{
		Required: false,
		Validate: validateKeyCost,
		Help:     "Cost of the key",
		Default:  bcrypt.DefaultCost,
	})

	addUserCommand := parser.NewCommand(
		"adduser", "Add a user for the specified cell")
	addUserCell := addUserCommand.String("c", "cell", &argparse.Options{
		Required: false,
		Validate: validateCell,
		Help:     "Name of the cell to add a user for",
		Default:  "queen",
	})

	delUserCommand := parser.NewCommand(
		"deluser", "Deletes the user for the specified cell")
	delUserCell := delUserCommand.String("c", "cell", &argparse.Options{
		Required: false,
		Help:     "Cell who's user will be deleted",
		Default:  "queen",
	})

	authUserCommand := parser.NewCommand(
		"authuser",
		"authorizes a user to access files for the specified cell",
	)
	authUserCell := authUserCommand.String("c", "cell", &argparse.Options{
		Required: false,
		Help:     "Cell to grant access to",
		Default:  "queen",
	})
	authUserUser := authUserCommand.String("u", "user", &argparse.Options{
		Required: true,
		Help:     "User to be given access",
	})

	ownCommand := parser.NewCommand(
		"own",
		"recursively gives ownership of a file or directory to the " +
		"specified cell",
	)
	ownCell := ownCommand.String("c", "cell", &argparse.Options{
		Required: false,
		Help:     "Cell to give ownership to",
		Default:  "queen",
	})
	ownFile := ownCommand.String("f", "file", &argparse.Options{
		Required: false,
		Help:     "File or directory to own",
		Default:  ".",
	})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}

	if newKeyCommand.Happened() {
		newKey(keyText, keyCost)
	} else if addUserCommand.Happened() {
		doAddUser(*addUserCell)
	} else if delUserCommand.Happened() {
		doDelUser(*delUserCell)
	} else if authUserCommand.Happened() {
		doAuthUser(*authUserUser, *authUserCell)
	} else if ownCommand.Happened() {
		doOwn(*ownFile, *ownCell)
	}
}
