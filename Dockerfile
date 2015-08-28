FROM ubuntu:14.04
MAINTAINER Boyd Hemphill <boyd@stackengine.com>

# Install the old and faithful stress package
RUN apt-get update && apt-get install --assume-yes stress

# Now use a go binary to wrap stress so we can launch an arbitrary number of
# containers on a node.
WORKDIR /code
ADD ./buildoutput/docker-stress stress-full
RUN chmod +x stress-full

ENTRYPOINT ["/code/stress-full"]