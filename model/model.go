package model

import "time"

type LoginInfo struct {
	UserID string `json:"user_id"`
}
type LoginCookie struct {
	DeviceToken  string
	RefreshToken string
	Token        string
}
type OrderHistory struct {
	StatusCode int `json:"status_code"`
	Data       struct {
		TotalCount string `json:"total_count"`
		Items      []struct {
			CurrentStatus struct {
				Code               int    `json:"code"`
				Message            string `json:"message"`
				Type               string `json:"type"`
				ChangedAt          Time   `json:"changedAt"`
				InternalStatusCode string `json:"internal_status_code"`
			} `json:"current_status"`
			OrderAddress    string `json:"order_address"`
			DeliveryAddress struct {
				ID                   int     `json:"id"`
				City                 string  `json:"city"`
				AddressLine1         string  `json:"address_line_1"`
				AddressLine2         string  `json:"address_line_2"`
				AddressLine3         string  `json:"address_line_3"`
				AddressLine4         string  `json:"address_line_4"`
				AddressLine5         string  `json:"address_line_5"`
				AddressOther         string  `json:"address_other"`
				Room                 string  `json:"room"`
				FlatNumber           string  `json:"flat_number"`
				Structure            string  `json:"structure"`
				Building             string  `json:"building"`
				Floor                string  `json:"floor"`
				District             string  `json:"district"`
				Postcode             string  `json:"postcode"`
				Company              string  `json:"company"`
				Latitude             float64 `json:"latitude"`
				Longitude            float64 `json:"longitude"`
				DeliveryInstructions string  `json:"delivery_instructions"`
			} `json:"delivery_address"`
			OrderID               int    `json:"order_id"`
			OrderCode             string `json:"order_code"`
			OrderedAt             Time   `json:"ordered_at"`
			ConfirmedDeliveryTime Time   `json:"confirmed_delivery_time"`
			TotalValue            int    `json:"total_value"`
			Subtotal              int    `json:"subtotal"`
			Vendor                struct {
				ID                         int     `json:"id"`
				Code                       string  `json:"code"`
				Name                       string  `json:"name"`
				Logo                       string  `json:"logo"`
				ServiceTaxPercentageAmount int     `json:"service_tax_percentage_amount"`
				ServiceFeePercentageAmount int     `json:"service_fee_percentage_amount"`
				VatPercentageAmount        int     `json:"vat_percentage_amount"`
				IsServiceFeeEnabled        bool    `json:"is_service_fee_enabled"`
				IsServiceTaxVisible        bool    `json:"is_service_tax_visible"`
				IsServiceTaxEnabled        bool    `json:"is_service_tax_enabled"`
				IsVatVisible               bool    `json:"is_vat_visible"`
				IsVatDisabled              bool    `json:"is_vat_disabled"`
				Address                    string  `json:"address"`
				Latitude                   float64 `json:"latitude"`
				Longitude                  float64 `json:"longitude"`
				PrimaryCuisineID           int     `json:"primary_cuisine_id"`
				Metadata                   struct {
					Timezone string `json:"timezone"`
				} `json:"metadata"`
				CityID           int    `json:"city_id"`
				Vertical         string `json:"vertical"`
				HeroListingImage string `json:"hero_listing_image"`
			} `json:"vendor"`
			Voucher       interface{} `json:"voucher"`
			OrderProducts []struct {
				Name                          string        `json:"name"`
				NameAttributes                Attributes    `json:"name_attributes"`
				QuantityAttributes            Attributes    `json:"quantity_attributes"`
				ToppingsAttributes            *Attributes   `json:"toppings_attributes"`
				SpecialInstructionsAttributes Attributes    `json:"special_instructions_attributes"`
				SubtitleAttributes            *Attributes   `json:"subtitle_attributes"`
				PriceAttributes               Attributes    `json:"price_attributes"`
				TotalPrice                    int           `json:"total_price"`
				Quantity                      int           `json:"quantity"`
				SoldOut                       bool          `json:"sold_out"`
				Tag                           interface{}   `json:"tag"`
				Status                        string        `json:"status"`
				WeightAttributes              interface{}   `json:"weight_attributes"`
				ReplacementProducts           []interface{} `json:"replacement_products"`
				ReplacementAttributes         interface{}   `json:"replacement_attributes"`
				IsMealForOne                  bool          `json:"is_meal_for_one"`
			} `json:"order_products"`
			SMSVerificationNeeded      bool    `json:"sms_verification_needed"`
			OrderAlreadyRated          bool    `json:"order_already_rated"`
			HasBeenReportedLate        bool    `json:"has_been_reported_late"`
			RiderTip                   int     `json:"rider_tip"`
			Charity                    int     `json:"charity"`
			ServiceFeeTotal            float64 `json:"service_fee_total"`
			DifferenceToMinimum        int     `json:"difference_to_minimum"`
			DifferenceToMinimumWithVat int     `json:"difference_to_minimum_with_vat"`
			ExpeditionType             string  `json:"expedition_type"`
			ServerTime                 Time    `json:"server_time"`
			PaymentTypeCode            string  `json:"payment_type_code"`
			Payment                    struct {
				Breakdown []struct {
					Code           string `json:"code"`
					Group          string `json:"group"`
					TranslationKey string `json:"translation_key"`
					Amount         int    `json:"amount"`
				} `json:"breakdown"`
			} `json:"payment"`
			DeliveryFeatures struct {
				ShowMap                bool `json:"show_map"`
				DisplayMap             bool `json:"display_map"`
				ShowCart               bool `json:"show_cart"`
				ShowStatuses           bool `json:"show_statuses"`
				ShowVendorContact      bool `json:"show_vendor_contact"`
				ShowVendorDeliveryTime bool `json:"show_vendor_delivery_time"`
				ShowRiderChat          bool `json:"show_rider_chat"`
				ShowPaymentBreakdown   bool `json:"show_payment_breakdown"`
				ShowHelpButton         bool `json:"show_help_button"`
			} `json:"delivery_features"`
			DeliveryProvider string `json:"delivery_provider"`
			IsCancellable    bool   `json:"is_cancellable"`
			StatusFlags      struct {
				IsPaid                         bool `json:"is_paid"`
				IsActive                       bool `json:"is_active"`
				IsPickedUp                     bool `json:"is_picked_up"`
				IsDelivered                    bool `json:"is_delivered"`
				IsCanceled                     bool `json:"is_canceled"`
				IsPreorder                     bool `json:"is_preorder"`
				IsPreorderActive               bool `json:"is_preorder_active"`
				IsPastOrder                    bool `json:"is_past_order"`
				IsCompleted                    bool `json:"is_completed"`
				IsReorderable                  bool `json:"is_reorderable"`
				IsReorderableAfterCancellation bool `json:"is_reorderable_after_cancellation"`
				IsRateable                     bool `json:"is_rateable"`
				IsMealForOne                   bool `json:"is_meal_for_one"`
			} `json:"status_flags"`
			DynamicFees []struct {
				TranslationKey string  `json:"translation_key"`
				Value          float64 `json:"value"`
				Style          string  `json:"style"`
				TextColor      string  `json:"text_color"`
				Tag            *struct {
					Type string `json:"type"`
					Text string `json:"text"`
				} `json:"tag"`
				InitialValue *int `json:"initial_value,omitempty"`
				Metadata     *struct {
					ShowIconType string `json:"show_icon_type"`
				} `json:"metadata,omitempty"`
				Placeholder struct{} `json:"placeholder"`
			} `json:"dynamic_fees"`
			Loyalty        interface{} `json:"loyalty"`
			StatusMessages struct {
				Titles []interface{} `json:"titles"`
			} `json:"status_messages"`
			StatusHistory  []interface{} `json:"status_history"`
			PaymentRefunds []interface{} `json:"payment_refunds"`
		} `json:"items"`
		Features struct {
			LegacyOrdersAvailable bool `json:"legacy_orders_available"`
			OosEnabled            bool `json:"oos_enabled"`
		} `json:"features"`
	} `json:"data"`
}

