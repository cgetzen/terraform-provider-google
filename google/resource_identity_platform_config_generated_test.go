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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIdentityPlatformConfig_identityPlatformConfigBasicExample(t *testing.T) {
	skipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        getTestOrgFromEnv(t),
		"billing_acct":  getTestBillingAccountFromEnv(t),
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIdentityPlatformConfig_identityPlatformConfigBasicExample(context),
			},
			{
				ResourceName:      "google_identity_platform_config.default",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIdentityPlatformConfig_identityPlatformConfigBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_project" "default" {
  project_id = "tf-test%{random_suffix}"
  name       = "tf-test%{random_suffix}"
  org_id     = "%{org_id}"
  billing_account =  "%{billing_acct}"
  labels = {
    firebase = "enabled"
  }
}

resource "google_project_service" "identitytoolkit" {
  project = google_project.default.project_id
  service = "identitytoolkit.googleapis.com"
}


resource "google_identity_platform_config" "default" {
  project = google_project.default.project_id
  autodelete_anonymous_users = true
}
`, context)
}
