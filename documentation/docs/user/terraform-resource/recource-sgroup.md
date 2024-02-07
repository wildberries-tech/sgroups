---
id: recource-sgroup
---

## **Описание**

sgroups_groups - это ресурс Terraform для определения групп безопасности.

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
            <td>name</td>
            <td>да</td>
            <td>string</td>
            <td>Имя группы безопасности.</td>
        </tr>
        <tr>
            <td>networks</td>
            <td>да</td>
            <td>string</td>
            <td>Имя сети которая включена в указанную группу безопасности.</td>
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
        <tr>
            <td>default_action</td>
            <td>да</td>
            <td>&quot;DROP\|ACCEPT&quot;</td>
            <td>Действие по умолчанию для группы безопасности. По умолчанию значение &quot;DROP&quot;. Может быть установлено в &quot;DROP&quot; или &quot;ACCEPT&quot;.</td>
        </tr>
        <tr>
            <td>icmp</td>
            <td>да</td>
            <td>object</td>
            <td>Структура, описывающая взаимодействие с ICMP-трафиком для IPv4 по умолчанию.</td>
        </tr>
        <tr>
            <td>icmp.logs</td>
            <td>нет</td>
            <td>bool</td>
            <td>Определяет, включены ли логи для ICMP-трафика. По умолчанию значение &quot;false&quot;. Установите в &quot;true&quot;, если требуется включить логирование.</td>
        </tr>
        <tr>
            <td>icmp.trace</td>
            <td>нет</td>
            <td>bool</td>
            <td>Определяет, включен ли трассировочный режим для ICMP-трафика. По умолчанию значение &quot;false&quot;. Установите в &quot;true&quot;, если требуется включить трассировку.</td>
        </tr>
        <tr>
            <td>icmp.type[]</td>
            <td>да</td>
            <td>array of strings</td>
            <td>Список, определяющий допустимые типы ICMP для IPv4 запросов.</td>
        </tr>
        <tr>
            <td>icmp6</td>
            <td>да</td>
            <td>object</td>
            <td>Структура, описывающая взаимодействие с ICMP-трафиком для IPv6 по умолчанию.</td>
        </tr>
        <tr>
            <td>icmp6.logs</td>
            <td>нет</td>
            <td>bool</td>
            <td>Определяет, включены ли логи для ICMP-трафика. По умолчанию значение &quot;false&quot;. Установите в &quot;true&quot;, если требуется включить логирование.</td>
        </tr>
        <tr>
            <td>icmp6.trace</td>
            <td>нет</td>
            <td>bool</td>
            <td>Определяет, включен ли трассировочный режим для ICMP-трафика. По умолчанию значение &quot;false&quot;. Установите в &quot;true&quot;, если требуется включить трассировку.</td>
        </tr>
        <tr>
            <td>icmp6.type[]</td>
            <td>да</td>
            <td>array of strings</td>
            <td>Список, определяющий допустимые типы ICMP для IPv6 запросов.</td>
        </tr>
    </tbody>
</table>

## **Пример использования**

```hcl
resource "sgroups_groups" "groups" {
    items = {
      key = {
        name = "sg-1"
        networks = "nw-2"
        logs = true
        trace = false
        default_action = "ACCEPT"
        icmp = {
          logs = false
          trace = false
          type = [0, 8]
          }
        icmp6 = {
          logs = false
          trace = false
          type = [0, 8]
          }
      }
    }
}
```

## **Ограничения**

* Каждая группа должна иметь уникальный ключ
* name и networks - длина значения не должна превышать 256 символов. Значение должно начинаться и заканчиваться символами без пробелов.
* type - значения в списке не должны повторяться. Значение должно быть цифрой от 0 до 255.