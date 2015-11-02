package main

import (
	"bufio"
	gcfg "code.google.com/p/gcfg"
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
	logger "github.com/shengkehua/xlog4go"
	"os"
	"strconv"
	"strings"
	"time"
	"userinfo"
)

const PID_MASK = 0xffffffffffff

func main() {
	// init log
	if err := logger.SetupLogWithConf(logFile); err != nil {
		fmt.Printf("log init fail: %s\n", err.Error())
		return
	}
	defer logger.Close()

	//read config file
	if err := gcfg.ReadFileInto(&cfg, confFile); err != nil {
		logger.Error("read conf failed, err:%s", err.Error())
		return
	}
	logger.Info("logFile=%v, confFile=%v, cfg=%v", logFile, confFile, cfg)

	//read data file
	path := "../data/hadoop.txt"
	logger.Debug("read file path=%v", path)
	file, err := os.Open(path)
	if err != nil {
		logger.Error("open file:%s fail, err=%v", path, err)
	}
	defer file.Close()

	//redis
	st := time.Now()
	conn, err := redis.DialTimeout("tcp", "10.10.16.99:6380", 0, 1*time.Second, 1*time.Second)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.Split(scanner.Text(), "\t")
		addr, err := parseLine(str)
		if err == nil {
			//logger.Debug("add=%v", addr)
			buffer, _ := proto.Marshal(addr)
			key := "goaddr" + str[0]
			conn.Do("SET", key, buffer)
		}
	}
	fmt.Println("consume time:", time.Since(st))

	if err := scanner.Err(); err != nil {
		logger.Error("read file:%s error, err=%v", path, err)
	}
}

func parseLine(str []string) (*userinfo.UserAddr, error) {
	//fmt.Println(str[0], str[1], str[2])
	addr := &userinfo.UserAddr{}
	if 3 != len(str) || "NULL" == str[1] || "NULL" == str[2] {
		return addr, errors.New("Error line!")
	}

	start := strings.Split(str[1], "#")
	end := strings.Split(str[2], "#")
	//fmt.Println(start[0], start[1], end[0], end[1])
	pid, _ := strconv.ParseInt(str[0], 10, 64)
	addr.Userid = proto.Int64(pid)
	slng, _ := strconv.ParseFloat(start[0], 64)
	addr.Homelng = proto.Float64(slng)
	slat, _ := strconv.ParseFloat(start[1], 64)
	addr.Homelat = proto.Float64(slat)
	elng, _ := strconv.ParseFloat(end[0], 64)
	addr.Corplng = proto.Float64(elng)
	elat, _ := strconv.ParseFloat(end[1], 64)
	addr.Corplat = proto.Float64(elat)
	return addr, nil
}

/*
    str1 := "test1"
    key1 := "test" + "1"
    str2 := "test2"
    key2 := "test" + "2"
    conn.Do("MSET", key1, str1, key2, str2)

   //buffer, _ := proto.Marshal(homeAddr1)
   //key := "goaddr" + strconv.FormatInt(passendid, 10)
   //_, err = conn.Do("SET", key, buffer)

   //home2 := &userinfo.UserAddr{}
   //buf1, _ := redis.Bytes(conn.Do("GET", key))
   //_ = proto.Unmarshal(buf1, home2)
   //fmt.Printf("buf1=%v\nhome2=%v\nerr=%v\n", buf1, home2, err)

   ////////////
   home3 := &userinfo.UserAddr{}
   key1 := "addr111066055" //"addr350863892480"
   buf2, _ := redis.Bytes(conn.Do("GET", key1))
   _ = proto.Unmarshal(buf2, home3)
   fmt.Printf("home2=%v\n", home3)
*/
