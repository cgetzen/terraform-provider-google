// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var ServiceManagementServiceIamSchema = map[string]*schema.Schema{
	"service_name": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type ServiceManagementServiceIamUpdater struct {
	serviceName string
	d           TerraformResourceData
	Config      *Config
}

func ServiceManagementServiceIamUpdaterProducer(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	if v, ok := d.GetOk("service_name"); ok {
		values["service_name"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"services/(?P<service_name>[^/]+)", "(?P<service_name>[^/]+)"}, d, config, d.Get("service_name").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ServiceManagementServiceIamUpdater{
		serviceName: values["service_name"],
		d:           d,
		Config:      config,
	}

	if err := d.Set("service_name", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting service_name: %s", err)
	}

	return u, nil
}

func ServiceManagementServiceIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	m, err := getImportIdQualifiers([]string{"services/(?P<service_name>[^/]+)", "(?P<service_name>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &ServiceManagementServiceIamUpdater{
		serviceName: values["service_name"],
		d:           d,
		Config:      config,
	}
	if err := d.Set("service_name", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting service_name: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *ServiceManagementServiceIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyServiceUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.userAgent)
	if err != nil {
		return nil, err
	}

	policy, err := sendRequest(u.Config, "POST", "", url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *ServiceManagementServiceIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyServiceUrl("setIamPolicy")
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.userAgent)
	if err != nil {
		return err
	}

	_, err = sendRequestWithTimeout(u.Config, "POST", "", url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *ServiceManagementServiceIamUpdater) qualifyServiceUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{ServiceManagementBasePath}}%s:%s", fmt.Sprintf("services/%s", u.serviceName), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *ServiceManagementServiceIamUpdater) GetResourceId() string {
	return fmt.Sprintf("services/%s", u.serviceName)
}

func (u *ServiceManagementServiceIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-servicemanagement-service-%s", u.GetResourceId())
}

func (u *ServiceManagementServiceIamUpdater) DescribeResource() string {
	return fmt.Sprintf("servicemanagement service %q", u.GetResourceId())
}
