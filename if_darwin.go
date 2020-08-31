package water

import "os"

func WrapTunFD(fd int, ifname string) *Interface {
	return &Interface{
		isTAP: false,
		name:  ifname,
		ReadWriteCloser: &tunReadCloser{
			f: os.NewFile(uintptr(fd), ifname),
		},
	}
}
