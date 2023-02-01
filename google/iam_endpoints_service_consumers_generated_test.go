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

func TestAccServiceManagementServiceConsumersIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":    randString(t, 10),
		"role":             "roles/servicemanagement.serviceController",
		"project_name":     getTestProjectFromEnv(),
		"consumer_project": getTestProjectFromEnv(),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceManagementServiceConsumersIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_endpoints_service_consumers_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("services/%s/consumers/%s roles/servicemanagement.serviceController", fmt.Sprintf("endpoint%s.endpoints.%s.cloud.goog", context["random_suffix"], context["project_name"]), context["project_name"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccServiceManagementServiceConsumersIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_endpoints_service_consumers_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("services/%s/consumers/%s roles/servicemanagement.serviceController", fmt.Sprintf("endpoint%s.endpoints.%s.cloud.goog", context["random_suffix"], context["project_name"]), context["project_name"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccServiceManagementServiceConsumersIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":    randString(t, 10),
		"role":             "roles/servicemanagement.serviceController",
		"project_name":     getTestProjectFromEnv(),
		"consumer_project": getTestProjectFromEnv(),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccServiceManagementServiceConsumersIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_endpoints_service_consumers_iam_member.foo",
				ImportStateId:     fmt.Sprintf("services/%s/consumers/%s roles/servicemanagement.serviceController user:admin@hashicorptest.com", fmt.Sprintf("endpoint%s.endpoints.%s.cloud.goog", context["random_suffix"], context["project_name"]), context["project_name"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccServiceManagementServiceConsumersIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix":    randString(t, 10),
		"role":             "roles/servicemanagement.serviceController",
		"project_name":     getTestProjectFromEnv(),
		"consumer_project": getTestProjectFromEnv(),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccServiceManagementServiceConsumersIamPolicy_basicGenerated(context),
			},
			{
				ResourceName:      "google_endpoints_service_consumers_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("services/%s/consumers/%s", fmt.Sprintf("endpoint%s.endpoints.%s.cloud.goog", context["random_suffix"], context["project_name"]), context["project_name"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccServiceManagementServiceConsumersIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_endpoints_service_consumers_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("services/%s/consumers/%s", fmt.Sprintf("endpoint%s.endpoints.%s.cloud.goog", context["random_suffix"], context["project_name"]), context["project_name"]),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccServiceManagementServiceConsumersIamMember_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_endpoints_service" "endpoints_service" {
  service_name = "endpoint%{random_suffix}.endpoints.%{project_name}.cloud.goog"
  project = "%{project_name}"
  grpc_config = <<EOF
type: google.api.Service
config_version: 3
name: endpoint%{random_suffix}.endpoints.%{project_name}.cloud.goog
usage:
  rules:
  - selector: endpoints.examples.bookstore.Bookstore.ListShelves
    allow_unregistered_calls: true
EOF
  protoc_output_base64 = "${filebase64("test-fixtures/test_api_descriptor.pb")}"
}

resource "google_endpoints_service_consumers_iam_member" "foo" {
  service_name = google_endpoints_service.endpoints_service.service_name
  consumer_project = "%{consumer_project}"
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccServiceManagementServiceConsumersIamPolicy_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_endpoints_service" "endpoints_service" {
  service_name = "endpoint%{random_suffix}.endpoints.%{project_name}.cloud.goog"
  project = "%{project_name}"
  grpc_config = <<EOF
type: google.api.Service
config_version: 3
name: endpoint%{random_suffix}.endpoints.%{project_name}.cloud.goog
usage:
  rules:
  - selector: endpoints.examples.bookstore.Bookstore.ListShelves
    allow_unregistered_calls: true
EOF
  protoc_output_base64 = "${filebase64("test-fixtures/test_api_descriptor.pb")}"
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_endpoints_service_consumers_iam_policy" "foo" {
  service_name = google_endpoints_service.endpoints_service.service_name
  consumer_project = "%{consumer_project}"
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccServiceManagementServiceConsumersIamPolicy_emptyBinding(context map[string]interface{}) string {
	return Nprintf(`
resource "google_endpoints_service" "endpoints_service" {
  service_name = "endpoint%{random_suffix}.endpoints.%{project_name}.cloud.goog"
  project = "%{project_name}"
  grpc_config = <<EOF
type: google.api.Service
config_version: 3
name: endpoint%{random_suffix}.endpoints.%{project_name}.cloud.goog
usage:
  rules:
  - selector: endpoints.examples.bookstore.Bookstore.ListShelves
    allow_unregistered_calls: true
EOF
  protoc_output_base64 = "${filebase64("test-fixtures/test_api_descriptor.pb")}"
}

data "google_iam_policy" "foo" {
}

resource "google_endpoints_service_consumers_iam_policy" "foo" {
  service_name = google_endpoints_service.endpoints_service.service_name
  consumer_project = "%{consumer_project}"
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccServiceManagementServiceConsumersIamBinding_basicGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_endpoints_service" "endpoints_service" {
  service_name = "endpoint%{random_suffix}.endpoints.%{project_name}.cloud.goog"
  project = "%{project_name}"
  grpc_config = <<EOF
type: google.api.Service
config_version: 3
name: endpoint%{random_suffix}.endpoints.%{project_name}.cloud.goog
usage:
  rules:
  - selector: endpoints.examples.bookstore.Bookstore.ListShelves
    allow_unregistered_calls: true
EOF
  protoc_output_base64 = "${filebase64("test-fixtures/test_api_descriptor.pb")}"
}

resource "google_endpoints_service_consumers_iam_binding" "foo" {
  service_name = google_endpoints_service.endpoints_service.service_name
  consumer_project = "%{consumer_project}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccServiceManagementServiceConsumersIamBinding_updateGenerated(context map[string]interface{}) string {
	return Nprintf(`
resource "google_endpoints_service" "endpoints_service" {
  service_name = "endpoint%{random_suffix}.endpoints.%{project_name}.cloud.goog"
  project = "%{project_name}"
  grpc_config = <<EOF
type: google.api.Service
config_version: 3
name: endpoint%{random_suffix}.endpoints.%{project_name}.cloud.goog
usage:
  rules:
  - selector: endpoints.examples.bookstore.Bookstore.ListShelves
    allow_unregistered_calls: true
EOF
  protoc_output_base64 = "${filebase64("test-fixtures/test_api_descriptor.pb")}"
}

resource "google_endpoints_service_consumers_iam_binding" "foo" {
  service_name = google_endpoints_service.endpoints_service.service_name
  consumer_project = "%{consumer_project}"
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
