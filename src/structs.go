package main

type Policy struct {
	PolicyName     string    `json:"PolicyName"`
	PolicyDocument PolicyDoc `json:"PolicyDocument"`
}

type PolicyDoc struct {
	Version   string    `json:"Version"`
	Statement Statement `json:"Statement"`
}

type Statement interface{}
