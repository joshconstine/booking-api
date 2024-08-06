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

	"github.com/gin-gonic/gin"
	"github.com/go-chi/chi/v5"

	"encoding/json"
)

type RentalController struct {
	rentalService     services.RentalService
	amenityService    services.AmenityService
	roomTypeService   services.RoomTypeService
	bedTypeService    services.BedTypeService
	rentalRoomService services.RentalRoomService
}

func NewRentalController(rentalService services.RentalService, amenityService services.AmenityService, roomTypeService services.RoomTypeService, bedTypeService services.BedTypeService, rentalRoomService services.RentalRoomService) *RentalController {
	return &RentalController{rentalService: rentalService, amenityService: amenityService, roomTypeService: roomTypeService, bedTypeService: bedTypeService, rentalRoomService: rentalRoomService}

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

	//If no room is selected, select the first room
	if roomInt == 0 {
		if len(rentalRooms) > 0 {
			roomInt = int(rentalRooms[0].ID)
		}
	}
	if roomInt != 0 {
		for _, r := range rentalRooms {
			if r.ID == uint(roomInt) {

				roomForm = request.UpdateRentalRoomRequest{
					ID:               r.ID,
					RentalID:         uint(rentalIdInt),
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
	}
	if len(params.Rooms) == 0 {

		return rentals.RentalBedroomsFormCreate(params, roomForm, errors, roomTypes, bedTypes).Render(r.Context(), w)
	}

	return rentals.RentalBedroomsForm(params, roomForm, errors, roomTypes, bedTypes).Render(r.Context(), w)
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
	return rentals.RentalBedroomsForm(step2params, bedroomForm, step2errors, roomTypes, bedTypes).Render(r.Context(), w)
	//http.Redirect(w, r, "/rentals", http.StatusSeeOther)

	//return rentals.RentalDetails().Render(r.Context(), w)
	return nil
}

func (controller *RentalController) Update(w http.ResponseWriter, r *http.Request) error {
	rentalId := chi.URLParam(r, "rentalId")
	bedrooms := r.FormValue("bedrooms")
	bathrooms := r.FormValue("bathrooms")
	address := r.FormValue("address")
	allowPets := r.FormValue("allowPets")
	allowInstantBooking := r.FormValue("allowInstantBooking")

	id, _ := strconv.Atoi(rentalId)
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
		AllowPets:           allowPets == "checked",
		AllowInstantBooking: allowInstantBooking == "checked",
	}

	errors := request.CreateRentalStep1Errors{}
	//

	amenities := getAmenitiesFromRequest(r)

	params.Amenities = amenities
	_, err := controller.rentalService.UpdateRental(params)

	amenitiesSorted := controller.amenityService.FindAllSorted()
	if err != nil {
		return err

	}

	return rentals.RentalForm(params, errors, amenitiesSorted).Render(r.Context(), w)
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

func (controller *RentalController) HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	// user := view.getAuthenticatedUser(r)
	// account, err := db.GetAccountByUserID(user.ID)
	// if err != nil {
	// 	return err
	// }
	// fmt.Printf("%+v\n", user.Account)

	rentalData := controller.rentalService.FindAll()

	return rentals.Index(rentalData).Render(r.Context(), w)
}
