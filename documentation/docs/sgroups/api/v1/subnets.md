---
id: subnets
---

# GET /v1/sg/\{sgName\}/subnets

## **Запрос**

`GET /v1/sg/{sgName}/subnets`

## **Ответ**

```json
 {
  "networks": [
     {
     "name": "nw-2",
     "network":  {
       "CIDR": "10.150.0.222/32"
      }
    } 
   ]
 }
```

## **Входные параметры**

<table>
    <thead>
        <tr>
            <th>№</th>
            <th>Параметр</th>
            <th>Тип данных</th>
            <th>Обязательность</th>
            <th>Описание</th>
            <th>Варианты значений</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1</td>
            <td>\{sgName\}</td>
            <td>string</td>
            <td>да</td>
            <td>уникальное имя sg</td>
            <td>SG-11</td>
        </tr>
    </tbody>
</table>

## **Проверки**

<table>
    <thead>
        <tr>
            <th>Параметр</th>
            <th>Условие</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>\{sgName\}</td>
            <td>\- длина значения не должна превышать 256 символов&lt;br /&gt;\- значение должно начинаться и заканчиваться символами без пробелов</td>
        </tr>
    </tbody>
</table>

## **Выходные параметры**

### **Положительный ответ**

<table>
    <thead>
        <tr>
            <th>№</th>
            <th>Параметр</th>
            <th>Тип данных</th>
            <th>Описание</th>
            <th>Варианты значений</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>1</td>
            <td>networks</td>
            <td>array of objects</td>
            <td></td>
            <td>\-</td>
        </tr>
        <tr>
            <td>1.1</td>
            <td>networks[].name</td>
            <td>string</td>
            <td>уникальное имя сети</td>
            <td>nw-1</td>
        </tr>
        <tr>
            <td>1.2</td>
            <td>networks[].network</td>
            <td>object</td>
            <td></td>
            <td>\-</td>
        </tr>
        <tr>
            <td>1.3</td>
            <td>networks[].network.CIDR</td>
            <td>string</td>
            <td></td>
            <td>10.150.0.220/32</td>
        </tr>
    </tbody>
</table>

### **Ответ с ошибками**

Код ошибки 404

* Указано неверное значение \{sgName\}

```json
 {
  "code": 5,
  "details":  [],
  "message": "SG 'string' is not found"
 }
```

* Ошибка в запросе

```json
 {
  "code": 5,
  "details":  [],
  "message": "Not Found"
 }
```

## **Описание интеграции**

```mermaid
sequenceDiagram
participant user as User
participant server as Server
participant db as Database

user->>server: Отобразить список доступных сетей связанных с SG

alt Ошибка в запросе
    server-->>user: Показать ошибку в запросе 
end

server->>db: Отправить запрос
db->>db: Проверка входящего запроса

alt Возникла ошибка при некорректном вводе SG
    db-->>server: Ответ с ошибкой
    server-->>user: Показать ошибку в запросе
end

db-->>server: Ответ со списком доступных сетей связанных с SG соответствующий запросу
server-->>user: Список доступных сетей связанных с SG соответствующий запросу
```