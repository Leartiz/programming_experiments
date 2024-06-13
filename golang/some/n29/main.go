package main

import (
	"fmt"

	tls "github.com/raff/tls-ext"
	psk "github.com/raff/tls-psk"
)

func getIdentity() string {
	return "key14936378984900"
}

func getKey(id string) ([]byte, error) {
	return []byte("zota-device"), nil
}

type FrameHeader struct {
}

func main() {
	var (
		config = &tls.Config{
			CipherSuites: []uint16{
				psk.TLS_PSK_WITH_AES_256_CBC_SHA,
			},
			Extra: psk.PSKConfig{
				GetKey:      getKey,
				GetIdentity: getIdentity,
			},
		}
	)

	// connect a client
	conn, err := tls.Dial("tcp", "192.168.154.252:9090", config)
	if err != nil {
		fmt.Printf("dial %v", err)
		return
	}

	if err = conn.Handshake(); err != nil {
		fmt.Printf("handshake %v", err)
		return
	}

	conn.Write([]byte("123"))
}
