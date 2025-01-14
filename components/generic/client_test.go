package generic_test

import (
	"context"
	"errors"
	"net"
	"testing"

	"github.com/edaniels/golog"
	componentpb "go.viam.com/api/component/generic/v1"
	"go.viam.com/test"
	"go.viam.com/utils"
	"go.viam.com/utils/rpc"
	"google.golang.org/grpc"

	"go.viam.com/rdk/components/generic"
	viamgrpc "go.viam.com/rdk/grpc"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/subtype"
	"go.viam.com/rdk/testutils"
	"go.viam.com/rdk/testutils/inject"
)

func TestClient(t *testing.T) {
	logger := golog.NewTestLogger(t)
	listener1, err := net.Listen("tcp", "localhost:0")
	test.That(t, err, test.ShouldBeNil)
	rpcServer, err := rpc.NewServer(logger, rpc.WithUnauthenticated())
	test.That(t, err, test.ShouldBeNil)

	workingGeneric := &inject.Generic{}
	failingGeneric := &inject.Generic{}

	workingGeneric.DoFunc = generic.EchoFunc
	failingGeneric.DoFunc = func(
		ctx context.Context,
		cmd map[string]interface{},
	) (
		map[string]interface{},
		error,
	) {
		return nil, errors.New("do failed")
	}

	resourceMap := map[resource.Name]interface{}{
		generic.Named(testGenericName): workingGeneric,
		generic.Named(failGenericName): failingGeneric,
	}
	genericSvc, err := subtype.New(resourceMap)
	test.That(t, err, test.ShouldBeNil)
	resourceSubtype := registry.ResourceSubtypeLookup(generic.Subtype)
	resourceSubtype.RegisterRPCService(context.Background(), rpcServer, genericSvc)

	go rpcServer.Serve(listener1)
	defer rpcServer.Stop()

	t.Run("Failing client", func(t *testing.T) {
		cancelCtx, cancel := context.WithCancel(context.Background())
		cancel()
		_, err = viamgrpc.Dial(cancelCtx, listener1.Addr().String(), logger)
		test.That(t, err, test.ShouldNotBeNil)
		test.That(t, err.Error(), test.ShouldContainSubstring, "canceled")
	})

	t.Run("client tests for working generic", func(t *testing.T) {
		conn, err := viamgrpc.Dial(context.Background(), listener1.Addr().String(), logger)
		test.That(t, err, test.ShouldBeNil)
		workingGenericClient := generic.NewClientFromConn(context.Background(), conn, testGenericName, logger)

		resp, err := workingGenericClient.DoCommand(context.Background(), generic.TestCommand)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, resp["cmd"], test.ShouldEqual, generic.TestCommand["cmd"])
		test.That(t, resp["data"], test.ShouldEqual, generic.TestCommand["data"])

		test.That(t, utils.TryClose(context.Background(), workingGenericClient), test.ShouldBeNil)
		test.That(t, conn.Close(), test.ShouldBeNil)
	})

	t.Run("client tests for failing generic", func(t *testing.T) {
		conn, err := viamgrpc.Dial(context.Background(), listener1.Addr().String(), logger)
		test.That(t, err, test.ShouldBeNil)
		failingGenericClient := generic.NewClientFromConn(context.Background(), conn, failGenericName, logger)
		test.That(t, err, test.ShouldBeNil)

		_, err = failingGenericClient.DoCommand(context.Background(), generic.TestCommand)
		test.That(t, err, test.ShouldNotBeNil)

		test.That(t, utils.TryClose(context.Background(), failingGenericClient), test.ShouldBeNil)
		test.That(t, conn.Close(), test.ShouldBeNil)
	})

	t.Run("dialed client tests for working generic", func(t *testing.T) {
		conn, err := viamgrpc.Dial(context.Background(), listener1.Addr().String(), logger)
		test.That(t, err, test.ShouldBeNil)
		client := resourceSubtype.RPCClient(context.Background(), conn, testGenericName, logger)
		workingGenericDialedClient, ok := client.(generic.Generic)
		test.That(t, ok, test.ShouldBeTrue)

		resp, err := workingGenericDialedClient.DoCommand(context.Background(), generic.TestCommand)
		test.That(t, err, test.ShouldBeNil)
		test.That(t, resp["cmd"], test.ShouldEqual, generic.TestCommand["cmd"])
		test.That(t, resp["data"], test.ShouldEqual, generic.TestCommand["data"])

		test.That(t, conn.Close(), test.ShouldBeNil)
	})
}

func TestClientDialerOption(t *testing.T) {
	logger := golog.NewTestLogger(t)
	listener, err := net.Listen("tcp", "localhost:0")
	test.That(t, err, test.ShouldBeNil)
	gServer := grpc.NewServer()
	injectGeneric := &inject.Generic{}

	genericSvc, err := subtype.New(map[resource.Name]interface{}{generic.Named(testGenericName): injectGeneric})
	test.That(t, err, test.ShouldBeNil)
	componentpb.RegisterGenericServiceServer(gServer, generic.NewServer(genericSvc))

	go gServer.Serve(listener)
	defer gServer.Stop()

	td := &testutils.TrackingDialer{Dialer: rpc.NewCachedDialer()}
	ctx := rpc.ContextWithDialer(context.Background(), td)
	conn1, err := viamgrpc.Dial(ctx, listener.Addr().String(), logger)
	test.That(t, err, test.ShouldBeNil)
	client1 := generic.NewClientFromConn(ctx, conn1, testGenericName, logger)
	test.That(t, td.NewConnections, test.ShouldEqual, 3)
	conn2, err := viamgrpc.Dial(ctx, listener.Addr().String(), logger)
	test.That(t, err, test.ShouldBeNil)
	client2 := generic.NewClientFromConn(ctx, conn2, testGenericName, logger)
	test.That(t, td.NewConnections, test.ShouldEqual, 3)

	err = utils.TryClose(context.Background(), client1)
	test.That(t, err, test.ShouldBeNil)
	err = utils.TryClose(context.Background(), client2)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, conn1.Close(), test.ShouldBeNil)
	test.That(t, conn2.Close(), test.ShouldBeNil)
}
