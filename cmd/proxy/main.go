package main

import (
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/slntopp/nocloud/pkg/graph"
	"github.com/slntopp/nocloud/pkg/nocloud"
	"github.com/slntopp/nocloud/pkg/nocloud/connectdb"
	"github.com/slntopp/nocloud/pkg/proxy"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	healthpb "github.com/slntopp/nocloud/pkg/health/proto"
)

var (
	log *zap.Logger

	arangodbHost string
	arangodbCred string

	SIGNING_KEY []byte
)

func init() {
	viper.AutomaticEnv()
	log = nocloud.NewLogger()

	viper.SetDefault("DB_HOST", "db:8529")
	viper.SetDefault("DB_CRED", "root:openSesame")

	viper.SetDefault("SIGNING_KEY", "seeeecreet")

	SIGNING_KEY = []byte(viper.GetString("SIGNING_KEY"))

	arangodbHost = viper.GetString("DB_HOST")
	arangodbCred = viper.GetString("DB_CRED")
}

func main() {
	defer func() {
		_ = log.Sync()
	}()

	log.Info("Setting up DB Connection")
	db := connectdb.MakeDBConnection(log, arangodbHost, arangodbCred)
	log.Info("DB connection established")

	ctrl := graph.NewServicesProvidersController(log, db)

	proxy.Setup(log, ctrl)

	r := mux.NewRouter()
	r.Use(AuthMiddleware)
	r.HandleFunc("/socket", proxy.Handler).Methods("GET")
	r.Use(mux.CORSMethodMiddleware(r))

	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("Failed to listen", zap.String("address", ":8000"), zap.Error(err))
	}
	s := grpc.NewServer()
	healthpb.RegisterInternalProbeServiceServer(s, NewHealthServer(log))

	go s.Serve(lis)

	/* #nosec */
	log.Fatal("Failed to serve proxy", zap.Error(http.ListenAndServe(":8080", r)))
}
