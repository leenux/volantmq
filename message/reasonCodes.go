package message

// ReasonCode contains return codes across all MQTT specs
type ReasonCode byte

// CodeIssuer who is message issuer
type CodeIssuer byte

// nolint: golint
const (
	CodeIssuerServer CodeIssuer = 0x00
	CodeIssuerClient
	CodeIssuerBoth
	CodeIssuerInvalid
)

// nolint: golint                                            // V3.1.1 \  V5.0
const ( ///////////////////////////////////////////////////////    |   \    |
	CodeSuccess                            ReasonCode = 0x00 //    |   \    |
	CodeRefusedUnacceptableProtocolVersion ReasonCode = 0x01 //    |   \    |
	CodeRefusedIdentifierRejected          ReasonCode = 0x02 //    |   \    |
	CodeRefusedServerUnavailable           ReasonCode = 0x03 //    |   \    |
	CodeRefusedBadUsernameOrPassword       ReasonCode = 0x04 //    |   \    |
	CodeRefusedNotAuthorized               ReasonCode = 0x05 // <--|   \    |
	CodeNoMatchingSubscribers              ReasonCode = 0x10 //        \    |
	CodeNoSubscriptionExisted              ReasonCode = 0x11 //        \    |
	CodeContinueAuthentication             ReasonCode = 0x18 //        \    |
	CodeReAuthenticate                     ReasonCode = 0x19 //        \    |
	CodeUnspecifiedError                   ReasonCode = 0x80 //        \    |
	CodeMalformedPacket                    ReasonCode = 0x81 //        \    |
	CodeProtocolError                      ReasonCode = 0x82 //        \    |
	CodeImplementationSpecificError        ReasonCode = 0x83 //        \    |
	CodeUnsupportedProtocol                ReasonCode = 0x84 //        \    |
	CodeInvalidClientID                    ReasonCode = 0x85 //        \    |
	CodeBadUserOrPassword                  ReasonCode = 0x86 //        \    |
	CodeNotAuthorized                      ReasonCode = 0x87 //        \    |
	CodeServerUnavailable                  ReasonCode = 0x88 //        \    |
	CodeServerBusy                         ReasonCode = 0x89 //        \    |
	CodeBanned                             ReasonCode = 0x8A //        \    |
	CodeServerShuttingDown                 ReasonCode = 0x8B //        \    |
	CodeBadAuthMethod                      ReasonCode = 0x8C //        \    |
	CodeSessionTakenOver                   ReasonCode = 0x8E //        \    |
	CodeKeepAliveTimeout                   ReasonCode = 0x8F //        \    |
	CodeTopicFilterNotValid                ReasonCode = 0x90 //        \    |
	CodePacketIdInUse                      ReasonCode = 0x91 //        \    |
	CodePacketIdNotFound                   ReasonCode = 0x92 //        \    |
	CodePacketTooLarge                     ReasonCode = 0x95 //        \    |
	CodeMessageRateTooHigh                 ReasonCode = 0x96 //        \    |
	CodeQuotaExceeded                      ReasonCode = 0x97 //        \    |
	CodeAdministrativeAction               ReasonCode = 0x98 //        \    |
	CodeDisconnectWithWillMessage          ReasonCode = 0x99 //        \    |
	CodeRetainUnavailable                  ReasonCode = 0x9A //        \    |
	CodeMaximumQoS                         ReasonCode = 0x9B //        \    |
	CodeUseAnotherServer                   ReasonCode = 0x9C //        \    |
	CodeServerMoved                        ReasonCode = 0x9D //        \    |
	CodeSharedSubscriptionNotSupported     ReasonCode = 0x9E //        \    |
	CodeConnectionRateExceeded             ReasonCode = 0x9F //        \    |
	CodeMaximumConnectTime                 ReasonCode = 0xA0 //        \    |
	CodeSubscriptionIdNotSupported         ReasonCode = 0xA1 //        \    |
	CodeWildcardSubscriptionsNotSupported  ReasonCode = 0xA2 //        \ <--|
)

