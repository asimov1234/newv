package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/asimov/newv/app/dispatcher"
	_ "github.com/asimov/newv/app/proxyman/inbound"
	_ "github.com/asimov/newv/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/asimov/newv/app/commander"
	_ "github.com/asimov/newv/app/log/command"
	_ "github.com/asimov/newv/app/proxyman/command"
	_ "github.com/asimov/newv/app/stats/command"

	// Developer preview services

	// Other optional features.
	_ "github.com/asimov/newv/app/dns"
	_ "github.com/asimov/newv/app/dns/fakedns"
	_ "github.com/asimov/newv/app/log"

	_ "github.com/asimov/newv/app/policy"
	_ "github.com/asimov/newv/app/reverse"
	_ "github.com/asimov/newv/app/router"
	_ "github.com/asimov/newv/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/asimov/newv/transport/internet/tagged/taggedimpl"

	// Developer preview features

	// Inbound and outbound proxies.
	_ "github.com/asimov/newv/proxy/blackhole"
	_ "github.com/asimov/newv/proxy/dns"
	_ "github.com/asimov/newv/proxy/dokodemo"
	_ "github.com/asimov/newv/proxy/freedom"
	_ "github.com/asimov/newv/proxy/http"
	_ "github.com/asimov/newv/proxy/loopback"
	_ "github.com/asimov/newv/proxy/shadowsocks"
	_ "github.com/asimov/newv/proxy/socks"
	_ "github.com/asimov/newv/proxy/trojan"
	_ "github.com/asimov/newv/proxy/vless/inbound"
	_ "github.com/asimov/newv/proxy/vless/outbound"
	_ "github.com/asimov/newv/proxy/vmess/inbound"
	_ "github.com/asimov/newv/proxy/vmess/outbound"
	_ "github.com/asimov/newv/proxy/wireguard"

	// Transports
	_ "github.com/asimov/newv/transport/internet/grpc"
	_ "github.com/asimov/newv/transport/internet/httpupgrade"

	_ "github.com/asimov/newv/transport/internet/reality"

	_ "github.com/asimov/newv/transport/internet/tcp"
	_ "github.com/asimov/newv/transport/internet/tls"
	_ "github.com/asimov/newv/transport/internet/udp"
	_ "github.com/asimov/newv/transport/internet/websocket"

	// Transport headers
	_ "github.com/asimov/newv/transport/internet/headers/http"
	_ "github.com/asimov/newv/transport/internet/headers/noop"
	_ "github.com/asimov/newv/transport/internet/headers/srtp"
	_ "github.com/asimov/newv/transport/internet/headers/tls"
	_ "github.com/asimov/newv/transport/internet/headers/utp"
	_ "github.com/asimov/newv/transport/internet/headers/wechat"
	_ "github.com/asimov/newv/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/asimov/newv/main/json"
	_ "github.com/asimov/newv/main/toml"
	_ "github.com/asimov/newv/main/yaml"

	// Load config from file or http(s)
	_ "github.com/asimov/newv/main/confloader/external"

	// Commands
	_ "github.com/asimov/newv/main/commands/all"
)
