// Code generated by me. DO NOT EDIT.

package goIdlCosNotification

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"
import __yaccidl__ "github.com/bhbosman/yaccidl"

// Struct declaration: "CosNotification::EventType", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosNotificationEventType struct {
	__goidl__.IdlObject
	DomainName string `json:"DomainName"`
	TypeName string `json:"TypeName"`
}

//noinspection GoSnakeCaseUsage
const CosNotificationEventTypeId_Const = "IDL:omg.org/CosNotification/EventType:1.0"

func (self *CosNotificationEventType) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosNotificationEventType) GoString() string {
	return self.String()
}

func (self *CosNotificationEventType) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.DomainName, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.TypeName, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationEventType) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationEventType) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.DomainName)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.TypeName)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosNotificationEventType_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosNotificationEventTypeHelper = CosNotificationEventType_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosNotificationEventTypeId_Const,
			__yaccidl__.IdlStruct,
			"CosNotification.idl.EventType.go",
			func() __goidl__.IIdlObject {
				return &CosNotificationEventType{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosNotificationEventType{}
			},
			__reflect__.TypeOf((*CosNotificationEventType)(nil))))
}