package service

import (
	"context"
	
	testifyMock "github.com/stretchr/testify/mock"
)

type ValidPasswordMock struct {
	testifyMock.Mock
}

func (vp *ValidPasswordMock) MinSize(ctx context.Context, value int, password string) error {
	args := vp.Called(ctx,value,password)
	return args.Error(0)
}

func (vp *ValidPasswordMock) MinUppercase(ctx context.Context, value int, password string) error {
	args := vp.Called(ctx,value,password)
	return args.Error(0)
}

func (vp *ValidPasswordMock) MinLowercase(ctx context.Context, value int, password string) error {
	args := vp.Called(ctx,value,password)
	return args.Error(0)
}

func (vp *ValidPasswordMock) MinDigit(ctx context.Context, value int, password string) error {
	args := vp.Called(ctx,value,password)
	return args.Error(0)
}

func (vp *ValidPasswordMock) MinSpecialChars(ctx context.Context, value int, password string) error {
	args := vp.Called(ctx,value,password)
	return args.Error(0)
}

func (vp *ValidPasswordMock) NoRepeted(ctx context.Context, value int, password string) error {
	args := vp.Called(ctx,value,password)
	return args.Error(0)
}