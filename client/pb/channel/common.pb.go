// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/channel/common.proto

package channel

import (
	msg "github.com/Mrs4s/MiraiGo/client/pb/msg"
)

type ChannelContentHead struct {
	Type    *uint64 `protobuf:"varint,1,opt"`
	SubType *uint64 `protobuf:"varint,2,opt"`
	Random  *uint64 `protobuf:"varint,3,opt"`
	Seq     *uint64 `protobuf:"varint,4,opt"`
	CntSeq  *uint64 `protobuf:"varint,5,opt"`
	Time    *uint64 `protobuf:"varint,6,opt"`
	Meta    []byte  `protobuf:"bytes,7,opt"`
}

func (x *ChannelContentHead) GetType() uint64 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *ChannelContentHead) GetSubType() uint64 {
	if x != nil && x.SubType != nil {
		return *x.SubType
	}
	return 0
}

func (x *ChannelContentHead) GetRandom() uint64 {
	if x != nil && x.Random != nil {
		return *x.Random
	}
	return 0
}

func (x *ChannelContentHead) GetSeq() uint64 {
	if x != nil && x.Seq != nil {
		return *x.Seq
	}
	return 0
}

func (x *ChannelContentHead) GetCntSeq() uint64 {
	if x != nil && x.CntSeq != nil {
		return *x.CntSeq
	}
	return 0
}

func (x *ChannelContentHead) GetTime() uint64 {
	if x != nil && x.Time != nil {
		return *x.Time
	}
	return 0
}

func (x *ChannelContentHead) GetMeta() []byte {
	if x != nil {
		return x.Meta
	}
	return nil
}

type DirectMessageMember struct {
	Uin             *uint64 `protobuf:"varint,1,opt"`
	Tinyid          *uint64 `protobuf:"varint,2,opt"`
	SourceGuildId   *uint64 `protobuf:"varint,3,opt"`
	SourceGuildName []byte  `protobuf:"bytes,4,opt"`
	NickName        []byte  `protobuf:"bytes,5,opt"`
	MemberName      []byte  `protobuf:"bytes,6,opt"`
	NotifyType      *uint32 `protobuf:"varint,7,opt"`
}

func (x *DirectMessageMember) GetUin() uint64 {
	if x != nil && x.Uin != nil {
		return *x.Uin
	}
	return 0
}

func (x *DirectMessageMember) GetTinyid() uint64 {
	if x != nil && x.Tinyid != nil {
		return *x.Tinyid
	}
	return 0
}

func (x *DirectMessageMember) GetSourceGuildId() uint64 {
	if x != nil && x.SourceGuildId != nil {
		return *x.SourceGuildId
	}
	return 0
}

func (x *DirectMessageMember) GetSourceGuildName() []byte {
	if x != nil {
		return x.SourceGuildName
	}
	return nil
}

func (x *DirectMessageMember) GetNickName() []byte {
	if x != nil {
		return x.NickName
	}
	return nil
}

func (x *DirectMessageMember) GetMemberName() []byte {
	if x != nil {
		return x.MemberName
	}
	return nil
}

func (x *DirectMessageMember) GetNotifyType() uint32 {
	if x != nil && x.NotifyType != nil {
		return *x.NotifyType
	}
	return 0
}

type ChannelEvent struct {
	Type    *uint64           `protobuf:"varint,1,opt"`
	Version *uint64           `protobuf:"varint,2,opt"`
	OpInfo  *ChannelMsgOpInfo `protobuf:"bytes,3,opt"`
}

func (x *ChannelEvent) GetType() uint64 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *ChannelEvent) GetVersion() uint64 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *ChannelEvent) GetOpInfo() *ChannelMsgOpInfo {
	if x != nil {
		return x.OpInfo
	}
	return nil
}

type ChannelExtInfo struct {
	FromNick            []byte                 `protobuf:"bytes,1,opt"`
	GuildName           []byte                 `protobuf:"bytes,2,opt"`
	ChannelName         []byte                 `protobuf:"bytes,3,opt"`
	Visibility          *uint32                `protobuf:"varint,4,opt"`
	NotifyType          *uint32                `protobuf:"varint,5,opt"`
	OfflineFlag         *uint32                `protobuf:"varint,6,opt"`
	NameType            *uint32                `protobuf:"varint,7,opt"`
	MemberName          []byte                 `protobuf:"bytes,8,opt"`
	Timestamp           *uint32                `protobuf:"varint,9,opt"`
	EventVersion        *uint64                `protobuf:"varint,10,opt"`
	Events              []*ChannelEvent        `protobuf:"bytes,11,rep"`
	FromRoleInfo        *ChannelRole           `protobuf:"bytes,12,opt"`
	FreqLimitInfo       *ChannelFreqLimitInfo  `protobuf:"bytes,13,opt"`
	DirectMessageMember []*DirectMessageMember `protobuf:"bytes,14,rep"`
}

