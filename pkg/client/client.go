package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	//"google.golang.org/grpc/credentials"

	"github.com/openinfradev/tks-contract/pkg/log"
	pb "github.com/openinfradev/tks-proto/tks_pb"
)

var (
	conn *grpc.ClientConn
	clusterInfoClient pb.ClusterInfoServiceClient
	cspInfoClient pb.CspInfoServiceClient
)

func RequestLogging() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		end := time.Now()

		log.Info(fmt.Sprintf("[GRPC:%s][START:%s][END:%s][ERR:%v]", method, start.Format(time.RFC3339), end.Format(time.RFC3339), err))
		log.Debug(fmt.Sprintf("[GRPC:%s][REQUEST %s][REPLY %s]", method, req, reply))
		
		return err
	}
}

func GetConnection(host string) (*grpc.ClientConn, error) {
	if conn == nil {
		_conn, err := grpc.Dial(
			host,
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(
				grpc_middleware.ChainUnaryClient(
					RequestLogging(),
				),
			),
		)
		if err != nil {
			return nil, err
		}
		conn = _conn
	} 
	return conn, nil
}

func GetClusterInfoClient(address string, port int, caller string) (pb.ClusterInfoServiceClient, error) {
	conn, err := GetConnection( fmt.Sprintf("%s:%d", address, port) )
	if err != nil {
		return nil, err
	} 

	clusterInfoClient = pb.NewClusterInfoServiceClient(conn)
	return clusterInfoClient, nil
}

func GetCspInfoClient(address string, port int, caller string) (pb.CspInfoServiceClient, error) {
	conn, err := GetConnection( fmt.Sprintf("%s:%d", address, port) )
	if err != nil {
		return nil, err
	}

	cspInfoClient = pb.NewCspInfoServiceClient(conn)
	return cspInfoClient, nil
}
