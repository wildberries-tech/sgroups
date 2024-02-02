---
id: main
---

# FW-IN

### Описание

Правило перехода в цепочку FW-IN-sgName с проверкой что трафик предназначен для указанной Security Group.

### Параметры

<table>
    <thead>
        <tr>
            <th>Шаблон параметра</th>
            <th>Структура</th>
            <th>Значение</th>
            <th>Описание</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>\{RuleType\}</td>
            <td>ip</td>
            <td>Наименование IPSet в котором описаны сети в Security Group</td>
            <td>Значение типа string, не должно содержать в себе пробелов</td>
        </tr>
        <tr>
            <td>\{DstSgroup\}</td>
            <td>
                <nobr>`dst @${IPSet(sgName)}`</nobr>
            </td>
            <td></td>
            <td></td>
        </tr>
        <tr>
            <td>\{sgName\}</td>
            <td></td>
            <td></td>
            <td>Название Security Group</td>
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
            <td>\{Verdict\}</td>
            <td>
                <nobr>`goto ${chainName}`</nobr>
            </td>
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
chain FW-IN {  
  # **********
  ${RuleType} ${DstSgroup} ${Counter} ${Verdict} FW-IN-${sgName}
  # **********
}
```

### Пример использования

```hcl
chain FW-IN {  
  # **********
  ip daddr @familyName-sgName_example counter packets 0 bytes 0 goto FW-IN-sgName_example
  # **********
}
```