// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: envoy/type/http_status.proto

package envoy_type

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/lyft/protoc-gen-validate/validate"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// HTTP response codes supported in Envoy.
// For more details: http://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml
type StatusCode int32

const (
	// Empty - This code not part of the HTTP status code specification, but it is needed for proto
	// `enum` type.
	StatusCode_Empty                         StatusCode = 0
	StatusCode_Continue                      StatusCode = 100
	StatusCode_OK                            StatusCode = 200
	StatusCode_Created                       StatusCode = 201
	StatusCode_Accepted                      StatusCode = 202
	StatusCode_NonAuthoritativeInformation   StatusCode = 203
	StatusCode_NoContent                     StatusCode = 204
	StatusCode_ResetContent                  StatusCode = 205
	StatusCode_PartialContent                StatusCode = 206
	StatusCode_MultiStatus                   StatusCode = 207
	StatusCode_AlreadyReported               StatusCode = 208
	StatusCode_IMUsed                        StatusCode = 226
	StatusCode_MultipleChoices               StatusCode = 300
	StatusCode_MovedPermanently              StatusCode = 301
	StatusCode_Found                         StatusCode = 302
	StatusCode_SeeOther                      StatusCode = 303
	StatusCode_NotModified                   StatusCode = 304
	StatusCode_UseProxy                      StatusCode = 305
	StatusCode_TemporaryRedirect             StatusCode = 307
	StatusCode_PermanentRedirect             StatusCode = 308
	StatusCode_BadRequest                    StatusCode = 400
	StatusCode_Unauthorized                  StatusCode = 401
	StatusCode_PaymentRequired               StatusCode = 402
	StatusCode_Forbidden                     StatusCode = 403
	StatusCode_NotFound                      StatusCode = 404
	StatusCode_MethodNotAllowed              StatusCode = 405
	StatusCode_NotAcceptable                 StatusCode = 406
	StatusCode_ProxyAuthenticationRequired   StatusCode = 407
	StatusCode_RequestTimeout                StatusCode = 408
	StatusCode_Conflict                      StatusCode = 409
	StatusCode_Gone                          StatusCode = 410
	StatusCode_LengthRequired                StatusCode = 411
	StatusCode_PreconditionFailed            StatusCode = 412
	StatusCode_PayloadTooLarge               StatusCode = 413
	StatusCode_URITooLong                    StatusCode = 414
	StatusCode_UnsupportedMediaType          StatusCode = 415
	StatusCode_RangeNotSatisfiable           StatusCode = 416
	StatusCode_ExpectationFailed             StatusCode = 417
	StatusCode_MisdirectedRequest            StatusCode = 421
	StatusCode_UnprocessableEntity           StatusCode = 422
	StatusCode_Locked                        StatusCode = 423
	StatusCode_FailedDependency              StatusCode = 424
	StatusCode_UpgradeRequired               StatusCode = 426
	StatusCode_PreconditionRequired          StatusCode = 428
	StatusCode_TooManyRequests               StatusCode = 429
	StatusCode_RequestHeaderFieldsTooLarge   StatusCode = 431
	StatusCode_InternalServerError           StatusCode = 500
	StatusCode_NotImplemented                StatusCode = 501
	StatusCode_BadGateway                    StatusCode = 502
	StatusCode_ServiceUnavailable            StatusCode = 503
	StatusCode_GatewayTimeout                StatusCode = 504
	StatusCode_HTTPVersionNotSupported       StatusCode = 505
	StatusCode_VariantAlsoNegotiates         StatusCode = 506
	StatusCode_InsufficientStorage           StatusCode = 507
	StatusCode_LoopDetected                  StatusCode = 508
	StatusCode_NotExtended                   StatusCode = 510
	StatusCode_NetworkAuthenticationRequired StatusCode = 511
)

