package ntp

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type NTPClient struct {
	Server  string
	Timeout time.Duration
}

const NtpEpochOffset = 2208988800

func NewClient(server string) (client *NTPClient) {
	client = &NTPClient{
		Server:  server,
		Timeout: 15 * time.Second,
	}
	return
}

func (client *NTPClient) GetTime() (t time.Time, err error) {
	conn, err := net.Dial("udp", client.Server)
	if err != nil {
		return t, fmt.Errorf("failed to connect: %v", err)
	}
	defer conn.Close()
	if err := conn.SetDeadline(time.Now().Add(client.Timeout)); err != nil {
		return t, fmt.Errorf("failed to set deadline: %v", err)
	}

	req := &NTPPacket{Header: 0x1B}

	if err := binary.Write(conn, binary.BigEndian, req); err != nil {
		return t, fmt.Errorf("failed to send request: %v", err)
	}

	rsp := &NTPPacket{}
	if err := binary.Read(conn, binary.BigEndian, rsp); err != nil {
		return t, fmt.Errorf("failed to read server response: %v", err)
	}
	secs := float64(rsp.TxTimeSec) - NtpEpochOffset
	nanos := (int64(rsp.TxTimeFrac) * 1e9) >> 32
	t = time.Unix(int64(secs), nanos)
	return
}
