// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CDEService.proto

package CDEService

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TileType int32

const (
	TileType_ImageTile     TileType = 0
	TileType_VideoTile     TileType = 1
	TileType_FeatureTile   TileType = 2
	TileType_AdvertiseTile TileType = 3
	TileType_CarouselTile  TileType = 4
)

var TileType_name = map[int32]string{
	0: "ImageTile",
	1: "VideoTile",
	2: "FeatureTile",
	3: "AdvertiseTile",
	4: "CarouselTile",
}

var TileType_value = map[string]int32{
	"ImageTile":     0,
	"VideoTile":     1,
	"FeatureTile":   2,
	"AdvertiseTile": 3,
	"CarouselTile":  4,
}

func (x TileType) String() string {
	return proto.EnumName(TileType_name, int32(x))
}

func (TileType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b71e49cbfefc313f, []int{0}
}

type SearchQuery struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *SearchQuery) Reset()         { *m = SearchQuery{} }
func (m *SearchQuery) String() string { return proto.CompactTextString(m) }
func (*SearchQuery) ProtoMessage()    {}
func (*SearchQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b71e49cbfefc313f, []int{0}
}

func (m *SearchQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchQuery.Unmarshal(m, b)
}
func (m *SearchQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchQuery.Marshal(b, m, deterministic)
}
func (m *SearchQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchQuery.Merge(m, src)
}
func (m *SearchQuery) XXX_Size() int {
	return xxx_messageInfo_SearchQuery.Size(m)
}
func (m *SearchQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchQuery.DiscardUnknown(m)
}

var xxx_messageInfo_SearchQuery proto.InternalMessageInfo

func (m *SearchQuery) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

type ContentTile struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	MediaUrl             string   `protobuf:"bytes,2,opt,name=mediaUrl,proto3" json:"mediaUrl,omitempty"`
	TileType             TileType `protobuf:"varint,3,opt,name=tileType,proto3,enum=CDEService.TileType" json:"tileType,omitempty"`
	Poster               []string `protobuf:"bytes,4,rep,name=poster,proto3" json:"poster,omitempty"`
	Portrait             []string `protobuf:"bytes,5,rep,name=portrait,proto3" json:"portrait,omitempty"`
	IsDetailPage         bool     `protobuf:"varint,6,opt,name=isDetailPage,proto3" json:"isDetailPage,omitempty"`
	PackageName          string   `protobuf:"bytes,7,opt,name=packageName,proto3" json:"packageName,omitempty"`
	ContentId            string   `protobuf:"bytes,8,opt,name=contentId,proto3" json:"contentId,omitempty"`
	Target               []string `protobuf:"bytes,9,rep,name=target,proto3" json:"target,omitempty"`
	RealeaseDate         string   `protobuf:"bytes,10,opt,name=realeaseDate,proto3" json:"realeaseDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *ContentTile) Reset()         { *m = ContentTile{} }
func (m *ContentTile) String() string { return proto.CompactTextString(m) }
func (*ContentTile) ProtoMessage()    {}
func (*ContentTile) Descriptor() ([]byte, []int) {
	return fileDescriptor_b71e49cbfefc313f, []int{1}
}

func (m *ContentTile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContentTile.Unmarshal(m, b)
}
func (m *ContentTile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContentTile.Marshal(b, m, deterministic)
}
func (m *ContentTile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentTile.Merge(m, src)
}
func (m *ContentTile) XXX_Size() int {
	return xxx_messageInfo_ContentTile.Size(m)
}
func (m *ContentTile) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentTile.DiscardUnknown(m)
}

var xxx_messageInfo_ContentTile proto.InternalMessageInfo

func (m *ContentTile) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ContentTile) GetMediaUrl() string {
	if m != nil {
		return m.MediaUrl
	}
	return ""
}

func (m *ContentTile) GetTileType() TileType {
	if m != nil {
		return m.TileType
	}
	return TileType_ImageTile
}

func (m *ContentTile) GetPoster() []string {
	if m != nil {
		return m.Poster
	}
	return nil
}

func (m *ContentTile) GetPortrait() []string {
	if m != nil {
		return m.Portrait
	}
	return nil
}

func (m *ContentTile) GetIsDetailPage() bool {
	if m != nil {
		return m.IsDetailPage
	}
	return false
}

func (m *ContentTile) GetPackageName() string {
	if m != nil {
		return m.PackageName
	}
	return ""
}

func (m *ContentTile) GetContentId() string {
	if m != nil {
		return m.ContentId
	}
	return ""
}

func (m *ContentTile) GetTarget() []string {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *ContentTile) GetRealeaseDate() string {
	if m != nil {
		return m.RealeaseDate
	}
	return ""
}

type MovieTile struct {
	RefId                string    `protobuf:"bytes,1,opt,name=ref_id,json=refId,proto3" json:"ref_id,omitempty"`
	Posters              *POSTERS  `protobuf:"bytes,2,opt,name=posters,proto3" json:"posters,omitempty"`
	Content              *CONTENT  `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Metadata             *METADATA `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
	TileType             TileType  `protobuf:"varint,5,opt,name=tileType,proto3,enum=CDEService.TileType" json:"tileType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-" bson:"-"`
	XXX_unrecognized     []byte    `json:"-" bson:"-"`
	XXX_sizecache        int32     `json:"-" bson:"-"`
}

func (m *MovieTile) Reset()         { *m = MovieTile{} }
func (m *MovieTile) String() string { return proto.CompactTextString(m) }
func (*MovieTile) ProtoMessage()    {}
func (*MovieTile) Descriptor() ([]byte, []int) {
	return fileDescriptor_b71e49cbfefc313f, []int{2}
}

func (m *MovieTile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MovieTile.Unmarshal(m, b)
}
func (m *MovieTile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MovieTile.Marshal(b, m, deterministic)
}
func (m *MovieTile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MovieTile.Merge(m, src)
}
func (m *MovieTile) XXX_Size() int {
	return xxx_messageInfo_MovieTile.Size(m)
}
func (m *MovieTile) XXX_DiscardUnknown() {
	xxx_messageInfo_MovieTile.DiscardUnknown(m)
}

var xxx_messageInfo_MovieTile proto.InternalMessageInfo

func (m *MovieTile) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *MovieTile) GetPosters() *POSTERS {
	if m != nil {
		return m.Posters
	}
	return nil
}

func (m *MovieTile) GetContent() *CONTENT {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *MovieTile) GetMetadata() *METADATA {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MovieTile) GetTileType() TileType {
	if m != nil {
		return m.TileType
	}
	return TileType_ImageTile
}

type POSTERS struct {
	Landscape            []string `protobuf:"bytes,1,rep,name=landscape,proto3" json:"landscape,omitempty"`
	Portrait             []string `protobuf:"bytes,2,rep,name=portrait,proto3" json:"portrait,omitempty"`
	Banner               []string `protobuf:"bytes,3,rep,name=banner,proto3" json:"banner,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *POSTERS) Reset()         { *m = POSTERS{} }
