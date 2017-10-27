package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bernardigiri/cfgcrypt/textValueEncryptor"
)

func explainUsage(cli *flag.FlagSet) {
	fmt.Println("cfgcrypt [textfile] ...")
	fmt.Println("\ttextfile\tText file to encrypt. (required)")
	cli.PrintDefaults()
}

func main() {
	cli := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	encodedKey := cli.String("key", "", "Base64 encoded encryption key, if not specified one will be generated")
	prefix := cli.String("prefix", "#{{", "Prefix string denoting start of value to be encrypted")
	postfix := cli.String("postfix", "}}#", "Post string denoting end of value to be encrypted")
	if len(os.Args) < 2 {
		os.Stderr.WriteString("Not enough arguments\n")
		explainUsage(cli)
		os.Exit(1)
	}
	textfile := os.Args[1]
	cli.Parse(os.Args[2:])
	if *prefix == "" ||
		*postfix == "" ||
		*prefix == *postfix {
		os.Stderr.WriteString("Invalid prefix/postfix\n")
		explainUsage(cli)
		os.Exit(1)
	}
	err := textValueEncryptor.EncryptTextFile(textfile, *prefix, *postfix, *encodedKey)
	if err != nil {
		msg := fmt.Sprintf("Error encrypting file \"%s\":\n%s\n", textfile, err.Error())
		os.Stderr.WriteString(msg)
		os.Exit(1)
	}
}
