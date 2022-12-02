package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Eli15x/password/src/models"
	"github.com/Eli15x/password/src/service"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
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

	if PasswordSetting.Rules == "" {
		c.String(http.StatusBadRequest, "not rules Passed")
		return
	}

	c.JSON(http.StatusOK, result)
}