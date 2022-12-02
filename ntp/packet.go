package ntp

import "encoding/binary"

type NTPPacket struct {
	Header         uint8
	Stratum        uint8
	Poll           int8
	Precision      int8
	RootDelay      uint32
	RootDispersion uint32
	ReferenceID    uint32
	RefTimeSec     uint32
	RefTimeFrac    uint32
	OrigTimeSec    uint32
	OrigTimeFrac   uint32
	RxTimeSec      uint32
	RxTimeFrac     uint32
	TxTimeSec      uint32
	TxTimeFrac     uint32
}

func NewPacket() (packet *NTPPacket) {
	packet = &NTPPacket{}
	return packet
}

func (packet *NTPPacket) Bytes() (data []byte, err error) {
	data = make([]byte, 48)
	h := uint8(0)
	// h |= (uint8(packet.LeapIndicator) << 6) & 0xC0
	// h |= (uint8(packet.Version) << 3) & 0x38
	// h |= (uint8(packet.Mode)) & 0x07
	data[0] = byte(h)
	data[1] = byte(packet.Stratum)
	data[2] = byte(packet.Poll)
	data[3] = byte(packet.Precision)

	// The remaining fields can just be copied in big endian order.
	binary.BigEndian.PutUint32(data[4:8], uint32(packet.RootDelay))
	binary.BigEndian.PutUint32(data[8:12], uint32(packet.RootDispersion))
	binary.BigEndian.PutUint32(data[12:16], uint32(packet.ReferenceID))
	// binary.BigEndian.PutUint64(data[16:24], uint64(packet.ReferenceTimestamp))
	// binary.BigEndian.PutUint64(data[24:32], uint64(packet.OriginTimestamp))
	// binary.BigEndian.PutUint64(data[32:40], uint64(packet.ReceiveTimestamp))
	// binary.BigEndian.PutUint64(data[40:48], uint64(packet.TransmitTimestamp))
	return
}
