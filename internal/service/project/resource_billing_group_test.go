package project_test

import (
	"fmt"
	"testing"

	acc "github.com/aiven/terraform-provider-aiven/internal/acctest"

	"github.com/aiven/aiven-go-client"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAivenBillingGroup_basic(t *testing.T) {
	resourceName := "aiven_billing_group.foo"
	rName := acctest.RandStringFromCharSet(10, acctest.CharSetAlphaNum)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acc.TestAccPreCheck(t) },
		ProviderFactories: acc.TestAccProviderFactories,
		CheckDestroy:      testAccCheckAivenBillingGroupResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCopyFromProjectBillingGroupResource(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("aiven_project.pr02", "billing_group"),
				),
			},
			{
				Config: testAccBillingGroupResource(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", fmt.Sprintf("test-acc-bg-%s", rName)),
				),
			},
			{
				Config: testCopyBillingGroupFromExistingOne(rName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("aiven_billing_group.bar2", "name", fmt.Sprintf("copy-test-acc-bg-%s", rName)),
					resource.TestCheckResourceAttr("aiven_billing_group.bar", "billing_currency", "EUR"),
					resource.TestCheckResourceAttr("aiven_billing_group.bar2", "billing_currency", "EUR"),
					resource.TestCheckResourceAttr("aiven_billing_group.bar2", "city", "Helsinki"),
					resource.TestCheckResourceAttr("aiven_billing_group.bar2", "company", "Aiven Oy"),
					resource.TestCheckResourceAttr("aiven_billing_group.bar2", "vat_id", "abc"),
				),
			},
		},
	})
}

func testAccCheckAivenBillingGroupResourceDestroy(s *terraform.State) error {
	c := acc.TestAccProvider.Meta().(*aiven.Client)

	// loop through the resources in state, verifying each billing group is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aiven_billing_group" {
			continue
		}

		db, err := c.BillingGroup.Get(rs.Primary.ID)
		if err != nil && !aiven.IsNotFound(err) && err.(aiven.Error).Status != 500 {
			return fmt.Errorf("error getting a billing group by id: %w", err)
		}

		if db != nil {
			return fmt.Errorf("billing group (%s) still exists", rs.Primary.ID)
		}
	}

	return nil
}

func testAccBillingGroupResource(name string) string {
	return fmt.Sprintf(`
resource "aiven_billing_group" "foo" {
  name           = "test-acc-bg-%s"
  billing_emails = ["ivan.savciuc+test1@aiven.fi", "ivan.savciuc+test2@aiven.fi"]
}

data "aiven_billing_group" "bar" {
  billing_group_id = aiven_billing_group.foo.id
}

resource "aiven_project" "pr1" {
  project       = "test-acc-pr-%s"
  billing_group = aiven_billing_group.foo.id

  depends_on = [aiven_billing_group.foo]
}`, name, name)
}

func testAccCopyFromProjectBillingGroupResource(name string) string {
	return fmt.Sprintf(`
resource "aiven_billing_group" "foo" {
  name             = "test-acc-bg-%s"
  billing_currency = "USD"
  vat_id           = "abc"
}

resource "aiven_project" "pr01" {
  project       = "test-acc-pr01-%s"
  billing_group = aiven_billing_group.foo.id
  depends_on    = [aiven_billing_group.foo]
}

resource "aiven_project" "pr02" {
  project           = "test-acc-p02-%s"
  copy_from_project = aiven_project.pr01.project

  depends_on = [aiven_project.pr01]
}`, name, name, name)
}

func testCopyBillingGroupFromExistingOne(name string) string {
	return fmt.Sprintf(`
resource "aiven_billing_group" "bar" {
  name             = "test-acc-bg-%s"
  billing_currency = "EUR"
  vat_id           = "abc"
  city             = "Helsinki"
  company          = "Aiven Oy"
}
resource "aiven_billing_group" "bar2" {
  name                    = "copy-test-acc-bg-%s"
  copy_from_billing_group = aiven_billing_group.bar.id
}`, name, name)
}
