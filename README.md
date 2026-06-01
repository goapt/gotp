> Fork: [github.com/xlzd/gotp](https://github.com/xlzd/gotp)

# GOTP - The Golang One-Time Password Library

[![build-status][build-status]][build-status] ![MIT License][license-badge]

GOTP is a Golang package for generating and verifying one-time passwords. It can be used to implement two-factor (2FA) or multi-factor (MFA) authentication methods in anywhere that requires users to log in.

Open MFA standards are defined in [RFC 4226][RFC 4226] (HOTP: An HMAC-Based One-Time Password Algorithm) and in [RFC 6238][RFC 6238] (TOTP: Time-Based One-Time Password Algorithm). GOTP implements server-side support for both of these standards.

GOTP was inspired by [PyOTP][PyOTP].

## Installation

```
go get github.com/goapt/gotp
```

## Usage

Check API docs at https://godoc.org/github.com/goapt/gotp

### Time-based OTPs

```go
totp := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO")
otp, err := totp.Now()           // current otp, e.g. "123456"
otp, err := totp.At(1524486261)  // otp of timestamp 1524486261, e.g. "123456"

// OTP verified for a given timestamp
ok, err := totp.Verify("492039", 1524486261)  // true
ok, err := totp.Verify("492039", 1520000000)  // false

// OTP verified with time window tolerance (handles clock drift)
ok, err := totp.VerifyWithWindow("492039", 1524486261, 1)  // true, allows ±1 interval

// generate a provisioning uri
totp.ProvisioningUri("demoAccountName", "issuerName")
// otpauth://totp/issuerName:demoAccountName?secret=4S62BZNFXXSZLCRO&issuer=issuerName
```

### Counter-based OTPs

```go
hotp := gotp.NewDefaultHOTP("4S62BZNFXXSZLCRO")
otp, err := hotp.At(0)  // e.g. "944181"
otp, err := hotp.At(1)  // e.g. "770975"

// OTP verified for a given counter
ok, err := hotp.Verify("944181", 0)  // true
ok, err := hotp.Verify("944181", 1)  // false

// generate a provisioning uri
hotp.ProvisioningUri("demoAccountName", "issuerName", 1)
// otpauth://hotp/issuerName:demoAccountName?secret=4S62BZNFXXSZLCRO&counter=1&issuer=issuerName
```

### Generate random secret

```go
secret, err := gotp.RandomSecret(16)  // e.g. "LMT4URYNZKEWZRAA"
```

### Google Authenticator Compatible

GOTP works with the Google Authenticator iPhone and Android app, as well as other OTP apps like Authy.
GOTP includes the ability to generate provisioning URIs for use with the QR Code
scanner built into these MFA client apps via `ProvisioningUri` method:

```go
gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO").ProvisioningUri("demoAccountName", "issuerName")
// otpauth://totp/issuerName:demoAccountName?secret=4S62BZNFXXSZLCRO&issuer=issuerName

gotp.NewDefaultHOTP("4S62BZNFXXSZLCRO").ProvisioningUri("demoAccountName", "issuerName", 1)
// otpauth://hotp/issuerName:demoAccountName?secret=4S62BZNFXXSZLCRO&counter=1&issuer=issuerName
```

This URL can then be rendered as a QR Code which can then be scanned and added to the users list of OTP credentials.

### Working example

Scan the following barcode with your phone's OTP app (e.g. Google Authenticator):

![Demo](https://user-images.githubusercontent.com/5506906/39129827-0f12b582-473e-11e8-9c19-5e4f071eed26.png)

Now run the following and compare the output:

```go
package main

import (
	"fmt"
	"log"

	"github.com/goapt/gotp"
)

func main() {
	otp, err := gotp.NewDefaultTOTP("4S62BZNFXXSZLCRO").Now()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current OTP is", otp)
}
```

## Changes from upstream

Compared to the original [xlzd/gotp](https://github.com/xlzd/gotp), this fork includes the following improvements:

- **Error handling**: All methods that could fail now return `(result, error)` instead of using `panic`. This includes `At`, `Verify`, `Now`, `NowWithExpiration`, `RandomSecret`, and `BuildUri`.
- **Bug fix**: `byteSecret()` no longer modifies the `OTP.secret` field (padding accumulation bug).
- **Bug fix**: `BuildUri` now properly percent-encodes special characters in the label path (e.g. `@` in email addresses) and uses `%20` for spaces in query parameters instead of `+`.
- **New feature**: `TOTP.VerifyWithWindow` supports time window tolerance for clock drift handling.
- **API change**: `Itob` renamed to `itob` (unexported, internal-only).
- **Documentation**: Added proper Go doc comments and package-level documentation in `doc.go`.

## License

GOTP is licensed under the [MIT License][License]

[RFC 4226]: https://tools.ietf.org/html/rfc4226
[RFC 6238]: https://tools.ietf.org/html/rfc6238
[PyOTP]: https://github.com/pyotp/pyotp
[build-status]: https://github.com/goapt/gotp/actions/workflows/go.yml/badge.svg
[license-badge]: https://img.shields.io/badge/license-MIT-green.svg
[License]: https://github.com/goapt/gotp/blob/master/LICENSE
