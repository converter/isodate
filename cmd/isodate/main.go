/*
isodate is a helper that prints ISO8601 format datetime for use at build time.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().UTC().Format(time.RFC3339))
}
