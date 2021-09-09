# build 
docker build -t zmq_sub_sample .

# run

# windows
docker run --rm -d -e PUB_ADDRESS=host.docker.internal --name zmq_sub_sample zmq_sub_sample

# linux
docker run --network host --rm -d -e PUB_ADDRESS=localhost --name zmq_sub_sample zmq_sub_sample

# logging
docker logs -f zmq_sub_sample

# stop
docker stop zmq_sub_sample

# remove
docker rm zmq_sub_sample