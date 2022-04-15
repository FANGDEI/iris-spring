FROM openjdk:8
VOLUME ["/tmp","/log"]
COPY ./target/springboot-08-ssmp-0.0.1-SNAPSHOT.jar ssmp-0.0.1-SNAPSHOT.jar
ENTRYPOINT ["java","-jar","-Xmx400m","./ssmp-0.0.1-SNAPSHOT.jar","&"]
