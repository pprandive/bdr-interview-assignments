package lib

import (
    "fmt"
)

const total_alphabets = 26

type Node struct{
    children [total_alphabets]*Node
    endOfWord bool
}

func NewNode()*Node{
    nn := new(Node)
    nn.endOfWord = false

    for i := 0; i< total_alphabets; i++{
        nn.children[i] = nil
    }
    return nn
}

func InsertNewNode(root *Node, key string) {

    temp := root
    for _, c := range key{
        index := c - 'a'
        if temp.children[index] == nil{
            temp.children[index] = NewNode()
        }

        temp = temp.children[index]
    }

    temp.endOfWord = true
}

func Search(root *Node, key string) bool{
    temp := root
    for _, c := range key{
        index := c - 'a'
        if temp.children[index] == nil{
            return false
        }

        temp = temp.children[index]
    }
    return temp.endOfWord
}

func IsEmpty(root *Node) bool{
    for i := 0; i< total_alphabets; i++{
        if root.children[i] != nil{
            return false
        }
    }
    return true
}

func Delete(root *Node, key string, dep int) *Node{
    if root == nil{
        return nil
    }
    fmt.Println("Inside delete")
    if dep == len(key){
        if root.endOfWord == true {
            root.endOfWord = false
        }

        if IsEmpty(root){
            root = nil
        }

        return root
    }

    index := key[dep] - 'a'
    root.children[index] = Delete(root.children[index], key, dep +1)

    if IsEmpty(root) && root.endOfWord == false{
        root = nil
    }

    return root
}
