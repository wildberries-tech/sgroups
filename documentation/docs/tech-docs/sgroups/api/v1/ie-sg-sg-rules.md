---
id: ie-sg-sg-rules
---

# POST v1/ie-sg-sg/rules

## Запрос

`POST v1/ie-sg-sg/rules`

<ul>
    <li className="text-justify">если в теле запроса указано одно или более значений в обоих массивах sg -\> sg_local, то получим ответ всех существующих комбинаций правил IE-SG-SG каждого указанного значения sg с каждым указанным значением sg_local (value-to-value)</li>
    <li className="text-justify">если в теле запроса один из массивов пустой а во втором указаны от одного и более значений, то получим ответ всех существующих комбинаций правил IE-SG-SG каждого указанного значения со всеми существующими (any-to-value, value-to-any)</li>
    <li className="text-justify">если в теле запроса указаны пустые массивы sg -\> sg_local, то получим ответ всех существующих комбинаций правил IE-SG-SG (any-to-any)</li>
    <li className="text-justify">если указано некорректное тело в запросе, то получим ответ всех существующих комбинаций правил IE-SG-SG (any-to-any)</li>
</ul>

```json
{
  "sg": ["sg-1"],
  "sg_local": ["sg-2"]
}
```

## Ответ

```json
{
  "rules": [
    {
      "sg": "sg-1",
      "sg_local": "sg-2",
      "logs": true,
      "trace": true,
      "ports": [
        {
          "d": "7800",
          "s": "4446"
        }
      ],
      "traffic": "Ingress",
      "transport": "TCP"
    }
  ]
}
```

## Входные параметры

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
            <td>sg</td>
            <td>array of strings</td>
            <td>да</td>
            <td>массив из имен источников SG</td>
            <td>sg-11</td>
        </tr>
        <tr>
            <td>2</td>
            <td>sg_local</td>
            <td>array of strings</td>
            <td>да</td>
            <td>массив из имен источников SG</td>
            <td>sg-12</td>
        </tr>
    </tbody>
</table>

## Проверки

<table>
    <thead>
        <tr>
            <th>Параметр</th>
            <th>Условие</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>sg</td>
            <td>\- длина значения не должна превышать 256 символов&lt;br /&gt;\- значение должно начинаться и заканчиваться символами без пробелов</td>
        </tr>
        <tr>
            <td>sg_local</td>
            <td>\- длина значения не должна превышать 256 символов&lt;br /&gt;\- значение должно начинаться и заканчиваться символами без пробелов</td>
        </tr>
    </tbody>
</table>

## Выходные параметры

### Положительный ответ

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
            <td>rules</td>
            <td>array of objects</td>
            <td></td>
            <td>\-</td>
        </tr>
        <tr>
            <td>1.2</td>
            <td>rules[].sg</td>
            <td>string</td>
            <td>название Security group</td>
            <td>sg-0</td>
        </tr>
        <tr>
            <td>1.3</td>
            <td>rules[].sg_local</td>
            <td>string</td>
            <td>название Security group</td>
            <td>sg-0</td>
        </tr>
        <tr>
            <td>1.4</td>
            <td>rules[].logs</td>
            <td>bool</td>
            <td>включено или выключено логирование (по умолчанию выключено)</td>
            <td>true/false</td>
        </tr>
        <tr>
            <td>1.5</td>
            <td>rules[].trace</td>
            <td>bool</td>
            <td>включена или выключена трассировка (по умолчанию выключена)</td>
            <td>true/false</td>
        </tr>
        <tr>
            <td>1.6</td>
            <td>rules[].ports</td>
            <td>array of objects</td>
            <td></td>
            <td>\-</td>
        </tr>
        <tr>
            <td>1.6.1</td>
            <td>rules[].ports[].d</td>
            <td>string</td>
            <td>значения портов входящего трафика</td>
            <td>&quot;7600-7700,7800&quot;</td>
        </tr>
        <tr>
            <td>1.6.2</td>
            <td>rules[].ports[].s</td>
            <td>string</td>
            <td>значения портов исходящего трафика</td>
            <td>&quot;4446&quot;</td>
        </tr>
        <tr>
            <td>1.7</td>
            <td>rules[].traffic</td>
            <td>string</td>
            <td>тип траффика (входящий/исходящий)</td>
            <td>&quot;Undef&quot;/&quot;Ingress&quot;/&quot;Egress&quot;</td>
        </tr>
        <tr>
            <td>1.8</td>
            <td>rules[].transport</td>
            <td>string</td>
            <td>метод передачи данных</td>
            <td>&quot;TCP&quot;/&quot;UDP&quot;</td>
        </tr>
    </tbody>
</table>

### Ответ с ошибками

Код ошибки 400

- Если sg или sg_local были указаны некорректно:
  \- ошибка, если значения были указаны не как массив, а как одно значение
  \- ошибка, если значение sg или sg_local не соответствует формату названия security group (длина значения не должна превышать 256 символов, значения должно начинаться и заканчиваться символами без пробелов, значение должно быть уникальным)

```json
{
  "code": 3,
  "details": [],
  "message": "proto: syntax error (line __): unexpected token \"string\""
}
```

Код ошибки 404

- Опечатка в имени метода

```json
{
  "code": 5,
  "details": [],
  "message": "Not Found"
}
```

## Описание интеграции

```mermaid
sequenceDiagram
participant user as User
participant server as Server
participant db as Database

user->>server: Отобразить список IE-SG-SG правил для входящего и исходящего траффика

alt Ошибка в запросе
    server-->>user: Показать ошибку в запросе
end

server->>db: Отправить запрос
db->>db: Проверка входящего запроса

alt sg и/или sg_local были указаны некорректно
    db-->>server: Ответ с ошибкой
    server-->>user: Показать ошибку в запросе
end

db-->>server: Ответ со списком IE-SG-SG правил входящего и исходящего траффика соответствующий запросу
server-->>user: Список IE-SG-SG правил входящего и исходящего траффика соответствующий запросу
```
