package controller

import (
	"gostore/service"
	"gostore/utils/helper"
	"gostore/utils/response"
	"net/http"
	"strconv"
)

type NotificationController struct {
	service service.NotificationService
}

func NewNotificationController(s service.NotificationService) *NotificationController {
	return &NotificationController{service: s}
}

func (n *NotificationController) GetAllNotifications(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	userId := ctx.Value(helper.GOSTORE_USERID).(int)
	categoryId, _ := strconv.Atoi(r.URL.Query().Get("categoryId"))

	notifications, err := n.service.GetAllNotifications(ctx, userId, categoryId)
	if err != nil {
		response.BuildErorResponse(w, err)
		return
	}

	response.BuildSuccesResponse(w, notifications, nil, "")
}
