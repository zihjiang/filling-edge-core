// +build 386 windows,amd64 windows



package common

import (
	"bytes"
	"encoding/binary"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/windows"
	"syscall"
	"time"
)

type SIDType uint32

//https://docs.microsoft.com/en-us/windows/desktop/CIMWin32Prov/win32-useraccount
var SIDTypeStringMap = map[SIDType]string{
	SIDType(1): "User",
	SIDType(2): "Group",
	SIDType(3): "Alias",
	SIDType(4): "Well Known Group",
	SIDType(5): "Alias",
	SIDType(6): "Deleted Account",
	SIDType(7): "Unknown",
	SIDType(8): "Computer",
}

func (s SIDType) GetSidTypeString() string {
	if mapping, stringMappingPresent := SIDTypeStringMap[s]; stringMappingPresent {
		return mapping
	}
	return ""
}

type SIDInfo struct {
	Name    string
	Domain  string
	SIDType SIDType
}

func GetSidInfo(sid *windows.SID) (*SIDInfo, error) {
	var sidInfo *SIDInfo
	account, domain, sidType, err := sid.LookupAccount("")
	if err == nil {
		sidInfo = &SIDInfo{Name: account, Domain: domain, SIDType: SIDType(sidType)}
	}
	return sidInfo, err
}

func ExtractString(byteData []byte) (string, error) {
	wordArray := make([]uint16, len(byteData)/2)
	err := binary.Read(bytes.NewReader(byteData), binary.LittleEndian, wordArray)
	if err != nil {
		log.WithError(err).Error("Error reading binary data")
		return "", err
	}
	return syscall.UTF16ToString(wordArray), nil
}

func ExtractStrings(byteData []byte, stringCount uint16) []string {
	strs := make([]string, 0, stringCount)
	wordArray := make([]uint16, len(byteData)/2)
	err := binary.Read(bytes.NewReader(byteData), binary.LittleEndian, wordArray)
	if err != nil {
		log.WithError(err).Fatal()
	}
	pos := 0
	for idx, value := range wordArray {
		if value == 0 {
			str := syscall.UTF16ToString(wordArray[pos:idx])
			strs = append(strs, str)
			pos = idx + 1
			stringCount--
			if stringCount == 0 {
				break
			}
		}
	}
	return strs
}

func ConvertTimeToLong(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond) / int64(time.Nanosecond)
}
