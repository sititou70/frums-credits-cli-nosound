FROM troian/golang-cross:v1.16.3

RUN apt update && apt install -y libasound2-dev
