// Copyright 2018 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package terraform_utils

import (
	"fmt"
	"waze/terraformer/terraform_utils/provider_wrapper"

	"github.com/hashicorp/terraform/terraform"
)

type TerraformResource struct {
	InstanceInfo  *terraform.InstanceInfo
	InstanceState *terraform.InstanceState
	ResourceName  string
	Item          interface{}
	Provider      string
}

func NewTerraformResource(ID, resourceName, resourceType, provider string, item interface{}, attributes map[string]string) TerraformResource {
	return TerraformResource{
		ResourceName: TfSanitize(resourceName),
		Item:         item,
		Provider:     provider,
		InstanceState: &terraform.InstanceState{
			ID:         ID,
			Attributes: attributes,
		},
		InstanceInfo: &terraform.InstanceInfo{
			Type: resourceType,
			Id:   fmt.Sprintf("%s.%s", resourceType, TfSanitize(resourceName)),
		},
	}
}

func (r *TerraformResource) Refresh(provider *provider_wrapper.ProviderWrapper) {
	r.InstanceState, _ = provider.Refresh(r.InstanceInfo, r.InstanceState)
}
