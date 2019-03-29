FROM gradle:latest as builder

COPY --chown=gradle:gradle . /home/gradle/src
WORKDIR /home/gradle/src
RUN gradle build

FROM openjdk:11-jre-slim

COPY --from=builder /home/gradle/src/build/libs/demo-0.0.1-SNAPSHOT.jar /app/
WORKDIR /app

CMD ["java", "-jar", "./demo-0.0.1-SNAPSHOT.jar"]
