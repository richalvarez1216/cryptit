package main

import (
	"fmt"

	"github.com/richalvarez1216/cryptit/decrypt"
	"github.com/richalvarez1216/cryptit/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("Kodekloud")
	fmt.Println(encrypt.Nimbus("Kodekloud"))
	fmt.Println(decrypt.Nimbus(encryptedStr))

}