func (x *ChannelExtInfo) GetFromNick() []byte {
	if x != nil {
		return x.FromNick
	}
	return nil
}

func (x *ChannelExtInfo) GetGuildName() []byte {
	if x != nil {
		return x.GuildName
	}
	return nil
}

func (x *ChannelExtInfo) GetChannelName() []byte {
	if x != nil {
		return x.ChannelName
	}
	return nil
}

func (x *ChannelExtInfo) GetVisibility() uint32 {
	if x != nil && x.Visibility != nil {
		return *x.Visibility
	}
	return 0
}

func (x *ChannelExtInfo) GetNotifyType() uint32 {
	if x != nil && x.NotifyType != nil {
		return *x.NotifyType
	}
	return 0
}

func (x *ChannelExtInfo) GetOfflineFlag() uint32 {
	if x != nil && x.OfflineFlag != nil {
		return *x.OfflineFlag
	}
	return 0
}

func (x *ChannelExtInfo) GetNameType() uint32 {
	if x != nil && x.NameType != nil {
		return *x.NameType
	}
	return 0
}

func (x *ChannelExtInfo) GetMemberName() []byte {
	if x != nil {
		return x.MemberName
	}
	return nil
}

func (x *ChannelExtInfo) GetTimestamp() uint32 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *ChannelExtInfo) GetEventVersion() uint64 {
	if x != nil && x.EventVersion != nil {
		return *x.EventVersion
	}
	return 0
}

func (x *ChannelExtInfo) GetEvents() []*ChannelEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *ChannelExtInfo) GetFromRoleInfo() *ChannelRole {
	if x != nil {
		return x.FromRoleInfo
	}
	return nil
}

func (x *ChannelExtInfo) GetFreqLimitInfo() *ChannelFreqLimitInfo {
	if x != nil {
		return x.FreqLimitInfo
	}
	return nil
}

func (x *ChannelExtInfo) GetDirectMessageMember() []*DirectMessageMember {
	if x != nil {
		return x.DirectMessageMember
	}
	return nil
}

type ChannelFreqLimitInfo struct {
	IsLimited      *uint32 `protobuf:"varint,1,opt"`
	LeftCount      *uint32 `protobuf:"varint,2,opt"`
	LimitTimestamp *uint64 `protobuf:"varint,3,opt"`
}

func (x *ChannelFreqLimitInfo) GetIsLimited() uint32 {
	if x != nil && x.IsLimited != nil {
		return *x.IsLimited
	}
	return 0
}

func (x *ChannelFreqLimitInfo) GetLeftCount() uint32 {
	if x != nil && x.LeftCount != nil {
		return *x.LeftCount
	}
	return 0
}

func (x *ChannelFreqLimitInfo) GetLimitTimestamp() uint64 {
	if x != nil && x.LimitTimestamp != nil {
		return *x.LimitTimestamp
	}
	return 0
}

type ChannelInfo struct {
	Id    *uint64 `protobuf:"varint,1,opt"`
	Name  []byte  `protobuf:"bytes,2,opt"`
	Color *uint32 `protobuf:"varint,3,opt"`
	Hoist *uint32 `protobuf:"varint,4,opt"`
}

func (x *ChannelInfo) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *ChannelInfo) GetName() []byte {
	if x != nil {
		return x.Name
	}
	return nil
}

func (x *ChannelInfo) GetColor() uint32 {
	if x != nil && x.Color != nil {
		return *x.Color
	}
	return 0
}

func (x *ChannelInfo) GetHoist() uint32 {
	if x != nil && x.Hoist != nil {
		return *x.Hoist
	}
	return 0
}

type ChannelLoginSig struct {
	Type  *uint32 `protobuf:"varint,1,opt"`
	Sig   []byte  `protobuf:"bytes,2,opt"`
	Appid *uint32 `protobuf:"varint,3,opt"`
}

func (x *ChannelLoginSig) GetType() uint32 {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return 0
}

func (x *ChannelLoginSig) GetSig() []byte {
	if x != nil {
		return x.Sig
	}
	return nil
}

func (x *ChannelLoginSig) GetAppid() uint32 {
	if x != nil && x.Appid != nil {
		return *x.Appid
	}
	return 0
}

