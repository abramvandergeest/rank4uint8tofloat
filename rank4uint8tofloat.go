package rank4uint8tofloat

import (

	// "time"

	"fmt"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

func init() {
	function.Register(&fnRank4uint8tofloat{})
}

type fnRank4uint8tofloat struct {
}

func (fnRank4uint8tofloat) Name() string {
	return "rank4uint8tofloat"
}

func (fnRank4uint8tofloat) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeArray,data.TypeFloat32,data.TypeFloat32}, false
}

func (fnRank4uint8tofloat) Eval(params ...interface{}) (interface{}, error) {
	// Converting Rank 4 tensor of uint8 to Rank 4 tensor of float32
	indata := params[0].([][][][]uint8)
	mean:=params[1].(float32)
	std:=params[2].(float32)
	fmt.Println("BLAH")
	var output [][][][]float32
	for _,row0 := range indata{
		var outrow0 [][][]float32
		for _,row1 :=range row0{
			var outrow1 [][]float32
			for _,row2 := range row1{
				var outrow2 []float32
					for i,item:=range row2{
						normeditem:=(float32(item)-mean)/std
						outrow2=append(outrow2,normeditem)
						if i==0{
							fmt.Println(normeditem)
						}
					}
				outrow1=append(outrow1,outrow2)

			}
			outrow0=append(outrow0,outrow1)
		}
		output=append(output,outrow0)
	}

	return output, nil
}
