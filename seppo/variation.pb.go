// Code generated by protoc-gen-go. DO NOT EDIT.
// source: variation.proto

package SeppoService

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Variation struct {
	Id          uint32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	SongId      uint32 `protobuf:"varint,2,opt,name=songId" json:"songId,omitempty"`
	LanguageId  uint32 `protobuf:"varint,3,opt,name=languageId" json:"languageId,omitempty"`
	AuthorId    uint32 `protobuf:"varint,4,opt,name=authorId" json:"authorId,omitempty"`
	CopyrightId uint32 `protobuf:"varint,5,opt,name=copyrightId" json:"copyrightId,omitempty"`
}

func (m *Variation) Reset()                    { *m = Variation{} }
func (m *Variation) String() string            { return proto.CompactTextString(m) }
func (*Variation) ProtoMessage()               {}
func (*Variation) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *Variation) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Variation) GetSongId() uint32 {
	if m != nil {
		return m.SongId
	}
	return 0
}

func (m *Variation) GetLanguageId() uint32 {
	if m != nil {
		return m.LanguageId
	}
	return 0
}

func (m *Variation) GetAuthorId() uint32 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

func (m *Variation) GetCopyrightId() uint32 {
	if m != nil {
		return m.CopyrightId
	}
	return 0
}

type SearchVariationsRequest struct {
	SearchWord           string   `protobuf:"bytes,1,opt,name=searchWord" json:"searchWord,omitempty"`
	SongDatabaseId       uint32   `protobuf:"varint,2,opt,name=songDatabaseId" json:"songDatabaseId,omitempty"`
	SongDatabaseFilterId uint32   `protobuf:"varint,3,opt,name=songDatabaseFilterId" json:"songDatabaseFilterId,omitempty"`
	TagId                uint32   `protobuf:"varint,4,opt,name=tagId" json:"tagId,omitempty"`
	LanguageId           uint32   `protobuf:"varint,5,opt,name=languageId" json:"languageId,omitempty"`
	Offset               uint32   `protobuf:"varint,6,opt,name=offset" json:"offset,omitempty"`
	Limit                uint32   `protobuf:"varint,7,opt,name=limit" json:"limit,omitempty"`
	ScheduleId           uint32   `protobuf:"varint,8,opt,name=scheduleId" json:"scheduleId,omitempty"`
	SkipVariationIds     []uint32 `protobuf:"varint,9,rep,packed,name=skipVariationIds" json:"skipVariationIds,omitempty"`
	OrderBy              uint32   `protobuf:"varint,10,opt,name=orderBy" json:"orderBy,omitempty"`
	SearchFrom           uint32   `protobuf:"varint,11,opt,name=searchFrom" json:"searchFrom,omitempty"`
}

func (m *SearchVariationsRequest) Reset()                    { *m = SearchVariationsRequest{} }
func (m *SearchVariationsRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchVariationsRequest) ProtoMessage()               {}
func (*SearchVariationsRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *SearchVariationsRequest) GetSearchWord() string {
	if m != nil {
		return m.SearchWord
	}
	return ""
}

func (m *SearchVariationsRequest) GetSongDatabaseId() uint32 {
	if m != nil {
		return m.SongDatabaseId
	}
	return 0
}

func (m *SearchVariationsRequest) GetSongDatabaseFilterId() uint32 {
	if m != nil {
		return m.SongDatabaseFilterId
	}
	return 0
}

func (m *SearchVariationsRequest) GetTagId() uint32 {
	if m != nil {
		return m.TagId
	}
	return 0
}

func (m *SearchVariationsRequest) GetLanguageId() uint32 {
	if m != nil {
		return m.LanguageId
	}
	return 0
}

func (m *SearchVariationsRequest) GetOffset() uint32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SearchVariationsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchVariationsRequest) GetScheduleId() uint32 {
	if m != nil {
		return m.ScheduleId
	}
	return 0
}

func (m *SearchVariationsRequest) GetSkipVariationIds() []uint32 {
	if m != nil {
		return m.SkipVariationIds
	}
	return nil
}

func (m *SearchVariationsRequest) GetOrderBy() uint32 {
	if m != nil {
		return m.OrderBy
	}
	return 0
}

func (m *SearchVariationsRequest) GetSearchFrom() uint32 {
	if m != nil {
		return m.SearchFrom
	}
	return 0
}

