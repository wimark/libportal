package libportal

import (
	"encoding/json"
	"errors"
	"github.com/globalsign/mgo/bson"
)

type AccessServerType string

const AccessServerTypeCiscoISG AccessServerType = "Cisco ISG"
const AccessServerTypeCiscoWLC AccessServerType = "Cisco WLC"
const AccessServerTypeCoovaChilly AccessServerType = "CoovaChilli"
const AccessServerTypeEcoBRAS AccessServerType = "EcoBRAS"
const AccessServerTypeWimark AccessServerType = "Wimark"

func (en AccessServerType) GetPtr() *AccessServerType { var v = en; return &v }

func (en AccessServerType) String() string {
	switch en {
	case AccessServerTypeCiscoISG:
		return "Cisco ISG"
	case AccessServerTypeCiscoWLC:
		return "Cisco WLC"
	case AccessServerTypeCoovaChilly:
		return "CoovaChilli"
	case AccessServerTypeEcoBRAS:
		return "EcoBRAS"
	case AccessServerTypeWimark:
		return "Wimark"
	}
	if len(en) == 0 {
		return "Wimark"
	}
	panic(errors.New("Invalid value of AccessServerType: " + string(en)))
}

func (en *AccessServerType) MarshalJSON() ([]byte, error) {
	switch *en {
	case AccessServerTypeCiscoISG:
		return json.Marshal("Cisco ISG")
	case AccessServerTypeCiscoWLC:
		return json.Marshal("Cisco WLC")
	case AccessServerTypeCoovaChilly:
		return json.Marshal("CoovaChilli")
	case AccessServerTypeEcoBRAS:
		return json.Marshal("EcoBRAS")
	case AccessServerTypeWimark:
		return json.Marshal("Wimark")
	}
	if len(*en) == 0 {
		return json.Marshal("Wimark")
	}
	return nil, errors.New("Invalid value of AccessServerType: " + string(*en))
}

func (en *AccessServerType) GetBSON() (interface{}, error) {
	switch *en {
	case AccessServerTypeCiscoISG:
		return "Cisco ISG", nil
	case AccessServerTypeCiscoWLC:
		return "Cisco WLC", nil
	case AccessServerTypeCoovaChilly:
		return "CoovaChilli", nil
	case AccessServerTypeEcoBRAS:
		return "EcoBRAS", nil
	case AccessServerTypeWimark:
		return "Wimark", nil
	}
	if len(*en) == 0 {
		return "Wimark", nil
	}
	return nil, errors.New("Invalid value of AccessServerType: " + string(*en))
}

func (en *AccessServerType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "Cisco ISG":
		*en = AccessServerTypeCiscoISG
		return nil
	case "Cisco WLC":
		*en = AccessServerTypeCiscoWLC
		return nil
	case "CoovaChilli":
		*en = AccessServerTypeCoovaChilly
		return nil
	case "EcoBRAS":
		*en = AccessServerTypeEcoBRAS
		return nil
	case "Wimark":
		*en = AccessServerTypeWimark
		return nil
	}
	if len(s) == 0 {
		*en = AccessServerTypeWimark
		return nil
	}
	return errors.New("Unknown AccessServerType: " + s)
}

func (en *AccessServerType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "Cisco ISG":
		*en = AccessServerTypeCiscoISG
		return nil
	case "Cisco WLC":
		*en = AccessServerTypeCiscoWLC
		return nil
	case "CoovaChilli":
		*en = AccessServerTypeCoovaChilly
		return nil
	case "EcoBRAS":
		*en = AccessServerTypeEcoBRAS
		return nil
	case "Wimark":
		*en = AccessServerTypeWimark
		return nil
	}
	if len(s) == 0 {
		*en = AccessServerTypeWimark
		return nil
	}
	return errors.New("Unknown AccessServerType: " + s)
}

type Locale string

const LocaleEn Locale = "en"
const LocaleRu Locale = "ru"
const LocaleTg Locale = "tg"

func (en Locale) GetPtr() *Locale { var v = en; return &v }

func (en Locale) String() string {
	switch en {
	case LocaleEn:
		return "en"
	case LocaleRu:
		return "ru"
	case LocaleTg:
		return "tg"
	}
	if len(en) == 0 {
		return "ru"
	}
	panic(errors.New("Invalid value of Locale: " + string(en)))
}

func (en *Locale) MarshalJSON() ([]byte, error) {
	switch *en {
	case LocaleEn:
		return json.Marshal("en")
	case LocaleRu:
		return json.Marshal("ru")
	case LocaleTg:
		return json.Marshal("tg")
	}
	if len(*en) == 0 {
		return json.Marshal("ru")
	}
	return nil, errors.New("Invalid value of Locale: " + string(*en))
}

func (en *Locale) GetBSON() (interface{}, error) {
	switch *en {
	case LocaleEn:
		return "en", nil
	case LocaleRu:
		return "ru", nil
	case LocaleTg:
		return "tg", nil
	}
	if len(*en) == 0 {
		return "ru", nil
	}
	return nil, errors.New("Invalid value of Locale: " + string(*en))
}

func (en *Locale) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "en":
		*en = LocaleEn
		return nil
	case "ru":
		*en = LocaleRu
		return nil
	case "tg":
		*en = LocaleTg
		return nil
	}
	if len(s) == 0 {
		*en = LocaleRu
		return nil
	}
	return errors.New("Unknown Locale: " + s)
}

func (en *Locale) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "en":
		*en = LocaleEn
		return nil
	case "ru":
		*en = LocaleRu
		return nil
	case "tg":
		*en = LocaleTg
		return nil
	}
	if len(s) == 0 {
		*en = LocaleRu
		return nil
	}
	return errors.New("Unknown Locale: " + s)
}

type PortalActionListType string

const PortalActionListTypeAddToAccessList PortalActionListType = "add_al"
const PortalActionListTypeAddToBlackList PortalActionListType = "add_bl"
const PortalActionListTypeDelFromAccessList PortalActionListType = "del_al"
const PortalActionListTypeDelFromBlackList PortalActionListType = "del_bl"
const PortalActionListTypeGetAccessList PortalActionListType = "get_al"
const PortalActionListTypeGetBlackList PortalActionListType = "get_bl"
const PortalActionListTypeNone PortalActionListType = "none"

func (en PortalActionListType) GetPtr() *PortalActionListType { var v = en; return &v }

func (en PortalActionListType) String() string {
	switch en {
	case PortalActionListTypeAddToAccessList:
		return "add_al"
	case PortalActionListTypeAddToBlackList:
		return "add_bl"
	case PortalActionListTypeDelFromAccessList:
		return "del_al"
	case PortalActionListTypeDelFromBlackList:
		return "del_bl"
	case PortalActionListTypeGetAccessList:
		return "get_al"
	case PortalActionListTypeGetBlackList:
		return "get_bl"
	case PortalActionListTypeNone:
		return "none"
	}
	if len(en) == 0 {
		return "none"
	}
	panic(errors.New("Invalid value of PortalActionListType: " + string(en)))
}

