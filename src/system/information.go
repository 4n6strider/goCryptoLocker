/*
  This package allows you to get a username
  and a list of drives.
*/

package information

import (
	"fmt"
	"os"
    "os/user"
    "path/filepath"
)


// Args
func Args() []string {
	return os.Args
}

// Exit
func Exit(code int) {
	os.Exit(code)
}

// Get user directory
func GetUserDir() string {
	user, err := user.Current()
    if err != nil {
        fmt.Println(err)
    }
	return user.HomeDir
}

// Get current executable file location
func ExecutableLocation() string {
	file, err := filepath.Abs(Args()[0])
    if err != nil {
        fmt.Println(err)
    }
	return file
}