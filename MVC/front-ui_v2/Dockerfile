# Этап 1: Сборка приложения
FROM node:21 AS builder
#FROM node:21

# Установите рабочую директорию в контейнере
WORKDIR /app

# Скопируйте файлы `package.json` и `package-lock.json` для установки зависимостей
COPY package*.json ./

# Установите зависимости
RUN npm install

# Скопируйте исходный код приложения в контейнер
COPY . .

# Соберите приложение
RUN npm run build
#CMD ["npm", "start"]
# Этап 2: Запуск с помощью Nginx
FROM nginx:stable-alpine

# Скопируйте собранное приложение из предыдущего шага
COPY --from=builder /app/dist /usr/share/nginx/html

# Откройте порт 80 для доступа к приложению
#EXPOSE 80

# Запустите Nginx и отслуживайте содержимое
CMD ["nginx", "-g", "daemon off;"]