var StatusCode_name = map[int32]string{
	0:   "Empty",
	100: "Continue",
	200: "OK",
	201: "Created",
	202: "Accepted",
	203: "NonAuthoritativeInformation",
	204: "NoContent",
	205: "ResetContent",
	206: "PartialContent",
	207: "MultiStatus",
	208: "AlreadyReported",
	226: "IMUsed",
	300: "MultipleChoices",
	301: "MovedPermanently",
	302: "Found",
	303: "SeeOther",
	304: "NotModified",
	305: "UseProxy",
	307: "TemporaryRedirect",
	308: "PermanentRedirect",
	400: "BadRequest",
	401: "Unauthorized",
	402: "PaymentRequired",
	403: "Forbidden",
	404: "NotFound",
	405: "MethodNotAllowed",
	406: "NotAcceptable",
	407: "ProxyAuthenticationRequired",
	408: "RequestTimeout",
	409: "Conflict",
	410: "Gone",
	411: "LengthRequired",
	412: "PreconditionFailed",
	413: "PayloadTooLarge",
	414: "URITooLong",
	415: "UnsupportedMediaType",
	416: "RangeNotSatisfiable",
	417: "ExpectationFailed",
	421: "MisdirectedRequest",
	422: "UnprocessableEntity",
	423: "Locked",
	424: "FailedDependency",
	426: "UpgradeRequired",
	428: "PreconditionRequired",
	429: "TooManyRequests",
	431: "RequestHeaderFieldsTooLarge",
	500: "InternalServerError",
	501: "NotImplemented",
	502: "BadGateway",
	503: "ServiceUnavailable",
	504: "GatewayTimeout",
	505: "HTTPVersionNotSupported",
	506: "VariantAlsoNegotiates",
	507: "InsufficientStorage",
	508: "LoopDetected",
	510: "NotExtended",
	511: "NetworkAuthenticationRequired",
}
var StatusCode_value = map[string]int32{
	"Empty":                         0,
	"Continue":                      100,
	"OK":                            200,
	"Created":                       201,
	"Accepted":                      202,
	"NonAuthoritativeInformation":   203,
	"NoContent":                     204,
	"ResetContent":                  205,
	"PartialContent":                206,
	"MultiStatus":                   207,
	"AlreadyReported":               208,
	"IMUsed":                        226,
	"MultipleChoices":               300,
	"MovedPermanently":              301,
	"Found":                         302,
	"SeeOther":                      303,
	"NotModified":                   304,
	"UseProxy":                      305,
	"TemporaryRedirect":             307,
	"PermanentRedirect":             308,
	"BadRequest":                    400,
	"Unauthorized":                  401,
	"PaymentRequired":               402,
	"Forbidden":                     403,
	"NotFound":                      404,
	"MethodNotAllowed":              405,
	"NotAcceptable":                 406,
	"ProxyAuthenticationRequired":   407,
	"RequestTimeout":                408,
	"Conflict":                      409,
	"Gone":                          410,
	"LengthRequired":                411,
	"PreconditionFailed":            412,
	"PayloadTooLarge":               413,
	"URITooLong":                    414,
	"UnsupportedMediaType":          415,
	"RangeNotSatisfiable":           416,
	"ExpectationFailed":             417,
	"MisdirectedRequest":            421,
	"UnprocessableEntity":           422,
	"Locked":                        423,
	"FailedDependency":              424,
	"UpgradeRequired":               426,
	"PreconditionRequired":          428,
	"TooManyRequests":               429,
	"RequestHeaderFieldsTooLarge":   431,
	"InternalServerError":           500,
	"NotImplemented":                501,
	"BadGateway":                    502,
	"ServiceUnavailable":            503,
	"GatewayTimeout":                504,
	"HTTPVersionNotSupported":       505,
	"VariantAlsoNegotiates":         506,
	"InsufficientStorage":           507,
	"LoopDetected":                  508,
	"NotExtended":                   510,
	"NetworkAuthenticationRequired": 511,
}

func (x StatusCode) String() string {
	return proto.EnumName(StatusCode_name, int32(x))
}
func (StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_http_status_e053fdba6856bd82, []int{0}
}

