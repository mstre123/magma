//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: lte/protos/oai/sgw_state.proto

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

// sgw_bearer_context_information
type SgwS8EpsBearerContextInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Imsi                []byte            `protobuf:"bytes,1,opt,name=imsi,proto3" json:"imsi,omitempty"`
	ImsiUnauthIndicator uint32            `protobuf:"varint,2,opt,name=imsi_unauth_indicator,json=imsiUnauthIndicator,proto3" json:"imsi_unauth_indicator,omitempty"`
	Msisdn              string            `protobuf:"bytes,3,opt,name=msisdn,proto3" json:"msisdn,omitempty"`
	LastKnownCellId     *Ecgi             `protobuf:"bytes,4,opt,name=last_known_cell_id,json=lastKnownCellId,proto3" json:"last_known_cell_id,omitempty"`
	Trxn                []byte            `protobuf:"bytes,5,opt,name=trxn,proto3" json:"trxn,omitempty"`
	Imsi64              uint64            `protobuf:"varint,6,opt,name=imsi64,proto3" json:"imsi64,omitempty"`
	MmeTeidS11          uint32            `protobuf:"varint,7,opt,name=mme_teid_s11,json=mmeTeidS11,proto3" json:"mme_teid_s11,omitempty"`
	MmeIpAddressS11     []byte            `protobuf:"bytes,8,opt,name=mme_ip_address_s11,json=mmeIpAddressS11,proto3" json:"mme_ip_address_s11,omitempty"`
	SgwTeidS11S4        uint32            `protobuf:"varint,9,opt,name=sgw_teid_s11_s4,json=sgwTeidS11S4,proto3" json:"sgw_teid_s11_s4,omitempty"`
	SgwIpAddressS11S4   []byte            `protobuf:"bytes,10,opt,name=sgw_ip_address_s11_s4,json=sgwIpAddressS11S4,proto3" json:"sgw_ip_address_s11_s4,omitempty"`
	PdnConnection       *SgwPdnConnection `protobuf:"bytes,11,opt,name=pdn_connection,json=pdnConnection,proto3" json:"pdn_connection,omitempty"`
}

func (x *SgwS8EpsBearerContextInfo) Reset() {
	*x = SgwS8EpsBearerContextInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_sgw_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SgwS8EpsBearerContextInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SgwS8EpsBearerContextInfo) ProtoMessage() {}

func (x *SgwS8EpsBearerContextInfo) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_sgw_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SgwS8EpsBearerContextInfo.ProtoReflect.Descriptor instead.
func (*SgwS8EpsBearerContextInfo) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_sgw_state_proto_rawDescGZIP(), []int{0}
}

func (x *SgwS8EpsBearerContextInfo) GetImsi() []byte {
	if x != nil {
		return x.Imsi
	}
	return nil
}

func (x *SgwS8EpsBearerContextInfo) GetImsiUnauthIndicator() uint32 {
	if x != nil {
		return x.ImsiUnauthIndicator
	}
	return 0
}

func (x *SgwS8EpsBearerContextInfo) GetMsisdn() string {
	if x != nil {
		return x.Msisdn
	}
	return ""
}

func (x *SgwS8EpsBearerContextInfo) GetLastKnownCellId() *Ecgi {
	if x != nil {
		return x.LastKnownCellId
	}
	return nil
}

func (x *SgwS8EpsBearerContextInfo) GetTrxn() []byte {
	if x != nil {
		return x.Trxn
	}
	return nil
}

func (x *SgwS8EpsBearerContextInfo) GetImsi64() uint64 {
	if x != nil {
		return x.Imsi64
	}
	return 0
}

func (x *SgwS8EpsBearerContextInfo) GetMmeTeidS11() uint32 {
	if x != nil {
		return x.MmeTeidS11
	}
	return 0
}

func (x *SgwS8EpsBearerContextInfo) GetMmeIpAddressS11() []byte {
	if x != nil {
		return x.MmeIpAddressS11
	}
	return nil
}

func (x *SgwS8EpsBearerContextInfo) GetSgwTeidS11S4() uint32 {
	if x != nil {
		return x.SgwTeidS11S4
	}
	return 0
}

func (x *SgwS8EpsBearerContextInfo) GetSgwIpAddressS11S4() []byte {
	if x != nil {
		return x.SgwIpAddressS11S4
	}
	return nil
}

func (x *SgwS8EpsBearerContextInfo) GetPdnConnection() *SgwPdnConnection {
	if x != nil {
		return x.PdnConnection
	}
	return nil
}

type SgwUeContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SgwBearerContext []*SgwS8EpsBearerContextInfo `protobuf:"bytes,1,rep,name=sgw_bearer_context,json=sgwBearerContext,proto3" json:"sgw_bearer_context,omitempty"`
}

func (x *SgwUeContext) Reset() {
	*x = SgwUeContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_sgw_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SgwUeContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SgwUeContext) ProtoMessage() {}

func (x *SgwUeContext) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_sgw_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SgwUeContext.ProtoReflect.Descriptor instead.
func (*SgwUeContext) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_sgw_state_proto_rawDescGZIP(), []int{1}
}

func (x *SgwUeContext) GetSgwBearerContext() []*SgwS8EpsBearerContextInfo {
	if x != nil {
		return x.SgwBearerContext
	}
	return nil
}

type SgwState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastTunnelId uint32 `protobuf:"varint,1,opt,name=last_tunnel_id,json=lastTunnelId,proto3" json:"last_tunnel_id,omitempty"`
	Gtpv1UTeid   uint32 `protobuf:"varint,2,opt,name=gtpv1u_teid,json=gtpv1uTeid,proto3" json:"gtpv1u_teid,omitempty"`
}

func (x *SgwState) Reset() {
	*x = SgwState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lte_protos_oai_sgw_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SgwState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SgwState) ProtoMessage() {}

func (x *SgwState) ProtoReflect() protoreflect.Message {
	mi := &file_lte_protos_oai_sgw_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SgwState.ProtoReflect.Descriptor instead.
func (*SgwState) Descriptor() ([]byte, []int) {
	return file_lte_protos_oai_sgw_state_proto_rawDescGZIP(), []int{2}
}

func (x *SgwState) GetLastTunnelId() uint32 {
	if x != nil {
		return x.LastTunnelId
	}
	return 0
}

func (x *SgwState) GetGtpv1UTeid() uint32 {
	if x != nil {
		return x.Gtpv1UTeid
	}
	return 0
}

var File_lte_protos_oai_sgw_state_proto protoreflect.FileDescriptor

var file_lte_protos_oai_sgw_state_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6c, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x61, 0x69,
	0x2f, 0x73, 0x67, 0x77, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0d, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69, 0x1a,
	0x21, 0x6c, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x61, 0x69, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x6c, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f,
	0x61, 0x69, 0x2f, 0x73, 0x74, 0x64, 0x5f, 0x33, 0x67, 0x70, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x6c, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x61, 0x69, 0x2f, 0x73, 0x70, 0x67, 0x77, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x03, 0x0a, 0x19, 0x53, 0x67, 0x77,
	0x53, 0x38, 0x45, 0x70, 0x73, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x6d, 0x73, 0x69, 0x12, 0x32, 0x0a, 0x15, 0x69, 0x6d,
	0x73, 0x69, 0x5f, 0x75, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x69, 0x63, 0x61,
	0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x13, 0x69, 0x6d, 0x73, 0x69, 0x55,
	0x6e, 0x61, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x64, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x6d, 0x73, 0x69, 0x73, 0x64, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x6d, 0x73, 0x69, 0x73, 0x64, 0x6e, 0x12, 0x40, 0x0a, 0x12, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x63, 0x65, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f,
	0x61, 0x69, 0x2e, 0x45, 0x63, 0x67, 0x69, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x4b, 0x6e, 0x6f,
	0x77, 0x6e, 0x43, 0x65, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x72, 0x78, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x74, 0x72, 0x78, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x6d, 0x73, 0x69, 0x36, 0x34, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x69, 0x6d,
	0x73, 0x69, 0x36, 0x34, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x6d, 0x65, 0x5f, 0x74, 0x65, 0x69, 0x64,
	0x5f, 0x73, 0x31, 0x31, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x6d, 0x65, 0x54,
	0x65, 0x69, 0x64, 0x53, 0x31, 0x31, 0x12, 0x2b, 0x0a, 0x12, 0x6d, 0x6d, 0x65, 0x5f, 0x69, 0x70,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x31, 0x31, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0f, 0x6d, 0x6d, 0x65, 0x49, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x31, 0x31, 0x12, 0x25, 0x0a, 0x0f, 0x73, 0x67, 0x77, 0x5f, 0x74, 0x65, 0x69, 0x64, 0x5f,
	0x73, 0x31, 0x31, 0x5f, 0x73, 0x34, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x73, 0x67,
	0x77, 0x54, 0x65, 0x69, 0x64, 0x53, 0x31, 0x31, 0x53, 0x34, 0x12, 0x30, 0x0a, 0x15, 0x73, 0x67,
	0x77, 0x5f, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x73, 0x31, 0x31,
	0x5f, 0x73, 0x34, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x73, 0x67, 0x77, 0x49, 0x70,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x31, 0x31, 0x53, 0x34, 0x12, 0x46, 0x0a, 0x0e,
	0x70, 0x64, 0x6e, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65,
	0x2e, 0x6f, 0x61, 0x69, 0x2e, 0x53, 0x67, 0x77, 0x50, 0x64, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x70, 0x64, 0x6e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x66, 0x0a, 0x0c, 0x53, 0x67, 0x77, 0x55, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x56, 0x0a, 0x12, 0x73, 0x67, 0x77, 0x5f, 0x62, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x28, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6c, 0x74, 0x65, 0x2e, 0x6f, 0x61, 0x69,
	0x2e, 0x53, 0x67, 0x77, 0x53, 0x38, 0x45, 0x70, 0x73, 0x42, 0x65, 0x61, 0x72, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x10, 0x73, 0x67, 0x77, 0x42,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x22, 0x51, 0x0a, 0x08,
	0x53, 0x67, 0x77, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x74, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x54, 0x75, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x67, 0x74, 0x70, 0x76, 0x31, 0x75, 0x5f, 0x74, 0x65, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x67, 0x74, 0x70, 0x76, 0x31, 0x75, 0x54, 0x65, 0x69, 0x64, 0x42,
	0x1f, 0x5a, 0x1d, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2f, 0x6c, 0x74, 0x65, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x61, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lte_protos_oai_sgw_state_proto_rawDescOnce sync.Once
	file_lte_protos_oai_sgw_state_proto_rawDescData = file_lte_protos_oai_sgw_state_proto_rawDesc
)

