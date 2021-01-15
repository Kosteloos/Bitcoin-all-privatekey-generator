package main

import (
	"fmt"
	"math/big"
	"net/http"
    "io/ioutil"
    "os"

	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
)

func getbalance(address string) string {
  response, err := http.Get("https://blockchain.info/q/getreceivedbyaddress/" + address)
  if err != nil {
      fmt.Printf("%s", err)
      os.Exit(1)
  }
  defer response.Body.Close()
  contents, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Printf("%s", err)
    os.Exit(1)
  }
  return string(contents)
}

func main() {
	// Print header
	fmt.Printf("%64s %34s %34s\n", "Private", "Public", "Public Compressed")

	// Initialise big numbers with small numbers
	count, one := big.NewInt(0), big.NewInt(1)
	count.SetString("102987336249554097029535217322581322789799900648198034993379397001115665086549",10)

	// Create a slice to pad our count to 32 bytes
	padded := make([]byte, 32)

	// Loop forever because we're never going to hit the end anyway
	for {
		// Increment our counter
		count.Add(count, one)

		// Copy count value's bytes to padded slice
		copy(padded[32-len(count.Bytes()):], count.Bytes())

		// Get public key
		_, public := btcec.PrivKeyFromBytes(btcec.S256(), padded)

		// Get compressed and uncompressed addresses
		caddr, _ := btcutil.NewAddressPubKey(public.SerializeCompressed(), &chaincfg.MainNetParams)
		uaddr, _ := btcutil.NewAddressPubKey(public.SerializeUncompressed(), &chaincfg.MainNetParams)

		// Print keys
		balanceu := getbalance(uaddr.EncodeAddress())
		balancec := getbalance(caddr.EncodeAddress())
		if balanceu > "0" {
		fmt.Printf("%x %34s %1s %34s %1s\n", padded, uaddr.EncodeAddress(), balanceu, caddr.EncodeAddress(), balancec)
		}
		if balancec > "0" {
		fmt.Printf("%x %34s %1s %34s %1s\n", padded, uaddr.EncodeAddress(), balanceu, caddr.EncodeAddress(), balancec)
		}
	}
}
