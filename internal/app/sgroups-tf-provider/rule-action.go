package provider

import (
	protos "github.com/wildberries-tech/sgroups/v2/pkg/api/sgroups"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
)

var (
	actionValidator = stringvalidator.OneOf(
		protos.RuleAction_DROP.String(),
		protos.RuleAction_ACCEPT.String())
)
