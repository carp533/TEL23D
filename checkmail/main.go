package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println(VerifyEmailAddress("Mary.Smith@wikipedia.org"))
	// fmt.Println(VerifyEmailAddress("abcd@gmail.yahoo@"))
	// fmt.Println(VerifyEmailAddress("abcd@gmail.yahoo"))
	AskUser()
}

func AskUser() {
	// TODO verwende bufio.NewScanner(os.Stdin)
	scanner := bufio.NewScanner(os.Stdin)
	// und scanner.Scan() / scanner.Text()
	fmt.Println("Bitte Email eingeben")
	for scanner.Scan() {
		email := scanner.Text()
		fmt.Println("Eingabe " + email)
		fmt.Println(VerifyEmailAddress(email))
	}

}

func VerifyEmailAddress(email string) *EmailVerInfo {
	syntaxok := CheckEmail(email)
	if !syntaxok {
		return &EmailVerInfo{Email: email, SyntaxOK: syntaxok, Valid: false}
	}
	index := strings.LastIndex(email, "@")
	username := email[:index]
	domain := strings.ToLower(email[index+1:])
	dominf, err := CheckDomain(domain)
	return &EmailVerInfo{Email: email, Username: username, Domain: domain, SyntaxOK: syntaxok, Valid: err == nil, DomainInfo: dominf, Error: err}
}
