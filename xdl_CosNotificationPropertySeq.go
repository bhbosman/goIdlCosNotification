// Code generated by me. DO NOT EDIT.

package goIdlCosNotification

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "CosNotification::PropertySeq", generatedBy by: "WriteStructSequenceDcl"
type CosNotificationPropertySeq struct {
	__goidl__.IdlObject
	Array []*CosNotificationProperty `json:"Array"`
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

func (self *CosNotificationPropertySeq) ReadValue(stream __goidl__.IReadAny) error {
	err := self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	var n uint32
	n, err = __goidl__.IdlUInt32Helper.Read(stream)
	if err != nil {
		return err
	}
	if n > 0 {
		self.Array = make([]*CosNotificationProperty, n)
		var i uint32
		for i = 0; i < n; i++ {
			self.Array[i] = &CosNotificationProperty{}
			err = self.Array[i].ReadValue(stream)
				if err != nil {
				return err
			}
		}
	}
	return nil
}

func (self *CosNotificationPropertySeq) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *CosNotificationPropertySeq) Write(stream __goidl__.IWriteAny) error {
	err := self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	err = __goidl__.IdlUInt32Helper.Write(stream, uint32(len(self.Array)))
	if err != nil {
	return err
	}
	if len(self.Array) > 0 {
		for _, item := range self.Array {
			err = item.Write(stream)
			if err != nil {
				return err
			}
		}
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
			__goidl__.SequenceType,
			"CosNotification.idl",
			"xdl_CosNotificationPropertySeq.go",
			func() __goidl__.IIdlObject {
				return &CosNotificationPropertySeq{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &CosNotificationPropertySeq{}
			},
			__reflect__.TypeOf((*CosNotificationPropertySeq)(nil))))
}