var packetTypeCodeMap = map[PacketType]map[ReasonCode]struct {
	iss  CodeIssuer
	desc string
}{
	CONNACK: {
		CodeSuccess:                            {iss: CodeIssuerServer, desc: "The Connection is accepted"},
		CodeRefusedUnacceptableProtocolVersion: {iss: CodeIssuerClient, desc: "The Server does not support the level of the MQTT protocol requested by the Client"},
		CodeRefusedIdentifierRejected:          {iss: CodeIssuerClient, desc: "The Client identifier is not allowed"},
		CodeRefusedServerUnavailable:           {iss: CodeIssuerClient, desc: "Server refused connection"},
		CodeRefusedBadUsernameOrPassword:       {iss: CodeIssuerClient, desc: "The data in the user name or password is malformed"},
		CodeRefusedNotAuthorized:               {iss: CodeIssuerClient, desc: "The Client is not authorized to connect"},
		CodeUnspecifiedError:                   {iss: CodeIssuerServer, desc: "The Server does not wish to reveal the reason for the failure, or none of the other Return Codes apply"},
		CodeMalformedPacket:                    {iss: CodeIssuerServer, desc: "Data within the CONNECT Packet was not consistent with this specification"},
		CodeImplementationSpecificError:        {iss: CodeIssuerServer, desc: "The CONNECT is valid but is not accepted by this Server"},
		CodeUnsupportedProtocol:                {iss: CodeIssuerServer, desc: "The Server does not support the level of the MQTT protocol requested by the Client"},
		CodeInvalidClientID:                    {iss: CodeIssuerServer, desc: "The Client Identifier is a valid string but is not allowed by the Server"},
		CodeBadUserOrPassword:                  {iss: CodeIssuerServer, desc: "The Server does not accept the username or password specified by the Client"},
		CodeNotAuthorized:                      {iss: CodeIssuerServer, desc: "The Client is not authorized to connect"},
		CodeServerUnavailable:                  {iss: CodeIssuerServer, desc: "The MQTT Server is not available"},
		CodeServerBusy:                         {iss: CodeIssuerServer, desc: "The Server is busy. Try again later"},
		CodeBanned:                             {iss: CodeIssuerServer, desc: "This Client has been banned by administrative action. Contact the server administrator"},
		CodeBadAuthMethod:                      {iss: CodeIssuerServer, desc: "The authentication method is not supported or does not match the authentication method currently in use"},
		CodeTopicFilterNotValid:                {iss: CodeIssuerServer, desc: "The Will Topic Is Invalid"},
		CodePacketTooLarge:                     {iss: CodeIssuerServer, desc: "The Connect Packet exceeded the maximum permissible size"},
		CodeUseAnotherServer:                   {iss: CodeIssuerServer, desc: "The Client should temporarily use another server"},
		CodeServerMoved:                        {iss: CodeIssuerServer, desc: "The Client should permanently use another server"},
		CodeConnectionRateExceeded:             {iss: CodeIssuerServer, desc: "The connection rate limit has been exceeded"},
	},
	PUBACK: {
		CodeSuccess:                     {iss: CodeIssuerBoth, desc: "The message is accepted. Publication of the QoS 1 message proceeds"},
		CodeNoMatchingSubscribers:       {iss: CodeIssuerBoth, desc: "The message is accepted but there are no subscribers"},
		CodeUnspecifiedError:            {iss: CodeIssuerBoth, desc: "The receiver does not accept the publish but either does not want to reveal the reason, or it does not match one of the other values"},
		CodeImplementationSpecificError: {iss: CodeIssuerBoth, desc: "The PUBLISH is valid but the receiver is not willing to accept it"},
		CodeNotAuthorized:               {iss: CodeIssuerBoth, desc: "The PUBLISH is not authorized"},
		CodeTopicFilterNotValid:         {iss: CodeIssuerBoth, desc: "The topic name is valid, but is not accepted"},
		CodeQuotaExceeded:               {iss: CodeIssuerBoth, desc: "An implementation imposed limit has been exceeded"},
	},
	PUBREC: {
		CodeSuccess:                     {iss: CodeIssuerBoth, desc: "The message is accepted. Publication of the QoS 2 message proceeds"},
		CodeNoMatchingSubscribers:       {iss: CodeIssuerBoth, desc: "The message is accepted but there are no subscribers"},
		CodeUnspecifiedError:            {iss: CodeIssuerBoth, desc: "The receiver does not accept the publish but either does not want to reveal the reason, or it does not match one of the other values"},
		CodeImplementationSpecificError: {iss: CodeIssuerBoth, desc: "The PUBLISH is valid but the receiver is not willing to accept it"},
		CodeNotAuthorized:               {iss: CodeIssuerBoth, desc: "The PUBLISH is not authorized"},
		CodeTopicFilterNotValid:         {iss: CodeIssuerBoth, desc: "The topic name is valid, but is not accepted"},
		CodePacketIdInUse:               {iss: CodeIssuerBoth, desc: "The PacketID is already in use. Possible mismatch in the session state between the Client and Server"},
		CodeQuotaExceeded:               {iss: CodeIssuerBoth, desc: "An implementation imposed limit has been exceeded"},
	},
	PUBREL: {
		CodeSuccess:          {iss: CodeIssuerBoth, desc: "Message released. Publication of QoS 2 message is complete"},
		CodePacketIdNotFound: {iss: CodeIssuerBoth, desc: "The PacketID is not known. Possible a mismatch between the Session state on the Client and Server"},
	},
	PUBCOMP: {
		CodeSuccess:          {iss: CodeIssuerBoth, desc: "Message released. Publication of QoS 2 message is complete"},
		CodePacketIdNotFound: {iss: CodeIssuerBoth, desc: "The PacketID is not known. Possible a mismatch between the Session state on the Client and Server"},
	},
	SUBACK: {
		0:                                     {iss: CodeIssuerBoth, desc: "The subscription is accepted and the maximum QoS sent will be QoS 0. This might be a lower QoS than was requested"}, // Maximum QoS 0
		1:                                     {iss: CodeIssuerBoth, desc: "The subscription is accepted and the maximum QoS sent will be QoS 1. This might be a lower QoS than was requested"}, // Maximum QoS 1
		2:                                     {iss: CodeIssuerBoth, desc: "The subscription is accepted and any received QoS will be sent to this subscription"},                               // Maximum QoS 2
		CodeUnspecifiedError:                  {iss: CodeIssuerBoth, desc: "The subscription is not accepted and the Server either does not wish to reveal the reason or none of the other Return Codes apply"},
		CodeImplementationSpecificError:       {iss: CodeIssuerBoth, desc: "The SUBSCRIBE is valid but the Server does not accept i"},
		CodeNotAuthorized:                     {iss: CodeIssuerBoth, desc: "The Client is not authorized to make this subscription"},
		CodeTopicFilterNotValid:               {iss: CodeIssuerBoth, desc: "The Topic Filter is correctly formed but is not allowed for this client"},
		CodePacketIdInUse:                     {iss: CodeIssuerBoth, desc: "The specified packet identifier is already in use"},
		CodeQuotaExceeded:                     {iss: CodeIssuerBoth, desc: "An implementation imposed limit has been exceeded"},
		CodeSharedSubscriptionNotSupported:    {iss: CodeIssuerBoth, desc: "The Server does not support shared subscriptions for this Client"},
		CodeSubscriptionIdNotSupported:        {iss: CodeIssuerBoth, desc: "The Server does not support subscription identifiers; the subscription is not accepted"},
		CodeWildcardSubscriptionsNotSupported: {iss: CodeIssuerBoth, desc: "The Server does not support Wildcard subscription; the subscription is not accepted"},
	},
	UNSUBACK: {
		CodeSuccess:                     {iss: CodeIssuerBoth, desc: "The subscription is deleted"},
		CodeNoSubscriptionExisted:       {iss: CodeIssuerBoth, desc: "No matching subscription existed"},
		CodeUnspecifiedError:            {iss: CodeIssuerBoth, desc: "The unsubscribe could not be completed and the Server either does not wish to reveal the reason or none of the other Return Codes apply"},
		CodeImplementationSpecificError: {iss: CodeIssuerBoth, desc: "The UNSUBSCRIBE is valid but the Server does not accept it"},
		CodeNotAuthorized:               {iss: CodeIssuerBoth, desc: "The client is not authorized to unsubscribe"},
		CodeTopicFilterNotValid:         {iss: CodeIssuerBoth, desc: "The topic filter is correctly formed but is not allowed for this client"},
		CodePacketIdInUse:               {iss: CodeIssuerBoth, desc: "The specified packet identifier is already in use"},
	},
	DISCONNECT: {
		CodeSuccess:                           {iss: CodeIssuerClient, desc: "Close the connection normally. Do not send the Will Message"},
		CodeUnspecifiedError:                  {iss: CodeIssuerBoth, desc: "The Connection is closed but the sender either does not wish to reveal the reason, or none of the other Return Codes apply"},
		CodeMalformedPacket:                   {iss: CodeIssuerBoth, desc: "The received packet does not conform to this specification"},
		CodeProtocolError:                     {iss: CodeIssuerBoth, desc: "An unexpected or out of order packet was received"},
		CodeImplementationSpecificError:       {iss: CodeIssuerBoth, desc: "The packet received is valid but cannot be processed by this implementation"},
		CodeNotAuthorized:                     {iss: CodeIssuerServer, desc: "The request is not authorized"},
		CodeServerBusy:                        {iss: CodeIssuerServer, desc: "The Server is busy and cannot continue processing this Client"},
		CodeServerShuttingDown:                {iss: CodeIssuerServer, desc: "The Server is shutting down"},
		CodeSessionTakenOver:                  {iss: CodeIssuerServer, desc: "Another Connection using the same ClientId has connected causing this Connection to be closed"},
		CodeKeepAliveTimeout:                  {iss: CodeIssuerServer, desc: "The Connection is closed because no Packet has been received for 1.5 times the Keep alive time"},
		CodeTopicFilterNotValid:               {iss: CodeIssuerBoth, desc: "The topic name or filter is valid, but is not accepted"},
		CodePacketTooLarge:                    {iss: CodeIssuerBoth, desc: "The packet size is too large"},
		CodeMessageRateTooHigh:                {iss: CodeIssuerBoth, desc: "The rate of publish is too high"},
		CodeQuotaExceeded:                     {iss: CodeIssuerBoth, desc: "An implementation imposed limit has been exceeded"},
		CodeAdministrativeAction:              {iss: CodeIssuerBoth, desc: "The Connection is closed due to an administrative action"},
		CodeDisconnectWithWillMessage:         {iss: CodeIssuerClient, desc: "The client wishes to disconnect but requires that the Server also publishes its Will message"},
		CodeRetainUnavailable:                 {iss: CodeIssuerServer, desc: "The Server has specified Retain unavailable in the CONNACK"},
		CodeMaximumQoS:                        {iss: CodeIssuerServer, desc: "The Client specified a QoS greater then the QoS specified in a Maximum QoS in the CONNACK"},
		CodeUseAnotherServer:                  {iss: CodeIssuerServer, desc: "The Client should temporarily change its Server"},
		CodeServerMoved:                       {iss: CodeIssuerServer, desc: "The Server is moved and the client should permanently change its server location"},
		CodeSharedSubscriptionNotSupported:    {iss: 0, desc: ""},
		CodeConnectionRateExceeded:            {iss: 0, desc: ""},
		CodeMaximumConnectTime:                {iss: CodeIssuerServer, desc: "The maximum connection time authorized for this connection has been exceeded"},
		CodeSubscriptionIdNotSupported:        {iss: CodeIssuerServer, desc: "The Server does not support subscription identifiers; the subscription is not accepted"},
		CodeWildcardSubscriptionsNotSupported: {iss: CodeIssuerServer, desc: "The Server does not support Wildcard subscription; the subscription is not accepted"},
	},
	AUTH: {
		CodeSuccess: {
			iss:  CodeIssuerServer,
			desc: "Authentication is successful",
		},
		CodeContinueAuthentication: {
			iss:  CodeIssuerBoth,
			desc: "Continue the authentication with another step",
		},
		CodeReAuthenticate: {
			iss:  CodeIssuerClient,
			desc: "Initiate a re-authentication",
		},
	},
}

