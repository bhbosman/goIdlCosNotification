// Code generated by me. DO NOT EDIT.

package goIdlCosNotification

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"
import __yaccidl__ "github.com/bhbosman/yaccidl"

// Struct declaration: "CosNotification::PropertySeq", generatedBy by: "WriteSeqStructDcl"
type CosNotificationPropertySeq struct {
	__goidl__.IdlObject
	Array []CosNotificationProperty `json:"Array"`
}

//noinspection GoSnakeCaseUsage
const CosNotificationPropertySeqId_Const = "IDL:CosNotification/PropertySeq:1.0"

func (self *CosNotificationPropertySeq) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosNotificationPropertySeq) GoString() string {
	return self.String()
}

func (self *CosNotificationPropertySeq) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosNotificationPropertySeq_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosNotificationPropertySeqHelper = CosNotificationPropertySeq_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosNotificationPropertySeqId_Const,
			__yaccidl__.IdlStruct,
			"CosNotification.idl.PropertySeq.go",
			func() __goidl__.IIdlObject {
				return &CosNotificationPropertySeq{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosNotificationPropertySeq{}
			},
			__reflect__.TypeOf((*CosNotificationPropertySeq)(nil))))
}
