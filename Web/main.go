package main

/* Web Request*/

/* GET , URL MANIPLATIONS

const myurl string = "https://www.google.com/"
func main() {
   GET REQUEST
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is %T", res) // getting actual response not copy of that because *
	fmt.Println()
	res1, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	content := string(res1) // convert  to string
	fmt.Println(content)
	defer res.Body.Close()

// Handling URLS -> Getting info
// 1) &url.URL {key : value}Create url
// 2) url.Parse tp get url info
res, _ := url.Parse(myurl) // parsing to get some important info
fmt.Println(res.Scheme)    // https
fmt.Println(res.Host)      // www.google.com
fmt.Println(res.Path)      // /
fmt.Println(res.Port())    //
fmt.Println(res.RawQuery)  //

qparams := res.Query()
fmt.Printf("The type of query params ar %T \n", qparams) // key value pair
fmt.Println(qparams)                                     // map
for key, value := range qparams {
	fmt.Printf("The key is %v and the value is %v \n", key, value)
}
//Creating URLS
	// https://www.google.com/search?q=golang
	// https://www.google.com/search?q=golang#go
	// https://www.google.com/search?q=golang&lang=en
	// https://www.google.com/search?q=golang&lang=en#go

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "google.com",
		Path:    "/",
		RawPath: "sahil",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
*/

/*-------------------------Creating WEB SERVER------------*/

func main() {

}
