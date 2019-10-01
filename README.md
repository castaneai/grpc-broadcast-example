# grpc-broadcast-example

gRPC broadcast example with BiDirectional streaming RPCs

## Usage

First, run `server/server.go` then gRPC server listening on `0.0.0.0:50051`.

```
$ go run server/server.go
```

Next run `client/client.go` with argument (client's name).

For example if you run 2 clients named 'Alice' and 'Bob', you will get the following output:

```
$ go run client/client.go Alice
```

```
$ go run client/client.go Bob
```

```
2019/10/01 21:00:04 new user: a85c7947-011e-4f0f-9805-ebe9ee2a9a3e
2019/10/01 21:00:04 broadcast: hello, I'm Alice
2019/10/01 21:00:05 broadcast: hello, I'm Alice
...
2019/10/01 21:00:20 new user: ac58baff-e689-4bfc-84e7-60149673079b
2019/10/01 21:00:20 broadcast: hello, I'm Bob
2019/10/01 21:00:21 broadcast: hello, I'm Alice
2019/10/01 21:00:21 broadcast: hello, I'm Bob
2019/10/01 21:00:22 broadcast: hello, I'm Alice
2019/10/01 21:00:22 broadcast: hello, I'm Bob
2019/10/01 21:00:23 broadcast: hello, I'm Alice
2019/10/01 21:00:23 broadcast: hello, I'm Bob
2019/10/01 21:00:24 broadcast: hello, I'm Alice
2019/10/01 21:00:24 broadcast: hello, I'm Bob
...
```
