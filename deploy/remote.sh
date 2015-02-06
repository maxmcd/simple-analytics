pid=`pidof kayobe`
if [[ $pid ]]; then
    echo killing server, process id: $pid
    kill $pid
fi

source ~/.profile
echo removing old files
rm -rf $GOPATH/src/github.com
go get -u github.com/maxmcd/kayobe
echo installing from github
cd $GOPATH/src/github.com/maxmcd/kayobe/
echo starting server...
go run kayobe.go &