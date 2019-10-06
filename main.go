package main

import (
	complexpb "complex"
	enumpb "enum_example"
	"fmt"
	"io/ioutil"
	"log"
	simplepb "simple"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {

	sm := doSimple()
	fmt.Println(sm)

	writeToFile("simple.bin", sm)
	fmt.Println("Data has been written.")

	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)

	fmt.Println("reading")
	fmt.Println(sm2)

	fmt.Println("convert to json")
	json := toJSON(sm2)
	fmt.Println(json)

	fmt.Println("convert from json to proto.Message")
	sm3 := &simplepb.SimpleMessage{}
	fromJSON(json, sm3)
	fmt.Println(sm3)

	fmt.Println("----- Enum ----")
	doEnum()

	fmt.Println("----- Complex ----")
	doComplex()
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   4,
			Name: "yes",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "no",
			},
		},
	}
	fmt.Println(cm)
}

func doEnum() {
	em := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_SUNDAY,
	}
	fmt.Println(em)
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to json", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) error {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Can't convert json to proto.Message", err)
		return err
	}
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Can't read file", err)
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Can't unmarshal", err2)
		return err2
	}
	return nil
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}
	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         1,
		IsSimple:   true,
		Name:       "My simp",
		SampleList: []int32{1, 3, 4},
	}

	sm.Name = "I renamed you"

	return &sm
}
