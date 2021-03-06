[![GoDoc](https://godoc.org/github.com/Apfel/openhmd-go?status.svg)](https://godoc.org/github.com/Apfel/openhmd-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/Apfel/openhmd-go)](https://goreportcard.com/report/github.com/Apfel/openhmd-go)
[![Build Status](https://travis-ci.org/Apfel/openhmd-go.svg?branch=master)](https://travis-ci.org/Apfel/openhmd-go)

# openhmd-go
[OpenHMD](http://www.openhmd.net/) API bindings for [Go](https://golang.org/).

```sh
go get github.com/Apfel/openhmd-go
```

This module requires OpenHMD. Click [here](http://www.openhmd.net/index.php/download/) for instructions.

## Example
This is OpenHMD's simple example, ported to Go using openhmd-go.

```go
package main

import (
	"log"
	"os"
	"strconv"
	"time"

	openhmd "github.com/Apfel/openhmd-go"
)

func main() {
	log.Printf("openhmd-go - Simple Example")
	var id int

	if len(os.Args) < 2 || os.Args[1] == "" {
		log.Fatalln("Please provide an device ID.")
	} else {
		id, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("Couldn't convert '%s' to an integer.\nError: %s\n", os.Args[1], err.Error())
		}
		log.Printf("Using ID %d.", id)
	}

	context := openhmd.CreateContext()
	if context == nil {
		log.Fatalln("Context couldn't be opened.")
	}

	if count := context.Probe(); count == 0 {
		log.Fatalln("No devices, quitting...")
	} else {
		log.Printf("Found device(s). Device count: %d\n", count)
	}

	device := context.ListOpenDevice(id)
	if device == nil || len(context.GetError()) != 0 {
		log.Fatalf("Device with ID %d couldn't be opened. Error: %s\n", id, context.GetError())
	} else {
		log.Printf("Opened device %s, vendor is %s. ID: %s\n", context.ListGetString(id, openhmd.StringValueProduct),
			context.ListGetString(id, openhmd.StringValueVendor), context.ListGetString(id, openhmd.StringValuePath))
	}

	c, width := device.GetInt(openhmd.IntValueScreenVerticalResolution, 1)
	if c != openhmd.StatusCodeOkay {
		log.Fatalf("Fetching Device width error.\nStatus Code: %d\n\n", c)
	}

	c, height := device.GetInt(openhmd.IntValueScreenHorizontalResolution, 1)
	if c != openhmd.StatusCodeOkay {
		log.Fatalf("Fetching Device height error.\nStatus Code: %d\n", c)
	}

	log.Printf("Device properties:\nResolution: %dx%d\n", width[0], height[0]) // I do know that this is rather poorly designed, but whatever

	for 1 == 1 {
		c, rot := device.GetFloat(openhmd.FloatValueRotationQuat, 4)
		if c != openhmd.StatusCodeOkay {
			log.Fatalf("Rotation - Error code %d\n", c)
		}
		c, pos := device.GetFloat(openhmd.FloatValuePositionVector, 3)
		if c != openhmd.StatusCodeOkay {
			log.Fatalf("Position - Error code %d\n", c)
		}
		log.Printf("Rotation: %f %f %f %f\nPosition: %f %f %f", rot[0], rot[1], rot[2], rot[3], pos[0], pos[1], pos[2])
		time.Sleep(time.Millisecond * 100)
		context.Update()
	}
}
```
