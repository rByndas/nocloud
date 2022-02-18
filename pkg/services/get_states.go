/*
Copyright © 2021-2022 Nikita Ivanovski info@slnt-opp.xyz

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package services

import (
	"context"

	pb "github.com/slntopp/nocloud/pkg/services/proto"
	sspb "github.com/slntopp/nocloud/pkg/statuses/proto"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	statuses_host string
)

func init() {
	viper.SetDefault("STATUSES_HOST", "statuses:8080")
	statuses_host = viper.GetString("STATUSES_HOST")
}

//Gets statuses Instanses of Servce from pkg/statuses
func (s *ServicesServiceServer) GetStates(ctx context.Context, request *pb.GetStatesRequest) (*pb.GetStatesResponse, error) {
	log := s.log.Named("TestServiceConfig")

	opts := []grpc.DialOption{
		grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(statuses_host, opts...)
	if err != nil {
		log.Fatal("fail to dial Statuses", zap.Error(err))
	}
	defer conn.Close()

	grpc_client := sspb.NewPostServiceClient(conn)

	service, err := s.Get(ctx, &pb.GetRequest{
		Uuid: request.Uuid,
	})
	if err != nil {
		log.Error("fail to get Services:", zap.Error(err))
	}

	resp, err := grpc_client.StateGet(ctx, service)
	if err != nil {
		log.Error("fail to send statuses:", zap.Error(err))
	}

	return resp, nil
}
