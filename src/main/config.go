package main

/*
   配置文件的结构体
*/

type Config struct {
	Demo struct {
		GoroutineCount int
		Port           string
	}
	Monitor struct {
		PrintFile     string
		PrintInterval int
	}
	Redis struct {
		Server           []string
		ConnectTimeoutMs int
		WriteTimeoutMs   int
		ReadTimeoutMs    int
		MaxIdle          int
		MaxActive        int
		IdleTimeoutS     int
		Password         string
		KeyPrefix        string
	}

	Memcache struct {
		Server           []string
		TimeoutMs        int
		EnableReadCache  bool
		EnableWriteCache bool
		KeyPrefix        string
	}
	Service struct {
		Switch   bool
		Interval uint32
		Host     []string
		Api      string
	}
	Watchdog map[string]*struct {
		Switch      bool
		Fluctuation float64
		Min         uint64
		Max         uint64
		Msg         string
	}
}
