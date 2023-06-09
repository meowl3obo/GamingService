package lolService

import (
	"net/http"
	"os"
	"regexp"
	"strings"

	. "gaming-service/model"
	. "gaming-service/model/lol"
	provider "gaming-service/provider/lol"
	transfer "gaming-service/transfer/lol"

	"github.com/gin-gonic/gin"
)

var (
	cacheItemsMap     = map[string]map[string]map[string]ItemResponse{}
	cacheSummonersMap = map[string]map[string]map[string]SummonerResponse{}
	cacheRolesMap     = map[string]map[string]map[string]RoleResponse{}
	cacheRoleDetails  = map[string]map[string]map[string]RoleDetailsResponse{}
)

func init() {
	cacheItemsMap = make(map[string]map[string]map[string]ItemResponse)
	cacheSummonersMap = make(map[string]map[string]map[string]SummonerResponse)
	cacheRolesMap = make(map[string]map[string]map[string]RoleResponse)
	cacheRoleDetails = make(map[string]map[string]map[string]RoleDetailsResponse)
}

// @Summary 取得全英雄
// @Tags LoL Assets
// @version 1.0
// @param version query string false "版本號"
// @param lang query string false "語系"
// @produce application/json
// @Router /api/lol/assets/roles [get]
func GetRoles(c *gin.Context) {
	version := handlerVersion(c.Query("version"))
	lang := handlerLang(c.Query("lang"))

	if checkRolesMapCache(version, lang) {
		c.JSON(http.StatusOK, setVersionResponse(version, cacheRolesMap[version][lang]))
		return
	}

	rolesDetails, statusCode, errObj := provider.GetRolesData(version, lang)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		rolesMap := transfer.ToRoleMap(rolesDetails)
		cacheRolesMap[version][lang] = rolesMap
		c.JSON(statusCode, setVersionResponse(version, rolesMap))
	}
}

// @Summary 取得單英雄詳細資料
// @Tags LoL Assets
// @version 1.0
// @param name path string true "英雄名稱"
// @param version query string false "版本號"
// @param lang query string false "語系"
// @produce application/json
// @Router /api/lol/assets/role/{name} [get]
func GetRole(c *gin.Context) {
	name := c.Param("name")
	version := handlerVersion(c.Query("version"))
	lang := handlerLang(c.Query("lang"))

	if checkRoleDetailsCache(version, lang, name) {
		c.JSON(http.StatusOK, setVersionResponse(version, cacheRoleDetails[version][lang][name]))
		return
	}

	roleDetails, statusCode, errObj := provider.GetRoleData(version, lang, name)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		versionReg, _ := regexp.Compile(`\.\d$`)
		mainVersion := versionReg.ReplaceAllString(version, "")
		skillDetails, _, _ := provider.GetRoleSkillDetails(mainVersion, strings.ToLower(name))

		roleDetails := transfer.ToRoleDetails(roleDetails, name, skillDetails)
		cacheRoleDetails[version][lang][name] = roleDetails
		c.JSON(statusCode, setVersionResponse(version, roleDetails))
	}
}

// @Summary 取得全部道具
// @Tags LoL Assets
// @version 1.0
// @param version query string false "版本號"
// @param lang query string false "語系"
// @produce application/json
// @Router /api/lol/assets/items [get]
func GetItems(c *gin.Context) {
	version := handlerVersion(c.Query("version"))
	lang := handlerLang(c.Query("lang"))

	if checkItemsMapCache(version, lang) {
		c.JSON(http.StatusOK, setVersionResponse(version, cacheItemsMap[version][lang]))
		return
	}

	items, statusCode, errObj := provider.GetItems(version, lang)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		itemsMap := transfer.ToItemMap(items)
		cacheItemsMap[version][lang] = itemsMap
		c.JSON(statusCode, setVersionResponse(version, itemsMap))
	}
}

// @Summary 取得全部召喚師技能
// @Tags LoL Assets
// @version 1.0
// @param version query string false "版本號"
// @param lang query string false "語系"
// @produce application/json
// @Router /api/lol/assets/summoners [get]
func GetSummoners(c *gin.Context) {
	version := handlerVersion(c.Query("version"))
	lang := handlerLang(c.Query("lang"))

	if checkSummonersMapCache(version, lang) {
		c.JSON(http.StatusOK, setVersionResponse(version, cacheSummonersMap[version][lang]))
		return
	}

	summonersDetails, statusCode, errObj := provider.GetSummoners(version, lang)

	if statusCode != 200 {
		c.JSON(statusCode, errObj)
	} else {
		summonersMap := transfer.ToSummonerMap(summonersDetails)
		cacheSummonersMap[version][lang] = summonersMap
		c.JSON(statusCode, setVersionResponse(version, summonersMap))
	}
}

func setVersionResponse[T any](version string, data T) VersionResponse[T] {
	response := VersionResponse[T]{
		Version: version,
		Data:    data,
	}
	return response
}

func handlerVersion(version string) string {
	if version == "" {
		return os.Getenv("LOL_VERSION")
	}
	return version
}

func handlerLang(lang string) string {
	if lang == "" {
		return "en_US"
	}
	return lang
}

func checkRolesMapCache(version string, lang string) bool {
	if cacheRolesMap[version] == nil {
		cacheRolesMap[version] = map[string]map[string]RoleResponse{}
	}
	if cacheRolesMap[version][lang] == nil {
		cacheRolesMap[version][lang] = map[string]RoleResponse{}
	}
	if len(cacheRolesMap[version][lang]) > 0 {
		return true
	}
	return false
}

func checkRoleDetailsCache(version string, lang string, name string) bool {
	if cacheRoleDetails[version] == nil {
		cacheRoleDetails[version] = map[string]map[string]RoleDetailsResponse{}
	}
	if cacheRoleDetails[version][lang] == nil {
		cacheRoleDetails[version][lang] = map[string]RoleDetailsResponse{}
	}
	_, found := cacheRoleDetails[version][lang][name]
	return found
}

func checkItemsMapCache(version string, lang string) bool {
	if cacheItemsMap[version] == nil {
		cacheItemsMap[version] = map[string]map[string]ItemResponse{}
	}
	if cacheItemsMap[version][lang] == nil {
		cacheItemsMap[version][lang] = map[string]ItemResponse{}
	}
	if len(cacheItemsMap[version][lang]) > 0 {
		return true
	}
	return false
}

func checkSummonersMapCache(version string, lang string) bool {
	if cacheSummonersMap[version] == nil {
		cacheSummonersMap[version] = map[string]map[string]SummonerResponse{}
	}
	if cacheSummonersMap[version][lang] == nil {
		cacheSummonersMap[version][lang] = map[string]SummonerResponse{}
	}
	if len(cacheSummonersMap[version][lang]) > 0 {
		return true
	}
	return false
}
