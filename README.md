## Необходимо для запуска:
* Docker Compose

## Запуск
Клонирование репозитория
```bash
git clone https://github.com/aleksandrzhigulin/VK.git
```
Переход в папку
```bash
cd VK
```
Сборка и запуск
```bash
docker-compose up --build
```
Дождаться запуска
![image](https://github.com/aleksandrzhigulin/VK/assets/66275482/24e1287f-bdb6-4cd1-ba03-ebe29fb42a8a)
Стек: Go (go-chi, database/sql), PostgreSQL, Docker, приложение работает на порте **8080**, postgresql на порте **5432**
# Описание API
## GET Методы
**GET localhost:8080/user/{user_id}/** - *возвращает айди пользователя, его имя и баланс* <br/> <br/>
**GET localhost:8080/user/history/{user_id}** - *возвращает выполненные пользователем задания* <br/> <br/>
**GET localhost:8080/quest/completed/{user_id}/{quest_id}** - *сигнализирует о выполнении пользователем с id {user_id} задания с id {quest_id} и начисляет награду* <br/> <br/>
## POST Методы
**POST localhost:8080/user/add** - *создаёт нового пользователя.* <br/>
Тело запроса:
```json
{
  "name": "{user_name}"
}
```
**POST localhost:8080/quest/add** - *создаёт новое задание* <br/>
Тело запроса:
```json
{
    "name": "{quest_name}",
    "cost": {quest_cost}
}
```
name - string <br/>
cost - int
## Примеры работы API
**GET localhost:8080/user/{user_id}/** 
![GET_USER](https://github.com/aleksandrzhigulin/VK/assets/66275482/090b5545-e1ef-4aba-9f07-abdd05b70b55)
<br/> <br/>
**GET localhost:8080/user/history/{user_id}**
![GET_USER_HISTORY](https://github.com/aleksandrzhigulin/VK/assets/66275482/c90a0557-1109-4ac3-9d1d-314a7484f0c4)
<br/> <br/>
**GET localhost:8080/quest/completed/{user_id}/{quest_id}**
![GET_QUEST_COMPLETED](https://github.com/aleksandrzhigulin/VK/assets/66275482/d1b8c382-cea4-4dc9-88ca-07e2b899d314)
![GET_QUEST_COMPLETED_ERROR](https://github.com/aleksandrzhigulin/VK/assets/66275482/92a5fae7-4426-4691-a8ec-9525de7ad51d)
<br/> <br/>
**POST localhost:8080/user/add**
![POST_USER_ADD](https://github.com/aleksandrzhigulin/VK/assets/66275482/faae24ef-97a4-4725-9b46-19bd7410dc40)
<br/> <br/>
**POST localhost:8080/quest/add**
![POST_QUEST_ADD](https://github.com/aleksandrzhigulin/VK/assets/66275482/9c5d77e3-4eca-438e-ad93-000c66785571)
