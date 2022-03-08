# ref:https://medium.com/@mccode/understanding-how-uid-and-gid-work-in-docker-containers-c37a01d01cf

FROM ubuntu:latest
RUN useradd -r -u 1000 appuser
USER appuser
ENTRYPOINT ["sleep", "infinity"]
