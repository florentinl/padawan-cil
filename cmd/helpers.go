package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"syscall"

	"golang.org/x/term"
)

func apiUri() string {
	if os.Getenv("PADAWAN_API_URI") != "" {
		return "https://" + os.Getenv("PADAWAN_API_URI")
	}
	return "https://padawan.kube.test.viarezo.fr"
}

func apiFQDN() string {
	if os.Getenv("PADAWAN_API_URI") != "" {
		return os.Getenv("PADAWAN_API_URI")
	}
	return "padawan.kube.test.viarezo.fr"
}

func setToken(token string) {
	err := ioutil.WriteFile("/tmp/padawan_token", []byte(token), 0600)
	if err != nil {
		fmt.Println("Error while saving token")
		os.Exit(1)
	}
}

func getToken() string {
	token, err := ioutil.ReadFile("/tmp/padawan_token")
	if err != nil {
		displayLoginMessage()
	}
	return string(token)
}

func displayLoginMessage() {
	fmt.Println("Error while loading token: please login with:")
	fmt.Println("$ padawan login")
	os.Exit(1)
}

func getPassword() string {
	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println("Error while reading password")
		os.Exit(1)
	}

	password := string(bytePassword)
	return password
}
