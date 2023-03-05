# Golang CLI Application for Server and IP Search

This is a simple CLI application written in Go that allows users to search for servers and their corresponding IP addresses using the native `net` package and the third-party `urfave/cli` package.


## Installation

To install the application, you will need to have Go installed on your machine. Once Go is installed, you can use the following command to download and install the application:

```bash

go get github.com/samuelharo97/cli-ip-finder

```

## Usage

To use the application, navigate to the directory where the application is installed and run the following command:

```bash

# to find IPs run:
./cli-ip-finder ip --host <hostname>

# to find Servers run:
./cli-ip-finder servers --host <hostname>

```
Replace <hostname> with the name of the server you wish to search for.

The application will return the IP address of the server, if it can be found.

## Contributing

If you wish to contribute to this project, please fork the repository and submit a pull request.