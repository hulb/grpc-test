package cmd

import (
	"fmt"
	"time"

	"github.com/hulb/grpc-test/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	serverAddr string

	clientCmd = &cobra.Command{
		Use:   "client",
		Short: "start grpc client and send request, print resp",
		RunE: func(cmd *cobra.Command, args []string) error {
			conn, err := grpc.Dial(serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				return err
			}
			client := pb.NewGreeterClient(conn)
			resp, err := client.SayHi(cmd.Context(), &pb.Request{
				Msg:  "hi",
				Time: timestamppb.New(time.Now()),
			})
			if err != nil {
				return err
			}

			respMsg := resp.GetBack()
			fmt.Println("resp msg: ", respMsg)
			respVlue := resp.GetValue()
			fields := respVlue.GetFields()
			for key, v := range fields {
				fmt.Println("key: ", key, " value kind: ", v.Kind)
			}

			return nil
		},
	}
)

func init() {
	clientCmd.Flags().StringVarP(&serverAddr, "server", "s", "", "server addr")
	clientCmd.MarkFlagRequired("server")
}
