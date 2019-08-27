package spider

import (
	"../conf"
	"../utils"
	"bytes"
	"io"
	"log"
	"os"
	"sync"
	"time"
)

type Saver struct {
	SavePath string
	SaveMethod string
}

type Param interface {

}

type SaveInterface interface {
	InitSaver(j conf.JsonConfig)
	SaveRule(content string)
	SaverInit()
	BinSave(fileName string, BContent []byte)
	StrSave(fileName string, SContent string)
}

func (s *Saver)InitSaver(j conf.JsonConfig)  {
	s.SaveMethod = j.SConfig.SaveMethod
	s.SavePath = j.SConfig.SavePath
}

func Save (group *sync.WaitGroup, saveInterface SaveInterface) {
	saveInterface.SaverInit()
	defer group.Done()
	for {

		select {
		case saveContent := <-saveChan:
			saveInterface.SaveRule(saveContent)
		case <-time.After(time.Duration(10) * time.Second):
			log.Println("[*]:  timeout, saver exit...")
			return
		}

	}
}

func (s *Saver)SaveRule(content string)  {

}

func (s *Saver)SaverInit() {
	if ! utils.Exists(s.SavePath){
		e := os.Mkdir(s.SavePath, 0755)
		if e != nil{
			log.Fatal("[*]:  保存文件夹创建失败")
		} else {
			log.Println("[*]:  文件夹创建成功")
		}
	}
}

func (s *Saver)BinSave(fileName string, BContent []byte) {
	if utils.Exists(s.SavePath+fileName){
		log.Fatal("[*]:  "+s.SavePath+fileName+"已存在")
	} else {
		f, err := os.Create(s.SavePath+fileName)
		if utils.Check(err) {
			_, err := io.Copy(f, bytes.NewReader(BContent))
			if ! utils.Check(err) {
				log.Fatal("[*]:  保存失败")
			}
		}
	}
}

func (s *Saver)StrSave(fileName string, SContent string)  {
	if utils.Exists(s.SavePath+fileName){
		log.Fatal("[*]:  "+s.SavePath+fileName+"已存在")
	} else {
		f, err := os.Create(s.SavePath + fileName)
		if utils.Check(err) {
			_, err := io.WriteString(f, SContent)
			if ! utils.Check(err) {
				log.Fatal("[*]:  保存失败")
			}
		}
	}
}