package main

import (
    "github.com/gorilla/websocket"
    "net/http"
    "log"
    "fmt"
    "encoding/base64"
    "time"
    "io/ioutil"
);

func main() {
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

    var img64 [] byte

    res := map[string]interface{}{}
    for {
        if err = ws.ReadJSON(&res); err != nil {
            if err.Error() == "EOF" {
                return
            }
            // ErrShortWrite means a write accepted fewer bytes than requested then failed to return an explicit error.
            if err.Error() == "unexpected EOF" {
                return
            }
            fmt.Println("Read : " + err.Error())
            return
        }

        res["a"] = "a"
        log.Println(res)

        for {
            files, _ := ioutil.ReadDir("./images");
            for _, f := range files {
                img64, _ = ioutil.ReadFile("./images/" + f.Name())
                str := base64.StdEncoding.EncodeToString(img64)
                res["img64"] = str

                if err = ws.WriteJSON(&res); err != nil {
                    fmt.Println("watch dir - Write : " + err.Error())
                    //return
                }
                time.Sleep(50 * time.Millisecond);
            }
            time.Sleep(50 * time.Millisecond);
        }
    }
}
