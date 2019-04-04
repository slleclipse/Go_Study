package valuecow

import (
	"errors"
	"fmt"
	"sync/atomic"
)

// ConcurrentArray 代表并发安全的整数数组接口
type ConcurrentArray interface {
	// Set 用于设置指定索引上的元素值
	Set(index uint32, elem int) (err error)
	// Get 用于获取指定索引上的元素值。
	Get(index uint32) (elem int, err error)
	// Len用于获取数组长度
	Len()uint32
}

type intArray struct {
	length uint32
	val atomic.Value
}

func NewConcurrentArray(length uint32) ConcurrentArray{
	array := intArray{}
	array.length = length
	array.val.Store(make([]int, array.length))
	return &array
}

func (array *intArray) Set(index uint32, elem int) (err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValue(); err != nil {
		return
	}

	// 不要这样做！否则会形成竞态条件！
	// oldArray := array.val.Load().([]int)
	// oldArray[index] = elem
	// array.val.Store(oldArray)

	newArray := make([]int, array.length)
	copy(newArray, array.val.Load().([]int))
	newArray[index] = elem
	array.val.Store(newArray)
	return
}

func (array *intArray) Get(index uint32) (elem int, err error) {
	if err = array.checkIndex(index); err != nil {
		return
	}
	if err = array.checkValue(); err != nil {
		return
	}
	elem = array.val.Load().([]int)[index]
	return
}

func (array *intArray) Len() uint32 {
	return array.length
}

func (array *intArray) checkIndex(index uint32) (err error){
	if index > array.length - 1 || index < 0 {
		return fmt.Errorf("Index out range [0, %d]", array.length)
	}
	return nil
}

func (array *intArray) checkValue() (err error) {
	v := array.val.Load()
	if v == nil {
		return errors.New("Invalid int array!")
	}
	return nil
}


