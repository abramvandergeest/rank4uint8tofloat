package rank4uint8tofloat

import (

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
	var norm [][][][]float32
	for _,row0 := range indata{
		var outrow0 [][][]float32
		for _,row1 :=range row0{
			var outrow1 [][]float32
			for _,row2 := range row1{
				var outrow2 []float32
					for _,item:=range row2{
						normeditem:=(float32(item)-mean)/std
						outrow2=append(outrow2,normeditem)
					}
				outrow1=append(outrow1,outrow2)

			}
			outrow0=append(outrow0,outrow1)
		}
		norm=append(norm,outrow0)
	}

	// swapping x and y
	var output [][][][]float32
	for i:=0;i<len(norm);i++{
		var rowA [][][]float32
		for j:=0;j<len(norm[0][0]);j++{
			var rowB [][]float32
			for k:=0;k<len(norm[0]);k++{
				var rowC []float32
				for l:=0;l<len(norm[0][0][0]);l++{
					rowC=append(rowC,norm[i][k][j][l])
				}
				rowB=append(rowB,rowC)
			}
			rowA=append(rowA,rowB)
		}
		output=append(output,rowA)
	}

	return output, nil
}
