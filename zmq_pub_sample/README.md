# build 
docker build -t zmq_pub_sample .

# run
docker run --rm -d -p 64230:64230 --name zmq_pub_sample zmq_pub_sample

# logging
docker logs -f zmq_pub_sample

# stop
docker stop zmq_pub_sample

# remove
docker rm zmq_pub_sample