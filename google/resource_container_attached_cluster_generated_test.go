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

func TestAccContainerAttachedCluster_containerAttachedClusterBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContainerAttachedClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccContainerAttachedCluster_containerAttachedClusterBasicExample(context),
			},
			{
				ResourceName:            "google_container_attached_cluster.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccContainerAttachedCluster_containerAttachedClusterBasicExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_project" "project" {
}

data "google_container_attached_versions" "versions" {
	location       = "us-west1"
	project        = data.google_project.project.project_id
}

resource "google_container_attached_cluster" "primary" {
  name     = "basic%{random_suffix}"
  location = "us-west1"
  project = data.google_project.project.project_id
  description = "Test cluster"
  distribution = "aks"
  oidc_config {
      issuer_url = "https://oidc.issuer.url"
  }
  platform_version = data.google_container_attached_versions.versions.valid_versions[0]
  fleet {
    project = "projects/${data.google_project.project.number}"
  }
}
`, context)
}

func TestAccContainerAttachedCluster_containerAttachedClusterFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContainerAttachedClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccContainerAttachedCluster_containerAttachedClusterFullExample(context),
			},
			{
				ResourceName:            "google_container_attached_cluster.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location"},
			},
		},
	})
}

func testAccContainerAttachedCluster_containerAttachedClusterFullExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_project" "project" {
}

data "google_container_attached_versions" "versions" {
	location       = "us-west1"
	project        = data.google_project.project.project_id
}

resource "google_container_attached_cluster" "primary" {
  name     = "basic%{random_suffix}"
  project = data.google_project.project.project_id
  location = "us-west1"
  description = "Test cluster"
  distribution = "aks"
  annotations = {
    label-one = "value-one"
  }
  authorization {
    admin_users = [ "user1@example.com", "user2@example.com"]
  }
  oidc_config {
      issuer_url = "https://oidc.issuer.url"
      jwks = base64encode("{\"keys\":[{\"use\":\"sig\",\"kty\":\"RSA\",\"kid\":\"testid\",\"alg\":\"RS256\",\"n\":\"somedata\",\"e\":\"AQAB\"}]}")
  }
  platform_version = data.google_container_attached_versions.versions.valid_versions[0]
  fleet {
    project = "projects/${data.google_project.project.number}"
  }
  logging_config {
    component_config {
      enable_components = ["SYSTEM_COMPONENTS", "WORKLOADS"]
    }
  }
  monitoring_config {
    managed_prometheus_config {
      enabled = true
    }
  }
}
`, context)
}

func TestAccContainerAttachedCluster_containerAttachedClusterIgnoreErrorsExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckContainerAttachedClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccContainerAttachedCluster_containerAttachedClusterIgnoreErrorsExample(context),
			},
			{
				ResourceName:            "google_container_attached_cluster.primary",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"location", "deletion_policy"},
			},
		},
	})
}

func testAccContainerAttachedCluster_containerAttachedClusterIgnoreErrorsExample(context map[string]interface{}) string {
	return Nprintf(`
data "google_project" "project" {
}

data "google_container_attached_versions" "versions" {
	location       = "us-west1"
	project        = data.google_project.project.project_id
}

resource "google_container_attached_cluster" "primary" {
  name     = "basic%{random_suffix}"
  location = "us-west1"
  project = data.google_project.project.project_id
  description = "Test cluster"
  distribution = "aks"
  oidc_config {
      issuer_url = "https://oidc.issuer.url"
  }
  platform_version = data.google_container_attached_versions.versions.valid_versions[0]
  fleet {
    project = "projects/${data.google_project.project.number}"
  }

  deletion_policy = "DELETE_IGNORE_ERRORS"
}
`, context)
}

func testAccCheckContainerAttachedClusterDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_container_attached_cluster" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{ContainerAttachedBasePath}}projects/{{project}}/locations/{{location}}/attachedClusters/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("ContainerAttachedCluster still exists at %s", url)
			}
		}

		return nil
	}
}
