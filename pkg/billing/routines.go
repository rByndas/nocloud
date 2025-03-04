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
package billing

import (
	"context"
	"time"

	hpb "github.com/slntopp/nocloud/pkg/health/proto"
	"github.com/slntopp/nocloud/pkg/nocloud/schema"
	sc "github.com/slntopp/nocloud/pkg/settings/client"
	settingspb "github.com/slntopp/nocloud/pkg/settings/proto"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	settingsClient settingspb.SettingsServiceClient
)

// Settings Key storing routine conf
const monFreqKey string = "billing-gen-transactions-routine"

type RoutineConf struct {
	Frequency int `json:"freq"` // Frequency in Seconds
}

var (
	defaultSetting = &sc.Setting[RoutineConf]{
		Value: RoutineConf{
			Frequency: 60,
		},
		Description: "Transactions Generating and Processing Routine Configuration",
		Public:      false,
	}
)

func init() {
	viper.AutomaticEnv()
	viper.SetDefault("SETTINGS_HOST", "settings:8000")
	host := viper.GetString("SETTINGS_HOST")

	conn, err := grpc.Dial(host, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	settingsClient = settingspb.NewSettingsServiceClient(conn)
}

func MakeConf(ctx context.Context, log *zap.Logger) (conf RoutineConf) {
	sc.Setup(log, ctx, &settingsClient)

	err := sc.Fetch(monFreqKey, &conf, defaultSetting)
	if err != nil {
		return defaultSetting.Value
	}

	return conf
}

func (s *BillingServiceServer) GenTransactionsRoutineState() []*hpb.RoutineStatus {
	return []*hpb.RoutineStatus{
		s.gen, s.proc,
	}
}

func (s *BillingServiceServer) GenTransactionsRoutine(ctx context.Context) {
	log := s.log.Named("Routine")

	conf := MakeConf(ctx, log)
	log.Info("Got Configuration", zap.Any("conf", conf))
	ticker := time.NewTicker(time.Second * time.Duration(conf.Frequency))
	tick := time.Now()
	for {
		log.Info("Starting Generating Transactions Sub-Routine", zap.Time("tick", tick))
		s.gen.Status.Status = hpb.Status_RUNNING
		s.gen.Status.Error = nil
		_, err := s.db.Query(ctx, generateTransactions, map[string]interface{}{
			"@transactions": schema.TRANSACTIONS_COL,
			"@instances":    schema.INSTANCES_COL,
			"@services":     schema.SERVICES_COL,
			"@records":      schema.RECORDS_COL,
			"@accounts":     schema.ACCOUNTS_COL,
			"permissions":   schema.PERMISSIONS_GRAPH.Name,
			"now":           tick.Unix(),
		})
		if err != nil {
			log.Error("Error Generating Transactions", zap.Error(err))
			s.gen.Status.Status = hpb.Status_HASERRS
			err_s := err.Error()
			s.gen.Status.Error = &err_s
		}
		s.gen.LastExecution = tick.Format("2006-01-02T15:04:05Z07:00")

		log.Info("Starting Processing Transactions Sub-Routine", zap.Time("tick", tick))
		s.proc.Status.Status = hpb.Status_RUNNING
		s.gen.Status.Error = nil
		_, err = s.db.Query(ctx, processTransactions, map[string]interface{}{
			"@transactions": schema.TRANSACTIONS_COL,
			"@accounts":     schema.ACCOUNTS_COL,
			"accounts":      schema.ACCOUNTS_COL,
			"now":           tick.Unix(),
		})
		if err != nil {
			log.Error("Error Processing Transactions", zap.Error(err))
			s.proc.Status.Status = hpb.Status_HASERRS
			err_s := err.Error()
			s.proc.Status.Error = &err_s
		}

		s.proc.LastExecution = tick.Format("2006-01-02T15:04:05Z07:00")
		tick = <-ticker.C
	}
}

const generateTransactions = `
FOR service IN @@services // Iterate over Services
	LET instances = (
        FOR i IN 2 OUTBOUND service
        GRAPH @permissions
        FILTER IS_SAME_COLLECTION(@@instances, i)
            RETURN i._key )

    LET account = LAST( // Find Service owner Account
    FOR node, edge, path IN 2
    INBOUND service
    GRAPH @permissions
    FILTER path.edges[*].role == ["owner","owner"]
    FILTER IS_SAME_COLLECTION(node, @@accounts)
        RETURN node
    )
    
    LET records = ( // Collect all unprocessed records
        FOR record IN @@records
        FILTER record.exec <= @now
        FILTER !record.processed
        FILTER record.instance IN instances
            UPDATE record._key WITH { processed: true } IN @@records RETURN NEW
    )
    
    FILTER LENGTH(records) > 0 // Skip if no Records (no empty Transaction)
    INSERT {
        exec: @now, // Timestamp in seconds
        processed: false,
        account: account._key,
        service: service._key,
        records: records[*]._key,
        total: SUM(records[*].total) // Calculate Total
    } IN @@transactions RETURN NEW
`

const processTransactions = `
FOR t IN @@transactions // Iterate over Transactions
FILTER t.exec <= @now
FILTER !t.processed
    LET account = DOCUMENT(CONCAT(@accounts, "/", t.account))
    UPDATE account WITH { balance: account.balance - t.total } IN @@accounts
    UPDATE t WITH { processed: true, proc: @now } IN @@transactions
`
