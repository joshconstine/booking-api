package controllers

import (
	"booking-api/constants"
	"booking-api/data/request"
	"booking-api/data/response"
	"booking-api/services"
	rentals "booking-api/view/rentals"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"

	"encoding/json"
)

type RentalController struct {
	rentalService      services.RentalService
	amenityService     services.AmenityService
	roomTypeService    services.RoomTypeService
	bedTypeService     services.BedTypeService
	rentalRoomService  services.RentalRoomService
	photoService       services.PhotoService
	entityPhotoService services.EntityPhotoService
}

func NewRentalController(rentalService services.RentalService, amenityService services.AmenityService, roomTypeService services.RoomTypeService, bedTypeService services.BedTypeService, rentalRoomService services.RentalRoomService, photoService services.PhotoService, entityPhotoService services.EntityPhotoService) *RentalController {
	return &RentalController{rentalService: rentalService, amenityService: amenityService, roomTypeService: roomTypeService, bedTypeService: bedTypeService, rentalRoomService: rentalRoomService, photoService: photoService, entityPhotoService: entityPhotoService}

}

type RentalListTemplateData struct {
	PageTitle string
	Rentals   []response.RentalResponse
}
type RentalTemplateData struct {
	PageTitle string
	Rental    response.RentalInformationResponse
}

func (controller *RentalController) FindAll(w http.ResponseWriter, r *http.Request) error {
	rentals := controller.rentalService.FindAll()

	rentalsJSON, err := json.Marshal(rentals)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(rentalsJSON)
	return nil

}

func (controller *RentalController) FindById(ctx *gin.Context) {
	rentalId := ctx.Param("rentalId")
	id, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(id))

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   rental,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *RentalController) CreateForm(w http.ResponseWriter, r *http.Request) error {
	params := request.CreateRentalStep1Params{}
	errors := request.CreateRentalStep1Errors{}
	amenities := controller.amenityService.FindAllSorted()
	return rentals.CreateRental(params, errors, amenities).Render(r.Context(), w)
}

func (controller *RentalController) InformationForm(w http.ResponseWriter, r *http.Request) error {

	rentalId := chi.URLParam(r, "rentalId")
	id, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(id))

	amenities := controller.amenityService.FindAllSorted()
	return rentals.RentalInformationForm(rental, amenities).Render(r.Context(), w)
}

func (controller *RentalController) AvailabilityForm(w http.ResponseWriter, r *http.Request) error {

	rentalId := chi.URLParam(r, "rentalId")
	id, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(id))

	params := request.CreateRentalStep3Params{
		RentalID: rental.ID,
	}
	var errors request.CreateRentalStep3Errors

	return rentals.RentalAvailabilityForm(params, errors).Render(r.Context(), w)
}
func makeBedroomName(existingBedrooms []response.RentalRoomResponse) string {
	count := 1
	for _, room := range existingBedrooms {
		if room.RoomType.ID == constants.ROOM_TYPE_BEDROOM_ID {
			count++
		}
	}
	return "Bedroom " + strconv.Itoa(count)
}

