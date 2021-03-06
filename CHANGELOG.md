# BMX

## 2.0.1-beta

### BUG FIXES

* Fixes a problem with password entry on darwin builds

### Improvements

* Update README documentation to be more correct with the `go get` paths

## 2.0.0-beta

### Improvements

* Rewrite in Golang

## 1.1.0

### Features

* BMX's write command now accepts an '--output' option (#128).

### BUG FIXES

* BMX no longer creates an empty 'credentials' directory in the CWD (#124).
* BMX no longer fails when '~/.aws/credentials' does not exist (#131).

## 1.0.0

### Features

* BMX is multi-platform: it runs on Linux and Windows.
* BMX maintains your Okta session for 12 hours: you enter your Okta password once a day, and BMX takes care of the rest.
* BMX supports TOTP and SMS MFA.
* BMX manages its own AWS STS tokens and never modifies ~/.aws/credentials without explicit direction from you. (See bmx write.)
* BMX has commands for renewing and forgetting its AWS STS tokens.
* BMX manages its AWS-CLI dependency.
