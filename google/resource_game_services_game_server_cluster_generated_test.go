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

func TestAccGameServicesGameServerCluster_gameServiceClusterBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"agones_cluster": "bootstrapped-agones-cluster",
		"random_suffix":  randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckGameServicesGameServerClusterDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGameServicesGameServerCluster_gameServiceClusterBasicExample(context),
			},
			{
				ResourceName:            "google_game_services_game_server_cluster.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"cluster_id", "realm_id", "location"},
			},
		},
	})
}

func testAccGameServicesGameServerCluster_gameServiceClusterBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_game_services_game_server_cluster" "default" {
    
  cluster_id = "%{agones_cluster}"
  realm_id   = google_game_services_realm.default.realm_id

  connection_info {
    gke_cluster_reference {
      cluster = "locations/us-west1/clusters/%{agones_cluster}"
    }
    namespace = "default"
  }
}

resource "google_game_services_realm" "default" {
  realm_id   = "realm%{random_suffix}"
  time_zone  = "PST8PDT"

  description = "Test Game Realm"
}
`, context)
}

func testAccCheckGameServicesGameServerClusterDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_game_services_game_server_cluster" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{GameServicesBasePath}}projects/{{project}}/locations/{{location}}/realms/{{realm_id}}/gameServerClusters/{{cluster_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("GameServicesGameServerCluster still exists at %s", url)
			}
		}

		return nil
	}
}
