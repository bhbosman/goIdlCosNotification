// Code generated by me. DO NOT EDIT.

package goIdlCosNotification

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"
import __yaccidl__ "github.com/bhbosman/yaccidl"

// Struct declaration: "CosNotification::FixedEventHeader", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosNotificationFixedEventHeader struct {
	__goidl__.IdlObject
	EventType CosNotificationEventType `json:"EventType"`
	EventName string `json:"EventName"`
}

//noinspection GoSnakeCaseUsage
const CosNotificationFixedEventHeaderId_Const = "IDL:omg.org/CosNotification/FixedEventHeader:1.0"

func (self *CosNotificationFixedEventHeader) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosNotificationFixedEventHeader) GoString() string {
	return self.String()
}

func (self *CosNotificationFixedEventHeader) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.EventType.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.EventName, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationFixedEventHeader) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationFixedEventHeader) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.EventType.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.EventName)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosNotificationFixedEventHeader_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosNotificationFixedEventHeaderHelper = CosNotificationFixedEventHeader_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosNotificationFixedEventHeaderId_Const,
			__yaccidl__.IdlStruct,
			"CosNotification.idl.FixedEventHeader.go",
			func() __goidl__.IIdlObject {
				return &CosNotificationFixedEventHeader{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosNotificationFixedEventHeader{}
			},
			__reflect__.TypeOf((*CosNotificationFixedEventHeader)(nil))))
}
