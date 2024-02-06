---
id: swarm-spec
---

# Terraform module

<div class="paragraph">
В папке `spec` находятся подкаталоги, которые служат описанием для Security Group (далее - `sgroup`). Для создания собственных `sgroup` необходимо создать вложенный подкаталог внутри `./spec`. Ожидается, что в этой иерархии будет как минимум три уровня подкаталогов. Каждый из этих подкаталогов описывает структуру окружения для конкретной `sgroup`, например: `./spec/superapp/dev/front/main.yaml`.
</div>

<div class="paragraph">
Правильный формат каталога для `sgroup` следующий: `<namespace>/<env>/<instance_name>/main.yaml`. Здесь `<namespace>` обозначает пространство имен, а `<instance_name>` - имя логической группы.
</div>

<div class="paragraph">
Для описания `sgroup`, в каждой конечной директории `./spec` создается файл в формате `*.yaml`, в котором описываются различные правила для `sgroup`.
</div>

## Security Group

### Описание

Security Group - это логическая группа для виртуального брандмауэра, которая включает набор инстансов или подсетей для фильтрации ingress/Egress правилами для сетевого трафика.

### Параметры ресурса

<table>
    <thead>
        <tr>
            <th>Название параметра</th>
            <th>Описание</th>
            <th>Значение по умолчанию</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>name</td>
            <td>Имя Security Group, соответствующей шаблону `/<namespace>/<env>/<instance_name>`</td>
            <td>""</td>
        </tr>
    </tbody>
</table>

### Пример использования

```yaml
name: "namespace/env/gitlab-runner"
```

### Ограничения

* name - обязательный параметр. Длина значения не должна превышать 256 символов. Значение должно начинаться и заканчиваться символами без пробелов.

## Список сетей (Networks)

### Описание

Networks - это список IP-адресов (CIDR), принадлежащих группе безопасности.

### Параметры ресурса

<table>
    <thead>
        <tr>
            <th>Название параметра</th>
            <th>Описание</th>
            <th>Значение по умолчанию</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>cidrs[]</td>
            <td>Список CIDR, связанных с Security Group</td>
            <td>[]</td>
        </tr>
    </tbody>
</table>

### Пример использования

```yaml
cidrs:
  - 10.160.152.130/32
  - 10.160.69.67/32
  - 10.20.56.4/32
```

### Ограничения

* Не должно быть пересечений в cidrs.

## Правила по умолчанию для Security Group (default rules)

### Описание

default rules - Структура, описывающая правила по умолчанию, для пакетов не соответствующих ни одному из установленных правил в цепочке.

### Параметры ресурса

<table>
    <thead>
        <tr>
            <th>Название параметра</th>
            <th>Описание</th>
            <th>Значение по умолчанию</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>default_rules</td>
            <td>Структура, описывающая правила по умолчанию, для пакетов не соответствующих ни одному из установленных правил в цепочке.</td>
            <td>{}</td>
        </tr>
        <tr>
            <td>default_rules.access</td>
            <td>Структура, описывающая взаимодействие с пакетами не соответствующими ни одному из установленных правил в таблице.</td>
            <td>{}</td>
        </tr>
        <tr>
            <td>default_rules.access.default.logs</td>
            <td>Включить/отключить логирование для пакетов, не соответствующих ни одному из установленных правил в цепочке.</td>
            <td>boolean</td>
        </tr>
        <tr>
            <td>default_rules.access.default.trace</td>
            <td>Включить/отключить трассировку для пакетов, не соответствующих ни одному из установленных правил в цепочке.</td>
            <td>boolean</td>
        </tr>
        <tr>
            <td>default_rules.access.default.action</td>
            <td>Определяет действие по умолчанию для пакетов, не соответствующих ни одному из установленных правил в цепочке. Допустимые значения: `ACCEPT`, `DROP`.</td>
            <td>`ACCEPT`</td>
        </tr>
        <tr>
            <td>default_rules.access.icmp</td>
            <td>Структура, описывающая взаимодействие с ICMP-трафиком по умолчанию. Для обработки ICMP-трафика добавляется соответствующее правило в начало цепочки.</td>
            <td>{}</td>
        </tr>
        <tr>
            <td>default_rules.access.icmp.logs</td>
            <td>Включить/отключить логирование ICMP-трафика.</td>
            <td>boolean</td>
        </tr>
        <tr>
            <td>default_rules.access.icmp.trace</td>
            <td>Включить/отключить трассировку ICMP-трафика.</td>
            <td>boolean</td>
        </tr>
        <tr>
            <td>default_rules.access.icmp.type[]</td>
            <td>Список, определяющий допустимые типы ICMP запросов.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>default_rules.access.icmp6</td>
            <td>Структура, описывающая взаимодействие с ICMPv6-трафиком по умолчанию. Для обработки ICMPv6-трафика добавляется соответствующее правило в начало цепочки.</td>
            <td>{}</td>
        </tr>
        <tr>
            <td>default_rules.access.icmp6.logs</td>
            <td>Включить/отключить логирование ICMPv6-трафика.</td>
            <td>boolean</td>
        </tr>
        <tr>
            <td>default_rules.access.icmp6.trace</td>
            <td>Включить/отключить трассировку ICMPv6-трафика.</td>
            <td>boolean</td>
        </tr>
        <tr>
            <td>default_rules.access.icmp6.type[]</td>
            <td>Список, определяющий допустимые типы ICMPv6 запросов.</td>
            <td>[]</td>
        </tr>
    </tbody>
