// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: lte/protos/oai/common_types.proto

package oai

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Ambr_BitrateUnitsAMBR int32

const (
	Ambr_BPS  Ambr_BitrateUnitsAMBR = 0
	Ambr_KBPS Ambr_BitrateUnitsAMBR = 1
)

// Enum value maps for Ambr_BitrateUnitsAMBR.
var (
	Ambr_BitrateUnitsAMBR_name = map[int32]string{
		0: "BPS",
		1: "KBPS",
	}
	Ambr_BitrateUnitsAMBR_value = map[string]int32{
		"BPS":  0,
		"KBPS": 1,
	}
)

func (x Ambr_BitrateUnitsAMBR) Enum() *Ambr_BitrateUnitsAMBR {
	p := new(Ambr_BitrateUnitsAMBR)
	*p = x
	return p
}

func (x Ambr_BitrateUnitsAMBR) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Ambr_BitrateUnitsAMBR) Descriptor() protoreflect.EnumDescriptor {
	return file_lte_protos_oai_common_types_proto_enumTypes[0].Descriptor()
}

func (Ambr_BitrateUnitsAMBR) Type() protoreflect.EnumType {
	return &file_lte_protos_oai_common_types_proto_enumTypes[0]
}

func (x Ambr_BitrateUnitsAMBR) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Ambr_BitrateUnitsAMBR.Descriptor instead.
func (Ambr_BitrateUnitsAMBR) EnumDescriptor() ([]byte, []int) {
	return file_lte_protos_oai_common_types_proto_rawDescGZIP(), []int{3, 0}
}

// guti_t
type Guti struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plmn    []byte `protobuf:"bytes,1,opt,name=plmn,proto3" json:"plmn,omitempty"`
	MmeGid  uint32 `protobuf:"varint,2,opt,name=mme_gid,json=mmeGid,proto3" json:"mme_gid,omitempty"`
	MmeCode uint32 `protobuf:"varint,3,opt,name=mme_code,json=mmeCode,proto3" json:"mme_code,omitempty"`
	MTmsi   uint32 `protobuf:"varint,4,opt,name=m_tmsi,json=mTmsi,proto3" json:"m_tmsi,omitempty"`
}

func (x *Guti) Reset() {
	*x = Guti{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_common_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Guti) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Guti) ProtoMessage() {}

func (x *Guti) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_common_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Guti.ProtoReflect.Descriptor instead.
func (*Guti) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_common_types_proto_rawDescGZIP(), []int{0}
}

func (x *Guti) GetPlmn() []byte {
	if x != nil {
		return x.Plmn
	}
	return nil
}

func (x *Guti) GetMmeGid() uint32 {
	if x != nil {
		return x.MmeGid
	}
	return 0
}

func (x *Guti) GetMmeCode() uint32 {
	if x != nil {
		return x.MmeCode
	}
	return 0
}

func (x *Guti) GetMTmsi() uint32 {
	if x != nil {
		return x.MTmsi
	}
	return 0
}

// ecgi_t
type Ecgi struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plmn []byte `protobuf:"bytes,1,opt,name=plmn,proto3" json:"plmn,omitempty"` //char[] as plmn_t is 3 bytes
	// serializing struct eci_t here without creating a new message
	EnbId  uint32 `protobuf:"varint,2,opt,name=enb_id,json=enbId,proto3" json:"enb_id,omitempty"`
	CellId uint32 `protobuf:"varint,3,opt,name=cell_id,json=cellId,proto3" json:"cell_id,omitempty"`
	Empty  uint32 `protobuf:"varint,4,opt,name=empty,proto3" json:"empty,omitempty"`
}

func (x *Ecgi) Reset() {
	*x = Ecgi{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_common_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ecgi) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ecgi) ProtoMessage() {}

func (x *Ecgi) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_common_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ecgi.ProtoReflect.Descriptor instead.
func (*Ecgi) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_common_types_proto_rawDescGZIP(), []int{1}
}

func (x *Ecgi) GetPlmn() []byte {
	if x != nil {
		return x.Plmn
	}
	return nil
}

func (x *Ecgi) GetEnbId() uint32 {
	if x != nil {
		return x.EnbId
	}
	return 0
}

func (x *Ecgi) GetCellId() uint32 {
	if x != nil {
		return x.CellId
	}
	return 0
}

func (x *Ecgi) GetEmpty() uint32 {
	if x != nil {
		return x.Empty
	}
	return 0
}

