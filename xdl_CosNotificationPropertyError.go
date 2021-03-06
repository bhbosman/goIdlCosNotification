// Code generated by me. DO NOT EDIT.

package goIdlCosNotification

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CosNotification::PropertyError", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosNotificationPropertyError struct {
	__goidl__.IdlObject
	Code uint32 `json:"Code"`
	Name string `json:"Name"`
	AvailableRange CosNotificationPropertyRange `json:"AvailableRange"`
}

//noinspection GoSnakeCaseUsage
const CosNotificationPropertyErrorId_Const = "IDL:omg.org/CosNotification/PropertyError:1.0"

func (self *CosNotificationPropertyError) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosNotificationPropertyError) GoString() string {
	return self.String()
}

func (self *CosNotificationPropertyError) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlEnum)
	self.Code, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Name, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.AvailableRange.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationPropertyError) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationPropertyError) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlEnum)
	err = __goidl__.IdlUInt32Helper.Write(stream, self.Code)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Name)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.AvailableRange.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosNotificationPropertyError_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosNotificationPropertyErrorHelper = CosNotificationPropertyError_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosNotificationPropertyErrorId_Const,
			__goidl__.StructType,
			"CosNotification.idl",
			"xdl_CosNotificationPropertyError.go",
			func() __goidl__.IIdlObject {
				return &CosNotificationPropertyError{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosNotificationPropertyError{}
			},
			__reflect__.TypeOf((*CosNotificationPropertyError)(nil))))
}
