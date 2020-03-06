// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spanner.proto

package spannerpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// State of the Tree.
// Mirrors trillian.TreeState.
type TreeState int32

const (
	TreeState_UNKNOWN_TREE_STATE TreeState = 0
	TreeState_ACTIVE             TreeState = 1
	TreeState_FROZEN             TreeState = 2
)

var TreeState_name = map[int32]string{
	0: "UNKNOWN_TREE_STATE",
	1: "ACTIVE",
	2: "FROZEN",
}

var TreeState_value = map[string]int32{
	"UNKNOWN_TREE_STATE": 0,
	"ACTIVE":             1,
	"FROZEN":             2,
}

func (x TreeState) String() string {
	return proto.EnumName(TreeState_name, int32(x))
}

func (TreeState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_879d3e919e93c6ba, []int{0}
}

// Type of the Tree.
// Mirrors trillian.TreeType.
type TreeType int32

const (
	TreeType_UNKNOWN TreeType = 0
	TreeType_LOG     TreeType = 1
	TreeType_MAP     TreeType = 2
)

var TreeType_name = map[int32]string{
	0: "UNKNOWN",
	1: "LOG",
	2: "MAP",
}

var TreeType_value = map[string]int32{
	"UNKNOWN": 0,
	"LOG":     1,
	"MAP":     2,
}

func (x TreeType) String() string {
	return proto.EnumName(TreeType_name, int32(x))
}

func (TreeType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_879d3e919e93c6ba, []int{1}
}

// Defines the preimage protection used for tree leaves / nodes.
// Eg, RFC6962 dictates a 0x00 prefix for leaves and 0x01 for nodes.
// Mirrors trillian.HashStrategy.
type HashStrategy int32

const (
	HashStrategy_UNKNOWN_HASH_STRATEGY HashStrategy = 0
	HashStrategy_RFC_6962              HashStrategy = 1
	HashStrategy_TEST_MAP_HASHER       HashStrategy = 2
	HashStrategy_OBJECT_RFC6962_SHA256 HashStrategy = 3
	HashStrategy_CONIKS_SHA512_256     HashStrategy = 4
	HashStrategy_CONIKS_SHA256         HashStrategy = 5
)

var HashStrategy_name = map[int32]string{
	0: "UNKNOWN_HASH_STRATEGY",
	1: "RFC_6962",
	2: "TEST_MAP_HASHER",
	3: "OBJECT_RFC6962_SHA256",
	4: "CONIKS_SHA512_256",
	5: "CONIKS_SHA256",
}

var HashStrategy_value = map[string]int32{
	"UNKNOWN_HASH_STRATEGY": 0,
	"RFC_6962":              1,
	"TEST_MAP_HASHER":       2,
	"OBJECT_RFC6962_SHA256": 3,
	"CONIKS_SHA512_256":     4,
	"CONIKS_SHA256":         5,
}

func (x HashStrategy) String() string {
	return proto.EnumName(HashStrategy_name, int32(x))
}

func (HashStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_879d3e919e93c6ba, []int{2}
}

// Supported hash algorithms.
// The numbering space is the same as for TLS, given in RFC 5246 s7.4.1.4.1. See
// http://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml#tls-parameters-18.
// Mirrors trillian.HashAlgorithm.
type HashAlgorithm int32

const (
	// No hash algorithm is used.
	HashAlgorithm_NONE HashAlgorithm = 0
	// SHA256 is used.
	HashAlgorithm_SHA256 HashAlgorithm = 4
)

var HashAlgorithm_name = map[int32]string{
	0: "NONE",
	4: "SHA256",
}

var HashAlgorithm_value = map[string]int32{
	"NONE":   0,
	"SHA256": 4,
}

func (x HashAlgorithm) String() string {
	return proto.EnumName(HashAlgorithm_name, int32(x))
}

func (HashAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_879d3e919e93c6ba, []int{3}
}

// Supported signature algorithms.
// The numbering space is the same as for TLS, given in RFC 5246 s7.4.1.4.1. See
// http://www.iana.org/assignments/tls-parameters/tls-parameters.xhtml#tls-parameters-16.
// Mirrors trillian.SignatureAlgorithm.
type SignatureAlgorithm int32

