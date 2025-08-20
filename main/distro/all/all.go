package all

import (
	// The following are necessary as they register handlers in their init functions.

	// Mandatory features. Can't remove unless there are replacements.
	_ "github.com/asimov1234/newv/app/dispatcher"
	_ "github.com/asimov1234/newv/app/proxyman/inbound"
	_ "github.com/asimov1234/newv/app/proxyman/outbound"

	// Default commander and all its services. This is an optional feature.
	_ "github.com/asimov1234/newv/app/commander"
	_ "github.com/asimov1234/newv/app/log/command"
	_ "github.com/asimov1234/newv/app/proxyman/command"
	_ "github.com/asimov1234/newv/app/stats/command"

	// Developer preview services

	// Other optional features.
	_ "github.com/asimov1234/newv/app/dns"
	_ "github.com/asimov1234/newv/app/dns/fakedns"
	_ "github.com/asimov1234/newv/app/log"

	_ "github.com/asimov1234/newv/app/policy"
	_ "github.com/asimov1234/newv/app/reverse"
	_ "github.com/asimov1234/newv/app/router"
	_ "github.com/asimov1234/newv/app/stats"

	// Fix dependency cycle caused by core import in internet package
	_ "github.com/asimov1234/newv/transport/internet/tagged/taggedimpl"

	// Developer preview features

	// Inbound and outbound proxies.
	_ "github.com/asimov1234/newv/proxy/blackhole"
	_ "github.com/asimov1234/newv/proxy/dns"
	_ "github.com/asimov1234/newv/proxy/dokodemo"
	_ "github.com/asimov1234/newv/proxy/freedom"
	_ "github.com/asimov1234/newv/proxy/http"
	_ "github.com/asimov1234/newv/proxy/loopback"
	_ "github.com/asimov1234/newv/proxy/shadowsocks"
	_ "github.com/asimov1234/newv/proxy/socks"
	_ "github.com/asimov1234/newv/proxy/trojan"
	_ "github.com/asimov1234/newv/proxy/vless/inbound"
	_ "github.com/asimov1234/newv/proxy/vless/outbound"
	_ "github.com/asimov1234/newv/proxy/vmess/inbound"
	_ "github.com/asimov1234/newv/proxy/vmess/outbound"
	_ "github.com/asimov1234/newv/proxy/wireguard"

	// Transports
	_ "github.com/asimov1234/newv/transport/internet/grpc"
	_ "github.com/asimov1234/newv/transport/internet/httpupgrade"

	_ "github.com/asimov1234/newv/transport/internet/reality"

	_ "github.com/asimov1234/newv/transport/internet/tcp"
	_ "github.com/asimov1234/newv/transport/internet/tls"
	_ "github.com/asimov1234/newv/transport/internet/udp"
	_ "github.com/asimov1234/newv/transport/internet/websocket"

	// Transport headers
	_ "github.com/asimov1234/newv/transport/internet/headers/http"
	_ "github.com/asimov1234/newv/transport/internet/headers/noop"
	_ "github.com/asimov1234/newv/transport/internet/headers/srtp"
	_ "github.com/asimov1234/newv/transport/internet/headers/tls"
	_ "github.com/asimov1234/newv/transport/internet/headers/utp"
	_ "github.com/asimov1234/newv/transport/internet/headers/wechat"
	_ "github.com/asimov1234/newv/transport/internet/headers/wireguard"

	// JSON & TOML & YAML
	_ "github.com/asimov1234/newv/main/json"
	_ "github.com/asimov1234/newv/main/toml"
	_ "github.com/asimov1234/newv/main/yaml"

	// Load config from file or http(s)
	_ "github.com/asimov1234/newv/main/confloader/external"

	// Commands
	_ "github.com/asimov1234/newv/main/commands/all"
)
