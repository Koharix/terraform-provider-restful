package restful

import (
	"testing"

	"log"
    "io/ioutil"
)

var (
	resBody = []byte(`{
		"name":"John",
		"age":30,
		"cars":[
		   {
			  "name":"Ford",
			  "models":[
				 "Fiesta",
				 "Focus",
				 "Mustang"
			  ]
		   },
		   {
			  "name":"BMW",
			  "models":[
				 "320",
				 "X3",
				 "X5"
			  ]
		   },
		   {
			  "name":"Fiat",
			  "models":[
				 "500",
				 "Panda"
			  ]
		   }
		]
	 }`)

	 validExpResBody = [][]byte{
		[]byte(`{
			"name":"John",
			"age":30,
			"cars":[
			   {
				  "name":"Ford",
				  "models":[
					 "Fiesta",
					 "Focus",
					 "Mustang"
				  ]
			   },
			   {
				  "name":"BMW",
				  "models":[
					 "320",
					 "X3",
					 "X5"
				  ]
			   },
			   {
				  "name":"Fiat",
				  "models":[
					 "500",
					 "Panda"
				  ]
			   }
			]
		 }`),
		[]byte(`{
			"age":30,
			"cars":[
			{
				"name":"Ford",
				"models":[
					"Fiesta",
					"Focus",
					"Mustang"
				]
			},
			{
				"name":"BMW",
				"models":[
					"320",
					"X3",
					"X5"
				]
			}
			]
		}`),
		[]byte(`{
			"age":30,
			"cars":[
			{
				"name":"BMW",
				"models":[
					"320",
					"X3",
					"X5"
				]
			}
			]
		}`),
		[]byte(`{
			"age":30
		}`),
	}

	invalidExpResBody = [][]byte{
		[]byte(`{
			"name":"Johnd",
			"age":30,
			"cars":[
			   {
				  "name":"Ford",
				  "models":[
					 "Fiesta",
					 "Focus",
					 "Mustang"
				  ]
			   },
			   {
				  "name":"BMW",
				  "models":[
					 "320",
					 "X3",
					 "X5"
				  ]
			   },
			   {
				  "name":"Fiat",
				  "models":[
					 "500",
					 "Panda"
				  ]
			   }
			]
		 }`),
		[]byte(`{
			"age":30,
			"cars":[
			{
				"name":"Ford",
				"models":[
					"Fiesta",
					"Focus",
					"Mustang"
				]
			},
			{
				"name":"BMW",
				"models":[
					"321",
					"X3",
					"X5"
				]
			}
			]
		}`),
		[]byte(`{
			"age":30,
			"cars":[
			{
				"name":"BMW",
				"modelz":[
					"320",
					"X3",
					"X5"
				]
			}
			]
		}`),
		[]byte(`{
			"age":"30"
		}`),
	}
)

func TestCompareStatusCode(t *testing.T){
	log.SetOutput(ioutil.Discard)
	err := compareStatusCode(200, 200)
	if err != nil {
		t.Error("Test Failed: Status codes are equal but compareStatusCode returned error")
	}

	err = compareStatusCode(200, 404)
	if err == nil {
		t.Error("Test Failed: Status codes are not equal but compareStatusCode did not return an error")
	}
}

func TestParseJson(t *testing.T){
	log.SetOutput(ioutil.Discard)
	m, _, _ := parseJson(resBody)
	if m["name"] != "John" {
		t.Error("Test Failed: m[\"name\"] != \"John\"")
	}
	if m["cars"].([]interface{})[1].(map[string]interface{})["models"].([]interface{})[2] != "X5" {
		t.Error("Test Failed: m[\"cars\"][1][\"models\"][2] != \"X5\"")
	}
}

func TestResChecker(t *testing.T){
	log.SetOutput(ioutil.Discard)
	err := resChecker(resBody, resBody, 200, 201)
		if err == nil {
			t.Error("Failed Test: compareStatusCode should have returned error.")
		}
	for _, expResBody := range validExpResBody {
		err := resChecker(resBody, expResBody, 200, 200)
		if err != nil {
			t.Error(err)
		}
	}	
	for _, expResBody := range invalidExpResBody {
		err := resChecker(resBody, expResBody, 200, 200)
		if err == nil {
			t.Error("Failed Test: compareMaps and compareSlices should have returned error.")
		}
	}
}