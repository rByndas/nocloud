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
package main

import (
	"context"
	"fmt"
	"net"

	billing "github.com/slntopp/nocloud/pkg/billing"
	pb "github.com/slntopp/nocloud/pkg/billing/proto"
	healthpb "github.com/slntopp/nocloud/pkg/health/proto"
	"github.com/slntopp/nocloud/pkg/nocloud"
	"github.com/slntopp/nocloud/pkg/nocloud/auth"
	"github.com/slntopp/nocloud/pkg/nocloud/connectdb"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	port string
	log *zap.Logger

	arangodbHost 	string
	arangodbCred 	string
	SIGNING_KEY		[]byte
	statesHost  string
)

func init() {
	viper.AutomaticEnv()
	log = nocloud.NewLogger()

	viper.SetDefault("PORT", "8080")

	viper.SetDefault("DB_HOST", "db:8529")
	viper.SetDefault("DB_CRED", "root:openSesame")
	viper.SetDefault("DRIVERS", "")
	viper.SetDefault("EXTENTION_SERVERS", "")
	viper.SetDefault("SIGNING_KEY", "seeeecreet")
	viper.SetDefault("STATES_HOST", "states:8080")

	port = viper.GetString("PORT")

	arangodbHost 	= viper.GetString("DB_HOST")
	arangodbCred 	= viper.GetString("DB_CRED")
	SIGNING_KEY 	= []byte(viper.GetString("SIGNING_KEY"))
	statesHost 	= viper.GetString("STATES_HOST")
}

func main() {
	defer func() {
		_ = log.Sync()
	}()

	log.Info("Setting up DB Connection")
	db := connectdb.MakeDBConnection(log, arangodbHost, arangodbCred)
	log.Info("DB connection established")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatal("Failed to listen", zap.String("address", port), zap.Error(err))
	}

	log.Debug("Init Connection with States", zap.String("host", statesHost))
	opts := []grpc.DialOption{
		grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(statesHost, opts...)
	if err != nil {
		log.Fatal("fail to dial States", zap.Error(err))
	}
	defer conn.Close()

	auth.SetContext(log, SIGNING_KEY)
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(log),
			grpc.UnaryServerInterceptor(auth.JWT_AUTH_INTERCEPTOR),
		)),
	)
	
	server := billing.NewBillingServiceServer(log, db)

	token, err := auth.MakeToken(schema.ROOT_ACCOUNT_KEY)
	if err != nil {
		log.Fatal("Can't generate token", zap.Error(err))
	}
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "bearer " + token)
	go server.GenTransactionsRoutine(ctx)
	pb.RegisterBillingServiceServer(s, server)

	records := billing.NewRecordsServiceServer(log, db)
	go records.Consume(ctx)
	pb.RegisterRecordsServiceServer(s, records)

	healthpb.RegisterInternalProbeServiceServer(s, NewHealthServer(log, server, records))

	log.Info(fmt.Sprintf("Serving gRPC on 0.0.0.0:%v", port), zap.Skip())
	log.Fatal("Failed to serve gRPC", zap.Error(s.Serve(lis)))
}