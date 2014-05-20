package main

import (
    "github.com/gorilla/websocket"
    "net/http"
    "log"
    "fmt"
    //"io/ioutil"
    "encoding/base64"
    "github.com/go-av/curl"
    "time"
);

func main() {

    //var img64 []byte
    //img64, _ = ioutil.ReadFile("/home/ubuntu/mygo/src/pushImage/google.png")
    //str := base64.StdEncoding.EncodeToString(img64)
    //fmt.Println(str)


    http.HandleFunc("/connws/", ConnWs)
    err := http.ListenAndServe(":9090", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }

}

func ConnWs(w http.ResponseWriter, r *http.Request) {
    ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
    if _, ok := err.(websocket.HandshakeError); ok {
        http.Error(w, "Not a websocket handshake", 400)
            return
    } else if err != nil {
        log.Println(err)
            return
    }

    res := map[string]interface{}{}
    for {
        if err = ws.ReadJSON(&res); err != nil {
            if err.Error() == "EOF" {
                return
            }
            // ErrShortWrite means that a write accepted fewer bytes than requested but failed to return an explicit error.
            if err.Error() == "unexpected EOF" {
                return
            }
            fmt.Println("Read : " + err.Error())
            return
        }

        res["a"] = "a"
        log.Println(res)

        for {
            //_, b := curl.Bytes("http://2d3bd0383620907d11324ede6e9f2b57.r0202.relay.yun.netgear.cn:80/ws/api/requestProxy/172.16.0.6/live/snapshot?p=cvJfXTMwNQRl8EGzMYvHgQl7H2x7bAcD80ckjIYr0K1IjSqRVpB1cY4FwsEhL3So")
            _, b := curl.Bytes("https://www.google.com.tw/images/srpr/logo11w.png")
            str2 := base64.StdEncoding.EncodeToString(b)
            res["img64"] = str2

            if err = ws.WriteJSON(&res); err != nil {
                fmt.Println("watch dir - Write : " + err.Error())
                //return
            }
            time.Sleep(50 * time.Millisecond);
        }
    }

}
