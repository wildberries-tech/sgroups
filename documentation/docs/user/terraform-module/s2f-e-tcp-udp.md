---
id: s2f-e-tcp-udp
---

# Sgroup to FQDN (e: tcp|udp)

## Описание

Правило для исходящего траффика в указанную Security Group для взаимодействия с указанными списком FQDN.

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

## Пример использования

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

## Ограничения

<ul>
    <li>Необходимо указать минимум одно значение в fqdnSet.</li>
    <li>fqdnSet - длина одного значения из списка не должна превышать 256 символов.</li>
    <li>Может быть использован только один протокол для текущего правила (TCP или UDP).</li>
    <li class="text-justify">Значения `rules.egress[].access.tcp/udp.ports_to[]` и `rules.egress[].access.tcp/udp.ports_from[]` портов должно находиться в интервале от 1 до 65535 Если значение не будет указано то будет использоваться весь диапазон портов.</li>
    <li class="text-justify">Значения `rules.egress[].access.tcp/udp.ports_from[]` портов не должно пересекаться в рамках текущего правила.</li>
</ul>