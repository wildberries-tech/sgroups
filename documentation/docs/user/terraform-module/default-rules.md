---
id: default-rules
---

# Правила по умолчанию для Security Group (default rules)

## Описание

default rules - Структура, описывающая правила по умолчанию, для пакетов не соответствующих ни одному из установленных правил в цепочке.

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

# Пример использования

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

## Ограничения

<ul>
    <li>`default_rules.access.default.action` - обязательный параметр. Допустимые значения: `ACCEPT`, `DROP`.</li>
    <li class="text-justify">`default_rules.access.icmp.type` и `default_rules.access.icmp6.type` - значения в списке не должны повторяться. Значение должно быть цифрой от 0 до 255.</li>
</ul>