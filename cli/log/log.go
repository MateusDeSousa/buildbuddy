package log

import (
	"flag"
	"log"
)

var (
	verbose = flag.Bool("bb_verbose", false, "If true, enable verbose buildbuddy logging.")
)

func Debugf(format string, v ...interface{}) {
	if !*verbose {
		return
	}
	log.Printf(format, v...)
}

func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}

func Fatal(format string) {
	log.Fatal(format)
}