const (
	// Anonymous signature scheme.
	SignatureAlgorithm_ANONYMOUS SignatureAlgorithm = 0
	// RSA signature scheme.
	SignatureAlgorithm_RSA SignatureAlgorithm = 1
	// ECDSA signature scheme.
	SignatureAlgorithm_ECDSA SignatureAlgorithm = 3
)

var SignatureAlgorithm_name = map[int32]string{
	0: "ANONYMOUS",
	1: "RSA",
	3: "ECDSA",
}

var SignatureAlgorithm_value = map[string]int32{
	"ANONYMOUS": 0,
	"RSA":       1,
	"ECDSA":     3,
}

func (x SignatureAlgorithm) String() string {
	return proto.EnumName(SignatureAlgorithm_name, int32(x))
}

func (SignatureAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_879d3e919e93c6ba, []int{4}
}

// LogStorageConfig holds settings which tune the storage implementation for
// a given log tree.
type LogStorageConfig struct {
	// num_unseq_buckets defines the length of the unsequenced time ring buffer.
	// This value must *never* be reduced for any provisioned tree.
	//
	// This value should be >= 1, and there's probably not much benefit in
	// raising it past about 4.
	// TODO(al): test what the effects of various values are here.
	NumUnseqBuckets int64 `protobuf:"varint,1,opt,name=num_unseq_buckets,json=numUnseqBuckets,proto3" json:"num_unseq_buckets,omitempty"`
	// num_merkle_buckets defines the number of individual buckets below each
	// unsequenced ring bucket.
	// This value may be changed at any time (so long as you understand the
	// impact it'll have on integration performace!)
	//
	// This value must lie in the range [1..256]
	NumMerkleBuckets     int64    `protobuf:"varint,2,opt,name=num_merkle_buckets,json=numMerkleBuckets,proto3" json:"num_merkle_buckets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogStorageConfig) Reset()         { *m = LogStorageConfig{} }
func (m *LogStorageConfig) String() string { return proto.CompactTextString(m) }
func (*LogStorageConfig) ProtoMessage()    {}
func (*LogStorageConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_879d3e919e93c6ba, []int{0}
}

func (m *LogStorageConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogStorageConfig.Unmarshal(m, b)
}
func (m *LogStorageConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogStorageConfig.Marshal(b, m, deterministic)
}
func (m *LogStorageConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogStorageConfig.Merge(m, src)
}
func (m *LogStorageConfig) XXX_Size() int {
	return xxx_messageInfo_LogStorageConfig.Size(m)
}
func (m *LogStorageConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_LogStorageConfig.DiscardUnknown(m)
}

var xxx_messageInfo_LogStorageConfig proto.InternalMessageInfo

func (m *LogStorageConfig) GetNumUnseqBuckets() int64 {
	if m != nil {
		return m.NumUnseqBuckets
	}
	return 0
}

func (m *LogStorageConfig) GetNumMerkleBuckets() int64 {
	if m != nil {
		return m.NumMerkleBuckets
	}
	return 0
}

// MapStorageConfig holds settings which tune the storage implementation for
// a given map tree.
type MapStorageConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MapStorageConfig) Reset()         { *m = MapStorageConfig{} }
func (m *MapStorageConfig) String() string { return proto.CompactTextString(m) }
func (*MapStorageConfig) ProtoMessage()    {}
func (*MapStorageConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_879d3e919e93c6ba, []int{1}
}

func (m *MapStorageConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapStorageConfig.Unmarshal(m, b)
}
func (m *MapStorageConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapStorageConfig.Marshal(b, m, deterministic)
}
func (m *MapStorageConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapStorageConfig.Merge(m, src)
}
func (m *MapStorageConfig) XXX_Size() int {
	return xxx_messageInfo_MapStorageConfig.Size(m)
}
func (m *MapStorageConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_MapStorageConfig.DiscardUnknown(m)
}

var xxx_messageInfo_MapStorageConfig proto.InternalMessageInfo

// TreeInfo stores information about a Trillian tree.
type TreeInfo struct {
	// tree_id is the ID of the tree, and is used as a primary key.
	TreeId int64 `protobuf:"varint,1,opt,name=tree_id,json=treeId,proto3" json:"tree_id,omitempty"`
	// key_id identifies the private key associated with this tree.
	KeyId int64 `protobuf:"varint,2,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	// name is a short name for this tree.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// description is a short free form text describing the tree.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// tree_type identifies whether this is a Log or a Map tree.
	TreeType TreeType `protobuf:"varint,5,opt,name=tree_type,json=treeType,proto3,enum=spannerpb.TreeType" json:"tree_type,omitempty"`
	// tree_state is the state of the tree.
	TreeState TreeState `protobuf:"varint,8,opt,name=tree_state,json=treeState,proto3,enum=spannerpb.TreeState" json:"tree_state,omitempty"`
	// hash_strategy is the hashing strategy used by the tree.
	HashStrategy HashStrategy `protobuf:"varint,9,opt,name=hash_strategy,json=hashStrategy,proto3,enum=spannerpb.HashStrategy" json:"hash_strategy,omitempty"`
	// hash_algorithm is the hash algorithm used by the tree.
	HashAlgorithm HashAlgorithm `protobuf:"varint,10,opt,name=hash_algorithm,json=hashAlgorithm,proto3,enum=spannerpb.HashAlgorithm" json:"hash_algorithm,omitempty"`
	// signature_algorithm is the signature algorithm used by the tree.
	SignatureAlgorithm SignatureAlgorithm `protobuf:"varint,11,opt,name=signature_algorithm,json=signatureAlgorithm,proto3,enum=spannerpb.SignatureAlgorithm" json:"signature_algorithm,omitempty"`
	// create_time_nanos is the creation timestamp of the tree, in nanos since
	// epoch.
	CreateTimeNanos int64 `protobuf:"varint,13,opt,name=create_time_nanos,json=createTimeNanos,proto3" json:"create_time_nanos,omitempty"`
	// update_time_nanos is the last update time of the tree, in nanos since
	// epoch.
	UpdateTimeNanos int64 `protobuf:"varint,14,opt,name=update_time_nanos,json=updateTimeNanos,proto3" json:"update_time_nanos,omitempty"`
	// private_key should be used to generate signatures for this tree.
	PrivateKey *any.Any `protobuf:"bytes,15,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	// public_key_der should be used to verify signatures produced by this tree.
	// It is the key in DER-encoded PKIX form.
	PublicKeyDer []byte `protobuf:"bytes,16,opt,name=public_key_der,json=publicKeyDer,proto3" json:"public_key_der,omitempty"`
	// config contains the log or map specific tree configuration.
	//
	// Types that are valid to be assigned to StorageConfig:
	//	*TreeInfo_LogStorageConfig
	//	*TreeInfo_MapStorageConfig
	StorageConfig isTreeInfo_StorageConfig `protobuf_oneof:"storage_config"`
	// max_root_duration_millis is the interval after which a new signed root is
	// produced even if there have been no submission.  If zero, this behavior is
	// disabled.
	MaxRootDurationMillis int64 `protobuf:"varint,17,opt,name=max_root_duration_millis,json=maxRootDurationMillis,proto3" json:"max_root_duration_millis,omitempty"`
	// If true the tree was soft deleted.
	Deleted bool `protobuf:"varint,18,opt,name=deleted,proto3" json:"deleted,omitempty"`
	// Time of tree deletion, if any.
	DeleteTimeNanos      int64    `protobuf:"varint,19,opt,name=delete_time_nanos,json=deleteTimeNanos,proto3" json:"delete_time_nanos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TreeInfo) Reset()         { *m = TreeInfo{} }
func (m *TreeInfo) String() string { return proto.CompactTextString(m) }
func (*TreeInfo) ProtoMessage()    {}
func (*TreeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_879d3e919e93c6ba, []int{2}
}

func (m *TreeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TreeInfo.Unmarshal(m, b)
}
func (m *TreeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TreeInfo.Marshal(b, m, deterministic)
}
func (m *TreeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreeInfo.Merge(m, src)
}
func (m *TreeInfo) XXX_Size() int {
	return xxx_messageInfo_TreeInfo.Size(m)
}
func (m *TreeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TreeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TreeInfo proto.InternalMessageInfo

func (m *TreeInfo) GetTreeId() int64 {
	if m != nil {
		return m.TreeId
	}
	return 0
}

func (m *TreeInfo) GetKeyId() int64 {
	if m != nil {
		return m.KeyId
	}
	return 0
}

func (m *TreeInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TreeInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TreeInfo) GetTreeType() TreeType {
	if m != nil {
		return m.TreeType
	}
	return TreeType_UNKNOWN
}

func (m *TreeInfo) GetTreeState() TreeState {
	if m != nil {
		return m.TreeState
	}
	return TreeState_UNKNOWN_TREE_STATE
}

func (m *TreeInfo) GetHashStrategy() HashStrategy {
	if m != nil {
		return m.HashStrategy
	}
	return HashStrategy_UNKNOWN_HASH_STRATEGY
}

func (m *TreeInfo) GetHashAlgorithm() HashAlgorithm {
	if m != nil {
		return m.HashAlgorithm
	}
	return HashAlgorithm_NONE
}

func (m *TreeInfo) GetSignatureAlgorithm() SignatureAlgorithm {
	if m != nil {
		return m.SignatureAlgorithm
	}
	return SignatureAlgorithm_ANONYMOUS
}

func (m *TreeInfo) GetCreateTimeNanos() int64 {
	if m != nil {
		return m.CreateTimeNanos
	}
	return 0
}

func (m *TreeInfo) GetUpdateTimeNanos() int64 {
	if m != nil {
		return m.UpdateTimeNanos
	}
	return 0
}

func (m *TreeInfo) GetPrivateKey() *any.Any {
	if m != nil {
		return m.PrivateKey
	}
	return nil
}

func (m *TreeInfo) GetPublicKeyDer() []byte {
	if m != nil {
		return m.PublicKeyDer
	}
	return nil
}

type isTreeInfo_StorageConfig interface {
	isTreeInfo_StorageConfig()
}

type TreeInfo_LogStorageConfig struct {
	LogStorageConfig *LogStorageConfig `protobuf:"bytes,6,opt,name=log_storage_config,json=logStorageConfig,proto3,oneof"`
}

type TreeInfo_MapStorageConfig struct {
	MapStorageConfig *MapStorageConfig `protobuf:"bytes,7,opt,name=map_storage_config,json=mapStorageConfig,proto3,oneof"`
}

func (*TreeInfo_LogStorageConfig) isTreeInfo_StorageConfig() {}

func (*TreeInfo_MapStorageConfig) isTreeInfo_StorageConfig() {}

func (m *TreeInfo) GetStorageConfig() isTreeInfo_StorageConfig {
	if m != nil {
		return m.StorageConfig
	}
	return nil
}

func (m *TreeInfo) GetLogStorageConfig() *LogStorageConfig {
	if x, ok := m.GetStorageConfig().(*TreeInfo_LogStorageConfig); ok {
		return x.LogStorageConfig
	}
	return nil
}

func (m *TreeInfo) GetMapStorageConfig() *MapStorageConfig {
	if x, ok := m.GetStorageConfig().(*TreeInfo_MapStorageConfig); ok {
		return x.MapStorageConfig
	}
	return nil
}

func (m *TreeInfo) GetMaxRootDurationMillis() int64 {
	if m != nil {
		return m.MaxRootDurationMillis
	}
	return 0
}

func (m *TreeInfo) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *TreeInfo) GetDeleteTimeNanos() int64 {
	if m != nil {
		return m.DeleteTimeNanos
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TreeInfo) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TreeInfo_LogStorageConfig)(nil),
		(*TreeInfo_MapStorageConfig)(nil),
	}
}

