# Project_template

Это шаблон для решения проектной работы. Структура этого файла повторяет структуру заданий. Заполняйте его по мере работы над решением.

# Задание 1. Анализ и планирование

<aside>

Чтобы составить документ с описанием текущей архитектуры приложения, можно часть информации взять из описания компании и условия задания. Это нормально.

</aside

### 1. Описание функциональности монолитного приложения


**Управление отоплением:**

- Пользователи могут включать и выключать отопление в доме удаленно
- Система поддерживает передачу команд управления от сервера к установленному в доме датчику
- Система разработана в разрезе монолитной архитектуры 
- Система поддерживает только синхронные вызовы
- Система не поддерживает возможность самостоятельного подключения устройств, только через услуги специалиста
- Система поддерживает только запрос данных о температуре со стороны сервера к установленному в доме датчику

**Мониторинг температуры:**

- Пользователи могут смотреть температуру в доме 
- Система поддерживает получение значений температуры путём взаимодействия сервера с датчиком 
- Система поддерживает получение значений температуры только по запросу от сервера 
- Система хранит состояние температуры в БД (Postgres)

### 2. Анализ архитектуры монолитного приложения

1. Монолитное приложение разработано на Go с использованием базы данных PostgreSQL
2. Взаимодействие между монолитом и внешними системами синхронное 
3. Инициатором обмена данными между приложением и внешними системами выступает сервер приложения


### 3. Определение доменов и границы контекстов

1. Домен управления отоплением в домах 
Работает синхронно по запросу со стороны сервера 

2. Домен мониторинга температуры в домах
Работает синхронно по запросу со стороны сервера и сохраняет данные в базу данных 

3. Домен сервисного обслуживания (не уверен насколько он нужен, но если поступил запрос и в монолите на этот счёт заложена бизнес-логика, то возможно есть смысл выделить его)

### **4. Проблемы монолитного решения**

- Целевая экосистема требует расширения функциональности, соответственно текущий монолит придется масштабировать, что проблематично
- Очень медленное синхронное взаимодействие между монолитным приложением и домами, после масштабирования
- Отсутствие очередей, что не дает возможность переноса определенных процессов в фоновый режим. Соостветственно растет нагрузка на систему.
- Подключение новый устроиств, с учётом монолитной архитектуры потребует значительноё ее доработки
- Команда описанная в задании слишком маленькая для работы над крупным монолитным приложением 
- Необходимость полного ретеста всей функциональности монолита, после внесения изменений, так как изменения могут повлиять на другу функциональность
- В случае сбоя в одной части монолитного-приложения, может повредить другие его части.
- В случае если измениться команда разработчиков/тестирования/аналитики и т.д, будет проблематично передавать экспертную оценку.

### 5. Визуализация контекста системы — диаграмма С4

Добавьте сюда диаграмму контекста в модели C4.

Чтобы добавить ссылку в файл Readme.md, нужно использовать синтаксис Markdown. Это делают так:

```markdown
http://127.0.0.1:8000/architecture/context/
```

# Задание 2. Проектирование микросервисной архитектуры

В этом задании вам нужно предоставить только диаграммы в модели C4. Мы не просим вас отдельно описывать получившиеся микросервисы и то, как вы определили взаимодействия между компонентами To-Be системы. Если вы правильно подготовите диаграммы C4, они и так это покажут.

**Диаграмма контейнеров (Containers)**

http://127.0.0.1:8000/architecture/container/

**Диаграмма компонентов (Components)**

1. User&House Service - http://127.0.0.1:8000/architecture/User%26HouseService/
2. Device Service - http://127.0.0.1:8000/architecture/DeviceService/
3. Heating Service - http://127.0.0.1:8000/architecture/HeatingService/
4. Lightning Service - http://127.0.0.1:8000/architecture/LightningService/
5. Surveillance Service - http://127.0.0.1:8000/architecture/SurveillanceService/
6. Notification Service - http://127.0.0.1:8000/architecture/NotificationService/

**Диаграмма кода (Code)**

Насколько я понимаю, диаграмма кода не является обязательной в разрезе C4 нотации 
И обычно делают диаграммы кода для самых сложных компонентов, у меня это: 
1. Компонент управления пользователями и домами - http://127.0.0.1:8000/architecture/User%26HouseServiceCode/
2. Компонент управления устройствами - http://127.0.0.1:8000/architecture/DeviceServiceCode/

Остальные компонентны идентичны друг другу и попроще, возможно не совсем целесообразно их отражать
Если я не прав, просьба отметить это на ревью - доработаю все что нужно

FYI: Не совсем сейчас понимаю, как сделать сервер статичным когда ревьюер будет проверять задание. Потому что если я закрою терминал, то сервер упадет и будет недоступен. Тут просьба подсказать, как сделать.

# Задание 3. Разработка ER-диаграммы

ER диаграммы были спроектированы под каждую БД микросервисов 

