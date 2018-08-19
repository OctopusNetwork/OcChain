package stun

import (
	"fmt"
	"net"
)

type response struct {
	packet      *packet // the original packet from the server
	serverAddr  *Host   // the address received packet
	changedAddr *Host   // parsed from packet
	mappedAddr  *Host   // parsed from packet, external addr of client NAT
	otherAddr   *Host   // parsed from packet, to replace changedAddr in RFC 5780
	identical   bool    // if mappedAddr is in local addr list
}

func newResponse(pkt *packet, conn net.PacketConn) *response {
	resp := &response{pkt, nil, nil, nil, nil, false}
	if pkt == nil {
		return resp
	}
	// RFC 3489 doesn't require the server return XOR mapped address.
	mappedAddr := pkt.getXorMappedAddr()
	if mappedAddr == nil {
		mappedAddr = pkt.getMappedAddr()
	}
	resp.mappedAddr = mappedAddr
	// compute identical
	localAddrStr := conn.LocalAddr().String()
	if mappedAddr != nil {
		mappedAddrStr := mappedAddr.String()
		resp.identical = isLocalAddress(localAddrStr, mappedAddrStr)
	}
	// compute changedAddr
	changedAddr := pkt.getChangedAddr()
	if changedAddr != nil {
		changedAddrHost := newHostFromStr(changedAddr.String())
		resp.changedAddr = changedAddrHost
	}
	// compute otherAddr
	otherAddr := pkt.getOtherAddr()
	if otherAddr != nil {
		otherAddrHost := newHostFromStr(otherAddr.String())
		resp.otherAddr = otherAddrHost
	}

	return resp
}

// String is only used for verbose mode output.
func (r *response) String() string {
	if r == nil {
		return "Nil"
	}
	return fmt.Sprintf("{packet nil: %v, local: %v, remote: %v, changed: %v, other: %v, identical: %v}",
		r.packet == nil,
		r.mappedAddr,
		r.serverAddr,
		r.changedAddr,
		r.otherAddr,
		r.identical)
}
