package schemas

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"k8s.io/kops/pkg/apis/kops"
)

func ExternalNetworkingSpec() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{},
	}
}

func ExpandExternalNetworkingSpec(in map[string]interface{}) kops.ExternalNetworkingSpec {
	if in == nil {
		panic("expand ExternalNetworkingSpec failure, in is nil")
	}
	return kops.ExternalNetworkingSpec{}
}

func FlattenExternalNetworkingSpecInto(in kops.ExternalNetworkingSpec, out map[string]interface{}) {
}

func FlattenExternalNetworkingSpec(in kops.ExternalNetworkingSpec) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenExternalNetworkingSpecInto(in, out)
	return out
}
