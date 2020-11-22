package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// args := os.Args

	generateTestFiles()

	println("Found these files:")
	listFiles()

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		cleanFileName(f)
	}

	println("New list of files")
	listFiles()
}

func generateTestFiles() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)

	f, err := os.Create("/tmp/dat2")

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}

func check(err struct{}) {
	if err != nil {
		log.Fatal(err)
	}
}

func cleanFileName(file os.FileInfo) {
	var oldName = file.Name()
	println("Cleaning: " + oldName)
}

func listFiles() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
	}
}
