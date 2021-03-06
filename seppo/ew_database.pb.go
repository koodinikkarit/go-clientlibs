// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ew_database.proto

package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EwDatabase struct {
	Id                             uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name                           string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	SongDatabaseId                 uint32 `protobuf:"varint,3,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	FilesystemPath                 string `protobuf:"bytes,4,opt,name=filesystemPath" json:"filesystemPath,omitempty"`
	MatiasClientId                 uint32 `protobuf:"varint,5,opt,name=matiasClientId" json:"matiasClientId,omitempty"`
	RemoveSongsFromEwDatabase      uint32 `protobuf:"varint,6,opt,name=removeSongsFromEwDatabase" json:"removeSongsFromEwDatabase,omitempty"`
	RemoveSongsFromSongDatabase    uint32 `protobuf:"varint,7,opt,name=removeSongsFromSongDatabase" json:"removeSongsFromSongDatabase,omitempty"`
	VariationVersionConflictAction uint32 `protobuf:"varint,8,opt,name=variationVersionConflictAction" json:"variationVersionConflictAction,omitempty"`
	Accepted                       bool   `protobuf:"varint,9,opt,name=accepted" json:"accepted,omitempty"`
}

func (m *EwDatabase) Reset()                    { *m = EwDatabase{} }
func (m *EwDatabase) String() string            { return proto.CompactTextString(m) }
func (*EwDatabase) ProtoMessage()               {}
func (*EwDatabase) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *EwDatabase) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *EwDatabase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EwDatabase) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *EwDatabase) GetFilesystemPath() string {
	if m != nil {
		return m.FilesystemPath
	}
	return ""
}

func (m *EwDatabase) GetMatiasClientId() uint32 {
	if m != nil {
		return m.MatiasClientId
	}
	return 0
}

func (m *EwDatabase) GetRemoveSongsFromEwDatabase() uint32 {
	if m != nil {
		return m.RemoveSongsFromEwDatabase
	}
	return 0
}

func (m *EwDatabase) GetRemoveSongsFromSongDatabase() uint32 {
	if m != nil {
		return m.RemoveSongsFromSongDatabase
	}
	return 0
}

func (m *EwDatabase) GetVariationVersionConflictAction() uint32 {
	if m != nil {
		return m.VariationVersionConflictAction
	}
	return 0
}

func (m *EwDatabase) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

type SearchEwDatabasesRequest struct {
	Offset         uint32 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit          uint32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	SongDatabaseId uint32 `protobuf:"varint,3,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	MatiasClientId uint32 `protobuf:"varint,4,opt,name=matiasClientId" json:"matiasClientId,omitempty"`
}

func (m *SearchEwDatabasesRequest) Reset()                    { *m = SearchEwDatabasesRequest{} }
func (m *SearchEwDatabasesRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchEwDatabasesRequest) ProtoMessage()               {}
func (*SearchEwDatabasesRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *SearchEwDatabasesRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SearchEwDatabasesRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchEwDatabasesRequest) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *SearchEwDatabasesRequest) GetMatiasClientId() uint32 {
	if m != nil {
		return m.MatiasClientId
	}
	return 0
}

type SearchEwDatabasesResponse struct {
	EwDatabases    []*EwDatabase `protobuf:"bytes,1,rep,name=ewDatabases" json:"ewDatabases,omitempty"`
	MaxEwDatabases uint32        `protobuf:"varint,2,opt,name=maxEwDatabases" json:"maxEwDatabases,omitempty"`
}

func (m *SearchEwDatabasesResponse) Reset()                    { *m = SearchEwDatabasesResponse{} }
func (m *SearchEwDatabasesResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchEwDatabasesResponse) ProtoMessage()               {}
func (*SearchEwDatabasesResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *SearchEwDatabasesResponse) GetEwDatabases() []*EwDatabase {
	if m != nil {
		return m.EwDatabases
	}
	return nil
}

func (m *SearchEwDatabasesResponse) GetMaxEwDatabases() uint32 {
	if m != nil {
		return m.MaxEwDatabases
	}
	return 0
}

type FetchEwDatabaseByIdRequest struct {
	EwDatabaseIds []uint32 `protobuf:"varint,1,rep,packed,name=ewDatabaseIds" json:"ewDatabaseIds,omitempty"`
}

func (m *FetchEwDatabaseByIdRequest) Reset()                    { *m = FetchEwDatabaseByIdRequest{} }
func (m *FetchEwDatabaseByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchEwDatabaseByIdRequest) ProtoMessage()               {}
func (*FetchEwDatabaseByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *FetchEwDatabaseByIdRequest) GetEwDatabaseIds() []uint32 {
	if m != nil {
		return m.EwDatabaseIds
	}
	return nil
}

