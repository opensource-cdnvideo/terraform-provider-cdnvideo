package provider

import (
	"context"
	"fmt"
	"terraform-provider-cdnvideo/internal/configuration"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &httpResource{}
	_ resource.ResourceWithConfigure = &httpResource{}
)

func NewHTTPResource() resource.Resource {
	return &httpResource{}
}

type httpResource struct {
	proxy *configuration.ConfigurationApiProxy
}

func (d *httpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_http"
}

// Create a new resource.
func (resource *httpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan CdnHttpResourceModel
	diag := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diag...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request from plan
	http_resource_request, diags := GenerateApiRequest(plan, ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Create new cdn http resource
	response, err := resource.proxy.CreateHttpResource(http_resource_request)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating cdn http resource",
			"Could not create cdn http resource, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Debug(ctx, "Created http resource")

	http_resource, err := resource.proxy.GetHttpResource(response.ResourceId)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting cdn http resource",
			"Could not getting cdn http resource, unexpected error: "+err.Error(),
		)
		return
	}

	// Map response body to schema and populate Computed attribute values
	plan.ID = types.StringValue(http_resource.ID)
	plan.Active = types.BoolPointerValue(http_resource.Active)
	plan.CreationTs = types.Int64Value(http_resource.CreationTs)
	plan.CdnDomain = types.StringValue(http_resource.CdnDomain)

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (resource *httpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var resource_id string
	diags := resp.State.GetAttribute(ctx, path.Root("id"), &resource_id)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	http_resource, err := resource.proxy.GetHttpResource(resource_id)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read cdn http resource",
			err.Error(),
		)
		return
	}
	tflog.Debug(ctx, "Successfully Read cdn http resource")

	// Map response body to model
	state, diags := GenerateState(http_resource, ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	tflog.Debug(ctx, "Successfully transfer response to model")

	// Set state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (resource *httpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CdnHttpResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	http_resource_request, diags := GenerateApiRequest(plan, ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	_, err := resource.proxy.UpdateHttpResource(http_resource_request, plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating cdn http resource",
			"Could not update cdn http resource, unexpected error: "+err.Error(),
		)
		return
	}

	http_resource, err := resource.proxy.GetHttpResource(plan.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error getting cdn http resource",
			"Could not get cdn http resource, unexpected error: "+err.Error(),
		)
		return
	}

	// Update resource state
	state, diags := GenerateState(http_resource, ctx)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (resource *httpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CdnHttpResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := resource.proxy.DeactivateHttpResource(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting cdn http resource",
			"Could not delete cdn http resource, unexpected error: "+err.Error(),
		)
		return
	}
}

// Configure adds the provider configured client to the resource.
func (resource *httpResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	proxy, ok := req.ProviderData.(*configuration.ConfigurationApiProxy)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *configuration.ConfigurationApiProxy, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	resource.proxy = proxy
}

func GenerateState(http_resource configuration.CdnHttpResource, ctx context.Context) (CdnHttpResourceModel, diag.Diagnostics) {
	servers, all_diags := types.MapValueFrom(ctx, ServersModel{}.AttributeTypes(), http_resource.Origin.Servers)

	locations, diags := types.MapValueFrom(ctx, LocationsModel{}.AttributeTypes(), http_resource.Locations)
	all_diags.Append(diags...)

	aws, diags := types.ObjectValueFrom(ctx, AWSModel{}.AttributeTypes(), http_resource.Origin.AWS)
	all_diags.Append(diags...)

	cache, diags := types.ObjectValueFrom(ctx, CacheModel{}.AttributeTypes(), http_resource.Cache)
	all_diags.Append(diags...)

	compress, diags := types.ObjectValueFrom(ctx, CompressModel{}.AttributeTypes(), http_resource.Compress)
	all_diags.Append(diags...)

	robots, diags := types.ObjectValueFrom(ctx, RobotsModel{}.AttributeTypes(), http_resource.Robots)
	all_diags.Append(diags...)

	auth, diags := types.ObjectValueFrom(ctx, AuthModel{}.AttributeTypes(), http_resource.Auth)
	all_diags.Append(diags...)

	headers, diags := types.ObjectValueFrom(ctx, HeadersModel{}.AttributeTypes(), http_resource.Headers)
	all_diags.Append(diags...)

	cors, diags := types.ObjectValueFrom(ctx, CorsModel{}.AttributeTypes(), http_resource.Cors)
	all_diags.Append(diags...)

	names, diags := types.SetValueFrom(ctx, types.StringType, http_resource.Names)
	all_diags.Append(diags...)

	limitations, diags := types.ObjectValueFrom(ctx, LimitationsModel{}.AttributeTypes(), http_resource.Limitations)
	all_diags.Append(diags...)

	packaging, diags := types.ObjectValueFrom(ctx, PackagingModel{}.AttributeTypes(), http_resource.Packaging)
	all_diags.Append(diags...)

	state := CdnHttpResourceModel{
		ID:         types.StringValue(http_resource.ID),
		Name:       types.StringValue(http_resource.Name),
		CdnDomain:  types.StringValue(http_resource.CdnDomain),
		Active:     types.BoolPointerValue(http_resource.Active),
		CreationTs: types.Int64Value(http_resource.CreationTs),
		Origin: OriginModel{
			Servers:        servers,
			Hostname:       types.StringPointerValue(http_resource.Origin.Hostname),
			SNIHostname:    types.StringPointerValue(http_resource.Origin.SNIHostname),
			HTTPS:          types.BoolPointerValue(http_resource.Origin.HTTPS),
			ReadTimeout:    types.StringPointerValue(http_resource.Origin.ReadTimeout),
			SendTimeout:    types.StringPointerValue(http_resource.Origin.SendTimeout),
			ConnectTimeout: types.StringPointerValue(http_resource.Origin.ConnectTimeout),
			AWS:            aws,
			S3Bucket:       types.StringPointerValue(http_resource.Origin.S3Bucket),
			SSLVerify:      types.BoolPointerValue(http_resource.Origin.SSLVerify),
		},
		Cache:              cache,
		Certificate:        types.Int64PointerValue(http_resource.Certificate),
		Tuning:             types.StringPointerValue(http_resource.Tuning),
		SliceSizeMegabytes: types.Int64PointerValue(http_resource.SliceSizeMegabytes),
		ModernTlsOnly:      types.BoolPointerValue(http_resource.ModernTlsOnly),
		StrongSslCiphers:   types.BoolPointerValue(http_resource.StrongSslCiphers),
		FollowRedirects:    types.BoolPointerValue(http_resource.FollowRedirects),
		NoHttp2:            types.BoolPointerValue(http_resource.NoHttp2),
		Http2Https:         types.BoolPointerValue(http_resource.Http2Https),
		HttpsOnly:          types.BoolPointerValue(http_resource.HttpsOnly),
		UseHttp3:           types.BoolPointerValue(http_resource.UseHttp3),
		Compress:           compress,
		Robots:             robots,
		Auth:               auth,
		Headers:            headers,
		Cors:               cors,
		Names:              names,
		Limitations:        limitations,
		IOSS:               types.BoolPointerValue(http_resource.IOSS),
		Packaging:          packaging,
		Locations:          locations,
	}

	return state, all_diags
}

