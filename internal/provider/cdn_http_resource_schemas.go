package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// httpResourceResourceModel maps the data source schema data.

// TODO: make separate module with all schemas. Divide schemas by modules
type CdnHttpResourceModel struct {
	ID                 types.String `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	CreationTs         types.Int64  `tfsdk:"creation_ts"`
	CdnDomain          types.String `tfsdk:"cdn_domain"`
	Active             types.Bool   `tfsdk:"active"`
	Origin             OriginModel  `tfsdk:"origin"`
	Cache              types.Object `tfsdk:"cache"`
	Certificate        types.Int64  `tfsdk:"certificate"`
	Tuning             types.String `tfsdk:"tuning"`
	SliceSizeMegabytes types.Int64  `tfsdk:"slice_size_megabytes"`
	ModernTlsOnly      types.Bool   `tfsdk:"modern_tls_only"`
	StrongSslCiphers   types.Bool   `tfsdk:"strong_ssl_ciphers"`
	FollowRedirects    types.Bool   `tfsdk:"follow_redirects"`
	NoHttp2            types.Bool   `tfsdk:"no_http2"`
	Http2Https         types.Bool   `tfsdk:"http2https"`
	HttpsOnly          types.Bool   `tfsdk:"https_only"`
	UseHttp3           types.Bool   `tfsdk:"use_http3"`
	Compress           types.Object `tfsdk:"compress"`
	Robots             types.Object `tfsdk:"robots"`
	Auth               types.Object `tfsdk:"auth"`
	Headers            types.Object `tfsdk:"headers"`
	Cors               types.Object `tfsdk:"cors"`
	Names              types.Set    `tfsdk:"names"`
	Limitations        types.Object `tfsdk:"limitations"`
	IOSS               types.Bool   `tfsdk:"ioss"`
	Packaging          types.Object `tfsdk:"packaging"`
	Locations          types.Map    `tfsdk:"locations"`
}

type OriginModel struct {
	Servers        types.Map    `tfsdk:"servers"`
	Hostname       types.String `tfsdk:"hostname"`
	HTTPS          types.Bool   `tfsdk:"https"`
	SNIHostname    types.String `tfsdk:"sni_hostname"`
	ReadTimeout    types.String `tfsdk:"read_timeout"`
	SendTimeout    types.String `tfsdk:"send_timeout"`
	ConnectTimeout types.String `tfsdk:"connect_timeout"`
	AWS            types.Object `tfsdk:"aws"`
	S3Bucket       types.String `tfsdk:"s3_bucket"`
	SSLVerify      types.Bool   `tfsdk:"ssl_verify"`
}

func (m OriginModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"servers": types.MapType{
			ElemType: ServersModel{}.AttributeTypes(),
		},
		"hostname":        types.StringType,
		"https":           types.BoolType,
		"sni_hostname":    types.StringType,
		"read_timeout":    types.StringType,
		"send_timeout":    types.StringType,
		"connect_timeout": types.StringType,
		"aws": types.ObjectType{
			AttrTypes: AWSModel{}.AttributeTypes(),
		},
		"s3_bucket":  types.StringType,
		"ssl_verify": types.BoolType,
	}
}

type ServersModel struct{}

func (m ServersModel) AttributeTypes() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"port":      types.Int64Type,
			"weight":    types.Int64Type,
			"max_fails": types.Int64Type,
			"backup":    types.BoolType,
		},
	}
}

type AWSModel struct{}

func (m AWSModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"auth": types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"access_key": types.StringType,
				"secret_key": types.StringType,
			},
		},
	}
}

type CacheModel struct{}

func (m CacheModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"disable":       types.BoolType,
		"consider_args": types.BoolType,
		"args_whitelist": types.SetType{
			ElemType: types.StringType,
		},
		"consider_cookies": types.BoolType,
		"cookies_whitelist": types.SetType{
			ElemType: types.StringType,
		},
		"valid": types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"c_2xx": types.StringType,
				"c_3xx": types.StringType,
				"c_4xx": types.StringType,
				"c_5xx": types.StringType,
				"force": types.BoolType,
			},
		},
		"use_stale": types.BoolType,
	}
}

type CompressModel struct{}

func (m CompressModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"brotli": types.BoolType,
		"gzip":   types.BoolType,
	}
}

type RobotsModel struct{}

func (m RobotsModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"type":           types.StringType,
		"robots_content": types.StringType,
	}
}

type AuthModel struct{}

func (m AuthModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"url":       types.StringType,
		"forbidden": types.BoolType,
		"md5": types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"secret":   types.StringType,
				"forever":  types.BoolType,
				"anywhere": types.BoolType,
			},
		},
	}

}

type HeadersModel struct{}

func (m HeadersModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"request": types.MapType{
			ElemType: types.StringType,
		},
		"response": types.MapType{
			ElemType: types.StringType,
		},
		"hide_in_response": types.SetType{
			ElemType: types.StringType,
		},
	}
}

type CorsModel struct{}

func (m CorsModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"domains": types.SetType{
			ElemType: types.StringType,
		},
		"headers": types.SetType{
			ElemType: types.StringType,
		},
		"expose": types.SetType{
			ElemType: types.StringType,
		},
		"methods": types.SetType{
			ElemType: types.StringType,
		},
		"credentials": types.BoolType,
		"max_age":     types.Int64Type,
		"disable":     types.BoolType,
	}
}

type TimesModel struct{}

func (m TimesModel) AttributeTypes() attr.Type {
	return types.SetType{
		ElemType: types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"start": types.StringType,
				"end":   types.StringType,
			},
		},
	}
}

type LimitationsModel struct{}

func (m LimitationsModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"geo": types.SetType{
			ElemType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"default_action": types.StringType,
					"exclude": types.SetType{
						ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"action":  types.StringType,
								"country": types.StringType,
								"region":  types.StringType,
							},
						},
					},
					"times": TimesModel{}.AttributeTypes(),
				},
			},
		},
		"ip": types.SetType{
			ElemType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"default_action": types.StringType,
					"exclude": types.SetType{
						ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"ip": types.StringType,
							},
						},
					},
					"times": TimesModel{}.AttributeTypes(),
				},
			},
		},
		"referer": types.SetType{
			ElemType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"default_action": types.StringType,
					"exclude": types.SetType{
						ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"referer": types.StringType,
							},
						},
					},
					"times": TimesModel{}.AttributeTypes(),
				},
			},
		},
		"useragent": types.SetType{
			ElemType: types.ObjectType{
				AttrTypes: map[string]attr.Type{
					"default_action": types.StringType,
					"exclude": types.SetType{
						ElemType: types.ObjectType{
							AttrTypes: map[string]attr.Type{
								"useragent": types.StringType,
							},
						},
					},
					"times": TimesModel{}.AttributeTypes(),
				},
			},
		},
	}
}

type LocationsModel struct{}

func (m LocationsModel) AttributeTypes() attr.Type {
	return types.ObjectType{
		AttrTypes: map[string]attr.Type{
			"cache": types.ObjectType{
				AttrTypes: CacheModel{}.AttributeTypes(),
			},
			"origin": types.ObjectType{
				AttrTypes: OriginModel{}.AttributeTypes(),
			},
			"auth": types.ObjectType{
				AttrTypes: AuthModel{}.AttributeTypes(),
			},
			"headers": types.ObjectType{
				AttrTypes: HeadersModel{}.AttributeTypes(),
			},
			"cors": types.ObjectType{
				AttrTypes: CorsModel{}.AttributeTypes(),
			},
			"limitations": types.ObjectType{
				AttrTypes: LimitationsModel{}.AttributeTypes(),
			},
			"compress": types.ObjectType{
				AttrTypes: CompressModel{}.AttributeTypes(),
			},
			"ioss": types.BoolType,
			"packaging": types.ObjectType{
				AttrTypes: PackagingModel{}.AttributeTypes(),
			},
			"rewrite": types.SetType{
				ElemType: types.ObjectType{
					AttrTypes: RewriteModel{}.AttributeTypes(),
				},
			},
			"return_http_status_code": types.Int64Type,
		},
	}
}

type PackagingModel struct{}

func (m PackagingModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"mp4": types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"output_protocols": types.SetType{
					ElemType: types.StringType,
				},
			},
		},
	}
}

func (m RewriteModel) AttributeTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"from": types.StringType,
		"to":   types.StringType,
		"flag": types.StringType,
	}
}

type RewriteModel struct{}

func (d *httpResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	// TODO: Maybe use resource plan modifier
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "HTTP resource ID",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"origin": OriginSchema(true, false),
			"name": schema.StringAttribute{
				Description: "Resource name",
				Required:    true,
			},
			"creation_ts": schema.Int64Attribute{
				Description: "Timestamp of resource creation",
				Computed:    true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"cdn_domain": schema.StringAttribute{
				Description: "CDN distribution domain",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"active": schema.BoolAttribute{
				Description: "Is the resource active",
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(true),
			},
			"cache": CacheSchema(),
			"certificate": schema.Int64Attribute{
				Description: "ID of the SSL Certificate to be bound to the resource",
				Optional:    true,
			},
			"tuning": schema.StringAttribute{
				Description: "Optimization of distribution. One of [default, large, live]",
				Optional:    true,
			},
			"slice_size_megabytes": schema.Int64Attribute{
				Description: "Slice size in MB (only for tuning=large)",
				Optional:    true,
			},
			"modern_tls_only": schema.BoolAttribute{
				Description: "Use only modern versions of TLS",
				Optional:    true,
			},
			"strong_ssl_ciphers": schema.BoolAttribute{
				Description: "Use strong SSL ciphers (requires modern_tls_only=true)",
				Optional:    true,
			},
			"follow_redirects": schema.BoolAttribute{
				Description: "Follow redirects",
				Optional:    true,
			},
			"no_http2": schema.BoolAttribute{
				Description: "Disable HTTP2",
				Optional:    true,
			},
			"http2https": schema.BoolAttribute{
				Description: "Automatically redirect HTTP to HTTPS on distribution",
				Optional:    true,
			},
			"https_only": schema.BoolAttribute{
				Description: "Use only HTTPS for distribution",
				Optional:    true,
			},
			"use_http3": schema.BoolAttribute{
				Description: "Use HTTP3",
				Optional:    true,
			},
			"compress":    CompressSchema(),
			"robots":      RobotsSchema(),
			"auth":        AuthSchema(),
			"headers":     HeadersSchema(),
			"cors":        CorsSchema(),
			"names":       NamesSchema(),
			"limitations": LimitationsSchema(),
			"ioss": schema.BoolAttribute{
				Description: "Image Optimization and Modification",
				Optional:    true,
			},
			"packaging": PackagingSchema(),
			"locations": schema.MapNestedAttribute{
				Description: "Rules for specific request paths",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"cache":       CacheSchema(),
						"origin":      OriginSchema(false, true),
						"auth":        AuthSchema(),
						"headers":     HeadersSchema(),
						"cors":        CorsSchema(),
						"limitations": LimitationsSchema(),
						"compress":    CompressSchema(),
						"ioss": schema.BoolAttribute{
							Description: "Image Optimization and Modification",
							Optional:    true,
						},
						"packaging": PackagingSchema(),
						"rewrite":   RewriteSchema(),
						"return_http_status_code": schema.Int64Attribute{
							Description: "HTTP code to respond instead of content",
							Optional:    true,
						},
					},
				},
			},
		},
	}
}

func CacheSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Description: "Cache settings",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"disable": schema.BoolAttribute{
				Description: "Do not cache content",
				Optional:    true,
			},
			"consider_args": schema.BoolAttribute{
				Description: "Consider query string in caching",
				Optional:    true,
			},
			"args_whitelist": schema.SetAttribute{
				Description: "List of query string parameters to consider when caching (requires cache.consider_args=true)",
				Optional:    true,
				ElementType: types.StringType,
			},
			"consider_cookies": schema.BoolAttribute{
				Description: "Consider cookies in caching",
				Optional:    true,
			},
			"cookies_whitelist": schema.SetAttribute{
				Description: "List of cookie to consider when caching (requires cache.consider_cookies=true)",
				Optional:    true,
				ElementType: types.StringType,
			},
			"valid": schema.SingleNestedAttribute{
				Description: "Cache time settings",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"c_2xx": schema.StringAttribute{
						Description: "Cache time for 2xx codes",
						Optional:    true,
					},
					"c_3xx": schema.StringAttribute{
						Description: "Cache time for 3xx codes",
						Optional:    true,
					},
					"c_4xx": schema.StringAttribute{
						Description: "Cache time for 4xx codes",
						Optional:    true,
					},
					"c_5xx": schema.StringAttribute{
						Description: "Cache time for 5xx codes",
						Optional:    true,
					},
					"force": schema.BoolAttribute{
						Description: "Ignore cache headers",
						Optional:    true,
					},
				},
			},
			"use_stale": schema.BoolAttribute{
				Description: "Enables/disables the ability to give outdated cached content if the origin is unavailable",
				Optional:    true,
			},
		},
	}
}

func OriginSchema(required, optional bool) schema.Attribute {
	return schema.SingleNestedAttribute{
		Description: "Content source (origin) settings",
		Required:    required,
		Optional:    optional,
		Attributes: map[string]schema.Attribute{
			"servers": schema.MapNestedAttribute{
				Description: "Origins description",
				Required:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"port": schema.Int64Attribute{
							Description: "Origin port",
							Optional:    true,
						},
						"weight": schema.Int64Attribute{
							Description: "Weight for balancing",
							Optional:    true,
						},
						"max_fails": schema.Int64Attribute{
							Description: "Number of failed attempts for balancing",
							Optional:    true,
						},
						"backup": schema.BoolAttribute{
							Description: "Is origin a backup?",
							Optional:    true,
						},
					},
				},
			},
			"hostname": schema.StringAttribute{
				Description: "Host header when requesting origin",
				Optional:    true,
			},
			"https": schema.BoolAttribute{
				Description: "Whether to use HTTPS when requesting origin",
				Optional:    true,
			},
			"sni_hostname": schema.StringAttribute{
				Description: "Allows the source to understand which certificate to use for connection if the source server provides multiple certificates (requires origin.https=true)",
				Optional:    true,
			},
			"read_timeout": schema.StringAttribute{
				Description: "Read timeout in seconds",
				Optional:    true,
			},
			"send_timeout": schema.StringAttribute{
				Description: "Send timeout in seconds",
				Optional:    true,
			},
			"connect_timeout": schema.StringAttribute{
				Description: "Connect timeout in seconds",
				Optional:    true,
			},
			"aws": schema.SingleNestedAttribute{
				Description: "Parameters for using AWS authorization when requesting origin",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"auth": schema.SingleNestedAttribute{
						Description: "Authorization keys",
						Required:    true,
						Attributes: map[string]schema.Attribute{
							"access_key": schema.StringAttribute{
								Required: true,
							},
							"secret_key": schema.StringAttribute{
								Required: true,
							},
						},
					},
				},
			},
			"s3_bucket": schema.StringAttribute{
				Description: "Allowed bucket (in case of specifying a common S3 domain as origin)",
				Optional:    true,
			},
			"ssl_verify": schema.BoolAttribute{
				Description: "Should check origins certificate (requires origin.https=true)",
				Optional:    true,
			},
		},
	}
}

func CompressSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Description: "Compression settings. This service is paid according to the tariffs indicated in dashboard.",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"brotli": schema.BoolAttribute{
				Description: "Use Brotli compression",
				Optional:    true,
			},
			"gzip": schema.BoolAttribute{
				Description: "Use Gzip compression",
				Optional:    true,
			},
		},
	}
}

func RobotsSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Description: "robots.txt settings",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"type": schema.StringAttribute{
				Description: "Type of robots.txt handling. One of [deny, custom, cached]",
				Required:    true,
			},
			"robots_content": schema.StringAttribute{
				Description: "Text of robots.txt (only for type=custom)",
				Optional:    true,
			},
		},
	}
}

func AuthSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Description: "User request authorization settings. This service is paid according to the tariffs indicated in dashboard.",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"md5": schema.SingleNestedAttribute{
				Description: "Local authorization settings (based on signature)",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"secret": schema.StringAttribute{
						Description: "Secret word",
						Optional:    true,
					},
					"forever": schema.BoolAttribute{
						Description: "No time limit",
						Optional:    true,
					},
					"anywhere": schema.BoolAttribute{
						Description: "Do not consider IP address",
						Optional:    true,
					},
				},
			},
			"url": schema.StringAttribute{
				Description: "URL of external authorization script",
				Optional:    true,
			},
			"forbidden": schema.BoolAttribute{
				Description: "Deny access",
				Optional:    true,
			},
		},
	}
}

func HeadersSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Description: "Header settings",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"request": schema.MapAttribute{
				Description: "Headers for request to origin",
				Optional:    true,
				ElementType: types.StringType,
			},
			"response": schema.MapAttribute{
				Description: "Headers for response to users",
				Optional:    true,
				ElementType: types.StringType,
			},
			"hide_in_response": schema.SetAttribute{
				Description: "List of headers specified on origin that CDN servers hide in the response",
				Optional:    true,
				ElementType: types.StringType,
			},
		},
	}
}

func CorsSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Description: "CORS settings",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"domains": schema.SetAttribute{
				Description: "Allowed domains",
				Optional:    true,
				ElementType: types.StringType,
			},
			"headers": schema.SetAttribute{
				Description: "Allowed request headers. Accept, Accept-Language, Content-Type, Content-Language are allowed by default.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"expose": schema.SetAttribute{
				Description: "Headers available to top-level APIs (Expose Headers). Cache-Control, Content-Language, Content-Type, Expires, Last-Modified, Pragma headers are allowed by default.",
				Optional:    true,
				ElementType: types.StringType,
			},
			// TODO: case sense
			"methods": schema.SetAttribute{
				Description: "Allowed methods. GET, HEAD, POST are allowed by default.",
				Optional:    true,
				ElementType: types.StringType,
			},
			"credentials": schema.BoolAttribute{
				Description: "Set the Access-Control-Allow-Credentials header",
				Optional:    true,
			},
			"max_age": schema.Int64Attribute{
				Description: "Preflight request response lifetime",
				Optional:    true,
			},
			"disable": schema.BoolAttribute{
				Description: "Disable CORS",
				Optional:    true,
			},
		},
	}
}

func NamesSchema() schema.Attribute {
	return schema.SetAttribute{
		Description: "CNAMEs for CDN domain",
		Optional:    true,
		ElementType: types.StringType,
	}
}

func TimesSchema() schema.Attribute {
	return schema.SetNestedAttribute{
		Description: "Restriction intervals",
		Required:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"start": schema.StringAttribute{
					Description: "Start of interval in ISO 8601-1:2019 format",
					Required:    true,
				},
				"end": schema.StringAttribute{
					Description: "End of interval in ISO 8601-1:2019 format",
					Required:    true,
				},
			},
		},
	}
}

func LimitationsSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Description: "Restriction of distribution by geography, IP, Referer or UserAgent. This service is paid according to the tariffs indicated in dashboard",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"geo": schema.SetNestedAttribute{
				Description: "Restriction of distribution by geography",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"default_action": schema.StringAttribute{
							Description: "Default action. One of [allow, deny]",
							Required:    true,
						},
						"exclude": schema.SetNestedAttribute{
							Description: "Exclusions",
							Required:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"action": schema.StringAttribute{
										Description: "Action. One of [allow, deny]",
										Required:    true,
									},
									"country": schema.StringAttribute{
										Description: "Country code in ISO 3166-1 alpha-2 format",
										Required:    true,
									},
									"region": schema.StringAttribute{
										Description: "Region code in ISO 3166-2 format or null",
										Required:    true,
									},
								},
							},
						},
						"times": TimesSchema(),
					},
				},
			},
			"ip": schema.SetNestedAttribute{
				Description: "Restriction of distribution by IP",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"default_action": schema.StringAttribute{
							Description: "Default action. One of [allow, deny]",
							Required:    true,
						},
						"exclude": schema.SetNestedAttribute{
							Description: "Exclusions",
							Required:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"ip": schema.StringAttribute{
										Description: "IP address in CIDR notation",
										Required:    true,
									},
								},
							},
						},
						"times": TimesSchema(),
					},
				},
			},
			"referer": schema.SetNestedAttribute{
				Description: "Restriction of distribution by Referer",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"default_action": schema.StringAttribute{
							Description: "Default action. One of [allow, deny]",
							Required:    true,
						},
						"exclude": schema.SetNestedAttribute{
							Description: "Exclusions",
							Required:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"referer": schema.StringAttribute{
										Description: "Referer (domain name or regexp)",
										Required:    true,
									},
								},
							},
						},
						"times": TimesSchema(),
					},
				},
			},
			"useragent": schema.SetNestedAttribute{
				Description: "Restriction of distribution by UserAgent",
				Optional:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"default_action": schema.StringAttribute{
							Description: "Default action. One of [allow, deny]",
							Required:    true,
						},
						"exclude": schema.SetNestedAttribute{
							Description: "Exclusions",
							Required:    true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"useragent": schema.StringAttribute{
										Description: "UserAgent or regexp",
										Required:    true,
									},
								},
							},
						},
						"times": TimesSchema(),
					},
				},
			},
		},
	}
}

func PackagingSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Description: "Video Converting",
		Optional:    true,
		Attributes: map[string]schema.Attribute{
			"mp4": schema.SingleNestedAttribute{
				Description: "Conversion parameters",
				Optional:    true,
				Attributes: map[string]schema.Attribute{
					"output_protocols": schema.SetAttribute{
						Description: "Formats in which videos are planned to be distributed. One of [MPEG-DASH, HLS]",
						Required:    true,
						ElementType: types.StringType,
					},
				},
			},
		},
	}
}

func RewriteSchema() schema.Attribute {
	return schema.SetNestedAttribute{
		Description: "Rewrite options for certain accounts",
		Optional:    true,
		NestedObject: schema.NestedAttributeObject{
			Attributes: map[string]schema.Attribute{
				"from": schema.StringAttribute{
					Description: "Rewrite option",
					Optional:    true,
				},
				"to": schema.StringAttribute{
					Description: "Rewrite option",
					Optional:    true,
				},
				"flag": schema.StringAttribute{
					Description: "Rewrite option",
					Optional:    true,
				},
			},
		},
	}
}
