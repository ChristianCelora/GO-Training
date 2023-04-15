# GO-Training

Simple execises to learn GOLang.

## Dev Env

I created a custom Docker container. 

To build it run:

```sh
docker build -t go-train-server .
```

To run it use:

```sh
docker run --rm -v "$(pwd)"/src:/go/src -it go-train-server
```

in depth:
 - the --rm delete the container once it finished the execution
 - the -v mount the folder with the code
 - the -it 

## Go Commands

### Help command

Always useful to use it.

```sh
go help           # basic
go help <command> # command help
```

### Simple run

To build and run with one command use:

```sh
go run src/hello_world/hello_world.go
```

### Compile and run!

To compile and create the executable file use the `build` command. 
The -o flag permit to specify the output path.

```sh
go build -o src/hello_world/hello_world src/hello_world/hello_world.go
./src/hello_world/hello_world # Run the executable
```