type FetchEwDatabaseByIdResponse struct {
	EwDatabases []*EwDatabase `protobuf:"bytes,1,rep,name=ewDatabases" json:"ewDatabases,omitempty"`
}

func (m *FetchEwDatabaseByIdResponse) Reset()                    { *m = FetchEwDatabaseByIdResponse{} }
func (m *FetchEwDatabaseByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchEwDatabaseByIdResponse) ProtoMessage()               {}
func (*FetchEwDatabaseByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *FetchEwDatabaseByIdResponse) GetEwDatabases() []*EwDatabase {
	if m != nil {
		return m.EwDatabases
	}
	return nil
}

type CreateEwDatabaseRequest struct {
	SongDatabaseId uint32 `protobuf:"varint,1,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	Name           string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	FilesystemPath string `protobuf:"bytes,3,opt,name=filesystemPath" json:"filesystemPath,omitempty"`
	MatiasClientId uint32 `protobuf:"varint,4,opt,name=matiasClientId" json:"matiasClientId,omitempty"`
}

func (m *CreateEwDatabaseRequest) Reset()                    { *m = CreateEwDatabaseRequest{} }
func (m *CreateEwDatabaseRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateEwDatabaseRequest) ProtoMessage()               {}
func (*CreateEwDatabaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *CreateEwDatabaseRequest) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *CreateEwDatabaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateEwDatabaseRequest) GetFilesystemPath() string {
	if m != nil {
		return m.FilesystemPath
	}
	return ""
}

func (m *CreateEwDatabaseRequest) GetMatiasClientId() uint32 {
	if m != nil {
		return m.MatiasClientId
	}
	return 0
}

type CreateEwDatabaseResponse struct {
	EwDatabase *EwDatabase `protobuf:"bytes,1,opt,name=ewDatabase" json:"ewDatabase,omitempty"`
}

func (m *CreateEwDatabaseResponse) Reset()                    { *m = CreateEwDatabaseResponse{} }
func (m *CreateEwDatabaseResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateEwDatabaseResponse) ProtoMessage()               {}
func (*CreateEwDatabaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *CreateEwDatabaseResponse) GetEwDatabase() *EwDatabase {
	if m != nil {
		return m.EwDatabase
	}
	return nil
}

type UpdateEwDatabaseRequest struct {
	EwDatabaseId                   uint32 `protobuf:"varint,1,opt,name=ewDatabaseId" json:"ewDatabaseId,omitempty"`
	Name                           string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	SongDatabaseId                 uint32 `protobuf:"varint,3,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	RemoveSongsFromEwDatabase      uint32 `protobuf:"varint,4,opt,name=removeSongsFromEwDatabase" json:"removeSongsFromEwDatabase,omitempty"`
	RemoveSongsFromSongDatabase    uint32 `protobuf:"varint,5,opt,name=removeSongsFromSongDatabase" json:"removeSongsFromSongDatabase,omitempty"`
	VariationVersionConflictAction uint32 `protobuf:"varint,6,opt,name=variationVersionConflictAction" json:"variationVersionConflictAction,omitempty"`
	Accepted                       bool   `protobuf:"varint,7,opt,name=accepted" json:"accepted,omitempty"`
}

func (m *UpdateEwDatabaseRequest) Reset()                    { *m = UpdateEwDatabaseRequest{} }
func (m *UpdateEwDatabaseRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateEwDatabaseRequest) ProtoMessage()               {}
func (*UpdateEwDatabaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *UpdateEwDatabaseRequest) GetEwDatabaseId() uint32 {
	if m != nil {
		return m.EwDatabaseId
	}
	return 0
}

func (m *UpdateEwDatabaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateEwDatabaseRequest) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *UpdateEwDatabaseRequest) GetRemoveSongsFromEwDatabase() uint32 {
	if m != nil {
		return m.RemoveSongsFromEwDatabase
	}
	return 0
}

func (m *UpdateEwDatabaseRequest) GetRemoveSongsFromSongDatabase() uint32 {
	if m != nil {
		return m.RemoveSongsFromSongDatabase
	}
	return 0
}

func (m *UpdateEwDatabaseRequest) GetVariationVersionConflictAction() uint32 {
	if m != nil {
		return m.VariationVersionConflictAction
	}
	return 0
}

func (m *UpdateEwDatabaseRequest) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

