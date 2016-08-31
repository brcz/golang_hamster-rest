package main

/*
*
* rate limit "token bucket" simple middleware
*/

import "github.com/abiosoft/river"

func rateTokenBucketMid(c *river.Context) {
    
    c.Next()
}