// TreeHead is the storage format for Trillian's commitment to a particular
// tree state.
type TreeHead struct {
	// tree_id identifies the tree this TreeHead is built from.
	TreeId int64 `protobuf:"varint,1,opt,name=tree_id,json=treeId,proto3" json:"tree_id,omitempty"`
	// ts_nanos is the nanosecond resolution timestamp at which the
	// TreeHead was created.
	TsNanos int64 `protobuf:"varint,2,opt,name=ts_nanos,json=tsNanos,proto3" json:"ts_nanos,omitempty"`
	// tree_size is the number of entries in the tree.
	TreeSize int64 `protobuf:"varint,3,opt,name=tree_size,json=treeSize,proto3" json:"tree_size,omitempty"`
	// root_hash is the root of the tree.
	RootHash []byte `protobuf:"bytes,4,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	// signature holds the raw digital signature across the serialized log_root
	// (not present) represented by the data in this TreeHead.
	Signature []byte `protobuf:"bytes,10,opt,name=signature,proto3" json:"signature,omitempty"`
	// tree_revision identifies the revision at which the TreeHead was created.
	TreeRevision         int64    `protobuf:"varint,6,opt,name=tree_revision,json=treeRevision,proto3" json:"tree_revision,omitempty"`
	Metadata             []byte   `protobuf:"bytes,9,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TreeHead) Reset()         { *m = TreeHead{} }
func (m *TreeHead) String() string { return proto.CompactTextString(m) }
func (*TreeHead) ProtoMessage()    {}
func (*TreeHead) Descriptor() ([]byte, []int) {
	return fileDescriptor_879d3e919e93c6ba, []int{3}
}

func (m *TreeHead) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TreeHead.Unmarshal(m, b)
}
func (m *TreeHead) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TreeHead.Marshal(b, m, deterministic)
}
func (m *TreeHead) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TreeHead.Merge(m, src)
}
func (m *TreeHead) XXX_Size() int {
	return xxx_messageInfo_TreeHead.Size(m)
}
func (m *TreeHead) XXX_DiscardUnknown() {
	xxx_messageInfo_TreeHead.DiscardUnknown(m)
}

