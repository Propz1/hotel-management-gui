package models

type CalendarUpdate struct {
	Hotels                 []SelectedHotels `json:"jsonSelectedHotels"`
	DateStartDay           int              `json:"startDay"`
	DateStartMonth         int              `json:"startMonth"`
	DateStartYear          int              `json:"startYear"`
	DateEndDay             int              `json:"endDay"`
	DateEndMonth           int              `json:"endMonth"`
	DateEndYear            int              `json:"endYear"`
	InstallationDate       int64            `json:"installationDate"`
	QuantityNumbers        int              `json:"numbersQuantity"`
	PriceRoom              int              `json:"priceRoom"`
	PriceBedTwin           int              `json:"priceBedTwin"`
	PriceBedTriple         int              `json:"priceBedTriple"`
	PriceBedQuadriple      int              `json:"priceBedQuadriple"`
	Cashback               int              `json:"cashback"`
	CashbackPercent        int              `json:"cashbackPercent"`
	CashbackNoCheckIn      int              `json:"cashbackNoCheckIn"`
	CashbackNoCheckPercent int              `json:"cashbackNoCheckInPercent"`
}

func NewCalendaUpdate() *CalendarUpdate {
	var updateCalendar CalendarUpdate
	return &updateCalendar
}

type SelectedHotels struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	City    string `json:"city"`
}

type MultipleSelectHotels struct {
	Name  string `json:"label"`
	Value string `json:"value"`
	City  string `json:"city"`
}

type GroupedHotels struct {
	City     string                 `json:"label"`
	CodeCity string                 `json:"code"`
	Hotels   []MultipleSelectHotels `json:"items"`
}

type ListGroupedHotels []GroupedHotels

func NewListGroupedHotels() *ListGroupedHotels {
	var lustGH ListGroupedHotels
	lustGH = make(ListGroupedHotels, 0)
	return &lustGH
}

type CalendarTable struct {
	Date                   int64  `json:"date"`
	InstallationDate       int64  `json:"installationDate"`
	Quantity               int    `json:"quantity"`
	HotelName              string `json:"hotelName"`
	HotelCity              string `json:"hotelCity"`
	HotelAddress           string `json:"hotelAddress"`
	HotelStars             int    `json:"hotelStars"`
	PriceRoom              int    `json:"priceRoom"`
	PriceBedTwin           int    `json:"priceBedTwin"`
	PriceBedTriple         int    `json:"priceBedTriple"`
	PriceBedQuadriple      int    `json:"priceBedQuadriple"`
	Cashback               int    `json:"cashback"`
	CashbackPercent        int    `json:"cashbackPercent"`
	CashbackNoCheckIn      int    `json:"cashbackNoCheckIn"`
	CashbackNoCheckPercent int    `json:"cashbackNoCheckInPercent"`
	AdministratorID        int64  `json:"administratorID"`
	AdministratorName      string `json:"administratorName"`
}

type ListCalendarTable []CalendarTable

func NewSliceCalendarTable() *ListCalendarTable {
	var listCT ListCalendarTable
	listCT = make(ListCalendarTable, 0)
	return &listCT
}

func NewSliceBookingRequests() *BR {
	var br BR
	br = make(BR, 0)
	return &br
}

type BR []BookingRequests

type BookingRequests struct {
	RequisitionNumber         int    `json:"requisitionNumber"`
	CheckInDate               int64  `json:"checkInDate"`
	CheckOutDate              int64  `json:"checkOutDate"`
	ApplicationSubmissionTime int64  `json:"applicationSubmissionTime"`
	ApplicationStatusDate     int64  `json:"applicationStatusDate"`
	Status                    string `json:"status"`
	Cost                      int    `json:"cost"`
	User                      User   `json:"user"`
	Hotel                     Hotel  `json:"hotel"`
}

type User struct {
	TelegramID int64  `json:"telegramId"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Telephone  string `json:"phone"`
}

type Hotel struct {
	Name     string `json:"name"`
	Code     string `json:"code"`
	City     string `json:"city"`
	CodeCity string `json:"codeCity"`
	Stars    int    `json:"stars"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Site     string `json:"site"`
}
