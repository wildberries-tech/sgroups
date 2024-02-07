---
id: recource-rule-ie-s2s-tcp-udp
---

## **Описание**

sgroups_ie_rules - это ресурс Terraform для определения правил безопасности на основе групп безопасности, протокола и портов.

## **Параметры Ресурса**

<table>
    <thead>
        <tr>
            <th>название</th>
            <th>обязательность</th>
            <th>тип данных</th>
            <th>доп. описание</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>traffic</td>
            <td>да</td>
            <td>&quot;ingress\|egress&quot;</td>
            <td>Направление трафика для данного правила. ingress - входящий траффик. egress - исходящий траффик.</td>
        </tr>
        <tr>
            <td>transport</td>
            <td>да</td>
            <td>&quot;TCP\|UDP&quot;</td>
            <td>Тип протокола для данного правила. Должен быть установлен в &quot;TCP&quot; или &quot;UDP&quot;.</td>
        </tr>
        <tr>
            <td>sg</td>
            <td>да</td>
            <td>string</td>
            <td>Имя исходной группы безопасности.</td>
        </tr>
        <tr>
            <td>sg_local</td>
            <td>да</td>
            <td>string</td>
            <td>Имя целевой группы безопасности.</td>
        </tr>
        <tr>
            <td>ports</td>
            <td>нет</td>
            <td>array of object</td>
            <td>Массив объектов, представляющих порты. Каждый объект содержит поля s (source) и d (destination).</td>
        </tr>
        <tr>
            <td>ports[].d</td>
            <td>нет</td>
            <td>string</td>
            <td>Значения s (source) портов в рамках одного правила не должны пересекаться</td>
        </tr>
        <tr>
            <td>ports[].s</td>
            <td>нет</td>
            <td>string</td>
            <td>Значения d (destination) портов в рамках одного правила могут пересекаться</td>
        </tr>
        <tr>
            <td>logs</td>
            <td>нет</td>
            <td>bool</td>
            <td>Определяет, включены ли логи для данного правила. По умолчанию значение &quot;false&quot;. Установите в &quot;true&quot;, если требуется включить логирование.</td>
        </tr>
        <tr>
            <td>trace</td>
            <td>нет</td>
            <td>bool</td>
            <td>Определяет, включен ли трассировочный режим для данного правила. По умолчанию значение &quot;false&quot;. Установите в &quot;true&quot;, если требуется включить трассировку.</td>
        </tr>
    </tbody>
</table>

## **Пример использования**

```hcl
resource "sgroups_ie_rules" "rules" { 
  items = { 
    "tcp:sg(sg_example)sg(sg_local_example-1)ingress" = {
      traffic     = "ingress"
      transport   = "tcp"
      sg          = "sg_example"
      sg_local    = "sg_local_example-1"
      ports = [
        {
          d = "80"
          s = ""
        }
      ]
      logs  = false
      trace = false
    }
  }
}
```

## **Ограничения**

* Каждое правило должно иметь уникальный ключ, чтобы избежать пересечений.
* Структура формы items должна соответствовать определенной последовательности которой необходимо придерживаться `"${transport}:sg${SG)}sg${sg_local}${traffic}"`
* sg и sg_local - длина значения не должна превышать 256 символов. Значение должно начинаться и заканчиваться символами без пробелов.
* Значения s (source) портов не должно пересекаться в рамках одного правила состоящего из уникального составного ключа (traffic, transport, sg, sg_local).
* Значения s (source) и d (distinct) портов должно находиться в интервале от 1 до 65535. Если значение не будет указано то будет использоваться весь диапазон портов.