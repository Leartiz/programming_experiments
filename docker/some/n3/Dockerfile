# Базовый образ
FROM node

LABEL maintainer joshua

# Выполняя git внутри контейнера!
RUN git clone https://github.com/docker-in-practice/todo

WORKDIR /todo

# `/dev/null` — это специальный файл устройства 
# в Unix-подобных операционных системах,
# который действует как "черная дыра" для данных.
# Любые данные, записанные в `/dev/null`,
# отбрасываются и не сохраняются.

# Влияет на фаловую систему.
RUN npm install > /dev/null

EXPOSE 8000

CMD npm start