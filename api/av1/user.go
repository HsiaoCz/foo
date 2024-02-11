package av1

import (
	"context"
	"strings"
	"time"

	"github.com/HsiaoCz/foo/etc"
	"github.com/HsiaoCz/foo/pb/pv1"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func UserSingup(c *fiber.Ctx) error {
	users := new(UserS)
	if err := c.BodyParser(users); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": err.Error(),
		})
	}
	if users.Password != users.RePassword {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": "please check the password and the re_password",
		})
	}
	if !strings.Contains(users.Email, "@") {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"Code":    fiber.ErrBadRequest.Code,
			"Message": "Please type a correct email",
		})
	}

	conn, err := grpc.Dial(etc.Conf.Addr.User, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.L().Error("grpc dial err:%v\n", zap.Error(err))
		return err
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	client := pv1.NewUserClient(conn)
	resp, err := client.UserSignup(ctx, &pv1.SignupRequest{
		Username: users.Username,
		Password: users.Password,
		Email:    users.Email,
		Phone:    users.Phone,
	})
	if err != nil {
		zap.L().Error("grpc user signup err:%v\n", zap.Error(err))
		return err
	}
	return c.Status(int(resp.GetCode())).JSON(fiber.Map{
		"Code":    resp.GetCode(),
		"Message": resp.GetMsg(),
	})
}

func UserLogin(c *fiber.Ctx) error {
	userL := new(UserL)
	if err := c.BodyParser(userL); err != nil {
		return err
	}
	conn, err := grpc.Dial(etc.Conf.Addr.User, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.L().Error("grpc dial err:%v\n", zap.Error(err))
		return err
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()
	client := pv1.NewUserClient(conn)
	resp, err := client.UserLogin(ctx, &pv1.LoginRequest{
		Username: userL.Username,
		Password: userL.Password,
		Email:    userL.Email,
		Phone:    userL.Phone,
	})
	if err != nil {
		zap.L().Error("grpc user login err:%v\n", zap.Error(err))
		return err
	}
	return c.Status(int(resp.GetCode())).JSON(fiber.Map{
		"Code":    resp.GetCode(),
		"Message": resp.GetMsg(),
	})
}

func UserModifed(c *fiber.Ctx) error {
	return nil
}
