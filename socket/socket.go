package socket

import (
	"fmt"
	"net"
	"strconv"
)

// See https://github.com/mwielgoszewski/logstash-filter-goaudit/blob/e6064b9852354fa862de4d23570ce0ab36a1595a/lib/logstash/filters/goaudit.rb#L283=

var AddressFamilies []string

func init() {
	AddressFamilies = []string{
		"unspecified",
		"local",
		"inet",
		"ax25",
		"ipx",
		"appletalk",
		"netrom",
		"bridge",
		"atmpvc",
		"x25",
		"inet6",
		"rose",
		"decnet",
		"netbeui",
		"security",
		"key",
		"netlink",
		"packet",
		"ash",
		"econet",
		"atmsvc",
		"rds",
		"sna",
		"irda",
		"pppox",
		"wanpipe",
		"llc",
		"ib",
		"mpls",
		"can",
		"tipc",
		"bluetooth",
		"iucv",
		"rxrpc",
		"isdn",
		"phonet",
		"ieee802154",
		"caif",
		"alg",
		"nfc",
		"vsock",
		"kcm",
		"qipcrtr",
	}
}

type Saddr struct {
	Family string
	IP     net.IP
	Port   int32
}

func Parse4(txt string) (*Saddr, error) {
	i, _ := strconv.ParseInt(txt[:2], 16, 64)
	ii, _ := strconv.ParseInt(txt[2:4], 16, 64)
	i += 256 * ii
	family := AddressFamilies[i]
	fmt.Println(family)
	switch i {
	case 1:
		return parseLocal(txt)
	case 2:
		return parseInet(txt)
	case 10:
		return parseInet6(txt)
	default:
		return nil, fmt.Errorf("family not handled : %s", family)
	}
}

func parseLocal(txt string) (*Saddr, error) {
	return nil, nil
}

func parseInet(txt string) (*Saddr, error) {
	i, _ := strconv.ParseInt(txt[4:6], 16, 64)
	ii, _ := strconv.ParseInt(txt[6:8], 16, 64)
	port := i*256 + ii
	ip := make(net.IP, 4)
	for i := 0; i < 4; i++ {
		ii, _ := strconv.ParseInt(txt[8+2*i:10+2*i], 16, 32)
		ip[i] = byte(ii)
	}
	return &Saddr{
		Family: AddressFamilies[2],
		IP:     ip,
		Port:   int32(port),
	}, nil
}

func parseInet6(txt string) (*Saddr, error) {
	return nil, nil
}
