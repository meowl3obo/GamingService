package lolService

import (
	"os"

	provider "gaming-service/provider/lol"
	transfer "gaming-service/transfer/lol"

	"github.com/gin-gonic/gin"
)

func GetRoles(c *gin.Context) {
	version := c.Query("version")
	lang := c.Query("lang")
	if version == "" {
		version = os.Getenv("LOL_VERSION")
	}
	if lang == "" {
		lang = "en_US"
	}

	rolesDetails, statusCode, errObj := provider.GetRolesData(version, lang)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, transfer.ToRoleMap(rolesDetails))
	}
}

func GetRole(c *gin.Context) {
	name := c.Param("name")
	version := c.Query("version")
	lang := c.Query("lang")
	if version == "" {
		version = os.Getenv("LOL_VERSION")
	}
	if lang == "" {
		lang = "en_US"
	}

	roleDetails, statusCode, errObj := provider.GetRoleData(version, lang, name)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, transfer.ToRoleDetails(roleDetails, name))
	}
}

func GetItems(c *gin.Context) {
	version := c.Query("version")
	lang := c.Query("lang")
	if version == "" {
		version = os.Getenv("LOL_VERSION")
	}
	if lang == "" {
		lang = "en_US"
	}

	items, statusCode, errObj := provider.GetItems(version, lang)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, transfer.ToItemMap(items))
	}
}

func GetSummoners(c *gin.Context) {
	version := c.Query("version")
	lang := c.Query("lang")
	if version == "" {
		version = os.Getenv("LOL_VERSION")
	}
	if lang == "" {
		lang = "en_US"
	}

	summonersDetails, statusCode, errObj := provider.GetSummoners(version, lang)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		c.JSON(statusCode, transfer.ToSummonerMap(summonersDetails))
	}
}
