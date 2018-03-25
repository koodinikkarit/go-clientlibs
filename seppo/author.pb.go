// Code generated by protoc-gen-go. DO NOT EDIT.
// source: author.proto

/*
Package SeppoService is a generated protocol buffer package.

It is generated from these files:
	author.proto
	copyright.proto
	ew_database.proto
	ew_song.proto
	language.proto
	log.proto
	matias_client.proto
	schedule.proto
	seppo_service.proto
	song.proto
	song_database.proto
	tag.proto
	variation.proto
	variation_version.proto

It has these top-level messages:
	Author
	FetchAuthorByIdRequest
	FetchAuthorByIdResponse
	SearchAuthorsRequest
	SearchAuthorsResponse
	CreateAuthorRequest
	CreateAuthorResponse
	UpdateAuthorRequest
	UpdateAuthorResponse
	RemoveAuthorRequest
	RemoveAuthorResponse
	Copyright
	FetchCopyrightByIdRequest
	FetchCopyrightByIdResponse
	SearchCopyrightsRequest
	SearchCopyrightsResponse
	CreateCopyrightRequest
	CreateCopyrightResponse
	UpdateCopyrightRequest
	UpdateCopyrightResponse
	RemoveCopyrightRequest
	RemoveCopyrightResponse
	EwDatabase
	SearchEwDatabasesRequest
	SearchEwDatabasesResponse
	FetchEwDatabaseByIdRequest
	FetchEwDatabaseByIdResponse
	CreateEwDatabaseRequest
	CreateEwDatabaseResponse
	UpdateEwDatabaseRequest
	UpdateEwDatabaseResponse
	RemoveEwDatabaseRequest
	RemoveEwDatabaseResponse
	EwSong
	Language
	SearchLanguagesRequest
	SearchLanguagesResponse
	FetchLanguageByIdRequest
	FetchLanguageByIdResponse
	CreateLanguageRequest
	CreateLanguageResponse
	UpdateLanguageRequest
	UpdateLanguageResponse
	RemoveLanguageRequest
	RemoveLanguageResponse
	Log
	SearchLogsRequest
	SearchLogsResponse
	MatiasClient
	SearchMatiasClientsRequest
	SearchMatiasClientsResponse
	FetchMatiasClientRequest
	FetchMatiasClientResponse
	FetchMatiasClientEwDatabasesRequest
	FetchMatiasClientEwDatabasesResponse
	UpdateMatiasClientRequest
	UpdateMatiasClientResponse
	Schedule
	SearchSchedulesRequest
	SearchSchedulesResponse
	FetchScheduleByIdRequest
	FetchScheduleByIdResponse
	CreateScheduleRequest
	CreateScheduleResponse
	UpdateScheduleRequest
	UpdateScheduleResponse
	RemoveScheduleRequest
	RemoveScheduleResponse
	Song
	SongDatabase
	SearchSongDatabasesRequest
	SearchSongDatabasesResponse
	FetchSongDatabaseByIdRequest
	FetchSongDatabaseByIdResponse
	CreateSongDatabaseRequest
	CreateSongDatabaseResponse
	UpdateSongDatabaseRequest
	UpdateSongDatabaseResponse
	RemoveSongDatabaseRequest
	RemoveSongDatabaseResponse
	Tag
	SearchTagsRequest
	SearchTagsResponse
	FetchTagByIdRequest
	FetchTagByIdResponse
	CreateTagRequest
	CreateTagResponse
	UpdateTagRequest
	UpdateTagResponse
	RemoveTagRequest
	RemoveTagResponse
	Variation
	SearchVariationsRequest
	SearchVariationsResponse
	FetchVariationByIdRequest
	FetchVariationByIdResponse
	CreateVariationRequest
	CreateVariationResponse
	UpdateVariationRequest
	UpdateVariationResponse
	RemoveVariationRequest
	RemoveVariationResponse
	VariationVersion
	FetchNewestVariationVersionByVariationIdRequest
	FetchNewestVariationVersionByVariationIdResponse
	FetchVariationVersionByIdRequest
	FetchVariationVersionByIdResponse
*/
package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Author struct {
	Id   uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Author) Reset()                    { *m = Author{} }
func (m *Author) String() string            { return proto.CompactTextString(m) }
func (*Author) ProtoMessage()               {}
func (*Author) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Author) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Author) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type FetchAuthorByIdRequest struct {
	AuthorIds []uint32 `protobuf:"varint,1,rep,packed,name=authorIds" json:"authorIds,omitempty"`
}

func (m *FetchAuthorByIdRequest) Reset()                    { *m = FetchAuthorByIdRequest{} }
func (m *FetchAuthorByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchAuthorByIdRequest) ProtoMessage()               {}
func (*FetchAuthorByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FetchAuthorByIdRequest) GetAuthorIds() []uint32 {
	if m != nil {
		return m.AuthorIds
	}
	return nil
}