</table>

### Пример использования

```yaml
default_rules:
  access:
    default:
      logs:   true
      trace:  true
      action: ACCEPT

    icmp:
      logs:  true
      trace: true
      type:  [0,8]

    icmp6:
      logs:  true
      trace: true
      type:  [0,8]
```

### Ограничения

* default_rules.access.default.action - обязательный параметр. Допустимые значения: `ACCEPT`, `DROP`.
* default_rules.access.icmp.type и default_rules.access.icmp6.type - значения в списке не должны повторяться. Значение должно быть цифрой от 0 до 255.

## Правила для Security Group для входящего и исходящего траффика

В рамках Security Group можно сформировать следующие типы взаимодействия:

* Ingress - трафик исходящий из инстанса в HBF правил описывающее разрешающее правило для входящего трафика инстанса. Для данного типа траффика можно сформировать следующие правила:
  * [cidrSet](#ingress-cidrset)
* Egress - трафик исходящий из инстанса в HBF правил описывающее разрешающее правило для исходящего трафика инстанса. Для данного типа траффика можно сформировать следующие правила:
  * [cidrSet](#egress-cidrset)
  * [fqdnSet](#egress-fqdnset)
* Security Group to Security Group (s2s) - правило для взаимодействия с внешней Security Group.
  * [sgroupSet](#security-group-to-security-group-s2s-sgroupset)

## Ingress - cidrSet {#ingress-cidrset}

### Описание

Правило для входящего траффика в указанную Security Group для взаимодействия с указанными списком CIDR.

### Параметры ресурса

<table>
    <thead>
        <tr>
            <th>Название параметра</th>
            <th>Описание</th>
            <th>Значение по умолчанию</th>
            <th></th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>rules</td>
            <td>Структура, содержащая описания сетевого взаимодействия (ingress, egress или s2s).</td>
            <td>{}</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[]</td>
            <td>Список, описывающий сетевое взаимодействие ingress (для входящего траффика). Для применения данных политик, добавляется соответствующее правило как на стороне инициатора.</td>
            <td>[]</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[].cidrSet[]</td>
            <td>Список, содержащий допустимые CIDR для ingress правила.</td>
            <td>[]</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[].log</td>
            <td>Включить/отключить логирование ingress трафика.</td>
            <td>boolean</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[].trace</td>
            <td>Включить/отключить трассировку ingress трафика .</td>
            <td>boolean</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[].access.tcp</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP либо с icmpIPv4.</td>
            <td>{}</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[].access.tcp.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[].access.tcp.ports_to[]</td>
            <td>Список destination портов, открытых для ingress трафика.</td>
            <td>[]</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[].access.tcp.ports_from[]</td>
            <td>Список source портов, открытых для ingress трафика.</td>
            <td>[]</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[].access.udp</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP.</td>
            <td>{}</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[].access.udp.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[].access.udp.ports_to[]</td>
            <td>Список destination портов, открытых для ingress трафика.</td>
            <td>[]</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[].access.udp.ports_from[]</td>
            <td>Список source портов, открытых для ingress трафика.</td>
            <td>[]</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.ingress[].access.icmpIPv4</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP либо с icmpIPv4.</td>
            <td>{}</td>
            <td>TODO</td>
        </tr>
        <tr>
            <td>rules.ingress[].access.icmpIPv4.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
            <td>TODO</td>
        </tr>
        <tr>
            <td>rules.ingress[].access.icmpIPv4.type[]</td>
            <td>Список, определяющий допустимые типы ICMP запросов.</td>
            <td>[]</td>
            <td>TODO</td>
        </tr>
        <tr>
            <td>rules.egress[].access.icmpIPv6</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6.</td>
            <td>{}</td>
            <td>TODO</td>
        </tr>
        <tr>
            <td>rules.egress[].access.icmpIPv6.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
            <td>TODO</td>
        </tr>
        <tr>
            <td>rules.egress[].access.icmpIPv6.type[]</td>
            <td>Список, определяющий допустимые типы ICMPv6 запросов.</td>
            <td>[]</td>
            <td>TODO</td>
        </tr>
    </tbody>
</table>

### Пример использования

```yaml
rules:
  ingress:
    - cidrSet:
        - 10.0.0.0/8
      logs:  true
      trace: true
      access:
        tcp:
          - description: ""
            ports_from:
              - 64231
            ports_to:
              - 443
              - 80
```

### Ограничения

* Необходимо указать минимум одно значение в cidrSet.
* Не должно быть пересечений в cidrSet.
* Может быть использован только один протокол для текущего правила (TCP или UDP)
* Значения rules.ingress[].tcp/udp.ports_to[] и rules.ingress[].tcp/udp.ports_from[] портов должно находиться в интервале от 1 до 65535. Если значение не будет указано то будет использоваться весь диапазон портов.
* Значения rules.ingress[].tcp/udp.ports_from[] портов не должно пересекаться в рамках текущего правила.
* **TODO:** rules.egress.access[].icmpIPv4/icmpIPv6.type[] - значения в списке не должны повторяться. Значение должно быть цифрой от 0 до 255.

## Egress - cidrSet {#egress-cidrset}

### Описание

Правило для исходящего траффика в указанную Security Group для взаимодействия с указанными списком CIDR.

### Параметры ресурса

<table>
    <thead>
        <tr>
            <th>Название параметра</th>
            <th>Описание</th>
            <th>Значение по умолчанию</th>
            <th></th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>rules</td>
            <td>Структура, содержащая описания сетевого взаимодействия (ingress, egress или s2s).</td>
            <td>{}</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[]</td>
            <td>Список, описывающий сетевое взаимодействие egress (для исходящего траффика). Для применения данных политик, добавляется соответствующее правило как на стороне инициатора.</td>
            <td>[]</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[].cidrSet[]</td>
            <td>Список, содержащий допустимые CIDR для egressправила.</td>
            <td>[]</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[].log</td>
            <td>Включить/отключить логирование egress трафика.</td>
            <td>boolean</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[].trace</td>
            <td>Включить/отключить трассировку egress трафика .</td>
            <td>boolean</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[].access.tcp</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6.</td>
            <td>{}</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[].access.tcp.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[].access.tcp.ports_to[]</td>
            <td>Список destination портов, открытых для egress трафика.</td>
            <td>[]</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[].access.tcp.ports_from[]</td>
            <td>Список source портов, открытых для egress трафика.</td>
            <td>[]</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[].access.udp</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6.</td>
            <td>{}</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[].access.udp.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[].access.udp.ports_to[]</td>
            <td>Список destination портов, открытых для egress трафика.</td>
            <td>[]</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[].access.udp.ports_from[]</td>
            <td>Список source портов, открытых для egress трафика.</td>
            <td>[]</td>
            <td></td>
        </tr>
        <tr>
            <td>rules.egress[].access.icmpIPv4</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6.</td>
            <td>{}</td>
            <td>TODO</td>
        </tr>
        <tr>
            <td>rules.egress[].access.icmpIPv4.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
            <td>TODO</td>
        </tr>
        <tr>
            <td>rules.egress[].access.icmpIPv4.type[]</td>
            <td>Список, определяющий допустимые типы ICMP запросов.</td>
            <td>[]</td>
            <td>TODO</td>
        </tr>
        <tr>
            <td>rules.egress[].access.icmpIPv6</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6.</td>
            <td>{}</td>
            <td>TODO</td>
        </tr>
        <tr>
            <td>rules.egress[].access.icmpIPv6.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
            <td>TODO</td>
        </tr>
        <tr>
            <td>rules.egress[].access.icmpIPv6.type[]</td>
            <td>Список, определяющий допустимые типы ICMPv6 запросов.</td>
            <td>[]</td>
            <td>TODO</td>
        </tr>
    </tbody>
</table>


### Пример использования

```yaml
rules:
  egress:
    - cidrSet:
        - 10.0.0.0/8
      logs:  true
      trace: true
      access:
        tcp:
          - description: ""
            ports_from:
              - 64231
            ports_to:
              - 443
              - 80
```

### Ограничения

* Необходимо указать минимум одно значение в cidrSet.
* Не должно быть пересечений в cidrSet.
* Может быть использован только один протокол для текущего правила (TCP или UDP).
* Значения rules.egress[].access.tcp/udp.ports_to[] и rules.egress[].access.tcp/udp.ports_from[] портов должно находиться в интервале от 1 до 65535. Если значение не будет указано то будет использоваться весь диапазон портов.
* Значения rules.egress[].access.tcp/udp.ports_from[] портов не должно пересекаться в рамках текущего правила.
* **TODO**\: rules.egress.access[].icmpIPv4/icmpIPv6.type[] - значения в списке не должны повторяться. Значение должно быть цифрой от 0 до 255.

## Egress - fqdnSet {#egress-fqdnset}

### Описание

Правило для исходящего траффика в указанную Security Group для взаимодействия с указанными списком FQDN.

### Параметры ресурса

<table>
    <thead>
        <tr>
            <th>Название параметра</th>
            <th>Описание</th>
            <th>Значение по умолчанию</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>rules</td>
            <td>Структура, содержащая описания сетевого взаимодействия (ingress, egress или s2s).</td>
            <td>{}</td>
        </tr>
        <tr>
            <td>rules.egress[]</td>
            <td>Список, описывающий сетевое взаимодействие egress (для исходящего траффика). Для применения данных политик, добавляется соответствующее правило как на стороне инициатора.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>rules.egress[].fqdnSet[]</td>
            <td>Список, содержащий fqdn для egressправила.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>rules.egress[].log</td>
            <td>Включить/отключить логирование egress трафика.</td>
            <td>boolean</td>
        </tr>
        <tr>
            <td>rules.egress[].trace</td>
            <td>Включить/отключить трассировку egress трафика .</td>
            <td>boolean</td>
        </tr>
        <tr>
            <td>rules.egress[].access.tcp</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP либо с типом UDP</td>
            <td>{}</td>
        </tr>
        <tr>
            <td>rules.egress[].access.tcp.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
        </tr>
        <tr>
            <td>rules.egress[].access.tcp.ports_to[]</td>
            <td>Список destination портов, открытых для egress трафика.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>rules.egress[].access.tcp.ports_from[]</td>
            <td>Список source портов, открытых для egress трафика.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>rules.egress[].access.udp</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP либо с типом UDP</td>
            <td>{}</td>
        </tr>
        <tr>
            <td>rules.egress[].access.udp.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
        </tr>
        <tr>
            <td>rules.egress[].access.udp.ports_to[]</td>
            <td>Список destination портов, открытых для egress трафика.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>rules.egress[].access.udp.ports_from[]</td>
            <td>Список source портов, открытых для egress трафика.</td>
            <td>[]</td>
        </tr>
    </tbody>
</table>

### Пример использования

```yaml
rules:
  egress:
    - fqdnSet:
        - zabbix-proxy.dl.wb.ru
        - z4-video-stream3.whs.wb.ru
      logs:  true
      trace: true
      access:
        tcp:
          - description: ""
            ports_from:
              - 64231
            ports_to:
              - 443
              - 80
```

### Ограничения

* Необходимо указать минимум одно значение в fqdnSet.
* fqdnSet - длина одного значения из списка не должна превышать 256 символов.
* Может быть использован только один протокол для текущего правила (TCP или UDP).
* Значения rules.egress[].access.tcp/udp.ports_to[] и rules.egress[].access.tcp/udp.ports_from[] портов должно находиться в интервале от 1 до 65535. Если значение не будет указано то будет использоваться весь диапазон портов.
* Значения rules.egress[].access.tcp/udp.ports_from[] портов не должно пересекаться в рамках текущего правила.

## Security Group to Security Group (s2s) - sgroupSet {#security-group-to-security-group-s2s-sgroupset}

### Описание

Правило для взаимодействия с внешней Security Group.

### Параметры ресурса

<table>
    <thead>
        <tr>
            <th>Название параметра</th>
            <th>Описание</th>
            <th>Значение по умолчанию</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>rules</td>
            <td>Структура, содержащая описания сетевого взаимодействия (ingress, egress или s2s).</td>
            <td>{}</td>
        </tr>
        <tr>
            <td>rules.s2s[]</td>
            <td>Список, описывающий сетевое взаимодействие текущей Security Group с внешней Security Group. Для применения данных политик, добавляется соответствующее правило как на стороне инициатора так и на стороне внешней Secyrity Group.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>rules.s2s[].sgroupSet[]</td>
            <td>Список, содержащий имена внешний Security Group.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>rules.s2s[].log</td>
            <td>Включить/отключить логирование для текущего правила.</td>
            <td>boolean</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.tcp</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP либо с типом UDP</td>
            <td>{}</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.tcp.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.tcp.ports_to[]</td>
            <td>Список destination портов, открытых для текущего правила.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.tcp.ports_from[]</td>
            <td>Список source портов, открытых для текущего правила.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.udp</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP либо с типом UDP</td>
            <td>{}</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.udp.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.udp.ports_to[]</td>
            <td>Список destination портов, открытых для текущего правила.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.udp.ports_from[]</td>
            <td>Список source портов, открытых для текущего правила.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.icmpIPv4</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6.</td>
            <td>{}</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.icmpIPv4.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.icmpIPv4.type[]</td>
            <td>Список, определяющий допустимые типы ICMP запросов.</td>
            <td>[]</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.icmpIPv6</td>
            <td>Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6.</td>
            <td>{}</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.icmpIPv6.description</td>
            <td>Формальное текстовое описание текущего правила.</td>
            <td>&quot;&quot;</td>
        </tr>
        <tr>
            <td>rules.s2s[].access.icmpIPv6.type[]</td>
            <td>Список, определяющий допустимые типы ICMPv6 запросов.</td>
            <td>[]</td>
        </tr>
    </tbody>
</table>

### Пример использования

```yaml
name: sg-example-from
rules:
  s2s:
    - sgroupSet:
        - sg-example-to
      logs:  true
      access:
        tcp:
          - description: ""
            ports_from:
              - 64231
            ports_to:
              - 443
              - 80
        udp:
          - description: ""
            ports_from:
              - 64231
            ports_to:
              - 443
              - 80
        icmpIPv4:
          - description: "" 
            Types:
                - 0
                - 8
        icmpIPv6:
          - description: "" 
            Types:
                - 0
                - 8
```


### Ограничения

* Необходимо указать минимум одно значение в sgroupSet.
* sgroupSet - длина одного значения из списка не должна превышать 256 символов. Значение должно начинаться и заканчиваться символами без пробелов.
* Может быть использован только один протокол для текущего правила (TCP или UDP).
* Значения rules.s2s[].access.tcp/udp.ports_to[] и rules.s2s[].access.tcp/udp.ports_from[] портов должно находиться в интервале от 1 до 65535. Если значение не будет указано то будет использоваться весь диапазон портов.
* Значения rules.s2s[].access.tcp/udp.ports_from[] портов не должно пересекаться в рамках текущего правила.
* rules.s2s[].access.icmpIPv4/icmpIPv6.type[] - значения в списке не должны повторяться. Значение должно быть цифрой от 0 до 255.