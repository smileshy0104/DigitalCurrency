package model

// Member 结构体
type Member struct {
	Id                         int64   `gorm:"column:id"`
	AliNo                      string  `gorm:"column:ali_no" default:"0"`
	QrCodeUrl                  string  `gorm:"column:qr_code_url"`
	AppealSuccessTimes         int64   `gorm:"column:appeal_success_times"`
	AppealTimes                int64   `gorm:"column:appeal_times"`
	ApplicationTime            int64   `gorm:"column:application_time"`
	Avatar                     string  `gorm:"column:avatar"`
	Bank                       string  `gorm:"column:bank"`
	Branch                     string  `gorm:"column:branch"`
	CardNo                     string  `gorm:"column:card_no"`
	CertifiedBusinessApplyTime int64   `gorm:"column:certified_business_apply_time"`
	CertifiedBusinessCheckTime int64   `gorm:"column:certified_business_check_time"`
	CertifiedBusinessStatus    int64   `gorm:"column:certified_business_status"`
	ChannelId                  int64   `gorm:"column:channel_id"`
	Email                      string  `gorm:"column:email"`
	FirstLevel                 int64   `gorm:"column:first_level"`
	GoogleDate                 int64   `gorm:"column:google_date"`
	GoogleKey                  string  `gorm:"column:google_key"`
	GoogleState                int64   `gorm:"column:google_state"`
	IdNumber                   string  `gorm:"column:id_number"`
	InviterId                  int64   `gorm:"column:inviter_id"`
	IsChannel                  int64   `gorm:"column:is_channel"`
	JyPassword                 string  `gorm:"column:jy_password"`
	LastLoginTime              int64   `gorm:"column:last_login_time"`
	City                       string  `gorm:"column:city"`
	Country                    string  `gorm:"column:country"`
	District                   string  `gorm:"column:district"`
	Province                   string  `gorm:"column:province"`
	LoginCount                 int64   `gorm:"column:login_count"`
	LoginLock                  int64   `gorm:"column:login_lock"`
	Margin                     string  `gorm:"column:margin"`
	MemberLevel                int64   `gorm:"column:member_level"`
	MobilePhone                string  `gorm:"column:mobile_phone"`
	Password                   string  `gorm:"column:password"`
	PromotionCode              string  `gorm:"column:promotion_code"`
	PublishAdvertise           int64   `gorm:"column:publish_advertise"`
	RealName                   string  `gorm:"column:real_name"`
	RealNameStatus             int64   `gorm:"column:real_name_status"`
	RegistrationTime           int64   `gorm:"column:registration_time"`
	Salt                       string  `gorm:"column:salt"`
	SecondLevel                int64   `gorm:"column:second_level"`
	SignInAbility              int64   `gorm:"column:sign_in_ability"`
	Status                     int64   `gorm:"column:status"`
	ThirdLevel                 int64   `gorm:"column:third_level"`
	Token                      string  `gorm:"column:token"`
	TokenExpireTime            int64   `gorm:"column:token_expire_time"`
	TransactionStatus          int64   `gorm:"column:transaction_status"`
	TransactionTime            int64   `gorm:"column:transaction_time"`
	Transactions               int64   `gorm:"column:transactions"`
	Username                   string  `gorm:"column:username"`
	QrWeCodeUrl                string  `gorm:"column:qr_we_code_url"`
	Wechat                     string  `gorm:"column:wechat"`
	Local                      string  `gorm:"column:local"`
	Integration                int64   `gorm:"column:integration"`
	MemberGradeId              int64   `gorm:"column:member_grade_id"`  // 等级id
	KycStatus                  int64   `gorm:"column:kyc_status"`       // kyc等级
	GeneralizeTotal            int64   `gorm:"column:generalize_total"` // 注册赠送积分
	InviterParentId            int64   `gorm:"column:inviter_parent_id"`
	SuperPartner               string  `gorm:"column:super_partner"`
	KickFee                    float64 `gorm:"column:kick_fee"`
	Power                      float64 `gorm:"column:power"`      // 个人矿机算力(每日维护)
	TeamLevel                  int64   `gorm:"column:team_level"` // 团队人数(每日维护)
	TeamPower                  float64 `gorm:"column:team_power"` // 团队矿机算力(每日维护)
	MemberLevelId              int64   `gorm:"column:member_level_id"`
}

// 表名
func (*Member) TableName() string {
	return "member"
}

// Member 相关的常量定义，用于标识会员的不同属性和状态
const (
	GENERAL        = iota // 普通会员
	REALNAME              // 实名会员
	IDENTIFICATION        // 认证商家
)

// Partner 类型的常量定义，用于区分合作伙伴的等级
const (
	NORMALPARTER = "0" // 普通合作伙伴
	SUPERPARTER  = "1" // 超级合作伙伴
	PSUPERPARTER = "2" // 更高级别的合作伙伴
)

// Member 状态常量定义
const (
	NORMAL  = iota // 正常状态
	ILLEGAL        // 非法状态
)

// FillSuperPartner 根据传入的 partner 字符串设置会员的超级合作伙伴状态和会员状态
// 参数 partner: 代表合作伙伴等级的字符串
func (m *Member) FillSuperPartner(partner string) {
	if partner == "" {
		m.SuperPartner = NORMALPARTER
		m.Status = NORMAL
	} else {
		if partner != NORMALPARTER {
			m.SuperPartner = partner
			m.Status = ILLEGAL
		}
	}
}

// MemberLevelStr 返回会员级别的字符串表示
// 返回值: 会员级别的中文描述
func (m *Member) MemberLevelStr() string {
	if m.MemberLevel == GENERAL {
		return "普通会员"
	}
	if m.MemberLevel == REALNAME {
		return "实名"
	}
	if m.MemberLevel == IDENTIFICATION {
		return "认证商家"
	}
	return ""
}

// MemberRate 根据会员的超级合作伙伴状态返回会员费率
// 返回值: 会员费率，以百分比表示
func (m *Member) MemberRate() int32 {
	if m.SuperPartner == NORMALPARTER {
		return 0
	}
	if m.SuperPartner == SUPERPARTER {
		return 1
	}
	if m.SuperPartner == PSUPERPARTER {
		return 2
	}
	return 0
}

// NewMember 创建并返回一个新的 Member 实例
// 返回值: 新的 Member 实例的指针
func NewMember() *Member {
	return &Member{}
}
