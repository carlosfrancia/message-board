package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
)

var flags = flag.NewFlagSet("My-app", flag.ExitOnError)
var att1Input string
var att2Input string
var logLevel string
var myMap map[string]bool

type appConfig struct {
	att1   string
	att2   string
	attMap map[string]bool
}

func init() {
	flags.StringVar(
		&logLevel, "log-level", "info", fmt.Sprintf("Log level, valid "+
			"values are %+v", log.AllLevels),
	)
	flags.StringVar(
		&att1Input, "att1", "att1Value", "The Att1",
	)
	flags.StringVar(
		&att2Input, "att2", "att2Value", "The Att2",
	)

	flags.Usage = func() {
		fmt.Printf("Usage of my-app:\n")
		flags.PrintDefaults()
		fmt.Println("\n my-app [-log-level (debug|info|error)")
		fmt.Println()
	}
	err := flags.Parse(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Starting
	ll, err := log.ParseLevel(logLevel)
	if err != nil {
		log.Fatalln(err)
	}
	log.SetLevel(ll)

	log.Info("Starting myApp")

	myMap = make(map[string]bool)
	myMap["myKey"] = true
	config, err := newAppConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Info("Hello world")
	log.WithFields(log.Fields{
		"Att1": config.att1,
		"Att2": config.att2,
	}).Info("My attributes")

	// Read a file
	// Create signature methods

	// Create a file
	myNewFile, err := os.Create("files/my_new_file.txt")
	if err != nil {
		fmt.Printf("failed writing to file '%s'", myNewFile.Name())
	}
	log.Info("My new file is $s: ", myNewFile.Name())
	fmt.Fprint(myNewFile, fmt.Sprintf("Starting a new file.\nAtt1 is %s \n", att1Input))

	defer myNewFile.Close()
	// option 1 as type file.

	tmpfile, err := os.Open("files/my_new_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer tmpfile.Close()

	// Readall returns []byte as below
	b, err := ioutil.ReadAll(tmpfile)
	fmt.Print(string(b))

	//Option 2 as type []byte

	contents, err := ioutil.ReadFile("files/my_file.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(contents))
}

func newAppConfig() (*appConfig, error) {

	return &appConfig{
		att1:   att1Input,
		att2:   att2Input,
		attMap: myMap,
	}, nil
}