func GenerateApiRequest(plan CdnHttpResourceModel, ctx context.Context) (configuration.CdnHttpResource, diag.Diagnostics) {
	opts := basetypes.ObjectAsOptions{UnhandledNullAsEmpty: true}

	// TODO: Check some strange pointer
	var aws *configuration.AWS = new(configuration.AWS)
	all_diags := plan.Origin.AWS.As(ctx, &aws, opts)

	servers := make(map[string]configuration.Servers)
	diags := plan.Origin.Servers.ElementsAs(ctx, &servers, false)
	all_diags.Append(diags...)

	var cache *configuration.Cache = new(configuration.Cache)
	diags = plan.Cache.As(ctx, &cache, opts)
	all_diags.Append(diags...)

	var compress *configuration.Compress = new(configuration.Compress)
	diags = plan.Compress.As(ctx, &compress, opts)
	all_diags.Append(diags...)

	var robots *configuration.Robots = new(configuration.Robots)
	diags = plan.Robots.As(ctx, &robots, opts)
	all_diags.Append(diags...)

	var auth *configuration.Auth = new(configuration.Auth)
	diags = plan.Auth.As(ctx, &auth, opts)
	all_diags.Append(diags...)

	var headers *configuration.Headers = new(configuration.Headers)
	diags = plan.Headers.As(ctx, &headers, opts)
	all_diags.Append(diags...)

	var cors *configuration.Cors = new(configuration.Cors)
	diags = plan.Cors.As(ctx, &cors, opts)
	all_diags.Append(diags...)

	names := make([]string, 0)
	diags = plan.Names.ElementsAs(ctx, &names, false)
	all_diags.Append(diags...)

	var limitations *configuration.Limitations = new(configuration.Limitations)
	diags = plan.Limitations.As(ctx, &limitations, opts)
	all_diags.Append(diags...)

	locations := make(map[string]configuration.Locations)
	diags = plan.Locations.ElementsAs(ctx, &locations, false)
	all_diags.Append(diags...)

	var packaging *configuration.Packaging = new(configuration.Packaging)
	diags = plan.Packaging.As(ctx, &packaging, opts)
	all_diags.Append(diags...)

	tflog.Info(ctx, "Parsed servers")
	http_resource_request := configuration.CdnHttpResource{
		Name: plan.Name.ValueString(),
		Origin: &configuration.Origin{
			Servers:        servers,
			Hostname:       plan.Origin.Hostname.ValueStringPointer(),
			SNIHostname:    plan.Origin.SNIHostname.ValueStringPointer(),
			HTTPS:          plan.Origin.HTTPS.ValueBoolPointer(),
			ReadTimeout:    plan.Origin.ReadTimeout.ValueStringPointer(),
			SendTimeout:    plan.Origin.SendTimeout.ValueStringPointer(),
			ConnectTimeout: plan.Origin.ConnectTimeout.ValueStringPointer(),
			AWS:            aws,
			S3Bucket:       plan.Origin.S3Bucket.ValueStringPointer(),
			SSLVerify:      plan.Origin.SSLVerify.ValueBoolPointer(),
		},
		Cache:              cache,
		Certificate:        plan.Certificate.ValueInt64Pointer(),
		Tuning:             plan.Tuning.ValueStringPointer(),
		SliceSizeMegabytes: plan.SliceSizeMegabytes.ValueInt64Pointer(),
		ModernTlsOnly:      plan.ModernTlsOnly.ValueBoolPointer(),
		StrongSslCiphers:   plan.StrongSslCiphers.ValueBoolPointer(),
		FollowRedirects:    plan.FollowRedirects.ValueBoolPointer(),
		NoHttp2:            plan.NoHttp2.ValueBoolPointer(),
		Http2Https:         plan.Http2Https.ValueBoolPointer(),
		HttpsOnly:          plan.HttpsOnly.ValueBoolPointer(),
		UseHttp3:           plan.UseHttp3.ValueBoolPointer(),
		Compress:           compress,
		Robots:             robots,
		Auth:               auth,
		Headers:            headers,
		Cors:               cors,
		Names:              names,
		Limitations:        limitations,
		IOSS:               plan.IOSS.ValueBoolPointer(),
		Packaging:          packaging,
		Locations:          locations,
	}
	return http_resource_request, all_diags
}