var xxx_messageInfo_TreeHead proto.InternalMessageInfo

func (m *TreeHead) GetTreeId() int64 {
	if m != nil {
		return m.TreeId
	}
	return 0
}

func (m *TreeHead) GetTsNanos() int64 {
	if m != nil {
		return m.TsNanos
	}
	return 0
}

func (m *TreeHead) GetTreeSize() int64 {
	if m != nil {
		return m.TreeSize
	}
	return 0
}

func (m *TreeHead) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func (m *TreeHead) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *TreeHead) GetTreeRevision() int64 {
	if m != nil {
		return m.TreeRevision
	}
	return 0
}

func (m *TreeHead) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func init() {
	proto.RegisterEnum("spannerpb.TreeState", TreeState_name, TreeState_value)
	proto.RegisterEnum("spannerpb.TreeType", TreeType_name, TreeType_value)
	proto.RegisterEnum("spannerpb.HashStrategy", HashStrategy_name, HashStrategy_value)
	proto.RegisterEnum("spannerpb.HashAlgorithm", HashAlgorithm_name, HashAlgorithm_value)
	proto.RegisterEnum("spannerpb.SignatureAlgorithm", SignatureAlgorithm_name, SignatureAlgorithm_value)
	proto.RegisterType((*LogStorageConfig)(nil), "spannerpb.LogStorageConfig")
	proto.RegisterType((*MapStorageConfig)(nil), "spannerpb.MapStorageConfig")
	proto.RegisterType((*TreeInfo)(nil), "spannerpb.TreeInfo")
	proto.RegisterType((*TreeHead)(nil), "spannerpb.TreeHead")
}

