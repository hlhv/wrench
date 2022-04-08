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

        // BUSYBOX
        adduser, err := exec.LookPath("adduser")
        if err == nil {
                addgroup, _ := exec.LookPath("addgroup")
                tryCommand (exec.Command (addgroup, fullName, "-S"),
                        "could not add group")
                tryCommand (exec.Command (adduser, fullName, "-SHDG", fullName),
                        "could not add user")
                return
        }

        // GNU
        useradd, err := exec.LookPath("useradd")
        if err == nil {
                tryCommand (exec.Command (
                        useradd, fullName, "-rUM",
                        "--shell", "/sbin/nologin"), "could not add user")
                return
        }

        fmt.Println("ERR your system does not support adding users.")
        os.Exit(1)
}

func doDelUser (name string) {
        fullName := "hlhv-" + name

        // BUSYBOX
        deluser, err := exec.LookPath("deluser")
        if err == nil {
                tryCommand (exec.Command (deluser, fullName, "--remove-home"),
                        "could not delete user")
                return
        }

        // GNU
        userdel, err := exec.LookPath("userdel")
        if err == nil {
                tryCommand (exec.Command (userdel, fullName, "-r"),
                        "could not delete user")
                groupdel, _ := exec.LookPath("groupdel")
                tryCommand (exec.Command (groupdel, fullName),
                        "could not delete group")
                return
        }

        fmt.Println("ERR your system does not support deleting users.")
        os.Exit(1)
}

func doAuthUser (user string, name string) {
        fullName := "hlhv-" + name

        // BUSYBOX
        adduser, err := exec.LookPath("adduser")
        if err == nil {
                tryCommand (exec.Command (adduser, user, fullName),
                        "could not add user to group " + fullName)
                return
        }

        // GNU
        useradd, err := exec.LookPath("usermod")
        if err == nil {
                tryCommand (exec.Command (useradd, "-a", "-g", fullName, user),
                        "could not add user to group " + fullName)
                return
        }

        fmt.Println("ERR your system does not support modifying users.")
        os.Exit(1)
}

func tryCommand (cmd *exec.Cmd, failReason string) {
        output, err := cmd.CombinedOutput()
        if err != nil {
                fmt.Println("ERR " + failReason + ":", string(output))
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
