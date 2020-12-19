Для повышения своих навыком в мире golang мне нужен был проект. Независимый, где я могу практиковать библиотеки, фреймворки или менять всё архитектуру проекта. Для этого я написал свой сервис авторизаций с чатом. Долгое время я работал фронтенде. Многие термины мне новинку в мире golang. К примеру тип логов, goconcurrency, goroutine. Но узнавая это я чувствую что двигаюсь, вид самое главное понимание того что пишешь.
Здесь собраны те стеке всегда интересовали меня, но не применял на практики. На конец-то нашёл время уделить на него. Познакомится ближе с swagger для документаций API.
Я знаю что нужно учить паттерны и придерживаться им. Когда у тебя мало практического опыта и в начальном этапе мало чем поможет. Да не спорю что повысить производительность проекта. Но мне кажется, если не столкнешься не поймешь.

Что частности в двигает такие проекты:
Многих собесах спрашиваю вот вы написали знаете golang, Есть ли проекты?
вот это написал сам читая книги, много статьи и смотря видео.
И готов защищать этот проект.



Первая этап:
- :heavy_check_mark: Проект должен придерживаться структуры project-layouts
- :heavy_check_mark: Проект должен запускаться через docker-compose
- :heavy_check_mark: Проект должен легко подключатся на любое бд в моём случай mongodb, через докер образ или внешний адресах
- :heavy_check_mark: Проект должен подключатся на другие технологий  redis, rabbitmq, через докер образ или внешний адресах
- :heavy_check_mark: Проекте должна использоваться swagger

Вторая задача:
- :heavy_check_mark: У проекта должна быть API авторизация 
- :heavy_check_mark: У проекта должна быть API регистрация
- :heavy_check_mark: У проекта должна быть API рефреш
- :heavy_check_mark: У проекта должна быть приватные API через JWT
- Тесты на все апи и бд


Третья задача:
- :heavy_check_mark: С помощью RMQ отправить команды
Это уже будет выполнять другой микросервис 
- Отправка на почту временный код активация при регистраций
- Смена пароля отправить временный код для возможности сменить пароль


Четвертая задача:
Предыстория: React тоже хотелось основательно собрать свой pack. Детально рассмотреть возможности webpack. Подключить Redux c typescript. Самом деле думаю typescript в React она тормозит проект виде написание. В Ангуляре хоть идёт по умолчанию но возможности этого языка в мире ООП раскрывается в этом фрейворке. 
- :heavy_check_mark: Поднять  React проект с нуля со своими библиотеками
- :heavy_check_mark: React, Redux подключаем API сервиса
- :heavy_check_mark: выносим в nginx

Пятая задача:
- Онлайн чат через websocket
- У проекта должно быть возможность создавать чат (то есть hub)
- Возможность подключаться в чат 
- Возможность ред. название чата
- Возможность удалять чат

Полезные ссылки:
https://medium.com/@theShiva5/creating-simple-login-api-using-go-and-mongodb-9b3c1c775d2f
https://www.nexmo.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr
https://www.digitalocean.com/community/tutorials/how-to-use-go-with-mongodb-using-the-mongodb-go-driver-ru
https://www.soberkoder.com/swagger-go-api-swaggo/
https://www.mongodb.com/blog/post/quick-start-golang--mongodb--how-to-create-documents
https://medium.com/better-programming/unit-testing-code-using-the-mongo-go-driver-in-golang-7166d1aa72c0
https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
https://github.com/b-pagis/medium-mongo-go-driver/blob/master/databases/database_test.go
Бесплатный смс https://www.twilio.com/
https://stackoverflow.com/questions/43631854/gracefully-shutdown-gorilla-server
https://dev.to/jacobsngoodwin/full-stack-memory-app-01-setup-go-server-with-reload-in-docker-62n



Требование проекту:
- Максимально упрощать код и понимания архитектуру самого проекта


mongodb:
```
db.createUser({
    user: 'admin',
    pwd: '*****',
    roles: [{ role: 'readWrite', db:'exclusive'}]
})
```


rabbitmq:
```
$ sudo rabbitmqctl add_user admin sample
Adding user "admin" ...

$ sudo rabbitmqctl set_user_tags admin administrator
Setting tags for user "cc-admin" to [administrator] ...

sudo rabbitmqctl change_password guest guest123
```
