package main

import (
	"fmt"
	. "github.com/iotaledger/iota.go/api"
	"github.com/iotaledger/iota.go/bundle"
	"github.com/iotaledger/iota.go/converter"
	"github.com/iotaledger/iota.go/trinary"
)

var node = "https://nodes.devnet.iota.org"
//var node = "http://0.0.0.0:14265"

const depth = 3
const minimumWeightMagnitude = 9

const address = trinary.Trytes("ZLGVEQ9JUZZWCZXLWVNTHBDX9G9KZTJP9VEERIIFHY9SIQKYBVAHIMLHXPQVE9IXFDDXNHQINXJDRPFDXNYVAPLZAW")

const seed = trinary.Trytes("JBN9ZRCOH9YRUGSWIQNZWAIFEZUBDUGTFPVRKXWPAUCEQQFS9NHPQLXCKZKRHVCCUZNF9CZZWKXRZVCWQ")

func main() {

	api, err := ComposeAPI(HTTPClientSettings{URI: node})
	must(err)

	var data = "{'message' : 'attack 2020.11.24 13:50'}"
	message, err := converter.ASCIIToTrytes(data)
	must(err)

	transfers := bundle.Transfers{
		{
			Address: address,
			Value: 0,
			Message: message,
		},
	}

	trytes, err := api.PrepareTransfers(seed, transfers, PrepareTransfersOptions{})
	must(err)

	myBundle, err := api.SendTrytes(trytes, depth, minimumWeightMagnitude)
	must(err)

	fmt.Println(bundle.TailTransactionHash(myBundle))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}