type ChannelMeta struct {
	FromUin  *uint64          `protobuf:"varint,1,opt"`
	LoginSig *ChannelLoginSig `protobuf:"bytes,2,opt"`
}

func (x *ChannelMeta) GetFromUin() uint64 {
	if x != nil && x.FromUin != nil {
		return *x.FromUin
	}
	return 0
}

func (x *ChannelMeta) GetLoginSig() *ChannelLoginSig {
	if x != nil {
		return x.LoginSig
	}
	return nil
}

type ChannelMsgContent struct {
	Head     *ChannelMsgHead     `protobuf:"bytes,1,opt"`
	CtrlHead *ChannelMsgCtrlHead `protobuf:"bytes,2,opt"`
	Body     *msg.MessageBody    `protobuf:"bytes,3,opt"`
	ExtInfo  *ChannelExtInfo     `protobuf:"bytes,4,opt"`
}

func (x *ChannelMsgContent) GetHead() *ChannelMsgHead {
	if x != nil {
		return x.Head
	}
	return nil
}

func (x *ChannelMsgContent) GetCtrlHead() *ChannelMsgCtrlHead {
	if x != nil {
		return x.CtrlHead
	}
	return nil
}

func (x *ChannelMsgContent) GetBody() *msg.MessageBody {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *ChannelMsgContent) GetExtInfo() *ChannelExtInfo {
	if x != nil {
		return x.ExtInfo
	}
	return nil
}

type ChannelMsgCtrlHead struct {
	IncludeUin []uint64 `protobuf:"varint,1,rep"`
	// repeated uint64 excludeUin = 2; // bytes?
	// repeated uint64 featureid = 3;
	OfflineFlag    *uint32          `protobuf:"varint,4,opt"`
	Visibility     *uint32          `protobuf:"varint,5,opt"`
	CtrlFlag       *uint64          `protobuf:"varint,6,opt"`
	Events         []*ChannelEvent  `protobuf:"bytes,7,rep"`
	Level          *uint64          `protobuf:"varint,8,opt"`
	PersonalLevels []*PersonalLevel `protobuf:"bytes,9,rep"`
	GuildSyncSeq   *uint64          `protobuf:"varint,10,opt"`
	MemberNum      *uint32          `protobuf:"varint,11,opt"`
	ChannelType    *uint32          `protobuf:"varint,12,opt"`
	PrivateType    *uint32          `protobuf:"varint,13,opt"`
}

func (x *ChannelMsgCtrlHead) GetIncludeUin() []uint64 {
	if x != nil {
		return x.IncludeUin
	}
	return nil
}

func (x *ChannelMsgCtrlHead) GetOfflineFlag() uint32 {
	if x != nil && x.OfflineFlag != nil {
		return *x.OfflineFlag
	}
	return 0
}

func (x *ChannelMsgCtrlHead) GetVisibility() uint32 {
	if x != nil && x.Visibility != nil {
		return *x.Visibility
	}
	return 0
}

func (x *ChannelMsgCtrlHead) GetCtrlFlag() uint64 {
	if x != nil && x.CtrlFlag != nil {
		return *x.CtrlFlag
	}
	return 0
}

func (x *ChannelMsgCtrlHead) GetEvents() []*ChannelEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

func (x *ChannelMsgCtrlHead) GetLevel() uint64 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

func (x *ChannelMsgCtrlHead) GetPersonalLevels() []*PersonalLevel {
	if x != nil {
		return x.PersonalLevels
	}
	return nil
}

func (x *ChannelMsgCtrlHead) GetGuildSyncSeq() uint64 {
	if x != nil && x.GuildSyncSeq != nil {
		return *x.GuildSyncSeq
	}
	return 0
}

func (x *ChannelMsgCtrlHead) GetMemberNum() uint32 {
	if x != nil && x.MemberNum != nil {
		return *x.MemberNum
	}
	return 0
}

func (x *ChannelMsgCtrlHead) GetChannelType() uint32 {
	if x != nil && x.ChannelType != nil {
		return *x.ChannelType
	}
	return 0
}

func (x *ChannelMsgCtrlHead) GetPrivateType() uint32 {
	if x != nil && x.PrivateType != nil {
		return *x.PrivateType
	}
	return 0
}

type ChannelMsgHead struct {
	RoutingHead *ChannelRoutingHead `protobuf:"bytes,1,opt"`
	ContentHead *ChannelContentHead `protobuf:"bytes,2,opt"`
}

