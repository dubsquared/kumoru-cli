# Kumoru CLI client

This repository holds the official [Kumoru.io](kumoru.io) CLI.

## Usage

Run `kumoru --help` to see usage.

## Installing

The CLI can be installed 2 ways:

### Retrieving a release

1. See the official [Releases](https://github.com/kumoru/kumoru-cli/releases) and download the binary for your system.
2. Unzip the archive
3. Place the kumoru binary in the directory of your choice.

### Building locally

If you prefer to get the latest code and build it your self, you'll need:

* Go 1.5
* `godeps`

1. Clone this repository
2. Run `make build` to build it for your local system

Alternatively you can run `make release` to start the full release process which includes tests. This will generate a linux and osx binary.

## Authors
Victor Palma <victor@kumoru.io>
Ryan Richard <ryan@kumoru.io>
