package motor_test

import (
	"context"
	"errors"
	"testing"

	"go.viam.com/test"
	"google.golang.org/protobuf/types/known/structpb"

	"go.viam.com/core/component/motor"
	"go.viam.com/core/config"
	pb "go.viam.com/core/proto/api/component/v1"
	"go.viam.com/core/resource"
	"go.viam.com/core/subtype"
	"go.viam.com/core/testutils/inject"
)

func newServer() (pb.MotorServiceServer, *inject.Motor, *inject.Motor, error) {
	injectMotor1 := &inject.Motor{}
	injectMotor2 := &inject.Motor{}

	resourceMap := map[resource.Name]interface{}{
		motor.Named("workingMotor"): injectMotor1,
		motor.Named("failingMotor"): injectMotor2,
		motor.Named("notAMotor"):    "not a motor",
	}

	injectSvc, err := subtype.New((resourceMap))
	if err != nil {
		return nil, nil, nil, err
	}
	return motor.NewServer(injectSvc), injectMotor1, injectMotor2, nil
}

func TestGetPIDConfig(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	req := pb.MotorServiceGetPIDConfigRequest{Name: "notAMotor"}
	resp, err := motorServer.GetPIDConfig(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.PIDFunc = func() motor.PID {
		return nil
	}
	req = pb.MotorServiceGetPIDConfigRequest{Name: "failingMotor"}
	resp, err = motorServer.GetPIDConfig(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.PIDFunc = func() motor.PID {
		return &motor.BasicPID{}
	}
	req = pb.MotorServiceGetPIDConfigRequest{Name: "failingMotor"}
	resp, err = motorServer.GetPIDConfig(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	var expectedConf *structpb.Struct
	workingMotor.PIDFunc = func() motor.PID {
		attrMap := config.AttributeMap{
			"someKey": 42,
		}
		pid, err := motor.CreatePID(&motor.PIDConfig{
			Type:       "basic",
			Attributes: attrMap,
		})
		if err != nil {
			panic("fix PID creation in TestGetPIDConfig")
		}
		attrStruct, err := structpb.NewStruct(attrMap)
		expectedConf = attrStruct
		if err != nil {
			panic("fix attribute creation in TestGetPIDConfig")
		}
		return pid
	}
	req = pb.MotorServiceGetPIDConfigRequest{Name: "workingMotor"}
	resp, err = motorServer.GetPIDConfig(context.Background(), &req)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, resp.GetPidConfig(), test.ShouldResemble, expectedConf)

}

func TestSetPIDConfig(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	req := pb.MotorServiceSetPIDConfigRequest{Name: "notAMotor"}
	resp, err := motorServer.SetPIDConfig(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.PIDFunc = func() motor.PID {
		return nil
	}
	req = pb.MotorServiceSetPIDConfigRequest{Name: "failingMotor"}
	resp, err = motorServer.SetPIDConfig(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	attrMap := config.AttributeMap{
		"Kp": 42.0,
		"Kd": 43.0,
	}
	pid, err := motor.CreatePID(&motor.PIDConfig{
		Type:       "basic",
		Attributes: attrMap,
	})
	if err != nil {
		panic("fix PID creation in TestGetPIDConfig")
	}

	workingMotor.PIDFunc = func() motor.PID {
		return pid
	}
	newAttrMap, err := structpb.NewStruct(config.AttributeMap{
		"Kp": 43.0,
	})
	test.That(t, err, test.ShouldBeNil)

	req = pb.MotorServiceSetPIDConfigRequest{
		Name:      "workingMotor",
		PidConfig: newAttrMap,
	}
	resp, err = motorServer.SetPIDConfig(context.Background(), &req)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, resp, test.ShouldNotBeNil)
	cfg, err := workingMotor.PID().Config(context.Background())
	test.That(t, err, test.ShouldBeNil)
	test.That(t, cfg.Attributes.Float64("Kp", 0.0), test.ShouldEqual, 43.0)
	test.That(t, cfg.Attributes.Float64("Kd", 0.0), test.ShouldEqual, 43.0)

}

func TestSetPower(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceSetPowerRequest{Name: "notAMotor"}
	resp, err := motorServer.SetPower(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.SetPowerFunc = func(ctx context.Context, powerPct float64) error {
		return errors.New("set power failed")
	}
	req = pb.MotorServiceSetPowerRequest{Name: "failingMotor", PowerPct: 0.5}
	resp, err = motorServer.SetPower(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.SetPowerFunc = func(ctx context.Context, powerPct float64) error {
		return nil
	}
	req = pb.MotorServiceSetPowerRequest{Name: "workingMotor", PowerPct: 0.5}
	resp, err = motorServer.SetPower(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldBeNil)
}

func TestGo(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceGoRequest{Name: "notAMotor"}
	resp, err := motorServer.Go(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.GoFunc = func(ctx context.Context, powerPct float64) error {
		return errors.New("motor go failed")
	}
	req = pb.MotorServiceGoRequest{Name: "failingMotor", PowerPct: 0.5}
	resp, err = motorServer.Go(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.GoFunc = func(ctx context.Context, powerPct float64) error {
		return nil
	}
	req = pb.MotorServiceGoRequest{Name: "workingMotor", PowerPct: 0.5}
	resp, err = motorServer.Go(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldBeNil)
}

func TestGoFor(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceGoForRequest{Name: "notAMotor"}
	resp, err := motorServer.GoFor(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.GoForFunc = func(ctx context.Context, rpm, rotations float64) error {
		return errors.New("go for failed")
	}
	req = pb.MotorServiceGoForRequest{Name: "failingMotor", Rpm: 42.0, Revolutions: 42.1}
	resp, err = motorServer.GoFor(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.GoForFunc = func(ctx context.Context, rpm, rotations float64) error {
		return nil
	}
	req = pb.MotorServiceGoForRequest{Name: "workingMotor", Rpm: 42.0, Revolutions: 42.1}
	resp, err = motorServer.GoFor(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldBeNil)
}

func TestPosition(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServicePositionRequest{Name: "notAMotor"}
	resp, err := motorServer.Position(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.PositionFunc = func(ctx context.Context) (float64, error) {
		return 0, errors.New("position unavailable")
	}
	req = pb.MotorServicePositionRequest{Name: "failingMotor"}
	resp, err = motorServer.Position(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.PositionFunc = func(ctx context.Context) (float64, error) {
		return 42.0, nil
	}
	req = pb.MotorServicePositionRequest{Name: "workingMotor"}
	resp, err = motorServer.Position(context.Background(), &req)
	test.That(t, resp.GetPosition(), test.ShouldEqual, 42.0)
	test.That(t, err, test.ShouldBeNil)
}

func TestPositionSupported(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServicePositionSupportedRequest{Name: "notAMotor"}
	resp, err := motorServer.PositionSupported(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.PositionSupportedFunc = func(ctx context.Context) (bool, error) {
		return false, errors.New("unable to determined if pos supported")
	}
	req = pb.MotorServicePositionSupportedRequest{Name: "failingMotor"}
	resp, err = motorServer.PositionSupported(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.PositionSupportedFunc = func(ctx context.Context) (bool, error) {
		return true, nil
	}
	req = pb.MotorServicePositionSupportedRequest{Name: "workingMotor"}
	resp, err = motorServer.PositionSupported(context.Background(), &req)
	test.That(t, resp.GetSupported(), test.ShouldBeTrue)
	test.That(t, err, test.ShouldBeNil)
}

func TestStop(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceStopRequest{Name: "notAMotor"}
	resp, err := motorServer.Stop(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.StopFunc = func(ctx context.Context) error {
		return errors.New("stop failed")
	}
	req = pb.MotorServiceStopRequest{Name: "failingMotor"}
	resp, err = motorServer.Stop(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.StopFunc = func(ctx context.Context) error {
		return nil
	}
	req = pb.MotorServiceStopRequest{Name: "workingMotor"}
	resp, err = motorServer.Stop(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldBeNil)
}

func TestIsOn(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceIsOnRequest{Name: "notAMotor"}
	resp, err := motorServer.IsOn(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.IsOnFunc = func(ctx context.Context) (bool, error) {
		return false, errors.New("could not determine if motor is on")
	}
	req = pb.MotorServiceIsOnRequest{Name: "failingMotor"}
	resp, err = motorServer.IsOn(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.IsOnFunc = func(ctx context.Context) (bool, error) {
		return true, nil
	}
	req = pb.MotorServiceIsOnRequest{Name: "workingMotor"}
	resp, err = motorServer.IsOn(context.Background(), &req)
	test.That(t, resp.GetIsOn(), test.ShouldBeTrue)
	test.That(t, err, test.ShouldBeNil)
}

func TestGoTo(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceGoToRequest{Name: "notAMotor"}
	resp, err := motorServer.GoTo(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.GoToFunc = func(ctx context.Context, rpm, position float64) error {
		return errors.New("go to failed")
	}
	req = pb.MotorServiceGoToRequest{Name: "failingMotor", Rpm: 20.0, Position: 2.5}
	resp, err = motorServer.GoTo(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.GoToFunc = func(ctx context.Context, rpm, position float64) error {
		return nil
	}
	req = pb.MotorServiceGoToRequest{Name: "workingMotor", Rpm: 20.0, Position: 2.5}
	resp, err = motorServer.GoTo(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldBeNil)
}

func TestResetZeroPosition(t *testing.T) {
	motorServer, workingMotor, failingMotor, _ := newServer()

	// fails on a bad motor
	req := pb.MotorServiceResetZeroPositionRequest{Name: "notAMotor"}
	resp, err := motorServer.ResetZeroPosition(context.Background(), &req)
	test.That(t, resp, test.ShouldBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	failingMotor.ResetZeroPositionFunc = func(ctx context.Context, offset float64) error {
		return errors.New("set to zero failed")
	}
	req = pb.MotorServiceResetZeroPositionRequest{Name: "failingMotor", Offset: 1.1}
	resp, err = motorServer.ResetZeroPosition(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldNotBeNil)

	workingMotor.ResetZeroPositionFunc = func(ctx context.Context, offset float64) error {
		return nil
	}
	req = pb.MotorServiceResetZeroPositionRequest{Name: "workingMotor", Offset: 1.1}
	resp, err = motorServer.ResetZeroPosition(context.Background(), &req)
	test.That(t, resp, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldBeNil)
}