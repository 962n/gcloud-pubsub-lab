FROM gcr.io/google.com/cloudsdktool/cloud-sdk:424.0.0-emulators
RUN apt-get update && \
    apt-get install -y netcat