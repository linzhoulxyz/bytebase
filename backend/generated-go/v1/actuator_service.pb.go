// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: v1/actuator_service.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// The request message for getting the theme resource.
type GetResourcePackageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResourcePackageRequest) Reset() {
	*x = GetResourcePackageRequest{}
	mi := &file_v1_actuator_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResourcePackageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResourcePackageRequest) ProtoMessage() {}

func (x *GetResourcePackageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResourcePackageRequest.ProtoReflect.Descriptor instead.
func (*GetResourcePackageRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{0}
}

// The theme resources.
type ResourcePackage struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The branding logo.
	Logo          []byte `protobuf:"bytes,1,opt,name=logo,proto3" json:"logo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourcePackage) Reset() {
	*x = ResourcePackage{}
	mi := &file_v1_actuator_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourcePackage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourcePackage) ProtoMessage() {}

func (x *ResourcePackage) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourcePackage.ProtoReflect.Descriptor instead.
func (*ResourcePackage) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{1}
}

func (x *ResourcePackage) GetLogo() []byte {
	if x != nil {
		return x.Logo
	}
	return nil
}

type SetupSampleRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SetupSampleRequest) Reset() {
	*x = SetupSampleRequest{}
	mi := &file_v1_actuator_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SetupSampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupSampleRequest) ProtoMessage() {}

func (x *SetupSampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupSampleRequest.ProtoReflect.Descriptor instead.
func (*SetupSampleRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{2}
}

type GetActuatorInfoRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetActuatorInfoRequest) Reset() {
	*x = GetActuatorInfoRequest{}
	mi := &file_v1_actuator_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetActuatorInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetActuatorInfoRequest) ProtoMessage() {}

func (x *GetActuatorInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetActuatorInfoRequest.ProtoReflect.Descriptor instead.
func (*GetActuatorInfoRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{3}
}

type UpdateActuatorInfoRequest struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The actuator to update.
	Actuator *ActuatorInfo `protobuf:"bytes,1,opt,name=actuator,proto3" json:"actuator,omitempty"`
	// The list of fields to update.
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateActuatorInfoRequest) Reset() {
	*x = UpdateActuatorInfoRequest{}
	mi := &file_v1_actuator_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateActuatorInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateActuatorInfoRequest) ProtoMessage() {}

func (x *UpdateActuatorInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateActuatorInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateActuatorInfoRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateActuatorInfoRequest) GetActuator() *ActuatorInfo {
	if x != nil {
		return x.Actuator
	}
	return nil
}

func (x *UpdateActuatorInfoRequest) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

type DeleteCacheRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCacheRequest) Reset() {
	*x = DeleteCacheRequest{}
	mi := &file_v1_actuator_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCacheRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCacheRequest) ProtoMessage() {}

func (x *DeleteCacheRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCacheRequest.ProtoReflect.Descriptor instead.
func (*DeleteCacheRequest) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{5}
}

// ServerInfo is the API message for server info.
// Actuator concept is similar to the Spring Boot Actuator.
type ActuatorInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// version is the bytebase's server version
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// git_commit is the git commit hash of the build
	GitCommit string `protobuf:"bytes,2,opt,name=git_commit,json=gitCommit,proto3" json:"git_commit,omitempty"`
	// readonly flag means if the Bytebase is running in readonly mode.
	Readonly bool `protobuf:"varint,3,opt,name=readonly,proto3" json:"readonly,omitempty"`
	// saas flag means if the Bytebase is running in SaaS mode, some features are not allowed to edit by users.
	Saas bool `protobuf:"varint,4,opt,name=saas,proto3" json:"saas,omitempty"`
	// demo flag means if the Bytebase is running in demo mode.
	Demo bool `protobuf:"varint,5,opt,name=demo,proto3" json:"demo,omitempty"`
	// host is the Bytebase instance host.
	Host string `protobuf:"bytes,6,opt,name=host,proto3" json:"host,omitempty"`
	// port is the Bytebase instance port.
	Port string `protobuf:"bytes,7,opt,name=port,proto3" json:"port,omitempty"`
	// external_url is the URL where user or webhook callback visits Bytebase.
	ExternalUrl string `protobuf:"bytes,8,opt,name=external_url,json=externalUrl,proto3" json:"external_url,omitempty"`
	// need_admin_setup flag means the Bytebase instance doesn't have any end users.
	NeedAdminSetup bool `protobuf:"varint,9,opt,name=need_admin_setup,json=needAdminSetup,proto3" json:"need_admin_setup,omitempty"`
	// disallow_signup is the flag to disable self-service signup.
	DisallowSignup bool `protobuf:"varint,10,opt,name=disallow_signup,json=disallowSignup,proto3" json:"disallow_signup,omitempty"`
	// last_active_time is the service last active time in UTC Time Format, any API calls will refresh this value.
	LastActiveTime *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=last_active_time,json=lastActiveTime,proto3" json:"last_active_time,omitempty"`
	// require_2fa is the flag to require 2FA for all users.
	Require_2Fa bool `protobuf:"varint,12,opt,name=require_2fa,json=require2fa,proto3" json:"require_2fa,omitempty"`
	// workspace_id is the identifier for the workspace.
	WorkspaceId string `protobuf:"bytes,13,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	// debug flag means if the debug mode is enabled.
	Debug              bool     `protobuf:"varint,15,opt,name=debug,proto3" json:"debug,omitempty"`
	UnlicensedFeatures []string `protobuf:"bytes,19,rep,name=unlicensed_features,json=unlicensedFeatures,proto3" json:"unlicensed_features,omitempty"`
	// disallow_password_signin is the flag to disallow user signin with email&password. (except workspace admins)
	DisallowPasswordSignin bool                        `protobuf:"varint,20,opt,name=disallow_password_signin,json=disallowPasswordSignin,proto3" json:"disallow_password_signin,omitempty"`
	PasswordRestriction    *PasswordRestrictionSetting `protobuf:"bytes,21,opt,name=password_restriction,json=passwordRestriction,proto3" json:"password_restriction,omitempty"`
	// docker flag means if the Bytebase instance is running in docker.
	Docker                 bool                     `protobuf:"varint,22,opt,name=docker,proto3" json:"docker,omitempty"`
	UserStats              []*ActuatorInfo_StatUser `protobuf:"bytes,23,rep,name=user_stats,json=userStats,proto3" json:"user_stats,omitempty"`
	ActivatedInstanceCount int32                    `protobuf:"varint,24,opt,name=activated_instance_count,json=activatedInstanceCount,proto3" json:"activated_instance_count,omitempty"`
	TotalInstanceCount     int32                    `protobuf:"varint,25,opt,name=total_instance_count,json=totalInstanceCount,proto3" json:"total_instance_count,omitempty"`
	EnableSample           bool                     `protobuf:"varint,26,opt,name=enable_sample,json=enableSample,proto3" json:"enable_sample,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *ActuatorInfo) Reset() {
	*x = ActuatorInfo{}
	mi := &file_v1_actuator_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActuatorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActuatorInfo) ProtoMessage() {}

