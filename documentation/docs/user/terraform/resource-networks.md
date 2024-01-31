---
id: resource-networks
---

## **Описание**

sgroups_networks - это ресурс Terraform для определения сетевых параметров, используемых в группах безопасности.

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
            <td>Имя сети.</td>
        </tr>
        <tr>
            <td>cidr</td>
            <td>да</td>
            <td>string</td>
            <td>CIDR блок для данной сети.</td>
        </tr>
    </tbody>
</table>

## **Пример использования**

```hcl
resource "sgroups_networks" "networks" {
    items = {
      key = {
        name = "nw-1"
        cidr = "10.0.0.0/24"
        }
    }
}
```

## **Ограничения**

* Каждая сеть должна иметь уникальный ключ.
* Не должно быть пересечений в CIDR.