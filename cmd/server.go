package cmd

import (
	"net"

	"github.com/hulb/grpc-test/pb"
	"github.com/hulb/grpc-test/server"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

var (
	listen    string
	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "grpc server",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger, err := zap.NewDevelopment()
			if err != nil {
				return err
			}

			ls, err := net.Listen("tcp", listen)
			if err != nil {
				return err
			}
			defer ls.Close()

			s := grpc.NewServer()
			pb.RegisterGreeterServer(s, server.NewServer(logger))

			logger.Sugar().Infoln("start grpc server on ", listen)
			return s.Serve(ls)
		},
	}
)

func init() {
	serverCmd.Flags().StringVarP(&listen, "listen", "l", "127.0.0.1:4444", "grpc server listen addr")
}
