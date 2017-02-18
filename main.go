package main

import (
   "net/http"
   "fmt"
   "io/ioutil"
)

func main() {
   resp , err := http.Get("http://api.theysaidso.com/qod.json")
   if err != nil {
     fmt.Println(err)
     return
   }

   defer resp.Body.Close()

   content, err := ioutil.ReadAll(resp.Body)
   if err != nil {
     fmt.Println(err)
      return
   }
   fmt.Println(string(content))
}

