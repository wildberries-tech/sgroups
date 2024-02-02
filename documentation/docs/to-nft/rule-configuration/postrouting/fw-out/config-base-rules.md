---
id: config-base-rules
---

# Config Base Rules

### Описание

### Параметры

<table>
    <thead>
        <tr>
            <th>Шаблон параметра</th>
            <th>Структура параметра</th>
            <th>Значение</th>
            <th>Описание</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>\{DstCIDR\}</td>
            <td>
                <nobr>`daddr {CIDR}`</nobr>
            </td>
            <td>Массив подсетей</td>
            <td>Список сетей в которые разрешаем трафик</td>
        </tr>
        <tr>
            <td>\{Verdict\}</td>
            <td>Accept</td>
            <td>
                <div>Не параметризированный</div>
                <br />
                <div class="text-justify">
                    <i>* Так как данное правило используется для проверки типа трафика то переход на другую цепочку правил происходит только с помощью goto.</i>
                </div>
                <i>Подробнее: [Verdict statement](../main.md#verdict-statement)</i>              
            </td>
            <td>Вердикт политики по пакетам данных</td>
        </tr>
    </tbody>
</table>

### Шаблон

```
chain FW-OUT {
  # **********
  ${DstCIDR} ${Verdict}
  # **********
}
```

### Пример использования

```
chain FW-OUT {
  # **********
  ip daddr { 1.1.1.1, 2.2.2.2} accept
  # **********
}
```


