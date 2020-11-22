# README [![Build Status](https://travis-ci.org/c3sr/model.svg?branch=master)](https://travis-ci.org/c3sr/model) [![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg)](https://github.com/emersion/stability-badges#experimental)

## Developing

Make sure to run `go generate` anytime you change one of the structures in the code.

## Testing

To run a mongodb on localhost (for `ranking_test.go`), try

    docker run -d -p27017:27017 mongo:3.0
