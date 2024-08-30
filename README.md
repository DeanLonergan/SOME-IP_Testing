# Testing vsomeip using Golang

The purpose of the repository is to develop a simple vsomeip service that can be used as a foundation to evaluate the suitability of using Golang to write BDD tests.

![vsomeip](https://github.com/user-attachments/assets/65705f6f-0aa6-4239-abb5-4d2d6a621da9)

Currently, the client/service notifications that are tested using Golang (using the [Godog](https://pkg.go.dev/github.com/cucumber/godog) framework:

- SUBSCRIBE ACK
  

## Build and Run
Assuming [vsomeip](https://github.com/COVESA/vsomeip) and [Go](https://go.dev/doc/install) have been installed, simply running the *build* or *test* scripts acquire dependencies, build the project and, in the case of *test*, run the Godog tests.

Building the project:
```
git clone git@github.com:DeanLonergan/SOME-IP_Testing.git
cd SOME-IP_Testing
python3 build.py
```
Running the tests, assuming the project has not been built (or the current build is out of date):
```
python3 test.py
```
Running the tests, assuming the build is up to date:
```
cd test
go test
```
<br>

**Note**:  The *test* script handles the installation and initialisation of godog, but this can also be achieved manually:
```
go mod init someip_testing
go get github.com/cucumber/godog/cmd/godog@latest
```

## Resources
The client and service are based on the following vsomeip guide by COVESA:\
https://github.com/COVESA/vsomeip/wiki/vsomeip-in-10-minutes

The routing manager (routingmanagerd.cdd) is from the COVESA vsomeip repository:\
https://github.com/COVESA/vsomeip/tree/master/examples/routingmanagerd
