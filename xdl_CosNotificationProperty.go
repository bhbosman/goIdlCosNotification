// Code generated by me. DO NOT EDIT.

package goIdlCosNotification

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CosNotification::Property", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosNotificationProperty struct {
	__goidl__.IdlObject
	Name string `json:"Name"`
	Value __goidl__.IdlAny `json:"Value"`
}

//noinspection GoSnakeCaseUsage
const CosNotificationPropertyId_Const = "IDL:omg.org/CosNotification/Property:1.0"

func (self *CosNotificationProperty) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosNotificationProperty) GoString() string {
	return self.String()
}

func (self *CosNotificationProperty) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Name, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(AnyType)
	self.Value, err = __goidl__.IdlAnyHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationProperty) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationProperty) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Name)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(AnyType)
	err = __goidl__.IdlAnyHelper.Write(stream, self.Value)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosNotificationProperty_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosNotificationPropertyHelper = CosNotificationProperty_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosNotificationPropertyId_Const,
			__goidl__.StructType,
			"CosNotification.idl",
			"xdl_CosNotificationProperty.go",
			func() __goidl__.IIdlObject {
				return &CosNotificationProperty{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosNotificationProperty{}
			},
			__reflect__.TypeOf((*CosNotificationProperty)(nil))))
}
