package _kyu

/*
	base91编码解码

	编码字符
	ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~"

	1、编码原理
	分成13bit然后？？？
*/

var (
	data = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~\"")
)

func Encode(d []byte) []byte {
	// [0-90]需要至多7个bit来表示

	// t 116
	// 00111010 00000000

	// 25 1

	return nil // do it!
}

func Decode(d []byte) []byte {
	return nil // do it!
}
