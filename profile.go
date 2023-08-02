package req

import (
	"github.com/imroc/req/v3/pkg/http2"
	utls "github.com/refraction-networking/utls"
)

type ClientProfile func(c *Client)

var http2SettingsChrome = []http2.Setting{
	{
		ID:  http2.SettingHeaderTableSize,
		Val: 65536,
	},
	{
		ID:  http2.SettingEnablePush,
		Val: 0,
	},
	{
		ID:  http2.SettingMaxConcurrentStreams,
		Val: 1000,
	},
	{
		ID:  http2.SettingInitialWindowSize,
		Val: 6291456,
	},
	{
		ID:  http2.SettingMaxHeaderListSize,
		Val: 262144,
	},
}

var chromeHeaderOrder = []string{
	"host",
	"pragma",
	"cache-control",
	"sec-ch-ua",
	"sec-ch-ua-mobile",
	"sec-ch-ua-platform",
	"upgrade-insecure-requests",
	"user-agent",
	"accept",
	"sec-fetch-site",
	"sec-fetch-mode",
	"sec-fetch-user",
	"sec-fetch-dest",
	"accept-encoding",
	"accept-language",
	"cookie",
}

var chromePseudoHeaderOrder = []string{
	":method",
	":authority",
	":scheme",
	":path",
}

var chromeHeaders = map[string]string{
	"pragma":                    "no-cache",
	"cache-control":             "no-cache",
	"sec-ch-ua":                 `"Chromium";v="106", "Google Chrome";v="106", "Not;A=Brand";v="99"`,
	"sec-ch-ua-mobile":          "?0",
	"sec-ch-ua-platform":        `"macOS"`,
	"upgrade-insecure-requests": "1",
	"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36",
	"accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9",
	"sec-fetch-site":            "none",
	"sec-fetch-mode":            "navigate",
	"sec-fetch-user":            "?1",
	"sec-fetch-dest":            "document",
	"accept-language":           "zh-CN,zh;q=0.9,en;q=0.8,zh-TW;q=0.7,it;q=0.6",
}

var ClientProfile_Chrome ClientProfile = func(c *Client) {
	c.SetTLSFingerprint(utls.HelloChrome_106_Shuffle).
		SetUserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36").
		SetCommonPseudoHeaderOder(chromePseudoHeaderOrder...).
		SetCommonHeaders(chromeHeaders).
		SetCommonHeaderOrder(chromeHeaderOrder...).
		SetHTTP2SettingsFrame(http2SettingsChrome...)
}
