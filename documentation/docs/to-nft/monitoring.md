---
id: monitoring
---

# Мониторинг

<div class="paragraph">
Информация о состоянии агента на различных узлах основывается по метрикам [из данной панели Grafana](https://grafana.wildberries.ru/d/wzJE8qKIk/hbf-agent?orgId=1&refresh=5s&from=now-12h&to=now). На ней также настроены уведомления, связанные с недоступностью или некорректной работой агента. Уведомления приходят в закрытый канал в Rocket.chat (#HBF_ALERTS). Таблица ниже описывает возможные для агента уведомления и причины их возникновения.
</div>

<table>
    <thead>
        <tr>
            <th>Уведомление</th>
            <th>Причина уведомления</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td>Статус health-check равен 0</td>
            <td>Метрика `healthcheck` возвращает значение 0 или по данной метрике отсутствуют значения</td>
        </tr>
        <tr>
            <td>Потребление CPU стало критически низким (менее 0.150 cores).</td>
            <td>Формула расчет потребления CPU связанная с метрикой `process_cpu_seconds_total` возвращает значение ниже 0.150 или по данной метрике отсутствуют значения</td>
        </tr>
        <tr>
            <td>Потребление RAM подозрительно выросло (более 65 Mb).</td>
            <td>Формула расчет потребления RAM связанная с метрикой `process_resident_memory_bytes` возвращает значение выше 65000000 или по данной метрике отсутствуют значения</td>
        </tr>
    </tbody>
</table>

## Метрики

<div class="paragraph">
Метрики по умолчанию доступны через порт **9660**. Пользователь может настроить порт и другие параметры в разделе "**telemetry**" файла конфигурации ( по умолчанию файл конфигурации находится в */opt/swarm/etc/to-nft* ). Для проверки доступности метрик рекомендуется выполнить команду `curl 0.0.0.0:9660/metrics`.
</div>

![metrics](../../static/img/nft.monitoring.aa.png)

### Описание метрик

Зеленым цветом <span class="green">выделены ключевые метрики</span>.

<table>
    <thead>
        <tr>
            <th>Название метрики</th>
            <th>Тип метрики</th>
            <th>Описание</th>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td class="green">agent_applied_configs</td>
            <td>counter</td>
            <td>Количество успешно примененных конфигураций</td>
        </tr>
        <tr>
            <td>go_gc_duration_seconds</td>
            <td>summary</td>
            <td>Сводка длительности паузы циклов сборки мусора</td>
        </tr>
        <tr>
            <td>go_gc_duration_seconds_sum</td>
            <td>counter</td>
            <td>Сумма длительности паузы циклов сборки мусора</td>
        </tr>
        <tr>
            <td>go_gc_duration_seconds_count</td>
            <td>counter</td>
            <td>Количество циклов сборки мусора</td>
        </tr>
        <tr>
            <td></td>
            <td></td>
            <td></td>
        </tr>
        <tr>
            <td>go_goroutines</td>
            <td>gauge</td>
            <td>Количество горутин, существующих в данный момент</td>
        </tr>
        <tr>
            <td class="green">go_info</td>
            <td>gauge</td>
            <td>Информация о среде выполнения Go</td>
        </tr>
        <tr>
            <td>go_memstats_alloc_bytes</td>
            <td>gauge</td>
            <td>Количество выделенных и еще используемых байтов</td>
        </tr>
        <tr>
            <td>go_memstats_alloc_bytes_total</td>
            <td>counter</td>
            <td>Общее количество выделенных байтов, даже если они освобождены</td>
        </tr>
        <tr>
            <td>go_memstats_buck_hash_sys_bytes</td>
            <td>gauge</td>
            <td>Количество байтов, используемых хэш-таблицей профилирования</td>
        </tr>
        <tr>
            <td>go_memstats_frees_total</td>
            <td>counter</td>
            <td>Общее количество освобождений</td>
        </tr>
        <tr>
            <td>go_memstats_gc_sys_bytes</td>
            <td>gauge</td>
            <td>Количество байтов, используемых для метаданных системы сборки мусора</td>
        </tr>
        <tr>
            <td>go_memstats_heap_alloc_bytes</td>
            <td>gauge</td>
            <td>Количество выделенных байтов кучи и еще используемых</td>
        </tr>
        <tr>
            <td>go_memstats_heap_idle_bytes</td>
            <td>gauge</td>
            <td>Количество байтов кучи в ожидании использования</td>
        </tr>
        <tr>
            <td>go_memstats_heap_inuse_bytes</td>
            <td>gauge</td>
            <td>Количество используемых байтов кучи</td>
        </tr>
        <tr>
            <td>go_memstats_heap_objects</td>
            <td>gauge</td>
            <td>Количество выделенных объектов</td>
        </tr>
        <tr>
            <td>go_memstats_heap_released_bytes</td>
            <td>gauge</td>
            <td>Количество байтов кучи, освобожденных в ОС</td>
        </tr>
        <tr>
            <td>go_memstats_heap_sys_bytes</td>
            <td>gauge</td>
            <td>Количество байтов кучи, полученных от системы</td>
        </tr>
        <tr>
            <td>go_memstats_last_gc_time_seconds</td>
            <td>gauge</td>
            <td>Количество секунд с 1970 года последней сборки мусора</td>
        </tr>
        <tr>
            <td>go_memstats_lookups_total</td>
            <td>counter</td>
            <td>Общее количество запросов указателей</td>
        </tr>
        <tr>
            <td>go_memstats_mallocs_total</td>
            <td>counter</td>
            <td>Общее количество выделений памяти</td>
        </tr>
        <tr>
            <td>go_memstats_mcache_inuse_bytes</td>
            <td>gauge</td>
            <td>Количество байтов, используемых структурами mcache</td>
        </tr>
        <tr>
            <td>go_memstats_mcache_sys_bytes</td>
            <td>gauge</td>
            <td>Количество байтов, используемых структурами mcache, полученных от системы</td>
        </tr>
        <tr>
            <td>go_memstats_mspan_inuse_bytes</td>
            <td>gauge</td>
            <td>Количество байтов, используемых структурами mspan</td>
        </tr>
        <tr>
            <td>go_memstats_mspan_sys_bytes</td>
            <td>gauge</td>
            <td>Количество байтов, используемых структурами mspan, полученных от системы</td>
        </tr>
        <tr>
            <td>go_memstats_next_gc_bytes</td>
            <td>gauge</td>
            <td>Количество байтов кучи, когда произойдет следующая сборка мусора</td>
        </tr>
        <tr>
            <td>go_memstats_other_sys_bytes</td>
            <td>gauge</td>
            <td>Количество байтов, используемых для других системных выделений</td>
        </tr>
        <tr>
            <td>go_memstats_stack_inuse_bytes</td>
            <td>gauge</td>
            <td>Количество байтов, используемых аллокатором стека</td>
        </tr>
        <tr>
            <td>go_memstats_stack_sys_bytes</td>
            <td>gauge</td>
            <td>Количество байтов, полученных от системы для аллокатора стека</td>
        </tr>
        <tr>
            <td>go_memstats_sys_bytes</td>
            <td>gauge</td>
            <td>Количество байтов, полученных от системы</td>
        </tr>
        <tr>
            <td>go_threads</td>
            <td>gauge</td>
            <td>Количество созданных потоков ОС</td>
        </tr>
        <tr>
            <td class="green">healthcheck</td>
            <td>gauge</td>
            <td>Индикатор проверки состояния процесса (0 или 1)</td>
        </tr>
        <tr>
            <td class="green">process_cpu_seconds_total</td>
            <td>counter</td>
            <td>Общее количество времени процессора, затраченного в секундах</td>
        </tr>
        <tr>
            <td>process_max_fds</td>
            <td>gauge</td>
            <td>Максимальное количество открытых файловых дескрипторов</td>
        </tr>
        <tr>
            <td>process_open_fds</td>
            <td>gauge</td>
            <td>Количество открытых файловых дескрипторов</td>
        </tr>
        <tr>
            <td class="green">process_resident_memory_bytes</td>
            <td>gauge</td>
            <td>Размер резидентной памяти в байтах</td>
        </tr>
        <tr>
            <td>process_start_time_seconds</td>
            <td>gauge</td>
            <td>Время запуска процесса с начала эпохи Unix в секундах</td>
        </tr>
        <tr>
            <td>process_virtual_memory_bytes</td>
            <td>gauge</td>
            <td>Размер виртуальной памяти в байтах</td>
        </tr>
        <tr>
            <td>process_virtual_memory_max_bytes</td>
            <td>gauge</td>
            <td>Максимальное количество виртуальной памяти в байтах</td>
        </tr>
        <tr>
            <td class="green">promhttp_metric_handler_errors_total</td>
            <td>counter</td>
            <td>Общее количество http ошибок, выявленных обработчиком</td>
        </tr>
    </tbody>
</table>