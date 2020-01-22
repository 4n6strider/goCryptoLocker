/*
  This package allows you to get a username
  and a list of drives.
*/

package information

import (
	"fmt"
    "os/user"
)

// Get user directory
func GetUserDir() string {
	user, err := user.Current()
    if err != nil {
        fmt.Println(err)
    }
	return user.HomeDir
}