package main

import (
        "os"
        "fmt"
        "errors"
        "strconv"
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

        err := parser.Parse(os.Args)
        if err != nil {
                fmt.Print(parser.Usage(err))
                os.Exit(1)
        }

        if newKeyCommand.Happened() {
                newKey(keyText, keyCost)
        }
}

func newKey (text *string, cost *int) {
        hashed, err := bcrypt.GenerateFromPassword([]byte(*text), *cost)
        if err != nil {
                fmt.Errorf("%s", err)
                os.Exit(1)
        }

        fmt.Println(string(hashed))
}

func validateKeyCost (args []string) (err error) {
        if len(args) < 1 { return errors.New("bug") }
        num, err := strconv.Atoi(args[0])
        if err != nil { return err }

        if num < bcrypt.MinCost {
                return errors.New(fmt.Sprint (
                        "Cost ", num, " is too low, must be at least ",
                        bcrypt.MinCost,
                ))
        }
        
        if num > bcrypt.MaxCost {
                return errors.New(fmt.Sprint (
                        "Cost ", num, " is too high, must be at most ",
                        bcrypt.MaxCost,
                ))
        }
        
        return nil
}
