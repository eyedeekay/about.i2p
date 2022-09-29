package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"cerca/util"

	"github.com/eyedeekay/about.i2p/about"
	"github.com/eyedeekay/onramp"
)

func readAllowlist(location string) []string {
	ed := util.Describe("read allowlist")
	data, err := os.ReadFile(location)
	ed.Check(err, "read file")
	list := strings.Split(strings.TrimSpace(string(data)), "\n")
	var processed []string
	for _, fullpath := range list {
		u, err := url.Parse(fullpath)
		if err != nil {
			continue
		}
		processed = append(processed, u.Host)
	}
	return processed
}

func complain(msg string) {
	fmt.Printf("cerca: %s\n", msg)
	os.Exit(0)
}

func main() {
	var allowlistLocation string
	var sessionKey string
	var dir string
	flag.StringVar(&allowlistLocation, "allowlist", "", "domains which can be used to read verification codes from during registration")
	flag.StringVar(&sessionKey, "authkey", "", "session cookies authentication key")
	flag.StringVar(&dir, "dir", "", "directory to run in")
	flag.Parse()
	if len(sessionKey) == 0 {
		complain("please pass a random session auth key with --authkey")
	} else if len(allowlistLocation) == 0 {
		//complain("please pass a file containing the verification code domain allowlist")
		allowlistLocation = "allow.txt"
		if err := ioutil.WriteFile(allowlistLocation, []byte(""), 0644); err != nil {
			panic(err)
		}
	}
	if _, err := os.Stat(allowlistLocation); os.IsNotExist(err) {
		if err := ioutil.WriteFile(allowlistLocation, []byte(""), 0644); err != nil {
			panic(err)
		}
	}

	garlic, err := onramp.NewGarlic("about.i2p", "127.0.0.1:7656", []string{})
	if err != nil {
		panic(err)
	}
	http.DefaultClient.Transport = &http.Transport{
		Dial: garlic.Dial,
	}
	allowList := readAllowlist(allowlistLocation)
	allowList = append(allowList, "*.i2p")
	allowList = append(allowList, "*.b32.i2p")
	if ln, err := garlic.ListenTLS(); err != nil {
		panic(err)
	} else {
		allowList = append(allowList, ln.Addr().String())
		if cercaServer, err := about.NewServer(allowList, sessionKey, dir); err != nil {
			panic(err)
		} else {
			if err := http.Serve(ln, cercaServer); err != nil {
				panic(err)
			} else {
				log.Println("Exited gracefully")
			}

		}
	}
}
