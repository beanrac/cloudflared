package mydns

import (
	"context"
	"net"
	"time"
)

// CustomResolver uses Cloudflare's DNS server (1.1.1.1)
var CustomResolver = &net.Resolver{
	PreferGo: true,
	Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
		d := net.Dialer{Timeout: time.Second * 5}
		return d.DialContext(ctx, "udp", "1.1.1.1:53")
	},
}

// LookupHost resolves a host using the custom DNS server
func LookupHost(name string) ([]string, error) {
	return CustomResolver.LookupHost(context.Background(), name)
}

// LookupTXT returns TXT records using the custom DNS server
func LookupTXT(name string) ([]string, error) {
	return CustomResolver.LookupTXT(context.Background(), name)
}

// LookupIP returns IP addresses using the custom DNS server
func LookupIP(name string) ([]net.IP, error) {
	return CustomResolver.LookupIP(context.Background(), "ip", name)
}

// LookupCNAME resolves the canonical name for the given host
func LookupCNAME(name string) (string, error) {
	return CustomResolver.LookupCNAME(context.Background(), name)
}

// LookupNS returns name servers for a domain using custom DNS
func LookupNS(name string) ([]*net.NS, error) {
	return CustomResolver.LookupNS(context.Background(), name)
}

// LookupSRV resolves SRV records for a service, protocol, and domain
func LookupSRV(service, proto, name string) (string, []*net.SRV, error) {
	return CustomResolver.LookupSRV(context.Background(), service, proto, name)
}
