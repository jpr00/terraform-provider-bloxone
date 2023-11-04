package acctest

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/provider"
)

var (
	// BloxOneClient will be used to do verification tests
	BloxOneClient *bloxoneclient.APIClient

	// ProtoV6ProviderFactories are used to instantiate a provider during
	// acceptance testing. The factory function will be invoked for every Terraform
	// CLI command executed to create a provider server to which the CLI can
	// reattach.
	ProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
		"bloxone": providerserver.NewProtocol6WithError(provider.New("test", "none")()),
	}
)

func PreCheck(t *testing.T) {
	cspURL := os.Getenv("BLOXONE_CSP_URL")
	if cspURL == "" {
		t.Fatal("BLOXONE_CSP_URL must be set for acceptance tests")
	}

	apiKey := os.Getenv("BLOXONE_API_KEY")
	if apiKey == "" {
		t.Fatal("BLOXONE_API_KEY must be set for acceptance tests")
	}

	var err error
	BloxOneClient, err = bloxoneclient.NewAPIClient(bloxoneclient.Configuration{
		ClientName: "acceptance-test",
		CSPURL:     cspURL,
		APIKey:     apiKey,
	})
	if err != nil {
		t.Fatal("Cannot create bloxone client")
	}
}