func (m *POSTERS) String() string { return proto.CompactTextString(m) }
func (*POSTERS) ProtoMessage()    {}
func (*POSTERS) Descriptor() ([]byte, []int) {
	return fileDescriptor_b71e49cbfefc313f, []int{3}
}

func (m *POSTERS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_POSTERS.Unmarshal(m, b)
}
func (m *POSTERS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_POSTERS.Marshal(b, m, deterministic)
}
func (m *POSTERS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_POSTERS.Merge(m, src)
}
func (m *POSTERS) XXX_Size() int {
	return xxx_messageInfo_POSTERS.Size(m)
}
func (m *POSTERS) XXX_DiscardUnknown() {
	xxx_messageInfo_POSTERS.DiscardUnknown(m)
}

var xxx_messageInfo_POSTERS proto.InternalMessageInfo

func (m *POSTERS) GetLandscape() []string {
	if m != nil {
		return m.Landscape
	}
	return nil
}

func (m *POSTERS) GetPortrait() []string {
	if m != nil {
		return m.Portrait
	}
	return nil
}

func (m *POSTERS) GetBanner() []string {
	if m != nil {
		return m.Banner
	}
	return nil
}

type METADATA struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	ReleaseDate          string   `protobuf:"bytes,2,opt,name=releaseDate,proto3" json:"releaseDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *METADATA) Reset()         { *m = METADATA{} }
func (m *METADATA) String() string { return proto.CompactTextString(m) }
func (*METADATA) ProtoMessage()    {}
func (*METADATA) Descriptor() ([]byte, []int) {
	return fileDescriptor_b71e49cbfefc313f, []int{4}
}

func (m *METADATA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_METADATA.Unmarshal(m, b)
}
func (m *METADATA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_METADATA.Marshal(b, m, deterministic)
}
func (m *METADATA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_METADATA.Merge(m, src)
}
func (m *METADATA) XXX_Size() int {
	return xxx_messageInfo_METADATA.Size(m)
}
func (m *METADATA) XXX_DiscardUnknown() {
	xxx_messageInfo_METADATA.DiscardUnknown(m)
}

var xxx_messageInfo_METADATA proto.InternalMessageInfo

func (m *METADATA) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *METADATA) GetReleaseDate() string {
	if m != nil {
		return m.ReleaseDate
	}
	return ""
}

type CONTENT struct {
	DetailPage           bool     `protobuf:"varint,1,opt,name=detailPage,proto3" json:"detailPage,omitempty"`
	Package              string   `protobuf:"bytes,2,opt,name=package,proto3" json:"package,omitempty"`
	Target               []string `protobuf:"bytes,3,rep,name=target,proto3" json:"target,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	PlaystoreUrl         string   `protobuf:"bytes,5,opt,name=playstoreUrl,proto3" json:"playstoreUrl,omitempty"`
	UseAlternate         bool     `protobuf:"varint,6,opt,name=useAlternate,proto3" json:"useAlternate,omitempty"`
	AlternateUrl         string   `protobuf:"bytes,7,opt,name=alternateUrl,proto3" json:"alternateUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" bson:"-"`
}

