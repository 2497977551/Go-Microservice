package main

import (
	"fmt"
	"io/ioutil"
	"log"

	Protobuf "github.com/MicroService-Su/ProtobufDemo"
	"google.golang.org/protobuf/proto"
)

func main() {
	p := Protobuf.Persons{
		Name:        "Joshua",
		PhoneNumber: []int64{123456789, 465456564},
		Addr:        "广东",
	}
	
	if err := writeToFile("test.bin",&p); err != nil{
		fmt.Println("写入文件失败")
	}
	fmt.Println("写入文件成功")
	pp := Protobuf.Persons{}
	if err := readFromFile("test.bin", &pp); err != nil {
		fmt.Println("读取文件失败")
	}
	fmt.Println("读取文件成功")
	fmt.Println(pp)
}

func readFromFile(fileName string, pb proto.Message) error {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("读取文件失败",err.Error())
	
	}
	if err := proto.Unmarshal(bytes, pb); err != nil {
		log.Fatalln("转化失败",err.Error())
	
	}
	return nil
}

func writeToFile(fileName string, pb proto.Message) error {
	bytes, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("proto序列化失败", err.Error())
		return err
	}
	if err := ioutil.WriteFile(fileName, bytes, 0644); err != nil {
		log.Fatalln("写入文件失败", err.Error())
		return err
	}
	return nil
}
