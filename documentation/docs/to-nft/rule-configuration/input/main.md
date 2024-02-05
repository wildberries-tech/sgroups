---
id: main
---

# INPUT

### Описание

Правило перехода в цепочку FW-IN с проверкой что трафик является входящим.

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
            <td>$\{Condition\}</td>
            <td>
                <nobr>`iifname != "lo"`</nobr>
            </td>
            <td>Не параметризированный</td>
            <td>Проверка что трафик является входящим</td>
        </tr>
        <tr>
            <td>$\{Counter\}</td>
            <td>
                <nobr>`counter packets 0 bytes 0`</nobr>
            </td>
            <td>Не параметризированный</td>
            <td>Счетчик, учитывает количество пройденных пакетов с количеством байтов переданной информации в рамках указанной цепочки правил</td>
        </tr>
        <tr>
            <td>$\{Verdict\}</td>
            <td>`goto`</td>
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
        <tr>
            <td>$\{Hook\}</td>
            <td>`hook {Name}`</td>
            <td>input</td>
            <td>Приоритет выполнения цепочки характерезующий стадию прохождения трафика</td>
        </tr>
                <tr>
            <td>$\{HookPriority\}</td>
            <td>`priority {IntValue}`</td>
            <td>0</td>
            <td>Приоритет выполнения цепочки одного типа</td>
        </tr>
    </tbody>
</table>

### Шаблон

```hcl
chain INPUT {  
    type filter ${Hook} ${HookPriority}; policy accept;
    # **********
    ${Condition} ${Counter} ${Verdict} FW-IN
    # **********
}
```

### Пример использования

```hcl
chain INPUT {
    type filter hook input priority 0; policy accept;
    # **********
    iifname != "lo" counter packets 0 bytes 0 goto FW-IN
    # **********
}
```