func (en *PortalActionListType) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalActionListTypeAddToAccessList:
		return json.Marshal("add_al")
	case PortalActionListTypeAddToBlackList:
		return json.Marshal("add_bl")
	case PortalActionListTypeDelFromAccessList:
		return json.Marshal("del_al")
	case PortalActionListTypeDelFromBlackList:
		return json.Marshal("del_bl")
	case PortalActionListTypeGetAccessList:
		return json.Marshal("get_al")
	case PortalActionListTypeGetBlackList:
		return json.Marshal("get_bl")
	case PortalActionListTypeNone:
		return json.Marshal("none")
	}
	if len(*en) == 0 {
		return json.Marshal("none")
	}
	return nil, errors.New("Invalid value of PortalActionListType: " + string(*en))
}

func (en *PortalActionListType) GetBSON() (interface{}, error) {
	switch *en {
	case PortalActionListTypeAddToAccessList:
		return "add_al", nil
	case PortalActionListTypeAddToBlackList:
		return "add_bl", nil
	case PortalActionListTypeDelFromAccessList:
		return "del_al", nil
	case PortalActionListTypeDelFromBlackList:
		return "del_bl", nil
	case PortalActionListTypeGetAccessList:
		return "get_al", nil
	case PortalActionListTypeGetBlackList:
		return "get_bl", nil
	case PortalActionListTypeNone:
		return "none", nil
	}
	if len(*en) == 0 {
		return "none", nil
	}
	return nil, errors.New("Invalid value of PortalActionListType: " + string(*en))
}

func (en *PortalActionListType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "add_al":
		*en = PortalActionListTypeAddToAccessList
		return nil
	case "add_bl":
		*en = PortalActionListTypeAddToBlackList
		return nil
	case "del_al":
		*en = PortalActionListTypeDelFromAccessList
		return nil
	case "del_bl":
		*en = PortalActionListTypeDelFromBlackList
		return nil
	case "get_al":
		*en = PortalActionListTypeGetAccessList
		return nil
	case "get_bl":
		*en = PortalActionListTypeGetBlackList
		return nil
	case "none":
		*en = PortalActionListTypeNone
		return nil
	}
	if len(s) == 0 {
		*en = PortalActionListTypeNone
		return nil
	}
	return errors.New("Unknown PortalActionListType: " + s)
}

func (en *PortalActionListType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "add_al":
		*en = PortalActionListTypeAddToAccessList
		return nil
	case "add_bl":
		*en = PortalActionListTypeAddToBlackList
		return nil
	case "del_al":
		*en = PortalActionListTypeDelFromAccessList
		return nil
	case "del_bl":
		*en = PortalActionListTypeDelFromBlackList
		return nil
	case "get_al":
		*en = PortalActionListTypeGetAccessList
		return nil
	case "get_bl":
		*en = PortalActionListTypeGetBlackList
		return nil
	case "none":
		*en = PortalActionListTypeNone
		return nil
	}
	if len(s) == 0 {
		*en = PortalActionListTypeNone
		return nil
	}
	return errors.New("Unknown PortalActionListType: " + s)
}

type PortalAdvertisementState string

const PortalAdvertisementStateChecked PortalAdvertisementState = "checked"
const PortalAdvertisementStateNeed PortalAdvertisementState = "none"

func (en PortalAdvertisementState) GetPtr() *PortalAdvertisementState { var v = en; return &v }

func (en PortalAdvertisementState) String() string {
	switch en {
	case PortalAdvertisementStateChecked:
		return "checked"
	case PortalAdvertisementStateNeed:
		return "none"
	}
	if len(en) == 0 {
		return "none"
	}
	panic(errors.New("Invalid value of PortalAdvertisementState: " + string(en)))
}

func (en *PortalAdvertisementState) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalAdvertisementStateChecked:
		return json.Marshal("checked")
	case PortalAdvertisementStateNeed:
		return json.Marshal("none")
	}
	if len(*en) == 0 {
		return json.Marshal("none")
	}
	return nil, errors.New("Invalid value of PortalAdvertisementState: " + string(*en))
}

func (en *PortalAdvertisementState) GetBSON() (interface{}, error) {
	switch *en {
	case PortalAdvertisementStateChecked:
		return "checked", nil
	case PortalAdvertisementStateNeed:
		return "none", nil
	}
	if len(*en) == 0 {
		return "none", nil
	}
	return nil, errors.New("Invalid value of PortalAdvertisementState: " + string(*en))
}

func (en *PortalAdvertisementState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "checked":
		*en = PortalAdvertisementStateChecked
		return nil
	case "none":
		*en = PortalAdvertisementStateNeed
		return nil
	}
	if len(s) == 0 {
		*en = PortalAdvertisementStateNeed
		return nil
	}
	return errors.New("Unknown PortalAdvertisementState: " + s)
}

func (en *PortalAdvertisementState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "checked":
		*en = PortalAdvertisementStateChecked
		return nil
	case "none":
		*en = PortalAdvertisementStateNeed
		return nil
	}
	if len(s) == 0 {
		*en = PortalAdvertisementStateNeed
		return nil
	}
	return errors.New("Unknown PortalAdvertisementState: " + s)
}

type PortalAdvertisementType string

const PortalAdvertisementTypeFlash PortalAdvertisementType = "flash"
const PortalAdvertisementTypeIframe PortalAdvertisementType = "iframe"
const PortalAdvertisementTypeImage PortalAdvertisementType = "image"
const PortalAdvertisementTypePoll PortalAdvertisementType = "poll"
const PortalAdvertisementTypePollIm PortalAdvertisementType = "poll_image"
const PortalAdvertisementTypePollQuality PortalAdvertisementType = "poll_quality"
const PortalAdvertisementTypePollUserData PortalAdvertisementType = "poll_user_data"
const PortalAdvertisementTypeVideo PortalAdvertisementType = "video"

func (en PortalAdvertisementType) GetPtr() *PortalAdvertisementType { var v = en; return &v }

func (en PortalAdvertisementType) String() string {
	switch en {
	case PortalAdvertisementTypeFlash:
		return "flash"
	case PortalAdvertisementTypeIframe:
		return "iframe"
	case PortalAdvertisementTypeImage:
		return "image"
	case PortalAdvertisementTypePoll:
		return "poll"
	case PortalAdvertisementTypePollIm:
		return "poll_image"
	case PortalAdvertisementTypePollQuality:
		return "poll_quality"
	case PortalAdvertisementTypePollUserData:
		return "poll_user_data"
	case PortalAdvertisementTypeVideo:
		return "video"
	}
	if len(en) == 0 {
		return "image"
	}
	panic(errors.New("Invalid value of PortalAdvertisementType: " + string(en)))
}