type FetchAuthorByIdResponse struct {
	Authors []*Author `protobuf:"bytes,1,rep,name=authors" json:"authors,omitempty"`
}

func (m *FetchAuthorByIdResponse) Reset()                    { *m = FetchAuthorByIdResponse{} }
func (m *FetchAuthorByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchAuthorByIdResponse) ProtoMessage()               {}
func (*FetchAuthorByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FetchAuthorByIdResponse) GetAuthors() []*Author {
	if m != nil {
		return m.Authors
	}
	return nil
}

type SearchAuthorsRequest struct {
	Offset uint32 `protobuf:"varint,1,opt,name=offset" json:"offset,omitempty"`
	Limit  uint32 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
}

func (m *SearchAuthorsRequest) Reset()                    { *m = SearchAuthorsRequest{} }
func (m *SearchAuthorsRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchAuthorsRequest) ProtoMessage()               {}
func (*SearchAuthorsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SearchAuthorsRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SearchAuthorsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type SearchAuthorsResponse struct {
	Authors []*Author `protobuf:"bytes,1,rep,name=authors" json:"authors,omitempty"`
}

func (m *SearchAuthorsResponse) Reset()                    { *m = SearchAuthorsResponse{} }
func (m *SearchAuthorsResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchAuthorsResponse) ProtoMessage()               {}
func (*SearchAuthorsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *SearchAuthorsResponse) GetAuthors() []*Author {
	if m != nil {
		return m.Authors
	}
	return nil
}

type CreateAuthorRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CreateAuthorRequest) Reset()                    { *m = CreateAuthorRequest{} }
func (m *CreateAuthorRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateAuthorRequest) ProtoMessage()               {}
func (*CreateAuthorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateAuthorRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreateAuthorResponse struct {
	Author *Author `protobuf:"bytes,1,opt,name=author" json:"author,omitempty"`
}

func (m *CreateAuthorResponse) Reset()                    { *m = CreateAuthorResponse{} }
func (m *CreateAuthorResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateAuthorResponse) ProtoMessage()               {}
func (*CreateAuthorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CreateAuthorResponse) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

type UpdateAuthorRequest struct {
	AuthorId uint32 `protobuf:"varint,1,opt,name=authorId" json:"authorId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *UpdateAuthorRequest) Reset()                    { *m = UpdateAuthorRequest{} }
func (m *UpdateAuthorRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateAuthorRequest) ProtoMessage()               {}
func (*UpdateAuthorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateAuthorRequest) GetAuthorId() uint32 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

func (m *UpdateAuthorRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateAuthorResponse struct {
	Author  *Author `protobuf:"bytes,1,opt,name=author" json:"author,omitempty"`
	Success bool    `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *UpdateAuthorResponse) Reset()                    { *m = UpdateAuthorResponse{} }
func (m *UpdateAuthorResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateAuthorResponse) ProtoMessage()               {}
func (*UpdateAuthorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateAuthorResponse) GetAuthor() *Author {
	if m != nil {
		return m.Author
	}
	return nil
}

func (m *UpdateAuthorResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RemoveAuthorRequest struct {
	AuthorId uint32 `protobuf:"varint,1,opt,name=authorId" json:"authorId,omitempty"`
}

func (m *RemoveAuthorRequest) Reset()                    { *m = RemoveAuthorRequest{} }
func (m *RemoveAuthorRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveAuthorRequest) ProtoMessage()               {}
func (*RemoveAuthorRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RemoveAuthorRequest) GetAuthorId() uint32 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

type RemoveAuthorResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *RemoveAuthorResponse) Reset()                    { *m = RemoveAuthorResponse{} }
func (m *RemoveAuthorResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveAuthorResponse) ProtoMessage()               {}
func (*RemoveAuthorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RemoveAuthorResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Author)(nil), "SeppoService.Author")
	proto.RegisterType((*FetchAuthorByIdRequest)(nil), "SeppoService.FetchAuthorByIdRequest")
	proto.RegisterType((*FetchAuthorByIdResponse)(nil), "SeppoService.FetchAuthorByIdResponse")
	proto.RegisterType((*SearchAuthorsRequest)(nil), "SeppoService.SearchAuthorsRequest")
	proto.RegisterType((*SearchAuthorsResponse)(nil), "SeppoService.SearchAuthorsResponse")
	proto.RegisterType((*CreateAuthorRequest)(nil), "SeppoService.CreateAuthorRequest")
	proto.RegisterType((*CreateAuthorResponse)(nil), "SeppoService.CreateAuthorResponse")
	proto.RegisterType((*UpdateAuthorRequest)(nil), "SeppoService.UpdateAuthorRequest")
	proto.RegisterType((*UpdateAuthorResponse)(nil), "SeppoService.UpdateAuthorResponse")
	proto.RegisterType((*RemoveAuthorRequest)(nil), "SeppoService.RemoveAuthorRequest")
	proto.RegisterType((*RemoveAuthorResponse)(nil), "SeppoService.RemoveAuthorResponse")
}

