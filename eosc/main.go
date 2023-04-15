package main

import (
	// Load all contracts here, so we can always read and decode
	// transactions with those contracts.
	_ "github.com/sebastianmontero/eos-go/msig"
	_ "github.com/sebastianmontero/eos-go/system"
	_ "github.com/sebastianmontero/eos-go/token"

	"github.com/sebastianmontero/eosc/eosc/cmd"
)

var version = "dev"

func init() {
	cmd.Version = version
}

func main() {
	cmd.Execute()
}
