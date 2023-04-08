/*
 MODULE: TTKUtils.go
 AUTHOR: Leo Schneider <schleo@outlook.com>
 DATE  : August 2017
 INFO  : This module handles generic and application wide helper functions
*/

package ttktools

import (
	"bufio"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base32"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// COLOR CONSTANTS
var (
	RED     = "\033[1;31m"
	GREEN   = "\033[1;32m"
	PURPLE  = "\033[1;35m"
	CYAN    = "\033[1;36m"
	BLUE    = "\033[1;34m"
	GRAY    = "\033[1;30m"
	LGREEN  = "\033[0;32m"
	LRED    = "\033[0;31m"
	LPURPLE = "\033[0;35m"
	NC      = "\033[0m"
)

// TermMessage prints a message on the terminal usinhg the specified color
func TermMessage(COLOR, message string) {
	fmt.Printf(COLOR + message + NC)
}

// PrintError prints a message on the terminal with a RED error indicator
func PrintError(message string) {
	fmt.Print(GRAY + "[" + LRED + "ERROR" + GRAY + "] " + NC + message + "\n")
}

// PrintInfo prints a message on the terminal using a GREEN info indicator
func PrintInfo(message string) {
	fmt.Print("\r" + GRAY + "[" + LGREEN + "INFO " + GRAY + "] " + NC + message + "\n")
}

// PrintMsg prints a message on the terminal using a grey title between brackets
func PrintMsg(color string, title string, message string) {
	fmt.Print(GRAY + "[" + color + title + GRAY + "] " + NC + message + "\n")
}

// PrintMsgInv prints a message on the terminal with a grey title and the message between brackets
func PrintMsgInv(color string, title string, message string) {
	fmt.Print(color + title + GRAY + " [ " + NC + message + GRAY + " ] " + NC)
}

// PrintData prints a message on the terminal with a grey title and the message between brackets and line break
func PrintData(color string, title string, message string) {
	fmt.Print(color + title + GRAY + " [ " + NC + message + GRAY + " ] " + NC + "\n")
}

// Run - Execute a filesystem command
func Run(command string, arguments string) (string, error) {
	out, err := exec.Command(command, arguments).Output()
	return string(out), err
}

// B64Encode returns a string encoded in Base64
func B64Encode(text string) string {
	return strings.ReplaceAll(base64.StdEncoding.EncodeToString([]byte(text)), "=", "")
}

// B64Decode returns a string decoded from Base64
func B64Decode(text string) (string, error) {
	data, err := base64.RawURLEncoding.DecodeString(text)
	if err != nil {
		data = []byte("")
	}
	return string(data), err
}

//B32Encode returns a string encoded in Base32
func B32Encode(text string) string {
	return base32.StdEncoding.EncodeToString([]byte(text))

}

//B32Decode returns a string decoded from Base32
func B32Decode(text string) (string, error) {
	data, err := base32.StdEncoding.WithPadding(-1).DecodeString(text)
	if err != nil {
		data = []byte("")
	}
	return string(data), err
}

// HMAC256 returns a HMAC SHA256 Signature
func HMAC256(payload string, secret string) string {
	key := []byte(secret)
	sig := hmac.New(sha256.New, key)
	sig.Write([]byte(payload))
	soma := string(sig.Sum(nil)[:])
	return B64Encode(soma)
}

//GetData load httprequest strings into struct fields based on their types
func GetData(data interface{}, value string) {
	if value != "" {
		switch data := data.(type) {
		case *int:
			*data, _ = strconv.Atoi(value)
		case *bool:
			*data, _ = strconv.ParseBool(value)
		case *string:
			*data = value
		}
	}
}

func checkOS() {
	if runtime.GOOS == "windows" {
		RED = ""
		GREEN = ""
		PURPLE = ""
		CYAN = ""
		BLUE = ""
		GRAY = ""
		LGREEN = ""
		LRED = ""
		LPURPLE = ""
		NC = ""
	}
}

func errCheck(source string, e error) bool {
	if e != nil {
		return false
	}
	return true
}

// Banner prints on the terminal the application name and version
func Banner(appname string, version string, build string, copy string) {
	TermMessage(CYAN, appname+" v"+version+" ("+build+")"+" "+GRAY+"("+runtime.GOOS+"/"+runtime.GOARCH+"/"+strconv.Itoa(strconv.IntSize)+"bits/"+strconv.Itoa(runtime.NumCPU())+" cores)\n")
	TermMessage(LGREEN, "Copyright "+copy+"\n\n")
}

// FileWrite writes a line to a text file
func FileWrite(file string, line string, append bool) bool {
	var f *os.File
	var err error
	if append {
		f, err = os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	} else {
		f, err = os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	}
	if !errCheck("TTKUtils", err) {
		return false
	}
	if _, err := f.Write([]byte(line + "\n")); err != nil {
		return errCheck("TTKUtils", err)
	}
	if err := f.Close(); err != nil {
		return errCheck("TTKUtils", err)
	}
	return true
}

// FileRead read a text file onto a string
func FileRead(file string) string {
	dat, err := ioutil.ReadFile(file)
	errCheck("TTKUtils", err)
	return string(dat)
}

// Dbg quick and dirty output of values on terminal
func Dbg(a ...interface{}) {
	fmt.Println(a...)
}

// IntToHex converts an int64 to a byte array
func IntToHex(num int) string {
	return fmt.Sprintf("%x", num)
}

// GetTimestamp retuns the current unix date incremented by sec Seconds
func GetTimestamp(sec int) int {
	return int(time.Now().Unix()) + sec
}

// SHA256 returns a SHA256 hash for a give text
func SHA256(text string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(text)))
}

