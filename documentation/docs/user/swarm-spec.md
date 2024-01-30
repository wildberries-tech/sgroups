---
id: swarm-spec
---

# Terraform module

В папке `spec` находятся подкаталоги, которые служат описанием для Security Group (далее - `sgroup`). Для создания собственных `sgroup` необходимо создать вложенный подкаталог внутри `./spec`. Ожидается, что в этой иерархии будет как минимум три уровня подкаталогов. Каждый из этих подкаталогов описывает структуру окружения для конкретной `sgroup`, например: `./spec/superapp/dev/front/main.yaml`.

Правильный формат каталога для `sgroup` следующий: `<namespace>/<env>/<instance_name>/main.yaml`. Здесь `<namespace>` обозначает пространство имен, а `<instance_name>` - имя логической группы.

Для описания `sgroup`, в каждой конечной директории `./spec` создается файл в формате `*.yaml`, в котором описываются различные правила для `sgroup`.

## Security Group

### Описание

Security Group - это логическая группа для виртуального брандмауэра, которая включает набор инстансов или подсетей для фильтрации ingress/Egress правилами для сетевого трафика.

### Параметры ресурса

| Название параметра | Описание | Значение по умолчанию |
| --- | --- | --- |
| name | Имя Security Group, соответствующей шаблону /\<namespace\>/\<env\>/\<instance_name\>. | "" |

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

| Название параметра | Описание | Значение по умолчанию |
| --- | --- | --- |
| cidrs[] | Список CIDR, связанных с Security Group | [] |

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

| Название параметра | Описание | Значение по умолчанию |
| --- | --- | --- |
| default_rules | Структура, описывающая правила по умолчанию, для пакетов не соответствующих ни одному из установленных правил в цепочке. | {} |
| default_rules.access | Структура, описывающая взаимодействие с пакетами не соответствующими ни одному из установленных правил в таблице. | {} |
| default_rules.access.default.logs | Включить/отключить логирование для пакетов, не соответствующих ни одному из установленных правил в цепочке. | boolean |
| default_rules.access.default.trace | Включить/отключить трассировку для пакетов, не соответствующих ни одному из установленных правил в цепочке. | boolean |
| default_rules.access.default.action | Определяет действие по умолчанию для пакетов, не соответствующих ни одному из установленных правил в цепочке. Допустимые значения: `ACCEPT`, `DROP`. | `ACCEPT` |
| default_rules.access.icmp | Структура, описывающая взаимодействие с ICMP-трафиком по умолчанию. Для обработки ICMP-трафика добавляется соответствующее правило в начало цепочки. | {} |
| default_rules.access.icmp.logs | Включить/отключить логирование ICMP-трафика. | boolean |
| default_rules.access.icmp.trace | Включить/отключить трассировку ICMP-трафика. | boolean |
| default_rules.access.icmp.type[] | Список, определяющий допустимые типы ICMP запросов. | [] |
| default_rules.access.icmp6 | Структура, описывающая взаимодействие с ICMPv6-трафиком по умолчанию. Для обработки ICMPv6-трафика добавляется соответствующее правило в начало цепочки. | {} |
| default_rules.access.icmp6.logs | Включить/отключить логирование ICMPv6-трафика. | boolean |
| default_rules.access.icmp6.trace | Включить/отключить трассировку ICMPv6-трафика. | boolean |
| default_rules.access.icmp6.type[] | Список, определяющий допустимые типы ICMPv6 запросов. | [] |

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

