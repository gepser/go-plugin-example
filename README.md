# <img src="./images/plugin.png" alt="plugin" width="50px" height="50px"> Go Plugin Example
A simple app just to try out the new Go plugins.

It simulates drivers for connecting to some databases, so each database should have their own plugin.

## Build the plugin(s)

Right now this only works on linux.

    go build -buildmode=plugin -o mysql.so mysql.go

## Run the program

Linux:

    go run main.go

Docker:

    docker build -t go-plugin-example . && docker run -ti go-plugin-example

If everything is ok, the output should look like this:

    [mysql]: Establishing connection to: http://localhost:3306/
    [mysql]: Inserting data: Golang
    [mysql]: Getting some data: Gepser

