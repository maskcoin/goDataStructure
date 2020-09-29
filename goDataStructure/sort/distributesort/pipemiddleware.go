package distributesort

import (
	"datastructure/sort/quicksort"
	"encoding/binary"
	"fmt"
	"io"
	"math/rand"
	"time"
)

var startTime time.Time // 构造时间

func Init() {
	startTime = time.Now()
}

func UseTime() {
	fmt.Println(time.Since(startTime))
}

//内存排序
func InMemorySort(in <-chan int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		var data []int // 创建一个数组，存储数据并排序
		for v := range in {
			data = append(data, v)
		}
		fmt.Println("数据读取完成,", time.Since(startTime))
		data = quicksort.QuickSort(data)
		for _, v := range data {
			out <- v
		}
		close(out)
	}()
	return out
}

//合并
func Merge(in1, in2 <-chan int) chan int {
	out := make(chan int)

	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2 {
			if !ok2 || (ok1 && v1 <=v2) {
				out<-v1
				v1, ok1 = <-in1
			} else {
				out<-v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()

	return out
}

//读取数据
func ReadSource(reader io.Reader, chunckSize int) <-chan int {
	out := make(chan int, 2048)
	go func() {
		buf := make([]byte, 8)
		readSize := 0
		for {
			n, err := reader.Read(buf)
			if n > 0 {
				readSize += n
				out <- int(binary.BigEndian.Uint64(buf))
			}
			if err == io.EOF || (chunckSize != -1 && readSize >= chunckSize) {
				break
			}
		}
		close(out)
	}()

	return out
}

//写入
func WriteSync(writer io.Writer, in <-chan int) {
	for v := range in {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(v)) // 转成网络上传输的大端字节序的切片
		writer.Write(buf)
	}
}

//生成随机数数组
func RandomeSource(count int) <-chan int {
	out := make(chan int, 1024)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int() // 压入随机数
		}
		close(out)
	}()
	return out
}

//多路合并
func MergeN(ins ...<-chan int) <-chan int {
	if len(ins) == 1 {
		return ins[0]
	} else {
		m := len(ins) / 2
		return Merge(MergeN(ins[:m]...), MergeN(ins[m:]...))
	}
}

//
func ArraySource(nums...int) <-chan int{
	var out = make(chan int, 1024)

	go func() {
		for _, v := range nums{
			out<-v
		}
		close(out)
	}()

	return out
}
