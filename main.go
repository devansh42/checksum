package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"flag"
	"fmt"
	"hash"
	"io"
	"log"
	"os"
)

func main() {
	var srcFile, hashMethod *string

	srcFile = flag.String("source", "", "Source File to check for checksum")
	hashMethod = flag.String("f", "md5", "Hash function to get checksum, available options are (md5,sha,sha1,sha2,sha3)")
	flag.Parse()
	for !flag.Parsed() {
	} // Waiting for arguments to password
	file, err := os.Open(*srcFile)
	if err != nil {
		log.Fatal("Couldn't open file due to : ", err.Error())
	}
	defer file.Close() // Closing the file after reading

	fmt.Print(getChecksum(file, *hashMethod))
}

func getChecksum(reader io.Reader, hashMethod string) string {
	br := bufio.NewReader(reader)

	var hasher hash.Hash
	switch hashMethod {

	case "md5":
		hasher = md5.New()
	case "sha", "sha1":
		hasher = sha1.New()
	case "sha2":
		hasher = sha256.New()
	case "sha3":
		hasher = sha512.New()
	default:
		log.Fatal("Unsupported hash functions : " + hashMethod)
	}
	_, err := io.Copy(hasher, br)
	if err != nil {
		log.Fatal("Error while calculating hash : " + err.Error())
	}
	return hex.EncodeToString(hasher.Sum(nil))
}
