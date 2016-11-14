package ImPdu

var PudHeadLen uint16 = 16

type PduHeader_t struct {
	length     uint32
	version    uint16
	flag       uint16
	service_id uint16
	command_id uint16
	seq_num    uint16
	reversed   uint16
}

type ImPdu struct {
	header PduHeader_t
	pudBuf []byte
}

func NewImPdu() *ImPdu {
	return &ImPdu{}

}

func (pdu *ImPdu) SetServiceId(sid uint16) {
	pdu.header.service_id = sid
}
func (pdu *ImPdu) SetCommandId(cid uint16) {
	pdu.header.command_id = cid
}
func (pdu *ImPdu) SetSeqNum(seq_num uint16) {
	pdu.header.seq_num = seq_num
}