// SHA1B64 return a SHA1 hash for a give text encoded in Base64
func SHA1B64(text string) string {
	hasher := sha1.New()
	hasher.Write([]byte(text))
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}

// DateToUnix get a date in dd/mm/yyyy and returns its unix date
func DateToUnix(date string) int {
	const layout = "02/01/2006"
	tm, _ := time.Parse(layout, date)
	ret, _ := strconv.Atoi(fmt.Sprintf("%d", tm.Unix()))
	return ret
}

// UnixToDate gets a unix date and returns a string with dd/mm/yyyy
func UnixToDate(date int) string {
	i, _ := strconv.ParseInt(strconv.Itoa(date), 10, 64)
	x := time.Unix(i, 0).Format(time.RFC1123)
	return x
}

// GetLocalIP returns the local ip address
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

// GetIP returns the IP address of a given hostname
func GetIP(host string) string {
	addr, err := net.LookupIP(host)
	if err != nil {
		return ""
	}
	return addr[0].String()
}

// MapToQuerystring converts a map into a HTTP querystring
func MapToQuerystring(params map[string]string) string {
	var ret string
	for k, v := range params {
		ret += k + "=" + v + "&"
	}
	return ret[0 : len(ret)-1]
}

// IsNumeric check if a specific string represents a number
func IsNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 32)
	return err == nil
}

// GetInput read user input from keyboard
func GetInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(LGREEN + prompt + NC)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(strings.TrimRight(input, "\n"))
}

// ShortTime accepts a unix date and returns the time in hh:mm:ss
func ShortTime(date int) string {
	i, err := strconv.ParseInt(strconv.Itoa(date), 10, 64)
	if err != nil {
		return ""
	}
	return time.Unix(i, 0).Format("15:04:05")
}

// GenCodeA generates a random alfabetic code of size n
func GenCodeA(n int) string {
	rn := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return GenCode(n, rn)
}

// GenCodeN generates a random numeric code of size n
func GenCodeN(n int) string {
	rn := "0123456789"
	return GenCode(n, rn)
}

// GenCodeAN generates a alphanumeric code of size n
func GenCodeAN(n int) string {
	rn := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return GenCode(n, rn)
}

// GenCode generates a code of size an based on the runes provided
func GenCode(n int, runes string) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune(runes)
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// Elapsed returns the milliseconds elapsed from start
func Elapsed(start time.Time) int64 {
	since := time.Since(start)
	return since.Milliseconds()
}

// GeoDistance returns the distance between two geographic coordinates
func GeoDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit ...string) float64 {
	const PI float64 = 3.141592653589793
	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)
	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)
	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)
	if dist > 1 {
		dist = 1
	}
	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515
	if len(unit) > 0 {
		if unit[0] == "K" {
			dist = dist * 1.609344
		} else if unit[0] == "N" {
			dist = dist * 0.8684
		}
	}
	return dist
}

// MemUsage returns a string with memory usage
func MemUsage() string {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return fmt.Sprintf("Alloc[%vM] TotalAlloc[%vM] Sys[%vM] NumGC[%v]", bToMb(m.Alloc), bToMb(m.TotalAlloc), bToMb(m.Sys), m.NumGC)
}

//byte to megabyte
func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
