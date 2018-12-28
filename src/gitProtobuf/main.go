package main

import (
	"bytes"
	"compress/gzip"
	"encoding/hex"
	"fmt"
	"io/ioutil"

	"github.com/gogo/protobuf/proto"
)

func main() {
	var iType IpcType = IpcType_positionVehicle
	var type1 *IpcType = &iType
	sCompanyId := "1"
	pos := OTIpc{
		CompanyId: &sCompanyId,
		Source:    &sCompanyId,
		IPCType:   type1,
	}
	pos.PositionVehicle = make([]*PositionVehicle, 0)
	sVehicleNo := "123465"
	var iVehicleRegionCode uint32 = 0
	var iPositionTime uint64 = 1
	var iLongitude uint64 = 2
	var iLatitude uint64 = 3
	var iSpeed uint32 = 4
	var iDirection uint32 = 5
	var iElevation int32 = 6
	var iMileage float32 = 7
	var iWarnStatus uint32 = 8
	var iVehStatus uint32 = 9
	sOrderId := ""
	posVeh := &PositionVehicle{
		CompanyId:         &sCompanyId,
		VehicleNo:         &sVehicleNo,
		VehicleRegionCode: &iVehicleRegionCode,
		PositionTime:      &iPositionTime,
		Longitude:         &iLongitude,
		Latitude:          &iLatitude,
		Speed:             &iSpeed,
		Direction:         &iDirection,
		Elevation:         &iElevation,
		Mileage:           &iMileage,
		WarnStatus:        &iWarnStatus,
		VehStatus:         &iVehStatus,
		OrderId:           &sOrderId,
	}
	pos.PositionVehicle = append(pos.PositionVehicle, posVeh)

	data, err := proto.Marshal(&pos)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(hex.EncodeToString(data))

	inData, _ := hex.DecodeString("1F8B0800000000000000E34AE2E2313432307071718E3036363017623490686A609CD4C0E18926C1F97253A37788A1A5B181C48E4B620A5BDFADBFCFAA317FE24C2B831B573F085A083984784846302630643054F1199B1A5A1A9A99999A18195B181B0100DAE398CF64000000")
	bufInData := bytes.NewBuffer(inData)
	gzipData, _ := gzip.NewReader(bufInData)
	rBuf, _ := ioutil.ReadAll(gzipData)
	fmt.Println(hex.EncodeToString(rBuf))
	proto.Unmarshal(rBuf, &pos)
	fmt.Println(pos)
}

func TestPosition() {
	// user := &UserInfo{
	// 	Message: "tiptok",
	// 	Length:  10,
	// 	Cnt:     20,
	// }
	// data, err := proto.Marshal(user)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(data)
	// tmpUser := &UserInfo{}
	// _ = proto.Unmarshal(data, tmpUser)
	// fmt.Println(tmpUser)

	dataStr := "0A0B313838363031383330353610B4CBE4DE05190000000000000000210000000000000000290000000000000000310000000000005940390000000000407F40403C4837500058B4CA09600068FFFFFFFFFFFFFFFFFF0170007A008001B4CBE4DE05"
	data, err := hex.DecodeString(dataStr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	pos := &TermPosition{}
	err = proto.Unmarshal(data, pos)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%s %v", "Proto 解析:", pos)
	}
}