func (x *ActuatorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActuatorInfo.ProtoReflect.Descriptor instead.
func (*ActuatorInfo) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{6}
}

func (x *ActuatorInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ActuatorInfo) GetGitCommit() string {
	if x != nil {
		return x.GitCommit
	}
	return ""
}

func (x *ActuatorInfo) GetReadonly() bool {
	if x != nil {
		return x.Readonly
	}
	return false
}

func (x *ActuatorInfo) GetSaas() bool {
	if x != nil {
		return x.Saas
	}
	return false
}

func (x *ActuatorInfo) GetDemo() bool {
	if x != nil {
		return x.Demo
	}
	return false
}

func (x *ActuatorInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ActuatorInfo) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ActuatorInfo) GetExternalUrl() string {
	if x != nil {
		return x.ExternalUrl
	}
	return ""
}

func (x *ActuatorInfo) GetNeedAdminSetup() bool {
	if x != nil {
		return x.NeedAdminSetup
	}
	return false
}

func (x *ActuatorInfo) GetDisallowSignup() bool {
	if x != nil {
		return x.DisallowSignup
	}
	return false
}

func (x *ActuatorInfo) GetLastActiveTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LastActiveTime
	}
	return nil
}

func (x *ActuatorInfo) GetRequire_2Fa() bool {
	if x != nil {
		return x.Require_2Fa
	}
	return false
}

func (x *ActuatorInfo) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *ActuatorInfo) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *ActuatorInfo) GetUnlicensedFeatures() []string {
	if x != nil {
		return x.UnlicensedFeatures
	}
	return nil
}

func (x *ActuatorInfo) GetDisallowPasswordSignin() bool {
	if x != nil {
		return x.DisallowPasswordSignin
	}
	return false
}

func (x *ActuatorInfo) GetPasswordRestriction() *PasswordRestrictionSetting {
	if x != nil {
		return x.PasswordRestriction
	}
	return nil
}

func (x *ActuatorInfo) GetDocker() bool {
	if x != nil {
		return x.Docker
	}
	return false
}

