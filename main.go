package main

import (
	"github.com/minjaelee0727/gocoin/cli"
	"github.com/minjaelee0727/gocoin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
