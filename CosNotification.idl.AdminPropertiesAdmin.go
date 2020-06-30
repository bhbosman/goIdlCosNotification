// Code generated by me. DO NOT EDIT.

package goIdlCosNotification

import __goidl__ "github.com/bhbosman/goidl"

// Interface declaration: "CosNotification::AdminPropertiesAdmin", generatedBy by: "WriteInterface"
type CosNotificationAdminPropertiesAdmin interface {
	// Interface operations
	// Original name: "get_admin"
	GetAdmin() (result *CosNotificationAdminProperties, err error)
	//Exceptions for : SetAdmin
	//	CosNotificationUnsupportedAdmin
	// Original name: "set_admin"
	SetAdmin(Admin *CosNotificationAdminProperties) (error error)
}

//noinspection GoSnakeCaseUsage
type CosNotificationAdminPropertiesAdmin_Helper struct {
}

//noinspection GoSnakeCaseUsage
const CosNotificationAdminPropertiesAdminId_Const = "IDL:omg.org/CosNotification/AdminPropertiesAdmin:1.0"

func (self CosNotificationAdminPropertiesAdmin_Helper) Id() string {
	return CosNotificationAdminPropertiesAdminId_Const
}

func (self CosNotificationAdminPropertiesAdmin_Helper) Read(stream __goidl__.IReadAny) (CosNotificationAdminPropertiesAdmin, error) {
	return nil, nil
}

func (self CosNotificationAdminPropertiesAdmin_Helper) Write(stream __goidl__.IWriteAny, v CosNotificationAdminPropertiesAdmin) error {
	return nil
}


//noinspection GoUnusedGlobalVariable
var CosNotificationAdminPropertiesAdminHelper = CosNotificationAdminPropertiesAdmin_Helper{}

func init() {
}
