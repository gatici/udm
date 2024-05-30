// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * Nudm_UECM
 *
 * Nudm Context Management Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package uecontextmanagement

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/omec-project/openapi"
	"github.com/omec-project/openapi/models"
	"github.com/omec-project/udm/logger"
	"github.com/omec-project/udm/producer"
	"github.com/omec-project/util/httpwrapper"
)

// DeregistrationSmfRegistrations - delete an SMF registration
func HTTPDeregistrationSmfRegistrations(c *gin.Context) {
	req := httpwrapper.NewRequest(c.Request, nil)
	req.Params["ueId"] = c.Params.ByName("ueId")
	req.Params["pduSessionId"] = c.Params.ByName("pduSessionId")

	rsp := producer.HandleDeregistrationSmfRegistrations(req)

	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
	if err != nil {
		logger.UecmLog.Errorln(err)
		problemDetails := models.ProblemDetails{
			Status: http.StatusInternalServerError,
			Cause:  "SYSTEM_FAILURE",
			Detail: err.Error(),
		}
		c.JSON(http.StatusInternalServerError, problemDetails)
	} else {
		c.Data(rsp.Status, "application/json", responseBody)
	}
}
