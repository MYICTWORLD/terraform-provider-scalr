package scalr

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	tfe "github.com/scalr/go-tfe"
)

func TestAccTFETeamAccess_basic(t *testing.T) {
	tmAccess := &tfe.TeamAccess{}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTFETeamAccessDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTFETeamAccess_basic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckTFETeamAccessExists(
						"scalr_team_access.foobar", tmAccess),
					testAccCheckTFETeamAccessAttributes(tmAccess),
					resource.TestCheckResourceAttr(
						"scalr_team_access.foobar", "access", "write"),
				),
			},
		},
	})
}

func TestAccTFETeamAccess_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckTFETeamAccessDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccTFETeamAccess_basic,
			},

			{
				ResourceName:        "scalr_team_access.foobar",
				ImportState:         true,
				ImportStateIdPrefix: "tst-terraform/workspace-test/",
				ImportStateVerify:   true,
			},
		},
	})
}

func testAccCheckTFETeamAccessExists(
	n string, tmAccess *tfe.TeamAccess) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		tfeClient := testAccProvider.Meta().(*tfe.Client)

		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No instance ID is set")
		}

		ta, err := tfeClient.TeamAccess.Read(ctx, rs.Primary.ID)
		if err != nil {
			return err
		}

		if ta == nil {
			return fmt.Errorf("TeamAccess not found")
		}

		*tmAccess = *ta

		return nil
	}
}

func testAccCheckTFETeamAccessAttributes(
	tmAccess *tfe.TeamAccess) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if tmAccess.Access != tfe.AccessWrite {
			return fmt.Errorf("Bad access: %s", tmAccess.Access)
		}
		return nil
	}
}

func testAccCheckTFETeamAccessDestroy(s *terraform.State) error {
	tfeClient := testAccProvider.Meta().(*tfe.Client)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "scalr_team_access" {
			continue
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No instance ID is set")
		}

		_, err := tfeClient.TeamAccess.Read(ctx, rs.Primary.ID)
		if err == nil {
			return fmt.Errorf("Team access %s still exists", rs.Primary.ID)
		}
	}

	return nil
}

const testAccTFETeamAccess_basic = `
resource "scalr_organization" "foobar" {
  name  = "tst-terraform"
  email = "admin@company.com"
}

resource "scalr_team" "foobar" {
  name         = "team-test"
  organization = "${scalr_organization.foobar.id}"
}

resource "scalr_workspace" "foobar" {
  name         = "workspace-test"
  organization = "${scalr_organization.foobar.id}"
}

resource "scalr_team_access" "foobar" {
  access       = "write"
  team_id      = "${scalr_team.foobar.id}"
  workspace_id = "${scalr_workspace.foobar.id}"
}`
