package libportal

import (
	"math/rand"
	"time"
)

// PortalUserAccount struct to represent user account for Profile
type PortalUserAccount struct {
	ID string `json:"id" bson:"_id"`

	// identity info
	Profile            string   `json:"profile" bson:"profile"`
	Identity           string   `json:"identity" bson:"identity"`
	AuthenticationCode string   `json:"authentication_code" bson:"authentication_code"`
	MACs               []string `json:"macs" bson:"macs"`

	// basic personal info
	Name    string `json:"name" bson:"name"`
	SurName string `json:"surname" bson:"surname"`

	DateOfBirth       string `json:"date_of_birth" bson:"date_of_birth"`
	DateOfBirthStruct struct {
		Day   int `json:"day" bson:"day"`
		Month int `json:"month" bson:"month"`
		Year  int `json:"year" bson:"year"`
	} `json:"date_of_birth_struct" bson:"date_of_birth_struct"`

	// personal info
	Email    string `json:"email" bson:"email"`
	Phone    string `json:"phone" bson:"phone"`
	Sex      string `json:"sex" bson:"sex"`
	City     string `json:"city" bson:"city"`
	HomeTown string `json:"home_town" bson:"home_town"`
	PhotoURL string `json:"photo_url" bson:"photo_url"`

	// data for external addintional use
	Passport       string      `json:"passport" bson:"passport"`
	AdditionalData interface{} `json:"additional_data" bson:"addition_data"`

	// info from social network after authorizations
	SocialNetwork map[string]AccountFromSocialNetwork `json:"social_network" bson:"social_network"`

	// is filled filled bt user and acceptable for pushes from portal
	Filled        bool `json:"filled" bson:"filled"`
	PushAgreement bool `json:"push_agreement" bson:"push_agreement"`

	// balance for paid internet
	Balance  int    `json:"balance" bson:"balance"`
	Currency string `json:"currency" bson:"currency"`

	// creation time
	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`

	// common statistics
	Visits map[string]int `json:"visits" bson:"visits"`

	// subscription data
	SubscribeID       string `json:"subscribe_id" bson:"subscribe_id"`
	CreateAtSubscribe int64  `json:"create_at_subscribe" bson:"create_at_subscribe"`
}

// PortalUserAccountShort for short represent account
type PortalUserAccountShort struct {
	Profile string `json:"profile" bson:"profile"`

	Identity string   `json:"identity" bson:"identity"`
	MACs     []string `json:"macs" bson:"macs"`

	Name    string `json:"name" bson:"name"`
	SurName string `json:"surname" bson:"surname"`
	Filled  bool   `json:"filled" bson:"filled"`

	PushAgreement bool `json:"push_agreement" bson:"push_agreement"`

	Balance  int    `json:"balance" bson:"balance"`
	Currency string `json:"currency" bson:"currency"`

	CreateAt int64 `json:"create_at" bson:"create_at"`

	SubscribeID       string `json:"subscribe_id"`
	CreateAtSubscribe int64  `json:"create_at_subscribe"`
}

// PortalUserVoucher struct to represent voucher
type PortalUserVoucher struct {
	ID string `json:"id" bson:"_id"`

	Account string `json:"account" bson:"account"`
	Profile string `json:"profile" bson:"profile"`

	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`

	StartAt  int64 `json:"start_at" bson:"start_at"`
	ExpireAt int64 `json:"expire_at" bson:"expire_at"`

	Code string `json:"code" bson:"code"`
	Used bool   `json:"used" bson:"used"`

	Identity string `json:"identity" bson:"identity"`
	Location string `json:"location" bson:"location"`

	// limits
	SpeedLimit   int   `json:"speed" bson:"speed"`
	TimeoutLimit int64 `json:"session" bson:"session"`
	TrafficLimit int   `json:"traffic" bson:"traffic"`

	// not needed anymore
	Plan string `json:"tariff" bson:"tariff"`
}

