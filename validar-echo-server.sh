#!/bin/bash

server_ip=server
server_port=12345
network=tp0_testing_net

message="Hello world!"
response=$(docker run --rm --network $network alpine sh -c "echo $message | nc $server_ip $server_port" 2> /dev/null)

if [ "$response" = "$message" ]
then
	echo "action: test_echo_server | result: success"
else
	echo "action: test_echo_server | result: fail"
fi