func (x *ChannelMsgHead) GetRoutingHead() *ChannelRoutingHead {
	if x != nil {
		return x.RoutingHead
	}
	return nil
}

func (x *ChannelMsgHead) GetContentHead() *ChannelContentHead {
	if x != nil {
		return x.ContentHead
	}
	return nil
}

type ChannelMsgMeta struct {
	AtAllSeq *uint64 `protobuf:"varint,1,opt"`
}

func (x *ChannelMsgMeta) GetAtAllSeq() uint64 {
	if x != nil && x.AtAllSeq != nil {
		return *x.AtAllSeq
	}
	return 0
}

type ChannelMsgOpInfo struct {
	OperatorTinyid *uint64 `protobuf:"varint,1,opt"`
	OperatorRole   *uint64 `protobuf:"varint,2,opt"`
	Reason         *uint64 `protobuf:"varint,3,opt"`
	Timestamp      *uint64 `protobuf:"varint,4,opt"`
	AtType         *uint64 `protobuf:"varint,5,opt"`
}

func (x *ChannelMsgOpInfo) GetOperatorTinyid() uint64 {
	if x != nil && x.OperatorTinyid != nil {
		return *x.OperatorTinyid
	}
	return 0
}

func (x *ChannelMsgOpInfo) GetOperatorRole() uint64 {
	if x != nil && x.OperatorRole != nil {
		return *x.OperatorRole
	}
	return 0
}

func (x *ChannelMsgOpInfo) GetReason() uint64 {
	if x != nil && x.Reason != nil {
		return *x.Reason
	}
	return 0
}

func (x *ChannelMsgOpInfo) GetTimestamp() uint64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *ChannelMsgOpInfo) GetAtType() uint64 {
	if x != nil && x.AtType != nil {
		return *x.AtType
	}
	return 0
}

type PersonalLevel struct {
	ToUin *uint64 `protobuf:"varint,1,opt"`
	Level *uint64 `protobuf:"varint,2,opt"`
}

func (x *PersonalLevel) GetToUin() uint64 {
	if x != nil && x.ToUin != nil {
		return *x.ToUin
	}
	return 0
}

func (x *PersonalLevel) GetLevel() uint64 {
	if x != nil && x.Level != nil {
		return *x.Level
	}
	return 0
}

type ChannelRole struct {
	Id   *uint64 `protobuf:"varint,1,opt"`
	Info []byte  `protobuf:"bytes,2,opt"`
	Flag *uint32 `protobuf:"varint,3,opt"`
}

func (x *ChannelRole) GetId() uint64 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *ChannelRole) GetInfo() []byte {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *ChannelRole) GetFlag() uint32 {
	if x != nil && x.Flag != nil {
		return *x.Flag
	}
	return 0
}

type ChannelRoutingHead struct {
	GuildId           *uint64 `protobuf:"varint,1,opt"`
	ChannelId         *uint64 `protobuf:"varint,2,opt"`
	FromUin           *uint64 `protobuf:"varint,3,opt"`
	FromTinyid        *uint64 `protobuf:"varint,4,opt"`
	GuildCode         *uint64 `protobuf:"varint,5,opt"`
	FromAppid         *uint64 `protobuf:"varint,6,opt"`
	DirectMessageFlag *uint32 `protobuf:"varint,7,opt"`
}

func (x *ChannelRoutingHead) GetGuildId() uint64 {
	if x != nil && x.GuildId != nil {
		return *x.GuildId
	}
	return 0
}

func (x *ChannelRoutingHead) GetChannelId() uint64 {
	if x != nil && x.ChannelId != nil {
		return *x.ChannelId
	}
	return 0
}

func (x *ChannelRoutingHead) GetFromUin() uint64 {
	if x != nil && x.FromUin != nil {
		return *x.FromUin
	}
	return 0
}

func (x *ChannelRoutingHead) GetFromTinyid() uint64 {
	if x != nil && x.FromTinyid != nil {
		return *x.FromTinyid
	}
	return 0
}

func (x *ChannelRoutingHead) GetGuildCode() uint64 {
	if x != nil && x.GuildCode != nil {
		return *x.GuildCode
	}
	return 0
}

func (x *ChannelRoutingHead) GetFromAppid() uint64 {
	if x != nil && x.FromAppid != nil {
		return *x.FromAppid
	}
	return 0
}

func (x *ChannelRoutingHead) GetDirectMessageFlag() uint32 {
	if x != nil && x.DirectMessageFlag != nil {
		return *x.DirectMessageFlag
	}
	return 0
}
