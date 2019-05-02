package ioT

import (
	"fmt"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	// destFile, err := os.Create("temp.txt")
	// if err != nil {
	// 	fmt.Println("Open File Error:", err.Error())
	// }
	// destFile.WriteString("2018-02-28 os.Create")

	opFile, err := os.OpenFile("temp.txt", os.O_RDWR, 666)
	/*os.Open -> readonly */
	if err != nil {
		fmt.Println("Open File Error:", err.Error())
	}
	opFile.WriteString("\n tik tik tik")
	opFile.WriteString("\n tok tok tok")
	t.Log("End")
}