| Название параметра | Описание | Значение по умолчанию |  |
| --- | --- | --- | --- |
| rules | Структура, содержащая описания сетевого взаимодействия (ingress, egress или s2s). | {} |  |
| rules.ingress[] | Список, описывающий сетевое взаимодействие ingress (для входящего траффика). Для применения данных политик, добавляется соответствующее правило как на стороне инициатора. | [] |  |
| rules.ingress[].cidrSet[] | Список, содержащий допустимые CIDR для ingress правила. | [] |  |
| rules.ingress[].log | Включить/отключить логирование ingress трафика. | boolean |  |
| rules.ingress[].trace | Включить/отключить трассировку ingress трафика . | boolean |  |
| rules.ingress[].access.tcp | Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP либо с icmpIPv4. | {} |  |
| rules.ingress[].access.tcp.description | Формальное текстовое описание текущего правила. | "" |  |
| rules.ingress[].access.tcp.ports_to[] | Список destination портов, открытых для ingress трафика. | [] |  |
| rules.ingress[].access.tcp.ports_from[] | Список source портов, открытых для ingress трафика. | [] |  |
| rules.ingress[].access.udp | Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP. | {} |  |
| rules.ingress[].access.udp.description | Формальное текстовое описание текущего правила. | "" |  |
| rules.ingress[].access.udp.ports_to[] | Список destination портов, открытых для ingress трафика. | [] |  |
| rules.ingress[].access.udp.ports_from[] | Список source портов, открытых для ingress трафика. | [] |  |
| rules.ingress[].access.icmpIPv4 | Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP либо с icmpIPv4. | {} | TODO |
| rules.ingress[].access.icmpIPv4.description | Формальное текстовое описание текущего правила. | "" | TODO |
| rules.ingress[].access.icmpIPv4.type[] | Список, определяющий допустимые типы ICMP запросов. | [] | TODO |
| rules.egress[].access.icmpIPv6 | Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6. | {} | TODO |
| rules.egress[].access.icmpIPv6.description | Формальное текстовое описание текущего правила. | "" | TODO |
| rules.egress[].access.icmpIPv6.type[] | Список, определяющий допустимые типы ICMPv6 запросов. | [] | TODO |

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

| Название параметра | Описание | Значение по умолчанию |  |
| --- | --- | --- | --- |
| rules | Структура, содержащая описания сетевого взаимодействия (ingress, egress или s2s). | {} |  |
| rules.egress[] | Список, описывающий сетевое взаимодействие egress (для исходящего траффика). Для применения данных политик, добавляется соответствующее правило как на стороне инициатора. | [] |  |
| rules.egress[].cidrSet[] | Список, содержащий допустимые CIDR для egressправила. | [] |  |
| rules.egress[].log | Включить/отключить логирование egress трафика. | boolean |  |
| rules.egress[].trace | Включить/отключить трассировку egress трафика . | boolean |  |
| rules.egress[].access.tcp | Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6. | {} |  |
| rules.egress[].access.tcp.description | Формальное текстовое описание текущего правила. | "" |  |
| rules.egress[].access.tcp.ports_to[] | Список destination портов, открытых для egress трафика. | [] |  |
| rules.egress[].access.tcp.ports_from[] | Список source портов, открытых для egress трафика. | [] |  |
| rules.egress[].access.udp | Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6. | {} |  |
| rules.egress[].access.udp.description | Формальное текстовое описание текущего правила. | "" |  |
| rules.egress[].access.udp.ports_to[] | Список destination портов, открытых для egress трафика. | [] |  |
| rules.egress[].access.udp.ports_from[] | Список source портов, открытых для egress трафика. | [] |  |
| rules.egress[].access.icmpIPv4 | Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6. | {} | TODO |
| rules.egress[].access.icmpIPv4.description | Формальное текстовое описание текущего правила. | "" | TODO |
| rules.egress[].access.icmpIPv4.type[] | Список, определяющий допустимые типы ICMP запросов. | [] | TODO |
| rules.egress[].access.icmpIPv6 | Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6. | {} | TODO |
| rules.egress[].access.icmpIPv6.description | Формальное текстовое описание текущего правила. | "" | TODO |
| rules.egress[].access.icmpIPv6.type[] | Список, определяющий допустимые типы ICMPv6 запросов. | [] | TODO |

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