func (controller *RentalController) NewBedroomForm(w http.ResponseWriter, r *http.Request) error {
	rentalId := chi.URLParam(r, "rentalId")
	params := request.CreateRentalStep2Params{}
	errors := request.CreateRentalStep2Errors{}
	rentalIdInt, _ := strconv.Atoi(rentalId)
	rentalRooms := controller.rentalRoomService.FindByRentalId(uint(rentalIdInt))
	params.Rooms = rentalRooms
	params.RentalID = uint(rentalIdInt)
	roomTypes := controller.roomTypeService.FindAll()
	bedTypes := controller.bedTypeService.FindAll()
	var roomForm request.UpdateRentalRoomRequest

	roomForm = request.UpdateRentalRoomRequest{
		RentalID:         uint(rentalIdInt),
		Name:             makeBedroomName(rentalRooms),
		Floor:            1,
		RentalRoomTypeID: constants.ROOM_TYPE_BEDROOM_ID,
	}
	return rentals.RentalBedroomsFormCreate(params, roomForm, errors, roomTypes, bedTypes).Render(r.Context(), w)
}
func GetParamsFromRooms(rooms []response.RentalRoomResponse, roomID *int, rentalID uint) request.UpdateRentalRoomRequest {
	var roomForm request.UpdateRentalRoomRequest
	for _, r := range rooms {
		if r.ID == uint(*roomID) {
			roomForm = request.UpdateRentalRoomRequest{
				ID:               r.ID,
				RentalID:         rentalID,
				Name:             r.Name,
				Description:      r.Description,
				Floor:            r.Floor,
				RentalRoomTypeID: r.RoomType.ID,
				Beds:             r.Beds,
			}

			for _, photo := range r.Photos {
				roomForm.Photos = append(roomForm.Photos, int(photo.ID))

			}

		}
	}
	return roomForm
}
func (controller *RentalController) BedroomForm(w http.ResponseWriter, r *http.Request) error {
	rentalId := chi.URLParam(r, "rentalId")
	room := chi.URLParam(r, "roomId")
	params := request.CreateRentalStep2Params{}
	errors := request.CreateRentalStep2Errors{}
	rentalIdInt, _ := strconv.Atoi(rentalId)
	roomInt, _ := strconv.Atoi(room)
	rentalRooms := controller.rentalRoomService.FindByRentalId(uint(rentalIdInt))
	params.Rooms = rentalRooms
	params.RentalID = uint(rentalIdInt)
	roomTypes := controller.roomTypeService.FindAll()
	bedTypes := controller.bedTypeService.FindAll()
	var roomForm request.UpdateRentalRoomRequest

	if roomInt == 0 {
		if len(rentalRooms) > 0 {
			roomInt = int(rentalRooms[0].ID)
		}
	}
	roomForm = GetParamsFromRooms(rentalRooms, &roomInt, uint(rentalIdInt))

	if len(params.Rooms) == 0 {

		return rentals.RentalBedroomsFormCreate(params, roomForm, errors, roomTypes, bedTypes).Render(r.Context(), w)
	}
	photos := controller.entityPhotoService.FindAllEntityPhotosForEntity(constants.RENTAL_ENTITY, uint(rentalIdInt))

	return rentals.RentalBedroomsForm(params, roomForm, errors, roomTypes, bedTypes, photos).Render(r.Context(), w)
}
func (controller *RentalController) Create(w http.ResponseWriter, r *http.Request) error {
	params := request.CreateRentalStep1Params{}
	bedroomsInt, _ := strconv.Atoi(r.FormValue("bedrooms"))
	guestsInt, _ := strconv.Atoi(r.FormValue("guests"))
	accountID := r.FormValue("accountID")
	accountIDInt, _ := strconv.Atoi(accountID)

	params.Name = r.FormValue("name")
	params.Address = r.FormValue("address")
	params.Description = r.FormValue("description")
	params.Bedrooms = uint(bedroomsInt)
	params.AccountID = uint(accountIDInt)
	params.Bathrooms, _ = strconv.ParseFloat(r.FormValue("bathrooms"), 32)
	params.Guests = uint(guestsInt)
	params.AllowInstantBooking, _ = strconv.ParseBool(r.FormValue("allowInstantBooking"))
	params.AllowPets, _ = strconv.ParseBool(r.FormValue("allowPets"))
	params.ParentProperty, _ = strconv.ParseBool(r.FormValue("parentProperty"))

	errors := request.CreateRentalStep1Errors{}

	amenities := getAmenitiesFromRequest(r)

	params.Amenities = amenities

	rental, err := controller.rentalService.CreateStep1(params)
	if err != nil {
		errors.Name = "Rental could not be created"
	}

	fmt.Println(rental)
	if errors.Name != "" {
		amenities := controller.amenityService.FindAllSorted()
		return rentals.CreateRental(params, errors, amenities).Render(r.Context(), w)
	}

	rentalId := chi.URLParam(r, "rentalId")
	step2params := request.CreateRentalStep2Params{}
	step2errors := request.CreateRentalStep2Errors{}
	rentalIdInt, _ := strconv.Atoi(rentalId)
	params.RentalID = uint(rentalIdInt)
	roomTypes := controller.roomTypeService.FindAll()
	bedTypes := controller.bedTypeService.FindAll()
	bedroomForm := request.UpdateRentalRoomRequest{Name: "Bedroom 1"}
	photos := controller.entityPhotoService.FindAllEntityPhotosForEntity(constants.RENTAL_ENTITY, uint(rentalIdInt))
	return rentals.RentalBedroomsForm(step2params, bedroomForm, step2errors, roomTypes, bedTypes, photos).Render(r.Context(), w)
	//http.Redirect(w, r, "/rentals", http.StatusSeeOther)

	//return rentals.RentalDetails().Render(r.Context(), w)
	return nil
}

