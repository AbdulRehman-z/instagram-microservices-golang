FROM node:alpine

#working directory
WORKDIR /usr/app

#install dependencies
COPY package.json .
RUN npm install --omit=dev
COPY . .

#default command
CMD ["npm", "start"]