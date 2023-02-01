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

func TestAccDialogflowCXFlow_dialogflowcxFlowFullExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDialogflowCXFlowDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDialogflowCXFlow_dialogflowcxFlowFullExample(context),
			},
			{
				ResourceName:            "google_dialogflow_cx_flow.basic_flow",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"parent"},
			},
		},
	})
}

func testAccDialogflowCXFlow_dialogflowcxFlowFullExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dialogflow_cx_agent" "agent" {
  display_name = "tf-test-dialogflowcx-agent%{random_suffix}"
  location = "global"
  default_language_code = "en"
  supported_language_codes = ["fr","de","es"]
  time_zone = "America/New_York"
  description = "Example description."
  avatar_uri = "https://cloud.google.com/_static/images/cloud/icons/favicons/onecloud/super_cloud.png"
  enable_stackdriver_logging = true
  enable_spell_correction    = true
	speech_to_text_settings {
		enable_speech_adaptation = true
	}
}


resource "google_dialogflow_cx_flow" "basic_flow" {
  parent       = google_dialogflow_cx_agent.agent.id
  display_name = "MyFlow"
  description  = "Test Flow"

  nlu_settings {
		classification_threshold = 0.3 
		model_type               = "MODEL_TYPE_STANDARD"
	}

  event_handlers {
		   event                    = "custom-event"
		   trigger_fulfillment {
			    return_partial_responses = false
				messages {
					text {
						text  = ["I didn't get that. Can you say it again?"]
					}
				}
		    }
		}

		event_handlers {
			event                    = "sys.no-match-default"
			trigger_fulfillment {
				 return_partial_responses = false
				 messages {
					 text {
						 text  = ["Sorry, could you say that again?"]
					 }
				 }
			 }
		 }

		 event_handlers {
			event                    = "sys.no-input-default"
			trigger_fulfillment {
				 return_partial_responses = false
				 messages {
					 text {
						 text  = ["One more time?"]
					 }
				 }
			 }
		 }
} 
`, context)
}

func testAccCheckDialogflowCXFlowDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dialogflow_cx_flow" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DialogflowCXBasePath}}{{parent}}/flows/{{name}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("DialogflowCXFlow still exists at %s", url)
			}
		}

		return nil
	}
}
