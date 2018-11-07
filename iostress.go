package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Version constant
const Version string = "1.0.0-beta"

// help message string
var helpMsg = `iostress = Creating stress on disks
usage: iostress [options]
options:
`

// print help
func printHelp() {
	fmt.Println(helpMsg)
	flag.PrintDefaults()
}

// Print Version
func printVersion() {
	fmt.Printf("cosuite api version %v\n", Version)
}

//TotalP stands for total active proccess number
var TotalP int

//CreateTempFolder cheks the temp directory. if it is not, creates a temp folder in the application running directory.
func CreateTempFolder() {
	path := "./temp"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0644)
	}
}

//Run starts write, read and delete proccess
//useRandomBytes bool, default false
//fileName : temp file
//arrayLength : length of byte array
//bufferSize : buffer size for writing proccess
//stressLevel : if it's full created asyncronized data production, default hard
func Run(useRandomBytes bool, fileName string, arrayLength int, bufferSize int, stressLevel string) {
	TotalP++
	var bArray []byte
	if useRandomBytes {
		bArray = GetByteArray(arrayLength, true)
	} else {
		bArray = GetByteArray(arrayLength, false)
	}
	if bufferSize == 0 {
		CreateFile(bArray, fileName)
	} else {
		CreateFileBufio(bArray, fileName, bufferSize)

	}
	if stressLevel == "hard" {
		DummyPro()
	}
	if stressLevel == "full" {
		go DummyPro()
	}

	ReadFile(fileName)
	//go DummyPro()
	DeleteFile(fileName)
	TotalP--
}

func main() {
	var versionFlag = flag.Bool("v", false, "output version information and exit.")
	var helpFlag = flag.Bool("h", false, "display this help dialog")
	var stressLevel = flag.String("l", "hard", "level of stress on cpu")
	var useRandomBytes = flag.Bool("r", false, "use random bytes")
	var bufferSize = flag.Int("b", 0, "if 0 not use buffered writer")
	var arrayLength = flag.Int("a", 16, "byte array length")
	flag.Parse()

	if *helpFlag == true {
		printHelp()
		os.Exit(0)
	}

	if *versionFlag == true {
		printVersion()
		os.Exit(0)
	}

	CreateTempFolder()

	i := 0
	for {
		fileName := "./temp/log." + strconv.Itoa(i) + ".tmp"
		Run(*useRandomBytes, fileName, *arrayLength, *bufferSize, *stressLevel)
		i++
	}
}

//GetByteArray creates byte array, random or ordered
//l : length of array
//r : bool, is random or not
func GetByteArray(l int, r bool) []byte {
	rand.Seed(time.Now().UTC().UnixNano())
	t := make([]byte, l)
	for i := 0; i < l; i++ {
		if r {
			t[i] = byte(rand.Intn(255))
		} else {
			t[i] = 1
		}
	}
	return t
}

//CreateFile creates test file
//b : bytes to write file
//fn : name of the file
func CreateFile(b []byte, fn string) {
	err := ioutil.WriteFile(fn, b, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

//CreateFileBufio creates test file with buffer
//b : bytes to write file
//fn : name of the file
//bs : buffer size
func CreateFileBufio(b []byte, fn string, bs int) {
	f, err := os.Create(fn)
	if err != nil {
		fmt.Println(err)
	}
	w := bufio.NewWriter(f)
	bufio.NewWriterSize(w, len(b))
	defer w.Flush()
	i := 0
	for i < len(b) {
		e := i + bs
		if e > len(b) {
			e = len(b)
		}
		bt := b[i:e]

		_, err2 := w.Write(bt)
		if err2 != nil {
			fmt.Println(err2)
		}
		i += bs
	}

}

//ReadFile reads file data
//fn : name of the file
func ReadFile(fn string) {
	_, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Println(err)
	}
}

//DeleteFile deletes file from disk
//fn : name of the file
func DeleteFile(fn string) {
	err := os.Remove(fn)
	if err != nil {
		log.Println(err)
	}
}

//DummyPro creates dummy data on memory
func DummyPro() {
	t := 0
	for i := 0; i < rand.Intn(1024*1024*1024); i++ {
		t += rand.Intn(1024 * 1024 * 1024)
	}
}
