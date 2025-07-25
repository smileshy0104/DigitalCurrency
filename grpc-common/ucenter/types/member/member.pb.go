// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: member.proto

package member

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MemberReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MemberId      int64                  `protobuf:"varint,3,opt,name=memberId,proto3" json:"memberId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *MemberReq) Reset() {
	*x = MemberReq{}
	mi := &file_member_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemberReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberReq) ProtoMessage() {}

func (x *MemberReq) ProtoReflect() protoreflect.Message {
	mi := &file_member_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberReq.ProtoReflect.Descriptor instead.
func (*MemberReq) Descriptor() ([]byte, []int) {
	return file_member_proto_rawDescGZIP(), []int{0}
}

func (x *MemberReq) GetMemberId() int64 {
	if x != nil {
		return x.MemberId
	}
	return 0
}

type MemberInfo struct {
	state                      protoimpl.MessageState `protogen:"open.v1"`
	Id                         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AliNo                      string                 `protobuf:"bytes,2,opt,name=aliNo,proto3" json:"aliNo,omitempty"`
	QrCodeUrl                  string                 `protobuf:"bytes,3,opt,name=qrCodeUrl,proto3" json:"qrCodeUrl,omitempty"`
	AppealSuccessTimes         int32                  `protobuf:"varint,4,opt,name=appealSuccessTimes,proto3" json:"appealSuccessTimes,omitempty"`
	AppealTimes                int32                  `protobuf:"varint,5,opt,name=appealTimes,proto3" json:"appealTimes,omitempty"`
	ApplicationTime            int64                  `protobuf:"varint,6,opt,name=applicationTime,proto3" json:"applicationTime,omitempty"`
	Avatar                     string                 `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Bank                       string                 `protobuf:"bytes,8,opt,name=bank,proto3" json:"bank,omitempty"`
	Branch                     string                 `protobuf:"bytes,9,opt,name=branch,proto3" json:"branch,omitempty"`
	CardNo                     string                 `protobuf:"bytes,10,opt,name=cardNo,proto3" json:"cardNo,omitempty"`
	CertifiedBusinessApplyTime int64                  `protobuf:"varint,11,opt,name=certifiedBusinessApplyTime,proto3" json:"certifiedBusinessApplyTime,omitempty"`
	CertifiedBusinessCheckTime int64                  `protobuf:"varint,12,opt,name=certifiedBusinessCheckTime,proto3" json:"certifiedBusinessCheckTime,omitempty"`
	CertifiedBusinessStatus    int32                  `protobuf:"varint,13,opt,name=certifiedBusinessStatus,proto3" json:"certifiedBusinessStatus,omitempty"`
	ChannelId                  int32                  `protobuf:"varint,14,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Email                      string                 `protobuf:"bytes,15,opt,name=email,proto3" json:"email,omitempty"`
	FirstLevel                 int32                  `protobuf:"varint,16,opt,name=firstLevel,proto3" json:"firstLevel,omitempty"`
	GoogleDate                 int64                  `protobuf:"varint,17,opt,name=googleDate,proto3" json:"googleDate,omitempty"`
	GoogleKey                  string                 `protobuf:"bytes,18,opt,name=googleKey,proto3" json:"googleKey,omitempty"`
	GoogleState                int32                  `protobuf:"varint,19,opt,name=googleState,proto3" json:"googleState,omitempty"`
	IdNumber                   string                 `protobuf:"bytes,20,opt,name=idNumber,proto3" json:"idNumber,omitempty"`
	InviterId                  int64                  `protobuf:"varint,21,opt,name=inviterId,proto3" json:"inviterId,omitempty"`
	IsChannel                  int32                  `protobuf:"varint,22,opt,name=isChannel,proto3" json:"isChannel,omitempty"`
	JyPassword                 string                 `protobuf:"bytes,23,opt,name=jyPassword,proto3" json:"jyPassword,omitempty"`
	LastLoginTime              int64                  `protobuf:"varint,24,opt,name=lastLoginTime,proto3" json:"lastLoginTime,omitempty"`
	City                       string                 `protobuf:"bytes,25,opt,name=city,proto3" json:"city,omitempty"`
	Country                    string                 `protobuf:"bytes,26,opt,name=country,proto3" json:"country,omitempty"`
	District                   string                 `protobuf:"bytes,27,opt,name=district,proto3" json:"district,omitempty"`
	Province                   string                 `protobuf:"bytes,28,opt,name=province,proto3" json:"province,omitempty"`
	LoginCount                 int32                  `protobuf:"varint,29,opt,name=loginCount,proto3" json:"loginCount,omitempty"`
	LoginLock                  int32                  `protobuf:"varint,30,opt,name=loginLock,proto3" json:"loginLock,omitempty"`
	Margin                     string                 `protobuf:"bytes,31,opt,name=margin,proto3" json:"margin,omitempty"`
	MemberLevel                int32                  `protobuf:"varint,32,opt,name=memberLevel,proto3" json:"memberLevel,omitempty"`
	MobilePhone                string                 `protobuf:"bytes,33,opt,name=mobilePhone,proto3" json:"mobilePhone,omitempty"`
	Password                   string                 `protobuf:"bytes,34,opt,name=password,proto3" json:"password,omitempty"`
	PromotionCode              string                 `protobuf:"bytes,35,opt,name=promotionCode,proto3" json:"promotionCode,omitempty"`
	PublishAdvertise           int32                  `protobuf:"varint,36,opt,name=publishAdvertise,proto3" json:"publishAdvertise,omitempty"`
	RealName                   string                 `protobuf:"bytes,37,opt,name=realName,proto3" json:"realName,omitempty"`
	RealNameStatus             int32                  `protobuf:"varint,38,opt,name=realNameStatus,proto3" json:"realNameStatus,omitempty"`
	RegistrationTime           int64                  `protobuf:"varint,39,opt,name=registrationTime,proto3" json:"registrationTime,omitempty"`
	Salt                       string                 `protobuf:"bytes,40,opt,name=salt,proto3" json:"salt,omitempty"`
	SecondLevel                int32                  `protobuf:"varint,41,opt,name=secondLevel,proto3" json:"secondLevel,omitempty"`
	SignInAbility              int32                  `protobuf:"varint,42,opt,name=signInAbility,proto3" json:"signInAbility,omitempty"`
	Status                     int32                  `protobuf:"varint,43,opt,name=status,proto3" json:"status,omitempty"`
	ThirdLevel                 int32                  `protobuf:"varint,44,opt,name=thirdLevel,proto3" json:"thirdLevel,omitempty"`
	Token                      string                 `protobuf:"bytes,45,opt,name=token,proto3" json:"token,omitempty"`
	TokenExpireTime            int64                  `protobuf:"varint,46,opt,name=tokenExpireTime,proto3" json:"tokenExpireTime,omitempty"`
	TransactionStatus          int32                  `protobuf:"varint,47,opt,name=transactionStatus,proto3" json:"transactionStatus,omitempty"`
	TransactionTime            int64                  `protobuf:"varint,48,opt,name=transactionTime,proto3" json:"transactionTime,omitempty"`
	Transactions               int32                  `protobuf:"varint,49,opt,name=transactions,proto3" json:"transactions,omitempty"`
	Username                   string                 `protobuf:"bytes,50,opt,name=username,proto3" json:"username,omitempty"`
	QrWeCodeUrl                string                 `protobuf:"bytes,51,opt,name=qrWeCodeUrl,proto3" json:"qrWeCodeUrl,omitempty"`
	Wechat                     string                 `protobuf:"bytes,52,opt,name=wechat,proto3" json:"wechat,omitempty"`
	Local                      string                 `protobuf:"bytes,53,opt,name=local,proto3" json:"local,omitempty"`
	Integration                int64                  `protobuf:"varint,54,opt,name=integration,proto3" json:"integration,omitempty"`
	MemberGradeId              int64                  `protobuf:"varint,55,opt,name=memberGradeId,proto3" json:"memberGradeId,omitempty"`
	KycStatus                  int32                  `protobuf:"varint,56,opt,name=kycStatus,proto3" json:"kycStatus,omitempty"`
	GeneralizeTotal            int64                  `protobuf:"varint,57,opt,name=generalizeTotal,proto3" json:"generalizeTotal,omitempty"`
	InviterParentId            int64                  `protobuf:"varint,58,opt,name=inviterParentId,proto3" json:"inviterParentId,omitempty"`
	SuperPartner               string                 `protobuf:"bytes,59,opt,name=superPartner,proto3" json:"superPartner,omitempty"`
	KickFee                    float64                `protobuf:"fixed64,60,opt,name=kickFee,proto3" json:"kickFee,omitempty"`
	Power                      float64                `protobuf:"fixed64,61,opt,name=power,proto3" json:"power,omitempty"`
	TeamLevel                  int32                  `protobuf:"varint,62,opt,name=teamLevel,proto3" json:"teamLevel,omitempty"`
	TeamPower                  float64                `protobuf:"fixed64,63,opt,name=teamPower,proto3" json:"teamPower,omitempty"`
	MemberLevelId              int64                  `protobuf:"varint,64,opt,name=memberLevelId,proto3" json:"memberLevelId,omitempty"`
	unknownFields              protoimpl.UnknownFields
	sizeCache                  protoimpl.SizeCache
}

func (x *MemberInfo) Reset() {
	*x = MemberInfo{}
	mi := &file_member_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MemberInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberInfo) ProtoMessage() {}

func (x *MemberInfo) ProtoReflect() protoreflect.Message {
	mi := &file_member_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberInfo.ProtoReflect.Descriptor instead.
func (*MemberInfo) Descriptor() ([]byte, []int) {
	return file_member_proto_rawDescGZIP(), []int{1}
}

func (x *MemberInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *MemberInfo) GetAliNo() string {
	if x != nil {
		return x.AliNo
	}
	return ""
}

func (x *MemberInfo) GetQrCodeUrl() string {
	if x != nil {
		return x.QrCodeUrl
	}
	return ""
}

func (x *MemberInfo) GetAppealSuccessTimes() int32 {
	if x != nil {
		return x.AppealSuccessTimes
	}
	return 0
}

func (x *MemberInfo) GetAppealTimes() int32 {
	if x != nil {
		return x.AppealTimes
	}
	return 0
}

func (x *MemberInfo) GetApplicationTime() int64 {
	if x != nil {
		return x.ApplicationTime
	}
	return 0
}

func (x *MemberInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *MemberInfo) GetBank() string {
	if x != nil {
		return x.Bank
	}
	return ""
}

func (x *MemberInfo) GetBranch() string {
	if x != nil {
		return x.Branch
	}
	return ""
}

func (x *MemberInfo) GetCardNo() string {
	if x != nil {
		return x.CardNo
	}
	return ""
}

func (x *MemberInfo) GetCertifiedBusinessApplyTime() int64 {
	if x != nil {
		return x.CertifiedBusinessApplyTime
	}
	return 0
}

func (x *MemberInfo) GetCertifiedBusinessCheckTime() int64 {
	if x != nil {
		return x.CertifiedBusinessCheckTime
	}
	return 0
}

func (x *MemberInfo) GetCertifiedBusinessStatus() int32 {
	if x != nil {
		return x.CertifiedBusinessStatus
	}
	return 0
}

func (x *MemberInfo) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *MemberInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *MemberInfo) GetFirstLevel() int32 {
	if x != nil {
		return x.FirstLevel
	}
	return 0
}

func (x *MemberInfo) GetGoogleDate() int64 {
	if x != nil {
		return x.GoogleDate
	}
	return 0
}

func (x *MemberInfo) GetGoogleKey() string {
	if x != nil {
		return x.GoogleKey
	}
	return ""
}

func (x *MemberInfo) GetGoogleState() int32 {
	if x != nil {
		return x.GoogleState
	}
	return 0
}

func (x *MemberInfo) GetIdNumber() string {
	if x != nil {
		return x.IdNumber
	}
	return ""
}

func (x *MemberInfo) GetInviterId() int64 {
	if x != nil {
		return x.InviterId
	}
	return 0
}

func (x *MemberInfo) GetIsChannel() int32 {
	if x != nil {
		return x.IsChannel
	}
	return 0
}

func (x *MemberInfo) GetJyPassword() string {
	if x != nil {
		return x.JyPassword
	}
	return ""
}

func (x *MemberInfo) GetLastLoginTime() int64 {
	if x != nil {
		return x.LastLoginTime
	}
	return 0
}

func (x *MemberInfo) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *MemberInfo) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *MemberInfo) GetDistrict() string {
	if x != nil {
		return x.District
	}
	return ""
}

func (x *MemberInfo) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *MemberInfo) GetLoginCount() int32 {
	if x != nil {
		return x.LoginCount
	}
	return 0
}

func (x *MemberInfo) GetLoginLock() int32 {
	if x != nil {
		return x.LoginLock
	}
	return 0
}

func (x *MemberInfo) GetMargin() string {
	if x != nil {
		return x.Margin
	}
	return ""
}

func (x *MemberInfo) GetMemberLevel() int32 {
	if x != nil {
		return x.MemberLevel
	}
	return 0
}

func (x *MemberInfo) GetMobilePhone() string {
	if x != nil {
		return x.MobilePhone
	}
	return ""
}

func (x *MemberInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *MemberInfo) GetPromotionCode() string {
	if x != nil {
		return x.PromotionCode
	}
	return ""
}

func (x *MemberInfo) GetPublishAdvertise() int32 {
	if x != nil {
		return x.PublishAdvertise
	}
	return 0
}

func (x *MemberInfo) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *MemberInfo) GetRealNameStatus() int32 {
	if x != nil {
		return x.RealNameStatus
	}
	return 0
}

func (x *MemberInfo) GetRegistrationTime() int64 {
	if x != nil {
		return x.RegistrationTime
	}
	return 0
}

func (x *MemberInfo) GetSalt() string {
	if x != nil {
		return x.Salt
	}
	return ""
}

func (x *MemberInfo) GetSecondLevel() int32 {
	if x != nil {
		return x.SecondLevel
	}
	return 0
}

func (x *MemberInfo) GetSignInAbility() int32 {
	if x != nil {
		return x.SignInAbility
	}
	return 0
}

func (x *MemberInfo) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *MemberInfo) GetThirdLevel() int32 {
	if x != nil {
		return x.ThirdLevel
	}
	return 0
}

func (x *MemberInfo) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *MemberInfo) GetTokenExpireTime() int64 {
	if x != nil {
		return x.TokenExpireTime
	}
	return 0
}

func (x *MemberInfo) GetTransactionStatus() int32 {
	if x != nil {
		return x.TransactionStatus
	}
	return 0
}

func (x *MemberInfo) GetTransactionTime() int64 {
	if x != nil {
		return x.TransactionTime
	}
	return 0
}

func (x *MemberInfo) GetTransactions() int32 {
	if x != nil {
		return x.Transactions
	}
	return 0
}

func (x *MemberInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *MemberInfo) GetQrWeCodeUrl() string {
	if x != nil {
		return x.QrWeCodeUrl
	}
	return ""
}

func (x *MemberInfo) GetWechat() string {
	if x != nil {
		return x.Wechat
	}
	return ""
}

func (x *MemberInfo) GetLocal() string {
	if x != nil {
		return x.Local
	}
	return ""
}

func (x *MemberInfo) GetIntegration() int64 {
	if x != nil {
		return x.Integration
	}
	return 0
}

func (x *MemberInfo) GetMemberGradeId() int64 {
	if x != nil {
		return x.MemberGradeId
	}
	return 0
}

func (x *MemberInfo) GetKycStatus() int32 {
	if x != nil {
		return x.KycStatus
	}
	return 0
}

func (x *MemberInfo) GetGeneralizeTotal() int64 {
	if x != nil {
		return x.GeneralizeTotal
	}
	return 0
}

func (x *MemberInfo) GetInviterParentId() int64 {
	if x != nil {
		return x.InviterParentId
	}
	return 0
}

func (x *MemberInfo) GetSuperPartner() string {
	if x != nil {
		return x.SuperPartner
	}
	return ""
}

func (x *MemberInfo) GetKickFee() float64 {
	if x != nil {
		return x.KickFee
	}
	return 0
}

func (x *MemberInfo) GetPower() float64 {
	if x != nil {
		return x.Power
	}
	return 0
}

func (x *MemberInfo) GetTeamLevel() int32 {
	if x != nil {
		return x.TeamLevel
	}
	return 0
}

func (x *MemberInfo) GetTeamPower() float64 {
	if x != nil {
		return x.TeamPower
	}
	return 0
}

func (x *MemberInfo) GetMemberLevelId() int64 {
	if x != nil {
		return x.MemberLevelId
	}
	return 0
}

var File_member_proto protoreflect.FileDescriptor

var file_member_proto_rawDesc = string([]byte{
	0x0a, 0x0c, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x27, 0x0a, 0x09, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x64, 0x22,
	0xbe, 0x10, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x61, 0x6c, 0x69, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61,
	0x6c, 0x69, 0x4e, 0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x55, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x2e, 0x0a, 0x12, 0x61, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12,
	0x61, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x65, 0x61, 0x6c, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x61, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x72, 0x61, 0x6e,
	0x63, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x12, 0x3e, 0x0a, 0x1a, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1a,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x1a, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1a,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x17, 0x63, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x17, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x65, 0x64, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49,
	0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0f, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x10, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x16, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x73, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x12, 0x1e, 0x0a, 0x0a, 0x6a, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x17,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6a, 0x79, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x18, 0x1b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x1c, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1d, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x61,
	0x72, 0x67, 0x69, 0x6e, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x61, 0x72, 0x67,
	0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x20, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x21, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x22, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x18, 0x23, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x72, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x41, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x18, 0x24, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x10, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x41, 0x64, 0x76, 0x65, 0x72,
	0x74, 0x69, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x25, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x26, 0x0a, 0x0e, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x26, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0e, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x27, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x73, 0x61, 0x6c, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x6f,
	0x6e, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x29, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73,
	0x65, 0x63, 0x6f, 0x6e, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x69,
	0x67, 0x6e, 0x49, 0x6e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x2a, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x41, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x2b, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x68, 0x69, 0x72,
	0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x2c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x68,
	0x69, 0x72, 0x64, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x2d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28,
	0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d,
	0x65, 0x18, 0x2e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x2f, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x30, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x31, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x71, 0x72, 0x57, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x55, 0x72, 0x6c, 0x18,
	0x33, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x72, 0x57, 0x65, 0x43, 0x6f, 0x64, 0x65, 0x55,
	0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x18, 0x34, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x77, 0x65, 0x63, 0x68, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f,
	0x63, 0x61, 0x6c, 0x18, 0x35, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x63, 0x61, 0x6c,
	0x12, 0x20, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x36, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x72, 0x61, 0x64,
	0x65, 0x49, 0x64, 0x18, 0x37, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x47, 0x72, 0x61, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x6b, 0x79, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x38, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6b, 0x79, 0x63,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x39, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x12, 0x28, 0x0a, 0x0f, 0x69, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x18, 0x3a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x69, 0x6e, 0x76, 0x69, 0x74,
	0x65, 0x72, 0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x75,
	0x70, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x3b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x73, 0x75, 0x70, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6b, 0x69, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x18, 0x3c, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x07, 0x6b, 0x69, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x77, 0x65,
	0x72, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x3e, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x65, 0x61, 0x6d, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x3f, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x09, 0x74, 0x65, 0x61, 0x6d, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x40, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x49, 0x64,
	0x32, 0x41, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x37, 0x0a, 0x0e, 0x46, 0x69,
	0x6e, 0x64, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x11, 0x2e, 0x6d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x12, 0x2e, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_member_proto_rawDescOnce sync.Once
	file_member_proto_rawDescData []byte
)

func file_member_proto_rawDescGZIP() []byte {
	file_member_proto_rawDescOnce.Do(func() {
		file_member_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_member_proto_rawDesc), len(file_member_proto_rawDesc)))
	})
	return file_member_proto_rawDescData
}

var file_member_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_member_proto_goTypes = []any{
	(*MemberReq)(nil),  // 0: member.MemberReq
	(*MemberInfo)(nil), // 1: member.MemberInfo
}
var file_member_proto_depIdxs = []int32{
	0, // 0: member.Member.FindMemberById:input_type -> member.MemberReq
	1, // 1: member.Member.FindMemberById:output_type -> member.MemberInfo
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_member_proto_init() }
func file_member_proto_init() {
	if File_member_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_member_proto_rawDesc), len(file_member_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_member_proto_goTypes,
		DependencyIndexes: file_member_proto_depIdxs,
		MessageInfos:      file_member_proto_msgTypes,
	}.Build()
	File_member_proto = out.File
	file_member_proto_goTypes = nil
	file_member_proto_depIdxs = nil
}
