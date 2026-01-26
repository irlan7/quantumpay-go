package p2pv1

type Config struct {
	ListenAddress string   // contoh: "0.0.0.0:7701"
	StaticPeers   []string // contoh: []{"1.2.3.4:7701"}
}
