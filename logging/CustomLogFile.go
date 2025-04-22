package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
)

func LogFile() {

	formatCustomLogFile()
	writeToMultipleFile()
}

func writeToMultipleFile() {
	flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile("myLog.log", flag, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	w := io.MultiWriter(file, os.Stderr)

	LtdFlags := log.Ldate | log.Ltime | log.Lshortfile
	logger := log.New(w, "myApp: ", LtdFlags)
	logger.Printf("This log will be appended to the file.")
	logger.Printf("Hello from custom log file ")
	logger.Printf("This log will be appended to the file.")
}

func formatCustomLogFile() {
	LOGFILE := path.Join(os.TempDir(), "mGo.log") // file will create on OS temp directory
	fmt.Println(LOGFILE)
	f, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return

	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	LtdFlags := log.Ldate | log.Lshortfile
	iLog := log.New(f, "LNum ", LtdFlags)
	iLog.Println("Mastering Go, 4th edition!")
	iLog.SetFlags(log.Lshortfile | log.LstdFlags)
	iLog.Println("Another log entry!")
}
