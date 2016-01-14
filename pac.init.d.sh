#!/bin/sh
#
# chkconfig: 2345 55 25
# Description: pac init.d script, put in /etc/init.d, chmod +x /etc/init.d/pac
#              For Debian, run: update-rc.d -f pac defaults
#              For CentOS, run: chkconfig --add pac
#
### BEGIN INIT INFO
# Provides:          pac
# Required-Start:    $all
# Required-Stop:     $all
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: pac init.d script
# Description:       OpenResty (aka. ngx_openresty) is a full-fledged web application server by bundling the standard pac core, lots of 3rd-party pac modules, as well as most of their external dependencies.
### END INIT INFO
#

DESC="Pac Daemon"
NAME=pac
PREFIX=root/pac/
DAEMON=/$PREFIX/bin/$NAME
CONF=/$PREFIX/etc/config.json
LOG=/$PREFIX/logs/
PID=/$PREFIX/logs/$NAME.pid
SCRIPT=/etc/init.d/$NAME

if [ ! -x "$DAEMON" ] || [ ! -f "$CONF" ]; then
    echo -e "\033[33m $DAEMON has no permission to run. \033[0m"
    echo -e "\033[33m Or $CONF doesn't exist. \033[0m"
    sleep 1
    exit 1
fi

do_start() {
    if [ -f $PID ]; then
        echo -e "\033[33m $PID already exists. \033[0m"
        #echo -e "\033[33m $DESC is already running or crashed. \033[0m"
        #echo -e "\033[32m $DESC Reopening $CONF ... \033[0m"
        #$DAEMON -s reopen -c $CONF
        #sleep 1
        #echo -e "\033[36m $DESC reopened. \033[0m"
    else
        echo -e "\033[32m $DESC Starting $CONF ... \033[0m"
        GIN_MODE=release $DAEMON -prefix $PREFIX -log_dir $LOG &
        sleep 1
        pgrep $NAME > $PID
        echo -e "\033[36m $DESC started. \033[0m"
    fi
}

do_stop() {
    if [ ! -f $PID ]; then
        echo -e "\033[33m $PID doesn't exist. \033[0m"
        echo -e "\033[33m $DESC isn't running. \033[0m"
    else
        echo -e "\033[32m $DESC Stopping $CONF ... \033[0m"
        pkill $NAME
        rm $PID
        sleep 1
        echo -e "\033[36m $DESC stopped. \033[0m"
    fi
}



do_info() {
    $DAEMON -h
}

case "$1" in
 start)
 do_start
 ;;
 stop)
 do_stop
 ;;
 restart)
 do_stop
 do_start
 ;;
 info)
 do_info
 ;;
 *)
 echo "Usage: $SCRIPT {start|stop|restart|info}"
 exit 2
 ;;
esac

exit 0

