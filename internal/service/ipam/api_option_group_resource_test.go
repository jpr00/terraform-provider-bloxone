package ipam_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/acctest"
)

func TestAccOptionGroupResource_basic(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_group.test"
	var v ipam.IpamsvcOptionGroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionGroupBasicConfig("option_group_test", "ip4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionGroupExists(context.Background(), resourceName, &v),
					// TODO: check and validate these
					resource.TestCheckResourceAttr(resourceName, "name", "option_group_test"),
					resource.TestCheckResourceAttr(resourceName, "protocol", "ip4"),
					// Test Read Only fields
					resource.TestCheckResourceAttrSet(resourceName, "created_at"),
					resource.TestCheckResourceAttrSet(resourceName, "id"),
					resource.TestCheckResourceAttrSet(resourceName, "updated_at"),
					// Test fields with default value
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOptionGroupResource_disappears(t *testing.T) {
	resourceName := "bloxone_dhcp_option_group.test"
	var v ipam.IpamsvcOptionGroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		CheckDestroy:             testAccCheckOptionGroupDestroy(context.Background(), &v),
		Steps: []resource.TestStep{
			{
				Config: testAccOptionGroupBasicConfig("option_group_test", "ip4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionGroupExists(context.Background(), resourceName, &v),
					testAccCheckOptionGroupDisappears(context.Background(), &v),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccOptionGroupResource_Comment(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_group.test_comment"
	var v ipam.IpamsvcOptionGroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionGroupComment("option_group_test", "ip4", "test comment"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment"),
				),
			},
			// Update and Read
			{
				Config: testAccOptionGroupComment("option_group_test", "ip4", "test comment update"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "comment", "test comment update"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOptionGroupResource_DhcpOptions(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_group.test_dhcp_options"
	var v ipam.IpamsvcOptionGroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionGroupDhcpOptions("option_group_test", "ip4", "option", "true"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.0.option_value", "true"),
					resource.TestCheckResourceAttrPair(resourceName, "dhcp_options.0.option_code", "bloxone_dhcp_option_code.test", "id"),
				),
			},
			// Update and Read
			{
				Config: testAccOptionGroupDhcpOptions("option_group_test", "ip4", "option", "false"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "dhcp_options.0.option_value", "false"),
					resource.TestCheckResourceAttrPair(resourceName, "dhcp_options.0.option_code", "bloxone_dhcp_option_code.test", "id"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOptionGroupResource_Name(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_group.test_name"
	var v ipam.IpamsvcOptionGroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionGroupName("option_group_test", "ip4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "option_group_test"),
				),
			},
			// Update and Read
			{
				Config: testAccOptionGroupName("option_group_test_1", "ip4"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "name", "option_group_test_1"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func TestAccOptionGroupResource_Tags(t *testing.T) {
	var resourceName = "bloxone_dhcp_option_group.test_tags"
	var v ipam.IpamsvcOptionGroup

	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read
			{
				Config: testAccOptionGroupTags("option_group_test", "ip4", map[string]string{
					"tag1": "value1",
					"tag2": "value2",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag1", "value1"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2"),
				),
			},
			// Update and Read
			{
				Config: testAccOptionGroupTags("option_group_test", "ip4", map[string]string{
					"tag2": "value2changed",
					"tag3": "value3",
				}),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckOptionGroupExists(context.Background(), resourceName, &v),
					resource.TestCheckResourceAttr(resourceName, "tags.tag2", "value2changed"),
					resource.TestCheckResourceAttr(resourceName, "tags.tag3", "value3"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccCheckOptionGroupExists(ctx context.Context, resourceName string, v *ipam.IpamsvcOptionGroup) resource.TestCheckFunc {
	// Verify the resource exists in the cloud
	return func(state *terraform.State) error {
		rs, ok := state.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("not found: %s", resourceName)
		}
		apiRes, _, err := acctest.BloxOneClient.IPAddressManagementAPI.
			OptionGroupAPI.
			OptionGroupRead(ctx, rs.Primary.ID).
			Execute()
		if err != nil {
			return err
		}
		if !apiRes.HasResult() {
			return fmt.Errorf("expected result to be returned: %s", resourceName)
		}
		*v = apiRes.GetResult()
		return nil
	}
}

func testAccCheckOptionGroupDestroy(ctx context.Context, v *ipam.IpamsvcOptionGroup) resource.TestCheckFunc {
	// Verify the resource was destroyed
	return func(state *terraform.State) error {
		_, httpRes, err := acctest.BloxOneClient.IPAddressManagementAPI.
			OptionGroupAPI.
			OptionGroupRead(ctx, *v.Id).
			Execute()
		if err != nil {
			if httpRes != nil && httpRes.StatusCode == http.StatusNotFound {
				// resource was deleted
				return nil
			}
			return err
		}
		return errors.New("expected to be deleted")
	}
}

func testAccCheckOptionGroupDisappears(ctx context.Context, v *ipam.IpamsvcOptionGroup) resource.TestCheckFunc {
	// Delete the resource externally to verify disappears test
	return func(state *terraform.State) error {
		_, err := acctest.BloxOneClient.IPAddressManagementAPI.
			OptionGroupAPI.
			OptionGroupDelete(ctx, *v.Id).
			Execute()
		if err != nil {
			return err
		}
		return nil
	}
}

func testAccOptionGroupBasicConfig(name, protocol string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_option_group" "test" {
    name = %q
    protocol = %q
}
`, name, protocol)
}

func testAccOptionGroupComment(name, protocol, comment string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_option_group" "test_comment" {
    name = %q
    protocol = %q
    comment = %q
}
`, name, protocol, comment)
}

func testAccOptionGroupDhcpOptions(name, protocol, type_, optValue string) string {
	config := fmt.Sprintf(`
resource "bloxone_dhcp_option_group" "test_dhcp_options" {
    name = %q
    protocol = %q
    dhcp_options = [
      {
       type = %q
       option_code = bloxone_dhcp_option_code.test.id
       option_value = %q
      }
    ]
}
`, name, protocol, type_, optValue)

	return strings.Join([]string{testAccOptionCodeBasicConfig("234", "test_dhcp_option_code", "boolean"), config}, "")
}

func testAccOptionGroupName(name, protocol string) string {
	return fmt.Sprintf(`
resource "bloxone_dhcp_option_group" "test_name" {
    name = %q
    protocol = %q
}
`, name, protocol)
}

func testAccOptionGroupTags(name, protocol string, tags map[string]string) string {
	tagsStr := "{\n"
	for k, v := range tags {
		tagsStr += fmt.Sprintf(`
		%s = %q
`, k, v)
	}
	tagsStr += "\t}"

	return fmt.Sprintf(`
resource "bloxone_dhcp_option_group" "test_tags" {
    name = %q
    protocol = %q
    tags = %s
}
`, name, protocol, tagsStr)
}
