package v1

import "time"

// GetTokenResponse .
type GetTokenResponse struct {
	Expire time.Time `json:"expire"` // วันเวลาที่ Token จะหมดอายุ
	Token  string    `json:"token"`  // Token ที่ใช้
}

// LanguageSupport .
type LanguageSupport string

//
const (
	LanguageTH LanguageSupport = "TH"
	LanguageEN LanguageSupport = "EN"
	LanguageCN LanguageSupport = "CN"
)

// TrackRequest .
type TrackRequest struct {
	Status   string          `json:"status"`
	Language LanguageSupport `json:"language"`
	Barcode  []string        `json:"barcode"`
}

// TrackResponse .
type TrackResponse struct {
	Response struct {
		Items map[string]struct {
			Barcode             string     `json:"barcode"`
			Status              string     `json:"status"`
			StatusDescription   string     `json:"status_description"`
			StatusDate          time.Time  `json:"status_date"`
			Location            string     `json:"location"`
			Postcode            string     `json:"postcode"`
			DeliveryStatus      *string    `json:"delivery_status"`
			DeliveryDescription *string    `json:"delivery_description"`
			DeliveryDatetime    *time.Time `json:"delivery_datetime"`
			ReceiverName        *string    `json:"receiver_name"`
			Signature           *string    `json:"signature"`
		} `json:"items"`
		TrackCount struct {
			TrackDate       string `json:"track_date"`
			CountNumber     int    `json:"count_number"`
			TrackCountLimit int    `json:"track_count_limit"`
		}
	} `json:"response"`
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
