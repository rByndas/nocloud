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

	"github.com/slntopp/nocloud/pkg/billing"
	pb "github.com/slntopp/nocloud/pkg/health/proto"
	"go.uber.org/zap"
)

const SERVICE = "Billing Machine"

type HealthServer struct {
	pb.UnimplementedInternalProbeServiceServer
	log *zap.Logger
	srv *billing.BillingServiceServer
	rec *billing.RecordsServiceServer
}

func NewHealthServer(log *zap.Logger, srv *billing.BillingServiceServer, rec *billing.RecordsServiceServer) *HealthServer {
	return &HealthServer{
		log: log, srv: srv, rec: rec,
	}
}

func (s *HealthServer) Service(_ context.Context, _ *pb.ProbeRequest) (*pb.ServingStatus, error) {
	return &pb.ServingStatus{
		Service: SERVICE,
		Status:  pb.Status_SERVING,
	}, nil
}

func (s *HealthServer) Routine(_ context.Context, _ *pb.ProbeRequest) (*pb.RoutinesStatus, error) {
	routines := append(s.srv.GenTransactionsRoutineState(), s.rec.ConsumerStatus)
	return &pb.RoutinesStatus{
		Routines: routines,
	}, nil
}
