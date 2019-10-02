# logman

[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fwymillerlinux%2Flogman.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fwymillerlinux%2Flogman?ref=badge_shield)

Collecting log files for the greater good!

This program lets you grab log files remotely and store them up into a compressed file on your local machine!

In all reality, you can literally grab any file you wish to grab of any remote machine.

## Installation

If you have Go installed, use `go get` for grabbing this repo:

`go get github.com/wymillerlinux/logman`

If you don't have Go installed, I will provide binaries soon.

## Running

Before you start the program, I would like to highlight the attraction to said program.
The configuration file, `config.yaml`, is where you set everything about where you want
to grab logs from. 

To run the `logman` program:

`go run .` in the root directory of the project.

To build the `logman` program:

`go build .` and then a binary will appear in the `bin` directory of your `$GOPATH`.

After that, just run the binary and you should be golden.

## Troubleshooting
The GitHub issue board is where I post issues with the program. If you any issue with the program, please post it there. Go panics are a great example of an issue you can post.

I'm new to Go, so forgive the long wait time for any bux fix.

## Contributing
If you would like to contribute, feel free to email me or send in pull requests. `todo.md` should have some things to get you rolling. There's quite a few items on there that I want added to make `logman` a finished product.
