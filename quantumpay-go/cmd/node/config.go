package main

import "flag"

type NodeConfig struct {
	NodeID   string
	RPCPort  string
	P2PPort  string
	DataDir  string
	Peers    string
}

func LoadConfig() NodeConfig {
	cfg := NodeConfig{}

	flag.StringVar(&cfg.NodeID, "node-id", "node1", "Node identifier")
	flag.StringVar(&cfg.RPCPort, "rpc-port", "8080", "RPC port")
	flag.StringVar(&cfg.P2PPort, "p2p-port", "7001", "P2P port")
	flag.StringVar(&cfg.DataDir, "data-dir", "./data/node1", "Data directory")
	flag.StringVar(&cfg.Peers, "peers", "", "Comma-separated peer list")

	flag.Parse()
	return cfg
}
