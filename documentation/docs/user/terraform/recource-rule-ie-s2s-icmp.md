---
id: recource-rule-ie-s2s-icmp
---

## **Описание**



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
            <td>sg_local</td>
            <td>да</td>
            <td>string</td>
            <td>Имя исходной группы безопасности.</td>
        </tr>
        <tr>
            <td>sg</td>
            <td>да</td>
            <td>string</td>
            <td>Имя целевой группы безопасности.</td>
        </tr>
        <tr>
            <td>ip_v</td>
            <td>да</td>
            <td>&quot;IPv4\|IPv6&quot;</td>
            <td>Версия IP для ICMP (IPv4 или IPv6).</td>
        </tr>
        <tr>
            <td>type</td>
            <td>да</td>
            <td>array of strings</td>
            <td>Список, определяющий допустимые типы ICMP запросов.</td>
        </tr>
        <tr>
            <td>traffic</td>
            <td>да</td>
            <td>&quot;ingress\|egress&quot;</td>
            <td>Направление трафика для данного правила. ingress - входящий траффик. egress - исходящий траффик.</td>
        </tr>
    </tbody>
</table>

## **Пример использования**

```hcl
resource "sgroups_icmp_rules" "rules" {
  items = {
    "sg(sg-example-local)sg(sg-example)icmp4 TODO" => {
        logs    = true
        trace   = true
        sg_local = sg-example-local
        sg       = sg-example
        ip_v    = IPv4
        type    = [0,8]
        traffic = ingress
    }
    "sg(sg-example-local)sg(sg-example)icmp6 TODO" => {
        logs    = true
        trace   = true
        sg_local = sg-example-local
        sg       = sg-example
        ip_v    = IPv6
        type    = [0,8]
        traffic = ingress
    }
  }
}
```

## **Ограничения**


