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
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccCloudIotDeviceRegistryIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
		"project":       getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIotDeviceRegistryIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloudiot_registry_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/registries/%s roles/viewer", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudiot-registry%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccCloudIotDeviceRegistryIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_cloudiot_registry_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/registries/%s roles/viewer", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudiot-registry%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudIotDeviceRegistryIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
		"project":       getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccCloudIotDeviceRegistryIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloudiot_registry_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/registries/%s roles/viewer user:admin@hashicorptest.com", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudiot-registry%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccCloudIotDeviceRegistryIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
		"role":          "roles/viewer",
		"project":       getTestProjectFromEnv(),
		"region":        getTestRegionFromEnv(),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCloudIotDeviceRegistryIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_cloudiot_registry_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/registries/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudiot-registry%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccCloudIotDeviceRegistryIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_cloudiot_registry_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/registries/%s", getTestProjectFromEnv(), getTestRegionFromEnv(), fmt.Sprintf("tf-test-cloudiot-registry%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCloudIotDeviceRegistryIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudiot_registry" "test-registry" {
  name     = "tf-test-cloudiot-registry%{random_suffix}"
}

resource "google_cloudiot_registry_iam_member" "foo" {
  project = google_cloudiot_registry.test-registry.project
  region = google_cloudiot_registry.test-registry.region
  name = google_cloudiot_registry.test-registry.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccCloudIotDeviceRegistryIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudiot_registry" "test-registry" {
  name     = "tf-test-cloudiot-registry%{random_suffix}"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_cloudiot_registry_iam_policy" "foo" {
  project = google_cloudiot_registry.test-registry.project
  region = google_cloudiot_registry.test-registry.region
  name = google_cloudiot_registry.test-registry.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccCloudIotDeviceRegistryIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudiot_registry" "test-registry" {
  name     = "tf-test-cloudiot-registry%{random_suffix}"
}

data "google_iam_policy" "foo" {
}

resource "google_cloudiot_registry_iam_policy" "foo" {
  project = google_cloudiot_registry.test-registry.project
  region = google_cloudiot_registry.test-registry.region
  name = google_cloudiot_registry.test-registry.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccCloudIotDeviceRegistryIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudiot_registry" "test-registry" {
  name     = "tf-test-cloudiot-registry%{random_suffix}"
}

resource "google_cloudiot_registry_iam_binding" "foo" {
  project = google_cloudiot_registry.test-registry.project
  region = google_cloudiot_registry.test-registry.region
  name = google_cloudiot_registry.test-registry.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccCloudIotDeviceRegistryIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_cloudiot_registry" "test-registry" {
  name     = "tf-test-cloudiot-registry%{random_suffix}"
}

resource "google_cloudiot_registry_iam_binding" "foo" {
  project = google_cloudiot_registry.test-registry.project
  region = google_cloudiot_registry.test-registry.region
  name = google_cloudiot_registry.test-registry.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
