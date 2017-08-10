package modxml

import (
	"strconv"
	"os"
	"io/ioutil"
	"encoding/xml"
	"log"

	simLog "GoSimulation/log"
)

var (
	ChardwareFile =  "./Config/Chardware.xml"
)

type indoor_root struct {
	XMLName xml.Name			`xml:"CMap_root"`
	Indoorhotspot indoorhotspot `xml:"IndoorHotspot"`
}
type indoorhotspot struct {
	XMLName xml.Name			`xml:"IndoorHotspot"`
	DXPoint string				`xml:"dXPoint,attr"`
	DYPoint string				`xml:"dYPoint,attr"`
	DLength string				`xml:"dLength,attr"`
	DWidth string				`xml:"dWidth,attr"`
	D string					`xml:"D,attr"`
	FirstOrientation string		`xml:"FirstOrientation,attr"`
	BSTypeName string			`xml:"BSTypeName,attr"`
	UERx uerx					`xml:"UERx"`
}
type sector_root struct {
	XMLName xml.Name			`xml:"CMap_root"`
	Sector sector 				`xml:"Sector"`
}
type sector struct {
	XMLName xml.Name			`xml:"Sector"`
	Num string 					`xml:"Num,attr"`
	DXPoint string				`xml:"dXPoint,attr"`
	DYPoint string				`xml:"dYPoint,attr"`
	DLength string				`xml:"dLength,attr"`
	D string					`xml:"D,attr"`
	FirstOrientation string		`xml:"FirstOrientation,attr"`
	BSTypeName string			`xml:"BSTypeName,attr"`
	UERx uerx					`xml:"UERx"`
}
type uerx struct {
	XMLName xml.Name			`xml:"UERx"`
	Num string					`xml:"Num,attr"`
	Distribution string			`xml:"Distribution,attr"`
}


func ModChardware(ueNum int, seedType string) error {
	logger := simLog.NewLog()
	content, err := ioutil.ReadFile(ChardwareFile)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	var xmlContent []byte
	switch seedType {
	case "indoor":
		var result indoor_root
		err = xml.Unmarshal(content, &result)
		if err != nil {
			log.Fatal(err)
			return err
		}
		result.Indoorhotspot.UERx.Num = strconv.Itoa(ueNum)

		// 写回Chardwre文件
		xmlContent, err = xml.MarshalIndent(result, "", "")
		if err != nil {
			logger.Fatal(err)
			return err
		}
	case "19sector":
		var result sector_root
		err = xml.Unmarshal(content, &result)
		if err != nil {
			log.Fatal(err)
			return err
		}
		result.Sector.UERx.Num = strconv.Itoa(ueNum)

		// 写回Chardwre文件
		xmlContent, err = xml.MarshalIndent(result, "", "")
		if err != nil {
			logger.Fatal(err)
			return err
		}
	}

	//加入头文件
	headerBytes := []byte(xml.Header)
	xmlOutput := append(headerBytes, xmlContent...)
	err = ioutil.WriteFile(ChardwareFile, xmlOutput, os.ModeAppend)
	if err != nil {
		logger.Fatal(err)
		return err
	}
	return nil
}

