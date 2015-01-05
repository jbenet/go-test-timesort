# go-test-timesort

go-test-timesort sorts `go-test` combined output by running time.
(maybe go-test can do this?)

Install:

```
go get github.com/jbenet/go-test-timesort
```

Use:

```
> go test ./... | go-test-timesort
ok    github.com/jbenet/go-ipfs/blocks  0.016s
ok    github.com/jbenet/go-ipfs/pin 0.016s
ok    github.com/jbenet/go-ipfs/p2p/crypto/secio  0.014s
ok    github.com/jbenet/go-ipfs/util/eventlog 0.013s
ok    github.com/jbenet/go-ipfs/blocks/bloom  0.012s
ok    github.com/jbenet/go-ipfs/routing/keyspace  0.008s
?     github.com/jbenet/go-ipfs [no test files]
?     github.com/jbenet/go-ipfs/blocks/blocksutil [no test files]
```