// eps_subscribed_qos_profile_t
type EpsSubscribedQosProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Qci uint32 `protobuf:"varint,1,opt,name=qci,proto3" json:"qci,omitempty"` // qci_t
	// serializing allocation_retention_priority_t here instead of adding a
	// new message type
	PriorityLevel           uint32 `protobuf:"varint,2,opt,name=priority_level,json=priorityLevel,proto3" json:"priority_level,omitempty"`                                 // priority_level_t
	PreEmptionVulnerability uint32 `protobuf:"varint,3,opt,name=pre_emption_vulnerability,json=preEmptionVulnerability,proto3" json:"pre_emption_vulnerability,omitempty"` // pre_emption_vulnerability_t
	PreEmptionCapability    uint32 `protobuf:"varint,4,opt,name=pre_emption_capability,json=preEmptionCapability,proto3" json:"pre_emption_capability,omitempty"`          // pre_emption_capability_t
}

func (x *EpsSubscribedQosProfile) Reset() {
	*x = EpsSubscribedQosProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_common_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EpsSubscribedQosProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EpsSubscribedQosProfile) ProtoMessage() {}

func (x *EpsSubscribedQosProfile) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_common_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EpsSubscribedQosProfile.ProtoReflect.Descriptor instead.
func (*EpsSubscribedQosProfile) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_common_types_proto_rawDescGZIP(), []int{2}
}

func (x *EpsSubscribedQosProfile) GetQci() uint32 {
	if x != nil {
		return x.Qci
	}
	return 0
}

func (x *EpsSubscribedQosProfile) GetPriorityLevel() uint32 {
	if x != nil {
		return x.PriorityLevel
	}
	return 0
}

func (x *EpsSubscribedQosProfile) GetPreEmptionVulnerability() uint32 {
	if x != nil {
		return x.PreEmptionVulnerability
	}
	return 0
}

func (x *EpsSubscribedQosProfile) GetPreEmptionCapability() uint32 {
	if x != nil {
		return x.PreEmptionCapability
	}
	return 0
}

// ambr_t
type Ambr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrUl uint64 `protobuf:"varint,1,opt,name=br_ul,json=brUl,proto3" json:"br_ul,omitempty"`
	BrDl uint64 `protobuf:"varint,2,opt,name=br_dl,json=brDl,proto3" json:"br_dl,omitempty"`
	// Unit (either bps or Kbps)
	BrUnit Ambr_BitrateUnitsAMBR `protobuf:"varint,3,opt,name=br_unit,json=brUnit,proto3,enum=magma.lte.oai.Ambr_BitrateUnitsAMBR" json:"br_unit,omitempty"`
}

func (x *Ambr) Reset() {
	*x = Ambr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_common_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ambr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ambr) ProtoMessage() {}

func (x *Ambr) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_common_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ambr.ProtoReflect.Descriptor instead.
func (*Ambr) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_common_types_proto_rawDescGZIP(), []int{3}
}

func (x *Ambr) GetBrUl() uint64 {
	if x != nil {
		return x.BrUl
	}
	return 0
}

func (x *Ambr) GetBrDl() uint64 {
	if x != nil {
		return x.BrDl
	}
	return 0
}

func (x *Ambr) GetBrUnit() Ambr_BitrateUnitsAMBR {
	if x != nil {
		return x.BrUnit
	}
	return Ambr_BPS
}

