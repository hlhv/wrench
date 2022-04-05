package main

import (
        "os"
        "fmt"
        "regexp"
        "errors"
        "os/exec"
)

func doAddUser (name string) {
        fullName := "hlhv-" + name

        var cmd *exec.Cmd

        adduser, err := exec.LookPath("adduser")
        if err == nil {
                // we are using adduser
                // adduser: -SHD
                cmd = exec.Command(adduser, fullName, "-SHD")
        } else {
                useradd, err := exec.LookPath("useradd")
                if err == nil {
                        // we are using useradd
                        // useradd: -rUM --shell /sbin/nologin
                        cmd = exec.Command (
                                useradd, fullName, "-rUM",
                                "--shell", "/sbin/nologin")
                }
        }

        if cmd == nil {
                fmt.Println("ERR your system does not support adding users.")
                os.Exit(1)
        }
        
        output, err := cmd.CombinedOutput()
        if err != nil {
                fmt.Println("ERR could not add user:", string(output))
                os.Exit(1)
        }
}

func doDelUser (name string) {
        fullName := "hlhv-" + name

        var cmd *exec.Cmd

        deluser, err := exec.LookPath("deluser")
        if err == nil {
                // we are using deluser
                // deluser: --remove-home
                cmd = exec.Command(deluser, fullName, "--remove-home")
        } else {
                userdel, err := exec.LookPath("userdel")
                if err == nil {
                        // we are using userdel
                        // userdel: -r
                        cmd = exec.Command(userdel, fullName, "-r")
                }
        }

        if cmd == nil {
                fmt.Println("ERR your system does not support deleting users.")
                os.Exit(1)
        }
        
        output, err := cmd.CombinedOutput()
        if err != nil {
                fmt.Println("ERR could not delete user:", string(output))
                os.Exit(1)
        }
}

func validateCell (args []string) (err error) {
        if len(args) < 1 { return errors.New("bug") }
        if err != nil { return err }
        
        cellRegex := regexp.MustCompile("^[a-z][-a-z0-9]*$")
        if !cellRegex.MatchString(args[0]) {
                return errors.New (
                        "\"" + args[0] + "\" is not a valid cell name. Must " +
                        "only contain lowercase letters a-z, dashes, numbers " +
                        "0-9, dashes, and must start with a lowercase letter.")
        }
        
        return nil
}
