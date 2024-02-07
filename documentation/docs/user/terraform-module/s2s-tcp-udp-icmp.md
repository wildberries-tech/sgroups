---
id: s2s-tcp-udp-icmp
---

# Sgroup to Sgroup (tcp|udp|icmp)

## Описание

Правило для взаимодействия с внешней Security Group.

## Параметры ресурса

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

## Пример использования

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


## Ограничения

<ul>
    <li>Необходимо указать минимум одно значение в sgroupSet.</li>
    <li class="text-justify">sgroupSet - длина одного значения из списка не должна превышать 256 символов. Значение должно начинаться и заканчиваться символами без пробелов.</li>
    <li>Может быть использован только один протокол для текущего правила (TCP или UDP).</li>
    <li class="text-justify">Значения `rules.s2s[].access.tcp/udp.ports_to[]` и `rules.s2s[].access.tcp/udp.ports_from[]` портов должно находиться в интервале от 1 до 65535. Если значение не будет указано то будет использоваться весь диапазон портов.</li>
    <li>Значения `rules.s2s[].access.tcp/udp.ports_from[]` портов не должно пересекаться в рамках текущего правила.</li>
    <li class="text-justify">`rules.s2s[].access.icmpIPv4/icmpIPv6.type[]` - значения в списке не должны повторяться. Значение должно быть цифрой от 0 до 255.</li>
</ul>