func (x *ActuatorInfo) GetUserStats() []*ActuatorInfo_StatUser {
	if x != nil {
		return x.UserStats
	}
	return nil
}

func (x *ActuatorInfo) GetActivatedInstanceCount() int32 {
	if x != nil {
		return x.ActivatedInstanceCount
	}
	return 0
}

func (x *ActuatorInfo) GetTotalInstanceCount() int32 {
	if x != nil {
		return x.TotalInstanceCount
	}
	return 0
}

func (x *ActuatorInfo) GetEnableSample() bool {
	if x != nil {
		return x.EnableSample
	}
	return false
}

type ActuatorInfo_StatUser struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserType      UserType               `protobuf:"varint,1,opt,name=user_type,json=userType,proto3,enum=bytebase.v1.UserType" json:"user_type,omitempty"`
	State         State                  `protobuf:"varint,2,opt,name=state,proto3,enum=bytebase.v1.State" json:"state,omitempty"`
	Count         int32                  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ActuatorInfo_StatUser) Reset() {
	*x = ActuatorInfo_StatUser{}
	mi := &file_v1_actuator_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ActuatorInfo_StatUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActuatorInfo_StatUser) ProtoMessage() {}

func (x *ActuatorInfo_StatUser) ProtoReflect() protoreflect.Message {
	mi := &file_v1_actuator_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActuatorInfo_StatUser.ProtoReflect.Descriptor instead.
func (*ActuatorInfo_StatUser) Descriptor() ([]byte, []int) {
	return file_v1_actuator_service_proto_rawDescGZIP(), []int{6, 0}
}

func (x *ActuatorInfo_StatUser) GetUserType() UserType {
	if x != nil {
		return x.UserType
	}
	return UserType_USER_TYPE_UNSPECIFIED
}

func (x *ActuatorInfo_StatUser) GetState() State {
	if x != nil {
		return x.State
	}
	return State_STATE_UNSPECIFIED
}

func (x *ActuatorInfo_StatUser) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_v1_actuator_service_proto protoreflect.FileDescriptor

const file_v1_actuator_service_proto_rawDesc = "" +
	"\n" +
	"\x19v1/actuator_service.proto\x12\vbytebase.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x17google/api/client.proto\x1a\x1fgoogle/api/field_behavior.proto\x1a\x1bgoogle/protobuf/empty.proto\x1a google/protobuf/field_mask.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x13v1/annotation.proto\x1a\x0fv1/common.proto\x1a\x18v1/setting_service.proto\x1a\x15v1/user_service.proto\"\x1b\n" +
	"\x19GetResourcePackageRequest\"%\n" +
	"\x0fResourcePackage\x12\x12\n" +
	"\x04logo\x18\x01 \x01(\fR\x04logo\"\x14\n" +
	"\x12SetupSampleRequest\"\x18\n" +
	"\x16GetActuatorInfoRequest\"\x99\x01\n" +
	"\x19UpdateActuatorInfoRequest\x12:\n" +
	"\bactuator\x18\x01 \x01(\v2\x19.bytebase.v1.ActuatorInfoB\x03\xe0A\x02R\bactuator\x12@\n" +
	"\vupdate_mask\x18\x02 \x01(\v2\x1a.google.protobuf.FieldMaskB\x03\xe0A\x02R\n" +
	"updateMask\"\x14\n" +
	"\x12DeleteCacheRequest\"\xe5\b\n" +
	"\fActuatorInfo\x12\x1d\n" +
	"\aversion\x18\x01 \x01(\tB\x03\xe0A\x03R\aversion\x12\"\n" +
	"\n" +
	"git_commit\x18\x02 \x01(\tB\x03\xe0A\x03R\tgitCommit\x12\x1f\n" +
	"\breadonly\x18\x03 \x01(\bB\x03\xe0A\x03R\breadonly\x12\x17\n" +
	"\x04saas\x18\x04 \x01(\bB\x03\xe0A\x03R\x04saas\x12\x17\n" +
	"\x04demo\x18\x05 \x01(\bB\x03\xe0A\x03R\x04demo\x12\x17\n" +
	"\x04host\x18\x06 \x01(\tB\x03\xe0A\x03R\x04host\x12\x17\n" +
	"\x04port\x18\a \x01(\tB\x03\xe0A\x03R\x04port\x12&\n" +
	"\fexternal_url\x18\b \x01(\tB\x03\xe0A\x03R\vexternalUrl\x12-\n" +
	"\x10need_admin_setup\x18\t \x01(\bB\x03\xe0A\x03R\x0eneedAdminSetup\x12,\n" +
	"\x0fdisallow_signup\x18\n" +
	" \x01(\bB\x03\xe0A\x03R\x0edisallowSignup\x12I\n" +
	"\x10last_active_time\x18\v \x01(\v2\x1a.google.protobuf.TimestampB\x03\xe0A\x03R\x0elastActiveTime\x12$\n" +
	"\vrequire_2fa\x18\f \x01(\bB\x03\xe0A\x03R\n" +
	"require2fa\x12&\n" +
	"\fworkspace_id\x18\r \x01(\tB\x03\xe0A\x03R\vworkspaceId\x12\x14\n" +
	"\x05debug\x18\x0f \x01(\bR\x05debug\x124\n" +
	"\x13unlicensed_features\x18\x13 \x03(\tB\x03\xe0A\x03R\x12unlicensedFeatures\x12=\n" +
	"\x18disallow_password_signin\x18\x14 \x01(\bB\x03\xe0A\x03R\x16disallowPasswordSignin\x12_\n" +
	"\x14password_restriction\x18\x15 \x01(\v2'.bytebase.v1.PasswordRestrictionSettingB\x03\xe0A\x03R\x13passwordRestriction\x12\x1b\n" +
	"\x06docker\x18\x16 \x01(\bB\x03\xe0A\x03R\x06docker\x12F\n" +
	"\n" +
	"user_stats\x18\x17 \x03(\v2\".bytebase.v1.ActuatorInfo.StatUserB\x03\xe0A\x03R\tuserStats\x12=\n" +
	"\x18activated_instance_count\x18\x18 \x01(\x05B\x03\xe0A\x03R\x16activatedInstanceCount\x125\n" +
	"\x14total_instance_count\x18\x19 \x01(\x05B\x03\xe0A\x03R\x12totalInstanceCount\x12(\n" +
	"\renable_sample\x18\x1a \x01(\bB\x03\xe0A\x03R\fenableSample\x1a~\n" +
	"\bStatUser\x122\n" +
	"\tuser_type\x18\x01 \x01(\x0e2\x15.bytebase.v1.UserTypeR\buserType\x12(\n" +
	"\x05state\x18\x02 \x01(\x0e2\x12.bytebase.v1.StateR\x05state\x12\x14\n" +
	"\x05count\x18\x03 \x01(\x05R\x05count2\xa4\x05\n" +
	"\x0fActuatorService\x12s\n" +
	"\x0fGetActuatorInfo\x12#.bytebase.v1.GetActuatorInfoRequest\x1a\x19.bytebase.v1.ActuatorInfo\" \xdaA\x00\x80\xea0\x01\x82\xd3\xe4\x93\x02\x13\x12\x11/v1/actuator/info\x12\xaa\x01\n" +
	"\x12UpdateActuatorInfo\x12&.bytebase.v1.UpdateActuatorInfoRequest\x1a\x19.bytebase.v1.ActuatorInfo\"Q\xdaA\x14actuator,update_mask\x8a\xea0\x0fbb.settings.set\x90\xea0\x01\x82\xd3\xe4\x93\x02\x1d:\bactuator2\x11/v1/actuator/info\x12\x82\x01\n" +
	"\vSetupSample\x12\x1f.bytebase.v1.SetupSampleRequest\x1a\x16.google.protobuf.Empty\":\x8a\xea0\x12bb.projects.create\x90\xea0\x01\x82\xd3\xe4\x93\x02\x1a\"\x18/v1/actuator:setupSample\x12f\n" +
	"\vDeleteCache\x12\x1f.bytebase.v1.DeleteCacheRequest\x1a\x16.google.protobuf.Empty\"\x1e\x80\xea0\x01\x82\xd3\xe4\x93\x02\x14*\x12/v1/actuator/cache\x12\x81\x01\n" +
	"\x12GetResourcePackage\x12&.bytebase.v1.GetResourcePackageRequest\x1a\x1c.bytebase.v1.ResourcePackage\"%\xdaA\x00\x80\xea0\x01\x82\xd3\xe4\x93\x02\x18\x12\x16/v1/actuator/resourcesB6Z4github.com/bytebase/bytebase/backend/generated-go/v1b\x06proto3"

var (
	file_v1_actuator_service_proto_rawDescOnce sync.Once
	file_v1_actuator_service_proto_rawDescData []byte
)

func file_v1_actuator_service_proto_rawDescGZIP() []byte {
	file_v1_actuator_service_proto_rawDescOnce.Do(func() {
		file_v1_actuator_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_v1_actuator_service_proto_rawDesc), len(file_v1_actuator_service_proto_rawDesc)))
	})
	return file_v1_actuator_service_proto_rawDescData
}

