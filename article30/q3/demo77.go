package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// 示例1。
	reader1 := strings.NewReader(
		"NewReader returns a new Reader reading from s. " +
			"It is similar to bytes.NewBufferString but more efficient and read-only.")
	fmt.Printf("The size of reader: %d\n", reader1.Size())
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))

	buf1 := make([]byte, 47)
	n, _ := reader1.Read(buf1)
	fmt.Printf("读取字符串： %q\n", buf1)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len()))
	fmt.Printf("size: %d \t len: %d\n", reader1.Size(), reader1.Len())
	fmt.Println()

	// 示例2。 ReadAt方法不会依据已读计数进行读取，也不会在读取后更新它
	buf2 := make([]byte, 21)
	offset1 := int64(64)
	n, _ = reader1.ReadAt(buf2, offset1)
	fmt.Printf("从第64位开始读取: %q\n", buf2)
	fmt.Printf("%d bytes were read. (call ReadAt, offset: %d)\n", n, offset1)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len())) // 还是47，ReadAt方法不更新已读计数
	fmt.Println()

	// 示例3。
	offset2 := int64(17)
	expectedIndex := reader1.Size() - int64(reader1.Len()) + offset2                 //47 + 17 = 64
	fmt.Printf("Seek with offset %d and whence %d ...\n", offset2, io.SeekCurrent)   // 17,1
	readingIndex, _ := reader1.Seek(offset2, io.SeekCurrent)                         //重新计数 ,之前47 +17 =64
	fmt.Printf("The reading index in reader: %d (returned by Seek)\n", readingIndex) // 64
	fmt.Printf("The reading index in reader: %d (computed by me)\n", expectedIndex)  // 64

	n, _ = reader1.Read(buf2)
	fmt.Printf("移动到64开始读 %q \n", buf2)
	fmt.Printf("%d bytes were read. (call Read)\n", n)
	fmt.Printf("The reading index in reader: %d\n",
		reader1.Size()-int64(reader1.Len())) //119-(119-47-17-21) = 85
}
