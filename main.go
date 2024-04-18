package main

import (
	db "DistributedKV/DB"
	web "DistributedKV/Web"
	config "DistributedKV/config"
	replication "DistributedKV/replication"
	"flag"
	"fmt"
	"log"
	"net/http"
	//"github.com/BurntSushi/toml"
)

var (
	//Creating a flag name "db-location", with default value "", and usage "The path to your DB"
	dbLocation = flag.String("db-location", "", "The path to your DB")
	httpAddr   = flag.String("http-addr", "127.0.0.1:8080", "HTTP Host address and port number")
	configFile = flag.String("config-file", "sharding.toml", "Config file for static sharding")
	shard      = flag.String("shard", "", "The name of the shard for the data")
	replica    = flag.Bool("replica", false, "run as a readonly replica")
)

func parseFlags() {

	flag.Parse()
	if *dbLocation == "" {
		log.Fatal("Must provide db location")
	}
	if *shard == "" {
		log.Fatal("Must provide shard")
	}
}

func main() {

	parseFlags()
	fmt.Println("Hello, Go!")

	c, err := config.ParseFile(*configFile)
	if err != nil {
		log.Fatalf("Error parsing config %q: %v", *configFile, err)
	}

	shards, err := config.ParseShards(c.Shard, *shard)
	if err != nil {
		log.Fatalf("Error parsing shards config: %v", err)
	}

	log.Printf("Shard count is %d, current shard: %d", shards.Count, shards.CurIdx)

	db, close, err := db.NewDatabase(*dbLocation, *replica)
	if err != nil {
		log.Fatalf("NewDatabase(%q): %v", *dbLocation, err)
	}
	defer close()

	if *replica {
		leaderAddr, ok := shards.Addrs[shards.CurIdx]
		if !ok {
			log.Fatalf("Could not find address for leader for shard %d", shards.CurIdx)
		}
		go replication.ClientLoop(db, leaderAddr)
	}

	srv := web.NewServer(db, shards)

	http.HandleFunc("/get", srv.GetHandler)
	http.HandleFunc("/set", srv.SetHandler)
	http.HandleFunc("/purge", srv.DeleteExtraKeysHandler)
	http.HandleFunc("/next-replication-key", srv.GetNextKeyForReplication)
	http.HandleFunc("/delete-replication-key", srv.DeleteReplicationKey)

	log.Fatal(http.ListenAndServe(*httpAddr, nil))
}
