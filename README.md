# golangbit
golang bit operation
> golang 的字节 位操作

+ GetBits(b byte)([]int)

    返回一个长度为8的0/1数组
+ GetByte([]int)(byte)

    将一个长度为8的0/1数组，转成byte

+ SetByteBit(srcbyte  byte,index int,value int)(byte)

    设置一个byte的index位的值为value(0/1)，并返回该byte
+ GetByteBit(srcbyte  byte,index int)(int)

    返回一个byte的index位的值(0/1)
