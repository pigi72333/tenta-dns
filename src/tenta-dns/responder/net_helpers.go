package responder

import (
	"tenta-dns/runtime"
	"fmt"
)

func hostInfo(v4 bool, net string, d *runtime.ServerDomain) (ip string, port int) {
	if v4 {
		ip = d.IPv4
	} else {
		ip = fmt.Sprintf("[%s]", d.IPv6)
	}
	if net == "tcp" {
		if d.DnsTcpPort <= runtime.PORT_UNSET {
			panic("Unable to start a TCP recursive DNS server without a valid TCP port")
		}
		port = d.DnsTcpPort
	} else if net == "udp" {
		if d.DnsUdpPort <= runtime.PORT_UNSET {
			panic("Unable to start a UDP recursive DNS server without a valid UDP port")
		}
		port = d.DnsUdpPort
	} else if net == "tls" {
		if d.DnsTlsPort <= runtime.PORT_UNSET {
			panic("Unable to start a TLS recursive DNS server without a valid TLS port")
		}
		port = d.DnsTlsPort
	} else {
		panic(fmt.Sprintf("Unknown network type %s", net))
	}
	return
}