type Time struct {
	Date     string `json:"date"`
	Timezone string `json:"timezone"`
}

type Attributes struct {
	Style        string      `json:"style"`
	TextColor    string      `json:"text_color"`
	Value        interface{} `json:"value"`
	Placeholders interface{} `json:"placeholders"`
}

// order

type Order struct {
	StatusCode int  `json:"status_code"`
	Data       Data `json:"data"`
}
type ChangedAt struct {
	Date               string    `json:"date"`
	Timezone           string    `json:"timezone"`
	DateTimezoneOffset time.Time `json:"date_timezone_offset"`
}
type StatusHistory struct {
	Code               int       `json:"code"`
	Message            string    `json:"message"`
	Type               string    `json:"type"`
	ChangedAt          ChangedAt `json:"changedAt"`
	InternalStatusCode string    `json:"internal_status_code"`
	DeclineReason      any       `json:"decline_reason"`
}
type CurrentStatus struct {
	Code               int       `json:"code"`
	Message            string    `json:"message"`
	Type               string    `json:"type"`
	ChangedAt          ChangedAt `json:"changedAt"`
	InternalStatusCode string    `json:"internal_status_code"`
	DeclineReason      any       `json:"decline_reason"`
}
type DeliveryAddress struct {
	ID                   int     `json:"id"`
	City                 string  `json:"city"`
	AddressLine1         string  `json:"address_line_1"`
	AddressLine2         string  `json:"address_line_2"`
	AddressLine3         string  `json:"address_line_3"`
	AddressLine4         string  `json:"address_line_4"`
	AddressLine5         string  `json:"address_line_5"`
	AddressOther         string  `json:"address_other"`
	Room                 string  `json:"room"`
	FlatNumber           string  `json:"flat_number"`
	Structure            string  `json:"structure"`
	Building             string  `json:"building"`
	Floor                string  `json:"floor"`
	District             string  `json:"district"`
	Postcode             string  `json:"postcode"`
	Company              string  `json:"company"`
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	DeliveryInstructions string  `json:"delivery_instructions"`
	Intercom             string  `json:"intercom"`
	Entrance             string  `json:"entrance"`
	Meta                 string  `json:"meta"`
	Block                string  `json:"block"`
	PhoneNumber          string  `json:"phone_number"`
}
type OrderedAt struct {
	Date               string    `json:"date"`
	Timezone           string    `json:"timezone"`
	DateTimezoneOffset time.Time `json:"date_timezone_offset"`
}
type ConfirmedDeliveryTime struct {
	Date               string    `json:"date"`
	Timezone           string    `json:"timezone"`
	DateTimezoneOffset time.Time `json:"date_timezone_offset"`
}
type Metadata struct {
	Timezone string `json:"timezone"`
}
type Vendor struct {
	ID                         int      `json:"id"`
	Code                       string   `json:"code"`
	Name                       string   `json:"name"`
	Logo                       string   `json:"logo"`
	ServiceTaxPercentageAmount int      `json:"service_tax_percentage_amount"`
	ServiceFeePercentageAmount int      `json:"service_fee_percentage_amount"`
	VatPercentageAmount        int      `json:"vat_percentage_amount"`
	IsServiceFeeEnabled        bool     `json:"is_service_fee_enabled"`
	IsServiceTaxVisible        bool     `json:"is_service_tax_visible"`
	IsServiceTaxEnabled        bool     `json:"is_service_tax_enabled"`
	IsVatVisible               bool     `json:"is_vat_visible"`
	IsVatDisabled              bool     `json:"is_vat_disabled"`
	Address                    string   `json:"address"`
	Latitude                   float64  `json:"latitude"`
	Longitude                  float64  `json:"longitude"`
	PrimaryCuisineID           int      `json:"primary_cuisine_id"`
	Metadata                   Metadata `json:"metadata"`
	CityID                     int      `json:"city_id"`
	Vertical                   string   `json:"vertical"`
	CustomerContactPhoneNumber any      `json:"customer_contact_phone_number"`
	HeroListingImage           string   `json:"hero_listing_image"`
}
type NameAttributes struct {
	Style        string `json:"style"`
	TextColor    string `json:"text_color"`
	Value        string `json:"value"`
	InitialValue any    `json:"initial_value"`
	Placeholders any    `json:"placeholders"`
}
type QuantityAttributes struct {
	Style        string `json:"style"`
	TextColor    string `json:"text_color"`
	Value        int    `json:"value"`
	InitialValue any    `json:"initial_value"`
	Placeholders any    `json:"placeholders"`
}
type ToppingsAttributes struct {
	Style        string `json:"style"`
	TextColor    string `json:"text_color"`
	Value        string `json:"value"`
	InitialValue any    `json:"initial_value"`
	Placeholders any    `json:"placeholders"`
}
type PriceAttributes struct {
	Style        string `json:"style"`
	TextColor    string `json:"text_color"`
	Value        int    `json:"value"`
	InitialValue any    `json:"initial_value"`
	Placeholders any    `json:"placeholders"`
}
type CollapsedAttributes struct {
	Style     string `json:"style"`
	TextColor string `json:"text_color"`
}
type OrderProducts struct {
	Name                          string              `json:"name"`
	NameAttributes                NameAttributes      `json:"name_attributes"`
	QuantityAttributes            QuantityAttributes  `json:"quantity_attributes"`
	ToppingsAttributes            ToppingsAttributes  `json:"toppings_attributes"`
	SpecialInstructionsAttributes any                 `json:"special_instructions_attributes"`
	SubtitleAttributes            any                 `json:"subtitle_attributes"`
	PriceAttributes               PriceAttributes     `json:"price_attributes"`
	TotalPrice                    int                 `json:"total_price"`
	Quantity                      int                 `json:"quantity"`
	SoldOut                       bool                `json:"sold_out"`
	Tag                           any                 `json:"tag"`
	Status                        string              `json:"status"`
	ContainerChargeAttributes     any                 `json:"container_charge_attributes"`
	CollapsedAttributes           CollapsedAttributes `json:"collapsed_attributes"`
	VariableWeightAttributes      any                 `json:"variable_weight_attributes"`
	ReplacementAttributes         any                 `json:"replacement_attributes"`
	ReplacementProducts           []any               `json:"replacement_products"`
	VariationID                   string              `json:"variation_id"`
}
type ServerTime struct {
	Date               string    `json:"date"`
	Timezone           string    `json:"timezone"`
	DateTimezoneOffset time.Time `json:"date_timezone_offset"`
}
type Breakdown struct {
	Code                 string `json:"code"`
	Group                string `json:"group"`
	TranslationKey       string `json:"translation_key"`
	Amount               int    `json:"amount"`
	Status               string `json:"status"`
	StatusTranslationKey string `json:"status_translation_key"`
}
type Payment struct {
	Breakdown []Breakdown `json:"breakdown"`
}
type DeliveryFeatures struct {
	ShowMap                bool `json:"show_map"`
	DisplayMap             bool `json:"display_map"`
	ShowCart               bool `json:"show_cart"`
	ShowStatuses           bool `json:"show_statuses"`
	ShowVendorContact      bool `json:"show_vendor_contact"`
	ShowVendorDeliveryTime bool `json:"show_vendor_delivery_time"`
	ShowRiderChat          bool `json:"show_rider_chat"`
	ShowPaymentBreakdown   bool `json:"show_payment_breakdown"`
	ShowHelpButton         bool `json:"show_help_button"`
	ShowVendorLocation     bool `json:"show_vendor_location"`
}
type DeliveryTimeRange struct {
	Range  string `json:"range"`
	Suffix string `json:"suffix"`
	Label  string `json:"label"`
}
type Titles struct {
	Name                    string `json:"name"`
	TranslationKey          string `json:"translation_key"`
	ImageURL                string `json:"image_url"`
	Active                  bool   `json:"active"`
	BarWeight               int    `json:"bar_weight"`
	IsFilled                bool   `json:"is_filled"`
	MediaURL                string `json:"media_url"`
	MediaType               string `json:"media_type"`
	PrimaryTranslationKey   string `json:"primary_translation_key"`
	SecondaryTranslationKey string `json:"secondary_translation_key"`
}
type StatusMessages struct {
	Subtitle any      `json:"subtitle"`
	Titles   []Titles `json:"titles"`
}
type StatusFlags struct {
	IsPaid                         bool `json:"is_paid"`
	IsActive                       bool `json:"is_active"`
	IsPickedUp                     bool `json:"is_picked_up"`
	IsDelivered                    bool `json:"is_delivered"`
	IsCanceled                     bool `json:"is_canceled"`
	IsPreorder                     bool `json:"is_preorder"`
	IsPreorderActive               bool `json:"is_preorder_active"`
	IsPastOrder                    bool `json:"is_past_order"`
	IsCompleted                    bool `json:"is_completed"`
	IsReorderable                  bool `json:"is_reorderable"`
	IsReorderableAfterCancellation bool `json:"is_reorderable_after_cancellation"`
}
type Placeholder struct {
}
type DynamicFees struct {
	TranslationKey string      `json:"translation_key"`
	Name           any         `json:"name"`
	Value          float64     `json:"value"`
	Style          string      `json:"style"`
	TextColor      string      `json:"text_color"`
	Tag            any         `json:"tag"`
	InitialValue   any         `json:"initial_value"`
	Metadata       any         `json:"metadata"`
	Placeholder    Placeholder `json:"placeholder"`
}
type Placeholders struct {
}
type CourierDetails struct {
	TitleTranslationKey string       `json:"title_translation_key"`
	Placeholders        Placeholders `json:"placeholders"`
	AvatarCount         int          `json:"avatar_count"`
}
type StatusMessage struct {
	PrimaryStatus    string `json:"primary_status"`
	SecondaryStatus  string `json:"secondary_status"`
	ActiveBar        int    `json:"active_bar"`
	BarWeights       []any  `json:"bar_weights"`
	IllustrationCode string `json:"illustration_code"`
}
type DeliveryTime struct {
	Label  string `json:"label"`
	Range  string `json:"range"`
	Suffix string `json:"suffix"`
}
type CompactDeliveryTime struct {
	Range  string `json:"range"`
	Suffix string `json:"suffix"`
}
type LiveActivity struct {
	StatusMessage       StatusMessage       `json:"status_message"`
	DeliveryTime        DeliveryTime        `json:"delivery_time"`
	CompactDeliveryTime CompactDeliveryTime `json:"compact_delivery_time"`
}
type Data struct {
	StatusHistory              []StatusHistory       `json:"status_history"`
	CurrentStatus              CurrentStatus         `json:"current_status"`
	DeclineReason              any                   `json:"decline_reason"`
	CancellationInfo           any                   `json:"cancellation_info"`
	VoucherInfo                any                   `json:"voucher_info"`
	RefundPolicy               any                   `json:"refund_policy"`
	OrderAddress               string                `json:"order_address"`
	DeliveryAddress            DeliveryAddress       `json:"delivery_address"`
	OrderID                    int                   `json:"order_id"`
	OrderCode                  string                `json:"order_code"`
	OrderedAt                  OrderedAt             `json:"ordered_at"`
	ConfirmedDeliveryTime      ConfirmedDeliveryTime `json:"confirmed_delivery_time"`
	TotalValue                 int                   `json:"total_value"`
	Subtotal                   int                   `json:"subtotal"`
	DeliveryFee                int                   `json:"delivery_fee"`
	Vendor                     Vendor                `json:"vendor"`
	Voucher                    any                   `json:"voucher"`
	OrderProducts              []OrderProducts       `json:"order_products"`
	SmsVerificationNeeded      bool                  `json:"sms_verification_needed"`
	OrderAlreadyRated          bool                  `json:"order_already_rated"`
	OrderRating                any                   `json:"order_rating"`
	HasBeenReportedLate        bool                  `json:"has_been_reported_late"`
	RiderTip                   int                   `json:"rider_tip"`
	Charity                    int                   `json:"charity"`
	ServiceFeeTotal            float64               `json:"service_fee_total"`
	DifferenceToMinimum        int                   `json:"difference_to_minimum"`
	DifferenceToMinimumWithVat int                   `json:"difference_to_minimum_with_vat"`
	ExpeditionType             string                `json:"expedition_type"`
	ServerTime                 ServerTime            `json:"server_time"`
	PaymentTypeCode            string                `json:"payment_type_code"`
	Payment                    Payment               `json:"payment"`
	Courier                    any                   `json:"courier"`
	ActiveCouriers             []any                 `json:"active_couriers"`
	DeliveryFeatures           DeliveryFeatures      `json:"delivery_features"`
	DeliveryProvider           string                `json:"delivery_provider"`
	IsCancellable              bool                  `json:"is_cancellable"`
	Eta                        int                   `json:"eta"`
	DeliveryTimeRange          DeliveryTimeRange     `json:"delivery_time_range"`
	StatusMessages             StatusMessages        `json:"status_messages"`
	StatusFlags                StatusFlags           `json:"status_flags"`
	DynamicFees                []DynamicFees         `json:"dynamic_fees"`
	Delay                      any                   `json:"delay"`
	Loyalty                    any                   `json:"loyalty"`
	OrderShortCode             int                   `json:"order_short_code"`
	StackedOrdersCount         int                   `json:"stacked_orders_count"`
	OrderExtra                 any                   `json:"order_extra"`
	Analytics                  any                   `json:"analytics"`
	Tracking                   any                   `json:"tracking"`
	IsDeliveryAddressUpdated   bool                  `json:"is_delivery_address_updated"`
	LastDeliveryAddress        any                   `json:"last_delivery_address"`
	ReplacementInfo            any                   `json:"replacement_info"`
	Refund                     any                   `json:"refund"`
	InitialTotalValue          any                   `json:"initial_total_value"`
	PriorityDelivery           any                   `json:"priority_delivery"`
	CourierDetails             CourierDetails        `json:"courier_details"`
	VendorContactDetails       any                   `json:"vendor_contact_details"`
	LiveActivity               LiveActivity          `json:"live_activity"`
	Map                        any                   `json:"map"`
}
