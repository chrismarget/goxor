package main

import (
	"github.com/chrismarget/goxor/config"
	"github.com/chrismarget/goxor/decrypt"
	"log"
	"os"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		log.Fatal(err)
	}

	err = decrypt.Decrypt(cfg.Key, cfg.In, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
