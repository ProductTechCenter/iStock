package service

import (
	"time"
	"errors"
)

type Manager struct {
	Match_que chan interface{}
	Clear_que chan interface{}
	Sequence_que chan interface{}
	Source_data_que chan interface{}
}

// global singleton instance
var manager *Manager

// 饿汉单例模式
func init() {
	manager = &Manager{
		make(chan interface{}),
		make(chan interface{}),
		make(chan interface{}),
		make(chan interface{}),
	}
}

func GetInstance() *Manager{
	return manager
}

// 懒汉单例模式
//var once sync.Once
//func GetInstance() *Manager {
//	once.Do(func() {
//		manager = &Manager {
//			make(chan interface{}),
//			make(chan interface{}),
//			make(chan interface{}),
//			make(chan interface{}),
//		}
//	})
//
//	return manager
//}

//func (this *Manager)sendMsg(val interface{}, )

// 发送一个msg到定序系统
// msg : 消息
// tmout : 超时时间
func Send2Senquence(msg interface{}, tmout int) error{
	ch := manager.Sequence_que

	select {
		case ch <- msg:
			return nil
		case <- time.After(time.Duration(tmout) * time.Second ):
			return errors.New("time out.")
	}
}

// 发送一个消息到 撮合系统
// msg : 消息
// tmout : 超时时间
func Send2Match(msg interface{}, tmout int) error{
	ch := manager.Match_que

	select {
	case ch <- msg:
		return nil
	case <- time.After(time.Duration(tmout) * time.Second ):
		return errors.New("time out.")
	}
}

// 发送一个消息到 清算系统
// msg : 消息
// tmout : 超时时间
func Send2Clearing(msg interface{}, tmout int) error{
	ch := manager.Clear_que

	select {
	case ch <- msg:
		return nil
	case <- time.After(time.Duration(tmout) * time.Second ):
		return errors.New("time out.")
	}
}

// 发送一个消息到 数据源
// msg : 消息
// tmout : 超时时间
func Send2Source(msg interface{}, tmout int) error{
	ch := manager.Source_data_que

	select {
	case ch <- msg:
		return nil
	case <- time.After(time.Duration(tmout) * time.Second ):
		return errors.New("time out.")
	}
}

//func Init() {
//	//  初始化各個系統
//	sequence.Init()
//	clearing.Init()
//	match.Init()
//	data_source.Init()
//}