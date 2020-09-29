package distributesort

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func CreateNetworkPipe(fileName string, fileSize int, chunkCount int) <-chan int {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pipe := RandomeSource(fileSize / 8)
	writer := bufio.NewWriter(file)
	WriteSync(writer, pipe)
	writer.Flush()

	file1, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file1.Close()

	chunkSize := fileSize / chunkCount
	var sortAddr []string
	Init()

	for i := 0; i < chunkSize; i++ {
		file1.Seek(int64(i*chunkSize), 0)
		source := ReadSource(bufio.NewReader(file1), chunkSize)
		addr := ":"+strconv.Itoa(7000+i) //开辟地址
		NetworkWrite(addr, InMemorySort(source))//写入到分布式主机
		sortAddr = append(sortAddr, addr)
	}

	var sortResults []<-chan int
	for _, addr:= range sortAddr{
		sortResults = append(sortResults, NetworkRead(addr))
	}

	return MergeN(sortResults...)
}

//1.本地 2.多线程 3.分布式

//多线程-调用中间件完成
func CreatePipe(fileName string, fileSize int, chunkCount int) <-chan int {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pipe := RandomeSource(fileSize / 8)
	writer := bufio.NewWriter(file)
	WriteSync(writer, pipe)
	writer.Flush()

	chunkSize := fileSize / chunkCount
	var sortResults []<-chan int // 排序结果
	Init()

	file, err = os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for i := 0; i < chunkCount; i++ {
		file.Seek(int64(i*chunkSize), 0) // 跳转文件指针
		source := ReadSource(bufio.NewReader(file), chunkSize)
		sortResults = append(sortResults, InMemorySort(source))
	}
	return MergeN(sortResults...)
}

func WriteToFile(in <-chan int, fileName string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	WriteSync(writer, in)
}

func ShowFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	in := ReadSource(bufio.NewReader(file), -1)
	count := 0
	for v := range in {
		fmt.Println(v)
		count++
		if count > 1000 {
			break
		}
	}
}
