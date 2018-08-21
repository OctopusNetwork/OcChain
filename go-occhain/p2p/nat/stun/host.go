package stun

import (
	"net"
	"strconv"
)

// Host defines the network address including address family, IP address and port.
type Host struct {
	family uint16
	ip     string
	port   uint16
}

func newHostFromStr(s string) *Host {
	udpAddr, err := net.ResolveUDPAddr("udp", s)
	if err != nil {
		return nil
	}
	host := new(Host)
	if udpAddr.IP.To4() != nil {
		host.family = attributeFamilyIPv4
	} else {
		host.family = attributeFamilyIPV6
	}
	host.ip = udpAddr.IP.String()
	host.port = uint16(udpAddr.Port)
	return host
}

// Family returns the family type of a host (IPv4 or IPv6).
func (h *Host) Family() uint16 {
	return h.family
}

// IP returns the internet protocol address of the host.
func (h *Host) IP() string {
	return h.ip
}

// Port returns the port number of the host.
func (h *Host) Port() uint16 {
	return h.port
}

// TransportAddr returns the transport layer address of the host.
func (h *Host) TransportAddr() string {
	return net.JoinHostPort(h.ip, strconv.Itoa(int(h.port)))
}

// String returns the string representation of the host address.
func (h *Host) String() string {
	return h.TransportAddr()
}
