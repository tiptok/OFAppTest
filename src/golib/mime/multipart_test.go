package mime

import (
	"bytes"
	"log"
	"mime/multipart"
	"testing"
)

func Test_multipart(t *testing.T){
	var bufReader bytes.Buffer
	mpWriter:=multipart.NewWriter(&bufReader)
	fw,err :=mpWriter.CreateFormFile("file_a","a.txt")
	if err!=nil{
		t.Fatal(err)
		return
	}
	fw.Write([]byte("this is a txt"))
	mpWriter.WriteField("name","tiprok")


	fwb,err :=mpWriter.CreateFormFile("file_a","b.txt")
	if err!=nil{
		t.Fatal(err)
		return
	}
	fwb.Write([]byte("this is b txt"))
	mpWriter.Close()



	//t.Log("write ",mpWriter.FormDataContentType())
	t.Log("write",bufReader.String())

	mpReader:=multipart.NewReader(&bufReader,mpWriter.Boundary())
	form,err :=mpReader.ReadForm(100)
	if err!=nil{
		t.Fatal(err)
	}
	for _,files :=range form.File{
		for i:=range files{
			file :=files[i]
			if f,err :=file.Open();err!=nil{
				return
			}else{
				bufTmp :=bytes.NewBuffer(nil)
				bufTmp.ReadFrom(f)
				log.Println(file.Filename," file:",bufTmp.String())
			}
		}
	}
	//log.Println("next part.")
	//for{
	//	if p,err :=mpReader.NextPart();(p!=nil && err==nil){
	//		log.Println(p.FileName())
	//	}else{
	//		break
	//	}
	//}
}