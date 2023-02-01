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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccMonitoringService_monitoringServiceCustomExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckMonitoringServiceDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccMonitoringService_monitoringServiceCustomExample(context),
			},
			{
				ResourceName:      "google_monitoring_custom_service.custom",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMonitoringService_monitoringServiceCustomExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_monitoring_custom_service" "custom" {
  service_id = "tf-test-custom-srv%{random_suffix}"
  display_name = "My Custom Service tf-test-custom-srv%{random_suffix}"

  telemetry {
  	resource_name = "//product.googleapis.com/foo/foo/services/test%{random_suffix}"
  }

  user_labels = {
    my_key       = "my_value"
    my_other_key = "my_other_value"
  }
}
`, context)
}

func testAccCheckMonitoringServiceDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_monitoring_custom_service" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{MonitoringBasePath}}v3/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil, isMonitoringConcurrentEditError)
			if err == nil {
				return fmt.Errorf("MonitoringService still exists at %s", url)
			}
		}

		return nil
	}
}
