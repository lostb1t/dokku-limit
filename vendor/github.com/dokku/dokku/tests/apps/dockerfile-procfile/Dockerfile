FROM node:4

EXPOSE 3001/udp
EXPOSE  3000/tcp
EXPOSE 3003

COPY . /app
WORKDIR /app
RUN npm install

CMD npm start
