package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
)

type logger interface {
	log()
}

type logData struct {
	message  string
	fileName string
}

func (lData *logData) log() {
	err := ioutil.WriteFile(lData.fileName, []byte(lData.message), fs.FileMode(0644))

	if err != nil {
		fmt.Println("Storing the log data failed!")
	}
}

type loggableString string

func (text loggableString) log() {
	fmt.Println(text)
}

func main() {
	userLog := &logData{"User logged in", "user-log.txt"}
	// do more work ...
	createLog(userLog)

	message := loggableString("[INFO] User created!")
	// do more work ...
	createLog(message)
	// createLog("test")

	outputValue("test")

	fmt.Println(double(5))
	fmt.Println(double(5.99))
	fmt.Println(double([]int{1, 2, 3}))
}

func createLog(data logger) {
	// do more stuff
	data.log()
}

func outputValue(value interface{}) {
	val, ok := value.(string)
	// val, ok := value.(loggableString)
	fmt.Println(val, ok)
}

func double(value interface{}) interface{} {
	switch val := value.(type) {
	case int:
		return val * 2
	case float64:
		return val * 2
	case []int: 
		return append(val, val...)
	default:
		return ""
	}
}
