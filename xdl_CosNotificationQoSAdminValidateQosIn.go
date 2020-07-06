// Code generated by me. DO NOT EDIT.

package goIdlCosNotification

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CosNotification::QoSAdmin_validate_qos_In", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type CosNotificationQoSAdminValidateQosIn struct {
	__goidl__.IdlObject
	RequiredQos CosNotificationQoSProperties `json:"RequiredQos"`
}

//noinspection GoSnakeCaseUsage
const CosNotificationQoSAdminValidateQosInId_Const = "IDL:CosNotification/QoSAdmin_validate_qos_In:1.0"

func (self *CosNotificationQoSAdminValidateQosIn) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *CosNotificationQoSAdminValidateQosIn) GoString() string {
	return self.String()
}

func (self *CosNotificationQoSAdminValidateQosIn) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.RequiredQos.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationQoSAdminValidateQosIn) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationQoSAdminValidateQosIn) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.RequiredQos.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type CosNotificationQoSAdminValidateQosIn_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var CosNotificationQoSAdminValidateQosInHelper = CosNotificationQoSAdminValidateQosIn_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			CosNotificationQoSAdminValidateQosInId_Const,
			__goidl__.StructType,
			"CosNotification.idl",
			"xdl_CosNotificationQoSAdminValidateQosIn.go",
			func() __goidl__.IIdlObject {
				return &CosNotificationQoSAdminValidateQosIn{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosNotificationQoSAdminValidateQosIn{}
			},
			__reflect__.TypeOf((*CosNotificationQoSAdminValidateQosIn)(nil))))
}
