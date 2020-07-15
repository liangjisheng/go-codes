FROM python:3.6
ADD . /docker
# ADD ../../src /src
WORKDIR /src
RUN pip install flask redis
CMD ["python", "app.py"]