package coredns

import (
	"log"
	"strconv"
	"strings"
	"time"
)

var nowFunc = time.Now

// generateSerial takes an old SOA serial number and increments it.
func generateSerial(oldSerial uint32) uint32 {
	// Serial numbers are in the format yyyymmddvv
	// where vv is a version count that starts at 01 each day.
	// Multiple serial numbers generated on the same day increase vv.
	// If the old serial number is not in this format, it gets replaced
	// with the new format. However if that would mean a new serial number
	// that is smaller than the old one, we punt and increment the old number.
	// At no time will a serial number == 0 be returned.

	original := oldSerial
	oldSerialStr := strconv.FormatUint(uint64(oldSerial), 10)
	var newSerial uint32

	// Make draft new serial number:
	today := nowFunc().UTC()
	todayStr := today.Format("20060102")
	version := uint32(1)
	todayNum, err := strconv.ParseUint(todayStr, 10, 32)
	if err != nil {
		log.Fatalf("new serial won't fit in 32 bits: %v", err)
	}
	draft := uint32(todayNum)*100 + version

	method := "none" // Used only in debugging.
	if oldSerial > draft {
		// If old_serial was really slow, upgrade to new yyyymmddvv standard:
		method = "o>d"
		newSerial = oldSerial + 1
		newSerial = oldSerial + 1
	} else if oldSerial == draft {
		// Edge case: increment old serial:
		method = "o=d"
		newSerial = draft + 1
	} else if len(oldSerialStr) != 10 {
		// If old_serial is wrong number of digits, upgrade to yyyymmddvv standard:
		method = "len!=10"
		newSerial = draft
	} else if strings.HasPrefix(oldSerialStr, todayStr) {
		// If old_serial just needs to be incremented:
		method = "prefix"
		newSerial = oldSerial + 1
	} else {
		// First serial number to be requested today:
		method = "default"
		newSerial = draft
	}

	if newSerial == 0 {
		// We never return 0 as the serial number.
		newSerial = 1
	}
	if oldSerial == newSerial {
		log.Fatalf("%v: old_serial == new_serial (%v == %v) draft=%v method=%v", original, oldSerial, newSerial, draft, method)
	}
	if oldSerial > newSerial {
		log.Fatalf("%v: old_serial > new_serial (%v > %v) draft=%v method=%v", original, oldSerial, newSerial, draft, method)
	}
	return newSerial
}
