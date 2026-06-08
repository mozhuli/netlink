//go:build go1.26

package nl

import (
	"syscall"

	"golang.org/x/sys/unix"
)

func newReceiveBuffer() []byte {
	return nil
}

func (s *NetlinkSocket) Receive() ([]syscall.NetlinkMessage, *unix.SockaddrNetlink, error) {
	var rb [RECEIVE_BUFFER_SIZE]byte

	nr, from, err := s.receiveRaw(rb[:])
	if err != nil {
		return nil, nil, err
	}

	return parseNetlinkReceiveResult(rb[:], nr, from)
}
