package handler

import (
	"context"
	"encoding/json"
	"net/http"
	//"fmt"

	"github.com/Eli15x/password/src/models"
	"github.com/Eli15x/password/src/service"
	"github.com/gin-gonic/gin"
)


func VerifyPassword(c *gin.Context) {

	PasswordSetting := &models.PasswordSetting{}
	err := json.NewDecoder(c.Request.Body).Decode(PasswordSetting)

	if err != nil {
		c.String(400, "%s", err)
		return
	}

	if PasswordSetting.Password == "" {
		c.String(http.StatusBadRequest, "not password Passed")
		return
	}

	if PasswordSetting.Rules == nil {
		c.String(http.StatusBadRequest, "not rules Passed")
		return
	}

	var noMatch []string
	verify := true
	for _, rule := range PasswordSetting.Rules {
		if rule.Rule == "minSize"  {
			err := service.GetInstanceValidPassword().MinSize(context.Background(),rule.Value,PasswordSetting.Password)
			if err != nil {
				noMatch = append(noMatch,"minSize")
				verify = false
			}
		}
		if rule.Rule == "minUppercase"  {
			err = service.GetInstanceValidPassword().MinUppercase(context.Background(),rule.Value,PasswordSetting.Password)
			if err != nil {
				noMatch = append(noMatch,"minUppercase")
				verify = false
			}
		}
		if rule.Rule == "minLowercase"  {
			err = service.GetInstanceValidPassword().MinLowercase(context.Background(),rule.Value,PasswordSetting.Password)
			if err != nil {
				noMatch = append(noMatch,"minLowercase")
				verify = false
			}
		}
		if rule.Rule == "minDigit"  {
			err = service.GetInstanceValidPassword().MinDigit(context.Background(),rule.Value,PasswordSetting.Password)
			if err != nil {
				noMatch = append(noMatch,"minDigit")
				verify = false
			}
		}
		if rule.Rule == "minSpecialChars"  {
			err = service.GetInstanceValidPassword().MinSpecialChars(context.Background(),rule.Value,PasswordSetting.Password)
			if err != nil {
				noMatch = append(noMatch,"minSpecialChars")
				verify = false
			}
		}
		if rule.Rule == "noRepeted"  {
			err = service.GetInstanceValidPassword().NoRepeted(context.Background(),rule.Value,PasswordSetting.Password)
			if err != nil {
				noMatch = append(noMatch,"noRepeted")
				verify = false
			}
		}
	
	}

	if noMatch == nil {
		noMatch = append(noMatch, "")
	}

	result := map[string]interface{}{
		"verify": verify,
		"noMatch": noMatch,
	}
	//for and redirect to function service.

	c.JSON(http.StatusOK,  result)
}