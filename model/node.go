package model

type Node struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	RaftAddress string `json:"raft_address"`
	Term        int32  `json:"term"`
	Role        string `json:"role"`
	LogTerm     int32  `json:"log_term"`
	LogIndex    int32  `json:"log_index"`
}

type NodeInfoResponse struct {
	LeaderName string  `json:"leader_name"`
	NodeList   []*Node `json:"servers"`
}
