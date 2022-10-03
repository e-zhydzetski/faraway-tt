# faraway-tt

[![Main Build Status](https://github.com/e-zhydzetski/faraway-tt/actions/workflows/main.yml/badge.svg)](https://github.com/e-zhydzetski/faraway-tt/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/e-zhydzetski/faraway-tt/branch/master/graph/badge.svg?token=Z7IWED0VRR)](https://codecov.io/gh/e-zhydzetski/faraway-tt)

## Task

Design and implement “Word of Wisdom” tcp server.
* TCP server should be protected from DDOS attacks with the Prof of Work (https://en.wikipedia.org/wiki/Proof_of_work), the challenge-response protocol should be used.
* The choice of the POW algorithm should be explained.
* After Prof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other collection of the quotes.
* Docker file should be provided both for the server and for the client that solves the POW challenge

## POW

As a first implementation of POW verification algorithm I chose bruteforce of [bcrypt](https://en.wikipedia.org/wiki/Bcrypt) hash:
* crypto function
* embedded salt, convenient
* configurable computational complexity
* popular, implementations are available for any lang

Desired key is a non-negative integer number, right border is unknown for a client. So client should iterate keys from zero to success.  
At the server side, right border used as a difficulty level.  
Currently, key is a crypto random number in `[0;difficulty)`

## Quoter source

[Quotable](https://github.com/lukePeavey/quotable)

## Dev environment

* [go](https://go.dev/)
* [docker](https://www.docker.com/)
* [golangci-lint](https://golangci-lint.run/)
* [task](https://taskfile.dev/)

## Demo

`task demo` - build images and setup docker-compose based env with multiple server and client instances