func (en *PortalAdvertisementType) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalAdvertisementTypeFlash:
		return json.Marshal("flash")
	case PortalAdvertisementTypeIframe:
		return json.Marshal("iframe")
	case PortalAdvertisementTypeImage:
		return json.Marshal("image")
	case PortalAdvertisementTypePoll:
		return json.Marshal("poll")
	case PortalAdvertisementTypePollIm:
		return json.Marshal("poll_image")
	case PortalAdvertisementTypePollQuality:
		return json.Marshal("poll_quality")
	case PortalAdvertisementTypePollUserData:
		return json.Marshal("poll_user_data")
	case PortalAdvertisementTypeVideo:
		return json.Marshal("video")
	}
	if len(*en) == 0 {
		return json.Marshal("image")
	}
	return nil, errors.New("Invalid value of PortalAdvertisementType: " + string(*en))
}

func (en *PortalAdvertisementType) GetBSON() (interface{}, error) {
	switch *en {
	case PortalAdvertisementTypeFlash:
		return "flash", nil
	case PortalAdvertisementTypeIframe:
		return "iframe", nil
	case PortalAdvertisementTypeImage:
		return "image", nil
	case PortalAdvertisementTypePoll:
		return "poll", nil
	case PortalAdvertisementTypePollIm:
		return "poll_image", nil
	case PortalAdvertisementTypePollQuality:
		return "poll_quality", nil
	case PortalAdvertisementTypePollUserData:
		return "poll_user_data", nil
	case PortalAdvertisementTypeVideo:
		return "video", nil
	}
	if len(*en) == 0 {
		return "image", nil
	}
	return nil, errors.New("Invalid value of PortalAdvertisementType: " + string(*en))
}

func (en *PortalAdvertisementType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "flash":
		*en = PortalAdvertisementTypeFlash
		return nil
	case "iframe":
		*en = PortalAdvertisementTypeIframe
		return nil
	case "image":
		*en = PortalAdvertisementTypeImage
		return nil
	case "poll":
		*en = PortalAdvertisementTypePoll
		return nil
	case "poll_image":
		*en = PortalAdvertisementTypePollIm
		return nil
	case "poll_quality":
		*en = PortalAdvertisementTypePollQuality
		return nil
	case "poll_user_data":
		*en = PortalAdvertisementTypePollUserData
		return nil
	case "video":
		*en = PortalAdvertisementTypeVideo
		return nil
	}
	if len(s) == 0 {
		*en = PortalAdvertisementTypeImage
		return nil
	}
	return errors.New("Unknown PortalAdvertisementType: " + s)
}

func (en *PortalAdvertisementType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "flash":
		*en = PortalAdvertisementTypeFlash
		return nil
	case "iframe":
		*en = PortalAdvertisementTypeIframe
		return nil
	case "image":
		*en = PortalAdvertisementTypeImage
		return nil
	case "poll":
		*en = PortalAdvertisementTypePoll
		return nil
	case "poll_image":
		*en = PortalAdvertisementTypePollIm
		return nil
	case "poll_quality":
		*en = PortalAdvertisementTypePollQuality
		return nil
	case "poll_user_data":
		*en = PortalAdvertisementTypePollUserData
		return nil
	case "video":
		*en = PortalAdvertisementTypeVideo
		return nil
	}
	if len(s) == 0 {
		*en = PortalAdvertisementTypeImage
		return nil
	}
	return errors.New("Unknown PortalAdvertisementType: " + s)
}

type PortalAuthenticationState string

const PortalAuthenticationStateChecked PortalAuthenticationState = "checked"
const PortalAuthenticationStateNeed PortalAuthenticationState = "need"
const PortalAuthenticationStateSent PortalAuthenticationState = "sent"

func (en PortalAuthenticationState) GetPtr() *PortalAuthenticationState { var v = en; return &v }

func (en PortalAuthenticationState) String() string {
	switch en {
	case PortalAuthenticationStateChecked:
		return "checked"
	case PortalAuthenticationStateNeed:
		return "need"
	case PortalAuthenticationStateSent:
		return "sent"
	}
	if len(en) == 0 {
		return "need"
	}
	panic(errors.New("Invalid value of PortalAuthenticationState: " + string(en)))
}

func (en *PortalAuthenticationState) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalAuthenticationStateChecked:
		return json.Marshal("checked")
	case PortalAuthenticationStateNeed:
		return json.Marshal("need")
	case PortalAuthenticationStateSent:
		return json.Marshal("sent")
	}
	if len(*en) == 0 {
		return json.Marshal("need")
	}
	return nil, errors.New("Invalid value of PortalAuthenticationState: " + string(*en))
}

func (en *PortalAuthenticationState) GetBSON() (interface{}, error) {
	switch *en {
	case PortalAuthenticationStateChecked:
		return "checked", nil
	case PortalAuthenticationStateNeed:
		return "need", nil
	case PortalAuthenticationStateSent:
		return "sent", nil
	}
	if len(*en) == 0 {
		return "need", nil
	}
	return nil, errors.New("Invalid value of PortalAuthenticationState: " + string(*en))
}

func (en *PortalAuthenticationState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "checked":
		*en = PortalAuthenticationStateChecked
		return nil
	case "need":
		*en = PortalAuthenticationStateNeed
		return nil
	case "sent":
		*en = PortalAuthenticationStateSent
		return nil
	}
	if len(s) == 0 {
		*en = PortalAuthenticationStateNeed
		return nil
	}
	return errors.New("Unknown PortalAuthenticationState: " + s)
}

func (en *PortalAuthenticationState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "checked":
		*en = PortalAuthenticationStateChecked
		return nil
	case "need":
		*en = PortalAuthenticationStateNeed
		return nil
	case "sent":
		*en = PortalAuthenticationStateSent
		return nil
	}
	if len(s) == 0 {
		*en = PortalAuthenticationStateNeed
		return nil
	}
	return errors.New("Unknown PortalAuthenticationState: " + s)
}

type PortalAuthenticationType string

const PortalAuthenticationTypeCallFront PortalAuthenticationType = "callfront"
const PortalAuthenticationTypeCallback PortalAuthenticationType = "callback"
const PortalAuthenticationTypeESIA PortalAuthenticationType = "esia"
const PortalAuthenticationTypeEmail PortalAuthenticationType = "email"
const PortalAuthenticationTypeHotelVoucher PortalAuthenticationType = "hotel_voucher"
const PortalAuthenticationTypeNone PortalAuthenticationType = "none"
const PortalAuthenticationTypeSMS PortalAuthenticationType = "sms"
const PortalAuthenticationTypeUserPass PortalAuthenticationType = "userpass"
const PortalAuthenticationTypeVoucher PortalAuthenticationType = "voucher"

