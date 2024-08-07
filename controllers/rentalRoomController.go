package controllers

import (
	"booking-api/constants"
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	rentals "booking-api/view/rentals"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type RentalRoomController struct {
	rentalRoomService  services.RentalRoomService
	roomTypeService    services.RoomTypeService
	bedTypeService     services.BedTypeService
	entityPhotoService services.EntityPhotoService
}

func NewRentalRoomController(rentalRoomService services.RentalRoomService, roomTypeService services.RoomTypeService, bedTypeService services.BedTypeService, entityPhotoService services.EntityPhotoService) *RentalRoomController {
	return &RentalRoomController{rentalRoomService: rentalRoomService, roomTypeService: roomTypeService, bedTypeService: bedTypeService, entityPhotoService: entityPhotoService}
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

func getBedsFromRequest(r *http.Request) []response.BedResponse {
	var Beds []response.BedResponse
	keyLead := "bed_type_id_"
	charLen := len(keyLead)
	for key, value := range r.Form {
		if len(key) > charLen && key[:charLen] == "bed_type_id_" {
			bedID := key[charLen:]
			bedIDInt, _ := strconv.Atoi(bedID)
			bedType := value[0]
			bedTypeIDInt, _ := strconv.Atoi(bedType)
			//Beds = append(Beds, response.BedResponse{ID: 0, BedTypeID: uint(bedIDInt)})
			Beds = append(Beds, response.BedResponse{ID: uint(bedIDInt), BedTypeID: uint(bedTypeIDInt)})
		}

	}
	return Beds
}
func setupParamsFromForm(r *http.Request) request.UpdateRentalRoomRequest {
	var updateParams request.UpdateRentalRoomRequest
	rentalId := chi.URLParam(r, "rentalId")
	room := chi.URLParam(r, "roomId")
	rentalIdInt, _ := strconv.Atoi(rentalId)
	roomInt, _ := strconv.Atoi(room)

	rentalRoomTypeID, _ := strconv.Atoi(r.FormValue("room_type_id"))
	updateParams.ID = uint(roomInt)
	updateParams.RentalID = uint(rentalIdInt)
	updateParams.Name = r.FormValue("name")
	updateParams.Description = r.FormValue("description")
	updateParams.Floor, _ = strconv.Atoi(r.FormValue("floor"))
	updateParams.RentalRoomTypeID = uint(rentalRoomTypeID)
	updateParams.Photos = []int{}

	updateParams.Beds = getBedsFromRequest(r)
	return updateParams
}
func (controller *RentalRoomController) Update(w http.ResponseWriter, r *http.Request) error {

	rentalId := chi.URLParam(r, "rentalId")
	rentalIdInt, _ := strconv.Atoi(rentalId)

	updateParams := setupParamsFromForm(r)

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

	photos := controller.entityPhotoService.FindAllEntityPhotosForEntity(constants.RENTAL_ENTITY, uint(rentalIdInt))
	return rentals.RentalBedroomsForm(params, updateParams, errors, roomTypes, bedTypes, photos).Render(r.Context(), w)
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
		ID:               result.ID,
		Name:             result.Name,
		Description:      result.Description,
		Floor:            result.Floor,
		RentalID:         uint(rentalIdInt),
		Beds:             result.Beds,
		RentalRoomTypeID: createRequest.RentalRoomTypeID,
	}

	params := request.CreateRentalStep2Params{}
	errors := request.CreateRentalStep2Errors{}
	rentalRooms := controller.rentalRoomService.FindByRentalId(uint(rentalIdInt))
	params.Rooms = rentalRooms
	params.RentalID = uint(rentalIdInt)
	roomTypes := controller.roomTypeService.FindAll()
	bedTypes := controller.bedTypeService.FindAll()

	photos := controller.entityPhotoService.FindAllEntityPhotosForEntity(constants.RENTAL_ENTITY, uint(rentalIdInt))
	return rentals.RentalBedroomsForm(params, updateParams, errors, roomTypes, bedTypes, photos).Render(r.Context(), w)
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

	roomForm := request.UpdateRentalRoomRequest{
		RentalID:         uint(rentalIdInt),
		Name:             makeBedroomName(rentalRooms),
		Floor:            1,
		RentalRoomTypeID: constants.ROOM_TYPE_BEDROOM_ID,
	}
	return rentals.RentalBedroomsFormCreate(params, roomForm, errors, roomTypes, bedTypes).Render(r.Context(), w)

}

func (controller *RentalRoomController) AddBedToRoom(w http.ResponseWriter, r *http.Request) error {

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
	updateParams.Beds = []response.BedResponse{}
	err := controller.rentalRoomService.AddBedToRoom(uint(roomInt), constants.BED_TYPE_TWIN_ID)

	if err != nil {
		return err
	}
	params := request.CreateRentalStep2Params{}

	errors := request.CreateRentalStep2Errors{}
	rentalRooms := controller.rentalRoomService.FindByRentalId(uint(rentalIdInt))
	params.Rooms = rentalRooms
	params.RentalID = uint(rentalIdInt)

	for _, room := range rentalRooms {
		if room.ID == uint(roomInt) {

			for _, bed := range room.Beds {
				updateParams.Beds = append(updateParams.Beds, bed)
			}
		}
	}
	roomTypes := controller.roomTypeService.FindAll()
	bedTypes := controller.bedTypeService.FindAll()

	photos := controller.entityPhotoService.FindAllEntityPhotosForEntity(constants.RENTAL_ENTITY, uint(rentalIdInt))
	return rentals.RentalBedroomsForm(params, updateParams, errors, roomTypes, bedTypes, photos).Render(r.Context(), w)
}
func (controller *RentalRoomController) DeleteBed(w http.ResponseWriter, r *http.Request) error {

	rentalId := chi.URLParam(r, "rentalId")
	bedId := chi.URLParam(r, "bedId")
	rentalIdInt, _ := strconv.Atoi(rentalId)
	bedIdInt, _ := strconv.Atoi(bedId)

	room := r.URL.Query().Get("roomId")
	roomInt, _ := strconv.Atoi(room)

	err := controller.rentalRoomService.DeleteBed(uint(bedIdInt))

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

	updateParams := GetParamsFromRooms(rentalRooms, &roomInt, uint(rentalIdInt))
	photos := controller.entityPhotoService.FindAllEntityPhotosForEntity(constants.RENTAL_ENTITY, uint(rentalIdInt))
	return rentals.RentalBedroomsForm(params, updateParams, errors, roomTypes, bedTypes, photos).Render(r.Context(), w)
}
