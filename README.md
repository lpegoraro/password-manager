# Password Manager

[![Build Status](https://travis-ci.com/lpegoraro/password-manager.svg?branch=master)](https://travis-ci.com/lpegoraro/password-manager) [![Join the chat at https://gitter.im/password-manager-go/community](https://badges.gitter.im/password-manager-go/community.svg)](https://gitter.im/password-manager-go/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

POC created to test and enhance Go skills

## Getting started

```bash
go build src/
```

## Usage

```bash
$ password-manager
Password Manager in Go version 0.0.1
Usage: `password_manager {COMMANDS} {OPTIONS}`
 The command list is the below
        help | -h: Prints this message
        version | -v: Print the version of the app
        get | -g {DESCRIPTION} {USERNAME} {OPTIONS}: Copy the password to the clipboard, for more information use `password_manager get help
        add | -a {DESCRIPTION} {USERNAME} {OPTIONS}: Add a new password entry, for more information use `password_manager add help
        config | -c {OPTIONS}: Configure encryption or password generation method
```

## TODO

- Create passwords in dictionaries with login and tags to know from where this is from
- Store in several configurable ways (configuration and multi-package implementation)
- Support most common types of encryption
- Provide safe web service in API
