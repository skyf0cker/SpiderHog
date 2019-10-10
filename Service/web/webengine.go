package web

import (
	"SpiderHog/Service/web/router"
	"log"
	"net/http"
	"sync"
)

func ServiceActive(group *sync.WaitGroup) {
	defer group.Done()
	router := router.InitRouter()

	s := &http.Server{
		//Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Addr:           ":3000",
		Handler:        router,
		//ReadTimeout:    setting.ServerSetting.ReadTimeout,
		//WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("gin Started...")

	if err := s.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
