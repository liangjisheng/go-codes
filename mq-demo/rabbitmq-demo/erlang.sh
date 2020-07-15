#!/bin/bash

curl -fsSL https://github.com/rabbitmq/signing-keys/releases/download/2.0/rabbitmq-release-signing-key.asc | sudo apt-key add -

# sudo apt-key adv --keyserver "hkps://keys.openpgp.org" --recv-keys "0x0A9AF2115F4687BD29803A206B73A36E6026DFCA"

sudo apt-get install apt-transport-https

vim /etc/apt/sources.list.d/bintray.erlang.list

# deb http://dl.bintray.com/rabbitmq-erlang/debian $distribution $component
# ubuntu16.04
deb http://dl.bintray.com/rabbitmq-erlang/debian xenial erlang
# ubuntu18.04
# deb http://dl.bintray.com/rabbitmq-erlang/debian bionic erlang

sudo apt-get update -y

sudo apt-get install -y erlang-base \
                        erlang-asn1 erlang-crypto erlang-eldap erlang-ftp erlang-inets \
                        erlang-mnesia erlang-os-mon erlang-parsetools erlang-public-key \
                        erlang-runtime-tools erlang-snmp erlang-ssl \
                        erlang-syntax-tools erlang-tftp erlang-tools erlang-xmerl