func (en PortalAuthenticationType) GetPtr() *PortalAuthenticationType { var v = en; return &v }

func (en PortalAuthenticationType) String() string {
	switch en {
	case PortalAuthenticationTypeCallFront:
		return "callfront"
	case PortalAuthenticationTypeCallback:
		return "callback"
	case PortalAuthenticationTypeESIA:
		return "esia"
	case PortalAuthenticationTypeEmail:
		return "email"
	case PortalAuthenticationTypeHotelVoucher:
		return "hotel_voucher"
	case PortalAuthenticationTypeNone:
		return "none"
	case PortalAuthenticationTypeSMS:
		return "sms"
	case PortalAuthenticationTypeUserPass:
		return "userpass"
	case PortalAuthenticationTypeVoucher:
		return "voucher"
	}
	if len(en) == 0 {
		return "sms"
	}
	panic(errors.New("Invalid value of PortalAuthenticationType: " + string(en)))
}

func (en *PortalAuthenticationType) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalAuthenticationTypeCallFront:
		return json.Marshal("callfront")
	case PortalAuthenticationTypeCallback:
		return json.Marshal("callback")
	case PortalAuthenticationTypeESIA:
		return json.Marshal("esia")
	case PortalAuthenticationTypeEmail:
		return json.Marshal("email")
	case PortalAuthenticationTypeHotelVoucher:
		return json.Marshal("hotel_voucher")
	case PortalAuthenticationTypeNone:
		return json.Marshal("none")
	case PortalAuthenticationTypeSMS:
		return json.Marshal("sms")
	case PortalAuthenticationTypeUserPass:
		return json.Marshal("userpass")
	case PortalAuthenticationTypeVoucher:
		return json.Marshal("voucher")
	}
	if len(*en) == 0 {
		return json.Marshal("sms")
	}
	return nil, errors.New("Invalid value of PortalAuthenticationType: " + string(*en))
}

func (en *PortalAuthenticationType) GetBSON() (interface{}, error) {
	switch *en {
	case PortalAuthenticationTypeCallFront:
		return "callfront", nil
	case PortalAuthenticationTypeCallback:
		return "callback", nil
	case PortalAuthenticationTypeESIA:
		return "esia", nil
	case PortalAuthenticationTypeEmail:
		return "email", nil
	case PortalAuthenticationTypeHotelVoucher:
		return "hotel_voucher", nil
	case PortalAuthenticationTypeNone:
		return "none", nil
	case PortalAuthenticationTypeSMS:
		return "sms", nil
	case PortalAuthenticationTypeUserPass:
		return "userpass", nil
	case PortalAuthenticationTypeVoucher:
		return "voucher", nil
	}
	if len(*en) == 0 {
		return "sms", nil
	}
	return nil, errors.New("Invalid value of PortalAuthenticationType: " + string(*en))
}

func (en *PortalAuthenticationType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "callfront":
		*en = PortalAuthenticationTypeCallFront
		return nil
	case "callback":
		*en = PortalAuthenticationTypeCallback
		return nil
	case "esia":
		*en = PortalAuthenticationTypeESIA
		return nil
	case "email":
		*en = PortalAuthenticationTypeEmail
		return nil
	case "hotel_voucher":
		*en = PortalAuthenticationTypeHotelVoucher
		return nil
	case "none":
		*en = PortalAuthenticationTypeNone
		return nil
	case "sms":
		*en = PortalAuthenticationTypeSMS
		return nil
	case "userpass":
		*en = PortalAuthenticationTypeUserPass
		return nil
	case "voucher":
		*en = PortalAuthenticationTypeVoucher
		return nil
	}
	if len(s) == 0 {
		*en = PortalAuthenticationTypeSMS
		return nil
	}
	return errors.New("Unknown PortalAuthenticationType: " + s)
}

func (en *PortalAuthenticationType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "callfront":
		*en = PortalAuthenticationTypeCallFront
		return nil
	case "callback":
		*en = PortalAuthenticationTypeCallback
		return nil
	case "esia":
		*en = PortalAuthenticationTypeESIA
		return nil
	case "email":
		*en = PortalAuthenticationTypeEmail
		return nil
	case "hotel_voucher":
		*en = PortalAuthenticationTypeHotelVoucher
		return nil
	case "none":
		*en = PortalAuthenticationTypeNone
		return nil
	case "sms":
		*en = PortalAuthenticationTypeSMS
		return nil
	case "userpass":
		*en = PortalAuthenticationTypeUserPass
		return nil
	case "voucher":
		*en = PortalAuthenticationTypeVoucher
		return nil
	}
	if len(s) == 0 {
		*en = PortalAuthenticationTypeSMS
		return nil
	}
	return errors.New("Unknown PortalAuthenticationType: " + s)
}

type PortalAuthorizationState string

const PortalAuthorizationStateChecked PortalAuthorizationState = "checked"
const PortalAuthorizationStateNeed PortalAuthorizationState = "none"

func (en PortalAuthorizationState) GetPtr() *PortalAuthorizationState { var v = en; return &v }

func (en PortalAuthorizationState) String() string {
	switch en {
	case PortalAuthorizationStateChecked:
		return "checked"
	case PortalAuthorizationStateNeed:
		return "none"
	}
	if len(en) == 0 {
		return "none"
	}
	panic(errors.New("Invalid value of PortalAuthorizationState: " + string(en)))
}

func (en *PortalAuthorizationState) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalAuthorizationStateChecked:
		return json.Marshal("checked")
	case PortalAuthorizationStateNeed:
		return json.Marshal("none")
	}
	if len(*en) == 0 {
		return json.Marshal("none")
	}
	return nil, errors.New("Invalid value of PortalAuthorizationState: " + string(*en))
}

func (en *PortalAuthorizationState) GetBSON() (interface{}, error) {
	switch *en {
	case PortalAuthorizationStateChecked:
		return "checked", nil
	case PortalAuthorizationStateNeed:
		return "none", nil
	}
	if len(*en) == 0 {
		return "none", nil
	}
	return nil, errors.New("Invalid value of PortalAuthorizationState: " + string(*en))
}

func (en *PortalAuthorizationState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "checked":
		*en = PortalAuthorizationStateChecked
		return nil
	case "none":
		*en = PortalAuthorizationStateNeed
		return nil
	}
	if len(s) == 0 {
		*en = PortalAuthorizationStateNeed
		return nil
	}
	return errors.New("Unknown PortalAuthorizationState: " + s)
}

