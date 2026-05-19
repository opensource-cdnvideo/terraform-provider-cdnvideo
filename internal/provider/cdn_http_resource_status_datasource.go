package provider

import (
	"context"
	"fmt"

	"terraform-provider-cdnvideo/internal/configuration"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &httpResourceStatusDataSource{}
	_ datasource.DataSourceWithConfigure = &httpResourceStatusDataSource{}
)

func NewHTTPResourceStatusDataSource() datasource.DataSource {
	return &httpResourceStatusDataSource{}
}

type httpResourceStatusDataSource struct {
	proxy *configuration.ConfigurationApiProxy
}

type httpResourceStatusModel struct {
	// Input
	ResourceID types.String `tfsdk:"resource_id"`
	// Computed
	Status  types.String `tfsdk:"status"`
	Message types.String `tfsdk:"message"`
}

func (d *httpResourceStatusDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_http_status"
}

func (d *httpResourceStatusDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches the configuration distribution status of a CDN HTTP resource. " +
			"Use this data source inside a check block to verify that settings have been fully " +
			"propagated across the CDN after a resource is created or updated.",
		Attributes: map[string]schema.Attribute{
			"resource_id": schema.StringAttribute{
				Description: "The ID of the CDN HTTP resource to check status for.",
				Required:    true,
			},
			"status": schema.StringAttribute{
				Description: "Current distribution status: Completed (active), Processing (applying settings), Error.",
				Computed:    true,
			},
			"message": schema.StringAttribute{
				Description: "Human-readable status description, e.g. 'dns processing' or 'configuration processing'.",
				Computed:    true,
			},
		},
	}
}

func (d *httpResourceStatusDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	proxy, ok := req.ProviderData.(*configuration.ConfigurationApiProxy)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *configuration.ConfigurationApiProxy, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.proxy = proxy
}

func (d *httpResourceStatusDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state httpResourceStatusModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	resourceStatus, err := d.proxy.GetHttpResourceStatus(state.ResourceID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error fetching CDN HTTP resource status",
			"Could not fetch status for resource "+state.ResourceID.ValueString()+": "+err.Error(),
		)
		return
	}

	state.Status = types.StringValue(resourceStatus.Status)
	state.Message = types.StringValue(resourceStatus.Message)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
