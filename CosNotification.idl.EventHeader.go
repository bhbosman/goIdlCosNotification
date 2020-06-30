// Code generated by me. DO NOT EDIT.

package goIdlCosNotification

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"
import __yaccidl__ "github.com/bhbosman/yaccidl"

// Struct declaration: "CosNotification::EventHeader", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosNotificationEventHeader struct {
	__goidl__.IdlObject
	FixedHeader CosNotificationFixedEventHeader `json:"FixedHeader"`
	VariableHeader CosNotificationOptionalHeaderFields `json:"VariableHeader"`
}

//noinspection GoSnakeCaseUsage
const CosNotificationEventHeaderId_Const = "IDL:omg.org/CosNotification/EventHeader:1.0"

func (self *CosNotificationEventHeader) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosNotificationEventHeader) GoString() string {
	return self.String()
}

func (self *CosNotificationEventHeader) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.FixedHeader.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.VariableHeader.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationEventHeader) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationEventHeader) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.FixedHeader.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.VariableHeader.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosNotificationEventHeader_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosNotificationEventHeaderHelper = CosNotificationEventHeader_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosNotificationEventHeaderId_Const,
			__yaccidl__.IdlStruct,
			"CosNotification.idl.EventHeader.go",
			func() __goidl__.IIdlObject {
				return &CosNotificationEventHeader{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosNotificationEventHeader{}
			},
			__reflect__.TypeOf((*CosNotificationEventHeader)(nil))))
}
