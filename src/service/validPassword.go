package service

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"unicode"
	//"strings"
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

	if len(password) < value {
		return errors.New("Error the password don't have the minSize")
	}
	return nil
}

func (vp *ValidPassword) MinUppercase(ctx context.Context, value int, password string) error {

	upperCase := 0
	for _, letter := range password {
       if unicode.IsUpper(letter) {
			upperCase++
	   }   
    }

	if value > upperCase{
		return errors.New("Error the password don't have the minUpperCase")
	}
	return nil
}

func (vp *ValidPassword) MinLowercase(ctx context.Context, value int, password string) error {
	lowerCase := 0
	for _, letter := range password {
       if unicode.IsLower(letter) {
			lowerCase++
	   }   
    }

	if value > lowerCase{
		return errors.New("Error the password don't have the minLowerCase")
	}

	return nil
}

func (vp *ValidPassword) MinDigit(ctx context.Context, value int, password string) error {
	digit := 0
	for _, letter := range password {
       if unicode.IsDigit(letter) {
			digit++
	   }   
    }

	if value > digit{
		return errors.New("Error the password don't have the min digits")
	}

	fmt.Println(digit)
	return nil
}

func (vp *ValidPassword) MinSpecialChars(ctx context.Context, value int, password string) error  {

	symbol := 0
	for _, letter := range password {
       if !unicode.IsDigit(letter) && !unicode.IsLetter(letter) {
			symbol++
	   }   
    }

	if value > symbol{
		return errors.New("Error the password don't have the min special chars")
	}

	return nil
}

func (vp *ValidPassword) NoRepeted(ctx context.Context, value int, password string) error {
	var letterFirst rune
	for _, letter := range password {
       if letterFirst == letter {
			return errors.New("Error the password have Repeted letters")
	   }  
	   letterFirst = letter 
    }

	return nil
}
