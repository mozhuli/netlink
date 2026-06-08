//go:build !go1.26

package nl

import (
	"syscall"

	"golang.org/x/sys/unix"
)

func newReceiveBuffer() []byte {
	return make([]byte, RECEIVE_BUFFER_SIZE)
}

func (s *NetlinkSocket) Receive() ([]syscall.NetlinkMessage, *unix.SockaddrNetlink, error) {
	s.receiveMu.Lock()
	defer s.receiveMu.Unlock()

	nr, from, err := s.receiveRaw(s.receiveBuffer)
	if err != nil {
		return nil, nil, err
	}

	return parseNetlinkReceiveResult(s.receiveBuffer, nr, from)
}