// HTTP status.
type HttpStatus struct {
	// Supplies HTTP response code.
	Code                 StatusCode `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.type.StatusCode" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *HttpStatus) Reset()         { *m = HttpStatus{} }
func (m *HttpStatus) String() string { return proto.CompactTextString(m) }
func (*HttpStatus) ProtoMessage()    {}
func (*HttpStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_http_status_e053fdba6856bd82, []int{0}
}
func (m *HttpStatus) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *HttpStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_HttpStatus.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *HttpStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpStatus.Merge(dst, src)
}
func (m *HttpStatus) XXX_Size() int {
	return m.Size()
}
func (m *HttpStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpStatus.DiscardUnknown(m)
}

var xxx_messageInfo_HttpStatus proto.InternalMessageInfo

func (m *HttpStatus) GetCode() StatusCode {
	if m != nil {
		return m.Code
	}
	return StatusCode_Empty
}

func init() {
	proto.RegisterType((*HttpStatus)(nil), "envoy.type.HttpStatus")
	proto.RegisterEnum("envoy.type.StatusCode", StatusCode_name, StatusCode_value)
}
func (m *HttpStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HttpStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Code != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHttpStatus(dAtA, i, uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintHttpStatus(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HttpStatus) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovHttpStatus(uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovHttpStatus(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHttpStatus(x uint64) (n int) {
	return sovHttpStatus(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HttpStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHttpStatus
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HttpStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HttpStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHttpStatus
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= (StatusCode(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHttpStatus(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHttpStatus
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHttpStatus(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHttpStatus
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHttpStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHttpStatus
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHttpStatus
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHttpStatus
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHttpStatus(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHttpStatus = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHttpStatus   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("envoy/type/http_status.proto", fileDescriptor_http_status_e053fdba6856bd82)
}

var fileDescriptor_http_status_e053fdba6856bd82 = []byte{
	// 922 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x49, 0x6f, 0x1c, 0x45,
	0x14, 0x4e, 0x4f, 0x65, 0x73, 0x25, 0x71, 0x5e, 0x2a, 0x4e, 0x6c, 0x42, 0xb0, 0xac, 0x1c, 0x10,
	0xe2, 0x60, 0x4b, 0xf0, 0x07, 0xb0, 0x1d, 0x3b, 0xb6, 0xf0, 0x4c, 0x46, 0xe3, 0x99, 0x5c, 0x51,
	0xb9, 0xeb, 0xcd, 0x4c, 0x29, 0x3d, 0xf5, 0x3a, 0xd5, 0x6f, 0xc6, 0x6e, 0x8e, 0xfc, 0x02, 0xf6,
	0x7d, 0x3d, 0xb0, 0x08, 0x25, 0x04, 0x04, 0x5c, 0x38, 0x71, 0x0c, 0x7b, 0x7e, 0x02, 0xf2, 0x8d,
	0x3b, 0x3b, 0x08, 0x50, 0xd5, 0x2c, 0xf1, 0x85, 0x5b, 0xd7, 0xab, 0xb7, 0x7c, 0xcb, 0xeb, 0x92,
	0x17, 0xd1, 0x0d, 0xa8, 0x5c, 0xe2, 0x32, 0xc7, 0xa5, 0x2e, 0x73, 0xfe, 0x44, 0xc1, 0x9a, 0xfb,
	0xc5, 0x62, 0xee, 0x89, 0x49, 0xc9, 0x78, 0xbb, 0x18, 0x6e, 0x2f, 0xcc, 0x0e, 0x74, 0x66, 0x8d,
	0x66, 0x5c, 0x1a, 0x7f, 0x0c, 0x93, 0x2e, 0xd5, 0xa4, 0xdc, 0x60, 0xce, 0xb7, 0x63, 0xa1, 0x7a,
	0x4c, 0x1e, 0x4e, 0xc9, 0xe0, 0x5c, 0xb2, 0x90, 0x3c, 0x34, 0xfd, 0xc8, 0xf9, 0xc5, 0x7b, 0x1d,
	0x16, 0x87, 0x19, 0xab, 0x64, 0x70, 0x65, 0xe6, 0x8b, 0x9f, 0xbe, 0x14, 0x47, 0x9e, 0x4a, 0x2a,
	0x0b, 0x87, 0xc6, 0x5f, 0x90, 0x34, 0x62, 0xe5, 0xc3, 0x9f, 0x4f, 0x49, 0x79, 0x2f, 0x55, 0x4d,
	0xc9, 0x23, 0x6b, 0xbd, 0x9c, 0x4b, 0x38, 0xa4, 0x4e, 0xca, 0xe3, 0xab, 0xe4, 0xd8, 0xba, 0x3e,
	0x82, 0x51, 0xc7, 0x64, 0xe5, 0xea, 0xe3, 0x70, 0x27, 0x51, 0x27, 0xe5, 0xb1, 0x55, 0x8f, 0x9a,
	0xd1, 0xc0, 0x57, 0x89, 0x3a, 0x25, 0x8f, 0x2f, 0xa7, 0x29, 0xe6, 0xe1, 0xf8, 0x75, 0xa2, 0x16,
	0xe4, 0xfd, 0x35, 0x72, 0xcb, 0x7d, 0xee, 0x92, 0xb7, 0xac, 0xd9, 0x0e, 0x70, 0xd3, 0xb5, 0xc9,
	0xf7, 0x34, 0x5b, 0x72, 0xf0, 0x4d, 0xa2, 0xa6, 0xe5, 0x54, 0x8d, 0x42, 0x5f, 0x74, 0x0c, 0xdf,
	0x26, 0xea, 0x8c, 0x3c, 0xd9, 0xc0, 0x02, 0x79, 0x1c, 0xfa, 0x2e, 0x51, 0x67, 0xe5, 0x74, 0x5d,
	0x7b, 0xb6, 0x3a, 0x1b, 0x07, 0xbf, 0x4f, 0x14, 0xc8, 0x13, 0xd5, 0x7e, 0xc6, 0x76, 0x88, 0x15,
	0x7e, 0x48, 0xd4, 0x8c, 0x3c, 0xbd, 0x9c, 0x79, 0xd4, 0xa6, 0x6c, 0x60, 0x4e, 0x3e, 0x20, 0xb8,
	0x9b, 0xa8, 0x13, 0xf2, 0xe8, 0x66, 0xb5, 0x55, 0xa0, 0x81, 0xfd, 0x98, 0x12, 0x8b, 0xf2, 0x0c,
	0x57, 0xbb, 0x64, 0x53, 0x2c, 0xe0, 0x66, 0x45, 0x9d, 0x93, 0x50, 0xa5, 0x01, 0x9a, 0x3a, 0xfa,
	0x9e, 0x76, 0xe8, 0x38, 0x2b, 0xe1, 0x56, 0x45, 0x49, 0x79, 0x64, 0x9d, 0xfa, 0xce, 0xc0, 0x47,
	0x95, 0x40, 0x6b, 0x1b, 0xf1, 0x2a, 0x77, 0xd1, 0xc3, 0xed, 0x4a, 0x18, 0x5e, 0x23, 0xae, 0x92,
	0xb1, 0x6d, 0x8b, 0x06, 0x3e, 0x8e, 0x09, 0xad, 0x02, 0xeb, 0x9e, 0xf6, 0x4a, 0xf8, 0xa4, 0xa2,
	0xce, 0xcb, 0x33, 0x4d, 0xec, 0xe5, 0xe4, 0xb5, 0x2f, 0x1b, 0x68, 0xac, 0xc7, 0x94, 0xe1, 0xd3,
	0x18, 0x9f, 0x4c, 0x99, 0xc4, 0x3f, 0xab, 0xa8, 0xd3, 0x52, 0xae, 0x68, 0xd3, 0xc0, 0x1b, 0x7d,
	0x2c, 0x18, 0x9e, 0x16, 0x41, 0x86, 0x96, 0xd3, 0x43, 0xdd, 0x9e, 0x44, 0x03, 0xcf, 0x88, 0x00,
	0xbe, 0xae, 0xcb, 0x5e, 0xac, 0xbc, 0xd1, 0xb7, 0x1e, 0x0d, 0x3c, 0x2b, 0x82, 0x7e, 0xeb, 0xe4,
	0x77, 0xac, 0x31, 0xe8, 0xe0, 0x39, 0x11, 0x80, 0xd4, 0x88, 0x87, 0xc0, 0x9f, 0x17, 0x91, 0x1b,
	0x72, 0x97, 0x4c, 0x8d, 0x78, 0x39, 0xcb, 0x68, 0x17, 0x0d, 0xbc, 0x20, 0x94, 0x92, 0xa7, 0x42,
	0x20, 0x3a, 0xa5, 0x77, 0x32, 0x84, 0x17, 0x45, 0xf0, 0x2a, 0xe2, 0x0f, 0x6e, 0xa1, 0x63, 0x9b,
	0x46, 0x8f, 0x26, 0xb3, 0x5e, 0x12, 0xc1, 0x88, 0x11, 0xc4, 0xa6, 0xed, 0x21, 0xf5, 0x19, 0x5e,
	0x8e, 0x03, 0x57, 0xc9, 0xb5, 0x33, 0x9b, 0x32, 0xbc, 0x22, 0xd4, 0x94, 0x3c, 0x7c, 0x85, 0x1c,
	0xc2, 0xab, 0x31, 0x7d, 0x0b, 0x5d, 0x87, 0xbb, 0x93, 0x1e, 0xaf, 0x09, 0x35, 0x2b, 0x55, 0xdd,
	0x63, 0x4a, 0xce, 0xd8, 0xd0, 0x7e, 0x5d, 0xdb, 0x0c, 0x0d, 0xbc, 0x3e, 0xa6, 0x97, 0x91, 0x36,
	0x4d, 0xa2, 0x2d, 0xed, 0x3b, 0x08, 0x6f, 0x88, 0x20, 0x4c, 0xab, 0xb1, 0x19, 0x22, 0xe4, 0x3a,
	0xf0, 0xa6, 0x50, 0xf7, 0xc9, 0x99, 0x96, 0x2b, 0xfa, 0xf9, 0xd0, 0xe1, 0x2a, 0x1a, 0xab, 0x9b,
	0x65, 0x8e, 0xf0, 0x96, 0x50, 0x73, 0xf2, 0x6c, 0x43, 0xbb, 0x0e, 0xd6, 0x88, 0xb7, 0x35, 0xdb,
	0xa2, 0x6d, 0x23, 0xb5, 0xb7, 0x45, 0x90, 0x7d, 0x6d, 0x2f, 0xc7, 0x94, 0xf5, 0x81, 0x99, 0xef,
	0x44, 0x30, 0x55, 0x5b, 0x0c, 0x6d, 0xc0, 0x89, 0xfc, 0xef, 0xc6, 0x56, 0x2d, 0x97, 0x7b, 0x4a,
	0xb1, 0x28, 0x42, 0x93, 0x35, 0xc7, 0x96, 0x4b, 0x78, 0x4f, 0x84, 0x7d, 0xda, 0xa2, 0xf4, 0x3a,
	0x1a, 0x78, 0x3f, 0xaa, 0x3b, 0x6c, 0x76, 0x19, 0x73, 0x74, 0x06, 0x5d, 0x5a, 0xc2, 0x07, 0x91,
	0x4a, 0x2b, 0xef, 0x78, 0x6d, 0x70, 0xc2, 0xfc, 0xc3, 0x88, 0xfc, 0x20, 0xf3, 0xc9, 0xd5, 0xcd,
	0x58, 0xd0, 0x24, 0xaa, 0x6a, 0x57, 0x8e, 0x30, 0x14, 0x70, 0x2b, 0x1a, 0x32, 0x3a, 0x6e, 0xa0,
	0x36, 0xe8, 0xd7, 0x2d, 0x66, 0xa6, 0x98, 0xa8, 0x73, 0x3b, 0xc2, 0xdc, 0x74, 0x8c, 0xde, 0xe9,
	0x6c, 0x1b, 0xfd, 0x00, 0xfd, 0x9a, 0xf7, 0xe4, 0xe1, 0xe7, 0xa8, 0x7d, 0x8d, 0x78, 0xb3, 0x97,
	0x67, 0x18, 0x36, 0x06, 0x0d, 0xfc, 0x22, 0x46, 0x5b, 0x76, 0x45, 0x33, 0xee, 0xea, 0x12, 0x7e,
	0x8d, 0xfc, 0x43, 0x9d, 0x4d, 0xb1, 0xe5, 0xf4, 0x40, 0xdb, 0x2c, 0x0a, 0xf6, 0x5b, 0x2c, 0x1f,
	0xa5, 0x8d, 0x9d, 0xfe, 0x5d, 0xa8, 0x8b, 0x72, 0x76, 0xa3, 0xd9, 0xac, 0x5f, 0x43, 0x5f, 0x58,
	0x72, 0x41, 0xe5, 0xb1, 0x0d, 0xf0, 0x87, 0x50, 0x17, 0xe4, 0xb9, 0x6b, 0xda, 0x5b, 0xed, 0x78,
	0x39, 0x2b, 0xa8, 0x86, 0x1d, 0x62, 0xab, 0x19, 0x0b, 0xf8, 0x73, 0x84, 0xb3, 0xe8, 0xb7, 0xdb,
	0x36, 0xb5, 0xe8, 0x78, 0x9b, 0xc9, 0xeb, 0x0e, 0xc2, 0x5f, 0x71, 0xcf, 0xb7, 0x88, 0xf2, 0xcb,
	0xc8, 0xd1, 0x02, 0xf8, 0x5b, 0x8c, 0x7e, 0xae, 0xb5, 0x3d, 0x0e, 0x8a, 0x1a, 0xf8, 0x47, 0xa8,
	0x4b, 0xf2, 0x81, 0x1a, 0xf2, 0x2e, 0xf9, 0xeb, 0xff, 0xb3, 0x9b, 0xff, 0x8a, 0x95, 0x07, 0xef,
	0xec, 0xcf, 0x27, 0x77, 0xf7, 0xe7, 0x93, 0x1f, 0xf7, 0xe7, 0x13, 0x39, 0x67, 0x69, 0xf8, 0xf6,
	0xe5, 0x61, 0xa3, 0x0f, 0x3c, 0x83, 0x3b, 0x47, 0xe3, 0xb3, 0xf9, 0xe8, 0x7f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x14, 0x5c, 0xc5, 0xb7, 0x7b, 0x05, 0x00, 0x00,
}
