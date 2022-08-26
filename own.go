package main

import "os"
import "fmt"
import "strconv"
import "os/user"
import "path/filepath"

func doOwn(path string, name string) {
	fullName := "hlhv-" + name

	userInfo, err := user.Lookup(fullName)
	uid, _ := strconv.Atoi(userInfo.Uid)
	gid, _ := strconv.Atoi(userInfo.Gid)

	if err != nil {
		fmt.Println("ERR could not get user info.")
		return
	}

	err = filepath.Walk(path, func(
		filePath string,
		file os.FileInfo,
		err error,
	) error {
		if err != nil {
			fmt.Println("ERR could not traverse filesystem:", err)
		}

		err = os.Chown(filePath, uid, gid)
		if err != nil {
			fmt.Println(
				"ERR could not take ownership of file:",
				err)
		}

		err = os.Chmod(filePath, 0770)
		if err != nil {
			fmt.Println(
				"ERR could not change mode of file:",
				err)
		}
		return nil
	})

	if err != nil {
		fmt.Println("ERR could not traverse filesystem.")
		return
	}
}
