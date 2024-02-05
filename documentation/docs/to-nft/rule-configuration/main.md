---
id: main
---

# Конфигурация правил (nftables cli)

### Описание

<div class="paragraph">
[*nftables*](https://wiki.nftables.org/) представляет собой инструмент для управления сетевыми правилами, предоставляя гибкую и эффективную систему фильтрации трафика.
</div>

## IPSet (set)

### Описание

<div class="paragraph">
IPSet представляет собой структуру данных, позволяющую хранить и организовывать множество элементов для использования в правилах nftables. Он используется для обработки больших списков элементов с минимальными затратами на ресурсы. В нашем случае IPSet формируется для описания массива IP-адресов, относящихся к FQDN или Security Groups, для дальнейшего использования в описаниях правил.
</div>

### Параметры

<table>
    <thead>
        <tr>
            <td>Наименование параметра</td>
            <td>Описание</td>
            <td>Значения</td>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>IPSet_Name</td>
            <td>Наименование IPSet</td>
            <td>
                Должно соответствовать одному из шаблонов:
                <ul>
                    <li><b>NetIPv4-$\{NAME\}</b> - для описания массивов IP адресов типа v4</li>
                    <li><b>NetIPv6-$\{NAME\}</b> - для описания массивов IP адресов типа v6</li>
                </ul>            
            </td>
        </tr>
        <tr>
            <td>type</td>
            <td>Описывает тип данных</td>
            <td>
                Могут быть установлены следующие значения:
                <ul>
                    <li><b>ipv4_addr</b> - для описания массивов IP адресов типа v4</li>
                    <li><b>ipv6_addr</b> - для описания массивов IP адресов типа v6</li>
                </ul>
            </td>
        </tr>
        <tr>
            <td>flags</td>
            <td>Описывает свойства IPSet</td>
            <td>
                Установлены следующие значения:
                <ul>
                    <li><b>constant</b> - флаг используется если значение элементов в множестве являются постоянными и не могут быть изменены</li>
                    <li><b>interval</b> - флаг используется для создания диапазона элементов множества</li>
                </ul>
            </td>
        </tr>
        <tr>
            <td>elements</td>
            <td>Указывает массив содержащихся в IPSet элементов подсетей (CIDR)</td>
            <td>Значения CIDR, в случае нескольких значений перечисляются через запятую</td>
        </tr>
    </tbody>
</table>

### Пример

```hcl
set NetIPv4-${NAME} {
    type ipv4_addr
    flags constant,interval
    elements = { 10.168.24.0/23 }
}
```



## Цепочки (Chains)

### Описание

<div class="paragraph">
В нашей реализации структуры мы вводим для INPUT и OUTPUT понятие 3-х цепочек. *Первая* цепочка является точкой входа для пакетов из сетевого стека, в ней указывается хук (input, prerouting, postrouting) и приоритет выполнения. *Вторая* цепочка используется для маршрутизации в последующие цепочки по принадлежности к той или иной Security Group. Последняя цепочка содержит наборы правил, относящиеся только к конкретной Security Group.
</div>

<table>
    <caption>Цепочки для входящего трафика</caption>
    <thead>
        <tr>
            <th>Наименование цепочки</th>
            <th>Описание</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>INPUT</td>
            <td><i>Первая</i> цепочка является точкой входа для пакетов из сетевого стека, в ней указывается хук (input) и приоритет выполнения 0 (filter)</td>
        </tr>
        <tr>
            <td>FW-IN</td>
            <td><i>Вторая</i> цепочка используется для маршрутизации в последующие цепочки по принадлежности к той или иной Security Group</td>
        </tr>
        <tr>
            <td>FW-IN-sgName (Security Group name)</td>
            <td><i>Третья</i> цепочка содержит наборы правил, относящиеся только к конкретной Security Group.</td>
        </tr>
    </tbody>
</table>

```hcl
table inet main-1705582480 {

chain FW-IN-exampleSG {
    # ******
    counter packets 0 bytes 0 accept
}

chain FW-IN {
    # ******
    ip daddr @NetIPv4-exampleSG counter packets   0 bytes 0 goto FW-IN-exampleSG
    # ******
  }

chain INPUT {
    type filter hook input priority 0; policy accept;
    ip daddr @NetIPv4-exampleSG counter packets   0 bytes 0 goto FW-IN
  }
}
```

<table>
    <caption>Цепочки для исходящего трафика</caption>
    <thead>
        <tr>
            <th>Наименование цепочки</th>
            <th>Описание</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>POSTROUTING</td>
            <td><i>Первая</i> цепочка является точкой входа для пакетов из сетевого стека, в ней указывается хук (postrouting) и приоритет выполнения 300</td>
        </tr>
        <tr>
            <td>FW-OUT</td>
            <td><i>Вторая</i> цепочка используется для маршрутизации в последующие цепочки по принадлежности к той или иной Security Group</td>
        </tr>
        <tr>
            <td>FW-OUT-sgName (Security Group name)</td>
            <td><i>Третья</i> цепочка содержит наборы правил, относящиеся только к конкретной Security Group.</td>
        </tr>
    </tbody>
</table>

## Verdict statement {#verdict-statement}

<table>
    <thead>
        <tr>
            <th>Значение</th>
            <th>Описание</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>accept</td>
            <td class="text-justify">Терминирующие правило которое разрещает трафик, попавший под это условие, и завершает обработку пакета в текущей таблице.</td>
        </tr>
        <tr>
            <td>drop</td>
            <td class="text-justify">Терминирующие правило которое запрещает трафик, попавший под это условие, и завершает обработку пакета в текущей таблице.</td>
        </tr>
        <tr>
            <td>goto <i>chain</i></td>
            <td class="text-justify">Переход на другую цепочку в рамках указанного правила. После завершения правил в этой цепочке обратно в ту же цепочку трафик попасть не сможет.</td>
        </tr>
        <tr>
            <td>
                <nobr>jump <i>chain</i></nobr>
            </td>
            <td class="text-justify">Переход на другую цепочку в рамках указанного правила. После завершения правил в этой цепочке трафик вернется в исходную цепочку.</td>
        </tr>
    </tbody>
</table>

## Дескрипторы которые могут использоваться при указании типа ICMP. {#icmp-descriptors}

<table>
    <thead>
        <tr>
            <th>Дескриптор</th>
            <th>Значение</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>echo-reply</td>
            <td>0</td>
        </tr>
        <tr>
            <td>destination-unreachable</td>
            <td>3</td>
        </tr>
        <tr>
            <td>source-quench</td>
            <td>4</td>
        </tr>
        <tr>
            <td>redirect</td>
            <td>5</td>
        </tr>
        <tr>
            <td>echo-request</td>
            <td>8</td>
        </tr>
        <tr>
            <td>router-advertisement</td>
            <td>9</td>
        </tr>
        <tr>
            <td>router-solicitation</td>
            <td>10</td>
        </tr>
        <tr>
            <td>time-exceeded</td>
            <td>11</td>
        </tr>
        <tr>
            <td>parameter-problem</td>
            <td>12</td>
        </tr>
        <tr>
            <td>timestamp-request</td>
            <td>13</td>
        </tr>
        <tr>
            <td>timestamp-reply</td>
            <td>14</td>
        </tr>
        <tr>
            <td>info-request</td>
            <td>15</td>
        </tr>
        <tr>
            <td>info-reply</td>
            <td>16</td>
        </tr>
        <tr>
            <td>address-mask-request</td>
            <td>17</td>
        </tr>
        <tr>
            <td>address-mask-reply</td>
            <td>18</td>
        </tr>
    </tbody>
</table>
