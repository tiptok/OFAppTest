package cryptoT

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"runtime"
	"testing"
)

func TestMd5(t *testing.T) {
	h := md5.New()
	h.Write([]byte("tiptok"))
	h.Write([]byte("ccc"))
	log.Println(hex.EncodeToString(h.Sum(nil)))

	log.Println("Home Dri:" + UserHomeDir())
}

func UserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		log.Println("HOMEDRIVE:" + os.Getenv("HOMEDRIVE") + " HOMEPATH:" + os.Getenv("HOMEPATH"))
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}
