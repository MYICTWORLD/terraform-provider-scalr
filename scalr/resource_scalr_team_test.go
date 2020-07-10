package scalr

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	tfe "github.com/scalr/go-scalr"
)

func TestAccTFETeam_basic(t *testing.T) {
	team := &tfe.Team{}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTFETeamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTFETeam_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTFETeamExists(
						"scalr_team.foobar", team),
					testAccCheckTFETeamAttributes(team),
					resource.TestCheckResourceAttr(
						"scalr_team.foobar", "name", "team-test"),
				),
			},
		},
	})
}

func TestAccTFETeam_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTFETeamDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTFETeam_basic,
			},

			{
				ResourceName:        "scalr_team.foobar",
				ImportState:         true,
				ImportStateIdPrefix: "tst-terraform/",
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCheckTFETeamExists(
	n string, team *tfe.Team) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		tfeClient := testAccProvider.Meta().(*tfe.Client)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No instance ID is set")
		}

		t, err := tfeClient.Teams.Read(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}

		if t == nil {
			return fmt.Errorf("Team not found")
		}

		*team = *t

		return nil
	}
}

func testAccCheckTFETeamAttributes(
	team *tfe.Team) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if team.Name != "team-test" {
			return fmt.Errorf("Bad name: %s", team.Name)
		}
		return nil
	}
}

func testAccCheckTFETeamDestroy(s *terraform.State) error {
	tfeClient := testAccProvider.Meta().(*tfe.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "scalr_team" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No instance ID is set")
		}

		_, err := tfeClient.Teams.Read(ctx, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Team %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccTFETeam_basic = `
resource "scalr_organization" "foobar" {
  name  = "tst-terraform"
  email = "admin@company.com"
}

resource "scalr_team" "foobar" {
  name         = "team-test"
  organization = "${scalr_organization.foobar.id}"
}`
