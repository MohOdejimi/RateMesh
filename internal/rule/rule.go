package rule

import (
    "fmt"
    "time"
)

type AlgorithmType string 

const (
    AlgorithmTokenBucket AlgorithmType =  "token_bucket"
    AlgorithmSlidingWindow AlgorithmType = "sliding_window"
    AlgorithmFixedWindow AlgorithmType = "fixed_window"
)

type Rule struct {
    ID        string       
    Name      string       
    Algorithm AlgorithmType
    Limit     int          
    Window    time.Duration 
    CreatedAt time.Time     
}

func (r *Rule) Validate() error {
    if r.ID == "" {
        return fmt.Errorf("rule ID is required")
    }
    if r.Name == "" {
        return fmt.Errorf("rule name is required")
    }
    if r.Algorithm != AlgorithmTokenBucket && r.Algorithm != AlgorithmSlidingWindow && r.Algorithm != AlgorithmFixedWindow {
        return fmt.Errorf("Invalid algorithm type: %s", r.Algorithm)
    }
    if r.Limit <= 0 {
        return fmt.Errorf("limit must be greater than zero")
    }
    if r.Window <= 0 {
        return fmt.Errorf("window duration must be greater than zero")
    }
    if r.CreatedAt.IsZero() {
        return fmt.Errorf("created-at timestamp must be set")
    }
    return nil
}
