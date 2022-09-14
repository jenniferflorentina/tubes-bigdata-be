## Setup:

### If building from source:

1. Once this repository is cloned and golang is installed in the system, navigate to this directory and run

```
go mod download
```

2. Once the dependencies are downloaded, using sample.env as reference either create a file called .env with the same keys or directly configure same keys as environment variables.
3. After the configuration and ensuring that the db is operational, run either

```
go build # to get the executable to run

or

go run main.go
```

