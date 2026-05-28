package rule

import (
	"testing"
	"time"
)

var tests = []struct {
	name string
	rule Rule
	wantErr bool
}{
	{
		name: "valid_rule",
		rule: Rule{
			ID: "a",
			Name: "api_rate_limit",
			Algorithm: AlgorithmTokenBucket,
			Limit: 1,
			Window: 60 * time.Second,
			CreatedAt: time.Now(),
		},
		wantErr: false,
	},
	{
		name: "missing_id",
		rule: Rule{
			ID: "",
			Name: "missing_id",
			Algorithm: AlgorithmTokenBucket,
			Limit: 1,
			Window: 60 * time.Second,						
			CreatedAt: time.Now(),	
		},		
		wantErr: true,
	},
	{
		name: "missing_name",
		rule: Rule{ 
			ID: "a",
			Name: "",
			Algorithm: AlgorithmTokenBucket,
			Limit: 1,
			Window: 60 * time.Second, 
			CreatedAt: time.Now(),
		},
		wantErr: true,
	},
	{
		name: "invalid_algorithm",
		rule: Rule{
			ID: "a",
			Name: "invalid_algorithm",
			Algorithm: "Invalid algorithm",
			Limit: 1,
			Window: 60 * time.Second,
			CreatedAt: time.Now(),
		},
		wantErr: true,
	},
	{
		name: "zero_limit",
		rule: Rule{
			ID: "a",
			Name: "zero_limit",
			Algorithm: AlgorithmTokenBucket,
			Limit: 0,
			Window: 60 * time.Second,
			CreatedAt: time.Now(),
		},
		wantErr: true,
	},
	{
		name: "negative_window",
		rule: Rule{
			ID: "a",
			Name: "negative_window",
			Algorithm: AlgorithmTokenBucket,
			Limit: 1,
			Window: -60 * time.Second,
			CreatedAt: time.Now(),
		},
		wantErr: true,			  
	},
	{
		name: "missing_created_at",
		rule: Rule{
			ID: "a",
			Name: "missing_created_at",
			Algorithm: AlgorithmTokenBucket,
			Limit: 1,
			Window: 60 * time.Second,
			CreatedAt: time.Time{},
		},
		wantErr: true,
	},
}

func TestRuleValidation(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.rule.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})			
	}
}