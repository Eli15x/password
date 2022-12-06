package service


import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	testifyMock "github.com/stretchr/testify/mock"
)

var (
	ValidPasswordTest *ValidPasswordMock
	ctx     context.Context
)

func Test_MinSize_StatusOk(t *testing.T) {
	ValidPasswordTest = &ValidPasswordMock{}
	ValidPasswordTest.On("MinSize",  testifyMock.Anything, testifyMock.Anything, testifyMock.Anything).Return(nil)
	err := ValidPasswordTest.MinSize(ctx, 5 , "teste")

	assert.Equal(t, err, nil)
}

func Test_MinUppercase_StatusOk(t *testing.T) {
	ValidPasswordTest = &ValidPasswordMock{}
	ValidPasswordTest.On("MinUppercase",  testifyMock.Anything, testifyMock.Anything, testifyMock.Anything).Return(nil)

	err := ValidPasswordTest.MinUppercase(ctx, 2 , "TesTe")

	assert.Equal(t, err, nil)
}

func Test_MinDigit_StatusOk(t *testing.T) {
	ValidPasswordTest = &ValidPasswordMock{}
	ValidPasswordTest.On("MinDigit",  testifyMock.Anything, testifyMock.Anything, testifyMock.Anything).Return(nil)

	err := ValidPasswordTest.MinDigit(ctx, 1 , "TesTe1")

	assert.Equal(t, err, nil)
}

func Test_MinSpecialChars_StatusOk(t *testing.T) {
	ValidPasswordTest = &ValidPasswordMock{}
	ValidPasswordTest.On("MinSpecialChars",  testifyMock.Anything, testifyMock.Anything, testifyMock.Anything).Return(nil)

	err := ValidPasswordTest.MinSpecialChars(ctx, 1 , "TesTe@")

	assert.Equal(t, err, nil)
}

func Test_Norepeted_StatusOk(t *testing.T) {
	ValidPasswordTest = &ValidPasswordMock{}
	ValidPasswordTest.On("NoRepeted",  testifyMock.Anything, testifyMock.Anything, testifyMock.Anything).Return(nil)

	err := ValidPasswordTest.NoRepeted(ctx, 1 , "TesTe")

	assert.Equal(t, err, nil)
}