[![Go](https://github.com/davrux/go-smtptester/actions/workflows/go.yml/badge.svg)](https://github.com/davrux/go-smtptester/actions/workflows/go.yml) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

# go-smtptester

Simple SMTP Server for Testing.

## How it works

All received mails are saved in a sync.Map with a key:

~~~~go
From+Recipient1+Recipient2
~~~~

Mails to the same sender and recipients will overwrite a previous
received mail, when the recipients slice has the same order as
in the mail received before.

## Example

See

~~~~sh
server_test.go
examples/simple/main.go
~~~~

for exmaple usage.
