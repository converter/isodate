/*
isodate is a helper that prints ISO8601 format datetime for use at build time.

Example output: 2018-03-12T15:15:50Z
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UTC().Format(time.RFC3339))
}
