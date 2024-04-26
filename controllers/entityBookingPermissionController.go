package controllers

import (
	"booking-api/constants"
	"booking-api/data/request"
	"booking-api/services"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type EntityBookingPermissionController struct {
	entityBookingPermissionService services.EntityBookingPermissionService
}

func NewEntityBookingPermissionController(entityBookingPermissionService services.EntityBookingPermissionService) *EntityBookingPermissionController {
	return &EntityBookingPermissionController{entityBookingPermissionService: entityBookingPermissionService}
}

func (e EntityBookingPermissionController) Update(w http.ResponseWriter, r *http.Request) error {
	var entityBookingPermission request.UpdateEntityBookingPermissionRequest

	ebpridINT, _ := strconv.Atoi(chi.URLParam(r, "entityBookingPermissionID"))

	entityBookingPermission.EntityBookingPermissionID = uint(ebpridINT)
	entityBookingPermission.InquiryStatusID = constants.INQUIRY_STATUS_DECLINED_ID

	err := e.entityBookingPermissionService.Update(entityBookingPermission)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