func init() { proto.RegisterFile("author.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcb, 0x4b, 0xc3, 0x30,
	0x1c, 0x26, 0x53, 0xbb, 0xed, 0xe7, 0xe6, 0x21, 0x8b, 0xb3, 0x88, 0x87, 0x91, 0xd3, 0x84, 0x51,
	0x7c, 0x80, 0x77, 0x75, 0x2a, 0xbb, 0xa6, 0x78, 0x15, 0x6a, 0xfb, 0x1b, 0x2b, 0xd8, 0xa5, 0x36,
	0xe9, 0xc0, 0xff, 0x5e, 0xc8, 0x43, 0x5b, 0xa7, 0x20, 0xbb, 0xf5, 0x6b, 0xbe, 0x57, 0xbe, 0x16,
	0x06, 0x49, 0xad, 0x57, 0xb2, 0x8a, 0xca, 0x4a, 0x6a, 0x49, 0x07, 0x31, 0x96, 0xa5, 0x8c, 0xb1,
	0xda, 0xe4, 0x29, 0xf2, 0x19, 0x04, 0xb7, 0xe6, 0x94, 0x1e, 0x41, 0x27, 0xcf, 0x42, 0x32, 0x21,
	0xd3, 0xa1, 0xe8, 0xe4, 0x19, 0xa5, 0xb0, 0xbf, 0x4e, 0x0a, 0x0c, 0x3b, 0x13, 0x32, 0xed, 0x0b,
	0xf3, 0xcc, 0x6f, 0x60, 0xfc, 0x88, 0x3a, 0x5d, 0x59, 0xc9, 0xdd, 0xc7, 0x22, 0x13, 0xf8, 0x5e,
	0xa3, 0xd2, 0xf4, 0x0c, 0xfa, 0x36, 0x65, 0x91, 0xa9, 0x90, 0x4c, 0xf6, 0xa6, 0x43, 0xf1, 0xfd,
	0x82, 0x2f, 0xe0, 0x64, 0x4b, 0xa7, 0x4a, 0xb9, 0x56, 0x48, 0x23, 0xe8, 0x5a, 0x9e, 0x95, 0x1d,
	0x5e, 0xb1, 0xa8, 0x59, 0x30, 0xb2, 0x12, 0xe1, 0x49, 0x7c, 0x0e, 0x2c, 0xc6, 0xa4, 0xf2, 0x5e,
	0xca, 0x17, 0x18, 0x43, 0x20, 0x97, 0x4b, 0x85, 0xda, 0x5d, 0xc1, 0x21, 0xca, 0xe0, 0xe0, 0x2d,
	0x2f, 0x72, 0x6d, 0xee, 0x31, 0x14, 0x16, 0xf0, 0x27, 0x38, 0xfe, 0xe1, 0xb2, 0x63, 0x9d, 0x73,
	0x18, 0xdd, 0x57, 0x98, 0x68, 0x74, 0x07, 0xae, 0x8d, 0x1f, 0x8f, 0x34, 0xc6, 0x9b, 0x03, 0x6b,
	0x53, 0x5d, 0xe4, 0x0c, 0x02, 0xeb, 0x66, 0xd8, 0x7f, 0x25, 0x3a, 0x0e, 0x7f, 0x80, 0xd1, 0x73,
	0x99, 0x6d, 0x05, 0x9e, 0x42, 0xcf, 0xcf, 0xed, 0x06, 0xf8, 0xc2, 0xbf, 0x7e, 0xc9, 0x17, 0x60,
	0x6d, 0x9b, 0x5d, 0xca, 0xd0, 0x10, 0xba, 0xaa, 0x4e, 0x53, 0x54, 0xca, 0x98, 0xf7, 0x84, 0x87,
	0xfc, 0x12, 0x46, 0x02, 0x0b, 0xb9, 0xf9, 0x7f, 0x4d, 0x7e, 0x01, 0xac, 0x2d, 0x71, 0x95, 0x1a,
	0x21, 0xa4, 0x15, 0xf2, 0x1a, 0x98, 0x3f, 0xfa, 0xfa, 0x33, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x91,
	0x26, 0x6b, 0xe1, 0x02, 0x00, 0x00,
}