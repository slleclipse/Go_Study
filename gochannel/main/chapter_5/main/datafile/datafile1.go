package datafile

import (
	"errors"
	"github.com/astaxie/beego/logs"
	"io"
	"os"
	"sync"
	"sync/atomic"
)

// Data 代表数据的类型
type Data []byte

// DataFile 代表数据文件的接口类型
type DataFile interface {
	// Read会读取一个数据块
	Read() (rsn int64, d Data, err error)

	// Write 会写入一个数据块
	Write(d Data) (wsn int64, err error)

	// Read会读取一个数据块
	Read2() (rsn int64, d Data, err error)

	// Read会读取一个数据块
	Read3() (rsn int64, d Data, err error)

	// Write 会写入一个数据块
	Write2(d Data) (wsn int64, err error)

	// Write 会写入一个数据块
	Write3(d Data) (wsn int64, err error)

	// RSN 会获取最后读取的数据块的序列号
	RSN() int64

	// WSN 会获取最后写入的数据块的序列号
	WSN() int64

	// DataLen 会获取数据块的长度
	DataLen() uint32

	// Close 会关闭数据文件
	Close() error
}

type myDataFile struct {
	f         *os.File     //文件
	fmutex    sync.RWMutex // 被用于文件的读写锁
	condition *sync.Cond   // 条件变量
	woffset   int64        // 写操作需要用到的偏移量
	roffset   int64        // 读操作需要用到的偏移量
	wmutex    sync.Mutex   // 写操作需要用到的互斥锁
	rmutex    sync.Mutex   // 读操作需要用到的互斥锁
	dataLen   uint32       // 数据块长度
}

func (df *myDataFile) Read() (rsn int64, d Data, err error) {
	// 读取并更新读取偏移量
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	// 读取一个数据块
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)

	for {
		df.fmutex.RLock()
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.fmutex.RUnlock()
				continue
			}
			df.fmutex.RUnlock()
			return
		}
		d = bytes
		logs.Info("Read data: ", bytes)
		df.fmutex.RUnlock()
		return
	}
}

// 使用条件变量
func (df *myDataFile) Read2() (rsn int64, d Data, err error) {
	// 读取并更新读取偏移量
	var offset int64
	df.rmutex.Lock()
	offset = df.roffset
	df.roffset += int64(df.dataLen)
	df.rmutex.Unlock()

	// 读取一个数据块
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()

	for {
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.condition.Wait()
				continue
			}
			return
		}
		d = bytes
		logs.Info("Read data: ", bytes)
		return
	}
}

// 使用条件变量
func (df *myDataFile) Read3() (rsn int64, d Data, err error) {
	// 读取并更新读取偏移量
	var offset int64
	for {
		offset = atomic.LoadInt64(&df.roffset)
		if atomic.CompareAndSwapInt64(&df.roffset, offset, offset + int64(df.dataLen)) {
			break
		}
	}

	// 读取一个数据块
	rsn = offset / int64(df.dataLen)
	bytes := make([]byte, df.dataLen)
	df.fmutex.RLock()
	defer df.fmutex.RUnlock()

	for {
		_, err = df.f.ReadAt(bytes, offset)
		if err != nil {
			if err == io.EOF {
				df.condition.Wait()
				continue
			}
			return
		}
		d = bytes
		logs.Info("Read data: ", bytes)
		return
	}
}

func (df *myDataFile) Write(d Data) (wsn int64, err error) {
	//读取并更新写偏移量
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	//写入一个数据块
	wsn = offset / int64(df.dataLen)
	var bytes []byte

	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	logs.Info("Writing data: ", bytes)
	_, err = df.f.Write(bytes)
	df.condition.Signal()
	return
}

// 使用条件变量
func (df *myDataFile) Write2(d Data) (wsn int64, err error) {
	//读取并更新写偏移量
	var offset int64
	df.wmutex.Lock()
	offset = df.woffset
	df.woffset += int64(df.dataLen)
	df.wmutex.Unlock()

	//写入一个数据块
	wsn = offset / int64(df.dataLen)
	var bytes []byte

	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	logs.Info("Writing data: ", bytes)
	_, err = df.f.Write(bytes)
	df.condition.Signal()
	return
}

// 使用原子操作
func (df *myDataFile) Write3(d Data) (wsn int64, err error) {
	//读取并更新写偏移量
	var offset int64
	for {
		offset = atomic.LoadInt64(&df.woffset)
		if atomic.CompareAndSwapInt64(&df.woffset, offset, offset + int64(df.dataLen)) {
			break
		}
	}

	//写入一个数据块
	wsn = offset / int64(df.dataLen)
	var bytes []byte

	if len(d) > int(df.dataLen) {
		bytes = d[0:df.dataLen]
	} else {
		bytes = d
	}
	df.fmutex.Lock()
	defer df.fmutex.Unlock()
	logs.Info("Writing data: ", bytes)
	_, err = df.f.Write(bytes)
	df.condition.Signal()
	return
}

func (df *myDataFile) RSN() int64 {
	df.rmutex.Lock()
	defer df.rmutex.Unlock()
	return df.roffset / int64(df.dataLen)
}

func (df *myDataFile) WSN() int64 {
	df.wmutex.Lock()
	defer df.wmutex.Unlock()
	return df.woffset / int64(df.dataLen)
}

func (df *myDataFile) DataLen() uint32 {
	return df.dataLen
}

func (df *myDataFile) Close() error {
	if df.f == nil {
		return nil
	}
	return df.f.Close()
}

func NewDataFile(path string, dataLen uint32) (DataFile, error) {
	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	if dataLen == 0 {
		return nil, errors.New("Invalid data length!")
	}

	df := &myDataFile{
		f:       f,
		dataLen: dataLen}
	df.condition = sync.NewCond(df.fmutex.RLocker())
	return df, nil
}
