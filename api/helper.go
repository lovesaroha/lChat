/*  Love Saroha
    lovesaroha1994@gmail.com (email address)
    https://www.lovesaroha.com (website)
    https://github.com/lovesaroha  (github)
*/
package api

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"regexp"
	"time"
	"unsafe"
)

// ValidateEmail function.
func inValidEmailAddress(s string) bool {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return !re.MatchString(s)
}

// This function encode given string in sha1.
func getSHA1Hash(s string) string {
	h := sha1.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// Constant letters defined for random string.
const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

// This function generate random string.
func generateRandomString(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return *(*string)(unsafe.Pointer(&b))
}

var letterAndNumbers = [62]string{"a", "9", "b", "0", "c", "1", "d", "2", "e", "3", "f", "4", "g", "5", "h", "6", "i", "7", "j", "8", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
var numberIndex int

// This function generate unique id.
func generateID() string {
	var currentTime = time.Now().UTC()
	numberIndex = (numberIndex + 1) % 62
	seconds := currentTime.Second()
	minutes := currentTime.Minute()
	hour := currentTime.Hour()
	day := currentTime.Day()
	month := currentTime.Month()
	year := (currentTime.Year() % 2000) % 62
	return generateRandomString(3) + letterAndNumbers[year] + letterAndNumbers[month] + letterAndNumbers[day] + letterAndNumbers[hour] + letterAndNumbers[seconds] + letterAndNumbers[minutes] + letterAndNumbers[numberIndex]
}