type SearchVariationsResponse struct {
	Variations    []*Variation `protobuf:"bytes,1,rep,name=variations" json:"variations,omitempty"`
	MaxVariations uint32       `protobuf:"varint,2,opt,name=maxVariations" json:"maxVariations,omitempty"`
}

func (m *SearchVariationsResponse) Reset()                    { *m = SearchVariationsResponse{} }
func (m *SearchVariationsResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchVariationsResponse) ProtoMessage()               {}
func (*SearchVariationsResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *SearchVariationsResponse) GetVariations() []*Variation {
	if m != nil {
		return m.Variations
	}
	return nil
}

func (m *SearchVariationsResponse) GetMaxVariations() uint32 {
	if m != nil {
		return m.MaxVariations
	}
	return 0
}

type FetchVariationByIdRequest struct {
	VariationIds []uint32 `protobuf:"varint,1,rep,packed,name=variationIds" json:"variationIds,omitempty"`
}

func (m *FetchVariationByIdRequest) Reset()                    { *m = FetchVariationByIdRequest{} }
func (m *FetchVariationByIdRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchVariationByIdRequest) ProtoMessage()               {}
func (*FetchVariationByIdRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

func (m *FetchVariationByIdRequest) GetVariationIds() []uint32 {
	if m != nil {
		return m.VariationIds
	}
	return nil
}

type FetchVariationByIdResponse struct {
	Variations []*Variation `protobuf:"bytes,1,rep,name=variations" json:"variations,omitempty"`
}

func (m *FetchVariationByIdResponse) Reset()                    { *m = FetchVariationByIdResponse{} }
func (m *FetchVariationByIdResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchVariationByIdResponse) ProtoMessage()               {}
func (*FetchVariationByIdResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{4} }

func (m *FetchVariationByIdResponse) GetVariations() []*Variation {
	if m != nil {
		return m.Variations
	}
	return nil
}

type CreateVariationRequest struct {
	Name            string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Text            string   `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	TagIds          []uint32 `protobuf:"varint,3,rep,packed,name=tagIds" json:"tagIds,omitempty"`
	SongDatabaseIds []uint32 `protobuf:"varint,4,rep,packed,name=songDatabaseIds" json:"songDatabaseIds,omitempty"`
}

func (m *CreateVariationRequest) Reset()                    { *m = CreateVariationRequest{} }
func (m *CreateVariationRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateVariationRequest) ProtoMessage()               {}
func (*CreateVariationRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{5} }

func (m *CreateVariationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateVariationRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *CreateVariationRequest) GetTagIds() []uint32 {
	if m != nil {
		return m.TagIds
	}
	return nil
}

func (m *CreateVariationRequest) GetSongDatabaseIds() []uint32 {
	if m != nil {
		return m.SongDatabaseIds
	}
	return nil
}

type CreateVariationResponse struct {
	Variation *Variation `protobuf:"bytes,1,opt,name=variation" json:"variation,omitempty"`
}

func (m *CreateVariationResponse) Reset()                    { *m = CreateVariationResponse{} }
func (m *CreateVariationResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateVariationResponse) ProtoMessage()               {}
func (*CreateVariationResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{6} }

func (m *CreateVariationResponse) GetVariation() *Variation {
	if m != nil {
		return m.Variation
	}
	return nil
}

type UpdateVariationRequest struct {
	VariationId           uint32   `protobuf:"varint,1,opt,name=variationId" json:"variationId,omitempty"`
	Name                  string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Text                  string   `protobuf:"bytes,3,opt,name=text" json:"text,omitempty"`
	SongId                uint32   `protobuf:"varint,4,opt,name=songId" json:"songId,omitempty"`
	LanguageId            uint32   `protobuf:"varint,5,opt,name=languageId" json:"languageId,omitempty"`
	AddTagIds             []uint32 `protobuf:"varint,6,rep,packed,name=addTagIds" json:"addTagIds,omitempty"`
	RemoveTagIds          []uint32 `protobuf:"varint,7,rep,packed,name=removeTagIds" json:"removeTagIds,omitempty"`
	AddSongDatabaseIds    []uint32 `protobuf:"varint,8,rep,packed,name=addSongDatabaseIds" json:"addSongDatabaseIds,omitempty"`
	RemoveSongDatabaseIds []uint32 `protobuf:"varint,9,rep,packed,name=removeSongDatabaseIds" json:"removeSongDatabaseIds,omitempty"`
	AuthorId              uint32   `protobuf:"varint,10,opt,name=authorId" json:"authorId,omitempty"`
}

func (m *UpdateVariationRequest) Reset()                    { *m = UpdateVariationRequest{} }
func (m *UpdateVariationRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateVariationRequest) ProtoMessage()               {}
func (*UpdateVariationRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{7} }

func (m *UpdateVariationRequest) GetVariationId() uint32 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

func (m *UpdateVariationRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateVariationRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *UpdateVariationRequest) GetSongId() uint32 {
	if m != nil {
		return m.SongId
	}
	return 0
}

func (m *UpdateVariationRequest) GetLanguageId() uint32 {
	if m != nil {
		return m.LanguageId
	}
	return 0
}

func (m *UpdateVariationRequest) GetAddTagIds() []uint32 {
	if m != nil {
		return m.AddTagIds
	}
	return nil
}

func (m *UpdateVariationRequest) GetRemoveTagIds() []uint32 {
	if m != nil {
		return m.RemoveTagIds
	}
	return nil
}

func (m *UpdateVariationRequest) GetAddSongDatabaseIds() []uint32 {
	if m != nil {
		return m.AddSongDatabaseIds
	}
	return nil
}

func (m *UpdateVariationRequest) GetRemoveSongDatabaseIds() []uint32 {
	if m != nil {
		return m.RemoveSongDatabaseIds
	}
	return nil
}

func (m *UpdateVariationRequest) GetAuthorId() uint32 {
	if m != nil {
		return m.AuthorId
	}
	return 0
}

type UpdateVariationResponse struct {
	Variation *Variation `protobuf:"bytes,1,opt,name=variation" json:"variation,omitempty"`
	Success   bool       `protobuf:"varint,2,opt,name=success" json:"success,omitempty"`
}

func (m *UpdateVariationResponse) Reset()                    { *m = UpdateVariationResponse{} }
func (m *UpdateVariationResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateVariationResponse) ProtoMessage()               {}
func (*UpdateVariationResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{8} }

func (m *UpdateVariationResponse) GetVariation() *Variation {
	if m != nil {
		return m.Variation
	}
	return nil
}

func (m *UpdateVariationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

type RemoveVariationRequest struct {
	VariationId uint32 `protobuf:"varint,1,opt,name=variationId" json:"variationId,omitempty"`
}

func (m *RemoveVariationRequest) Reset()                    { *m = RemoveVariationRequest{} }
func (m *RemoveVariationRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveVariationRequest) ProtoMessage()               {}
func (*RemoveVariationRequest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{9} }

func (m *RemoveVariationRequest) GetVariationId() uint32 {
	if m != nil {
		return m.VariationId
	}
	return 0
}

type RemoveVariationResponse struct {
	Success bool `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
}

func (m *RemoveVariationResponse) Reset()                    { *m = RemoveVariationResponse{} }
func (m *RemoveVariationResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveVariationResponse) ProtoMessage()               {}
func (*RemoveVariationResponse) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{10} }

