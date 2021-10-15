package main

import (
	"github.com/minjaelee0727/gocoin/explorer"
	"github.com/minjaelee0727/gocoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
