// go-test-timesort sorts `go-test` combined output by running time.
//
// Note: maybe go-test can do this? writing this was faster than checking docs.
// that's how cool golang + stdlib is.
//
// Install:
//
//    go get github.com/jbenet/go-test-timesort
//
// Use:
//
//   > go test ./... | go-test-timesort
//   ok    github.com/jbenet/go-ipfs/blocks  0.016s
//   ok    github.com/jbenet/go-ipfs/pin 0.016s
//   ok    github.com/jbenet/go-ipfs/p2p/crypto/secio  0.014s
//   ok    github.com/jbenet/go-ipfs/util/eventlog 0.013s
//   ok    github.com/jbenet/go-ipfs/blocks/bloom  0.012s
//   ok    github.com/jbenet/go-ipfs/routing/keyspace  0.008s
//   ?     github.com/jbenet/go-ipfs [no test files]
//   ?     github.com/jbenet/go-ipfs/blocks/blocksutil [no test files]
//
package main
