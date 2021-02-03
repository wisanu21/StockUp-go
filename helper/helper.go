package helper

import (

	// "gitlab.com/fireexitco/cmu-aqi-monitor/backend/config"
	// "gitlab.com/fireexitco/cmu-aqi-monitor/backend/models"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ" // 52 possibilities
	letterIdxBits = 6                                                      // 6 bits to represent 64 possibilities / indexes
	letterIdxMask = 1<<letterIdxBits - 1                                   // All 1-bits, as many as letterIdxBits
)

func RandStringRunes(n int) string {
	result := make([]byte, n)
	var random = 1.3
	for {
		bufferSize := int(float64(n) * random)
		for i, j, randomBytes := 0, 0, []byte{}; i < n; j++ {
			if j%bufferSize == 0 {
				randomBytes = SecureRandomBytes(bufferSize)
			}
			if idx := int(randomBytes[j%n] & letterIdxMask); idx < len(letterBytes) {
				result[i] = letterBytes[idx]
				i++
			}
		}
		// var tokenData models.UserToken
		// random = random + 0.3
		// newToken := string(result)
		// config.GetDB().Where("token = \"" + newToken + "\"").Find(&tokenData)
		// fmt.Println(tokenData)
		// if tokenData.Token == "" {
		break
		// }
		// fmt.Println(string(result))
	}

	return string(result)

}

func SecureRandomBytes(length int) []byte {
	var randomBytes = make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatal("Unable to generate random bytes")
	}
	return randomBytes
}

var path = Concatenation(Concatenation("logs/", time.Now().Format("01-02-2006")), ".txt")

func Log(text string) {
	// path =
	if !FileExists(path) {
		CreateFile(path)
	}

	f, err := os.OpenFile(path,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	logger := log.New(f, "", log.LstdFlags)
	logger.Println(text)
}

func CreateFile(path string) {
	// check if file exists

	var _, err = os.Stat(path)
	// create file if not exists
	if os.IsNotExist(err) {
		var file, errr = os.Create(path)
		if errr == nil {
			return
		}
		defer file.Close()
	}

	fmt.Println("File Created Successfully", path)
}

func Concatenation(str1, srtr2 string) string {
	var buf bytes.Buffer
	buf.WriteString(str1)
	// append string 'world!' to buf var
	buf.WriteString(srtr2)
	return buf.String()
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
