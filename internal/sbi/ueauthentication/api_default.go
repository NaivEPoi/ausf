/*
 * AUSF API
 *
 * OpenAPI specification for AUSF
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ueauthentication

import (
	"net/http"

	"github.com/gin-gonic/gin"

	ausf_context "github.com/free5gc/ausf/internal/context"
	"github.com/free5gc/ausf/internal/logger"
	"github.com/free5gc/ausf/internal/sbi/producer"
	"github.com/free5gc/openapi"
	"github.com/free5gc/openapi/models"
	"github.com/free5gc/util/httpwrapper"
)

// HTTPEapAuthMethod -
func HTTPEapAuthMethod(ctx *gin.Context) {
	scopes := []string{"nausf-auth"}
	_, oauth_err := openapi.CheckOAuth(ctx.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && ausf_context.GetSelf().OAuth {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}

	var eapSessionReq models.EapSession

	requestBody, err := ctx.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.Auth5gAkaComfirmLog.Errorf("Get Request Body error: %+v", err)
		ctx.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	err = openapi.Deserialize(&eapSessionReq, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.Auth5gAkaComfirmLog.Errorln(problemDetail)
		ctx.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := httpwrapper.NewRequest(ctx.Request, eapSessionReq)
	req.Params["authCtxId"] = ctx.Param("authCtxId")

	rsp := producer.HandleEapAuthComfirmRequest(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.Auth5gAkaComfirmLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		ctx.Data(rsp.Status, "application/json", responseBody)
	}
}

// HTTPUeAuthenticationsAuthCtxID5gAkaConfirmationPut -
func HTTPUeAuthenticationsAuthCtxID5gAkaConfirmationPut(ctx *gin.Context) {
	scopes := []string{"nausf-auth"}
	_, oauth_err := openapi.CheckOAuth(ctx.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && ausf_context.GetSelf().OAuth {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}

	var confirmationData models.ConfirmationData

	requestBody, err := ctx.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.Auth5gAkaComfirmLog.Errorf("Get Request Body error: %+v", err)
		ctx.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	err = openapi.Deserialize(&confirmationData, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.Auth5gAkaComfirmLog.Errorln(problemDetail)
		ctx.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := httpwrapper.NewRequest(ctx.Request, confirmationData)
	req.Params["authCtxId"] = ctx.Param("authCtxId")

	rsp := producer.HandleAuth5gAkaComfirmRequest(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.Auth5gAkaComfirmLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		ctx.Data(rsp.Status, "application/json", responseBody)
	}
}

// HTTPUeAuthenticationsPost -
func HTTPUeAuthenticationsPost(ctx *gin.Context) {
	scopes := []string{"nausf-auth"}
	_, oauth_err := openapi.CheckOAuth(ctx.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && ausf_context.GetSelf().OAuth {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}

	var authInfo models.AuthenticationInfo

	requestBody, err := ctx.GetRawData()
	if err != nil {
		problemDetail := models.ProblemDetails{
			Title:  "System failure",
			Status: http.StatusInternalServerError,
			Detail: err.Error(),
			Cause:  "SYSTEM_FAILURE",
		}
		logger.UeAuthPostLog.Errorf("Get Request Body error: %+v", err)
		ctx.JSON(http.StatusInternalServerError, problemDetail)
		return
	}

	err = openapi.Deserialize(&authInfo, requestBody, "application/json")
	if err != nil {
		problemDetail := "[Request Body] " + err.Error()
		rsp := models.ProblemDetails{
			Title:  "Malformed request syntax",
			Status: http.StatusBadRequest,
			Detail: problemDetail,
		}
		logger.UeAuthPostLog.Errorln(problemDetail)
		ctx.JSON(http.StatusBadRequest, rsp)
		return
	}

	req := httpwrapper.NewRequest(ctx.Request, authInfo)

	rsp := producer.HandleUeAuthPostRequest(req)

	for key, value := range rsp.Header {
		ctx.Header(key, value[0])
	}
	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.UeAuthPostLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		ctx.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		ctx.Data(rsp.Status, "application/json", responseBody)
	}
}
