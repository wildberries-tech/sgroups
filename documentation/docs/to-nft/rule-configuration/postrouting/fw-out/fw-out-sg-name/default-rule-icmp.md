---
id: default-rule-icmp
---

# Sgroup Default Rule (icmp)

### Описание

Правило в обычной цепочке для исходящего трафика по ICMP. Имеет первый приоритет в цепочке FW-OUT-sgName.

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
            <td>$\{Trace\}</td>
            <td>`nftrace set`</td>
            <td>
                <ul>
                    <li><b>1</b> - трассировка включена</li>
                    <li><b>0</b> - трассировка выключена</li>
                </ul>            
            </td>
            <td>Трассировка указанного правила (опциональна, можно включить/выключить)</td>
        </tr>
        <tr>
            <td>$\{Transport\}</td>
            <td>`icmp`</td>
            <td></td>
            <td>Указывает на протокол транспортного уровня</td>
        </tr>
        <tr>
            <td>$\{TypeList\}</td>
            <td>
                <nobr>`type ${TypeList}`</nobr>
            </td>
            <td>
                <div class="text-justify">
                    Список который содержит от 0 до 255 элементов в строковом представлении и указывает на числовой код ICMP
                </div>
                <i>Подробнее: [дескрипторы для ICMP](../../../main.md#icmp-descriptors)</i>            
            </td>
            <td>Значение кодов типа ICMP</td>
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
            <td>$\{Log\}</td>
            <td>
                <nobr>`log level debug flags ip options`</nobr>
            </td>
            <td>Не параметризированный</td>
            <td>Логирование указанного правила (опциональна, можно включить/выключить)</td>
        </tr>
        <tr>
            <td>$\{Verdict\}</td>
            <td>`Accept`</td>
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
chain FW-OUT-sgName {
    # **********
    ${Trace} ${RuleType} ${Counter} ${Log} ${Verdict}
    # **********
}
```

### Пример использования

```hcl
chain FW-OUT-sgname_example {
    # **********
    nftrace set 1 icmp type { echo-reply, echo-request } counter packets 0 bytes 0 log level debug flags ip options accept
    # **********
}
```