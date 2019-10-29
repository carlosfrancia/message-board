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

type aFile struct {
	filename string
	ext      string
}

type point struct {
	x, y int
}

// p := point{1, 2}

type myArrayVector []struct {
	key, value string
}

// Another way same thing

func createArray() {

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	println(s)
	myArray := []struct {
		key   string
		value string
	}{
		{"mykey", "myValue"},
		{"mykey2", "myValue2"},
	}
	println(myArray)

	// create an array of length 10
	pow := make([]int, 10)
	//iterate

	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	// Slice of slices

	pic := make([][]int, 10) /* type declaration. coordenate y */
	for i := range pic {
		pic[i] = make([]int, 5) /* again the type? Go needs it. coordenate x */
		for j := range pic[i] {
			pic[i][j] = int((i ^ j))
		}
	}
}

func (f aFile) createFile() {
	myNewFile, err := os.Create(f.filename + "." + f.ext)
	if err != nil {
		fmt.Printf("failed writing to file '%s'", myNewFile.Name())
	}
	defer myNewFile.Close()
}

func (f aFile) writeInFile(text string) {
	log.Info("filename: " + f.filename + "." + f.ext)
	writtingFile, err := os.OpenFile(f.filename+"."+f.ext, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer writtingFile.Close()
	fmt.Fprintln(writtingFile, text)
}

func (f aFile) readFile() {
	tmpfile, err := os.Open(fmt.Sprintf("%s.%s", f.filename, f.ext))

	log.Info(fmt.Sprintf("Printing file: %s.%s", f.filename, f.ext))
	//	tmpfile, err := os.Open(f.filename + "." + f.ext)
	if err != nil {
		log.Fatal(err)
	}
	defer tmpfile.Close()
	b, err := ioutil.ReadAll(tmpfile)
	fmt.Print(string(b))
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

func main2() {
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
	log.Info("My new file is: ", myNewFile.Name())
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

	aNewFile := aFile{
		filename: "files/signatureLoveFile",
		ext:      "txt",
	}
	aNewFile.createFile()
	aNewFile.writeInFile("My new content")
	aNewFile.writeInFile("Sandra is a girl")
	aNewFile.writeInFile("Carlitos is her boyfriend")
	aNewFile.readFile()

}

func newAppConfig() (*appConfig, error) {

	return &appConfig{
		att1:   att1Input,
		att2:   att2Input,
		attMap: myMap,
	}, nil
}
