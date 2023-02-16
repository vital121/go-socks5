package socks5

import (
	"net"

	"golang.org/x/net/context"
)

// NameResolver is used to implement custom name resolution
type NameResolver interface {
	Resolve(ctx context.Context, name string) (context.Context, net.IP, error)
}

// DNSResolver uses the system DNS to resolve host names
type DNSResolver struct{}

func (d DNSResolver) Resolve(ctx context.Context, name string) (context.Context, net.IP, error) {
	addr, err := net.ResolveIPAddr("ip", name)
	log.Printf("Resolve addr %s\n", addr)
	log.Printf("Resolve net.IP %s\n", net.IP)
	if err != nil {
		return ctx, nil, err
	}
	return ctx, addr.IP, err
}
