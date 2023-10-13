package main

import (
	"github.com/mileapp_screening/cmd"
	_ "github.com/mileapp_screening/utils/env"
	_ "github.com/mileapp_screening/utils/log"
)

func main() {
	cmd.Execute()
}
