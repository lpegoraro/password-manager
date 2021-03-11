# Password Manager

[![Build Status](https://travis-ci.com/lpegoraro/password-manager.svg?branch=master)](https://travis-ci.com/lpegoraro/password-manager)
[![Join the chat at https://gitter.im/password-manager-go/community](https://badges.gitter.im/password-manager-go/community.svg)](https://gitter.im/password-manager-go/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![codecov](https://codecov.io/gh/lpegoraro/password-manager/branch/master/graph/badge.svg)](https://codecov.io/gh/lpegoraro/password-manager)

POC created to test and enhance Go skills

## Getting started

For Development

```bash
go build -a -o $GOPATH/bin/password-manager github.com/lpegoraro/password-manager/password-manager
```

For Usage

```bash
go get github.com/lpegoraro/password-manager/password-manager
```

Using Docker

```bash
cd password-manager
sudo docker run --rm -v "$PWD":/go/src/github.com/lpegoraro/password-manager -w /go/src/github.com/lpegoraro/password-manager golang:latest  go build -a -o $GOPATH/bin/password-manager ./password-manager/ && sudo docker build -t lpegoraro/password-manager:latest .
```

## Usage

```bash
$ password-manager
Password Manager in Go version 0.1.0
Usage: `password_manager {COMMANDS} {OPTIONS}`
 The command list is the below
	help | -h: Prints this message
	version | -v: Print the version of the app
	get | -g {DESCRIPTION} {USERNAME} {OPTIONS}: Copy the password to the clipboard, for more information use `password_manager get help
	add | -a {DESCRIPTION} {USERNAME} {OPTIONS}: Add a new password entry, for more information use `password_manager add help
	config | -c {METHOD} {SEED} {FACTOR} {STORAGE_TYPE}: Configure encryption or password generation method
	 | 	 "Method": Type of password, please choose from the following {uuid | cert | custom }
	 | 	 "Seed": Any passfrase you would like
	 | 	 "Factor": Given the Method uuid, you can choose between 4 and 5
	 | 	 	   Given the Method cert you can choose the algorithym for the password creation
	 | 	 "Storage Type": Only supporting "NOT_ENCRYPTED_FILE" storage at the moment, you can choose
	 | 	 	   You can choose output also, but you will need to manually configure in the settings since this
	 | 	 	is a development feature only.
```

## Serving

These are the diagrams exemplifying the flow to connect:

Preparation

```sequence
Tenant->PwdMngr: Upload PEM Key
PwdMngr->Tenant: Fingerprint
Tenant->PwdMngr: Request for Serve with Fingerprint
PwdMngr->Tenant: Port to connect
```

Serving

```flow
st=>start: Tenant sends Request for Serve with given Fingerprint
publickey=>operation: From Tenant Name retrieve key
storage=>operation: Create configuration and storage, adding key for encryption

st->publickey->storage
```

Add/Get Password Flow

```flow
st=>start: Tenant requests Add Password
conf=>operation: Retrieve configuration
command=>operation: Send command to PwdManager
bundle=>operation: Assemble response and encrypt from given public Key
respond=>operation: Sends encrypted Response


st->conf->command->bundle->respond
```

## TODO

- Store in several configurable ways (configuration and multi-package
  implementation)
- Support most common types of encryption and more password generation methods
- Provide safe web service in API
