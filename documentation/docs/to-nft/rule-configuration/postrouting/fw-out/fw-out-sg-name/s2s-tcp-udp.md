---
id: s2s-tcp-udp
---

# Sgroup to Sgroup (tcp|udp)

### Описание

Правило в обычной цепочке для исходящего трафика от другой Security Group. Имеет второй приоритет в цепочке FW-OUT-sgName.

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
            <td>$\{DstSgroup\}</td>
            <td>
                <nobr>`saddr @${IPSet(sgName)}`</nobr>            
            </td>
            <td>Наименование IPSet в котором описаны сети в Security Group</td>
            <td>Значение типа string, не должно содержать в себе пробелов</td>
        </tr>
        <tr>
            <td>$\{Transport\}</td>
            <td>`tcp` | `udp`</td>
            <td>протокол передачи данных в цепочке правил.</td>
            <td>Одно из двух значений `tcp` | `udp`</td>
        </tr>
        <tr>
            <td>$\{RuleType\}</td>
            <td>`ip`</td>
            <td></td>
            <td>Описывает, что принимает трафик типа ip</td>
        </tr>
        <tr>
            <td>$\{SrcPorts\}</td>
            <td>`sport {}`</td>
            <td>Набор целочисленных значений от 0 до 65535З</td>
            <td class="text-justify">Значения `sport` (source port). Может быть как одно значение, как и множество значений портов. В случае если одно значение у порта то передается значение либо как целочисленное значение либо как название порта. Если передается массив значений портов то они должны быть внутри `{}` перечислены через запятую.</td>
        </tr>
        <tr>
            <td>$\{DstPorts\}</td>
            <td>`dport {}`</td>
            <td>Набор целочисленных значений от 0 до 65535З</td>
            <td class="text-justify">Значения `dport` (destination port). Может быть как одно значение, как и множество значений портов. В случае если одно значение у порта то передается значение либо как целочисленное значение либо как название порта. Если передается массив значений портов то они должны быть внутри `{}` перечислены через запятую.</td>
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
    ${Trace} ${RuleType} ${DstSgroup} ${Transport} ${SrcPorts} ${DstPorts} ${Counter} ${Log} ${Verdict}
    # **********
}
```

### Пример использования

```hcl
chain FW-OUT-sgname_example {
    # **********
    nftrace set 1 ip daddr NetIPv4-sgName_example tcp dport { 80, 443 } sport { 80, 443 } counter packets 0 bytes 0 log level debug flags ip options accept
    nftrace set 1 ip daddr NetIPv6-sgName_example tcp dport { 80, 443 } sport { 80, 443 } counter packets 0 bytes 0 log level debug flags ip options accept
    # **********
}
```