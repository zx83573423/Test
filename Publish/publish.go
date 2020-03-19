package Publish

import (
	"sync"
	"time"
)

type (
	subscriber chan interface{}         //订阅者通道
	topicFunc  func(v interface{}) bool //主题过滤器
)

type Publisher struct {
	m           sync.RWMutex             //锁
	buffer      int                      //订阅数据长度
	timeout     time.Duration            //超时时间
	subscribers map[subscriber]topicFunc //通道和订阅过滤器
}

//构建一个发布者对象
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		m:           sync.RWMutex{},
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

//添加一个订阅者，订阅全部主题
func (pThis *Publisher) Subscribe() chan interface{} {
	return pThis.SubscribeTopic(nil)
}

//添加一个新的订阅者，订阅过滤器筛选过后的主题
func (pThis *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, pThis.buffer)
	pThis.m.Lock()
	pThis.subscribers[ch] = topic
	pThis.m.Unlock()
	return ch
}

//退出订阅
func (pThis *Publisher) Evict(sub chan interface{}) {
	pThis.m.Lock()
	defer pThis.m.Unlock()
	delete(pThis.subscribers, sub)
	close(sub)
}

//发布一个主题
func (pThis *Publisher) Publish(v interface{}) {
	pThis.m.Lock()
	pThis.m.Unlock()

	var wg sync.WaitGroup
	for sub, topic := range pThis.subscribers {
		wg.Add(1)
		pThis.sendTopic(sub, topic, v, &wg)
	}

	wg.Wait()
}

//关闭发布者对象，同时关闭所有订阅者通道
func (pThis *Publisher) Close() {
	pThis.m.Lock()
	defer pThis.m.Unlock()
	for sub := range pThis.subscribers {
		delete(pThis.subscribers, sub)
		close(sub)
	}
}

func (pThis *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(pThis.timeout):
	}
}
