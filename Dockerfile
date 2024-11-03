# Stage 1: Build the application
FROM maven:3.9.4-eclipse-temurin-21 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Maven project file
COPY pom.xml .

# Download project dependencies
RUN mvn dependency:go-offline

# Copy the rest of the project files
COPY src ./src

# Build the WAR file
RUN mvn clean package -DskipTests

# Stage 2: Run the application
FROM tomcat:10.1.14-jdk21-temurin

# Remove the default Tomcat web apps
RUN rm -rf /usr/local/tomcat/webapps/*

# Copy the WAR file from the build stage
COPY --from=build /app/target/*.war /usr/local/tomcat/webapps/ROOT.war

# Expose port 8080
EXPOSE 8080

# Start Tomcat server
CMD ["catalina.sh", "run"]