func (m *RemoveVariationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*Variation)(nil), "SeppoService.Variation")
	proto.RegisterType((*SearchVariationsRequest)(nil), "SeppoService.SearchVariationsRequest")
	proto.RegisterType((*SearchVariationsResponse)(nil), "SeppoService.SearchVariationsResponse")
	proto.RegisterType((*FetchVariationByIdRequest)(nil), "SeppoService.FetchVariationByIdRequest")
	proto.RegisterType((*FetchVariationByIdResponse)(nil), "SeppoService.FetchVariationByIdResponse")
	proto.RegisterType((*CreateVariationRequest)(nil), "SeppoService.CreateVariationRequest")
	proto.RegisterType((*CreateVariationResponse)(nil), "SeppoService.CreateVariationResponse")
	proto.RegisterType((*UpdateVariationRequest)(nil), "SeppoService.UpdateVariationRequest")
	proto.RegisterType((*UpdateVariationResponse)(nil), "SeppoService.UpdateVariationResponse")
	proto.RegisterType((*RemoveVariationRequest)(nil), "SeppoService.RemoveVariationRequest")
	proto.RegisterType((*RemoveVariationResponse)(nil), "SeppoService.RemoveVariationResponse")
}

func init() { proto.RegisterFile("variation.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0xd2, 0xae, 0x6b, 0x5e, 0xd7, 0x0d, 0x59, 0xa3, 0x35, 0x13, 0x42, 0x95, 0x85, 0x50,
	0xc5, 0xa1, 0x87, 0x0d, 0x84, 0xc4, 0x05, 0x69, 0xa0, 0x4a, 0xb9, 0xa1, 0x94, 0xc1, 0xd9, 0x8b,
	0xbd, 0x36, 0xd0, 0xd6, 0xc1, 0x76, 0xab, 0xf5, 0xca, 0x3f, 0xe0, 0x0f, 0xf2, 0x57, 0x40, 0x71,
	0x52, 0xc7, 0x69, 0xa2, 0x49, 0x68, 0xb7, 0xbc, 0xef, 0x3d, 0x3f, 0x7f, 0xdf, 0xe7, 0x67, 0x07,
	0xce, 0xb6, 0x54, 0x26, 0x54, 0x27, 0x62, 0x3d, 0x49, 0xa5, 0xd0, 0x02, 0x9d, 0xcc, 0x78, 0x9a,
	0x8a, 0x19, 0x97, 0xdb, 0x24, 0xe6, 0xe4, 0xb7, 0x07, 0xc1, 0xd7, 0x7d, 0x05, 0x3a, 0x05, 0x3f,
	0x61, 0xd8, 0x1b, 0x79, 0xe3, 0x7e, 0xe4, 0x27, 0x0c, 0x0d, 0xa0, 0xa3, 0xc4, 0x7a, 0x1e, 0x32,
	0xec, 0x1b, 0xac, 0x88, 0xd0, 0x0b, 0x80, 0x25, 0x5d, 0xcf, 0x37, 0x74, 0xce, 0x43, 0x86, 0x5b,
	0x26, 0xe7, 0x20, 0xe8, 0x02, 0xba, 0x74, 0xa3, 0x17, 0x42, 0x86, 0x0c, 0xb7, 0x4d, 0xd6, 0xc6,
	0x68, 0x04, 0xbd, 0x58, 0xa4, 0x3b, 0x99, 0xcc, 0x17, 0x3a, 0x64, 0xf8, 0xc8, 0xa4, 0x5d, 0x88,
	0xfc, 0xf5, 0x61, 0x38, 0xe3, 0x54, 0xc6, 0x0b, 0xcb, 0x4c, 0x45, 0xfc, 0xe7, 0x86, 0x2b, 0x9d,
	0xed, 0xac, 0x4c, 0xea, 0x9b, 0x90, 0x39, 0xd3, 0x20, 0x72, 0x10, 0xf4, 0x0a, 0x4e, 0x33, 0x8e,
	0x9f, 0xa8, 0xa6, 0xb7, 0x54, 0x71, 0xcb, 0xfc, 0x00, 0x45, 0x97, 0x70, 0xee, 0x22, 0xd3, 0x64,
	0xa9, 0xb9, 0xb4, 0x5a, 0x1a, 0x73, 0xe8, 0x1c, 0x8e, 0x34, 0x9d, 0x5b, 0x49, 0x79, 0x70, 0xe0,
	0xc5, 0x51, 0xcd, 0x8b, 0x01, 0x74, 0xc4, 0xdd, 0x9d, 0xe2, 0x1a, 0x77, 0x72, 0x0f, 0xf3, 0x28,
	0xeb, 0xb6, 0x4c, 0x56, 0x89, 0xc6, 0xc7, 0x79, 0x37, 0x13, 0x18, 0x7d, 0xf1, 0x82, 0xb3, 0xcd,
	0x32, 0xeb, 0xd6, 0xcd, 0xbb, 0x95, 0x08, 0x7a, 0x0d, 0x4f, 0xd4, 0x8f, 0x24, 0xb5, 0xc6, 0x84,
	0x4c, 0xe1, 0x60, 0xd4, 0x1a, 0xf7, 0xa3, 0x1a, 0x8e, 0x30, 0x1c, 0x0b, 0xc9, 0xb8, 0xbc, 0xde,
	0x61, 0x30, 0x8d, 0xf6, 0x61, 0xe9, 0xe2, 0x54, 0x8a, 0x15, 0xee, 0x15, 0xbb, 0x58, 0x84, 0xec,
	0x00, 0xd7, 0x0f, 0x40, 0xa5, 0x62, 0xad, 0x38, 0x7a, 0x07, 0x60, 0x47, 0x4a, 0x61, 0x6f, 0xd4,
	0x1a, 0xf7, 0x2e, 0x87, 0x13, 0x77, 0xa8, 0x26, 0x76, 0x55, 0xe4, 0x94, 0xa2, 0x97, 0xd0, 0x5f,
	0xd1, 0xfb, 0xb2, 0x63, 0x71, 0x32, 0x55, 0x90, 0x7c, 0x80, 0x67, 0x53, 0xae, 0x9d, 0x9d, 0xaf,
	0x77, 0x21, 0xdb, 0x9f, 0x3e, 0x81, 0x93, 0xad, 0xab, 0xdc, 0x33, 0xca, 0x2b, 0x18, 0xb9, 0x81,
	0x8b, 0xa6, 0x06, 0x8f, 0x64, 0x4f, 0x7e, 0x79, 0x30, 0xf8, 0x28, 0x39, 0xd5, 0xbc, 0xcc, 0x17,
	0xac, 0x10, 0xb4, 0xd7, 0x74, 0xc5, 0x8b, 0x69, 0x34, 0xdf, 0x19, 0xa6, 0xf9, 0xbd, 0x36, 0x1a,
	0x83, 0xc8, 0x7c, 0x67, 0x93, 0x60, 0x46, 0x46, 0xe1, 0x96, 0xe1, 0x5d, 0x44, 0x68, 0x0c, 0x67,
	0xd5, 0xe9, 0x54, 0xb8, 0x6d, 0x0a, 0x0e, 0x61, 0xf2, 0x19, 0x86, 0x35, 0x0e, 0x85, 0xb0, 0xb7,
	0x10, 0x58, 0xb6, 0x86, 0xc9, 0x03, 0xba, 0xca, 0x4a, 0xf2, 0xc7, 0x87, 0xc1, 0x4d, 0xca, 0x9a,
	0x64, 0x8d, 0xa0, 0xe7, 0x18, 0x5b, 0xbc, 0x0a, 0x2e, 0x64, 0x85, 0xfb, 0x0d, 0xc2, 0x5b, 0x55,
	0xe1, 0xc5, 0x33, 0xd2, 0x7e, 0xe0, 0x19, 0xa9, 0x5f, 0x9d, 0xe7, 0x10, 0x50, 0xc6, 0xbe, 0xe4,
	0x9e, 0x75, 0x8c, 0x25, 0x25, 0x90, 0x0d, 0x83, 0xe4, 0x2b, 0xb1, 0xe5, 0x45, 0xc1, 0x71, 0x3e,
	0x0c, 0x2e, 0x86, 0x26, 0x80, 0x28, 0x63, 0xb3, 0x03, 0x77, 0xbb, 0xa6, 0xb2, 0x21, 0x83, 0xde,
	0xc0, 0xd3, 0x7c, 0xfd, 0xe1, 0x92, 0xfc, 0x8e, 0x35, 0x27, 0x2b, 0xcf, 0x1d, 0x54, 0x9f, 0x3b,
	0xf2, 0x1d, 0x86, 0x35, 0x7f, 0x1f, 0x75, 0x64, 0xd9, 0xb5, 0x56, 0x9b, 0x38, 0xe6, 0x2a, 0xbf,
	0x41, 0xdd, 0x68, 0x1f, 0x92, 0xf7, 0x30, 0x88, 0x0c, 0xc1, 0xff, 0x3f, 0x4b, 0x72, 0x05, 0xc3,
	0xda, 0xda, 0x82, 0xa7, 0xb3, 0xa1, 0x57, 0xd9, 0xf0, 0xb6, 0x63, 0x7e, 0x29, 0x57, 0xff, 0x02,
	0x00, 0x00, 0xff, 0xff, 0xa1, 0xea, 0x2b, 0xcf, 0x65, 0x06, 0x00, 0x00,
}
