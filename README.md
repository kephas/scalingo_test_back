# Scalingo Backend technical test

Here's how to install and run the server:

```sh
$ go get github.com/kephas/scalingo_test_back
$ scalingo_test_back
```

The program accepts 3 command-line arguments:

 * `--port/-p` for the HTTP port to listen to (default: 8000)
 * `--limit/-l` for the number of results to get from Github (default: 100)
 * `--workers/-w` for the number of parallel workers (default: 10)
