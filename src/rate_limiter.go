package main

/*
*
* rate limit "token bucket" simple middleware
 */

import (
        "time"
        "net/http"
        )

import (
        "github.com/abiosoft/river"
        )

const (
       requestSize = 10 // request size in tokens
       bucketSize = 15
)
var tokenBucket *Bucket

func rateTokenBucketMid(c *river.Context) {
    
    if tokenBucket == nil {
        //create bucket and fill with tokents
        tokenBucket = NewBucket(1 * time.Second, bucketSize)
        tokenBucket.AddTokens(bucketSize)
    }
    
    // every request cost requestSize tokens.
    // ask to spend tokens before processing request itself
    // if tokens are not enought than TooManyRequests error triggered out
    bucketNotEmpty := <-tokenBucket.SpendToken(requestSize)
    
    if bucketNotEmpty != nil {
        c.Render(http.StatusTooManyRequests, river.M{"error": "Rate limit reached. Try later."})
        return
    }
    
    c.Register(bucketNotEmpty)
	c.Next()
}

