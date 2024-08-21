terraform {
  required_providers {
    cdnvideo = {
      source = "opensource-cdnvideo/cdnvideo"
    }
  }
}

provider "cdnvideo" {
  account_name = "account_name"
  username     = "example@example.ru"
  password     = "password"
}

resource "cdnvideo_http" "edu" {
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
  name   = "testname"
  active = true
  cache = {
    disable       = false
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
  certificate        = 1
  tuning             = "default"
  modern_tls_only    = false
  strong_ssl_ciphers = false
  follow_redirects   = false
  no_http2           = false
  http2https         = false
  https_only         = false
  use_http3          = false
  compress = {
    brotli = false
    gzip   = true
  }
  robots = {
    type = "deny"
  }
  auth = {
    md5 = {
      secret   = "string"
      forever  = false
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
    max_age     = 120
    disable     = false
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
            action  = "deny"
            country = "RU"
            region  = "BEL"
          }
        ]
        times = [
          {
            start = "2024-01-01T00:00:00Z"
            end   = "2024-01-02T00:00:00Z"
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
            end   = "2024-01-02T00:00:00Z"
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
            end   = "2024-01-02T00:00:00Z"
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
            end   = "2024-01-02T00:00:00Z"
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
        disable       = false
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
          secret   = "string"
          forever  = false
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
        max_age     = 120
        disable     = false
      }
      limitations = {
        geo = [
          {
            default_action = "allow"
            exclude = [
              {
                action  = "deny"
                country = "RU"
                region  = "BEL"
              }
            ]
            times = [
              {
                start = "2024-01-01T00:00:00Z"
                end   = "2024-01-02T00:00:00Z"
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
                end   = "2024-01-02T00:00:00Z"
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
                end   = "2024-01-02T00:00:00Z"
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
                end   = "2024-01-02T00:00:00Z"
              }
            ]
          }
        ]
      }
      compress = {
        brotli = false
        gzip   = true
      }
      ioss = false
      packaging = {
        mp4 = {
          output_protocols = [
            "MPEG-DASH"
          ]
        }
      }
      # rewrite = [
      #   {
      #     from = "^/cdn/.+(/_video_.+)"
      #     to   = "$1"
      #     flag = "break"
      #   }
      # ]
      return_http_status_code = 403
    }
  }
}

output "edu_resource" {
  value = cdnvideo_http.edu
}