// apn_configuration_t
type ApnConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContextIdentifier uint32                   `protobuf:"varint,1,opt,name=context_identifier,json=contextIdentifier,proto3" json:"context_identifier,omitempty"`
	IpAddress         []string                 `protobuf:"bytes,2,rep,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	PdnType           uint32                   `protobuf:"varint,3,opt,name=pdn_type,json=pdnType,proto3" json:"pdn_type,omitempty"`
	ServiceSelection  []byte                   `protobuf:"bytes,4,opt,name=service_selection,json=serviceSelection,proto3" json:"service_selection,omitempty"`
	SubscribedQos     *EpsSubscribedQosProfile `protobuf:"bytes,5,opt,name=subscribed_qos,json=subscribedQos,proto3" json:"subscribed_qos,omitempty"`
	Ambr              *Ambr                    `protobuf:"bytes,6,opt,name=ambr,proto3" json:"ambr,omitempty"`
}

func (x *ApnConfig) Reset() {
	*x = ApnConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_common_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApnConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApnConfig) ProtoMessage() {}

func (x *ApnConfig) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_common_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApnConfig.ProtoReflect.Descriptor instead.
func (*ApnConfig) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_common_types_proto_rawDescGZIP(), []int{4}
}

func (x *ApnConfig) GetContextIdentifier() uint32 {
	if x != nil {
		return x.ContextIdentifier
	}
	return 0
}

func (x *ApnConfig) GetIpAddress() []string {
	if x != nil {
		return x.IpAddress
	}
	return nil
}

func (x *ApnConfig) GetPdnType() uint32 {
	if x != nil {
		return x.PdnType
	}
	return 0
}

func (x *ApnConfig) GetServiceSelection() []byte {
	if x != nil {
		return x.ServiceSelection
	}
	return nil
}

func (x *ApnConfig) GetSubscribedQos() *EpsSubscribedQosProfile {
	if x != nil {
		return x.SubscribedQos
	}
	return nil
}

func (x *ApnConfig) GetAmbr() *Ambr {
	if x != nil {
		return x.Ambr
	}
	return nil
}

// apn_config_profile
type ApnConfigProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContextIdentifier uint32       `protobuf:"varint,1,opt,name=context_identifier,json=contextIdentifier,proto3" json:"context_identifier,omitempty"`
	AllApnConfInd     uint32       `protobuf:"varint,2,opt,name=all_apn_conf_ind,json=allApnConfInd,proto3" json:"all_apn_conf_ind,omitempty"` // all_apn_conf_ind_t
	ApnConfigs        []*ApnConfig `protobuf:"bytes,3,rep,name=apn_configs,json=apnConfigs,proto3" json:"apn_configs,omitempty"`
}

func (x *ApnConfigProfile) Reset() {
	*x = ApnConfigProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_common_types_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApnConfigProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApnConfigProfile) ProtoMessage() {}

func (x *ApnConfigProfile) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_common_types_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApnConfigProfile.ProtoReflect.Descriptor instead.
func (*ApnConfigProfile) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_common_types_proto_rawDescGZIP(), []int{5}
}

func (x *ApnConfigProfile) GetContextIdentifier() uint32 {
	if x != nil {
		return x.ContextIdentifier
	}
	return 0
}

func (x *ApnConfigProfile) GetAllApnConfInd() uint32 {
	if x != nil {
		return x.AllApnConfInd
	}
	return 0
}

func (x *ApnConfigProfile) GetApnConfigs() []*ApnConfig {
	if x != nil {
		return x.ApnConfigs
	}
	return nil
}

var File_lte_protos_oai_common_types_proto protoreflect.FileDescriptor

var file_lte_protos_oai_common_types_proto_rawDesc = []byte{
	0x0a, 0x21, 0x6c, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x61, 0x69,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f,
	0x61, 0x69, 0x22, 0x65, 0x0a, 0x04, 0x47, 0x75, 0x74, 0x69, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c,
	0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x70, 0x6c, 0x6d, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x6d, 0x6d, 0x65, 0x5f, 0x67, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x06, 0x6d, 0x6d, 0x65, 0x47, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6d, 0x65, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6d, 0x6d, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x6d, 0x5f, 0x74, 0x6d, 0x73, 0x69, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x05, 0x6d, 0x54, 0x6d, 0x73, 0x69, 0x22, 0x60, 0x0a, 0x04, 0x45, 0x63, 0x67,
	0x69, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x6d, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x70, 0x6c, 0x6d, 0x6e, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x62, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x6e, 0x62, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x63,
	0x65, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xc4, 0x01, 0x0a, 0x17,
	0x45, 0x70, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x51, 0x6f, 0x73,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x71, 0x63, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x71, 0x63, 0x69, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x72, 0x69,
	0x6f, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x0d, 0x70, 0x72, 0x69, 0x6f, 0x72, 0x69, 0x74, 0x79, 0x4c, 0x65, 0x76, 0x65, 0x6c,
	0x12, 0x3a, 0x0a, 0x19, 0x70, 0x72, 0x65, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x76, 0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x17, 0x70, 0x72, 0x65, 0x45, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x56,
	0x75, 0x6c, 0x6e, 0x65, 0x72, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x34, 0x0a, 0x16,
	0x70, 0x72, 0x65, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x61, 0x70, 0x61,
	0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x70, 0x72,
	0x65, 0x45, 0x6d, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x22, 0x96, 0x01, 0x0a, 0x04, 0x41, 0x6d, 0x62, 0x72, 0x12, 0x13, 0x0a, 0x05, 0x62,
	0x72, 0x5f, 0x75, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x62, 0x72, 0x55, 0x6c,
	0x12, 0x13, 0x0a, 0x05, 0x62, 0x72, 0x5f, 0x64, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x62, 0x72, 0x44, 0x6c, 0x12, 0x3d, 0x0a, 0x07, 0x62, 0x72, 0x5f, 0x75, 0x6e, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c,
	0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x41, 0x6d, 0x62, 0x72, 0x2e, 0x42, 0x69, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x73, 0x41, 0x4d, 0x42, 0x52, 0x52, 0x06, 0x62, 0x72,
	0x55, 0x6e, 0x69, 0x74, 0x22, 0x25, 0x0a, 0x10, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x55,
	0x6e, 0x69, 0x74, 0x73, 0x41, 0x4d, 0x42, 0x52, 0x12, 0x07, 0x0a, 0x03, 0x42, 0x50, 0x53, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4b, 0x42, 0x50, 0x53, 0x10, 0x01, 0x22, 0x99, 0x02, 0x0a, 0x09,
	0x41, 0x70, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x64, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x64, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4d, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x5f, 0x71, 0x6f,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e,
	0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x45, 0x70, 0x73, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x64, 0x51, 0x6f, 0x73, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52,
	0x0d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x51, 0x6f, 0x73, 0x12, 0x27,
	0x0a, 0x04, 0x61, 0x6d, 0x62, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x41, 0x6d, 0x62,
	0x72, 0x52, 0x04, 0x61, 0x6d, 0x62, 0x72, 0x22, 0xa5, 0x01, 0x0a, 0x10, 0x41, 0x70, 0x6e, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x2d, 0x0a, 0x12,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x11, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78,
	0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x10, 0x61,
	0x6c, 0x6c, 0x5f, 0x61, 0x70, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x5f, 0x69, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x61, 0x6c, 0x6c, 0x41, 0x70, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x49, 0x6e, 0x64, 0x12, 0x39, 0x0a, 0x0b, 0x61, 0x70, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x41, 0x70, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0a, 0x61, 0x70, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x42,
	0x1f, 0x5a, 0x1d, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6c, 0x74, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x61, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lte_protos_oai_common_types_proto_rawDescOnce sync.Once
	file_lte_protos_oai_common_types_proto_rawDescData = file_lte_protos_oai_common_types_proto_rawDesc
)

func file_lte_protos_oai_common_types_proto_rawDescGZIP() []byte {
	file_lte_protos_oai_common_types_proto_rawDescOnce.Do(func() {
		file_lte_protos_oai_common_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_lte_protos_oai_common_types_proto_rawDescData)
	})
	return file_lte_protos_oai_common_types_proto_rawDescData
}

var file_lte_protos_oai_common_types_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_lte_protos_oai_common_types_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_lte_protos_oai_common_types_proto_goTypes = []interface{}{
	(Ambr_BitrateUnitsAMBR)(0),      // 0: magma.lte.oai.Ambr.BitrateUnitsAMBR
	(*Guti)(nil),                    // 1: magma.lte.oai.Guti
	(*Ecgi)(nil),                    // 2: magma.lte.oai.Ecgi
	(*EpsSubscribedQosProfile)(nil), // 3: magma.lte.oai.EpsSubscribedQosProfile
	(*Ambr)(nil),                    // 4: magma.lte.oai.Ambr
	(*ApnConfig)(nil),               // 5: magma.lte.oai.ApnConfig
	(*ApnConfigProfile)(nil),        // 6: magma.lte.oai.ApnConfigProfile
}
var file_lte_protos_oai_common_types_proto_depIdxs = []int32{
	0, // 0: magma.lte.oai.Ambr.br_unit:type_name -> magma.lte.oai.Ambr.BitrateUnitsAMBR
	3, // 1: magma.lte.oai.ApnConfig.subscribed_qos:type_name -> magma.lte.oai.EpsSubscribedQosProfile
	4, // 2: magma.lte.oai.ApnConfig.ambr:type_name -> magma.lte.oai.Ambr
	5, // 3: magma.lte.oai.ApnConfigProfile.apn_configs:type_name -> magma.lte.oai.ApnConfig
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_lte_protos_oai_common_types_proto_init() }
func file_lte_protos_oai_common_types_proto_init() {
	if File_lte_protos_oai_common_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lte_protos_oai_common_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Guti); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lte_protos_oai_common_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ecgi); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lte_protos_oai_common_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EpsSubscribedQosProfile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lte_protos_oai_common_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ambr); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lte_protos_oai_common_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApnConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_lte_protos_oai_common_types_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApnConfigProfile); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_lte_protos_oai_common_types_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lte_protos_oai_common_types_proto_goTypes,
		DependencyIndexes: file_lte_protos_oai_common_types_proto_depIdxs,
		EnumInfos:         file_lte_protos_oai_common_types_proto_enumTypes,
		MessageInfos:      file_lte_protos_oai_common_types_proto_msgTypes,
	}.Build()
	File_lte_protos_oai_common_types_proto = out.File
	file_lte_protos_oai_common_types_proto_rawDesc = nil
	file_lte_protos_oai_common_types_proto_goTypes = nil
	file_lte_protos_oai_common_types_proto_depIdxs = nil
}