// PortalTariffPlan struct to represent paid plans (tariffs)
type PortalTariffPlan struct {
	ID string `json:"id" bson:"_id"`

	Profile string `json:"profile" bson:"profile"`

	// tariff plan code be for vouchers or for subscriptions
	Type TariffPlanType `json:"type" bson:"type"`

	// name of plan
	Name string `json:"name" bson:"name"`

	// limits
	SpeedLimit   int   `json:"speed" bson:"speed"`
	TimeoutLimit int64 `json:"session" bson:"session"`
	TrafficLimit int   `json:"traffic" bson:"traffic"`

	// how much to pay
	Amount   int    `json:"amount" bson:"amount"`
	Currency string `json:"currency" bson:"currency"`

	// data for subscriptions
	Reccuring         bool   `json:"recurring" bson:"recurring"`
	ReccuringInterval string `json:"recurring_period" bson:"recurring_period"`

	// friendly description on different languages
	Description map[Locale]string `json:"description" bson:"description"`

	// service info
	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
	Location string    `json:"location" bson:"location"`
}

const (
	letterBytes           = "abcdefghijklmnopqrstuvwxyz1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	voucherMinLength      = 6
	hotelVoucherMinLength = 4
)

// GenerateVoucher function to generate unique
// alphanumeric voucher code with provided len.
// Example: GenerateVoucher(6) -> a1b-2c3
func GenerateVoucher(length int) string {
	if length < voucherMinLength {
		length = voucherMinLength
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

// GenerateHotelVoucher function to generate unique
// alphanumeric voucher code with provided len.
// Example: GenerateVoucher(6) -> a1b-2c3
func GenerateHotelVoucher(length int) string {
	if length < hotelVoucherMinLength {
		length = hotelVoucherMinLength
	}

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

// PortalPaymentSystem struct for represent PS
// (Humo, Megafon Life, Robokassa, etc)
type PortalPaymentSystem struct {
	ID string `json:"id" bson:"_id"`

	Name string                  `json:"name" bson:"name"`
	Type PortalPaymentSystemType `json:"type" bson:"type"`

	// for External
	Merchant   string `json:"merchant" bson:"merchant"`
	Profile    string `json:"profile" bson:"profile"`
	TranUnique bool   `json:"tran_unique" bson:"tran_unique"`

	// for buttons / guides
	Image string `json:"image" bson:"image"`
	Text  string `json:"text" bson:"text"`
	Head  string `json:"head" bson:"head"`

	// only for ExternalURL
	TemplateURL  string `json:"template_url" bson:"template_url"`
	TemplateHash string `json:"template_hash" bson:"template_hash"`
	HashKey      string `json:"hash_key" bson:"hash_key"`

	// will be deprecated
	Identity string `json:"identity" bson:"identity"`
}

// PortalTransaction struct for represent balance transactions
type PortalTransaction struct {
	ID string `json:"id" bson:"_id"`

	Profile string `json:"profile" bson:"profile"`

	Account  string `json:"account" bson:"account"`
	Identity string `json:"identity" bson:"identity"`

	Value    int    `json:"balance" bson:"balance"`
	Currency string `json:"currency" bson:"currency"`

	Type string `json:"type" bson:"type"`
	Fill bool   `json:"fill" bson:"fill"`

	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`
}

type AccountFromSocialNetwork struct {
	ID           string   `json:"id" bson:"id"`
	Name         string   `json:"first_name" bson:"first_name"`
	LastName     string   `json:"last_name" bson:"last_name"`
	Email        string   `json:"email" bson:"email"`
	DateOfBirth  string   `json:"birthday" bson:"birthday"`
	Sex          string   `json:"sex" bson:"sex"`
	City         string   `json:"city" bson:"city"`
	Universities []string `json:"universities" bson:"universities"`
	HomeTown     string   `json:"home_town" bson:"home_town"`
	PhotoURL     string   `json:"photo_url" bson:"photo_url"`
}

// PortalUserHotelVoucher struct to represent voucher
type PortalUserHotelVoucher struct {
	ID string `json:"id" bson:"_id"`

	Account string `json:"account" bson:"account"`
	Profile string `json:"profile" bson:"profile"`

	Create   time.Time `json:"create" bson:"create"`
	CreateAt int64     `json:"create_at" bson:"create_at"`

	StartAt  int64 `json:"start_at" bson:"start_at"`
	ExpireAt int64 `json:"expire_at" bson:"expire_at"`

	Code string `json:"code" bson:"code"`
	Used bool   `json:"used" bson:"used"`

	// limits
	SpeedLimit   int   `json:"speed" bson:"speed"`
	TimeoutLimit int64 `json:"session" bson:"session"`
	TrafficLimit int   `json:"traffic" bson:"traffic"`
}
