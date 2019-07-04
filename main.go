package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: gobfuscate [flags] out_path")
		flag.PrintDefaults()
		os.Exit(1)
	}

	//pkgName := flag.Args()[0]
	outPath := flag.Args()[0]

	if !obfuscate(outPath) {
		os.Exit(1)
	}
	log.Println("ok")
}

func obfuscate(outPath string) bool {

	log.Println("Obfuscating files...")
	if err := ObfuscateStrings(outPath); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to obfuscate strings:", err)
		return false
	}

	return  true
}

func encryptComponents(pkgName string, enc *Encrypter) string {
	comps := strings.Split(pkgName, "/")
	for i, comp := range comps {
		comps[i] = enc.Encrypt(comp)
	}
	return strings.Join(comps, "/")
}