func (m *CONTENT) Reset()         { *m = CONTENT{} }
func (m *CONTENT) String() string { return proto.CompactTextString(m) }
func (*CONTENT) ProtoMessage()    {}
func (*CONTENT) Descriptor() ([]byte, []int) {
	return fileDescriptor_b71e49cbfefc313f, []int{5}
}

func (m *CONTENT) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CONTENT.Unmarshal(m, b)
}
func (m *CONTENT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CONTENT.Marshal(b, m, deterministic)
}
func (m *CONTENT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CONTENT.Merge(m, src)
}
func (m *CONTENT) XXX_Size() int {
	return xxx_messageInfo_CONTENT.Size(m)
}
func (m *CONTENT) XXX_DiscardUnknown() {
	xxx_messageInfo_CONTENT.DiscardUnknown(m)
}

var xxx_messageInfo_CONTENT proto.InternalMessageInfo

func (m *CONTENT) GetDetailPage() bool {
	if m != nil {
		return m.DetailPage
	}
	return false
}

func (m *CONTENT) GetPackage() string {
	if m != nil {
		return m.Package
	}
	return ""
}

func (m *CONTENT) GetTarget() []string {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *CONTENT) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *CONTENT) GetPlaystoreUrl() string {
	if m != nil {
		return m.PlaystoreUrl
	}
	return ""
}

func (m *CONTENT) GetUseAlternate() bool {
	if m != nil {
		return m.UseAlternate
	}
	return false
}

func (m *CONTENT) GetAlternateUrl() string {
	if m != nil {
		return m.AlternateUrl
	}
	return ""
}

func init() {
	proto.RegisterEnum("CDEService.TileType", TileType_name, TileType_value)
	proto.RegisterType((*SearchQuery)(nil), "CDEService.SearchQuery")
	proto.RegisterType((*ContentTile)(nil), "CDEService.ContentTile")
	proto.RegisterType((*MovieTile)(nil), "CDEService.MovieTile")
	proto.RegisterType((*POSTERS)(nil), "CDEService.POSTERS")
	proto.RegisterType((*METADATA)(nil), "CDEService.METADATA")
	proto.RegisterType((*CONTENT)(nil), "CDEService.CONTENT")
}

func init() { proto.RegisterFile("CDEService.proto", fileDescriptor_b71e49cbfefc313f) }

