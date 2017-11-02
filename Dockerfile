FROM scratch
MAINTAINER github.com/tydavis

COPY hello-world-docker /

CMD ["/hello-world-docker"]
