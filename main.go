package main

import (
	"github.com/abhimanyu003/probe/cmd"
)

// version specify version of application using ldflags
var version = "dev"

func main() {
	cmd.Version = version
	cmd.Execute()
}
