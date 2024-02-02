---
id: main
---

### Описание

Правило перехода в цепочку FW-OUT с проверкой что трафик является исходящим.

### Параметры

<table>
    <thead>
        <tr>
            <th>Шаблон параметра</th>
            <th>Структура</th>
            <th>Значение</th>
            <th>Описание</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>\{Condition\}</td>
            <td>
                <nobr>`oifname != "lo"`</nobr>
            </td>
            <td>Не параметризированный</td>
            <td>Проверка что трафик является исходящим</td>
        </tr>
        <tr>
            <td>\{Counter\}</td>
            <td>
                <nobr>`counter packets 0 bytes 0`</nobr>
            </td>
            <td>Не параметризированный</td>
            <td>Счетчик, учитывает количество пройденных пакетов с количеством байтов переданной информации в рамках указанной цепочки правил</td>
        </tr>
        <tr>
            <td>\{Verdict\}</td>
            <td>goto</td>
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

```hcl
chain FW-IN {  
  # **********
  ${Condition} ${Counter} ${Verdict} FW-OUT
  # **********
}
```

### Пример использования

```hcl
chain POSTROUTING {
    type filter hook postrouting priority 300; policy accept;
    # **********
		oifname != "lo" counter packets 0 bytes 0 goto FW-OUT
		# **********
	}
```