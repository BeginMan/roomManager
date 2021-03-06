package roomManager

//更换房间标记
//因为前面已经修改了RoomID，所以这里只需要把接受者节点的指针重新对接到新房间的节点上，就可以了
//新的房间更新逻辑：
//1.扫描所有列，找到长度小于最大长度的，加入队尾
//2.修改结点中索引信息
func changeRoom(roomInfo *RoomInfo, node *ReciveNode) {
	roomInfo.Lock.Lock()
	defer roomInfo.Lock.Unlock()
	//如果房间没有列，则创建一个列
	if len(roomInfo.Rows) == 0 {
		row := &RowList{}
		row.Nodes = make([]*ReciveNode, 0, ROW_LENGTH)
		roomInfo.Rows = append(roomInfo.Rows, row)
	}
	//设定是否添加完成标记
	addSuccess := false
	//寻找一个未满的列，把新的节点加进去
	for _, v := range roomInfo.Rows {
		if len(v.Nodes) < ROW_LENGTH {
			v.Nodes = append(v.Nodes, node)
			addSuccess = true
			break
		}
	}
	//如果前面所有的列都是满的，则创建一个新列，把节点添加进去
	if !addSuccess {
		row := &RowList{}
		row.Nodes = make([]*ReciveNode, 0, ROW_LENGTH)
		roomInfo.Rows = append(roomInfo.Rows, row)
		row.Nodes = append(row.Nodes, node)
	}
}
