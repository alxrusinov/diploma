package model

type Process string

const (
	New        Process = "NEW"
	Registered Process = "REGISTERED"
	Invalid    Process = "INVALID"
	Processing Process = "PROCESSING"
	Processed  Process = "PROCESSED"
)
