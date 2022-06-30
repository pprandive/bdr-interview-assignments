package lib

import (
    _ "fmt"
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

