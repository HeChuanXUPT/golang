package main

import (
	"fmt"
	"unsafe"

	"github.com/mdlayher/netlink"
	"github.com/mdlayher/netlink/nlenc"
	"golang.org/x/sys/unix"
)

const TCP_MAX_STATES = 12
const TCPDIAG_GETSOCK = 18

type InetDiagSockId struct {
	SourcePort    uint16
	DestPort      uint16
	SourceAddress [4]uint32
	DestAddress   [4]uint32
	If            uint32
	Cookie        [2]uint32
}

type InetDiagReq struct {
	Family    uint8
	SourceLen uint8
	DestLen   uint8
	Ext       uint8
	Id        InetDiagSockId
	States    uint32
	Dbs       uint32
}

type InetDiagMsgData struct {
	Family  uint8
	State   uint8
	Timer   uint8
	Retrans uint8
	Id      InetDiagSockId
	Expires uint32
	Rqueue  uint32
	Wqueue  uint32
	Uid     uint32
	Inode   uint32
}

// MarshalBinary marshals a Message into a byte slice.
func (m InetDiagReq) MarshalBinary() []byte {
	ml := int(unsafe.Sizeof(m))
	b := make([]byte, ml)
	nlenc.PutUint8(b[0:1], m.Family)
	nlenc.PutUint32(b[52:56], m.States)
	return b
}

func getSockets() {
	c, err := netlink.Dial(unix.NETLINK_INET_DIAG, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	r := InetDiagReq{
		Family: unix.AF_INET,
		States: 1<<TCP_MAX_STATES - 1, // all status
	}

	req := netlink.Message{
		Header: netlink.Header{
			Flags: netlink.HeaderFlagsRoot | netlink.HeaderFlagsMatch |
				netlink.HeaderFlagsRequest | netlink.HeaderFlagsAcknowledge,
			Type: TCPDIAG_GETSOCK,
		},
		Data: r.MarshalBinary(),
	}

	// Perform a request, receive replies, and validate the replies
	msgs, err := c.Execute(req)
	if err != nil {
		panic(err)
	}

	fmt.Println("get msg count:", len(msgs))

	for _, m := range msgs {
		// m.Data type is InetDiagMsgData
		var data *InetDiagMsgData = *(**InetDiagMsgData)(unsafe.Pointer(&m.Data))
		fmt.Println(data)
	}
}

func main() {
	getSockets()
}