func file_lte_protos_oai_sgw_state_proto_rawDescGZIP() []byte {
	file_lte_protos_oai_sgw_state_proto_rawDescOnce.Do(func() {
		file_lte_protos_oai_sgw_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_lte_protos_oai_sgw_state_proto_rawDescData)
	})
	return file_lte_protos_oai_sgw_state_proto_rawDescData
}

var file_lte_protos_oai_sgw_state_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_lte_protos_oai_sgw_state_proto_goTypes = []interface{}{
	(*SgwS8EpsBearerContextInfo)(nil), // 0: magma.lte.oai.SgwS8EpsBearerContextInfo
	(*SgwUeContext)(nil),              // 1: magma.lte.oai.SgwUeContext
	(*SgwState)(nil),                  // 2: magma.lte.oai.SgwState
	(*Ecgi)(nil),                      // 3: magma.lte.oai.Ecgi
	(*SgwPdnConnection)(nil),          // 4: magma.lte.oai.SgwPdnConnection
}
var file_lte_protos_oai_sgw_state_proto_depIdxs = []int32{
	3, // 0: magma.lte.oai.SgwS8EpsBearerContextInfo.last_known_cell_id:type_name -> magma.lte.oai.Ecgi
	4, // 1: magma.lte.oai.SgwS8EpsBearerContextInfo.pdn_connection:type_name -> magma.lte.oai.SgwPdnConnection
	0, // 2: magma.lte.oai.SgwUeContext.sgw_bearer_context:type_name -> magma.lte.oai.SgwS8EpsBearerContextInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_lte_protos_oai_sgw_state_proto_init() }
func file_lte_protos_oai_sgw_state_proto_init() {
	if File_lte_protos_oai_sgw_state_proto != nil {
		return
	}
	file_lte_protos_oai_common_types_proto_init()
	file_lte_protos_oai_std_3gpp_types_proto_init()
	file_lte_protos_oai_spgw_state_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_lte_protos_oai_sgw_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SgwS8EpsBearerContextInfo); i {
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
		file_lte_protos_oai_sgw_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SgwUeContext); i {
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
		file_lte_protos_oai_sgw_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SgwState); i {
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
			RawDescriptor: file_lte_protos_oai_sgw_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lte_protos_oai_sgw_state_proto_goTypes,
		DependencyIndexes: file_lte_protos_oai_sgw_state_proto_depIdxs,
		MessageInfos:      file_lte_protos_oai_sgw_state_proto_msgTypes,
	}.Build()
	File_lte_protos_oai_sgw_state_proto = out.File
	file_lte_protos_oai_sgw_state_proto_rawDesc = nil
	file_lte_protos_oai_sgw_state_proto_goTypes = nil
	file_lte_protos_oai_sgw_state_proto_depIdxs = nil
}