var fileDescriptor_b71e49cbfefc313f = []byte{
	// 621 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xf9, 0xce, 0xa4, 0x85, 0xb2, 0x94, 0xb2, 0xaa, 0x2a, 0x14, 0x99, 0x4b, 0x54, 0x89,
	0xb6, 0x2a, 0x37, 0x6e, 0x21, 0x09, 0x52, 0x0f, 0xfd, 0xc0, 0x31, 0x5c, 0x90, 0x8a, 0xa6, 0xf1,
	0x34, 0xac, 0x70, 0xbc, 0x66, 0xbd, 0xa9, 0x94, 0x2b, 0x7f, 0x81, 0xdf, 0xc4, 0x2f, 0xe0, 0xc8,
	0xb5, 0x3f, 0x04, 0xed, 0xae, 0x9d, 0xd8, 0x95, 0x90, 0xb8, 0xf9, 0xbd, 0x79, 0xd9, 0x9d, 0x7d,
	0xf3, 0x32, 0xb0, 0x33, 0x1a, 0x4f, 0xa6, 0xa4, 0xee, 0xc4, 0x8c, 0x8e, 0x52, 0x25, 0xb5, 0x64,
	0xb0, 0x61, 0xf6, 0x0f, 0xe6, 0x52, 0xce, 0x63, 0x3a, 0xc6, 0x54, 0x1c, 0x63, 0x92, 0x48, 0x8d,
	0x5a, 0xc8, 0x24, 0x73, 0x4a, 0xff, 0x15, 0xf4, 0xa6, 0x84, 0x6a, 0xf6, 0xf5, 0xc3, 0x92, 0xd4,
	0x8a, 0xed, 0x42, 0xf3, 0xbb, 0xf9, 0xe0, 0x5e, 0xdf, 0x1b, 0x74, 0x03, 0x07, 0xfc, 0x5f, 0x35,
	0xe8, 0x8d, 0x64, 0xa2, 0x29, 0xd1, 0xa1, 0x88, 0xc9, 0xa8, 0xb4, 0xd0, 0x31, 0x15, 0x2a, 0x0b,
	0xd8, 0x3e, 0x74, 0x16, 0x14, 0x09, 0xfc, 0xa8, 0x62, 0x5e, 0xb3, 0x85, 0x35, 0x66, 0x27, 0xd0,
	0xd1, 0x22, 0xa6, 0x70, 0x95, 0x12, 0xaf, 0xf7, 0xbd, 0xc1, 0xe3, 0xd3, 0xdd, 0xa3, 0x52, 0xd7,
	0x61, 0x5e, 0x0b, 0xd6, 0x2a, 0xb6, 0x07, 0xad, 0x54, 0x66, 0x9a, 0x14, 0x6f, 0xf4, 0xeb, 0x83,
	0x6e, 0x90, 0x23, 0x73, 0x4b, 0x2a, 0x95, 0x56, 0x28, 0x34, 0x6f, 0xda, 0xca, 0x1a, 0x33, 0x1f,
	0xb6, 0x44, 0x36, 0x26, 0x8d, 0x22, 0xbe, 0xc2, 0x39, 0xf1, 0x56, 0xdf, 0x1b, 0x74, 0x82, 0x0a,
	0xc7, 0xfa, 0xd0, 0x4b, 0x71, 0xf6, 0x0d, 0xe7, 0x74, 0x81, 0x0b, 0xe2, 0x6d, 0xdb, 0x68, 0x99,
	0x62, 0x07, 0xd0, 0x9d, 0xb9, 0xc7, 0x9e, 0x45, 0xbc, 0x63, 0xeb, 0x1b, 0xc2, 0xf4, 0xa5, 0x51,
	0xcd, 0x49, 0xf3, 0xae, 0xeb, 0xcb, 0x21, 0x73, 0xb7, 0x22, 0x8c, 0x09, 0x33, 0x1a, 0xa3, 0x26,
	0x0e, 0xf6, 0x87, 0x15, 0xce, 0xbf, 0xf7, 0xa0, 0x7b, 0x2e, 0xef, 0x04, 0x59, 0x17, 0x9f, 0x43,
	0x4b, 0xd1, 0xed, 0x17, 0x11, 0x15, 0x36, 0x2a, 0xba, 0x3d, 0x8b, 0xd8, 0x6b, 0x68, 0xbb, 0xa7,
	0x66, 0xd6, 0xc5, 0xde, 0xe9, 0xb3, 0xb2, 0x53, 0x57, 0x97, 0xd3, 0x70, 0x12, 0x4c, 0x83, 0x42,
	0x63, 0xe4, 0x79, 0x73, 0xd6, 0xd8, 0x07, 0xf2, 0xd1, 0xe5, 0x45, 0x38, 0xb9, 0x08, 0x83, 0x42,
	0x63, 0x06, 0xb1, 0x20, 0x8d, 0x11, 0x6a, 0xe4, 0x0d, 0xab, 0xaf, 0x0c, 0xe2, 0x7c, 0x12, 0x0e,
	0xc7, 0xc3, 0x70, 0x18, 0xac, 0x55, 0x95, 0xd1, 0x35, 0xff, 0x67, 0x74, 0xfe, 0x67, 0x68, 0xe7,
	0x6d, 0x1a, 0x2f, 0x63, 0x4c, 0xa2, 0x6c, 0x86, 0xa9, 0x49, 0x8b, 0x31, 0x6c, 0x43, 0x54, 0x66,
	0x59, 0x7b, 0x30, 0xcb, 0x3d, 0x68, 0xdd, 0x60, 0x92, 0x90, 0xe2, 0x75, 0xe7, 0xb3, 0x43, 0xfe,
	0x3b, 0xe8, 0x14, 0x4d, 0xfe, 0x23, 0x87, 0x7d, 0xe8, 0x29, 0xda, 0x0c, 0xc2, 0x45, 0xb1, 0x4c,
	0xf9, 0x7f, 0x3c, 0x68, 0xe7, 0xce, 0xb0, 0x97, 0x00, 0xd1, 0x26, 0x31, 0x9e, 0x4d, 0x4c, 0x89,
	0x61, 0x1c, 0xda, 0x79, 0x38, 0xf2, 0x93, 0x0a, 0x58, 0x4a, 0x42, 0xbd, 0x92, 0x04, 0x06, 0x0d,
	0x6d, 0xcc, 0x6a, 0x58, 0xb9, 0xfd, 0x36, 0xe9, 0x48, 0x63, 0x5c, 0x65, 0x5a, 0x2a, 0x32, 0xff,
	0x8f, 0xa6, 0x4b, 0x47, 0x99, 0x33, 0x9a, 0x65, 0x46, 0xc3, 0x58, 0x93, 0x4a, 0x4c, 0xe3, 0x79,
	0x7a, 0xcb, 0x9c, 0xd1, 0x60, 0x01, 0xcc, 0x39, 0x2e, 0xbe, 0x15, 0xee, 0xf0, 0x1a, 0x3a, 0xc5,
	0x50, 0xd8, 0x36, 0x74, 0xcf, 0x16, 0x38, 0xb7, 0x81, 0xdb, 0x79, 0x64, 0xe0, 0x27, 0x11, 0x91,
	0xb4, 0xd0, 0x63, 0x4f, 0xa0, 0xf7, 0x9e, 0x50, 0x2f, 0x95, 0xab, 0xd7, 0xd8, 0x53, 0xd8, 0x1e,
	0x46, 0x77, 0xa4, 0xb4, 0xc8, 0x1c, 0x55, 0x67, 0x3b, 0xb0, 0x35, 0x42, 0x25, 0x97, 0x19, 0xc5,
	0x96, 0x69, 0x9c, 0x5e, 0x43, 0x69, 0xbd, 0xb0, 0x2b, 0x68, 0xb9, 0x05, 0xc2, 0x5e, 0x94, 0x63,
	0x51, 0x5a, 0x2a, 0xfb, 0x95, 0x42, 0x69, 0x8f, 0xf8, 0xec, 0xc7, 0xef, 0xfb, 0x9f, 0xb5, 0x2d,
	0xbf, 0x7d, 0x9c, 0x59, 0xf9, 0x5b, 0xef, 0xf0, 0xc4, 0xbb, 0x69, 0xd9, 0xcd, 0xf4, 0xe6, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x38, 0x11, 0x91, 0xd7, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CDEServiceClient is the client API for CDEService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CDEServiceClient interface {
	Search(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (CDEService_SearchClient, error)
}

type cDEServiceClient struct {
	cc *grpc.ClientConn
}

func NewCDEServiceClient(cc *grpc.ClientConn) CDEServiceClient {
	return &cDEServiceClient{cc}
}

func (c *cDEServiceClient) Search(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (CDEService_SearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CDEService_serviceDesc.Streams[0], "/CDEService.CDEService/Search", opts...)
	if err != nil {
		return nil, err
	}
	x := &cDEServiceSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CDEService_SearchClient interface {
	Recv() (*ContentTile, error)
	grpc.ClientStream
}

type cDEServiceSearchClient struct {
	grpc.ClientStream
}

func (x *cDEServiceSearchClient) Recv() (*ContentTile, error) {
	m := new(ContentTile)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CDEServiceServer is the server API for CDEService service.
type CDEServiceServer interface {
	Search(*SearchQuery, CDEService_SearchServer) error
}

// UnimplementedCDEServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCDEServiceServer struct {
}

func (*UnimplementedCDEServiceServer) Search(req *SearchQuery, srv CDEService_SearchServer) error {
	return status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterCDEServiceServer(s *grpc.Server, srv CDEServiceServer) {
	s.RegisterService(&_CDEService_serviceDesc, srv)
}

func _CDEService_Search_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CDEServiceServer).Search(m, &cDEServiceSearchServer{stream})
}

type CDEService_SearchServer interface {
	Send(*ContentTile) error
	grpc.ServerStream
}

type cDEServiceSearchServer struct {
	grpc.ServerStream
}

func (x *cDEServiceSearchServer) Send(m *ContentTile) error {
	return x.ServerStream.SendMsg(m)
}

var _CDEService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "CDEService.CDEService",
	HandlerType: (*CDEServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Search",
			Handler:       _CDEService_Search_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "CDEService.proto",
}
