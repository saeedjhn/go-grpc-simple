package main

import (
	"log"
	"net"
	"time"

	grpcctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	pb "github.com/saeedjhn/go-grpc/blog/goproto"
	"github.com/saeedjhn/go-grpc/blog/server/interceptor"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"

	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
)

const (
	_addr                  = "0.0.0.0:50001"
	_maxConnectionIdle     = 5 * time.Minute
	_timeout               = 30 * time.Second
	_maxConnectionAge      = 2 * time.Hour
	_time                  = 1 * time.Minute
	_maxConnectionAgeGrace = 1 * time.Second
)

func main() {
	listen, err := net.Listen("tcp", _addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}

	log.Printf("Listening  on %s\n", _addr)

	intercep := interceptor.New()

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		// Closes idle connections after 5 minutes of inactivity.
		MaxConnectionIdle: _maxConnectionIdle,
		// Waits up to 20 seconds for the client to respond to a Keepalive probe before considering the connection dead.
		Timeout: _timeout,
		// Forces the connection to be refreshed after 2 hours, regardless of its state (active or idle).
		MaxConnectionAge: _maxConnectionAge,
		// Sends a Keepalive probe every 1 minute to check the health of the connection.
		Time: _time,
		// Typical Settings for MaxConnectionAgeGrace
		// Short Grace Period (1–5 seconds):
		//
		// Use when you want connections to refresh quickly and prioritize connection closure over
		// completing in-flight requests.
		// Suitable for real-time systems where connections need to be refreshed frequently to rebalance
		// load or apply new configurations.
		// Long Grace Period (30–60 seconds):
		//
		// Use when you want to ensure in-flight RPCs have sufficient time to finish before closing the connection.
		// Suitable for systems handling longer-running requests or batch operations.
		MaxConnectionAgeGrace: _maxConnectionAgeGrace,
	}),
		grpc.UnaryInterceptor(intercep.Logger),
		grpc.ChainUnaryInterceptor(
			grpcctxtags.UnaryServerInterceptor(),
			grpcprometheus.UnaryServerInterceptor,
			grpcrecovery.UnaryServerInterceptor(),
		),
	)

	pb.RegisterBlogServiceServer(server, &Server{})

	// if environment != configs.Production {
	reflection.Register(server)
	// }

	if err = server.Serve(listen); err != nil {
		log.Fatalf("Failed to server %v\n", err)
	}
}
