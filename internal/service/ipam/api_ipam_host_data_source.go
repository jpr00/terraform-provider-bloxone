package ipam

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	bloxoneclient "github.com/infobloxopen/bloxone-go-client/client"
	"github.com/infobloxopen/bloxone-go-client/ipam"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/flex"
	"github.com/infobloxopen/terraform-provider-bloxone/internal/utils"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &IpamHostDataSource{}

func NewIpamHostDataSource() datasource.DataSource {
	return &IpamHostDataSource{}
}

// IpamHostDataSource defines the data source implementation.
type IpamHostDataSource struct {
	client *bloxoneclient.APIClient
}

func (d *IpamHostDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_" + "ipam_hosts"
}

type IpamsvcIpamHostModelWithFilter struct {
	Filters    types.Map  `tfsdk:"filters"`
	TagFilters types.Map  `tfsdk:"tag_filters"`
	Results    types.List `tfsdk:"results"`
}

func (m *IpamsvcIpamHostModelWithFilter) FlattenResults(ctx context.Context, from []ipam.IpamsvcIpamHost, diags *diag.Diagnostics) {
	if len(from) == 0 {
		return
	}
	m.Results = flex.FlattenFrameworkListNestedBlock(ctx, from, IpamsvcIpamHostAttrTypes, diags, FlattenIpamsvcIpamHost)
}

func (d *IpamHostDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Retrieves information about existing IPAM Hosts.\n\nThe IPAM host object represents any network connected equipment that is assigned one or more IP Addresses.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.MapAttribute{
				Description: "Filter are used to return a more specific list of results. Filters can be used to match resources by specific attributes, e.g. name. If you specify multiple filters, the results returned will have only resources that match all the specified filters.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"tag_filters": schema.MapAttribute{
				Description: "Tag Filters are used to return a more specific list of results filtered by tags. If you specify multiple filters, the results returned will have only resources that match all the specified filters.",
				ElementType: types.StringType,
				Optional:    true,
			},
			"results": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: utils.DataSourceAttributeMap(IpamsvcIpamHostResourceSchemaAttributes, &resp.Diagnostics),
				},
				Computed: true,
			},
		},
	}
}

func (d *IpamHostDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*bloxoneclient.APIClient)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *bloxoneclient.APIClient, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

func (d *IpamHostDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data IpamsvcIpamHostModelWithFilter

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	allResults, err := utils.ReadWithPages(func(offset, limit int32) ([]ipam.IpamsvcIpamHost, error) {
		apiRes, _, err := d.client.IPAddressManagementAPI.
			IpamHostAPI.
			IpamHostList(ctx).
			Filter(flex.ExpandFrameworkMapFilterString(ctx, data.Filters, &resp.Diagnostics)).
			Tfilter(flex.ExpandFrameworkMapFilterString(ctx, data.TagFilters, &resp.Diagnostics)).
			Offset(offset).
			Limit(limit).
			Execute()
		if err != nil {
			resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read IpamHost, got error: %s", err))
			return nil, err
		}
		return apiRes.GetResults(), nil
	})
	if err != nil {
		return
	}

	data.FlattenResults(ctx, allResults, &resp.Diagnostics)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
