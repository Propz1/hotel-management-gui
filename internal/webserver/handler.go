package webserver

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"hotelBotGUI/internal/handlers"
	"hotelBotGUI/internal/models"

	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"

	"github.com/jackc/pgx/v5"
	pgxpool "github.com/jackc/pgx/v5/pgxpool"
	zrlog "github.com/rs/zerolog/log"
)

var (
	httpStatus        int
	bookingRequests   = models.NewSliceBookingRequests()
	listGroupedHotels = models.NewListGroupedHotels()
)

type handler struct {
}

func NewHandler() handlers.Handler {
	return &handler{}
}

func (h *handler) Register(router *chi.Mux) {

	// Mounting the subrouter
	//router.Mount("/orders", OrderRoutes())

	///////////////////////////////////////////////////////////////////////////////

	// GET endpoint
	router.Get("/getRequisitions", GetRequisitions)
	router.Get("/getGroupedHotels", GetGroupedHotels)
	router.Get("/getCalendarTable", GetCalendarTable)

	//router.Get("/users/{userId}", GetUser)

	// router.Route("/users", func(r chi.Router) {
	// 	r.Get("/{userId}", GetUser)
	// r.Post("/", createUser)
	// r.Put("/deactivate", deactivateUser)
	// })

	// 	// Apply auth middleware to only `GET /users/{id}`
	// router.Group(func(r chi.Router) {
	// 	r.Use(AuthMiddleware)
	// 	r.Get("/users/{id}")
	// })

	// router.Get("/files/*", func(w http.ResponseWriter, r *http.Request) {
	// 	filepath := chi.URLParam(r, "*")
	// 	http.ServeFile(w, r, filepath)
	// })

	// router.Get("/search", func(w http.ResponseWriter, r *http.Request) {
	// 	query := r.URL.Query().Get("query")
	// 	w.Write([]byte("Search query: " + query))
	// })

	// router.Get("/greet", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Content-Type", "application/json")
	// 	w.WriteHeader(http.StatusOK)
	// 	json.NewEncoder(w).Encode(map[string]string{"message": "Hello, Golang Chi!"})
	// })

	// router.Get("/user/{userID}", func(w http.ResponseWriter, r *http.Request) {

	// 	userID := chi.URLParam(r, "userID")

	// 	if userID == "" {
	// 		http.Error(w, "User ID is required", http.StatusBadRequest)
	// 		return
	// 	}

	// 	// Assuming GetUser returns nil if the user doesn't exist
	// 	// user := getUser(userID)
	// 	// if user == nil {
	// 	// 	http.NotFound(w, r)
	// 	// 	return
	// 	// }

	// 	// Handle the user found case here
	// })

	//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// POST endpoint
	// router.Post("/users", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Create a new user"))
	// })

	// router.Post("/users", func(w http.ResponseWriter, r *http.Request) {
	// 	var user models.User
	// 	json.NewDecoder(r.Body).Decode(&user)
	// 	w.Write([]byte("Added user: " + user.Name))
	// })
	///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
	// PUT endpoint

	router.Route("/bookingRequest", func(r chi.Router) {
		r.Options("/", OptionsHundler)
		r.Patch("/", UpdateBookingRequest)
	})

	router.Route("/updateDataForCalendar", func(r chi.Router) {
		r.Options("/", OptionsHundler)
		r.Put("/", UpdateCalendar)
	})

	// 	num := chi.URLParam(r, "requisitionNumber")
	// 	fmt.Printf("\nrequisitionNumber = %v\n", num)
	// 	w.Write([]byte("Update: " + num))
	// })

	// // Subrouters:
	//  router.Route("/bookingRequest{requisitionNumber}", func(r chi.Router) {
	// // 	r.Use(ArticleCtx)
	// // 	r.Get("/", getArticle)                                          // GET /articles/123
	//  	r.Put("/", updateArticle)                                       // PUT /articles/123
	// 	r.Options("/", updateArticle)
	// // 	r.Delete("/", deleteArticle)                                    // DELETE /articles/123
	//    })

	// 	// r.Route("/{requisitionNumber}", func(r chi.Router) {
	// 	// 	//r.Use(ArticleCtx)            // Load the *Article on the request context
	// 	// 	//r.Get("/", GetArticle)       // GET /articles/123
	// 	// 	r.Put("/", UpdateBookingRequest)
	// 	// 	//r.Delete("/", DeleteArticle)
	// 	// })

	// 	// GET /articles/whats-up
	// 	//r.With(ArticleCtx).Get("/{articleSlug:[a-z-]+}", GetArticle)
	// })

	// router.Options("/bookingRequest/{requisitionNumber}", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
	// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	// 	w.Header().Set("Access-Control-Allow-Headers", "X-Custom-Header, X-PINGOTHER, Content-Type")

	// 	w.WriteHeader(http.StatusOK)
	// })

	// router.Put("/bookingRequest/{requisitionNumber}", func(w http.ResponseWriter, r *http.Request) {

	// 	//num := chi.URLParam(r, "requisitionNumber")

	// 	num := r.URL.Query().Get("requisitionNumber")

	// 	fmt.Printf("\nrequisitionNumber = %v\n", num)

	// 	fmt.Printf("\nr.Context() = %v\n", r.Context())

	// 	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// 	w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
	// 	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, OPTIONS")
	// 	w.Header().Set("Access-Control-Allow-Headers", "X-Custom-Header, X-PINGOTHER, Content-Type")

	// 	w.WriteHeader(http.StatusOK)

	// 	//w.Write([]byte("Update: "))
	// })

	/////////////////////////////////////////////////////////////////////////////////////////////
	// DELETE endpoint
	// router.Delete("/users/{userId}", func(w http.ResponseWriter, r *http.Request) {
	// 	userId := chi.URLParam(r, "userId")
	// 	w.Write([]byte("Delete user: " + userId))
	// })

	// router.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == "POST" {
	// 		// обработать POST запрос
	// 	} else if r.Method == "PUT" {
	// 		// обработать PUT запрос
	// 	} else if r.Method == "DELETE" {
	// 		// обработать DELETE запрос
	// 	}
	// })

}

