package roomManager

import "sync"

type RowList struct {
	Lock        sync.RWMutex //方式写入的时候冲突
	FrontNode   *ReciveNode  //第一个节点的指针地址
	BackNode    *ReciveNode  //最后一个节点的指针地址d
	Length      uint64       //当前行节点数量
	CurrentRoom *RoomInfo    //当前所在房间指针地址
}