var codeDescMap = map[ReasonCode]string{
	CodeSuccess:                            "Operation success",
	CodeRefusedUnacceptableProtocolVersion: "The Server does not support the level of the MQTT protocol requested by the Client",
	CodeRefusedIdentifierRejected:          "The Client identifier is not allowed",
	CodeRefusedServerUnavailable:           "Server refused connection",
	CodeRefusedBadUsernameOrPassword:       "The data in the user name or password is malformed",
	CodeRefusedNotAuthorized:               "The Client is not authorized to connect",
	CodeNoMatchingSubscribers:              "The message is accepted but there are no subscribers",
	CodeNoSubscriptionExisted:              "No matching subscription existed",
	CodeContinueAuthentication:             "Continue the authentication with another step",
	CodeReAuthenticate:                     "Initiate a re-authentication",
	CodeUnspecifiedError:                   "Return code not specified by application",
	CodeMalformedPacket:                    "Data within the Packet was not consistent with this specification",
	CodeProtocolError:                      "",
	CodeImplementationSpecificError:        "",
	CodeUnsupportedProtocol:                "",
	CodeInvalidClientID:                    "",
	CodeBadUserOrPassword:                  "",
	CodeNotAuthorized:                      "",
	CodeServerUnavailable:                  "",
	CodeServerBusy:                         "",
	CodeBanned:                             "",
	CodeServerShuttingDown:                 "",
	CodeBadAuthMethod:                      "",
	CodeSessionTakenOver:                   "",
	CodeKeepAliveTimeout:                   "",
	CodeTopicFilterNotValid:                "",
	CodePacketIdInUse:                      "",
	CodePacketIdNotFound:                   "",
	CodePacketTooLarge:                     "",
	CodeMessageRateTooHigh:                 "",
	CodeQuotaExceeded:                      "",
	CodeAdministrativeAction:               "",
	CodeDisconnectWithWillMessage:          "",
	CodeRetainUnavailable:                  "",
	CodeMaximumQoS:                         "",
	CodeUseAnotherServer:                   "",
	CodeServerMoved:                        "",
	CodeSharedSubscriptionNotSupported:     "",
	CodeConnectionRateExceeded:             "",
	CodeMaximumConnectTime:                 "",
	CodeSubscriptionIdNotSupported:         "",
	CodeWildcardSubscriptionsNotSupported:  "",
}

