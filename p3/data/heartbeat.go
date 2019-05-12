package data

//HeartBeatData heatbeat data strut
type HeartBeatData struct {
	IfNewBlock  bool   `json:"ifNewBlock"`
	ID          int32  `json:"id"`
	BlockJSON   string `json:"blockJson"`
	PeerMapJSON string `json:"peerMapJson"`
	Addr        string `json:"addr"`
	Hops        int32  `json:"hops"`
}

//NewHeartBeatData initial the heartbeat data
func NewHeartBeatData(ifNewBlock bool, id int32, blockJSON string, peerMapJSON string, addr string) HeartBeatData {
	heartBeatData := HeartBeatData{IfNewBlock: ifNewBlock, ID: id, BlockJSON: blockJSON, PeerMapJSON: peerMapJSON, Addr: addr, Hops: 3}
	return heartBeatData
}
