package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// WafPSMLog waf p s m log
// swagger:model WafPSMLog
type WafPSMLog struct {

	// Actions generated by this rule. Enum options - WAF_ACTION_NO_OP, WAF_ACTION_BLOCK, WAF_ACTION_ALLOW_PARAMETER. Field introduced in 18.2.3.
	Actions []string `json:"actions,omitempty"`

	// Name of the PSM group. Field introduced in 18.2.3.
	GroupName *string `json:"group_name,omitempty"`

	// UUID of the PSM group. Field introduced in 18.2.3.
	GroupUUID *string `json:"group_uuid,omitempty"`

	// Name of the PSM location. Field introduced in 18.2.3.
	Location *string `json:"location,omitempty"`

	// Transaction data that matched the rule. Field introduced in 18.2.3.
	Matches []*WafRuleMatchData `json:"matches,omitempty"`

	// Rule ID of the matching rule. Field introduced in 18.2.3.
	RuleID *string `json:"rule_id,omitempty"`

	// Name of the matching rule. Field introduced in 18.2.3.
	RuleName *string `json:"rule_name,omitempty"`
}
