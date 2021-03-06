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
	AuthorId             uint32   `protobuf:"varint,12,opt,name=authorId" json:"authorId,omitempty"`
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

func (m *SearchVariationsRequest) GetAuthorId() uint32 {
	if m != nil {
		return m.AuthorId
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
	// 610 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x8e, 0xd3, 0x30,
	0x10, 0x56, 0x9a, 0x6e, 0xdb, 0x4c, 0xdb, 0x5d, 0x64, 0x2d, 0xad, 0x59, 0x21, 0x54, 0x45, 0x08,
	0x55, 0x1c, 0x7a, 0xd8, 0x05, 0x21, 0x71, 0x41, 0x5a, 0x50, 0xa5, 0xdc, 0x50, 0xca, 0xc2, 0xd9,
	0x1b, 0x7b, 0xdb, 0x40, 0x5b, 0x07, 0xdb, 0xad, 0xb6, 0x57, 0xde, 0x80, 0x33, 0xef, 0xc6, 0xb3,
	0xa0, 0x38, 0xa9, 0xe3, 0xfc, 0x68, 0x25, 0xb4, 0xb7, 0xcc, 0x37, 0xe3, 0xf1, 0x37, 0xdf, 0xcc,
	0x38, 0x70, 0xb6, 0x27, 0x22, 0x26, 0x2a, 0xe6, 0xdb, 0x59, 0x22, 0xb8, 0xe2, 0x68, 0xb0, 0x60,
	0x49, 0xc2, 0x17, 0x4c, 0xec, 0xe3, 0x88, 0xf9, 0xbf, 0x1d, 0xf0, 0xbe, 0x1e, 0x23, 0xd0, 0x29,
	0xb4, 0x62, 0x8a, 0x9d, 0x89, 0x33, 0x1d, 0x86, 0xad, 0x98, 0xa2, 0x11, 0x74, 0x24, 0xdf, 0x2e,
	0x03, 0x8a, 0x5b, 0x1a, 0xcb, 0x2d, 0xf4, 0x02, 0x60, 0x4d, 0xb6, 0xcb, 0x1d, 0x59, 0xb2, 0x80,
	0x62, 0x57, 0xfb, 0x2c, 0x04, 0x5d, 0x40, 0x8f, 0xec, 0xd4, 0x8a, 0x8b, 0x80, 0xe2, 0xb6, 0xf6,
	0x1a, 0x1b, 0x4d, 0xa0, 0x1f, 0xf1, 0xe4, 0x20, 0xe2, 0xe5, 0x4a, 0x05, 0x14, 0x9f, 0x68, 0xb7,
	0x0d, 0xf9, 0x7f, 0x5c, 0x18, 0x2f, 0x18, 0x11, 0xd1, 0xca, 0x30, 0x93, 0x21, 0xfb, 0xb9, 0x63,
	0x52, 0xa5, 0x37, 0x4b, 0xed, 0xfa, 0xc6, 0x45, 0xc6, 0xd4, 0x0b, 0x2d, 0x04, 0xbd, 0x82, 0xd3,
	0x94, 0xe3, 0x27, 0xa2, 0xc8, 0x2d, 0x91, 0xcc, 0x30, 0xaf, 0xa0, 0xe8, 0x12, 0xce, 0x6d, 0x64,
	0x1e, 0xaf, 0x15, 0x13, 0xa6, 0x96, 0x46, 0x1f, 0x3a, 0x87, 0x13, 0x45, 0x96, 0xa6, 0xa4, 0xcc,
	0xa8, 0x68, 0x71, 0x52, 0xd3, 0x62, 0x04, 0x1d, 0x7e, 0x77, 0x27, 0x99, 0xc2, 0x9d, 0x4c, 0xc3,
	0xcc, 0x4a, 0xb3, 0xad, 0xe3, 0x4d, 0xac, 0x70, 0x37, 0xcb, 0xa6, 0x0d, 0x5d, 0x5f, 0xb4, 0x62,
	0x74, 0xb7, 0x4e, 0xb3, 0xf5, 0xb2, 0x6c, 0x05, 0x82, 0x5e, 0xc3, 0x13, 0xf9, 0x23, 0x4e, 0x8c,
	0x30, 0x01, 0x95, 0xd8, 0x9b, 0xb8, 0xd3, 0x61, 0x58, 0xc3, 0x11, 0x86, 0x2e, 0x17, 0x94, 0x89,
	0xeb, 0x03, 0x06, 0x9d, 0xe8, 0x68, 0x16, 0x2a, 0xce, 0x05, 0xdf, 0xe0, 0x7e, 0x7e, 0x8b, 0x41,
	0x4a, 0xfd, 0x1b, 0x94, 0xfb, 0xe7, 0x1f, 0x00, 0xd7, 0x9b, 0x23, 0x13, 0xbe, 0x95, 0x0c, 0xbd,
	0x03, 0x30, 0xe3, 0x26, 0xb1, 0x33, 0x71, 0xa7, 0xfd, 0xcb, 0xf1, 0xcc, 0x1e, 0xb8, 0x99, 0x39,
	0x15, 0x5a, 0xa1, 0xe8, 0x25, 0x0c, 0x37, 0xe4, 0xbe, 0xc8, 0x98, 0x77, 0xad, 0x0c, 0xfa, 0x1f,
	0xe0, 0xd9, 0x9c, 0x29, 0xeb, 0xe6, 0xeb, 0x43, 0x40, 0x8f, 0x93, 0xe1, 0xc3, 0x60, 0x6f, 0xab,
	0xe2, 0x68, 0x55, 0x4a, 0x98, 0x7f, 0x03, 0x17, 0x4d, 0x09, 0x1e, 0xc9, 0xde, 0xff, 0xe5, 0xc0,
	0xe8, 0xa3, 0x60, 0x44, 0xb1, 0xc2, 0x9f, 0xb3, 0x42, 0xd0, 0xde, 0x92, 0x0d, 0xcb, 0x27, 0x55,
	0x7f, 0xa7, 0x98, 0x62, 0xf7, 0x4a, 0xd7, 0xe8, 0x85, 0xfa, 0x3b, 0x9d, 0x12, 0x3d, 0x4e, 0x12,
	0xbb, 0x9a, 0x77, 0x6e, 0xa1, 0x29, 0x9c, 0x95, 0x27, 0x57, 0xe2, 0xb6, 0x0e, 0xa8, 0xc2, 0xfe,
	0x67, 0x18, 0xd7, 0x38, 0xe4, 0x85, 0xbd, 0x05, 0xcf, 0xb0, 0xd5, 0x4c, 0x1e, 0xa8, 0xab, 0x88,
	0xf4, 0xff, 0xb6, 0x60, 0x74, 0x93, 0xd0, 0xa6, 0xb2, 0x26, 0xd0, 0xb7, 0x84, 0xcd, 0x5f, 0x0c,
	0x1b, 0x32, 0x85, 0xb7, 0x1a, 0x0a, 0x77, 0xcb, 0x85, 0xe7, 0x4f, 0x4c, 0xfb, 0x81, 0x27, 0xa6,
	0xbe, 0x56, 0xcf, 0xc1, 0x23, 0x94, 0x7e, 0xc9, 0x34, 0xeb, 0x68, 0x49, 0x0a, 0x20, 0x1d, 0x06,
	0xc1, 0x36, 0x7c, 0xcf, 0xf2, 0x80, 0x6e, 0x36, 0x0c, 0x36, 0x86, 0x66, 0x80, 0x08, 0xa5, 0x8b,
	0x8a, 0xba, 0x3d, 0x1d, 0xd9, 0xe0, 0x41, 0x6f, 0xe0, 0x69, 0x76, 0xbe, 0x7a, 0x24, 0xdb, 0xbf,
	0x66, 0x67, 0x69, 0x95, 0xa0, 0xb2, 0x4a, 0xdf, 0x61, 0x5c, 0xd3, 0xf7, 0x51, 0x2d, 0x4b, 0x57,
	0x5e, 0xee, 0xa2, 0x88, 0xc9, 0x6c, 0x83, 0x7a, 0xe1, 0xd1, 0xf4, 0xdf, 0xc3, 0x28, 0xd4, 0x04,
	0xff, 0xbf, 0x97, 0xfe, 0x15, 0x8c, 0x6b, 0x67, 0x73, 0x9e, 0xd6, 0x85, 0x4e, 0xe9, 0xc2, 0xdb,
	0x8e, 0xfe, 0xdd, 0x5c, 0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x25, 0x9e, 0x04, 0xcb, 0x81, 0x06,
	0x00, 0x00,
}
