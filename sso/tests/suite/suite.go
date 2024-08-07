package suite

import (
	"context"
	"net"
	"strconv"
	"testing"

	ssov1 "github.com/SicParv1sMagna/sso-service/gen/go/sso"
	"github.com/SicParv1sMagna/sso-service/sso/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T
	Cfg        *config.Config
	AuthClient ssov1.AuthClient
}

const configPath = "../../config/local_tests_config.yaml"

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoadPath(configPath)

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	grpcAddress := net.JoinHostPort("::", strconv.Itoa(cfg.GRPC.Port))

	cc, err := grpc.DialContext(context.Background(),
		grpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	authClient := ssov1.NewAuthClient(cc)

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		AuthClient: authClient,
	}
}
