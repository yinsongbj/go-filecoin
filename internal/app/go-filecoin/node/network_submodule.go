package node

import (
	"github.com/filecoin-project/go-filecoin/internal/pkg/net"
	exchange "github.com/ipfs/go-ipfs-exchange-interface"
	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/routing"
	libp2pps "github.com/libp2p/go-libp2p-pubsub"
)

// NetworkSubmodule enhances the `Node` with networking capabilities.
type NetworkSubmodule struct {
	NetworkName string

	host host.Host

	// TODO: do we need a second host? (see: https://github.com/filecoin-project/go-filecoin/issues/3477)
	PeerHost host.Host

	// Router is a router from IPFS
	Router routing.Routing

	fsub *libp2pps.PubSub

	// TODO: split chain bitswap from storage bitswap (issue: ???)
	bitswap exchange.Interface

	Network *net.Network
}