func (en *PortalAuthorizationState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "checked":
		*en = PortalAuthorizationStateChecked
		return nil
	case "none":
		*en = PortalAuthorizationStateNeed
		return nil
	}
	if len(s) == 0 {
		*en = PortalAuthorizationStateNeed
		return nil
	}
	return errors.New("Unknown PortalAuthorizationState: " + s)
}

type PortalAuthorizationType string

const PortalAuthorizationTypeExtVoucher PortalAuthorizationType = "ext_voucher"
const PortalAuthorizationTypeFB PortalAuthorizationType = "facebook"
const PortalAuthorizationTypeFree PortalAuthorizationType = "free"
const PortalAuthorizationTypeHotelVoucher PortalAuthorizationType = "hotel_voucher"
const PortalAuthorizationTypeIG PortalAuthorizationType = "instagram"
const PortalAuthorizationTypeNone PortalAuthorizationType = "none"
const PortalAuthorizationTypeSkip PortalAuthorizationType = "skip"
const PortalAuthorizationTypeSponsor PortalAuthorizationType = "sponsor"
const PortalAuthorizationTypeStaff PortalAuthorizationType = "staff"
const PortalAuthorizationTypeSubscription PortalAuthorizationType = "subscription"
const PortalAuthorizationTypeVK PortalAuthorizationType = "vk"
const PortalAuthorizationTypeVoucher PortalAuthorizationType = "voucher"

func (en PortalAuthorizationType) GetPtr() *PortalAuthorizationType { var v = en; return &v }

func (en PortalAuthorizationType) String() string {
	switch en {
	case PortalAuthorizationTypeExtVoucher:
		return "ext_voucher"
	case PortalAuthorizationTypeFB:
		return "facebook"
	case PortalAuthorizationTypeFree:
		return "free"
	case PortalAuthorizationTypeHotelVoucher:
		return "hotel_voucher"
	case PortalAuthorizationTypeIG:
		return "instagram"
	case PortalAuthorizationTypeNone:
		return "none"
	case PortalAuthorizationTypeSkip:
		return "skip"
	case PortalAuthorizationTypeSponsor:
		return "sponsor"
	case PortalAuthorizationTypeStaff:
		return "staff"
	case PortalAuthorizationTypeSubscription:
		return "subscription"
	case PortalAuthorizationTypeVK:
		return "vk"
	case PortalAuthorizationTypeVoucher:
		return "voucher"
	}
	if len(en) == 0 {
		return "none"
	}
	panic(errors.New("Invalid value of PortalAuthorizationType: " + string(en)))
}

func (en *PortalAuthorizationType) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalAuthorizationTypeExtVoucher:
		return json.Marshal("ext_voucher")
	case PortalAuthorizationTypeFB:
		return json.Marshal("facebook")
	case PortalAuthorizationTypeFree:
		return json.Marshal("free")
	case PortalAuthorizationTypeHotelVoucher:
		return json.Marshal("hotel_voucher")
	case PortalAuthorizationTypeIG:
		return json.Marshal("instagram")
	case PortalAuthorizationTypeNone:
		return json.Marshal("none")
	case PortalAuthorizationTypeSkip:
		return json.Marshal("skip")
	case PortalAuthorizationTypeSponsor:
		return json.Marshal("sponsor")
	case PortalAuthorizationTypeStaff:
		return json.Marshal("staff")
	case PortalAuthorizationTypeSubscription:
		return json.Marshal("subscription")
	case PortalAuthorizationTypeVK:
		return json.Marshal("vk")
	case PortalAuthorizationTypeVoucher:
		return json.Marshal("voucher")
	}
	if len(*en) == 0 {
		return json.Marshal("none")
	}
	return nil, errors.New("Invalid value of PortalAuthorizationType: " + string(*en))
}

func (en *PortalAuthorizationType) GetBSON() (interface{}, error) {
	switch *en {
	case PortalAuthorizationTypeExtVoucher:
		return "ext_voucher", nil
	case PortalAuthorizationTypeFB:
		return "facebook", nil
	case PortalAuthorizationTypeFree:
		return "free", nil
	case PortalAuthorizationTypeHotelVoucher:
		return "hotel_voucher", nil
	case PortalAuthorizationTypeIG:
		return "instagram", nil
	case PortalAuthorizationTypeNone:
		return "none", nil
	case PortalAuthorizationTypeSkip:
		return "skip", nil
	case PortalAuthorizationTypeSponsor:
		return "sponsor", nil
	case PortalAuthorizationTypeStaff:
		return "staff", nil
	case PortalAuthorizationTypeSubscription:
		return "subscription", nil
	case PortalAuthorizationTypeVK:
		return "vk", nil
	case PortalAuthorizationTypeVoucher:
		return "voucher", nil
	}
	if len(*en) == 0 {
		return "none", nil
	}
	return nil, errors.New("Invalid value of PortalAuthorizationType: " + string(*en))
}

func (en *PortalAuthorizationType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ext_voucher":
		*en = PortalAuthorizationTypeExtVoucher
		return nil
	case "facebook":
		*en = PortalAuthorizationTypeFB
		return nil
	case "free":
		*en = PortalAuthorizationTypeFree
		return nil
	case "hotel_voucher":
		*en = PortalAuthorizationTypeHotelVoucher
		return nil
	case "instagram":
		*en = PortalAuthorizationTypeIG
		return nil
	case "none":
		*en = PortalAuthorizationTypeNone
		return nil
	case "skip":
		*en = PortalAuthorizationTypeSkip
		return nil
	case "sponsor":
		*en = PortalAuthorizationTypeSponsor
		return nil
	case "staff":
		*en = PortalAuthorizationTypeStaff
		return nil
	case "subscription":
		*en = PortalAuthorizationTypeSubscription
		return nil
	case "vk":
		*en = PortalAuthorizationTypeVK
		return nil
	case "voucher":
		*en = PortalAuthorizationTypeVoucher
		return nil
	}
	if len(s) == 0 {
		*en = PortalAuthorizationTypeNone
		return nil
	}
	return errors.New("Unknown PortalAuthorizationType: " + s)
}

func (en *PortalAuthorizationType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ext_voucher":
		*en = PortalAuthorizationTypeExtVoucher
		return nil
	case "facebook":
		*en = PortalAuthorizationTypeFB
		return nil
	case "free":
		*en = PortalAuthorizationTypeFree
		return nil
	case "hotel_voucher":
		*en = PortalAuthorizationTypeHotelVoucher
		return nil
	case "instagram":
		*en = PortalAuthorizationTypeIG
		return nil
	case "none":
		*en = PortalAuthorizationTypeNone
		return nil
	case "skip":
		*en = PortalAuthorizationTypeSkip
		return nil
	case "sponsor":
		*en = PortalAuthorizationTypeSponsor
		return nil
	case "staff":
		*en = PortalAuthorizationTypeStaff
		return nil
	case "subscription":
		*en = PortalAuthorizationTypeSubscription
		return nil
	case "vk":
		*en = PortalAuthorizationTypeVK
		return nil
	case "voucher":
		*en = PortalAuthorizationTypeVoucher
		return nil
	}
	if len(s) == 0 {
		*en = PortalAuthorizationTypeNone
		return nil
	}
	return errors.New("Unknown PortalAuthorizationType: " + s)
}

