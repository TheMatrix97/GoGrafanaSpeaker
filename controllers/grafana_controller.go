package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thematrix97/go-grafana-speaker/services"
	"github.com/thematrix97/go-grafana-speaker/types"
)

func ProcessGrafanaEvent(ctx *gin.Context) (types.GrafanaEvent, *types.HTTPError) {
	var grafanaEvent types.GrafanaEvent
	var err error
	if err = ctx.ShouldBindJSON(&grafanaEvent); err == nil {
		err = services.PlayNotification(services.GetEnvVariable("NOTIFICATION_SOUND_FILE"))
		if err != nil {
			return grafanaEvent, &types.HTTPError{Code: http.StatusInternalServerError, Msg: "Cannot play sound file :("}
		}
	} else {
		return grafanaEvent, &types.HTTPError{Code: http.StatusBadRequest, Msg: "Bad Request"}
	}
	return grafanaEvent, nil
}
