package common

type SerializableQueue interface {
	Push(interface{})
	Pop() interface{}
}

type PriorityQueue struct {
}

func (pq *PriorityQueue) Push(obj interface{}) {

}

func (pq *PriorityQueue) Pop() interface{} {
	return nil
}

type PickleFifoDiskQueue struct {
}

func (pq *PickleFifoDiskQueue) Push(obj interface{}) {

}

func (pq *PickleFifoDiskQueue) Pop() interface{} {
	return nil
}

type PickleLifoDiskQueue struct {
}

func (pq *PickleLifoDiskQueue) Push(obj interface{}) {

}

func (pq *PickleLifoDiskQueue) Pop() interface{} {
	return nil
}

type MarshalFifoDiskQueue struct {
}

func (mq *MarshalFifoDiskQueue) Push(obj interface{}) {

}

func (mq *MarshalFifoDiskQueue) Pop() interface{} {
	return nil
}

type MarshalLifoDiskQueue struct {
}

func (mq *MarshalLifoDiskQueue) Push(obj interface{}) {

}

func (mq *MarshalLifoDiskQueue) Pop() interface{} {
	return nil
}

type FifoMemoryQueue struct {
}

func (mq *FifoMemoryQueue) Push(obj interface{}) {

}

func (mq *FifoMemoryQueue) Pop() interface{} {
	return nil
}

type LifoMemoryQueue struct {
}

func (mq *LifoMemoryQueue) Push(obj interface{}) {

}

func (mq *LifoMemoryQueue) Pop() interface{} {
	return nil
}
