package main

import (
    "fmt"
    "net/http"
    _ "sync"

    "github.com/gin-gonic/gin"

    "dict/lib"
)

type dictionary struct{

    root *lib.Node
    //Read write lock
    //dictionaryGuard sync.RWMutex
}

var dict dictionary

func main(){

    dict = dictionary{}
    dict.root = lib.NewNode()

    router := gin.Default()
    router.GET("/insert/:key", insert)
    router.GET("/search/:key", search)
    router.GET("/delete/:key", delete)

    router.Run("localhost:8080")

    fmt.Println("Listening on post 8080")
}

// getAlbums responds with the list of all albums as JSON.
func insert(c *gin.Context) {

    key := c.Param("key")

    //dict.dictionaryGuard.Lock() //Write lock
    lib.InsertNewNode(dict.root, key)
    //dict.dictionaryGuard.Unlock()

    c.IndentedJSON(http.StatusOK, key)
}

func search(c *gin.Context) {

    key := c.Param("key")

    //dict.dictionaryGuard.RLock()  //A read lock
    found := lib.Search(dict.root, key)
    //dict.dictionaryGuard.RUnlock()

    message := ""
    if found == true {
        message = key+" is found."
    } else{
        message = key+" is not found."
    }
    c.IndentedJSON(http.StatusOK, gin.H{"message": message})
}

func delete(c *gin.Context) {
    key := c.Param("key")

    //dict.dictionaryGuard.Lock()  //A read lock
    lib.Delete(dict.root, key, 0)
    //dict.dictionaryGuard.Unlock()

    /*message := ""
    if found == true {
        message = key+" is found."
    } else{
        message = key+" is not found."
    }*/
    c.IndentedJSON(http.StatusOK, "")
}
