package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jnnkrdb/gokv/global"
)

func main() {

	// activate logging
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)

	fmt.Printf("Hello world!")

	for {

		log.Printf("New Line, HomeDir: %s, BinDir: %s", global.GOKV_HOME, global.GOKV_BINARY_PATH)

		time.Sleep(30 * time.Second)
	}
}
