// Code generated by me. DO NOT EDIT.

package goIdlCosNotification

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CosNotification::NamedPropertyRange", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosNotificationNamedPropertyRange struct {
	__goidl__.IdlObject
	Name string `json:"Name"`
	Range CosNotificationPropertyRange `json:"Range"`
}

//noinspection GoSnakeCaseUsage
const CosNotificationNamedPropertyRangeId_Const = "IDL:omg.org/CosNotification/NamedPropertyRange:1.0"

func (self *CosNotificationNamedPropertyRange) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosNotificationNamedPropertyRange) GoString() string {
	return self.String()
}

func (self *CosNotificationNamedPropertyRange) ReadValue(stream __goidl__.IReadAny) error {
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
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.Range.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationNamedPropertyRange) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationNamedPropertyRange) Write(stream __goidl__.IWriteAny) error {
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
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.Range.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosNotificationNamedPropertyRange_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosNotificationNamedPropertyRangeHelper = CosNotificationNamedPropertyRange_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosNotificationNamedPropertyRangeId_Const,
			__goidl__.StructType,
			"CosNotification.idl",
			"xdl_CosNotificationNamedPropertyRange.go",
			func() __goidl__.IIdlObject {
				return &CosNotificationNamedPropertyRange{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosNotificationNamedPropertyRange{}
			},
			__reflect__.TypeOf((*CosNotificationNamedPropertyRange)(nil))))
}
