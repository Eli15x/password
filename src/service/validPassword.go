package service

import (
	"context"
	//"errors"
	//"fmt"
	"sync"
	//"time"

	//"github.com/Eli15x/password/src/models"
	//"github.com/fatih/structs"
	//"github.com/labstack/gommon/log"
	//"go.mongodb.org/mongo-driver/bson"
)

var (
	instanceValidPassword CommandValidPassword
	onceValidPassword     sync.Once
)

type CommandValidPassword interface {
	MinSize(ctx context.Context, value int, password string) error
	MinUppercase(ctx context.Context, value int, password string) error
	MinLowercase(ctx context.Context, value int, password string) error
	MinDigit(ctx context.Context, value int, password string) error
	MinSpecialChars(ctx context.Context, value int, password string) error
	NoRepeted(ctx context.Context, value int, password string) error
}

type ValidPassword struct{}

func GetInstanceValidPassword() CommandValidPassword {
	onceValidPassword.Do(func() {
		instanceValidPassword = &ValidPassword{}
	})
	return instanceValidPassword
}

func (vp *ValidPassword) MinSize(ctx context.Context, value int, password string) error {

	return nil
}

func (vp *ValidPassword) MinUppercase(ctx context.Context, value int, password string) error {

	return nil
}

func (vp *ValidPassword) MinLowercase(ctx context.Context, value int, password string) error {
	return nil
}

func (vp *ValidPassword) MinDigit(ctx context.Context, value int, password string) error {

	return nil
}

func (vp *ValidPassword) MinSpecialChars(ctx context.Context, value int, password string) error  {

	return nil
}

func (vp *ValidPassword) NoRepeted(ctx context.Context, value int, password string) error {

	return nil
}
