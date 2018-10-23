package main

import "github.com/vjosso/transfer.sh/cmd"

func main() {
	app := cmd.New()
	app.RunAndExitOnError()
}
