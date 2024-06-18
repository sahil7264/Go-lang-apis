package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var myurl string = "http://localhost:4000"

func creategetRequest() {
	res, err := http.Get(myurl) // get
	if err != nil {             // check errr
		fmt.Println("Error Something")
		panic(err)
	}
	// print res
	fmt.Println("Status :", res.StatusCode)
	fmt.Println("Content Length :", res.ContentLength)

	defer res.Body.Close()                   // close connections
	content, err := ioutil.ReadAll(res.Body) // read data
	if err != nil {
		panic(err)
	}
	// fmt.Println(content)

	var resString strings.Builder
	bytecount, _ := resString.Write(content) // just to get control over response such to use sting builder methods
	fmt.Println(bytecount)
	fmt.Println(resString.String())
}
func createpostRequest() {
	myurl = myurl + "/post"
	fmt.Println(myurl)

	// Create dummy json data
	body := strings.NewReader(
		`
	{
     	"a":"4",
		"b":"6"
	}
	`,
	)
	res, err := http.Post(myurl, "application/json", body) // content-type is app/json
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(res.Body) // read data same as await we are doing
	fmt.Println("Status :", res.StatusCode)
	fmt.Println("Response :", string(content))
}

func performFormRequest() {
	myurl = myurl + "/form"
	data := url.Values{}
	data.Add("fname", "sahil")
	data.Add("fmname", "rajendra")
	data.Add("lname", "shile")

	res, err := http.PostForm(myurl, data) // it knows that form data in json
	if err != nil {
		panic(err)
	}

	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))

	defer res.Body.Close()
}

/* Create JSON Data*/

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Password string   `json:"-"` // - entire remove from json
	Platform string   `json:"platform"`
	Tags     []string `json:"tags,omitempty"` // omitempty basically not show if null no value
}

func encodeJson() {
	fmt.Println("Creating Json Data")
	lcCoarse := []course{
		{"CP", 99, "pass@123", "codeforces", []string{"Dsa"}},
		{"MERN", 99, "pass@123", "codeforces", []string{"Dsa"}},
		{"PERN", 99, "pass@123", "codeforces", []string{"Dsa"}},
		{"JAVA", 99, "pass@123", "codeforces", []string{"Dsa"}},
		{"Cloud", 99, "pass@123", "codeforces", nil},
	}

	fmt.Println(lcCoarse)

	// packaging this data as json
	finalJson, err := json.Marshal(lcCoarse)
	if err != nil {
		panic(err)
	}

	//Easy to read format , Indentation . \t giving space when indentation
	finalJsonNew, err := json.MarshalIndent(lcCoarse, "c", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(finalJson)
	fmt.Printf("%s\n", finalJson)
	fmt.Printf("%s\n", finalJsonNew)
	// As stuct having data key start with capital but json having small
	// So use alice in struct `json:name`
}

func decodeJson() {

	//Data we get from web is in byte format
	jsonData := []byte(
		`
		{
		              "name": "CP",
		              "price": 99,
		              "platform": "codeforces",
		              "tags": [
		                      "Dsa"
		              ]
		      }
		`,
	)

	var lcCoarse course
	//check json valid or not
	checkValid := json.Valid(jsonData)
	fmt.Println(checkValid)
	if checkValid {
		json.Unmarshal(jsonData, &lcCoarse)
		fmt.Print(lcCoarse)
	} else {
		fmt.Println("Not Valid Json")
	}
	// fmt.Println(string(jsonData))

	//Some case we want to add data to key value pair not struct
	fmt.Println()
	var mydata map[string]interface{} // value can to differed to use interface
	json.Unmarshal(jsonData, &mydata)
	fmt.Println(mydata)

	for i, e := range mydata {
		fmt.Printf("key : %v and value is %v \n", i, e)
		fmt.Printf("key : %T and value is %T \n", i, e)
	}
}
func main() {
	fmt.Println("Hello Sahil")
	// creategetRequest()
	// createpostRequest()
	// performFormRequest()
	// encodeJson()
	decodeJson()
}
