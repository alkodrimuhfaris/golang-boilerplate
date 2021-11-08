package main

import (
	"fmt"
	"os"
	"papitupi-web/src/infrastructure"
	_ "papitupi-web/src/infrastructure/timezone"
	"time"
)

func main() {
	// setting timezone
	os.Setenv("TZ", "Asia/Jakarta")

	fmt.Println(time.Now())

	infrastructure.Infrastructure()
}
