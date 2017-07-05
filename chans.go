package roomManager

import "sync"

var messageChannel map[string]chan nodeMessage
var messageChannelLock sync.RWMutex

func init() {
	messageChannel = make(map[string]chan nodeMessage)
}

func sendMessageToChannel(roomId string, nm nodeMessage) error {
	messageChannelLock.RLock()
	//如果房间不存在，创建一个房间
	if c, ok := messageChannel[roomId]; ok {
		c <- nm
		messageChannelLock.RUnlock()
	} else {
		messageChannelLock.RUnlock()
		messageChannelLock.Lock()
		//创建房间通道
		messageChannel[roomId] = make(chan nodeMessage, 1024)
		messageChannel[roomId] <- nm
		messageChannelLock.Unlock()
		//创建房间实例
		roomObj := &RoomInfo{}
		roomObj.RoomID = roomId
		roomObj.Rows = make([]*RowList, 0, 1024)

		go daemonReciver(messageChannel[roomId], roomObj)
	}
	return nil
}
