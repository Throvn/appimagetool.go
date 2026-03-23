package appimagetoolgo

import (
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Taken from https://stackoverflow.com/a/28371044/10408987
func copyFile(src string, dst string) {
	// Read all content of src to data, may cause OOM for a large file.
	data, err := os.ReadFile(src)
	Check(err)
	// Write data to dst
	err = os.WriteFile(dst, data, 0644)
	Check(err)
}

func safeFileBase(path string) string {
	dir, file := filepath.Split(path)

	// Remove file extension, if exists.
	fileParts := strings.Split(file, ".")
	if len(fileParts) > 1 {
		file = strings.Join(fileParts[:len(fileParts)-1], ".")
	}

	return filepath.Join(dir, file)
}

func CreateAppImage(path string, appImageEngine string) {
	fileName := safeFileBase(path) + ".AppImage"

	copyFile(appImageEngine, fileName)

	outFileName := safeFileBase(path) + ".squashfs"

	os.Remove(outFileName)
	CreateSquashFSFromFolder(path, outFileName)

	AppendToFile(outFileName, fileName)
	hash := CalculateMD5(fileName)
	UpdateMD5(fileName, hash)
	fmt.Println(hex.EncodeToString(hash))
	MakeExecutable(fileName)
	hash = CalculateSha256(fileName)

	err := os.Remove(outFileName)
	Check(err)

	os.Remove(appImageEngine)
	Check(err)

	fmt.Printf("Created %s\n", fileName)
}
