FROM maven:3.9.4-eclipse-temurin-21

# Set the working directory
WORKDIR /app

# Copy the pom.xml and download dependencies
COPY pom.xml .
RUN mvn dependency:go-offline

# Copy the rest of your application
COPY . .

# Expose the application's port
EXPOSE 8080

# Run the application using Maven
CMD ["mvn", "spring-boot:run"]
