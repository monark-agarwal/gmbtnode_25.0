/*
skycoin daemon
*/
package main

/*
CODE GENERATED AUTOMATICALLY WITH FIBER COIN CREATOR
AVOID EDITING THIS MANUALLY
*/

import (
	"flag"
	_ "net/http/pprof"
	"os"

	"github.com/skycoin/skycoin/src/readable"
	"github.com/skycoin/skycoin/src/skycoin"
	"github.com/skycoin/skycoin/src/util/logging"
)

var (
	// Version of the node. Can be set by -ldflags
	Version = "0.25.0"
	// Commit ID. Can be set by -ldflags
	Commit = ""
	// Branch name. Can be set by -ldflags
	Branch = ""
	// ConfigMode (possible values are "", "STANDALONE_CLIENT").
	// This is used to change the default configuration.
	// Can be set by -ldflags
	ConfigMode = ""

	logger = logging.MustGetLogger("main")

	// CoinName name of coin
	CoinName = "skycoin"

	// GenesisSignatureStr hex string of genesis signature
	GenesisSignatureStr = "09b4507e67c9e26d907f236dd46010d3cce010fe9768f6f2fd236b6618f1a86a5a3dc1f10ac1f111db93a4c128837ba74d943926b2f4861a70be526a4b3b94cf00"
	// GenesisAddressStr genesis address string
	GenesisAddressStr = "2c7eaGNtR3QtWWt4DMxzK82rbY3RSaBKiBf"
	// BlockchainPubkeyStr pubic key string
	BlockchainPubkeyStr = "0338fb95bb41f807d34905b2ec9ce16fdb18e9679c52457936e3445645a5c40aee"
	// BlockchainSeckeyStr empty private key string
	BlockchainSeckeyStr = ""

	// GenesisTimestamp genesis block create unix time
	GenesisTimestamp uint64 = 1620822930
	// GenesisCoinVolume represents the coin capacity
	GenesisCoinVolume uint64 = 5000e12

	// DefaultConnections the default trust node addresses
	DefaultConnections = []string{
	"88.198.193.236:5200",
        "157.90.154.25:5201",
        "157.90.154.25:5202",
        "157.90.154.25:5203",
        "95.217.220.148:5200",
        "157.90.160.253:5200",
        "116.203.120.94:5200",
        "157.90.154.25:5200",}

	nodeConfig = skycoin.NewNodeConfig(ConfigMode, skycoin.NodeParameters{
		CoinName:                       CoinName,
		GenesisSignatureStr:            GenesisSignatureStr,
		GenesisAddressStr:              GenesisAddressStr,
		GenesisCoinVolume:              GenesisCoinVolume,
		GenesisTimestamp:               GenesisTimestamp,
		BlockchainPubkeyStr:            BlockchainPubkeyStr,
		BlockchainSeckeyStr:            BlockchainSeckeyStr,
		DefaultConnections:             DefaultConnections,
		PeerListURL:                    "http://coin4.glbrain.com/blockchain/peers.txt",
		Port:                           5200,
		WebInterfacePort:               5220,
		DataDirectory:                  "$HOME/.gmbtcoin",
		UnconfirmedBurnFactor:          2,
		UnconfirmedMaxTransactionSize:  32768,
		UnconfirmedMaxDropletPrecision: 6,
		CreateBlockBurnFactor:          2,
		CreateBlockMaxTransactionSize:  32768,
		CreateBlockMaxDropletPrecision: 6,
		MaxBlockSize:                   32768,
	})

	parseFlags = true
)

func init() {
	nodeConfig.RegisterFlags()
}

func main() {
	if parseFlags {
		flag.Parse()
	}

	// create a new fiber coin instance
	coin := skycoin.NewCoin(skycoin.Config{
		Node: nodeConfig,
		Build: readable.BuildInfo{
			Version: Version,
			Commit:  Commit,
			Branch:  Branch,
		},
	}, logger)

	// parse config values
	if err := coin.ParseConfig(); err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	// run fiber coin node
	if err := coin.Run(); err != nil {
		os.Exit(1)
	}
}
