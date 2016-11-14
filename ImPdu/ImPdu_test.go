package ImPdu

import (
	"testing"
)

func TestCreateImPdu(t *testing.T) {

	pdu := NewImPdu()
	pdu.SetCommandId(11)
	pdu.SetSeqNum(1)
	pdu.SetServiceId(1)
	t.Logf("%v", pdu)
}
