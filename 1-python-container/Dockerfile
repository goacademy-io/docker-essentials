FROM bitnami/python:3
LABEL maintainer="info@goacademy.io"
WORKDIR goacademy
COPY . .
# The EXPOSE instruction informs Docker that the container listens on the specified network ports at runtime
# It doesn't open the container port, it just like documenting the port
EXPOSE 8000
CMD ["python","server.py"]