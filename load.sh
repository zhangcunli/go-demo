#!/bin/sh

status(){
    pgrep SearchServer
    if [ $? -ne 0 ] ; then
        echo "Service Server is closed" 
    else
        echo "Service Server is running..."
    fi
}

start(){
    pgrep SearchServer
    if [ $? -eq 0 ] ; then
        echo "Service Server is already running..." 
    else
        if [ ! -d log ];then
            mkdir log
        fi

        if [ ! -d "bin" ]; then 
            echo "dir bin is not exsit, SearchServer is not started, exit..." 
            exit 1
        fi
        cd bin
        mkdir -p ../status/goDemo
        export LD_LIBRARY_PATH=../lib
        #setsid ./supervise.demo -u ../status/goDemo ./goDemo </dev/null &>/dev/null &
        ./goDemo
        cd -
    fi
}

stop(){
    killall -9 supervise.SearchServer SearchServer
}

restart(){
    stop
    sleep 1
    start
}

case "$1" in
    'start')
        start
        ;;
    'stop')
        stop
        ;;
    'status')
        status
        ;;
    'restart')
        restart
        ;;
    *)
    echo "Usage: $0 {start|stop|restart|status}"
    exit 1
        ;;
    esac
	
