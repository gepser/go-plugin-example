# go-plugin-calculator
A simple calculator just to try out the new Go plugins.

## Build the plugin(s)

Right now this only works on linux.

    go build -buildmode=plugin -o add.so add.go

## Run the program

    go run main.go

If everything is ok, the output should be:

    3

## Next steps

- Add a Dockerfile so this can be executed everywhere.
- Add a Makefile with commands to build the plugins.
- Add more plugins.
- Move plugins to a folder.
- Discover the plugins recursively (inside a folder by convention or by input parameters or both).
- Possibly make the plugins comply with an interface.