func OrderRoutes() chi.Router {
	r := chi.NewRouter()
	// r.Get("/", ListOrders)        // List orders
	r.Get("/getRequisitions", GetRequisitions)
	// r.Get("/{orderID}", GetOrder) // Get a specific order
	return r
}

// func GetUser(w http.ResponseWriter, r *http.Request) {
// 	userId := chi.URLParam(r, "userId")
// 	w.Write([]byte("Retrieve user: " + userId))
// }

// func getUser(userID int64) {
// return

// }

func OptionsHundler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "https://77d3-176-117-0-227.ngrok-free.app")
	//w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "X-Custom-Header, X-PINGOTHER, Content-Type")

	w.WriteHeader(http.StatusOK)
}

func GetCalendarTable(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://77d3-176-117-0-227.ngrok-free.app")
	//w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(map[string]string{"message": "Hello, Golang Chi!"})

	listCalendarTable := models.NewSliceCalendarTable()

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	dbpool.Config().MaxConns = 3

	if err != nil {
		zrlog.Error().Msg(fmt.Sprintf("Creat pgxpool failed: %+v\n", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer dbpool.Close()

	err = getCalendarFromDB(ctx, "All", listCalendarTable, dbpool)

	if err != nil {
		zrlog.Error().Msg(fmt.Sprintf("Connect to db is fail: %+v\n", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(*listCalendarTable); err != nil {
		zrlog.Error().Msg(fmt.Sprintf("Somthing rong in json.NewEncoder(w).Encode(*listCalendarTable): %+v\n", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

}

func GetGroupedHotels(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://77d3-176-117-0-227.ngrok-free.app")
	//w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
	w.WriteHeader(http.StatusOK)

	*listGroupedHotels = nil

	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()

	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	dbpool.Config().MaxConns = 12

	if err != nil {
		zrlog.Error().Msg(fmt.Sprintf("Creat pgxpool failed: %+v\n", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer dbpool.Close()

	err = getHotelsFromDB(context.Background(), "GroupedHotels", dbpool)

	if err != nil {
		zrlog.Error().Msg(fmt.Sprintf("Connect to db is fail: %+v\n", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(*listGroupedHotels); err != nil {
		zrlog.Error().Msg(fmt.Sprintf("Somthing rong in json.NewEncoder(w).Encode(*listGroupedHotels): %+v\n", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetRequisitions(w http.ResponseWriter, r *http.Request) {

	//host := os.Getenv("WEBSERVER_HOST")
	//port := os.Getenv("WEBSERVER_PORT")

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://77d3-176-117-0-227.ngrok-free.app")
	//w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
	w.WriteHeader(http.StatusOK)
	//json.NewEncoder(w).Encode(map[string]string{"message": "Hello, Golang Chi!"})

	*bookingRequests = nil

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dbpool, err := pgxpool.New(ctx, os.Getenv("DATABASE_URL"))
	dbpool.Config().MaxConns = 12

	if err != nil {
		zrlog.Error().Msg(fmt.Sprintf("Creat pgxpool failed: %+v\n", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer dbpool.Close()

	err = getRequisitionsFromDB(ctx, "AllBookingRequestes", dbpool)

	if err != nil {
		zrlog.Error().Msg(fmt.Sprintf("Connect to db is fail: %+v\n", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

	if err := json.NewEncoder(w).Encode(*bookingRequests); err != nil {
		zrlog.Error().Msg(fmt.Sprintf("Somthing rong in json.NewEncoder(w).Encode(*bookingRequests): %+v\n", err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
	}

	// jsonResp, err := json.Marshal(*bookingRequests)
	// fmt.Printf("\n%v\n\n", jsonResp)
	// w.Write(jsonResp)

}

func getHotelsFromDB(ctx context.Context, toggle string, dbpool *pgxpool.Pool) error {

	var querySQL string

	switch {
	case toggle == "GroupedHotels":
		querySQL = fmt.Sprintf("SELECT COALESCE(hotel.city,'')as city, COALESCE(hotel.name,'') as name, COALESCE(hotel.stars,0) as stars, COALESCE(hotel.address, '') as address, COALESCE(hotel.code_city,'') as code_city, COALESCE(hotel.phone,'') as phone, COALESCE(hotel.site, '') as site, COALESCE(hotel.code, '') as code_hotel, COUNT(hotel.name) as count_hotel, COUNT(hotel.city) as count_city FROM \"Hotel\" AS hotel GROUP BY ROLLUP (hotel.city, (hotel.city, hotel.name, hotel.stars, hotel.address, hotel.code_city, hotel.phone, hotel.site, hotel.code)) ORDER BY hotel.city, hotel.name DESC")
	default:
		querySQL = fmt.Sprintf("SELECT COALESCE(hotel.city,'')as city, COALESCE(hotel.name,'') as name, COALESCE(hotel.stars,0) as stars, COALESCE(hotel.address, '') as address, COALESCE(hotel.code_city,'') as code_city, COALESCE(hotel.phone,'') as phone, COALESCE(hotel.site, '') as site, COALESCE(hotel.code, '') as code_hotel, COUNT(hotel.name) as count_hotel, COUNT(hotel.city) as count_city FROM \"Hotel\" AS hotel GROUP BY ROLLUP (hotel.city, (hotel.city, hotel.name, hotel.stars, hotel.address, hotel.code_city, hotel.phone, hotel.site, hotel.code)) ORDER BY hotel.city, hotel.name DESC")
	}

	row, err := dbpool.Query(ctx, querySQL)
	if err != nil {
		return fmt.Errorf("func getHotelsFromDB(), dbpool.Query is fail %w", err)
	}

	var city string
	var name string
	var stars int
	var address string
	var codeCity string
	var phone string
	var site string
	var codeHotel string
	var countHotelsInTheCity int
	var countCities int

	var currentCity string
	var currentCodeCity string
	var countHotels int = 1000000
	var sliceHotels []models.MultipleSelectHotels

	for row.Next() {

		err := row.Scan(&city, &name, &stars, &address, &codeCity, &phone, &site, &codeHotel, &countHotelsInTheCity, &countCities)

		if err != nil {
			return fmt.Errorf("func getHotelsFromDB(), scan datas of row is fail %w", err)
		}

		if city == "" && name == "" {
			continue
		}

		if city != "" && name == "" {

			countHotels = countHotelsInTheCity
			sliceHotels = nil
			currentCity = city
			currentCodeCity = codeCity

		} else {

			hotels := models.MultipleSelectHotels{
				Name:  name,
				Value: address,
				City:  currentCity,
			}

			sliceHotels = append(sliceHotels, hotels)

			countHotels--
		}

		if countHotels == 0 {

			if len(sliceHotels) != 0 {

				groupedHotels := models.GroupedHotels{
					City:     currentCity,
					CodeCity: currentCodeCity,
					Hotels:   sliceHotels,
				}

				*listGroupedHotels = append(*listGroupedHotels, groupedHotels)
			}
		}

	}

	return nil

}

func getRequisitionsFromDB(ctx context.Context, toggle string, dbpool *pgxpool.Pool) error {

	var querySQL string
	var statusBookingRequest string

	switch {
	case toggle == "NewBookingRequestes":
		statusBookingRequest = "new"
		querySQL = fmt.Sprintf("SELECT br.requisition_number, br.check_in_date, br.check_out_date, br.application_submission_time, br.status, br.application_status_date, br.cost, cl.name, cl.email,cl.telephone FROM \"Booking_Request\" AS br LEFT JOIN \"Client\" AS cl ON br.user_telegram_id=cl.user_telegram_id WHERE br.status='%v'", statusBookingRequest)
	case toggle == "RejectedBookingRequestes":
		statusBookingRequest = "rejected"
		querySQL = fmt.Sprintf("SELECT br.requisition_number, br.check_in_date, br.check_out_date, br.application_submission_time, br.status, br.application_status_date, br.cost, cl.name, cl.email,cl.telephone FROM \"Booking_Request\" AS br LEFT JOIN \"Client\" AS cl ON br.user_telegram_id=cl.user_telegram_id WHERE br.status='%v'", statusBookingRequest)
	case toggle == "CanselledBookingRequestes":
		statusBookingRequest = "canselled"
		querySQL = fmt.Sprintf("SELECT br.requisition_number, br.check_in_date, br.check_out_date, br.application_submission_time, br.status, br.application_status_date, br.cost, cl.name, cl.email,cl.telephone FROM \"Booking_Request\" AS br LEFT JOIN \"Client\" AS cl ON br.user_telegram_id=cl.user_telegram_id WHERE br.status='%v'", statusBookingRequest)
	case toggle == "AllBookingRequestes":
		querySQL = fmt.Sprintf("SELECT br.requisition_number, br.check_in_date, br.check_out_date, br.application_submission_time, br.status, br.application_status_date, br.cost, cl.name, cl.email,cl.telephone,cl.user_telegram_id, COALESCE(br.hotel_name,'') as hotel_name, COALESCE(ht.code,'') as hotel_code, COALESCE(ht.city,'') as hotel_city, COALESCE(ht.code_city,'') as hotel_city, COALESCE(ht.address, '') as hotel_address, COALESCE(ht.stars,0) as hotel_stars, COALESCE(ht.phone,'') as hotel_phone FROM \"Booking_Request\" AS br LEFT JOIN \"Client\" AS cl ON br.user_telegram_id=cl.user_telegram_id LEFT JOIN \"Hotel\" AS ht ON (br.city=ht.city AND br.hotel_name=ht.name)")

	default:
		querySQL = fmt.Sprintf("SELECT br.requisition_number, br.check_in_date, br.check_out_date, br.application_submission_time, br.status, br.application_status_date, br.cost, cl.name, cl.email,cl.telephone,cl.user_telegram_id, COALESCE(br.hotel_name,'') as hotel_name, COALESCE(ht.code,'') as hotel_code, COALESCE(ht.city,'') as hotel_city, COALESCE(ht.code_city,'') as hotel_city, COALESCE(ht.address, '') as hotel_address, COALESCE(ht.stars,0) as hotel_stars, COALESCE(ht.phone,'') as hotel_phone FROM \"Booking_Request\" AS br LEFT JOIN \"Client\" AS cl ON br.user_telegram_id=cl.user_telegram_id LEFT JOIN \"Hotel\" AS ht ON (br.city=ht.city AND br.hotel_name=ht.name)")
	}

	row, err := dbpool.Query(ctx, querySQL)
	if err != nil {
		return fmt.Errorf("func GetRequisitions(), dbpool.Query is fail %w", err)
	}

	var requisitionNumber int
	var checkInDate int64
	var checkOutDate int64
	var applicationSubmissionTime int64
	var applicationStatusDate int64
	var cost int
	var status string
	var userName string
	var userEmail string
	var userPhone string
	var userTelegramID int64
	var hotelName string
	var hotelCode string
	var hotelAddress string
	var hotelCity string
	var hotelCodeCity string
	var hotelPhone string
	var hotelStars int

	for row.Next() {

		err := row.Scan(&requisitionNumber, &checkInDate, &checkOutDate, &applicationSubmissionTime, &status, &applicationStatusDate, &cost, &userName, &userEmail, &userPhone, &userTelegramID, &hotelName,
			&hotelCode, &hotelCity, &hotelCodeCity, &hotelAddress, &hotelStars, &hotelPhone)

		if err != nil {
			return fmt.Errorf("func GetRequisitions(), scan datas of row is fail %w", err)
		}

		guest := models.User{
			TelegramID: userTelegramID,
			Name:       userName,
			Email:      userEmail,
			Telephone:  userPhone,
		}

		hotel := models.Hotel{
			Name:     hotelName,
			Code:     hotelCode,
			City:     hotelCity,
			CodeCity: hotelCodeCity,
			Stars:    hotelStars,
			Address:  hotelAddress,
			Phone:    hotelPhone,
		}

		br := models.BookingRequests{
			RequisitionNumber:         requisitionNumber,
			CheckInDate:               checkInDate,
			CheckOutDate:              checkOutDate,
			ApplicationSubmissionTime: applicationSubmissionTime,
			ApplicationStatusDate:     applicationStatusDate,
			Status:                    status,
			Cost:                      cost,
			User:                      guest,
			Hotel:                     hotel,
		}

		*bookingRequests = append(*bookingRequests, br)
	}

	return nil

}

func getCalendarFromDB(ctx context.Context, toggle string, listCalendarTable *models.ListCalendarTable, dbpool *pgxpool.Pool) error {

	var querySQL string

	switch {
	case toggle == "All":
		querySQL = fmt.Sprintf("SELECT COALESCE(date,0) as date, COALESCE(installation_date,0) as installation_date, COALESCE(hotel_name,'') as hotel_name, COALESCE(hotel_city, '') as hotel_city, COALESCE(hotel_address,'') as hotel_address, COALESCE(ht.stars,0) as stars, COALESCE(quantity,0) as quantity, COALESCE(price_room,0) as price_room, COALESCE(price_bed_twin, 0) as price_bed_twin, COALESCE(price_bed_triple,0) as price_bed_triple, COALESCE(price_bed_quadriple,0) as price_bed_quadriple, COALESCE(cashback,0) as cashback, COALESCE(cashback_percent,0) as cashback_percent, COALESCE(cashback_no_check_in,0) as cashback_no_check_in, COALESCE(cashback_no_check_in_percent,0) as cashback_no_check_in_percent, COALESCE(administrator_id, 0) as administrator_id,  COALESCE(adm.name,'') as administrator_name FROM \"%v\" as cl LEFT JOIN \"%v\" as adm ON cl.administrator_id=adm.id LEFT JOIN \"%v\" as ht ON (cl.hotel_name=ht.name AND cl.hotel_city=ht.city AND cl.hotel_address=ht.address) ORDER BY cl.date", "Calendar", "Administrator", "Hotel")
	default:
		querySQL = fmt.Sprintf("SELECT COALESCE(date,0) as date, COALESCE(installation_date,0) as installation_date, COALESCE(hotel_name,'') as hotel_name, COALESCE(hotel_city, '') as hotel_city, COALESCE(hotel_address,'') as hotel_address, COALESCE(ht.stars,0) as stars, COALESCE(quantity,0) as quantity, COALESCE(price_room,0) as price_room, COALESCE(price_bed_twin, 0) as price_bed_twin, COALESCE(price_bed_triple,0) as price_bed_triple, COALESCE(price_bed_quadriple,0) as price_bed_quadriple, COALESCE(cashback,0) as cashback, COALESCE(cashback_percent,0) as cashback_percent, COALESCE(cashback_no_check_in,0) as cashback_no_check_in, COALESCE(cashback_no_check_in_percent,0) as cashback_no_check_in_percent, COALESCE(administrator_id, 0) as administrator_id,  COALESCE(adm.name,'') as administrator_name FROM \"%v\" as cl LEFT JOIN \"%v\" as adm ON cl.administrator_id=adm.id LEFT JOIN \"%v\" as ht ON (cl.hotel_name=ht.name AND cl.hotel_city=ht.city AND cl.hotel_address=ht.address) ORDER BY cl.date", "Calendar", "Administrator", "Hotel")
	}

	rows, err := dbpool.Query(ctx, querySQL)
	if err != nil {
		return fmt.Errorf("func getCalendarFromDB(), dbpool.Query is fail %w", err)
	}

	var date int64
	var installationDate int64
	var hotelName string
	var hotelCity string
	var hotelAddress string
	var hotelStars int
	var quantity int
	var priceRoom int
	var priceBedTwin int
	var priceBedTriple int
	var priceBedQuadriple int
	var cashback int
	var cashbackPercent int
	var cashbackNoCheckIn int
	var cashbackNoCheckInPercent int
	var administratorID int64
	var administratorName string

	for rows.Next() {

		err := rows.Scan(&date, &installationDate, &hotelName, &hotelCity, &hotelAddress, &hotelStars, &quantity, &priceRoom, &priceBedTwin, &priceBedTriple, &priceBedQuadriple, &cashback, &cashbackPercent, &cashbackNoCheckIn,
			&cashbackNoCheckInPercent, &administratorID, &administratorName)

		if err != nil {
			return fmt.Errorf("func getCalendarFromDB(), scan datas of row is fail %w", err)
		}

		cal := models.CalendarTable{
			Date:                   date,
			InstallationDate:       installationDate,
			Quantity:               quantity,
			HotelName:              hotelName,
			HotelCity:              hotelCity,
			HotelAddress:           hotelAddress,
			HotelStars:             hotelStars,
			PriceRoom:              priceRoom,
			PriceBedTwin:           priceBedTwin,
			PriceBedTriple:         priceBedTwin,
			PriceBedQuadriple:      priceBedQuadriple,
			Cashback:               cashback,
			CashbackPercent:        cashbackPercent,
			CashbackNoCheckIn:      cashbackNoCheckIn,
			CashbackNoCheckPercent: cashbackNoCheckInPercent,
			AdministratorID:        administratorID,
			AdministratorName:      administratorName,
		}

		*listCalendarTable = append(*listCalendarTable, cal)
	}

	return nil
}

func UpdateCalendar(w http.ResponseWriter, r *http.Request) {

	var badRequest bool

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://77d3-176-117-0-227.ngrok-free.app")
	//w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
	w.WriteHeader(http.StatusOK)

	var updateCalendar = models.NewCalendaUpdate()

	err := json.NewDecoder(r.Body).Decode(updateCalendar)

	if err != nil {
		zrlog.Error().Msg(fmt.Sprintf("json.NewDecoder(): %+v\n", err.Error()))
		badRequest = true
		w.WriteHeader(http.StatusBadRequest)
	}

	if ((updateCalendar.DateEndDay < updateCalendar.DateStartDay) && (updateCalendar.DateEndMonth == updateCalendar.DateStartMonth) && (updateCalendar.DateEndYear == updateCalendar.DateStartYear)) || (updateCalendar.DateEndMonth < updateCalendar.DateStartMonth) || (updateCalendar.DateEndYear < updateCalendar.DateStartYear) {
		zrlog.Error().Msg("End date < start date")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Not allowed, that the end date was less the start date!"))
	}

	if !badRequest {

		startDate := time.Date(updateCalendar.DateStartYear, time.Month(updateCalendar.DateStartMonth), updateCalendar.DateStartDay, 0, 0, 0, 0, time.Local)
		endDate := time.Date(updateCalendar.DateEndYear, time.Month(updateCalendar.DateEndMonth), updateCalendar.DateEndDay, 0, 0, 0, 0, time.Local)

		//TEST
		fmt.Printf("\nНачало периода %v\n", startDate)
		fmt.Printf("\nНачало периода %v\n", startDate.UnixNano())

		//TEST
		fmt.Printf("\nКонец периода %v\n", endDate)
		fmt.Printf("\nКонец периода %v\n", endDate.UnixNano())
		//TEST

		diff := endDate.Sub(startDate)

		//diff.Hours()// number of Hours
		//diff.Nanoseconds()// number of Nanoseconds
		//diff.Minutes()// number of Minutes
		//diff.Seconds()// number of Seconds

		numbersOfDays := diff.Hours() / 24

		//TEST
		fmt.Printf("numbersOfDays => %v", numbersOfDays)
		//TEST

		var periodDates []int64

		periodDates = append(periodDates, startDate.UnixNano())

		if numbersOfDays > 0 {

			for i := 1; i <= int(numbersOfDays); i++ {
				periodDates = append(periodDates, startDate.AddDate(0, 0, i).UnixNano())
			}

		}

		errG, _ := errgroup.WithContext(context.Background())
		errG.SetLimit(5)

		for _, date := range periodDates {

			day := date

			errG.Go(func() error {
				err := dbUpdateCalendar(context.Background(), day, updateCalendar)
				return err
			})
		}

		err := errG.Wait()

		if err != nil {
			zrlog.Error().Msg(fmt.Sprintf("func dbUpdateCalendar(): %+v\n", err.Error()))
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

	}

}

func dbUpdateCalendar(ctx context.Context, day int64, updateCalendar *models.CalendarUpdate) error {

	errGr, _ := errgroup.WithContext(ctx)
	errGr.SetLimit(5)

	for _, hotel := range updateCalendar.Hotels {

		h := hotel

		errGr.Go(func() error {
			err := dbUpdateRecord(ctx, day, h.Name, h.Address, h.City, updateCalendar)
			return err
		})

	}

	errr := errGr.Wait()

	// 	// rows.Next()

	return errr
}

func dbUpdateRecord(ctx context.Context, day int64, hotelName string, hotelAddress string, hotelCity string, updateCalendar *models.CalendarUpdate) error {

	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))

	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	querySQL := fmt.Sprintf("with updated as (UPDATE \"%v\" SET date='%v',installation_date='%v', hotel_name='%v', hotel_address='%v', hotel_city='%v', quantity='%v', price_room='%v', price_bed_twin='%v', price_bed_triple='%v', price_bed_quadriple='%v',cashback='%v', cashback_percent='%v', cashback_no_check_in='%v', cashback_no_check_in_percent='%v' WHERE (date=$1 AND hotel_name=$2 AND hotel_address=$3 AND hotel_city=$4) RETURNING *) INSERT INTO \"%v\" (date, installation_date, hotel_name, hotel_address, hotel_city,quantity, price_room, price_bed_twin, price_bed_triple, price_bed_quadriple, cashback, cashback_percent, cashback_no_check_in, cashback_no_check_in_percent) SELECT '%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v','%v' WHERE NOT EXISTS (SELECT 1 FROM updated WHERE (date=$1 AND hotel_name=$2 AND hotel_address=$3 AND hotel_city=$4));", "Calendar", day, time.Now().UnixNano(), hotelName, hotelAddress, hotelCity, updateCalendar.QuantityNumbers, updateCalendar.PriceRoom, updateCalendar.PriceBedTwin, updateCalendar.PriceBedTriple, updateCalendar.PriceBedQuadriple, updateCalendar.Cashback, updateCalendar.CashbackPercent, updateCalendar.CashbackNoCheckIn, updateCalendar.CashbackNoCheckPercent, "Calendar", day, time.Now().UnixNano(), hotelName, hotelAddress, hotelCity, updateCalendar.QuantityNumbers, updateCalendar.PriceRoom, updateCalendar.PriceBedTwin, updateCalendar.PriceBedTriple, updateCalendar.PriceBedQuadriple, updateCalendar.Cashback, updateCalendar.CashbackPercent, updateCalendar.CashbackNoCheckIn, updateCalendar.CashbackNoCheckPercent)

	rows, _ := conn.Query(ctx, querySQL, day, hotelName, hotelAddress, hotelCity)
	rows.Close()

	if rows.Err() != nil {
		return fmt.Errorf("func dbUpdateRecord() for Calendar table, query is faile: %w", rows.Err())
	}

	return nil
}

func UpdateBookingRequest(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "https://77d3-176-117-0-227.ngrok-free.app")
	//w.Header().Set("Access-Control-Allow-Origin", "https://localhost:5173")
	w.WriteHeader(http.StatusOK)

	requisitionNumber := r.URL.Query().Get("requisitionNumber")
	status := r.URL.Query().Get("status")

	switch {
	case requisitionNumber != "" && status != "":

		_, err := strconv.Atoi(requisitionNumber)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			zrlog.Error().Msg(fmt.Sprintf("Unable to connect to database: %+v\n", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		}
		defer conn.Close(context.Background())

		err = dbUpdateRequisition(context.Background(), requisitionNumber, status, conn)

		if err != nil {
			zrlog.Error().Msg(fmt.Sprintf("query UPDATE to db is failed: %+v\n", err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		}

	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}

func dbUpdateRequisition(ctx context.Context, requisitionNumber string, status string, conn *pgx.Conn) error {

	var statusDate int64

	reqNum, _ := strconv.Atoi(requisitionNumber)

	query := fmt.Sprintf("UPDATE %s SET status='%v', application_status_date='%v' WHERE requisition_number=$1 RETURNING application_status_date", "\"Booking_Request\"", status, time.Now().UnixNano())

	err := conn.QueryRow(context.Background(), query, reqNum).Scan(&statusDate)

	if err != nil {
		return fmt.Errorf("dbUpdateRequisition, query to db is faile: %w", err)
	}

	// row, err := dbpool.Query(ctx, fmt.Sprintf("UPDATE %s SET status='%v', application_status_date='%v' WHERE requisition_number=$1 RETURNING application_status_date", "\"Booking_Request\"", status, time.Now().UnixNano()), reqNum)

	// if err != nil {
	// 	return fmt.Errorf("dbUpdateRequisition, query to db is faile: %w", err)
	// }

	// if row.Next() {
	// 	return row.Err()
	// }

	return nil
}
