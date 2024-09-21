package model

type Process string

const (
	Registered Process = "REGISTERED"
	Invalid    Process = "INVALID"
	Processing Process = "PROCESSING"
	Processed  Process = "PROCESSED"
)
