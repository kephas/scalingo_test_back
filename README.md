# Scalingo Backend technical test

After cloning, here's how to setup and run the server:

```sh
$ bower install
$ go build -o gh-search
$ ./gh-search
```

The program accepts 3 command-line arguments:

 * `--port/-p` for the HTTP port to listen to (default: 8000)
 * `--limit/-l` for the number of results to get from Github (default: 100)
 * `--workers/-w` for the number of parallel workers (default: 10)
