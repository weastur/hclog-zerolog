// main package is used to demonstrate how to use hclog-zerolog with go-hclog
package main

import (
	"github.com/hashicorp/raft"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	hclogzerolog "github.com/weastur/hclog-zerolog"
)

const ComponentCtxKey = "component"

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.With().Str(ComponentCtxKey, "core").Logger()

	log.Info().Msg("Starting")

	raftLogger := log.With().Str(ComponentCtxKey, "raft").Logger()

	config := raft.DefaultConfig()
	config.LocalID = raft.ServerID("node1")
	config.Logger = hclogzerolog.New(raftLogger)

	logs := raft.NewInmemStore()
	stable := raft.NewInmemStore()

	snapshots := raft.NewInmemSnapshotStore()
	addr, transport := raft.NewInmemTransport("node1")

	r, err := raft.NewRaft(config, nil, logs, stable, snapshots, transport)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create raft")
	}

	future := r.BootstrapCluster(raft.Configuration{
		Servers: []raft.Server{
			{
				ID:      config.LocalID,
				Address: addr,
			},
		},
	})
	if err := future.Error(); err != nil {
		log.Fatal().Err(err).Msg("failed to bootstrap cluster")
	}

	log.Info().Msg("Started")
	select {} // Block forever
}
