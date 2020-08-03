# Golang+Vue websocket chat

This is golang+vue application with built-in websocket chat.

Application features:
  * Token auth
  * View a list of all users
  * Add user as friend
  * Making conversations with friends
  * Chatting

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Check that you have installed:
  * [golang 1.14+](https://golang.org/) - A open source programming language that makes it easy to build simple, reliable, and efficient software.
  * [gcc/g++](https://gcc.gnu.org/) - The GNU Compiler Collection includes front ends for C, C++, Objective-C, Fortran, Ada, Go, and D, as well as libraries for these languages (libstdc++,...).

### Installing

1. Install gcc/g++
    * For ***Windows*** install [tdm-gcc](https://jmeubank.github.io/tdm-gcc/) and restart shell
    * For ***Linux***(Ubuntu) run `sudo apt-get install build-essential` to install [gcc/g++](https://gcc.gnu.org/)
2. Go to the project folder
3. Run `go build chat.go`
4. Rename file **config.example.yaml** to **config.yaml**
5. Run `./chat migrate` for database creation and table migration
6. Run `./chat serve` to run app in browser
7. Open http://localhost:4545 in your browser.

## App screenshots
<p>
  <img src="/screenshots/sign_in.png" alt="SignIn" width="400" hspace="10">
  <img src="/screenshots/new_chat.png" alt="New chat" width="400" hspace="10">
</p>
<p>
  <img src="/screenshots/chats.png" alt="Chats" width="400" hspace="10">
  <img src="/screenshots/friends.png" alt="Friends" width="400" hspace="10">
</p>

## Built With

* [Golang](https://golang.org/) - A open source programming language that makes it easy to build simple, reliable, and efficient software
* [Gorilla websocket](https://github.com/gorilla/websocket) - A fast, well-tested and widely used WebSocket implementation for Go
* [Vue](https://vuejs.org/) - The Progressive JavaScript Framework
* [Vuetify](https://vuetifyjs.com/en/) - Material Design Component Framework

