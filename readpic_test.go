package excelize

import (
	"fmt"

	"io/ioutil"
	"log"
	"testing"
)

func TestPic(t *testing.T) {
	// 打开文件
	f, err := OpenFile("./Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
		return
	}

	fmt.Println(f.GetSheetMap()[1])
	for index, name := range f.GetSheetMap() {
		log.Println(index, name)
		// file, _, err := f.GetPictureByRowCol(name, 41, 0)
		// log.Println(file, err)
		allPics := f.GetAllPictures(name)
		log.Println(allPics)
		for _, v := range allPics {
			err := ioutil.WriteFile(v.Filename, v.Bytes, 0644)
			if err != nil {
				fmt.Println(err)
			}
			break
		}
		// // // Get value from cell by given worksheet name and axis.
		// // cell, err := f.GetCellValue("整体订货单", "B2")
		// // if err != nil {
		// // 	fmt.Println(err)
		// // 	return
		// // }
		// // fmt.Println(cell)
		// // Get all the rows in the Sheet1.
		// rows, _ := f.GetRows(name)
		// for _, row := range rows {
		// 	for _, colCell := range row {
		// 		fmt.Print(colCell, "\t")
		// 	}
		// 	fmt.Println()
		// }
		// // file, raw, err := f.GetPicture(f.GetSheetMap()[1], "D10")
		// file, raw, err := f.GetPictureByRowCol(f.GetSheetMap()[1], 45, 1)
		// if err != nil {
		// 	fmt.Println("GetPicture Error:", err)
		// 	t.FailNow()
		// }
		// fmt.Println(file, raw, err)
		// err = ioutil.WriteFile(file, raw, 0644)
		// if err != nil {
		// 	fmt.Println(err)
		// }
	}

}
