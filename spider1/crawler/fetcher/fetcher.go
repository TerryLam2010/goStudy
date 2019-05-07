package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string,)([]byte,error){
	<-rateLimiter
	cookieJar, _ := cookiejar.New(nil)
	req,err := http.NewRequest("GET",url,nil)
	client := &http.Client{
		Jar: cookieJar,
	}
	req.Header.Add("User-Agent","Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.62 Safari/537.36")
	resp, err := client.Do(req)
	//resp,err := http.Get("https://www.zhenai.com/zhenghun")
	//req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
		fmt.Errorf("wrong status code : %d",resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewEncoder())
	return ioutil.ReadAll(utf8Reader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding  {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Printf("Fetcher error: %v",err)
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}