package config

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	keyStringFlag  = "ks"
	keyStringUsage = "key string"
	keyFileFlag    = "kf"
	keyFileUsage   = "key file"
)

type Config struct {
	Key []byte
	In  io.Reader
}

func usage(msg string) {
	flag.Usage()
	_, err := fmt.Fprintf(os.Stderr, "\n%s\n", msg)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(1)
}

func cliHasArg(arg string) bool {
	var result bool
	flag.Visit(func(f *flag.Flag) {
		if f.Name == arg {
			result = true
		}
	})
	return result
}

func Get() (Config, error) {
	keyString := flag.String(keyStringFlag, "", keyStringUsage)
	keyFile := flag.String(keyFileFlag, "", keyFileUsage)
	flag.Parse()

	if cliHasArg(keyStringFlag) && cliHasArg(keyFileFlag) {
		usage(fmt.Sprintf("-%s and -%s arguments are mutually exclusive", keyStringFlag, keyFileFlag))
	}

	if !cliHasArg(keyStringFlag) && !cliHasArg(keyFileFlag) {
		usage(fmt.Sprintf("one of '-%s <keystring>' or '-%s <keyfile>' must be specified", keyStringFlag, keyFileFlag))
	}

	// Get key directly from CLI argument
	var key []byte
	if cliHasArg(keyStringFlag) {
		key = []byte(*keyString)
	}

	// Get key from file specified in CLI argument
	var err error
	if cliHasArg(keyFileFlag) {
		key, err = ioutil.ReadFile(*keyFile)
		if err != nil {
			return Config{}, err
		}
	}

	if len(key) == 0 {
		return Config{}, errors.New("key cannot be empty")
	}

	cfg := Config {
		Key: key,
	}

	args := flag.Args()
	switch len(args) {
	case 0:
		cfg.In = os.Stdin
	case 1:
		in, err := os.Open(args[0])
		if err != nil {
			return Config{}, err
		}
		cfg.In = in
	default:
		return Config{}, errors.New("only one input file allowed")
	}

	return cfg, nil
}