var file_v1_actuator_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_v1_actuator_service_proto_goTypes = []any{
	(*GetResourcePackageRequest)(nil),  // 0: bytebase.v1.GetResourcePackageRequest
	(*ResourcePackage)(nil),            // 1: bytebase.v1.ResourcePackage
	(*SetupSampleRequest)(nil),         // 2: bytebase.v1.SetupSampleRequest
	(*GetActuatorInfoRequest)(nil),     // 3: bytebase.v1.GetActuatorInfoRequest
	(*UpdateActuatorInfoRequest)(nil),  // 4: bytebase.v1.UpdateActuatorInfoRequest
	(*DeleteCacheRequest)(nil),         // 5: bytebase.v1.DeleteCacheRequest
	(*ActuatorInfo)(nil),               // 6: bytebase.v1.ActuatorInfo
	(*ActuatorInfo_StatUser)(nil),      // 7: bytebase.v1.ActuatorInfo.StatUser
	(*fieldmaskpb.FieldMask)(nil),      // 8: google.protobuf.FieldMask
	(*timestamppb.Timestamp)(nil),      // 9: google.protobuf.Timestamp
	(*PasswordRestrictionSetting)(nil), // 10: bytebase.v1.PasswordRestrictionSetting
	(UserType)(0),                      // 11: bytebase.v1.UserType
	(State)(0),                         // 12: bytebase.v1.State
	(*emptypb.Empty)(nil),              // 13: google.protobuf.Empty
}
var file_v1_actuator_service_proto_depIdxs = []int32{
	6,  // 0: bytebase.v1.UpdateActuatorInfoRequest.actuator:type_name -> bytebase.v1.ActuatorInfo
	8,  // 1: bytebase.v1.UpdateActuatorInfoRequest.update_mask:type_name -> google.protobuf.FieldMask
	9,  // 2: bytebase.v1.ActuatorInfo.last_active_time:type_name -> google.protobuf.Timestamp
	10, // 3: bytebase.v1.ActuatorInfo.password_restriction:type_name -> bytebase.v1.PasswordRestrictionSetting
	7,  // 4: bytebase.v1.ActuatorInfo.user_stats:type_name -> bytebase.v1.ActuatorInfo.StatUser
	11, // 5: bytebase.v1.ActuatorInfo.StatUser.user_type:type_name -> bytebase.v1.UserType
	12, // 6: bytebase.v1.ActuatorInfo.StatUser.state:type_name -> bytebase.v1.State
	3,  // 7: bytebase.v1.ActuatorService.GetActuatorInfo:input_type -> bytebase.v1.GetActuatorInfoRequest
	4,  // 8: bytebase.v1.ActuatorService.UpdateActuatorInfo:input_type -> bytebase.v1.UpdateActuatorInfoRequest
	2,  // 9: bytebase.v1.ActuatorService.SetupSample:input_type -> bytebase.v1.SetupSampleRequest
	5,  // 10: bytebase.v1.ActuatorService.DeleteCache:input_type -> bytebase.v1.DeleteCacheRequest
	0,  // 11: bytebase.v1.ActuatorService.GetResourcePackage:input_type -> bytebase.v1.GetResourcePackageRequest
	6,  // 12: bytebase.v1.ActuatorService.GetActuatorInfo:output_type -> bytebase.v1.ActuatorInfo
	6,  // 13: bytebase.v1.ActuatorService.UpdateActuatorInfo:output_type -> bytebase.v1.ActuatorInfo
	13, // 14: bytebase.v1.ActuatorService.SetupSample:output_type -> google.protobuf.Empty
	13, // 15: bytebase.v1.ActuatorService.DeleteCache:output_type -> google.protobuf.Empty
	1,  // 16: bytebase.v1.ActuatorService.GetResourcePackage:output_type -> bytebase.v1.ResourcePackage
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_v1_actuator_service_proto_init() }
func file_v1_actuator_service_proto_init() {
	if File_v1_actuator_service_proto != nil {
		return
	}
	file_v1_annotation_proto_init()
	file_v1_common_proto_init()
	file_v1_setting_service_proto_init()
	file_v1_user_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_v1_actuator_service_proto_rawDesc), len(file_v1_actuator_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_actuator_service_proto_goTypes,
		DependencyIndexes: file_v1_actuator_service_proto_depIdxs,
		MessageInfos:      file_v1_actuator_service_proto_msgTypes,
	}.Build()
	File_v1_actuator_service_proto = out.File
	file_v1_actuator_service_proto_goTypes = nil
	file_v1_actuator_service_proto_depIdxs = nil
}