type PortalOSType string

const PortalOSTypeAndroid PortalOSType = "Android"
const PortalOSTypeIOS PortalOSType = "iPhone OS"
const PortalOSTypeLinux PortalOSType = "Linux"
const PortalOSTypeMacOs PortalOSType = "Mac Os"
const PortalOSTypeNone PortalOSType = "none"
const PortalOSTypeWindows PortalOSType = "Windows"

func (en PortalOSType) GetPtr() *PortalOSType { var v = en; return &v }

func (en PortalOSType) String() string {
	switch en {
	case PortalOSTypeAndroid:
		return "Android"
	case PortalOSTypeIOS:
		return "iPhone OS"
	case PortalOSTypeLinux:
		return "Linux"
	case PortalOSTypeMacOs:
		return "Mac Os"
	case PortalOSTypeNone:
		return "none"
	case PortalOSTypeWindows:
		return "Windows"
	}
	if len(en) == 0 {
		return "none"
	}
	panic(errors.New("Invalid value of PortalOSType: " + string(en)))
}

func (en *PortalOSType) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalOSTypeAndroid:
		return json.Marshal("Android")
	case PortalOSTypeIOS:
		return json.Marshal("iPhone OS")
	case PortalOSTypeLinux:
		return json.Marshal("Linux")
	case PortalOSTypeMacOs:
		return json.Marshal("Mac Os")
	case PortalOSTypeNone:
		return json.Marshal("none")
	case PortalOSTypeWindows:
		return json.Marshal("Windows")
	}
	if len(*en) == 0 {
		return json.Marshal("none")
	}
	return nil, errors.New("Invalid value of PortalOSType: " + string(*en))
}

func (en *PortalOSType) GetBSON() (interface{}, error) {
	switch *en {
	case PortalOSTypeAndroid:
		return "Android", nil
	case PortalOSTypeIOS:
		return "iPhone OS", nil
	case PortalOSTypeLinux:
		return "Linux", nil
	case PortalOSTypeMacOs:
		return "Mac Os", nil
	case PortalOSTypeNone:
		return "none", nil
	case PortalOSTypeWindows:
		return "Windows", nil
	}
	if len(*en) == 0 {
		return "none", nil
	}
	return nil, errors.New("Invalid value of PortalOSType: " + string(*en))
}

func (en *PortalOSType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "Android":
		*en = PortalOSTypeAndroid
		return nil
	case "iPhone OS":
		*en = PortalOSTypeIOS
		return nil
	case "Linux":
		*en = PortalOSTypeLinux
		return nil
	case "Mac Os":
		*en = PortalOSTypeMacOs
		return nil
	case "none":
		*en = PortalOSTypeNone
		return nil
	case "Windows":
		*en = PortalOSTypeWindows
		return nil
	}
	if len(s) == 0 {
		*en = PortalOSTypeNone
		return nil
	}
	return errors.New("Unknown PortalOSType: " + s)
}

func (en *PortalOSType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "Android":
		*en = PortalOSTypeAndroid
		return nil
	case "iPhone OS":
		*en = PortalOSTypeIOS
		return nil
	case "Linux":
		*en = PortalOSTypeLinux
		return nil
	case "Mac Os":
		*en = PortalOSTypeMacOs
		return nil
	case "none":
		*en = PortalOSTypeNone
		return nil
	case "Windows":
		*en = PortalOSTypeWindows
		return nil
	}
	if len(s) == 0 {
		*en = PortalOSTypeNone
		return nil
	}
	return errors.New("Unknown PortalOSType: " + s)
}

type PortalPaymentSystemType string

const PortalPaymentSystemTypeExternal PortalPaymentSystemType = "ext"
const PortalPaymentSystemTypeExternalURL PortalPaymentSystemType = "ext_url"
const PortalPaymentSystemTypeGuide PortalPaymentSystemType = "guide"

func (en PortalPaymentSystemType) GetPtr() *PortalPaymentSystemType { var v = en; return &v }

func (en PortalPaymentSystemType) String() string {
	switch en {
	case PortalPaymentSystemTypeExternal:
		return "ext"
	case PortalPaymentSystemTypeExternalURL:
		return "ext_url"
	case PortalPaymentSystemTypeGuide:
		return "guide"
	}
	if len(en) == 0 {
		return "guide"
	}
	panic(errors.New("Invalid value of PortalPaymentSystemType: " + string(en)))
}

func (en *PortalPaymentSystemType) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalPaymentSystemTypeExternal:
		return json.Marshal("ext")
	case PortalPaymentSystemTypeExternalURL:
		return json.Marshal("ext_url")
	case PortalPaymentSystemTypeGuide:
		return json.Marshal("guide")
	}
	if len(*en) == 0 {
		return json.Marshal("guide")
	}
	return nil, errors.New("Invalid value of PortalPaymentSystemType: " + string(*en))
}

func (en *PortalPaymentSystemType) GetBSON() (interface{}, error) {
	switch *en {
	case PortalPaymentSystemTypeExternal:
		return "ext", nil
	case PortalPaymentSystemTypeExternalURL:
		return "ext_url", nil
	case PortalPaymentSystemTypeGuide:
		return "guide", nil
	}
	if len(*en) == 0 {
		return "guide", nil
	}
	return nil, errors.New("Invalid value of PortalPaymentSystemType: " + string(*en))
}

func (en *PortalPaymentSystemType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ext":
		*en = PortalPaymentSystemTypeExternal
		return nil
	case "ext_url":
		*en = PortalPaymentSystemTypeExternalURL
		return nil
	case "guide":
		*en = PortalPaymentSystemTypeGuide
		return nil
	}
	if len(s) == 0 {
		*en = PortalPaymentSystemTypeGuide
		return nil
	}
	return errors.New("Unknown PortalPaymentSystemType: " + s)
}

func (en *PortalPaymentSystemType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ext":
		*en = PortalPaymentSystemTypeExternal
		return nil
	case "ext_url":
		*en = PortalPaymentSystemTypeExternalURL
		return nil
	case "guide":
		*en = PortalPaymentSystemTypeGuide
		return nil
	}
	if len(s) == 0 {
		*en = PortalPaymentSystemTypeGuide
		return nil
	}
	return errors.New("Unknown PortalPaymentSystemType: " + s)
}

type PortalReportObjectType string

