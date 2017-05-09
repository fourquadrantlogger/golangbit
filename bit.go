package golangbit

func GetBits( srcbyte  byte)(result []int){
	result=make([]int,8)
	for index:=uint(0);index<8;index++{
		mask:=byte( 1 <<index)
		b:=srcbyte&mask
		if(b==0){
			result[index]= 0
		}else {
			result[index]= 1
		}
	}
	return
}
func GetByte(srcbit []int)( result  byte){
	if(len(srcbit)!=8){
		panic("srcbit len must =8")
	}

	for index:=uint(0);index<8;index++{
		v:=srcbit[index]
		if(v==1||v==0){
			if(v==0){

			}else {
				result+=(1<<index)
			}
		}else {
			panic("bit in array must = 0 or 1")
		}
	}
	return
}
func SetByteBit(srcbyte  byte,index int,value int)(byte){
	if(index>=8 || index<0){
		panic("index error")
	}
	if(value!=0||value!=1){
		panic("value must =1 or 0")
	}
	bits:=GetBits(srcbyte)
	bits[index]=value

	return GetByte(bits)
}
func GetByteBit( srcbyte  byte,index byte)( int){
	bits:=GetBits(srcbyte)
	return bits[index]
}
