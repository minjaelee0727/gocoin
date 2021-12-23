package main

import (
	"github.com/minjaelee0727/gocoin/cli"
	"github.com/minjaelee0727/gocoin/db"
)

func usage() {

}

func main() {
	defer db.Close()
	db.InitDB()
	cli.Start()
}