| Название параметра | Описание | Значение по умолчанию |
| --- | --- | --- |
| rules | Структура, содержащая описания сетевого взаимодействия (ingress, egress или s2s). | {} |
| rules.egress[] | Список, описывающий сетевое взаимодействие egress (для исходящего траффика). Для применения данных политик, добавляется соответствующее правило как на стороне инициатора. | [] |
| rules.egress[].fqdnSet[] | Список, содержащий fqdn для egressправила. | [] |
| rules.egress[].log | Включить/отключить логирование egress трафика. | boolean |
| rules.egress[].trace | Включить/отключить трассировку egress трафика . | boolean |
| rules.egress[].access.tcp | Тип протокола для данного правила, можно использовать либо блок с типом TCP либо с типом UDP | {} |
| rules.egress[].access.tcp.description | Формальное текстовое описание текущего правила. | "" |
| rules.egress[].access.tcp.ports_to[] | Список destination портов, открытых для egress трафика. | [] |
| rules.egress[].access.tcp.ports_from[] | Список source портов, открытых для egress трафика. | [] |
| rules.egress[].access.udp | Тип протокола для данного правила, можно использовать либо блок с типом TCP либо с типом UDP | {} |
| rules.egress[].access.udp.description | Формальное текстовое описание текущего правила. | "" |
| rules.egress[].access.udp.ports_to[] | Список destination портов, открытых для egress трафика. | [] |
| rules.egress[].access.udp.ports_from[] | Список source портов, открытых для egress трафика. | [] |

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

| Название параметра | Описание | Значение по умолчанию |
| --- | --- | --- |
| rules | Структура, содержащая описания сетевого взаимодействия (ingress, egress или s2s). | {} |
| rules.s2s[] | Список, описывающий сетевое взаимодействие текущей Security Group с внешней Security Group. Для применения данных политик, добавляется соответствующее правило как на стороне инициатора так и на стороне внешней Secyrity Group. | [] |
| rules.s2s[].sgroupSet[] | Список, содержащий имена внешний Security Group. | [] |
| rules.s2s[].log | Включить/отключить логирование для текущего правила. | boolean |
| rules.s2s[].access.tcp | Тип протокола для данного правила, можно использовать либо блок с типом TCP либо с типом UDP | {} |
| rules.s2s[].access.tcp.description | Формальное текстовое описание текущего правила. | "" |
| rules.s2s[].access.tcp.ports_to[] | Список destination портов, открытых для текущего правила. | [] |
| rules.s2s[].access.tcp.ports_from[] | Список source портов, открытых для текущего правила. | [] |
| rules.s2s[].access.udp | Тип протокола для данного правила, можно использовать либо блок с типом TCP либо с типом UDP | {} |
| rules.s2s[].access.udp.description | Формальное текстовое описание текущего правила. | "" |
| rules.s2s[].access.udp.ports_to[] | Список destination портов, открытых для текущего правила. | [] |
| rules.s2s[].access.udp.ports_from[] | Список source портов, открытых для текущего правила. | [] |
| rules.s2s[].access.icmpIPv4 | Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6. | {} |
| rules.s2s[].access.icmpIPv4.description | Формальное текстовое описание текущего правила. | "" |
| rules.s2s[].access.icmpIPv4.type[] | Список, определяющий допустимые типы ICMP запросов. | [] |
| rules.s2s[].access.icmpIPv6 | Тип протокола для данного правила, можно использовать либо блок с типом TCP, либо с UDP, либо с icmpIPv4 либо с icmpIPv6. | {} |
| rules.s2s[].access.icmpIPv6.description | Формальное текстовое описание текущего правила. | "" |
| rules.s2s[].access.icmpIPv6.type[] | Список, определяющий допустимые типы ICMPv6 запросов. | [] |

### Пример использования

```yaml
rules:
  s2s:
    - sgroupSet:
        - namespace/env/gitlab-runner
      logs:  true
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

* Необходимо указать минимум одно значение в sgroupSet.
* sgroupSet - длина одного значения из списка не должна превышать 256 символов. Значение должно начинаться и заканчиваться символами без пробелов.
* Может быть использован только один протокол для текущего правила (TCP или UDP).
* Значения rules.s2s[].access.tcp/udp.ports_to[] и rules.s2s[].access.tcp/udp.ports_from[] портов должно находиться в интервале от 1 до 65535. Если значение не будет указано то будет использоваться весь диапазон портов.
* Значения rules.s2s[].access.tcp/udp.ports_from[] портов не должно пересекаться в рамках текущего правила.
* rules.s2s[].access.icmpIPv4/icmpIPv6.type[] - значения в списке не должны повторяться. Значение должно быть цифрой от 0 до 255.