package main

type Company struct {
    ID        int   `json:"id,omitempty"`
    Name string   `json:"name,omitempty"`
    Status  string   `json:"status,omitempty"`
}

type Companies []Company
