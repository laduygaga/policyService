package models

var Rules = map[string]string {
	"reject_null_sender" : "reject_null_sender",
	"reject_sender_login_mismatch	" : "reject_sender_login_mismatch",
	"default" : "default",
	"accept" : "accept",
	"whitelist" : "whitelist",
	"discard" : "discard",
	"reject" : "reject",
	"reject_forged_sender" : "reject_forged_sender",
	"reject_blacklisted" : "reject_blacklisted",
	"reject_not_authorized" : "reject_not_authorized",
	"reject_message_size_exceeded" : "reject_message_size_exceeded",
	"reject_blacklisted_rdns" : "reject_blacklisted_rdns",
	"reject_exceed_msg_size" : "reject_exceed_msg_size",
	"reject_exceed_max_msgs" : "reject_exceed_max_msgs",
	"reject_exceed_max_quota" : "reject_exceed_max_quota",
	"greylisting" : "greylisting",
}

type Rule struct {
	Id int
	Name string
	Default_value string
	Type string
	Description string
}

type PolicyClass struct {
	Id int
	Name string
	Description string
}

type PolicyClassRule struct {
	Id int
	Policy_class_id int
	Rule_id int
	Default_value string
}

type User struct {
	Id int
	Email string
}

type UserPolicyClass struct {
	// Id int
	User_id int
	Policy_class_id int
	Priority int
}

type User_rule struct {
	User_id int
	Rule_id int
	Value int
	Email string
}

type Usages struct {
	Id int
	User_id int
	Rule_id int
	Value string
}

type PayloadNested struct {
	Reverse_client_name string
	Ccert_pubkey_fingerprint string
	Ccert_subject string
	Sender_domain string
	Sasl_sender string
	Protocol_state string
	Encryption_protocol string
	Ccert_issuer string
	Client_address string
	Size string
	Protocol_name string
	Sasl_username_domain string
	Client_name string
	Policy_context string
	Helo_name string
	Etrn_domain string
	Instance string
	Encryption_keysize string
	Encryption_cipher string
	Ccert_fingerprint string
	Recipient_count string
	Client_port string
	Queue_id string
	Sasl_method string
	Recipient string
	Sasl_username string
	Stress string
	Sender string
	Request string
	Recipient_domain string
	}

type RequestPayload struct {
	Recipient string
	Smtp_session_data PayloadNested
	Sender string
	Sasl_username_domain string
	Conn_iredapd string
	Conn_amavisd string
	Sender_domain string
	Conn_vmail string
	Recipient_domain string
	Recipient_without_ext string
	Sender_without_ext string
	Sasl_username string
	Client_address string
}