type UpdateEwDatabaseResponse struct {
	EwDatabase *EwDatabase `protobuf:"bytes,1,opt,name=ewDatabase" json:"ewDatabase,omitempty"`
	Success    bool        `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *UpdateEwDatabaseResponse) Reset()                    { *m = UpdateEwDatabaseResponse{} }
func (m *UpdateEwDatabaseResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateEwDatabaseResponse) ProtoMessage()               {}
func (*UpdateEwDatabaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *UpdateEwDatabaseResponse) GetEwDatabase() *EwDatabase {
	if m != nil {
		return m.EwDatabase
	}
	return nil
}

func (m *UpdateEwDatabaseResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RemoveEwDatabaseRequest struct {
	EwDatabaseId uint32 `protobuf:"varint,1,opt,name=ewDatabaseId" json:"ewDatabaseId,omitempty"`
}

func (m *RemoveEwDatabaseRequest) Reset()                    { *m = RemoveEwDatabaseRequest{} }
func (m *RemoveEwDatabaseRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveEwDatabaseRequest) ProtoMessage()               {}
func (*RemoveEwDatabaseRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *RemoveEwDatabaseRequest) GetEwDatabaseId() uint32 {
	if m != nil {
		return m.EwDatabaseId
	}
	return 0
}

type RemoveEwDatabaseResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *RemoveEwDatabaseResponse) Reset()                    { *m = RemoveEwDatabaseResponse{} }
func (m *RemoveEwDatabaseResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveEwDatabaseResponse) ProtoMessage()               {}
func (*RemoveEwDatabaseResponse) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *RemoveEwDatabaseResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*EwDatabase)(nil), "SeppoService.EwDatabase")
	proto.RegisterType((*SearchEwDatabasesRequest)(nil), "SeppoService.SearchEwDatabasesRequest")
	proto.RegisterType((*SearchEwDatabasesResponse)(nil), "SeppoService.SearchEwDatabasesResponse")
	proto.RegisterType((*FetchEwDatabaseByIdRequest)(nil), "SeppoService.FetchEwDatabaseByIdRequest")
	proto.RegisterType((*FetchEwDatabaseByIdResponse)(nil), "SeppoService.FetchEwDatabaseByIdResponse")
	proto.RegisterType((*CreateEwDatabaseRequest)(nil), "SeppoService.CreateEwDatabaseRequest")
	proto.RegisterType((*CreateEwDatabaseResponse)(nil), "SeppoService.CreateEwDatabaseResponse")
	proto.RegisterType((*UpdateEwDatabaseRequest)(nil), "SeppoService.UpdateEwDatabaseRequest")
	proto.RegisterType((*UpdateEwDatabaseResponse)(nil), "SeppoService.UpdateEwDatabaseResponse")
	proto.RegisterType((*RemoveEwDatabaseRequest)(nil), "SeppoService.RemoveEwDatabaseRequest")
	proto.RegisterType((*RemoveEwDatabaseResponse)(nil), "SeppoService.RemoveEwDatabaseResponse")
}

func init() { proto.RegisterFile("ew_database.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcd, 0x6e, 0x13, 0x31,
	0x10, 0xd6, 0xe6, 0xbf, 0xd3, 0xa6, 0x12, 0x16, 0x22, 0x6e, 0x2b, 0xa1, 0xc8, 0x42, 0x28, 0xa7,
	0x1c, 0x80, 0x03, 0x42, 0x20, 0x41, 0x03, 0x91, 0x72, 0x43, 0x1b, 0x40, 0xe2, 0x84, 0xdc, 0xf5,
	0xa4, 0xb5, 0x94, 0xb5, 0x97, 0xb5, 0x9b, 0xd2, 0x13, 0x2f, 0xc0, 0x03, 0xf0, 0x02, 0x3c, 0x16,
	0xef, 0x82, 0xe2, 0xdd, 0xb0, 0xce, 0x66, 0x1b, 0x28, 0xe1, 0x16, 0x7f, 0xf9, 0x66, 0xfc, 0x7d,
	0xfe, 0x32, 0x13, 0xb8, 0x83, 0x57, 0x9f, 0x04, 0xb7, 0xfc, 0x8c, 0x1b, 0x1c, 0x26, 0xa9, 0xb6,
	0x9a, 0x1c, 0x4c, 0x31, 0x49, 0xf4, 0x14, 0xd3, 0x85, 0x8c, 0x90, 0x7d, 0xab, 0x03, 0xbc, 0xb9,
	0x7a, 0x9d, 0x53, 0xc8, 0x21, 0xd4, 0xa4, 0xa0, 0x41, 0x3f, 0x18, 0x74, 0xc3, 0x9a, 0x14, 0x84,
	0x40, 0x43, 0xf1, 0x18, 0x69, 0xad, 0x1f, 0x0c, 0xf6, 0x42, 0xf7, 0x99, 0x3c, 0x84, 0x43, 0xa3,
	0xd5, 0xf9, 0xaa, 0x66, 0x22, 0x68, 0xdd, 0xf1, 0x4b, 0xe8, 0x92, 0x37, 0x93, 0x73, 0x34, 0xd7,
	0xc6, 0x62, 0xfc, 0x96, 0xdb, 0x0b, 0xda, 0x70, 0x5d, 0x4a, 0xe8, 0x92, 0x17, 0x73, 0x2b, 0xb9,
	0x19, 0xcd, 0x25, 0x2a, 0x3b, 0x11, 0xb4, 0x99, 0xf5, 0x5b, 0x47, 0xc9, 0x73, 0x38, 0x4a, 0x31,
	0xd6, 0x0b, 0x9c, 0x6a, 0x75, 0x6e, 0xc6, 0xa9, 0x8e, 0x0b, 0xe1, 0xb4, 0xe5, 0x4a, 0x6e, 0x26,
	0x90, 0x97, 0x70, 0x52, 0xfa, 0x72, 0xea, 0xc9, 0xa5, 0x6d, 0x57, 0xbf, 0x8d, 0x42, 0xc6, 0x70,
	0x7f, 0xc1, 0x53, 0xc9, 0xad, 0xd4, 0xea, 0x03, 0xa6, 0x46, 0x6a, 0x35, 0xd2, 0x6a, 0x36, 0x97,
	0x91, 0x7d, 0x15, 0x2d, 0x41, 0xda, 0x71, 0x4d, 0xfe, 0xc0, 0x22, 0xc7, 0xd0, 0xe1, 0x51, 0x84,
	0x89, 0x45, 0x41, 0xf7, 0xfa, 0xc1, 0xa0, 0x13, 0xfe, 0x3e, 0xb3, 0xef, 0x01, 0xd0, 0x29, 0xf2,
	0x34, 0xba, 0x28, 0xa4, 0x9b, 0x10, 0x3f, 0x5f, 0xa2, 0xb1, 0xe4, 0x1e, 0xb4, 0xf4, 0x6c, 0x66,
	0xd0, 0xe6, 0x01, 0xe5, 0x27, 0x72, 0x17, 0x9a, 0x73, 0x19, 0x4b, 0xeb, 0x52, 0xea, 0x86, 0xd9,
	0xe1, 0x36, 0x31, 0x95, 0x9e, 0xbf, 0x51, 0xf5, 0xfc, 0xec, 0x2b, 0x1c, 0x55, 0x28, 0x33, 0x89,
	0x56, 0x06, 0xc9, 0x33, 0xd8, 0xc7, 0x02, 0xa6, 0x41, 0xbf, 0x3e, 0xd8, 0x7f, 0x44, 0x87, 0xfe,
	0x4f, 0x6d, 0x58, 0xd4, 0x85, 0x3e, 0x39, 0x13, 0xf0, 0xc5, 0xeb, 0x9a, 0xfb, 0x28, 0xa1, 0xec,
	0x14, 0x8e, 0xc7, 0x68, 0xfd, 0xfb, 0x4f, 0xaf, 0x27, 0x62, 0xf5, 0x38, 0x0f, 0xa0, 0x5b, 0x34,
	0x9d, 0x88, 0x4c, 0x43, 0x37, 0x5c, 0x07, 0xd9, 0x47, 0x38, 0xa9, 0xec, 0xb1, 0xbb, 0x0d, 0xf6,
	0x23, 0x80, 0xde, 0x28, 0x45, 0x6e, 0xd1, 0x63, 0xe4, 0xe2, 0x36, 0xb3, 0x08, 0x2a, 0xb3, 0xb8,
	0x61, 0xdc, 0x4a, 0x63, 0x54, 0xff, 0xcb, 0x31, 0xaa, 0xce, 0xf1, 0x1d, 0xd0, 0x4d, 0x99, 0xb9,
	0xff, 0xa7, 0x00, 0x85, 0x25, 0xa7, 0x71, 0x9b, 0x7d, 0x8f, 0xcb, 0x7e, 0xd6, 0xa0, 0xf7, 0x3e,
	0x11, 0x95, 0xee, 0x19, 0x1c, 0xf8, 0x29, 0xe4, 0xde, 0xd7, 0xb0, 0x9d, 0x16, 0xcd, 0xd6, 0xc5,
	0xd0, 0xd8, 0x71, 0x31, 0x34, 0xff, 0xc7, 0x62, 0x68, 0xdd, 0x7a, 0x31, 0xb4, 0x4b, 0x8b, 0x41,
	0x01, 0xdd, 0x7c, 0xde, 0x5d, 0x53, 0x23, 0x14, 0xda, 0xe6, 0x32, 0x8a, 0xd0, 0x64, 0x33, 0xd7,
	0x09, 0x57, 0x47, 0xf6, 0x02, 0x7a, 0xa1, 0xb3, 0xfc, 0x4f, 0x71, 0xb2, 0x27, 0x40, 0x37, 0xcb,
	0x73, 0xb9, 0xde, 0xa5, 0xc1, 0xda, 0xa5, 0x67, 0x2d, 0xf7, 0x0f, 0xf5, 0xf8, 0x57, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xd1, 0x5d, 0x47, 0x4a, 0xb6, 0x06, 0x00, 0x00,
}