// PacketTypeDir check direction of packet type
func (c ReasonCode) PacketTypeDir(p PacketType) (CodeIssuer, error) {
	pT, ok := packetTypeCodeMap[p]
	if !ok {
		return CodeIssuerInvalid, ErrInvalidMessageType
	}

	if issuer, ok := pT[c]; ok {
		return issuer.iss, ErrInvalidReturnCode
	}

	return CodeIssuerInvalid, nil
}

// Value convert reason code to byte type
func (c ReasonCode) Value() byte {
	return byte(c)
}

// IsValid check either reason code is valid across all MQTT specs
func (c ReasonCode) IsValid() bool {
	if _, ok := codeDescMap[c]; ok {
		return true
	}
	return false
}

// IsValidV3 check either reason code is valid for MQTT V3.1/V3.1.1 or not
func (c ReasonCode) IsValidV3() bool {
	if c >= CodeSuccess && c <= CodeRefusedNotAuthorized {
		return true
	}
	return false
}

// IsValidV5 check either reason code is valid for MQTT V5.0 or not
func (c ReasonCode) IsValidV5() bool {
	if c == CodeSuccess || (c >= CodeNoMatchingSubscribers && c <= CodeWildcardSubscriptionsNotSupported) {
		return true
	}
	return false
}

// IsValidForType check either reason code is valid for giver packet type
func (c ReasonCode) IsValidForType(p PacketType) bool {
	pT, ok := packetTypeCodeMap[p]
	if !ok {
		return false
	}

	if _, ok = pT[c]; !ok {
		return false
	}

	return true
}

// Error returns the description of the ReturnCode
func (c ReasonCode) Error() string {
	if s, ok := codeDescMap[c]; ok {
		return s
	}

	return "Unknown error"
}

// Desc return code description
func (c ReasonCode) Desc() string {
	return c.Error()
}