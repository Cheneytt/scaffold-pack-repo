#!/bin/sh

set -e

version="$(cat plugin.yaml | grep "version" | cut -d '"' -f 2)"
echo "Installing scaffold-pack-repo ${version} ..."

# Find correct archive name
unameOut="$(uname -s)"

case "${unameOut}" in
    Linux*)     os=Linux;;
    Darwin*)    os=Darwin;;
    CYGWIN*)    os=Cygwin;;
    MINGW*)     os=windows;;
    *)          os="UNKNOWN:${unameOut}"
esac

arch=`uname -m`
url="https://github.com/Cheneytt/scaffold-pack-repo/releases/download/${version}/scaffold-pack-repo_${version}_${os}_${arch}.tar.gz"

if [ "$url" = "" ]
then
    echo "Unsupported OS / architecture: ${os}_${arch}"
    exit 1
fi

echo "Downloading archive from $url"

filename=`echo ${url} | sed -e "s/^.*\///g"`

# Download archive
if [ -n $(command -v wget) ]
then
    wget $url
elif [ -n $(command -v curl) ]
then
    curl -SL -O $url
else
    echo "Need wget or curl"
    exit -1
fi

# Install bin
rm -rf bin && tar xzf $filename > /dev/null && rm -f $filename

echo "Done."