const PortalReportObjectTypeAdAndPollsReport PortalReportObjectType = "ad_and_polls"
const PortalReportObjectTypeClientsReport PortalReportObjectType = "clients"
const PortalReportObjectTypeDateOfBirthReport PortalReportObjectType = "date_of_birth"
const PortalReportObjectTypeNone PortalReportObjectType = "none"
const PortalReportObjectTypeSessionReport PortalReportObjectType = "session"
const PortalReportObjectTypeStatReport PortalReportObjectType = "stat"

func (en PortalReportObjectType) GetPtr() *PortalReportObjectType { var v = en; return &v }

func (en PortalReportObjectType) String() string {
	switch en {
	case PortalReportObjectTypeAdAndPollsReport:
		return "ad_and_polls"
	case PortalReportObjectTypeClientsReport:
		return "clients"
	case PortalReportObjectTypeDateOfBirthReport:
		return "date_of_birth"
	case PortalReportObjectTypeNone:
		return "none"
	case PortalReportObjectTypeSessionReport:
		return "session"
	case PortalReportObjectTypeStatReport:
		return "stat"
	}
	if len(en) == 0 {
		return "none"
	}
	panic(errors.New("Invalid value of PortalReportObjectType: " + string(en)))
}

func (en *PortalReportObjectType) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalReportObjectTypeAdAndPollsReport:
		return json.Marshal("ad_and_polls")
	case PortalReportObjectTypeClientsReport:
		return json.Marshal("clients")
	case PortalReportObjectTypeDateOfBirthReport:
		return json.Marshal("date_of_birth")
	case PortalReportObjectTypeNone:
		return json.Marshal("none")
	case PortalReportObjectTypeSessionReport:
		return json.Marshal("session")
	case PortalReportObjectTypeStatReport:
		return json.Marshal("stat")
	}
	if len(*en) == 0 {
		return json.Marshal("none")
	}
	return nil, errors.New("Invalid value of PortalReportObjectType: " + string(*en))
}

func (en *PortalReportObjectType) GetBSON() (interface{}, error) {
	switch *en {
	case PortalReportObjectTypeAdAndPollsReport:
		return "ad_and_polls", nil
	case PortalReportObjectTypeClientsReport:
		return "clients", nil
	case PortalReportObjectTypeDateOfBirthReport:
		return "date_of_birth", nil
	case PortalReportObjectTypeNone:
		return "none", nil
	case PortalReportObjectTypeSessionReport:
		return "session", nil
	case PortalReportObjectTypeStatReport:
		return "stat", nil
	}
	if len(*en) == 0 {
		return "none", nil
	}
	return nil, errors.New("Invalid value of PortalReportObjectType: " + string(*en))
}

func (en *PortalReportObjectType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "ad_and_polls":
		*en = PortalReportObjectTypeAdAndPollsReport
		return nil
	case "clients":
		*en = PortalReportObjectTypeClientsReport
		return nil
	case "date_of_birth":
		*en = PortalReportObjectTypeDateOfBirthReport
		return nil
	case "none":
		*en = PortalReportObjectTypeNone
		return nil
	case "session":
		*en = PortalReportObjectTypeSessionReport
		return nil
	case "stat":
		*en = PortalReportObjectTypeStatReport
		return nil
	}
	if len(s) == 0 {
		*en = PortalReportObjectTypeNone
		return nil
	}
	return errors.New("Unknown PortalReportObjectType: " + s)
}

func (en *PortalReportObjectType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "ad_and_polls":
		*en = PortalReportObjectTypeAdAndPollsReport
		return nil
	case "clients":
		*en = PortalReportObjectTypeClientsReport
		return nil
	case "date_of_birth":
		*en = PortalReportObjectTypeDateOfBirthReport
		return nil
	case "none":
		*en = PortalReportObjectTypeNone
		return nil
	case "session":
		*en = PortalReportObjectTypeSessionReport
		return nil
	case "stat":
		*en = PortalReportObjectTypeStatReport
		return nil
	}
	if len(s) == 0 {
		*en = PortalReportObjectTypeNone
		return nil
	}
	return errors.New("Unknown PortalReportObjectType: " + s)
}

type PortalResponseStatus string

const PortalResponseStatusError PortalResponseStatus = "error"
const PortalResponseStatusSuccess PortalResponseStatus = "success"

func (en PortalResponseStatus) GetPtr() *PortalResponseStatus { var v = en; return &v }

func (en PortalResponseStatus) String() string {
	switch en {
	case PortalResponseStatusError:
		return "error"
	case PortalResponseStatusSuccess:
		return "success"
	}
	if len(en) == 0 {
		return "success"
	}
	panic(errors.New("Invalid value of PortalResponseStatus: " + string(en)))
}

func (en *PortalResponseStatus) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalResponseStatusError:
		return json.Marshal("error")
	case PortalResponseStatusSuccess:
		return json.Marshal("success")
	}
	if len(*en) == 0 {
		return json.Marshal("success")
	}
	return nil, errors.New("Invalid value of PortalResponseStatus: " + string(*en))
}

func (en *PortalResponseStatus) GetBSON() (interface{}, error) {
	switch *en {
	case PortalResponseStatusError:
		return "error", nil
	case PortalResponseStatusSuccess:
		return "success", nil
	}
	if len(*en) == 0 {
		return "success", nil
	}
	return nil, errors.New("Invalid value of PortalResponseStatus: " + string(*en))
}

func (en *PortalResponseStatus) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "error":
		*en = PortalResponseStatusError
		return nil
	case "success":
		*en = PortalResponseStatusSuccess
		return nil
	}
	if len(s) == 0 {
		*en = PortalResponseStatusSuccess
		return nil
	}
	return errors.New("Unknown PortalResponseStatus: " + s)
}

func (en *PortalResponseStatus) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "error":
		*en = PortalResponseStatusError
		return nil
	case "success":
		*en = PortalResponseStatusSuccess
		return nil
	}
	if len(s) == 0 {
		*en = PortalResponseStatusSuccess
		return nil
	}
	return errors.New("Unknown PortalResponseStatus: " + s)
}

type PortalSignType string

const PortalSignTypeEqual PortalSignType = "equal"
const PortalSignTypeGreater PortalSignType = "greater"
const PortalSignTypeLess PortalSignType = "less"

func (en PortalSignType) GetPtr() *PortalSignType { var v = en; return &v }

func (en PortalSignType) String() string {
	switch en {
	case PortalSignTypeEqual:
		return "equal"
	case PortalSignTypeGreater:
		return "greater"
	case PortalSignTypeLess:
		return "less"
	}
	if len(en) == 0 {
		return "greater"
	}
	panic(errors.New("Invalid value of PortalSignType: " + string(en)))
}

