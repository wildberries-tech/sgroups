---
id: default-rule-all
---

# Sgroup Default Rule (all)

### Описание

Правило по умолчанию в цепочке Security Group для входящего трафика. Имеет последний приоритет в цепочке FW-IN-sgName и отрабатывает в самом конце.

### Параметры

<table>
    <thead>
        <tr>
            <th>Наименование параметра</th>
            <th>Структура параметра</th>
            <th>Описание</th>
            <th>Значения</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>\{Trace\}</td>
            <td>nftrace set</td>
            <td>
                <ul>
                    <li><b>1</b> - трассировка включена</li>
                    <li><b>0</b> - трассировка выключена</li>
                </ul>               
            </td>
            <td>Трассировка указанного правила (опциональна, можно включить/выключить)</td>
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
            <td>\{Log\}</td>
            <td>log level debug flags ip options</td>
            <td>Не параметризированный&lt;br/&gt;</td>
            <td>Логирование указанного правила (опциональна, можно включить/выключить)</td>
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

```hcl
chain FW-IN-sgName {
    # **********
		${Trace} ${Counter} ${Log} ${Verdict}
		# **********
}
```

### Пример использования

```hcl
chain FW-IN-sgname_example {
  # **********
  nftrace set 1 counter packets 0 bytes 0 log level debug flags ip options accept
  # **********
}
```