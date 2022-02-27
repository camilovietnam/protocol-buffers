package main

import (
	"fmt"
	"go_gen/src/go_gen/complexpb"
	"go_gen/src/go_gen/enumpb"
	"go_gen/src/go_gen/simplepb"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	sm := doSimple()
	writeToFile("simple.bin", sm)
	readFromFile("simple.bin", sm)
	jsonDemo(sm)
	doEnum()
	doComplex()
}

func doComplex() {
	fmt.Println("\n>>> Do Complex: ")
	cm := &complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   666,
			Name: "My message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{Id: 1, Name: "Batman Begins"},
			{Id: 2, Name: "The Dark Knight"},
			{Id: 3, Name: "The Dark Knight Rises"},
		},
	}

	fmt.Println(cm)
}

func doEnum() {
	fmt.Println("\n>>> Do Enum: ")
	em := enumpb.EnumMessage{
		Id:           69,
		DayOfTheWeek: enumpb.DayOfTheWeek_FRIDAY,
	}

	fmt.Println("The day of the week is : ", em.GetDayOfTheWeek())
}

func jsonDemo(sm proto.Message) {
	sm2 := &simplepb.SimpleMessage{}

	asJSON := writeToJSON(sm)
	fmt.Println("Message as JSON: ", asJSON)

	readFromJSON(asJSON, sm2)
	fmt.Println("The JSON, converted to message: ", sm2)
}

func writeToJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}

	out, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Error writing to JSON", err)
	}

	return out
}

func readFromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Error reading from JSON string", err)
	}
}

func readFromFile(fn string, pb proto.Message) {
	in, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalln("Could not read from file", err)
		return
	}

	if err = proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Could not de-serialise", err)
		return
	}
}

func writeToFile(fn string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise", err)
		return
	}

	err = ioutil.WriteFile(fn, out, 0644)
	if err != nil {
		log.Fatalln("Can't write to file", err)
		return
	}
}

func doSimple() *simplepb.SimpleMessage {
	fmt.Println("\n>>> Do Simple:")
	sm := simplepb.SimpleMessage{
		Id:         123,
		IsSimple:   false,
		Name:       "My simple message",
		SampleList: []int32{3, 6, 9, 12},
	}

	echo := fmt.Sprintf(`
	The message Name: %s
	The message ID: %d
	Is it a simple message? %t
	`, sm.GetName(), sm.GetId(), sm.GetIsSimple())
	fmt.Println(echo)

	return &sm
}
