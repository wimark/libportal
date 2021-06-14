package libportal

const (
	CollAccessServer = "portal_access_server"
)

// AccessServer
type AccessServer struct {
	ID string `json:"id" bson:"_id"`

	Enable bool             `json:"enable" bson:"enable"`
	Name   string           `json:"name" bson:"name"`
	Type   AccessServerType `json:"type" bson:"type"`

	NasID      string `json:"nas_id" bson:"nas_id"`
	NasIP      string `json:"nas_ip" bson:"nas_ip"`
	ExternalIP string `json:"external_ip" bson:"external_ip"`

	Location    string `json:"location" bson:"location"`
	Description string `json:"desctiption" bson:"description"`

	Login    string `json:"login" bson:"login"`
	Password string `json:"password" bson:"password"`

	PreAuthVSA  string `json:"pre_auth_vsa" bson:"pre_auth_vsa"`
	PostAuthVSA string `json:"post_auth_vsa" bson:"post_auth_vsa"`
	CoaVSA      string `json:"coa_vsa" bson:"coa_vsa"`

	CoaSecret string `json:"coa_secret" bson:"coa_secret"`
	CoaPort   int    `json:"coa_port" bson:"coa_port"`
	MacAuth   bool   `json:"mac_auth" bson:"mac_auth"`
}
