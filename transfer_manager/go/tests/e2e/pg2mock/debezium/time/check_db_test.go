package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/doublecloud/transfer/transfer_manager/go/internal/logger"
	"github.com/doublecloud/transfer/transfer_manager/go/pkg/abstract"
	server "github.com/doublecloud/transfer/transfer_manager/go/pkg/abstract/model"
	pgcommon "github.com/doublecloud/transfer/transfer_manager/go/pkg/providers/postgres"
	"github.com/doublecloud/transfer/transfer_manager/go/pkg/providers/postgres/pgrecipe"
	"github.com/doublecloud/transfer/transfer_manager/go/tests/helpers"
	"github.com/stretchr/testify/require"
)

var (
	Source = *pgrecipe.RecipeSource(pgrecipe.WithPrefix(""), pgrecipe.WithInitDir("init_source"))
)

func init() {
	_ = os.Setenv("YC", "1") // to not go to vanga
	Source.WithDefaults()
}

//---------------------------------------------------------------------------------------------------------------------

func TestSnapshotAndReplication(t *testing.T) {
	defer require.NoError(t, helpers.CheckConnections(
		helpers.LabeledPort{Label: "PG source", Port: Source.Port},
	))

	container := helpers.NewTestCaseContainer()
	container.AddCase(newContainerTimeWithTZ())
	container.AddCase(newContainerTime())
	container.Initialize(t)

	//------------------------------------------------------------------------------

	sinker := &helpers.MockSink{}
	target := server.MockDestination{
		SinkerFactory: func() abstract.Sinker { return sinker },
		Cleanup:       server.DisabledCleanup,
	}
	transfer := helpers.MakeTransfer("fake", &Source, &target, abstract.TransferTypeSnapshotAndIncrement)

	sinker.PushCallback = func(input []abstract.ChangeItem) {
		for _, el := range input {
			container.AddChangeItem(t, &el)
		}
	}

	worker := helpers.Activate(t, transfer)
	defer worker.Close(t)

	//-----------------------------------------------------------------------------------------------------------------
	// execute SQL statements

	srcConn, err := pgcommon.MakeConnPoolFromSrc(&Source, logger.Log)
	require.NoError(t, err)
	defer srcConn.Close()

	container.ExecStatement(context.Background(), t, srcConn)

	//-----------------------------------------------------------------------------------------------------------------

	for {
		time.Sleep(time.Second)

		if container.IsEnoughChangeItems(t) {
			break
		}
	}

	container.Check(t)
}