func (controller *RentalController) Update(w http.ResponseWriter, r *http.Request) error {

	timeout := 30 * time.Second
	var cancelFn func()
	if timeout > 0 {
		// ctx, cancelFn = ctx.WithTimeout(ctx, timeout)
	}

	if cancelFn != nil {
		defer cancelFn()
	}
	err := r.ParseMultipartForm(10 * 1024 * 1024) // 10 MB limit
	if err != nil {
		http.Error(w, "Failed to parse multipart form", http.StatusInternalServerError)

	}

	err = r.ParseForm()

	rentalId := chi.URLParam(r, "rentalId")
	id, _ := strconv.Atoi(rentalId)
	// Process multiple files
	files := r.MultipartForm.File["photo"]
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Failed to open file", http.StatusInternalServerError)
			return err
		}
		defer file.Close()

		rentalId := chi.URLParam(r, "rentalId")
		id, _ := strconv.Atoi(rentalId)

		photoResult := controller.photoService.AddPhoto(&file, fileHeader, constants.RENTAL_ENTITY, id)
		_ = controller.entityPhotoService.AddPhotoToEntity(photoResult.ID, constants.RENTAL_ENTITY, uint(id))
	}

	bedrooms := r.FormValue("bedrooms")
	bathrooms := r.FormValue("bathrooms")
	address := r.FormValue("address")
	allowPets := r.FormValue("allowPets")
	allowInstantBooking := r.FormValue("allowInstantBooking")

	bedroomsInt, _ := strconv.Atoi(bedrooms)
	bathroomsFloat, _ := strconv.ParseFloat(bathrooms, 64)
	guests := r.FormValue("guests")

	guestsInt, _ := strconv.Atoi(guests)
	params := request.CreateRentalStep1Params{
		RentalID:            uint(id),
		Name:                r.FormValue("name"),
		Description:         r.FormValue("description"),
		Bedrooms:            uint(bedroomsInt),
		Address:             address,
		Guests:              uint(guestsInt),
		Bathrooms:           bathroomsFloat,
		AllowPets:           allowPets == "on",
		AllowInstantBooking: allowInstantBooking == "on",
	}

	errors := request.CreateRentalStep1Errors{}

	amenities := getAmenitiesFromRequest(r)

	params.Amenities = amenities
	_, err = controller.rentalService.UpdateRental(params)
	rental := controller.rentalService.FindById(params.RentalID)

	amenitiesSorted := controller.amenityService.FindAllSorted()
	if err != nil {
		return err

	}

	return rentals.RentalForm(params, errors, amenitiesSorted, rental.Photos).Render(r.Context(), w)
}
func getAmenitiesFromRequest(r *http.Request) []response.AmenityResponse {
	var amenities []uint
	for key, value := range r.Form {
		if len(key) > 8 && key[:8] == "amenity_" {
			amenityId, _ := strconv.Atoi(value[0])
			amenities = append(amenities, uint(amenityId))
		}
	}
	var responseAmenities []response.AmenityResponse
	var amenity response.AmenityResponse
	for _, amenityId := range amenities {
		amenity.ID = uint(amenityId)
		responseAmenities = append(responseAmenities, amenity)

	}
	return responseAmenities
}
func (controller *RentalController) HandleRentalDetail(w http.ResponseWriter, r *http.Request) error {

	// dateParam := chi.URLParam(r, "date")

	rentalId := chi.URLParam(r, "rentalId")
	id, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(id))

	return rentals.RentalInformationResponse(rental).Render(r.Context(), w)
}

