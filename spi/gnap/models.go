/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gnap

import (
	"encoding/json"

	"github.com/hyperledger/aries-framework-go/pkg/doc/jose/jwk"
)

// AuthRequest https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-2
type AuthRequest struct {
	// TODO: single TokenRequest is treated like a slice of one element.
	AccessToken []*TokenRequest  `json:"access_token,omitempty"`
	Client      *RequestClient   `json:"client,omitempty"`
	Interact    *RequestInteract `json:"interact,omitempty"`
}

// RequestClient https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-2.3
type RequestClient struct {
	IsReference bool
	Ref         string
	Key         *ClientKey `json:"key"`
}

// ClientKey https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-7.1.1
type ClientKey struct {
	Proof string  `json:"proof"`
	JWK   jwk.JWK `json:"jwk"`
}

// TokenRequest https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-2.1
type TokenRequest struct {
	Access []TokenAccess `json:"access"`
	Label  string        `json:"label,omitempty"`
	Flags  []AccessFlag  `json:"flags,omitempty"`
}

// TokenAccess represents a GNAP token access descriptor, either as a string reference or as an object.
//
// see: https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-8
type TokenAccess struct {
	IsReference bool
	Ref         string
	Type        string `json:"type"`
	Raw         json.RawMessage
}

// RequestInteract https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-2.5
type RequestInteract struct {
	Start  []string      `json:"start"`
	Finish RequestFinish `json:"finish"`
}

// RequestFinish https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-2.5.2
type RequestFinish struct {
	Method string `json:"method"`
	URI    string `json:"uri"`
	Nonce  string `json:"nonce"`
}

// AuthResponse https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-3
type AuthResponse struct {
	Continue    ResponseContinue `json:"continue,omitempty"`
	AccessToken []AccessToken    `json:"access_token,omitempty"`
	Interact    ResponseInteract `json:"interact,omitempty"`
	InstanceID  string           `json:"instance_id,omitempty"`
	Subject     Subject          `json:"subject,omitempty"`
}

// ResponseContinue https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-3.1
type ResponseContinue struct {
	URI         string      `json:"uri,omitempty"`
	AccessToken AccessToken `json:"access_token,omitempty"`
	Wait        int         `json:"wait,omitempty"`
}

// ResponseInteract https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-3.3
type ResponseInteract struct {
	Redirect string `json:"redirect,omitempty"`
	Finish   string `json:"finish,omitempty"`
}

// Subject https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-3.4
type Subject struct {
	SubIDs     []SubjectID        `json:"sub_ids,omitempty"`
	Assertions []SubjectAssertion `json:"assertions,omitempty"`
}

// SubjectID https://www.ietf.org/archive/id/draft-ietf-secevent-subject-identifiers-09.txt
type SubjectID struct {
	ID     string `json:"id,omitempty"`
	Format string `json:"format,omitempty"`
}

// SubjectAssertion https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-3.4
type SubjectAssertion struct {
	Value  string `json:"value,omitempty"`
	Format string `json:"format,omitempty"`
}

// AccessToken https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-3.2.1
type AccessToken struct {
	Value   string        `json:"value,omitempty"`
	Label   string        `json:"label,omitempty"`
	Manage  string        `json:"manage,omitempty"`
	Access  []TokenAccess `json:"access,omitempty"`
	Expires int64         `json:"expires_in,omitempty"` // integer value in seconds.
	Key     string        `json:"key,omitempty"`
	Flags   []AccessFlag  `json:"flags,omitempty"`
}

// ContinueRequest https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-5.1
type ContinueRequest struct {
	InteractRef string `json:"interact_ref"`
}

// ErrorResponse https://www.ietf.org/archive/id/draft-ietf-gnap-core-protocol-09.html#section-3.6
type ErrorResponse struct {
	Error       string `json:"error"`
	Description string `json:"error_description,omitempty"`
}

// IntrospectRequest https://www.ietf.org/archive/id/draft-ietf-gnap-resource-servers-01.html#section-3.3
type IntrospectRequest struct {
	AccessToken    string         `json:"access_token"`
	Proof          string         `json:"proof"`
	Access         []TokenAccess  `json:"access,omitempty"`
	ResourceServer *RequestClient `json:"resource_server,omitempty"`
}

// IntrospectResponse https://www.ietf.org/archive/id/draft-ietf-gnap-resource-servers-01.html#section-3.3
type IntrospectResponse struct {
	Active      bool              `json:"active"`
	Access      []TokenAccess     `json:"access,omitempty"`
	Key         *ClientKey        `json:"key,omitempty"`
	Flags       []AccessFlag      `json:"flags,omitempty"`
	SubjectData map[string]string `json:"subject_data,omitempty"`
}

type AccessFlag string

const (
	Bearer  AccessFlag = "bearer"
	Durable            = "durable"
	Split              = "split"
)
