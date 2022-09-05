package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"
)

/*
Let it be just a model structs with json and yml tags
Adapter will convert it in both sides
*/
type OrderClient struct {
	Name     string `json:"name",yml:"name"`
	Lastname string `json:"lastname",yml:"lastname"`
	Age      uint   `json:"age",yml:"age"`
	UserUid  string `json:"user_uid",yml:"user_uid"`
}

type OrderData struct {
	OrderUid    string      `json:"order_uid",yml:"order_uid"`
	Destination string      `json:"destination",yml:"destination"`
	Weight      int         `json:"weigth",yml:"weigth"`
	Client      OrderClient `json:"client",yml:"client"`
}

func (od *OrderData) GetJSON() ([]byte, error) {
	return json.MarshalIndent(*od, "", "\t")
}

func (od *OrderData) GetYaML() ([]byte, error) {
	return yaml.Marshal(*od)
}

func FromJSON(d []byte) (OrderData, error) {
	od := new(OrderData)
	err := json.Unmarshal(d, &od)
	return *od, err
}

func FromYAML(d []byte) (OrderData, error) {
	od := new(OrderData)
	err := yaml.Unmarshal(d, &od)
	return *od, err
}

func main() {
	od := OrderData{
		OrderUid:    "asd123",
		Destination: "Izhevsk",
		Weight:      228,
		Client: OrderClient{
			Name:     "Ivan",
			Lastname: "Lebedev",
			Age:      23,
			UserUid:  "9182387jasihfj",
		},
	}

	fyaml, _ := os.Create("order.yml")

	ymlbytes, _ := od.GetYaML()
	fyaml.Write(ymlbytes)
	fyaml.Close()

	var odd OrderData
	raw, _ := ioutil.ReadFile("order.yml")
	yaml.Unmarshal(raw, &odd)
	fmt.Printf("%v\n", odd)

	fjson, _ := os.Create("order.json")

	jsbytes, _ := odd.GetJSON()
	fjson.Write(jsbytes)

	fjson.Close()

}
