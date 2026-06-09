FROM eclipse-temurin:17-jdk-jammy

WORKDIR /tmp
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && \
  apt-get -y upgrade && \
  apt-get install -y gcc g++ make sqlite3 default-mysql-client && \
  apt-get clean

RUN useradd --uid=1001 --create-home isucon
USER isucon

RUN mkdir -p /home/isucon/webapp/java
WORKDIR /home/isucon/webapp/java
COPY --chown=isucon:isucon ./ /home/isucon/webapp/java/

CMD ["./gradlew", "bootRun"]
