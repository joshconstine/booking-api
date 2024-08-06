package controllers

import (
	"booking-api/data/request"
	"booking-api/services"
	rentals "booking-api/view/rentals"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type RentalRoomController struct {
	rentalRoomService services.RentalRoomService
	roomTypeService   services.RoomTypeService
	bedTypeService    services.BedTypeService
}

func NewRentalRoomController(rentalRoomService services.RentalRoomService, roomTypeService services.RoomTypeService, bedTypeService services.BedTypeService) *RentalRoomController {
	return &RentalRoomController{rentalRoomService: rentalRoomService, roomTypeService: roomTypeService, bedTypeService: bedTypeService}
}

//func (controller *RentalRoomController) FindAll(ctx *gin.Context) {
//	rentalRooms := controller.rentalRoomService.FindAll()
//
//	webResponse := response.Response{
//		Code:   200,
//		Status: "Ok",
//		Data:   rentalRooms,
//	}
//
//	ctx.Header("Content-Type", "application/json")
//	ctx.JSON(http.StatusOK, webResponse)
//
//}
//
//func (controller *RentalRoomController) FindById(ctx *gin.Context) {
//	rentalRoomId := ctx.Param("rentalRoomId")
//	id, _ := strconv.Atoi(rentalRoomId)
//
//	rentalRoom := controller.rentalRoomService.FindById(uint(id))
//
//	webResponse := response.Response{
//		Code:   200,
//		Status: "Ok",
//		Data:   rentalRoom,
//	}
//
//	ctx.Header("Content-Type", "application/json")
//	ctx.JSON(http.StatusOK, webResponse)
//}

//func (controller *RentalRoomController) Create(ctx *gin.Context) {
//	var rentalRoomCreateRequest request.RentalRoomCreateRequest
//
//	err := ctx.ShouldBindJSON(&rentalRoomCreateRequest)
//
//	if err != nil {
//		webResponse := response.Response{
//			Code:   400,
//			Status: "Bad Request",
//			Data:   err.Error(),
//		}
//
//		ctx.Header("Content-Type", "application/json")
//		ctx.JSON(http.StatusBadRequest, webResponse)
//		return
//	}
//
//	rentalRoom, err := controller.rentalRoomService.Create(rentalRoomCreateRequest)
//
//	if err != nil {
//		webResponse := response.Response{
//			Code:   500,
//			Status: "Internal Server Error",
//			Data:   err.Error(),
//		}
//
//		ctx.Header("Content-Type", "application/json")
//		ctx.JSON(http.StatusBadRequest, webResponse)
//		return
//
//	}
//	webResponse := response.Response{
//		Code:   http.StatusCreated,
//		Status: "Created",
//		Data:   rentalRoom,
//	}
//
//	ctx.Header("Content-Type", "application/json")
//	ctx.JSON(http.StatusOK, webResponse)
//}

func (controller *RentalRoomController) Update(w http.ResponseWriter, r *http.Request) error {

	rentalId := chi.URLParam(r, "rentalId")
	room := chi.URLParam(r, "roomId")
	rentalIdInt, _ := strconv.Atoi(rentalId)
	roomInt, _ := strconv.Atoi(room)

	var updateParams request.UpdateRentalRoomRequest
	rentalRoomTypeID, _ := strconv.Atoi(r.FormValue("room_type_id"))
	updateParams.ID = uint(roomInt)
	updateParams.RentalID = uint(rentalIdInt)
	updateParams.Name = r.FormValue("name")
	updateParams.Description = r.FormValue("description")
	updateParams.Floor, _ = strconv.Atoi(r.FormValue("floor"))
	updateParams.RentalRoomTypeID = uint(rentalRoomTypeID)
	updateParams.Photos = []int{}
	updateParams.Beds = []int{}

	_, err := controller.rentalRoomService.Update(updateParams)

	if err != nil {
		return err
	}
	params := request.CreateRentalStep2Params{}
	errors := request.CreateRentalStep2Errors{}
	rentalRooms := controller.rentalRoomService.FindByRentalId(uint(rentalIdInt))
	params.Rooms = rentalRooms
	params.RentalID = uint(rentalIdInt)
	roomTypes := controller.roomTypeService.FindAll()
	bedTypes := controller.bedTypeService.FindAll()

	return rentals.RentalBedroomsForm(params, updateParams, errors, roomTypes, bedTypes).Render(r.Context(), w)
}

func (controller *RentalRoomController) Create(w http.ResponseWriter, r *http.Request) error {

	rentalId := chi.URLParam(r, "rentalId")
	rentalIdInt, _ := strconv.Atoi(rentalId)

	var createRequest request.RentalRoomCreateRequest
	rentalRoomTypeID, _ := strconv.Atoi(r.FormValue("room_type_id"))
	createRequest.RentalID = uint(rentalIdInt)
	createRequest.Name = r.FormValue("name")
	createRequest.Description = r.FormValue("description")
	createRequest.Floor, _ = strconv.Atoi(r.FormValue("floor"))
	createRequest.RentalRoomTypeID = uint(rentalRoomTypeID)
	createRequest.Photos = []int{}
	createRequest.Beds = []int{}

	result, err := controller.rentalRoomService.Create(createRequest)

	if err != nil {
		return err
	}

	updateParams := request.UpdateRentalRoomRequest{
		ID:          result.ID,
		Name:        result.Name,
		Description: result.Description,
		Floor:       result.Floor,
		RentalID:    uint(rentalIdInt),
	}

	for _, beds := range result.Beds {
		updateParams.Beds = append(updateParams.Beds, int(beds.ID))
	}

	params := request.CreateRentalStep2Params{}
	errors := request.CreateRentalStep2Errors{}
	rentalRooms := controller.rentalRoomService.FindByRentalId(uint(rentalIdInt))
	params.Rooms = rentalRooms
	params.RentalID = uint(rentalIdInt)
	roomTypes := controller.roomTypeService.FindAll()
	bedTypes := controller.bedTypeService.FindAll()

	return rentals.RentalBedroomsForm(params, updateParams, errors, roomTypes, bedTypes).Render(r.Context(), w)
}
func (controller *RentalRoomController) Delete(w http.ResponseWriter, r *http.Request) error {

	rentalId := chi.URLParam(r, "rentalId")
	room := chi.URLParam(r, "roomId")
	rentalIdInt, _ := strconv.Atoi(rentalId)
	roomInt, _ := strconv.Atoi(room)

	err := controller.rentalRoomService.Delete(uint(roomInt))

	if err != nil {
		return err
	}

	params := request.CreateRentalStep2Params{}
	errors := request.CreateRentalStep2Errors{}
	rentalRooms := controller.rentalRoomService.FindByRentalId(uint(rentalIdInt))
	params.Rooms = rentalRooms
	params.RentalID = uint(rentalIdInt)
	roomTypes := controller.roomTypeService.FindAll()
	bedTypes := controller.bedTypeService.FindAll()

	return rentals.RentalBedroomsForm(params, request.UpdateRentalRoomRequest{}, errors, roomTypes, bedTypes).Render(r.Context(), w)
}
