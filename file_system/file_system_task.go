package main

import (
    "os"
    "io/ioutil"
    "strconv"
)

func main() {
    binaryData, err := ioutil.ReadFile("counter.txt")
    if err != nil {
	file, _ := os.Create("counter.txt")
	os.Chmod("counter.txt", 0777)
	file.Write([]byte {'1'})
	file.Close()
    } else {
	data := string(binaryData)
	num, err := strconv.Atoi(data)
	if err != nil {
	    os.Exit(1)
	} else {
	    num++
	    file, _ := os.OpenFile("counter.txt", os.O_RDWR, 0777)
	    file.WriteAt([]byte(strconv.Itoa(num)), 0)
	    file.Close()
	}
    }
}