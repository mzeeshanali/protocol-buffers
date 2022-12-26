package main

import (
	"fmt"
	pb "github.com/mzeeshanali/go-proto-buff/proto"
	"google.golang.org/protobuf/proto"
	"reflect"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:        321,
		Name:      "john doe",
		Fruits:    []string{"apple", "mangoes"},
		Purchased: true,
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{Id: 2, Name: "abc"},
		MultipleDummies: []*pb.Dummy{
			{Id: 2, Name: "abc"},
			{Id: 3, Name: "xyz"},
			{Id: 4, Name: "lmn"},
			{Id: 5, Name: "zxy"},
		},
	}
}
func doEnumeration() *pb.Enum {
	return &pb.Enum{
		EyeColor: pb.EyeColor_EYE_COLOR_BROWN,
	}
}

func doOneOfs(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Errorf("unexpected type %v: ", x)
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myId1": {Id: 1},
			"myId2": {Id: 12},
			"myId3": {Id: 31},
		},
	}
}

func doFile(p proto.Message) {
	path := "simple.bin"
	writeFile(path, p)

	message := &pb.Simple{}
	readFile(path, message)
	fmt.Println(message)
}

func doToJson(message proto.Message) string {
	jsonString := ToJson(message)
	return jsonString
}

func doFromJson(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	FromJson(jsonString, message)
	return message
}

func main() {
	//fmt.Println(doSimple())
	//fmt.Println(doComplex())
	//fmt.Println(doEnumeration())
	//fmt.Println("this should be id: ")
	//doOneOfs(&pb.Result_Id{Id: 1})
	//fmt.Println("this should be message: ")
	//doOneOfs(&pb.Result_Message{Message: "hello"})
	//fmt.Println(doMap())
	//doFile(doSimple())

	jsonString := doToJson(doComplex())
	message := doFromJson(jsonString, reflect.TypeOf(pb.Complex{}))

	fmt.Println(jsonString)
	fmt.Println(message)

}
