package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	flag "github.com/spf13/pflag"

	. "github.com/logrusorgru/aurora"
)

// version flags for main function
var (
	appname     string
	version     string
	date        string
	description string = "Golang App to display time zone information."
)

// CheckError is a function to print out errors
func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

// clearScreen simple clears the terminal screen of any existing text
func clearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// main is the primary entry point for the app
func main() {

	//clearScreen()

	// set the option flagsd
	var ver = flag.BoolP("version", "v", false, "prints app version information")
	flag.Parse()

	// only print the version informationif the user asks for it.
	if *ver {
		fmt.Println("\nApp Name .....: ", Cyan(appname))
		fmt.Println("Version ......: ", Cyan(version))
		fmt.Println("Build Date ...: ", Cyan(date))
		fmt.Println("Description ..: ", Cyan(description))
		fmt.Println()
		os.Exit(0)
	}

	tNow := time.Now()

	fmt.Println(Cyan("\nCurrent Local Time Zone"))
	fmt.Println(strings.Repeat("-", 55))

	// Local Time Now in Unix Timestamp format
	tUnix := tNow.Unix()
	fmt.Println("Unix.Time:      ", tUnix)

	// Local Time Now from unix timestamp
	tLocal := time.Unix(tUnix, 0)
	fmt.Println("Time.Local:     ", tLocal)

	// UTC Time from unix timestamp
	tUtc := time.Unix(tUnix, 0).UTC()
	fmt.Println("Time.UTC:       ", tUtc)

	// Location Source: https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	fmt.Println(Cyan("\nLocal Time For A Specific Zone"))
	fmt.Println(strings.Repeat("-", 58))

	eastern, e := time.LoadLocation("US/Eastern")
	CheckError(e)

	central, e := time.LoadLocation("US/Central")
	CheckError(e)

	mountain, e := time.LoadLocation("US/Mountain")
	CheckError(e)

	pacific, e := time.LoadLocation("US/Pacific")
	CheckError(e)

	alaska, e := time.LoadLocation("US/Alaska")
	CheckError(e)

	hawaii, e := time.LoadLocation("US/Hawaii")
	CheckError(e)

	fmt.Println("US/Eastern:     ", tNow.In(eastern))
	fmt.Println("US/Central:     ", tNow.In(central))
	fmt.Println("US/Mountain:    ", tNow.In(mountain))
	fmt.Println("US/Pacific:     ", tNow.In(pacific))
	fmt.Println("US/Alaska:      ", tNow.In(alaska))
	fmt.Println("US/Hawaii:      ", tNow.In(hawaii))
	/*
	   Goloang Time Format Examples
	*/
	fmt.Println(Cyan("\nGolang Time Package Format Examples"))
	fmt.Println(strings.Repeat("-", 55))
	fmt.Println("ANSIC:          ", tNow.Format(time.ANSIC))
	fmt.Println("UnixDate:       ", tNow.Format(time.UnixDate))
	fmt.Println("RubyDate:       ", tNow.Format(time.RubyDate))
	fmt.Println("RFC822:         ", tNow.Format(time.RFC822))
	fmt.Println("RFC822Z:        ", tNow.Format(time.RFC822Z))
	fmt.Println("RFC850:         ", tNow.Format(time.RFC850))
	fmt.Println("RFC1123:        ", tNow.Format(time.RFC1123))
	fmt.Println("RFC1123Z:       ", tNow.Format(time.RFC1123Z))
	fmt.Println("RFC3339:        ", tNow.Format(time.RFC3339))
	fmt.Println("Kitchen:        ", tNow.Format(time.Kitchen))
	fmt.Println("StampMicro:     ", tNow.Format(time.StampMicro))
	fmt.Println("StampMilli:     ", tNow.Format(time.StampMilli))
	fmt.Println("StampNano:      ", tNow.Format(time.StampNano))
	fmt.Println()
}
