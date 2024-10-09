#!/bin/bash

# start server before client

../testing/start_server.sh > server.log &

sleep 5


../testing/start_client.sh > client.log
