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
	"context"
	"log"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func init() {
	resource.AddTestSweepers("ComputeNodeGroup", &resource.Sweeper{
		Name: "ComputeNodeGroup",
		F:    testSweepComputeNodeGroup,
	})
}

// At the time of writing, the CI only passes us-central1 as the region
func testSweepComputeNodeGroup(region string) error {
	resourceName := "ComputeNodeGroup"
	log.Printf("[INFO][SWEEPER_LOG] Starting sweeper for %s", resourceName)

	config, err := sharedConfigForRegion(region)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error getting shared config for region: %s", err)
		return err
	}

	err = config.LoadAndValidate(context.Background())
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error loading: %s", err)
		return err
	}

	t := &testing.T{}
	billingId := getTestBillingAccountFromEnv(t)

	// Setup variables to replace in list template
	d := &ResourceDataMock{
		FieldsInSchema: map[string]interface{}{
			"project":         config.Project,
			"region":          region,
			"location":        region,
			"zone":            "-",
			"billing_account": billingId,
		},
	}

	listTemplate := strings.Split("https://compute.googleapis.com/compute/v1/projects/{{project}}/aggregated/nodeGroups", "?")[0]
	listUrl, err := replaceVars(d, config, listTemplate)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] error preparing sweeper list url: %s", err)
		return nil
	}

	res, err := sendRequest(config, "GET", config.Project, listUrl, config.userAgent, nil)
	if err != nil {
		log.Printf("[INFO][SWEEPER_LOG] Error in response from request %s: %s", listUrl, err)
		return nil
	}

	resourceList, ok := res["items"]
	if !ok {
		log.Printf("[INFO][SWEEPER_LOG] Nothing found in response.")
		return nil
	}
	var rl []interface{}
	zones := resourceList.(map[string]interface{})
	// Loop through every zone in the list response
	for _, zonesValue := range zones {
		zone := zonesValue.(map[string]interface{})
		for k, v := range zone {
			// Zone map either has resources or a warning stating there were no resources found in the zone
			if k != "warning" {
				resourcesInZone := v.([]interface{})
				rl = append(rl, resourcesInZone...)
			}
		}
	}

	log.Printf("[INFO][SWEEPER_LOG] Found %d items in %s list response.", len(rl), resourceName)
	// Keep count of items that aren't sweepable for logging.
	nonPrefixCount := 0
	for _, ri := range rl {
		obj := ri.(map[string]interface{})
		if obj["name"] == nil {
			log.Printf("[INFO][SWEEPER_LOG] %s resource name was nil", resourceName)
			return nil
		}

		name := GetResourceNameFromSelfLink(obj["name"].(string))
		// Skip resources that shouldn't be sweeped
		if !isSweepableTestResource(name) {
			nonPrefixCount++
			continue
		}

		deleteTemplate := "https://compute.googleapis.com/compute/v1/projects/{{project}}/zones/{{zone}}/nodeGroups/{{name}}"
		if obj["zone"] == nil {
			log.Printf("[INFO][SWEEPER_LOG] %s resource zone was nil", resourceName)
			return nil
		}
		zone := GetResourceNameFromSelfLink(obj["zone"].(string))
		deleteTemplate = strings.Replace(deleteTemplate, "{{zone}}", zone, -1)

		deleteUrl, err := replaceVars(d, config, deleteTemplate)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] error preparing delete url: %s", err)
			return nil
		}
		deleteUrl = deleteUrl + name

		// Don't wait on operations as we may have a lot to delete
		_, err = sendRequest(config, "DELETE", config.Project, deleteUrl, config.userAgent, nil)
		if err != nil {
			log.Printf("[INFO][SWEEPER_LOG] Error deleting for url %s : %s", deleteUrl, err)
		} else {
			log.Printf("[INFO][SWEEPER_LOG] Sent delete request for %s resource: %s", resourceName, name)
		}
	}

	if nonPrefixCount > 0 {
		log.Printf("[INFO][SWEEPER_LOG] %d items were non-sweepable and skipped.", nonPrefixCount)
	}

	return nil
}
