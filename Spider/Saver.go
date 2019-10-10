package Spider

import (
	"SpiderHog/conf"
	"SpiderHog/utils"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

var SaverMachine *Saver
//var SLog *logrus.Logger = utils.GetStdoutLogger()
var SLog *logrus.Logger = utils.GetFileLogger()


func init() {
	Config := conf.GetConfigure()
	_ = Config.SavConfig // now is useless
	SaverMachine = &Saver{}
}

type Saver struct {
	SaveEngine SaveApi
}

type SaveApi interface {
	Save(interface{}) (error)
}

func (s *Saver)ActiveEngine(SaveEngine SaveApi) {
	s.SaveEngine = SaveEngine
}

func (s *Saver)Saved(group *sync.WaitGroup) {
	defer group.Done()
	for {
		select {
		case content := <-SaveChan:
			SLog.Println("[*]:  begin saving content...")
			s.SaveEngine.Save(content)
		case <-time.After(time.Duration(100) * time.Second):
			SLog.Println("[*]:  timeout, saver exit...")
			return
		}
	}
}