1. User&House Service - http://127.0.0.1:8000/architecture/ERD/User%26HouseServiceERD/
2. Device Service - http://127.0.0.1:8000/architecture/ERD/DeviceServiceERD/
3. Heating Service - http://127.0.0.1:8000/architecture/ERD/HeatingServiceERD/
4. Lightning Service - http://127.0.0.1:8000/architecture/ERD/LightningServiceERD/
5. Surveillance Service - http://127.0.0.1:8000/architecture/ERD/SurveillanceService/
6. Notification Service - http://127.0.0.1:8000/architecture/ERD/NotificationServiceERD/

# Задание 4. Создание и документирование API

### 1. Тип API
1. Решил сделать единый эндпоинт через REST, который будут дергать потребители "Теплого дома" 
Под капотом уже идет разделение и маршрутизация через API Gateway в сторону нужного микросервиса исходя от запроса потребителя
Так как у нас несколько потребителей и чтобы все безболезнено имели доступ, лучше использовать REST
2. Внутри думал сначала сделать все тоже через REST но так как у нас есть использование видеонаблюдения, а в таком случае как будто gRPC целесообразнее - то решил попробовать все контракты сделать через gRPC. Хотя ранее на работе я с gRPC не сталкивался и нам хватало REST, я решил все таки попробовать, думаю будет куча ошибок - но в первую очередь тут хотел попробовать спроектировать такие интеграции.
У gRPC тут есть несколько плюсов: 
- Он выше в производительности чем REST
- Строгий контракт 
- Стриминг (который, как мне кажется сильно важен например для видеонаблюдения)
- Микросервисы написаны на Go, что более совместимо с gRPC и кодогенерация будет осуществляться проще

### 2. Документация API
1. REST Синхронный - https://jdegre.github.io/editor/?url=https://raw.githubusercontent.com/DmitrakovPV/YP-Projects/refs/heads/warmhouse/docs/docs/api/rest/User%26HouseService.yaml 
2. gRPC - для gRPC к сожалению нет WEB интерфейса для превью, только если оборачивать в REST
	- user: https://raw.githubusercontent.com/DmitrakovPV/YP-Projects/refs/heads/warmhouse/docs/docs/api/gRPC/user.proto
	- house: https://raw.githubusercontent.com/DmitrakovPV/YP-Projects/refs/heads/warmhouse/docs/docs/api/gRPC/house.proto
	- device: https://raw.githubusercontent.com/DmitrakovPV/YP-Projects/refs/heads/warmhouse/docs/docs/api/gRPC/device.proto
	- heating: https://raw.githubusercontent.com/DmitrakovPV/YP-Projects/refs/heads/warmhouse/docs/docs/api/gRPC/heating.proto
	- lightning: https://raw.githubusercontent.com/DmitrakovPV/YP-Projects/refs/heads/warmhouse/docs/docs/api/gRPC/lightning.proto
	- access: https://raw.githubusercontent.com/DmitrakovPV/YP-Projects/refs/heads/warmhouse/docs/docs/api/gRPC/access.proto
	- surveliance: https://raw.githubusercontent.com/DmitrakovPV/YP-Projects/refs/heads/warmhouse/docs/docs/api/gRPC/surveillance.proto
	- schedule: https://raw.githubusercontent.com/DmitrakovPV/YP-Projects/refs/heads/warmhouse/docs/docs/api/gRPC/schedule.proto
	- pbuf: https://raw.githubusercontent.com/DmitrakovPV/YP-Projects/refs/heads/warmhouse/docs/docs/api/gRPC/pbuf.yaml
	- pbufGen: https://raw.githubusercontent.com/DmitrakovPV/YP-Projects/refs/heads/warmhouse/docs/docs/api/gRPC/pbufGen.yaml
3. AsyncAPI - https://studio.asyncapi.com/?url=https://raw.githubusercontent.com/DmitrakovPV/YP-Projects/refs/heads/warmhouse/docs/docs/api/eventsAPI.yaml

# Задание 5. Работа с docker и docker-compose

Перейдите в apps.

Там находится приложение-монолит для работы с датчиками температуры. В README.md описано как запустить решение.

Вам нужно:

1) сделать простое приложение temperature-api на любом удобном для вас языке программирования, которое при запросе /temperature?location= будет отдавать рандомное значение температуры.

Locations - название комнаты, sensorId - идентификатор названия комнаты

```
	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}
```
Пункт выполнен - смотреть содержимое папки temperature-api

2) Приложение следует упаковать в Docker и добавить в docker-compose. Порт по умолчанию должен быть 8081

Пункт выполнен - но у меня порт 8080 был занят PlantUML контейнером, новый сервис повесил на 8081 
Но порт откуда должен запускаться монолит я поменял на 8088 просьба обратить внимание, когда будете дергать запросы в postman


3) Кроме того для smart_home приложения требуется база данных - добавьте в docker-compose файл настройки для запуска postgres с указанием скрипта инициализации ./smart_home/init.sql

Для проверки можно использовать Postman коллекцию smarthome-api.postman_collection.json и вызвать:

- Create Sensor
- Get All Sensors

Должно при каждом вызове отображаться разное значение температуры

Ревьюер будет проверять точно так же.

Пункт выполнен - при вызове  Get All Sensors возвращается список всех сенсоров после их создания Create Sensor


