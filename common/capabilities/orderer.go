/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package capabilities

import (
	cb "github.com/hyperledger/fabric/protos/common"
)

const (
	ordererTypeName = "Orderer"

	// OrdererV1_1 is the capabilties string for standard new non-backwards compatible fabric v1.1 orderer capabilities.
	OrdererV1_1 = "V1.1"
)

// OrdererProvider provides capabilities information for orderer level config.
type OrdererProvider struct {
	*registry
	v11BugFixes bool
}

// NewOrdererProvider creates an orderer capabilities provider.
func NewOrdererProvider(capabilities map[string]*cb.Capability) *OrdererProvider {
	cp := &OrdererProvider{}
	cp.registry = newRegistry(cp, capabilities)
	_, cp.v11BugFixes = capabilities[OrdererV1_1]
	return cp
}

// Type returns a descriptive string for logging purposes.
func (cp *OrdererProvider) Type() string {
	return ordererTypeName
}

// HasCapability returns true if the capability is supported by this binary.
func (cp *OrdererProvider) HasCapability(capability string) bool {
	switch capability {
	// Add new capability names here
	case OrdererV1_1:
		return true
	default:
		return false
	}
}

// SetChannelModPolicyDuringCreate specifies whether the v1.0 undesirable behavior of setting the /Channel
// group's mod_policy to "" should be fixed or not.
func (cp *OrdererProvider) SetChannelModPolicyDuringCreate() bool {
	return cp.v11BugFixes
}
