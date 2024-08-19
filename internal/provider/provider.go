package provider

import (
	"context"
	"os"

	"terraform-provider-cdnvideo/internal/configuration"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &cdnvideoProvider{}
)

// ProviderModel maps provider schema data to a Go type.
type cdnvideoProviderModel struct {
	AccountName types.String `tfsdk:"account_name"`
	Username    types.String `tfsdk:"username"`
	Password    types.String `tfsdk:"password"`
}

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &cdnvideoProvider{
			version: version,
		}
	}
}

// cdnvideoProvider is the provider implementation.
type cdnvideoProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *cdnvideoProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "cdnvideo"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *cdnvideoProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_name": schema.StringAttribute{
				Optional: true,
			},
			"username": schema.StringAttribute{
				Optional: true,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

// Configure prepares an API client for data sources and resources.
func (p *cdnvideoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring Configuration api client")

	var config cdnvideoProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.AccountName.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("account_name"),
			"Unknown CDNVideo API AccountName",
			"The provider cannot create the CDNVideo API client as there is an unknown configuration value for the CDNVideo API account_name. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the CDN_ACCOUNT_NAME environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown CDNVideo API username",
			"The provider cannot create the CDNVideo API client as there is an unknown configuration value for the CDNVideo API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the CDN_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown CDNVideo API Password",
			"The provider cannot create the CDNVideo API client as there is an unknown configuration value for the CDNVideo API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the CDN_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	account_name := os.Getenv("CDN_ACCOUNT_NAME")
	username := os.Getenv("CDN_USERNAME")
	password := os.Getenv("CDN_PASSWORD")

	if !config.AccountName.IsNull() {
		account_name = config.AccountName.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if account_name == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("account_name"),
			"Missing CDNVideo API AccountName",
			"The provider cannot create the CDNVideo API client as there is a missing or empty value for the CDNVideo API account_name. "+
				"Set the account_name value in the configuration or use the CDN_ACCOUNT_NAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing CDNVideo API Username",
			"The provider cannot create the CDNVideo API client as there is a missing or empty value for the CDNVideo API username. "+
				"Set the username value in the configuration or use the CDN_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing CDNVideo API Password",
			"The provider cannot create the CDNVideo API client as there is a missing or empty value for the CDNVideo API password. "+
				"Set the password value in the configuration or use the CDN_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "cdn_account_name", account_name)
	ctx = tflog.SetField(ctx, "cdn_username", username)
	ctx = tflog.SetField(ctx, "cdn_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "cdn_password")

	tflog.Debug(ctx, "Creating CDNVideo client")

	// Create a new CDNVideo client using the configuration values
	configuration_proxy, err := configuration.NewProxy(&username, &password, &account_name)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create CDNVideo API Client",
			"An unexpected error occurred when creating the CDNVideo API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"CDNVideo Client Error: "+err.Error(),
		)
		return
	}

	// Make the CDNVideo client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = configuration_proxy
	resp.ResourceData = configuration_proxy
	tflog.Info(ctx, "Configured success client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *cdnvideoProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *cdnvideoProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewHTTPResource,
	}
}
