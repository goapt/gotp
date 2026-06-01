package main

import (
	"fmt"
	"log"

	"github.com/goapt/gotp"
)

func main() {
	secret, err := gotp.RandomSecret(16)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Random secret:", secret)
	defaultTOTPUsage()
	defaultHOTPUsage()
}

func defaultTOTPUsage() {
	otp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")

	now, err := otp.Now()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("current one-time password is:", now)

	at0, err := otp.At(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("one-time password of timestamp 0 is:", at0)

	uri, err := otp.ProvisioningUri("demoAccountName", "issuerName")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uri)

	ok, err := otp.Verify("179394", 1524485781)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("verify result:", ok)
}

func defaultHOTPUsage() {
	otp := gotp.NewDefaultHOTP("4S62BZNFXXSZLCRO")

	at0, err := otp.At(0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("one-time password of counter 0 is:", at0)

	uri, err := otp.ProvisioningUri("demoAccountName", "issuerName", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(uri)

	ok, err := otp.Verify("944181", 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("verify result:", ok)
}