func (en *PortalSignType) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalSignTypeEqual:
		return json.Marshal("equal")
	case PortalSignTypeGreater:
		return json.Marshal("greater")
	case PortalSignTypeLess:
		return json.Marshal("less")
	}
	if len(*en) == 0 {
		return json.Marshal("greater")
	}
	return nil, errors.New("Invalid value of PortalSignType: " + string(*en))
}

func (en *PortalSignType) GetBSON() (interface{}, error) {
	switch *en {
	case PortalSignTypeEqual:
		return "equal", nil
	case PortalSignTypeGreater:
		return "greater", nil
	case PortalSignTypeLess:
		return "less", nil
	}
	if len(*en) == 0 {
		return "greater", nil
	}
	return nil, errors.New("Invalid value of PortalSignType: " + string(*en))
}

func (en *PortalSignType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "equal":
		*en = PortalSignTypeEqual
		return nil
	case "greater":
		*en = PortalSignTypeGreater
		return nil
	case "less":
		*en = PortalSignTypeLess
		return nil
	}
	if len(s) == 0 {
		*en = PortalSignTypeGreater
		return nil
	}
	return errors.New("Unknown PortalSignType: " + s)
}

func (en *PortalSignType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "equal":
		*en = PortalSignTypeEqual
		return nil
	case "greater":
		*en = PortalSignTypeGreater
		return nil
	case "less":
		*en = PortalSignTypeLess
		return nil
	}
	if len(s) == 0 {
		*en = PortalSignTypeGreater
		return nil
	}
	return errors.New("Unknown PortalSignType: " + s)
}

type PortalUserState string

const PortalUserStateAdvertise PortalUserState = "advertise"
const PortalUserStateAuthenticate PortalUserState = "authenticate"
const PortalUserStateAuthorize PortalUserState = "authorize"
const PortalUserStateNew PortalUserState = "new"
const PortalUserStatePass PortalUserState = "pass"

func (en PortalUserState) GetPtr() *PortalUserState { var v = en; return &v }

func (en PortalUserState) String() string {
	switch en {
	case PortalUserStateAdvertise:
		return "advertise"
	case PortalUserStateAuthenticate:
		return "authenticate"
	case PortalUserStateAuthorize:
		return "authorize"
	case PortalUserStateNew:
		return "new"
	case PortalUserStatePass:
		return "pass"
	}
	if len(en) == 0 {
		return "new"
	}
	panic(errors.New("Invalid value of PortalUserState: " + string(en)))
}

func (en *PortalUserState) MarshalJSON() ([]byte, error) {
	switch *en {
	case PortalUserStateAdvertise:
		return json.Marshal("advertise")
	case PortalUserStateAuthenticate:
		return json.Marshal("authenticate")
	case PortalUserStateAuthorize:
		return json.Marshal("authorize")
	case PortalUserStateNew:
		return json.Marshal("new")
	case PortalUserStatePass:
		return json.Marshal("pass")
	}
	if len(*en) == 0 {
		return json.Marshal("new")
	}
	return nil, errors.New("Invalid value of PortalUserState: " + string(*en))
}

func (en *PortalUserState) GetBSON() (interface{}, error) {
	switch *en {
	case PortalUserStateAdvertise:
		return "advertise", nil
	case PortalUserStateAuthenticate:
		return "authenticate", nil
	case PortalUserStateAuthorize:
		return "authorize", nil
	case PortalUserStateNew:
		return "new", nil
	case PortalUserStatePass:
		return "pass", nil
	}
	if len(*en) == 0 {
		return "new", nil
	}
	return nil, errors.New("Invalid value of PortalUserState: " + string(*en))
}

func (en *PortalUserState) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "advertise":
		*en = PortalUserStateAdvertise
		return nil
	case "authenticate":
		*en = PortalUserStateAuthenticate
		return nil
	case "authorize":
		*en = PortalUserStateAuthorize
		return nil
	case "new":
		*en = PortalUserStateNew
		return nil
	case "pass":
		*en = PortalUserStatePass
		return nil
	}
	if len(s) == 0 {
		*en = PortalUserStateNew
		return nil
	}
	return errors.New("Unknown PortalUserState: " + s)
}

func (en *PortalUserState) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "advertise":
		*en = PortalUserStateAdvertise
		return nil
	case "authenticate":
		*en = PortalUserStateAuthenticate
		return nil
	case "authorize":
		*en = PortalUserStateAuthorize
		return nil
	case "new":
		*en = PortalUserStateNew
		return nil
	case "pass":
		*en = PortalUserStatePass
		return nil
	}
	if len(s) == 0 {
		*en = PortalUserStateNew
		return nil
	}
	return errors.New("Unknown PortalUserState: " + s)
}

type TariffPlanType string

const TariffPlanTypeSubscription TariffPlanType = "subscription"
const TariffPlanTypeVoucher TariffPlanType = "voucher"

func (en TariffPlanType) GetPtr() *TariffPlanType { var v = en; return &v }

func (en TariffPlanType) String() string {
	switch en {
	case TariffPlanTypeSubscription:
		return "subscription"
	case TariffPlanTypeVoucher:
		return "voucher"
	}
	if len(en) == 0 {
		return "voucher"
	}
	panic(errors.New("Invalid value of TariffPlanType: " + string(en)))
}

func (en *TariffPlanType) MarshalJSON() ([]byte, error) {
	switch *en {
	case TariffPlanTypeSubscription:
		return json.Marshal("subscription")
	case TariffPlanTypeVoucher:
		return json.Marshal("voucher")
	}
	if len(*en) == 0 {
		return json.Marshal("voucher")
	}
	return nil, errors.New("Invalid value of TariffPlanType: " + string(*en))
}

func (en *TariffPlanType) GetBSON() (interface{}, error) {
	switch *en {
	case TariffPlanTypeSubscription:
		return "subscription", nil
	case TariffPlanTypeVoucher:
		return "voucher", nil
	}
	if len(*en) == 0 {
		return "voucher", nil
	}
	return nil, errors.New("Invalid value of TariffPlanType: " + string(*en))
}

func (en *TariffPlanType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	switch s {
	case "subscription":
		*en = TariffPlanTypeSubscription
		return nil
	case "voucher":
		*en = TariffPlanTypeVoucher
		return nil
	}
	if len(s) == 0 {
		*en = TariffPlanTypeVoucher
		return nil
	}
	return errors.New("Unknown TariffPlanType: " + s)
}

func (en *TariffPlanType) SetBSON(v bson.Raw) error {
	var s string
	if err := v.Unmarshal(&s); err != nil {
		return err
	}
	switch s {
	case "subscription":
		*en = TariffPlanTypeSubscription
		return nil
	case "voucher":
		*en = TariffPlanTypeVoucher
		return nil
	}
	if len(s) == 0 {
		*en = TariffPlanTypeVoucher
		return nil
	}
	return errors.New("Unknown TariffPlanType: " + s)
}
