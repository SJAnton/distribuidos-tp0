#!/bin/bash
filename=$1
client_amount=$2

> "$filename"

echo "name: tp0" >> "$filename"
echo "services:" >> "$filename"

cat << EOF >> "$filename"
  server:
    container_name: server
    image: server:latest
    entrypoint: python3 /main.py
    environment:
      - PYTHONUNBUFFERED=1
    networks:
      - testing_net
    volumes:
      - ./server/config.ini:/config.ini
EOF

echo >> "$filename"

for i in $(seq 1 $client_amount)
do

cat << EOF >> "$filename"
  client$i:
    container_name: client$i
    image: client:latest
    entrypoint: /client
    environment:
      - CLI_ID=$i
      - NOMBRE="Nombre$i"
      - APELLIDO="Apellido$i"
      - DOCUMENTO=$((30000000 + $i))
      - NACIMIENTO=$((1950 + $i))-01-01
      - NUMERO=$((5000 + $i))
    networks:
      - testing_net
    depends_on:
      - server
    volumes:
      - ./client/config.yaml:/config.yaml
EOF

echo >> "$filename"
done

cat << EOF >> "$filename"
networks:
  testing_net:
    ipam:
      driver: default
      config:
        - subnet: 172.25.125.0/24
EOF
