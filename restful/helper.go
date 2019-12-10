package main

import (
	"fmt"
	"reflect"
	"time"
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"errors"
	"strings"
	"log"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func restCall(method string, uri string, headers []interface{},  reqBody []byte) ([]byte, int, error) {
	
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(method, uri, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, 0, err
	}

	for _, value := range headers {
		header := strings.Split(value.(string), ":")
		req.Header.Set(header[0], header[1])
	}
	
	res, err := client.Do(req)
	if err != nil {
		return nil, 0, err
	}

	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, 0, err
	}

	log.Println("Response Status Code: " + fmt.Sprint(res.StatusCode))
	log.Println("Response Body: " + string(resBody))

	return resBody, res.StatusCode, nil
}

func resChecker(resBody []byte, expResBody []byte, resStatusCode int, expResStatusCode int) error {
	log.Println("function resChecker() started")

	if expResStatusCode != 0 {
		err := compareStatusCode(resStatusCode, expResStatusCode)
		if err != nil {
			return err
		}
	}

	log.Println("Expected response: " + string(expResBody))

	if len(expResBody) > 0 {
		expResMap, expResSlice, err := parseJson(expResBody)
		if err != nil {
			return err
		}
		resMap, resSlice, err := parseJson(resBody)
		if err != nil {
			return err
		} else if resMap != nil {
			err = compareMaps(expResMap, resMap)
		} else if resSlice !=  nil {
			err = compareSlices(expResSlice, resSlice)
		}
		if err != nil {
			return errors.New(err.Error() + "\n" + string(expResBody) + "\nNOT IN\n" + string(resBody))
		}
	}
	return nil
}

func compareStatusCode(resStatusCode int, expResStatusCode int) error {
	log.Println("function compareStatusCode() started")
	if resStatusCode != expResStatusCode {
		return errors.New("ERROR! Expected Status Code: " + fmt.Sprint(expResStatusCode) + "; Returned Status Code: " + fmt.Sprint(resStatusCode))
	}
	return nil
}

func compareMaps(map1 map[string]interface{}, map2 map[string]interface{}) error {
	for key, _ := range map1 {
		if reflect.ValueOf(map1[key]).Kind() == reflect.Slice &&  reflect.ValueOf(map2[key]).Kind() == reflect.Slice {
			err := compareSlices(map1[key].([]interface{}), map2[key].([]interface{}))
			if err != nil {
				return err
			}
		} else if reflect.ValueOf(map1[key]).Kind() == reflect.Map &&  reflect.ValueOf(map2[key]).Kind() == reflect.Map {
			err := compareMaps(map1[key].(map[string]interface{}), map2[key].(map[string]interface{}))
			if err != nil {
				return err
			}
		} else {
			if map1[key] != map2[key] {
				return errors.New("ERROR!")
			}
		}
	}
	return nil
}

func compareSlices(slice1 []interface{}, slice2 []interface{}) error {
	for index1, _ := range slice1 {
		for index2, _ := range slice2 {
			if reflect.ValueOf(slice1[index1]).Kind() == reflect.Slice &&  reflect.ValueOf(slice2[index2]).Kind()  == reflect.Slice {
				err := compareSlices(slice1[index1].([]interface{}), slice2[index2].([]interface{}))
				if err == nil {
					break
				} else if err != nil && len(slice2) == index2 + 1 {
					return err
				}
			} else if reflect.ValueOf(slice1[index1]).Kind() == reflect.Map &&  reflect.ValueOf(slice2[index2]).Kind()  == reflect.Map {
				err := compareMaps(slice1[index1].(map[string]interface{}), slice2[index2].(map[string]interface{}))
				if err == nil {
					break
				} else if err != nil && len(slice2) == index2 + 1 {
					return err
				}
			} else if slice1[index1] == slice2[index2] {
				break
			} 
			if len(slice2) == index2 + 1 {
				return errors.New("ERROR!")
			}
		}
	}
	return nil
}

func parseJson(jsonBytes []byte) (map[string]interface{}, []interface{}, error) {
	log.Println("function parseJson() started")
	var jsonMap map[string]interface{}
	err := json.Unmarshal(jsonBytes, &jsonMap)
	if err != nil {
		var jsonSlice []interface{}
		err := json.Unmarshal(jsonBytes, &jsonSlice)
		if err != nil {
			return nil, nil, err
		}
		return nil, jsonSlice, nil
	}
	return jsonMap, nil, nil
}

func setOutputs(d *schema.ResourceData, resBody []byte, jsonOutputs []interface{}) error {
	log.Println("function setOutputs() started")
	var s []string
	for _, json := range jsonOutputs {
		resMap, _, err := parseJson(resBody)
		if err != nil {
			return err
		}
		s = append(s, resMap[json.(string)].(string))
	}
	d.Set("outputs", s)
	return nil
}