func (controller *RentalController) HandleRentalAdminDetail(w http.ResponseWriter, r *http.Request) error {

	rentalId := chi.URLParam(r, "rentalId")
	id, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(id))

	amenities := controller.amenityService.FindAllSorted()
	return rentals.RentalAdmin(rental, amenities).Render(r.Context(), w)
}
func (controller *RentalController) HandleRentalAdminDetailAvailability(w http.ResponseWriter, r *http.Request) error {

	rentalId := chi.URLParam(r, "rentalId")
	id, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(id))
	params := request.CreateRentalStep3Params{}
	errors := request.CreateRentalStep3Errors{}
	return rentals.RentalAdminAvailability(rental, params, errors).Render(r.Context(), w)
}

func (controller *RentalController) HandleRentalAdminDetailRooms(w http.ResponseWriter, r *http.Request) error {
	rentalId := chi.URLParam(r, "rentalId")
	rentalIdInt, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(rentalIdInt))
	var updateParams request.UpdateRentalRoomRequest
	params := request.CreateRentalStep2Params{}
	errors := request.CreateRentalStep2Errors{}
	rentalRooms := controller.rentalRoomService.FindByRentalId(uint(rentalIdInt))
	params.Rooms = rentalRooms
	params.RentalID = uint(rentalIdInt)

	room := r.URL.Query().Get("roomId")
	roomInt, _ := strconv.Atoi(room)
	//If no room is selected, select the first room
	if roomInt == 0 {
		if len(rentalRooms) > 0 {
			roomInt = int(rentalRooms[0].ID)
		}
	}
	updateParams = GetParamsFromRooms(rentalRooms, &roomInt, uint(rentalIdInt))

	roomTypes := controller.roomTypeService.FindAll()
	bedTypes := controller.bedTypeService.FindAll()
	return rentals.RentalAdminRooms(rental, params, updateParams, errors, roomTypes, bedTypes).Render(r.Context(), w)

}
func (controller *RentalController) HandleRentalAdminDetailRoomsCreate(w http.ResponseWriter, r *http.Request) error {
	rentalId := chi.URLParam(r, "rentalId")
	rentalIdInt, _ := strconv.Atoi(rentalId)

	rental := controller.rentalService.FindById(uint(rentalIdInt))
	params := request.CreateRentalStep2Params{}
	errors := request.CreateRentalStep2Errors{}
	rentalRooms := controller.rentalRoomService.FindByRentalId(uint(rentalIdInt))
	params.Rooms = rentalRooms
	params.RentalID = uint(rentalIdInt)

	updateParams := request.UpdateRentalRoomRequest{
		RentalID:         uint(rentalIdInt),
		Name:             makeBedroomName(rentalRooms),
		Floor:            1,
		RentalRoomTypeID: constants.ROOM_TYPE_BEDROOM_ID,
	}

	roomTypes := controller.roomTypeService.FindAll()
	bedTypes := controller.bedTypeService.FindAll()
	return rentals.RentalAdminRoomsCreate(rental, params, updateParams, errors, roomTypes, bedTypes).Render(r.Context(), w)

}
func (controller *RentalController) HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	// user := view.getAuthenticatedUser(r)
	// account, err := db.GetAccountByUserID(user.ID)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%+v\n", user.Account)

	rentalData := controller.rentalService.FindAll()

	return rentals.IndexRentals(rentalData).Render(r.Context(), w)
}
