package entity

import "time"

type Member struct {
	ID                         int64     `json:"id"`
	Avatar                     string    `json:"avatar"`
	Bank                       string    `json:"bank"`
	Branch                     string    `json:"branch"`
	CardNo                     string    `json:"card_no"`
	CertifiedBusinessApplyTime time.Time `json:"certified_business_apply_time"`
	CertifiedBusinessCheckTime time.Time `json:"certified_business_check_time"`
	CertifiedBusinessStatus    int64     `json:"certified_business_status"`
	ChannelId                  int64     `json:"channel_id"`
	Email                      string    `json:"email"`
	FirstLevel                 int64     `json:"first_level"`
	GoogleDate                 time.Time `json:"google_date"`
	GoogleKey                  string    `json:"google_key"`
	GoogleState                int64     `json:"google_state"`
	IdNumber                   string    `json:"id_number"`
	InviterId                  int64     `json:"inviter_id"`
	IsChannel                  int64     `json:"is_channel"`
	JyPassword                 string    `json:"jy_password"`
	LastLoginTime              time.Time `json:"last_login_time"`
	City                       string    `json:"city"`
	Country                    string    `json:"country"`
	District                   string    `json:"district"`
	Province                   string    `json:"province"`
	LoginCount                 int64     `json:"login_count"`
	LoginLock                  int64     `json:"login_lock"`
	Margin                     string    `json:"margin"`
	MemberLevel                int64     `json:"member_level"`
	MobilePhone                string    `json:"mobile_phone"`
	Password                   string    `json:"password"`
	PromotionCode              string    `json:"promotion_code"`
	PublishAdvertise           int64     `json:"publish_advertise"`
	RealName                   string    `json:"real_name"`
	RealNameStatus             int64     `json:"real_name_status"`
	RegistrationTime           time.Time `json:"registration_time"`
	Salt                       string    `json:"salt"`
	SecondLevel                int64     `json:"second_level"`
	SignInAbility              int64     `json:"sign_in_ability"`
	Status                     int64     `json:"status"`
	ThirdLevel                 int64     `json:"third_level"`
	Token                      string    `json:"token"`
	TokenExpireTime            time.Time `json:"token_expire_time"`
	TransactionStatus          int64     `json:"transaction_status"`
	TransactionTime            time.Time `json:"transaction_time"`
	Transactions               int64     `json:"transactions"`
	Username                   string    `json:"username"`
	QrWeCodeUrl                string    `json:"qr_we_code_url"`
	Wechat                     string    `json:"wechat"`
	Local                      string    `json:"local"`
	Integration                int64     `json:"integration"`
	MemberGradeId              int64     `json:"member_grade_id"`
	KycStatus                  int64     `json:"kyc_status"`
	GeneralizeTotal            int64     `json:"generalize_total"`
	InviterParentId            int64     `json:"inviter_parent_id"`
}
