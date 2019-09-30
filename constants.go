// Copyright (c) 2019 Sulaeman (me@sulaeman.com), All rights reserved.
// This source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package telesign

import "time"

const (
	version = "1.0.0"

	// EnvStandard is the standard env
	EnvStandard = "Standard"
	// EnvEnterprise is the enterprise env
	EnvEnterprise = "Enterprise"

	// MessageARN type
	MessageARN = "ARN"
	// MessageMKT type
	MessageMKT = "MKT"
	// MessageOTP type
	MessageOTP = "OTP"

	// UcidAtck For use in a 2FA situation like updating an account, or trying to log in.
	UcidAtck = "ATCK"
	// UcidBacf For creating an account on somebody's service in a situation where the service may be vulnerable to bulk attacks or individual fraudsters.
	UcidBacf = "BACF"
	// UcidBacs For creating an account on somebody's service in a situation where the service may be vulnerable to bulk attacks or individual spammers.
	UcidBacs = "BACS"
	// UcidChbk For use in a situation such as someone trying to buy something expensive or unusual on your platform, and you want to verify that it is really them.
	UcidChbk = "CHBK"
	// UcidCldr Calendar Event
	UcidCldr = "CLDR"
	// UcidLead For use in a situation where you require a person to enter their personal information in order to obtain information about something like a loan or realty or a school, and you want to check if the person is bogus or not.
	UcidLead = "LEAD"
	// UcidOthr For use in a 2FA situation like updating an account, or trying to log in.
	UcidOthr = "OTHR"
	// UcidPwrt For use in a situation where a password reset is required.
	UcidPwrt = "PWRT"
	// UcidResv For use in a situation where you have end users making reservations and not showing up, and you want to be able to include some kind of phone verification in that loop.
	UcidResv = "RESV"
	// UcidRxpf For use in situations where you are trying to prevent prescription fraud.
	UcidRxpf = "RXPF"
	// UcidShip For use in situations where you are sending a shipping notification.
	UcidShip = "SHIP"
	// UcidThef For use in situations where you are trying to prevent an end user from deactivating or redirecting a phone number in order to take over someone else's identity.
	UcidThef = "THEF"
	// UcidTrvf For use in situations where you are transferring money, and you want to check to see if it is approved by sending a text message or phone call to your end user. This is similar to CHBK, but is specifically for a money transaction.
	UcidTrvf = "TRVF"
	// UcidUnkn is the same as OTHR.
	UcidUnkn = "UNKN"

	// PhoneTypeFixedLineCode number
	PhoneTypeFixedLineCode = "1"
	// PhoneTypeFixedLine name
	PhoneTypeFixedLine = "FIXED_LINE"
	// PhoneTypeMobileCode number
	PhoneTypeMobileCode = "2"
	// PhoneTypeMobile name
	PhoneTypeMobile = "MOBILE"
	// PhoneTypePrepaidCode number
	PhoneTypePrepaidCode = "3"
	// PhoneTypePrepaid name
	PhoneTypePrepaid = "PREPAID"
	// PhoneTypeTollFreeCode number
	PhoneTypeTollFreeCode = "4"
	// PhoneTypeTollFree name
	PhoneTypeTollFree = "TOLL_FREE"
	// PhoneTypeVoipCode number
	PhoneTypeVoipCode = "5"
	// PhoneTypeVoip name
	PhoneTypeVoip = "VOIP"
	// PhoneTypePagerCode number
	PhoneTypePagerCode = "6"
	// PhoneTypePager name
	PhoneTypePager = "PAGER"
	// PhoneTypePayphoneCode number
	PhoneTypePayphoneCode = "7"
	// PhoneTypePayphone name
	PhoneTypePayphone = "PAYPHONE"
	// PhoneTypeInvalidCode number
	PhoneTypeInvalidCode = "8"
	// PhoneTypeInvalid name
	PhoneTypeInvalid = "INVALID"
	// PhoneTypeRestrictedPremiumCode number
	PhoneTypeRestrictedPremiumCode = "9"
	// PhoneTypeRestrictedPremium name
	PhoneTypeRestrictedPremium = "RESTRICTED_PREMIUM"
	// PhoneTypePersonalCode number
	PhoneTypePersonalCode = "10"
	// PhoneTypePersonal name
	PhoneTypePersonal = "PERSONAL"
	// PhoneTypeVoicemailCode number
	PhoneTypeVoicemailCode = "11"
	// PhoneTypeVoicemail name
	PhoneTypeVoicemail = "VOICEMAIL"
	// PhoneTypeOtherCode number
	PhoneTypeOtherCode = "12"
	// PhoneTypeOther name
	PhoneTypeOther = "OTHER"

	// CallForwardActionBlock name
	CallForwardActionBlock = "block"

	// AccountLifecycleEventCreate name
	AccountLifecycleEventCreate = "create"

	// Active name
	Active = "ACTIVE"
	// Reachable name
	Reachable = "REACHABLE"
	// Unavailable name
	Unavailable = "UNAVAILABLE"

	defaultHTTPTimeout = 20
	timeFormat         = time.RFC1123Z
	authMethod         = "HMAC-SHA256"

	domain           = "https://rest-api.telesign.com"
	domainEnterprise = "https://rest-ww.telesign.com"
)
