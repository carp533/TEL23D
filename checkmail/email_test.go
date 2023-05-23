package main

import "fmt"

// prüft Syntaktische Korrektheit einer E-Mail
func ExampleCheckEmail() {
	mails := []string{
		"ç$€§/az@gmail.com",
		"abcd@gmail_yahoo.com",
		"abcd@gmail-yahoo.com",
		"abcd@gmailyahoo",
		"abcd@gmail.yahoo"}
	for _, mail := range mails {
		fmt.Println(CheckEmail(mail))
	}
	//Output:
	// true
	// false
	// true
	// false
	// true
}
