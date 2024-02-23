# GoAuthJWTService
тестовое задание на junior go developer

# Доступные эндпоинты

1) GET   /auth/token — endpoint для генерации пары токенов

2) POST  /auth/token/refresh — endpoint для проверки refresh_token и выдачи новых пар токенов

3) GET   /swagger/* - Swagger (в сваггере лучше обьясняется все возможные выводы и необходимые параметры для проекта)

### Swagger UI:

http://localhost:8080/docs/index.html


#### Запуск локально у себя (лучше make run):
    make local
    
#### Запуск в докере(легче в использовании и меньше мороки):
    make run
