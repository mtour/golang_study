package main

import "fmt"
import (
        "snowflake4go"
        "time"
)

func main() {
        fmt.Println(snowflake4go.MaxWorkerId)


        sf,_ := snowflake4go.NewSnowFlake(1)
        for i := 1; i < 100; i++ {
                time.Sleep(time.Nanosecond*1000*1000)
                fmt.Println(sf.Next())
        }
}
