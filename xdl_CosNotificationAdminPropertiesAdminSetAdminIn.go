// Code generated by me. DO NOT EDIT.

package goIdlCosNotification

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CosNotification::AdminPropertiesAdmin_set_admin_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosNotificationAdminPropertiesAdminSetAdminIn struct {
	__goidl__.IdlObject
	Admin CosNotificationAdminProperties `json:"Admin"`
}

//noinspection GoSnakeCaseUsage
const CosNotificationAdminPropertiesAdminSetAdminInId_Const = "IDL:CosNotification/AdminPropertiesAdmin_set_admin_In:1.0"

func (self *CosNotificationAdminPropertiesAdminSetAdminIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosNotificationAdminPropertiesAdminSetAdminIn) GoString() string {
	return self.String()
}

func (self *CosNotificationAdminPropertiesAdminSetAdminIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.Admin.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationAdminPropertiesAdminSetAdminIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationAdminPropertiesAdminSetAdminIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.Admin.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosNotificationAdminPropertiesAdminSetAdminIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosNotificationAdminPropertiesAdminSetAdminInHelper = CosNotificationAdminPropertiesAdminSetAdminIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosNotificationAdminPropertiesAdminSetAdminInId_Const,
			__goidl__.StructType,
			"CosNotification.idl",
			"xdl_CosNotificationAdminPropertiesAdminSetAdminIn.go",
			func() __goidl__.IIdlObject {
				return &CosNotificationAdminPropertiesAdminSetAdminIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosNotificationAdminPropertiesAdminSetAdminIn{}
			},
			__reflect__.TypeOf((*CosNotificationAdminPropertiesAdminSetAdminIn)(nil))))
}