func init() {
	proto.RegisterFile("spanner.proto", fileDescriptor_879d3e919e93c6ba)
}

var fileDescriptor_879d3e919e93c6ba = []byte{
	// 958 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xed, 0x6e, 0x1a, 0x47,
	0x14, 0xf5, 0x1a, 0x0c, 0xcb, 0x35, 0xd8, 0xe3, 0x71, 0xdc, 0xac, 0x93, 0x56, 0x42, 0x6e, 0x2b,
	0x51, 0x54, 0x41, 0xeb, 0xc8, 0x8e, 0xa2, 0x54, 0xaa, 0xd6, 0x18, 0x07, 0x9b, 0xb0, 0x44, 0xb3,
	0xeb, 0x56, 0xc9, 0x9f, 0xd5, 0xc0, 0x8e, 0x61, 0xe5, 0xfd, 0xea, 0xee, 0x6c, 0x14, 0xf2, 0x0c,
	0xfd, 0xd3, 0xc7, 0xec, 0x5b, 0x54, 0x33, 0xb3, 0x60, 0x8c, 0xd5, 0x7f, 0x33, 0xe7, 0x9c, 0x7b,
	0xaf, 0xe7, 0xfa, 0x9c, 0x05, 0x1a, 0x59, 0x42, 0xa3, 0x88, 0xa5, 0x9d, 0x24, 0x8d, 0x79, 0x8c,
	0x6b, 0xc5, 0x35, 0x99, 0xbc, 0x38, 0x9e, 0xc5, 0xf1, 0x2c, 0x60, 0x5d, 0x49, 0x4c, 0xf2, 0xbb,
	0x2e, 0x8d, 0x16, 0x4a, 0x75, 0x12, 0x00, 0x7a, 0x1f, 0xcf, 0x6c, 0x1e, 0xa7, 0x74, 0xc6, 0x7a,
	0x71, 0x74, 0xe7, 0xcf, 0x70, 0x1b, 0x0e, 0xa2, 0x3c, 0x74, 0xf3, 0x28, 0x63, 0x7f, 0xb9, 0x93,
	0x7c, 0x7a, 0xcf, 0x78, 0x66, 0x68, 0x4d, 0xad, 0x55, 0x22, 0xfb, 0x51, 0x1e, 0xde, 0x0a, 0xfc,
	0x42, 0xc1, 0xf8, 0x67, 0xc0, 0x42, 0x1b, 0xb2, 0xf4, 0x3e, 0x60, 0x2b, 0xf1, 0xb6, 0x14, 0xa3,
	0x28, 0x0f, 0x47, 0x92, 0x28, 0xd4, 0x27, 0x18, 0xd0, 0x88, 0x26, 0x8f, 0xa6, 0x9d, 0xfc, 0x5d,
	0x05, 0xdd, 0x49, 0x19, 0xbb, 0x8e, 0xee, 0x62, 0xfc, 0x1c, 0xaa, 0x3c, 0x65, 0xcc, 0xf5, 0xbd,
	0x62, 0x60, 0x45, 0x5c, 0xaf, 0x3d, 0x7c, 0x04, 0x95, 0x7b, 0xb6, 0x10, 0xb8, 0xea, 0xbd, 0x73,
	0xcf, 0x16, 0xd7, 0x1e, 0xc6, 0x50, 0x8e, 0x68, 0xc8, 0x8c, 0x52, 0x53, 0x6b, 0xd5, 0x88, 0x3c,
	0xe3, 0x26, 0xec, 0x7a, 0x2c, 0x9b, 0xa6, 0x7e, 0xc2, 0xfd, 0x38, 0x32, 0xca, 0x92, 0x5a, 0x87,
	0xf0, 0x2f, 0x50, 0x93, 0x53, 0xf8, 0x22, 0x61, 0xc6, 0x4e, 0x53, 0x6b, 0xed, 0x9d, 0x1e, 0x76,
	0x56, 0xeb, 0xea, 0x88, 0xbf, 0xc6, 0x59, 0x24, 0x8c, 0xe8, 0xbc, 0x38, 0xe1, 0x57, 0x00, 0xb2,
	0x22, 0xe3, 0x94, 0x33, 0x43, 0x97, 0x25, 0xcf, 0x36, 0x4a, 0x6c, 0xc1, 0x11, 0xd9, 0x59, 0x1e,
	0xf1, 0x6f, 0xd0, 0x98, 0xd3, 0x6c, 0xee, 0x66, 0x3c, 0xa5, 0x9c, 0xcd, 0x16, 0x46, 0x4d, 0xd6,
	0x3d, 0x5f, 0xab, 0x1b, 0xd0, 0x6c, 0x6e, 0x17, 0x34, 0xa9, 0xcf, 0xd7, 0x6e, 0xf8, 0x77, 0xd8,
	0x93, 0xd5, 0x34, 0x98, 0xc5, 0xa9, 0xcf, 0xe7, 0xa1, 0x01, 0xb2, 0xdc, 0xd8, 0x28, 0x37, 0x97,
	0x3c, 0x91, 0xd3, 0x56, 0x57, 0x6c, 0xc1, 0x61, 0xe6, 0xcf, 0x22, 0xca, 0xf3, 0x94, 0xad, 0x75,
	0xd9, 0x95, 0x5d, 0xbe, 0x5b, 0xeb, 0x62, 0x2f, 0x55, 0x0f, 0xad, 0x70, 0xf6, 0x04, 0x13, 0xb6,
	0x98, 0xa6, 0x8c, 0x72, 0xe6, 0x72, 0x3f, 0x64, 0x6e, 0x44, 0xa3, 0x38, 0x33, 0x1a, 0xca, 0x16,
	0x8a, 0x70, 0xfc, 0x90, 0x59, 0x02, 0x16, 0xda, 0x3c, 0xf1, 0x36, 0xb4, 0x7b, 0x4a, 0xab, 0x88,
	0x07, 0xed, 0x19, 0xec, 0x26, 0xa9, 0xff, 0x59, 0x88, 0xef, 0xd9, 0xc2, 0xd8, 0x6f, 0x6a, 0xad,
	0xdd, 0xd3, 0x67, 0x1d, 0xe5, 0xd9, 0xce, 0xd2, 0xb3, 0x1d, 0x33, 0x5a, 0x10, 0x28, 0x84, 0x43,
	0xb6, 0xc0, 0x3f, 0xc0, 0x5e, 0x92, 0x4f, 0x02, 0x7f, 0x2a, 0xaa, 0x5c, 0x8f, 0xa5, 0x06, 0x6a,
	0x6a, 0xad, 0x3a, 0xa9, 0x2b, 0x74, 0xc8, 0x16, 0x97, 0x2c, 0xc5, 0x43, 0xc0, 0x41, 0x3c, 0x73,
	0x33, 0x65, 0x39, 0x77, 0x2a, 0x3d, 0x67, 0x54, 0xe4, 0x8c, 0x97, 0x6b, 0x3b, 0xd8, 0x0c, 0xc1,
	0x60, 0x8b, 0xa0, 0x60, 0x33, 0x18, 0x43, 0xc0, 0x21, 0x4d, 0x36, 0x9b, 0x55, 0x9f, 0x34, 0xdb,
	0xf4, 0xb8, 0x68, 0x16, 0x6e, 0x60, 0xf8, 0x35, 0x18, 0x21, 0xfd, 0xe2, 0xa6, 0x71, 0xcc, 0x5d,
	0x2f, 0x4f, 0xa9, 0x70, 0xa6, 0x1b, 0xfa, 0x41, 0xe0, 0x67, 0xc6, 0x81, 0xdc, 0xd4, 0x51, 0x48,
	0xbf, 0x90, 0x38, 0xe6, 0x97, 0x05, 0x3b, 0x92, 0x24, 0x36, 0xa0, 0xea, 0xb1, 0x80, 0x71, 0xe6,
	0x19, 0xb8, 0xa9, 0xb5, 0x74, 0xb2, 0xbc, 0x8a, 0xad, 0xab, 0xe3, 0xfa, 0xd6, 0x0f, 0xd5, 0xd6,
	0x15, 0xb1, 0xda, 0xfa, 0x05, 0x82, 0xbd, 0xc7, 0xef, 0xb8, 0x29, 0xeb, 0x75, 0xd4, 0x38, 0xf9,
	0x57, 0x53, 0x71, 0x1c, 0x30, 0xea, 0xfd, 0x7f, 0x1c, 0x8f, 0x41, 0xe7, 0x59, 0x31, 0x40, 0x05,
	0xb2, 0xca, 0x33, 0xf5, 0xef, 0x7c, 0x59, 0x84, 0x2b, 0xf3, 0xbf, 0xaa, 0x5c, 0x96, 0x54, 0x8e,
	0x6c, 0xff, 0x2b, 0x13, 0xa4, 0x7c, 0xb0, 0x70, 0xaa, 0x4c, 0x66, 0x9d, 0xe8, 0x02, 0x10, 0x46,
	0xc6, 0xdf, 0x42, 0x6d, 0x65, 0x3b, 0x69, 0xf6, 0x3a, 0x79, 0x00, 0xf0, 0xf7, 0xd0, 0x90, 0x7d,
	0x53, 0xf6, 0xd9, 0xcf, 0x44, 0xb0, 0x2b, 0xb2, 0x77, 0x5d, 0x80, 0xa4, 0xc0, 0xf0, 0x0b, 0xd0,
	0x43, 0xc6, 0xa9, 0x47, 0x39, 0x95, 0x69, 0xab, 0x93, 0xd5, 0xfd, 0xa6, 0xac, 0xef, 0xa0, 0xca,
	0x4d, 0x59, 0xd7, 0x51, 0xed, 0xa6, 0xac, 0x57, 0x91, 0xde, 0x7e, 0x0b, 0xb5, 0x55, 0x70, 0xf1,
	0x37, 0x80, 0x6f, 0xad, 0xa1, 0x35, 0xfe, 0xd3, 0x72, 0x1d, 0xd2, 0xef, 0xbb, 0xb6, 0x63, 0x3a,
	0x7d, 0xb4, 0x85, 0x01, 0x2a, 0x66, 0xcf, 0xb9, 0xfe, 0xa3, 0x8f, 0x34, 0x71, 0xbe, 0x22, 0xe3,
	0x4f, 0x7d, 0x0b, 0x6d, 0xb7, 0x7f, 0x52, 0x7b, 0x92, 0x9f, 0x87, 0x5d, 0xa8, 0x16, 0xb5, 0x68,
	0x0b, 0x57, 0xa1, 0xf4, 0x7e, 0xfc, 0x0e, 0x69, 0xe2, 0x30, 0x32, 0x3f, 0xa0, 0xed, 0xf6, 0x3f,
	0x1a, 0xd4, 0xd7, 0x93, 0x8e, 0x8f, 0xe1, 0x68, 0x39, 0x6b, 0x60, 0xda, 0x03, 0xd7, 0x76, 0x88,
	0xe9, 0xf4, 0xdf, 0x7d, 0x44, 0x5b, 0xb8, 0x0e, 0x3a, 0xb9, 0xea, 0xb9, 0xe7, 0x6f, 0xce, 0x4f,
	0x91, 0x86, 0x0f, 0x61, 0xdf, 0xe9, 0xdb, 0x8e, 0x3b, 0x32, 0x3f, 0x48, 0x65, 0x9f, 0xa0, 0x6d,
	0x51, 0x3d, 0xbe, 0xb8, 0xe9, 0xf7, 0x1c, 0x97, 0x5c, 0xf5, 0x84, 0xd0, 0xb5, 0x07, 0xe6, 0xe9,
	0xd9, 0x39, 0x2a, 0xe1, 0x23, 0x38, 0xe8, 0x8d, 0xad, 0xeb, 0xa1, 0x2d, 0xa0, 0xb3, 0x5f, 0x4f,
	0x5d, 0x01, 0x97, 0xf1, 0x01, 0x34, 0x1e, 0x60, 0x01, 0xed, 0xb4, 0x7f, 0x84, 0xc6, 0xa3, 0xaf,
	0x07, 0xd6, 0xa1, 0x6c, 0x8d, 0xad, 0xe2, 0xc5, 0x85, 0xac, 0xdc, 0x7e, 0x0d, 0xf8, 0xe9, 0xe7,
	0x01, 0x37, 0xa0, 0x66, 0x5a, 0x63, 0xeb, 0xe3, 0x68, 0x7c, 0x6b, 0xab, 0x17, 0x13, 0xdb, 0x44,
	0x1a, 0xae, 0xc1, 0x4e, 0xbf, 0x77, 0x69, 0x9b, 0xa8, 0x74, 0xf1, 0xf6, 0xd3, 0x9b, 0x99, 0xcf,
	0xe7, 0xf9, 0xa4, 0x33, 0x8d, 0xc3, 0x6e, 0xf1, 0x03, 0xc4, 0x53, 0x61, 0x61, 0x1a, 0x75, 0x0b,
	0xeb, 0x75, 0xa7, 0x41, 0x9c, 0x7b, 0x45, 0x70, 0xba, 0xab, 0x00, 0x4d, 0x2a, 0x32, 0xf5, 0xaf,
	0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0xba, 0x57, 0xcd, 0x18, 0xd3, 0x06, 0x00, 0x00,
}
