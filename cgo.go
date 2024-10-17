package util

import "unsafe"

func Bytes2Pointer(data []byte) unsafe.Pointer {
	if len(data) == 0 {
		return nil
	}
	return unsafe.Pointer(&data[0])
}

func Bool2Byte(b bool) byte {
	if b {
		return 1
	}
	return 0
}

func StrLen(buf []byte) int {
	strCnt := 0
	for _, v := range buf {
		if v == 0 {
			break
		}
		strCnt++
	}
	return strCnt
}
