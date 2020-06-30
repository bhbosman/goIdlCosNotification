// Code generated by me. DO NOT EDIT.

package goIdlCosNotification

import __goidl__ "github.com/bhbosman/goidl"

// Enum declaration: "CosNotification::QoSError_code", generatedBy by: "WriteEnumDcl"
type CosNotificationQoSErrorCode uint32

//noinspection GoUnusedConst
const (
	CosNotificationQoSErrorCodeUnsupportedProperty CosNotificationQoSErrorCode = 0
	CosNotificationQoSErrorCodeUnavailableProperty CosNotificationQoSErrorCode = 1
	CosNotificationQoSErrorCodeUnsupportedValue    CosNotificationQoSErrorCode = 2
	CosNotificationQoSErrorCodeUnavailableValue    CosNotificationQoSErrorCode = 3
	CosNotificationQoSErrorCodeBadProperty         CosNotificationQoSErrorCode = 4
	CosNotificationQoSErrorCodeBadType             CosNotificationQoSErrorCode = 5
	CosNotificationQoSErrorCodeBadValue            CosNotificationQoSErrorCode = 6
)

//noinspection GoSnakeCaseUsage
type CosNotificationQoSErrorCode_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CosNotificationQoSErrorCodeId_Const = "IDL:omg.org/CosNotification/QoSError_code:1.0"

func (self CosNotificationQoSErrorCode_Helper) Id() string {
	return CosNotificationQoSErrorCodeId_Const
}

func (self CosNotificationQoSErrorCode_Helper) Read(stream __goidl__.IReadAny) (uint32, error) {
	result, err := __goidl__.IdlUInt32Helper.Read(stream)
	return result, err
}

func (self CosNotificationQoSErrorCode_Helper) Write(stream __goidl__.IWriteAny, v uint32) error {
	return __goidl__.IdlUInt32Helper.Write(stream, v)
}


//noinspection GoUnusedGlobalVariable
var CosNotificationQoSErrorCodeHelper = CosNotificationQoSErrorCode_Helper{}

func init() {
}
