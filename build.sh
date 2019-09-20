#!/bin/sh
 
set -e
rm -rf bin

ROOTDIR=`pwd`
pack=pack
packfile=lkchain
tarfile=lkchain-linux-x64.tar.gz
packdst=pack/$packfile

function do_pack()
{
    if [ ! -f db_kv.tar ] ;then
        echo "not found db_kv.tar"
        echo "./scripts/download_init_db.sh"
        ./scripts/download_init_db.sh
    fi

    echo "pack"
    cd $ROOTDIR
    rm -rf $pack
    mkdir -p $packdst
    mkdir -p $packdst/bin
    mkdir -p $packdst/data
    mkdir -p $packdst/sbin
	mkdir -p $packdst/init

    chmod 777 bin/lkchain
    chmod 777 scripts/start.sh
    chmod 777 scripts/monitor.sh
    tar xf db_kv.tar -C $packdst/init
    chmod 744 -R pack/lkchain/init/db
    cp bin/lkchain $packdst/bin
    cp scripts/start.sh $packdst/sbin
    cp scripts/start-pre-testnet.sh $packdst/sbin
    cp scripts/monitor.sh $packdst/sbin
    cd pack ; tar zcf $tarfile $packfile ; echo "done $tarfile";
}


echo "start build ...."
# build chain
make build
do_pack
echo "build success!"
