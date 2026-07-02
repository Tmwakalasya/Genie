package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

const banner = `
  ____ _____ _   _ ___ _____ 
 / ___| ____| \ | |_ _| ____|
| |  _|  _| |  \| || ||  _|  
| |_| | |___| |\  || || |___ 
 \____|_____|_| \_|___|_____|
`

var bootLines = []string{
	"[    0.000000] Genie BIOS v0.1.0 — control plane bootstrap",
	"[    0.000412] CPU: 1 goroutine online (more conjured on demand)",
	"[    0.001093] Initializing VM store ............ OK",
	"[    0.002247] Loading lifecycle state machine .. OK",
	"[    0.002251]   pending -> booting -> running -> terminated",
	"[    0.002254]   booting -> failed  (loud error, or timeout for the silent ones)",
	"[    0.004120] Arming quota enforcer ............ OK",
	"[    0.004988] Rate limiters .................... NOT INSTALLED (arrives Phase 3)",
	"[    0.006301] Cloud provisioner ................ FAKE (gets real in Phase 5)",
	"[    0.007777] Rubbing lamp ..................... OK",
	"[    0.008128] Genie is awake. Your wish is my command.",
}

func printBoot() {
	fmt.Print(banner)
	for _, line := range bootLines {
		fmt.Println(line)
		time.Sleep(60 * time.Millisecond) // the theatrical pause
	}
	fmt.Println()
}

func main() {
	printBoot()

	router := gin.Default()
	router.Run("localhost:8080")
}
