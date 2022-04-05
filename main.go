package main

import (
        "os"
        "fmt"
        "golang.org/x/crypto/bcrypt"
        "github.com/akamensky/argparse"
)

func main () {
        parser := argparse.NewParser (
                "", "Swiss army knife for configuring HLHV")

        newKeyCommand := parser.NewCommand ("newkey", "Generate a new key")
        keyText := newKeyCommand.String ("t", "text", &argparse.Options {
                Required: true,
                Help:     "Contents of the key as text",
        })
        keyCost := newKeyCommand.Int ("c", "cost", &argparse.Options {
                Required: false,
                Validate: validateKeyCost,
                Help:     "Cost of the key",
                Default:  bcrypt.DefaultCost,
        })

        addUserCommand := parser.NewCommand (
                "adduser", "Add a user for the specified cell")
        addUserCell := addUserCommand.String ("c", "cell", &argparse.Options {
                Required: false,
                Validate: validateCell,
                Help:     "Name of the cell to add a user for",
                Default:  "queen",
        })
        
        delUserCommand := parser.NewCommand (
                "deluser", "Deletes the user for the specified cell")
        delUserCell := delUserCommand.String ("c", "cell", &argparse.Options {
                Required: false,
                Help:     "Cell who's user will be deleted",
                Default:  "queen",
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
        }
}
