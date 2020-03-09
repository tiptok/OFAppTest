package main

import (
	"github.com/disintegration/imaging"
	image2 "image"
	//"image/png"
	"log"
	"path/filepath"
)

func main() {
	var (
		path = "E:\\个人\\照片\\1576414574(1).png"
		w    = 100
		h    = 120
	)
	if err := resize(path, w, h); err != nil {
		log.Println(err)
	}
}

func resize(path string, w, h int) (err error) {
	var (
		image image2.Image
		//newfile *os.File
	)
	if image, err = imaging.Open(path); err != nil {
		return
	}
	//thumb:=imaging.Thumbnail(image,w,h,imaging.Lanczos)
	//dst :=imaging.New(w,h,color.NRGBA{0, 0, 0, 0})
	//dst =imaging.Paste(dst,thumb,image2.Pt(0,0))
	//dir,file:=filepath.Split(path)
	//if newfile,err =os.Create(filepath.Join(dir,"thumb"+file));err!=nil{
	//	return
	//}
	//if err= jpeg.Encode(newfile,dst,&jpeg.Options{Quality:50});err!=nil{
	//	return
	//}

	dir, file := filepath.Split(path)
	//if newfile,err =os.Create(filepath.Join(dir,"thumb"+file));err!=nil{
	//	return
	//}
	image = imaging.Resize(image, w, h, imaging.Lanczos)
	if err = imaging.Save(image, filepath.Join(dir, "thumb_"+file)); err != nil {
		return
	}
	return
}
