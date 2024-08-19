package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestResource(t *testing.T) {
	resource_name := "cdnvideo_http.edu"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + `
				resource "cdnvideo_http" "edu" {
					origin = {
						servers = {
							"google.com" = {
								port = 443
							}
						}
					}
					name = "testname"
			
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check set options
					resource.TestCheckResourceAttr(resource_name, "origin.servers.google.com.port", "443"),
					resource.TestCheckResourceAttr(resource_name, "name", "testname"),

					// Check computed options
					resource.TestCheckResourceAttrSet(resource_name, "id"),
					resource.TestCheckResourceAttrSet(resource_name, "active"),
					resource.TestCheckResourceAttrSet(resource_name, "cdn_domain"),
					resource.TestCheckResourceAttrSet(resource_name, "creation_ts"),

					// Check all other options not set
					resource.TestCheckNoResourceAttr(resource_name, "origin.servers.google.com.weight"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.servers.google.com.max_fails"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.servers.google.com.backup"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.hostname"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.https"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.read_timeout"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.send_timeout"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.connect_timeout"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.aws"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.s3_bucket"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.ssl_verify"),
					resource.TestCheckNoResourceAttr(resource_name, "cache"),
					resource.TestCheckNoResourceAttr(resource_name, "certificate"),
					resource.TestCheckNoResourceAttr(resource_name, "tuning"),
					resource.TestCheckNoResourceAttr(resource_name, "modern_tls_only"),
					resource.TestCheckNoResourceAttr(resource_name, "strong_ssl_ciphers"),
					resource.TestCheckNoResourceAttr(resource_name, "follow_redirects"),
					resource.TestCheckNoResourceAttr(resource_name, "no_http2"),
					resource.TestCheckNoResourceAttr(resource_name, "http2https"),
					resource.TestCheckNoResourceAttr(resource_name, "https_only"),
					resource.TestCheckNoResourceAttr(resource_name, "use_http3"),
					resource.TestCheckNoResourceAttr(resource_name, "compress"),
					resource.TestCheckNoResourceAttr(resource_name, "robots"),
					resource.TestCheckNoResourceAttr(resource_name, "auth"),
					resource.TestCheckNoResourceAttr(resource_name, "headers"),
					resource.TestCheckNoResourceAttr(resource_name, "cors"),
					resource.TestCheckNoResourceAttr(resource_name, "names"),
					resource.TestCheckNoResourceAttr(resource_name, "limitations"),
					resource.TestCheckNoResourceAttr(resource_name, "ioss"),
					resource.TestCheckNoResourceAttr(resource_name, "packaging"),
					resource.TestCheckNoResourceAttr(resource_name, "locations"),
				),
			},
			// Check full configuration
			{
				Config: providerConfig + `
				resource "cdnvideo_http" "edu" {
					origin = {
						servers = {
							"google.com" = {
								port = 443
								weight = 1
								max_fails = 10
								backup = false
							}
							"storage.yandexcloud.net" = {}
						}
						hostname = "string"
						https = true
						sni_hostname    = "custom-host.com"
						read_timeout = "10s"
						send_timeout = "10s"
						connect_timeout = "10s"
						aws = {
							auth = {
								access_key = "string"
								secret_key = "string"
							}
						}
						s3_bucket = "string"
						ssl_verify = false
					}
					name = "testname"
					active = true
					cache = {
						disable = false
						consider_args = true
						args_whitelist = [
							"param1"
						]
						consider_cookies = true
						cookies_whitelist = [
							"param1"
						]		
						valid = {
							c_2xx = "1d"
							c_3xx = "1d"
							c_4xx = "1s"
							c_5xx = "1s"
							force = false
						}
						use_stale = false
					}
					certificate = 1
					tuning = "default"
					modern_tls_only = false
					strong_ssl_ciphers = false
					follow_redirects = false
					no_http2 = false
					http2https = false
					https_only = false
					use_http3 = false
					compress = {
						brotli = false
						gzip = true
					}
					robots = {
						type = "deny"
					}
					auth = {
						md5 = {
							secret = "string"
							forever = false
							anywhere = false
						}
					}
					headers = {
						request = {
							header_name = "header_value"
						}
						response = {
							header_name = "header_value"
						}
						hide_in_response = [
							"header-to-hide"
						]				  
					}
					cors = {
						domains = [
							"example.com"
						]
						headers = [
							"string"
						]
						expose = [
							"string"
						]
						methods = [
							"STRING"
						]
						credentials = true
						max_age = 120
						disable = false
					}
					names = [
						"cdn.test.com"
					]
					limitations = {
						geo = [
							{
								default_action = "allow"
								exclude = [
									{
										action = "deny"
										country = "RU"
										region = "BEL"
									}
								]
								times = [
									{
										start = "2024-01-01T00:00:00Z"
										end = "2024-01-02T00:00:00Z"
									}
								]
							}
						]
						ip = [
							{
								default_action = "allow"
								exclude = [
									{
										ip = "192.168.0.1/24"
									}
								]
								times = [
									{
										start = "2024-01-01T00:00:00Z"
										end = "2024-01-02T00:00:00Z"
									}
								]
							}
						]
						referer = [
							{
								default_action = "allow"
								exclude = [
									{
										referer = "*.ru"
									}
								]
								times = [
									{
										start = "2024-01-01T00:00:00Z"
										end = "2024-01-02T00:00:00Z"
									}
								]
							}
						]
						useragent = [
							{
								default_action = "allow"
								exclude = [
									{
										useragent = "browser_name"
									}
								]
								times = [
									{
										start = "2024-01-01T00:00:00Z"
										end = "2024-01-02T00:00:00Z"
									}
								]
							}
						]
					}
					ioss = false
					packaging = {
						mp4 = {
							output_protocols = [
								"MPEG-DASH"
							]
						}
					}
					locations = {
						"path_to_content" = {
							cache = {
								disable = false
								consider_args = true
						        args_whitelist = [
									"param1"
								]
								consider_cookies = true
								cookies_whitelist = [
									"param1"
								]
								valid = {
									c_2xx = "1d"
									c_3xx = "1d"
									c_4xx = "1s"
									c_5xx = "1s"
									force = false
								}
								use_stale = false
							}
							origin = {
								servers = {
									"google.com" = {
										port      = 443
										weight    = 1
										max_fails = 10
										backup    = false
									}
									"storage.yandexcloud.net" = {}
								}
								hostname        = "string"
								https           = true
								sni_hostname    = "custom-host.com"
								read_timeout    = "10s"
								send_timeout    = "10s"
								connect_timeout = "10s"
								aws = {
									auth = {
										access_key = "string"
										secret_key = "string"
									}
								}
								s3_bucket  = "string"
								ssl_verify = false
							}
							auth = {
								md5 = {
									secret = "string"
									forever = false
									anywhere = false
								}
							}
							headers = {
								request = {
									header_name = "header_value"
								}
								response = {
									header_name = "header_value"
								}
								hide_in_response = [
									"header-to-hide"
								]						  
							}
							cors = {
								domains = [
									"example.com"
								]
								headers = [
									"string"
								]
								expose = [
									"string"
								]
								methods = [
									"STRING"
								]
								credentials = true
								max_age = 120
								disable = false
							}
							limitations = {
								geo = [
									{
										default_action = "allow"
										exclude = [
											{
												action = "deny"
												country = "RU"
												region = "BEL"
											}
										]
										times = [
											{
												start = "2024-01-01T00:00:00Z"
												end = "2024-01-02T00:00:00Z"
											}
										]
									}
								]
								ip = [
									{
										default_action = "allow"
										exclude = [
											{
												ip = "192.168.0.1/24"
											}
										]
										times = [
											{
												start = "2024-01-01T00:00:00Z"
												end = "2024-01-02T00:00:00Z"
											}
										]
									}
								]
								referer = [
									{
										default_action = "allow"
										exclude = [
											{
												referer = "*.ru"
											}
										]
										times = [
											{
												start = "2024-01-01T00:00:00Z"
												end = "2024-01-02T00:00:00Z"
											}
										]
									}
								]
								useragent = [
									{
										default_action = "allow"
										exclude = [
											{
												useragent = "browser_name"
											}
										]
										times = [
											{
												start = "2024-01-01T00:00:00Z"
												end = "2024-01-02T00:00:00Z"
											}
										]
									}
								]
							}
							compress = {
								brotli = false
								gzip = true
							}	
							ioss = false
							packaging = {
								mp4 = {
									output_protocols = [
										"MPEG-DASH"
									]
								}
							}
							rewrite = [
								{
									from = "^/cdn/.+(/_video_.+)"
									to = "$1"
									flag = "break" 
								}
							]
							return_http_status_code = 403
						}
					}
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check set options
					resource.TestCheckResourceAttr(resource_name, "origin.servers.google.com.port", "443"),
					resource.TestCheckResourceAttr(resource_name, "origin.servers.google.com.weight", "1"),
					resource.TestCheckResourceAttr(resource_name, "origin.servers.google.com.max_fails", "10"),
					resource.TestCheckResourceAttr(resource_name, "origin.servers.google.com.backup", "false"),
					resource.TestCheckResourceAttrSet(resource_name, "origin.servers.storage.yandexcloud.net.%"),
					resource.TestCheckResourceAttr(resource_name, "origin.hostname", "string"),
					resource.TestCheckResourceAttr(resource_name, "origin.https", "true"),
					resource.TestCheckResourceAttr(resource_name, "origin.sni_hostname", "custom-host.com"),
					resource.TestCheckResourceAttr(resource_name, "origin.read_timeout", "10s"),
					resource.TestCheckResourceAttr(resource_name, "origin.send_timeout", "10s"),
					resource.TestCheckResourceAttr(resource_name, "origin.connect_timeout", "10s"),
					resource.TestCheckResourceAttr(resource_name, "origin.aws.auth.access_key", "string"),
					resource.TestCheckResourceAttr(resource_name, "origin.aws.auth.secret_key", "string"),
					resource.TestCheckResourceAttr(resource_name, "origin.s3_bucket", "string"),
					resource.TestCheckResourceAttr(resource_name, "origin.ssl_verify", "false"),
					resource.TestCheckResourceAttr(resource_name, "name", "testname"),
					resource.TestCheckResourceAttr(resource_name, "cache.disable", "false"),
					resource.TestCheckResourceAttr(resource_name, "cache.consider_args", "true"),
					resource.TestCheckResourceAttr(resource_name, "cache.args_whitelist.0", "param1"),
					resource.TestCheckResourceAttr(resource_name, "cache.consider_cookies", "true"),
					resource.TestCheckResourceAttr(resource_name, "cache.cookies_whitelist.0", "param1"),
					resource.TestCheckResourceAttr(resource_name, "cache.use_stale", "false"),
					resource.TestCheckResourceAttr(resource_name, "cache.valid.c_2xx", "1d"),
					resource.TestCheckResourceAttr(resource_name, "cache.valid.c_3xx", "1d"),
					resource.TestCheckResourceAttr(resource_name, "cache.valid.c_4xx", "1s"),
					resource.TestCheckResourceAttr(resource_name, "cache.valid.c_5xx", "1s"),
					resource.TestCheckResourceAttr(resource_name, "cache.valid.force", "false"),
					resource.TestCheckResourceAttr(resource_name, "certificate", "1"),
					resource.TestCheckResourceAttr(resource_name, "tuning", "default"),
					resource.TestCheckResourceAttr(resource_name, "modern_tls_only", "false"),
					resource.TestCheckResourceAttr(resource_name, "strong_ssl_ciphers", "false"),
					resource.TestCheckResourceAttr(resource_name, "follow_redirects", "false"),
					resource.TestCheckResourceAttr(resource_name, "no_http2", "false"),
					resource.TestCheckResourceAttr(resource_name, "http2https", "false"),
					resource.TestCheckResourceAttr(resource_name, "https_only", "false"),
					resource.TestCheckResourceAttr(resource_name, "use_http3", "false"),
					resource.TestCheckResourceAttr(resource_name, "compress.brotli", "false"),
					resource.TestCheckResourceAttr(resource_name, "compress.gzip", "true"),
					resource.TestCheckResourceAttr(resource_name, "robots.type", "deny"),
					resource.TestCheckResourceAttr(resource_name, "auth.md5.secret", "string"),
					resource.TestCheckResourceAttr(resource_name, "auth.md5.forever", "false"),
					resource.TestCheckResourceAttr(resource_name, "auth.md5.anywhere", "false"),
					resource.TestCheckResourceAttr(resource_name, "headers.request.header_name", "header_value"),
					resource.TestCheckResourceAttr(resource_name, "headers.response.header_name", "header_value"),
					resource.TestCheckResourceAttr(resource_name, "headers.hide_in_response.0", "header-to-hide"),
					resource.TestCheckResourceAttr(resource_name, "cors.domains.0", "example.com"),
					resource.TestCheckResourceAttr(resource_name, "cors.headers.0", "string"),
					resource.TestCheckResourceAttr(resource_name, "cors.expose.0", "string"),
					resource.TestCheckResourceAttr(resource_name, "cors.methods.0", "STRING"),
					resource.TestCheckResourceAttr(resource_name, "cors.credentials", "true"),
					resource.TestCheckResourceAttr(resource_name, "cors.max_age", "120"),
					resource.TestCheckResourceAttr(resource_name, "cors.disable", "false"),
					resource.TestCheckResourceAttr(resource_name, "names.0", "cdn.test.com"),
					resource.TestCheckResourceAttr(resource_name, "limitations.geo.0.default_action", "allow"),
					resource.TestCheckResourceAttr(resource_name, "limitations.geo.0.exclude.0.action", "deny"),
					resource.TestCheckResourceAttr(resource_name, "limitations.geo.0.exclude.0.country", "RU"),
					resource.TestCheckResourceAttr(resource_name, "limitations.geo.0.exclude.0.region", "BEL"),
					resource.TestCheckResourceAttr(resource_name, "limitations.geo.0.times.0.start", "2024-01-01T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "limitations.geo.0.times.0.end", "2024-01-02T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "limitations.ip.0.default_action", "allow"),
					resource.TestCheckResourceAttr(resource_name, "limitations.ip.0.exclude.0.ip", "192.168.0.1/24"),
					resource.TestCheckResourceAttr(resource_name, "limitations.ip.0.times.0.start", "2024-01-01T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "limitations.ip.0.times.0.end", "2024-01-02T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "limitations.referer.0.default_action", "allow"),
					resource.TestCheckResourceAttr(resource_name, "limitations.referer.0.exclude.0.referer", "*.ru"),
					resource.TestCheckResourceAttr(resource_name, "limitations.referer.0.times.0.start", "2024-01-01T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "limitations.referer.0.times.0.end", "2024-01-02T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "limitations.useragent.0.default_action", "allow"),
					resource.TestCheckResourceAttr(resource_name, "limitations.useragent.0.exclude.0.useragent", "browser_name"),
					resource.TestCheckResourceAttr(resource_name, "limitations.useragent.0.times.0.start", "2024-01-01T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "limitations.useragent.0.times.0.end", "2024-01-02T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "ioss", "false"),
					resource.TestCheckResourceAttr(resource_name, "packaging.mp4.output_protocols.0", "MPEG-DASH"),

					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cache.disable", "false"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cache.consider_args", "true"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cache.args_whitelist.0", "param1"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cache.consider_cookies", "true"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cache.cookies_whitelist.0", "param1"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cache.use_stale", "false"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cache.valid.c_2xx", "1d"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cache.valid.c_3xx", "1d"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cache.valid.c_4xx", "1s"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cache.valid.c_5xx", "1s"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cache.valid.force", "false"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.servers.google.com.port", "443"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.servers.google.com.weight", "1"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.servers.google.com.max_fails", "10"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.servers.google.com.backup", "false"),
					resource.TestCheckResourceAttrSet(resource_name, "locations.path_to_content.origin.servers.storage.yandexcloud.net.%"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.hostname", "string"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.https", "true"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.sni_hostname", "custom-host.com"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.read_timeout", "10s"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.send_timeout", "10s"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.connect_timeout", "10s"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.aws.auth.access_key", "string"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.aws.auth.secret_key", "string"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.s3_bucket", "string"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.origin.ssl_verify", "false"),

					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.auth.md5.secret", "string"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.auth.md5.forever", "false"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.auth.md5.anywhere", "false"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.headers.request.header_name", "header_value"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.headers.response.header_name", "header_value"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.headers.hide_in_response.0", "header-to-hide"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cors.domains.0", "example.com"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cors.headers.0", "string"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cors.expose.0", "string"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cors.methods.0", "STRING"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cors.credentials", "true"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cors.max_age", "120"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.cors.disable", "false"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.geo.0.default_action", "allow"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.geo.0.exclude.0.action", "deny"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.geo.0.exclude.0.country", "RU"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.geo.0.exclude.0.region", "BEL"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.geo.0.times.0.start", "2024-01-01T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.geo.0.times.0.end", "2024-01-02T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.ip.0.default_action", "allow"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.ip.0.exclude.0.ip", "192.168.0.1/24"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.ip.0.times.0.start", "2024-01-01T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.ip.0.times.0.end", "2024-01-02T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.referer.0.default_action", "allow"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.referer.0.exclude.0.referer", "*.ru"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.referer.0.times.0.start", "2024-01-01T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.referer.0.times.0.end", "2024-01-02T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.useragent.0.default_action", "allow"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.useragent.0.exclude.0.useragent", "browser_name"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.useragent.0.times.0.start", "2024-01-01T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.limitations.useragent.0.times.0.end", "2024-01-02T00:00:00Z"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.ioss", "false"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.packaging.mp4.output_protocols.0", "MPEG-DASH"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.rewrite.0.from", "^/cdn/.+(/_video_.+)"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.rewrite.0.to", "$1"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.rewrite.0.flag", "break"),
					resource.TestCheckResourceAttr(resource_name, "locations.path_to_content.return_http_status_code", "403"),

					// Check computed options
					resource.TestCheckResourceAttrSet(resource_name, "id"),
					resource.TestCheckResourceAttrSet(resource_name, "active"),
					resource.TestCheckResourceAttrSet(resource_name, "cdn_domain"),
					resource.TestCheckResourceAttrSet(resource_name, "creation_ts"),
				),
			},
			// Check remove full configuration
			{
				Config: providerConfig + `
				resource "cdnvideo_http" "edu" {
					origin = {
						servers = {
							"google.com" = {
								port = 443
							}
						}
					}
					name = "testname"
			
				}`,
				Check: resource.ComposeAggregateTestCheckFunc(
					// Check set options
					resource.TestCheckResourceAttr(resource_name, "origin.servers.google.com.port", "443"),
					resource.TestCheckResourceAttr(resource_name, "name", "testname"),

					// Check computed options
					resource.TestCheckResourceAttrSet(resource_name, "id"),
					resource.TestCheckResourceAttrSet(resource_name, "active"),
					resource.TestCheckResourceAttrSet(resource_name, "cdn_domain"),
					resource.TestCheckResourceAttrSet(resource_name, "creation_ts"),

					// Check all other options not set
					resource.TestCheckNoResourceAttr(resource_name, "origin.servers.google.com.weight"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.servers.google.com.max_fails"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.servers.google.com.backup"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.hostname"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.https"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.read_timeout"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.send_timeout"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.connect_timeout"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.aws"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.s3_bucket"),
					resource.TestCheckNoResourceAttr(resource_name, "origin.ssl_verify"),
					resource.TestCheckNoResourceAttr(resource_name, "cache"),
					resource.TestCheckNoResourceAttr(resource_name, "certificate"),
					resource.TestCheckNoResourceAttr(resource_name, "tuning"),
					resource.TestCheckNoResourceAttr(resource_name, "modern_tls_only"),
					resource.TestCheckNoResourceAttr(resource_name, "strong_ssl_ciphers"),
					resource.TestCheckNoResourceAttr(resource_name, "follow_redirects"),
					resource.TestCheckNoResourceAttr(resource_name, "no_http2"),
					resource.TestCheckNoResourceAttr(resource_name, "http2https"),
					resource.TestCheckNoResourceAttr(resource_name, "https_only"),
					resource.TestCheckNoResourceAttr(resource_name, "use_http3"),
					resource.TestCheckNoResourceAttr(resource_name, "compress"),
					resource.TestCheckNoResourceAttr(resource_name, "robots"),
					resource.TestCheckNoResourceAttr(resource_name, "auth"),
					resource.TestCheckNoResourceAttr(resource_name, "headers"),
					resource.TestCheckNoResourceAttr(resource_name, "cors"),
					resource.TestCheckNoResourceAttr(resource_name, "names"),
					resource.TestCheckNoResourceAttr(resource_name, "limitations"),
					resource.TestCheckNoResourceAttr(resource_name, "ioss"),
					resource.TestCheckNoResourceAttr(resource_name, "packaging"),
					resource.TestCheckNoResourceAttr(resource_name, "locations"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}
