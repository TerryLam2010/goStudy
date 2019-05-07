package parser

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestParseCitylist(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s",contents)
	const resultSize = 470
	expectedUrls := []string {"http://www.zhenai.com/zhenghun/aba","http://www.zhenai.com/zhenghun/baicheng1","http://www.zhenai.com/zhenghun/cangzhou"}
	//expectedCities := []string {"","",""}
	result := ParseCitylist(contents)
	if len(result.Requestes) != resultSize {
		t.Errorf("result should hava %d requests;but had %d",resultSize,len(result.Requestes))
	}

	for i,url := range expectedUrls {
		if result.Requestes[i].Url != url {
			t.Errorf("expected url #%d : %s;but was %s",i,url,result.Requestes[i].Url)
		}
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should hava %d requests;but had %d",resultSize,len(result.Items))
	}
}
