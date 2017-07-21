package main

import (
	"log"
	"net/http"
	"io"
	"encoding/json"
	"math/rand"
	"strconv"
	"os"
	"bufio"
	"strings"
)

var (
    data map[string]string = make(map[string]string)
)

func main() {
	for k,_:=range data{
		//log.Println(k,v)
		http.HandleFunc("/rest/get/"+k, getHandler)
		http.HandleFunc("/rest/set/"+k, setHandler)
	}
	http.ListenAndServe(":8089",nil)
}

func getHandler(resp http.ResponseWriter,req *http.Request)  {
	uri := ""
	uris := strings.Split(req.RequestURI,"/")
	if len(uris) > 0{
		uri = uris[len(uris)-1]
	}else {
		uri = req.RequestURI
	}
	log.Println("get Handler uri=",uri,"value=",data[uri])
	resp.Header().Add("Content-Type","application/json")
	resp.Header().Add("Access-Control-Allow-Origin","*")
	if uri == "dataList"{
		lm := make([]map[string]string,0)
		_ = json.Unmarshal([]byte(data[uri]),&lm)
		for _,l:=range lm{
			for k1,v1:=range l{
				_ = v1
				l[k1] = strconv.Itoa(rand.Intn(100))
			}
		}
		bs,_:=json.Marshal(lm)
		io.WriteString(resp,string(bs))
	}else {
		io.WriteString(resp,data[uri])
	}
}

func setHandler(resp http.ResponseWriter,req *http.Request)  {
	uri := ""
	uris := strings.Split(req.RequestURI,"/")
	if len(uris) > 0{
		uri = uris[len(uris)-1]
	}else {
		uri = req.RequestURI
	}

	log.Println("set Handler uri=",uri,"value=",data[uri])
	if uri == "dataList"{
	}else {
		req.ParseForm()
		log.Println(req.Form)
		uriData := req.Form.Get(uri)
		log.Println("new value = ",uriData)
		data[uri] = uriData
	}
	resp.Header().Add("Content-Type","application/json")
	resp.Header().Add("Access-Control-Allow-Origin","*")
	io.WriteString(resp,`{"success":"ok"}`)
}

func init(){
	file, err := os.Open("data.txt")
	defer file.Close()
	if err != nil {
		log.Println(err)
	}
	r := bufio.NewReader(file)
	for {
		b, _, err := r.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		s := strings.TrimSpace(string(b))
		kv := strings.Split(s, "=")
		if len(kv)==2{
			data[kv[0]] = kv[1]
		}
	}
}

