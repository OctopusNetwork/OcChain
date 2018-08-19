
package stun

import (
	"testing"
)

func TestNewPacketFromBytes(t *testing.T) {
	b := make([]byte, 23)
	_, err := newPacketFromBytes(b)
	if err == nil {
		t.Errorf("newPacketFromBytes error")
	}
	b = make([]byte, 24)
	_, err = newPacketFromBytes(b)
	if err != nil {
		t.Errorf("newPacketFromBytes error")
	}
}

func TestNewPacket(t *testing.T) {
	_, err := newPacket()
	if err != nil {
		t.Errorf("newPacket error")
	}
}

func TestPacketAll(t *testing.T) {
	p, err := newPacket()
	if err != nil {
		t.Errorf("newPacket error")
	}
	p.addAttribute(*newChangeReqAttribute(true, true))
	p.addAttribute(*newSoftwareAttribute("aaa"))
	p.addAttribute(*newFingerprintAttribute(p))
	pkt, err := newPacketFromBytes(p.bytes())
	if err != nil {
		t.Errorf("newPacketFromBytes error")
	}
	if pkt.types != 0 {
		t.Errorf("newPacketFromBytes error")
	}
	if pkt.length < 24 {
		t.Errorf("newPacketFromBytes error")
	}
}
