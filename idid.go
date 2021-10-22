package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"time"
)

const Version = 1.3

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func version() {
	fmt.Println("+------ Version: ", Version, " ------+")
	os.Exit(0)
}

func read(path string) {
	dat, err := ioutil.ReadFile(path)
	check(err)
	fmt.Println("+------------------------------------------------+")
	fmt.Println("|                 Read the file:                 |")
	fmt.Println("+------------------------------------------------+\n")
	fmt.Print(string(dat))
	os.Exit(0)
}

func write(path string, msg string) {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	defer f.Close()

	currentTime := time.Now()
	scanner := bufio.NewScanner(f)
	dateAlreadyPresent := false
	dateToWrite := "============ [ " + currentTime.Format("02/01/2006") + " ] ============"
	for scanner.Scan() {
		if scanner.Text() == dateToWrite {
			dateAlreadyPresent = true
		}
	}

	if !dateAlreadyPresent {
		_, err = f.WriteString(dateToWrite + "\n")
		check(err)
	}
	_, err = f.WriteString("* ")
	_, err = f.WriteString(msg)
	_, err = f.WriteString("\n")
	check(err)

	f.Sync()
}

func main() {

	usr, err := user.Current()
	check(err)
	home := usr.HomeDir + "/idid"

	var (
		pathPtr    = flag.String("path", home, "Path for file")
		messagePtr = flag.String("msg", "", "Message to write")
		showPtr    = flag.Bool("show", false, "Use to show the file")
		verPtr     = flag.Bool("version", false, "Show the version")
	)
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Println("No arguments\n")
		fmt.Println("Try lunching with -h to show how to use")
		os.Exit(3)
	}

	if *verPtr {
		version()
	}

	if *showPtr {
		read(*pathPtr)
	}

	write(*pathPtr, *messagePtr)
}
