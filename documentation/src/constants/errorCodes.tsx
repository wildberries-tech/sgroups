import React from "react";
import { TerrorCodes } from "../customTypes/errorCodes";

export const ERRORCODES: TerrorCodes = {

    apiTypo: [
        <details>
            <summary>Ошибка в запросе</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>404</td>
                <td>5</td>
                <td>NOT_FOUND</td>
                <td>Ошибка в запросе</td>
            </tr>
        </table>
        </details>
    ],
    
    networkExist: [
        <details>
            <summary>Сеть с таким названием уже существует</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>13</td>
                <td>INTERNAL</td>
                <td>Сеть с таким названием уже существует</td>
            </tr>
        </table>
        </details>
    ],

    networkNotExist: [
        <details>
            <summary>Добавление несуществующей сети</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>13</td>
                <td>INTERNAL</td>
                <td>Добавление несуществующей сети</td>
            </tr>
        </table>
        </details>
    ],

    networkTypo: [
        <details>
            <summary>Некорректное значение поля Networks</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>3</td>
                <td>INVALID_ARGUMENT</td>
                <td>Некорректное значение поля Networks</td>
            </tr>
        </table>
        </details>
    ],

    sgExist: [
        <details>
            <summary>Security Group с таким названием уже существует</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>13</td>
                <td>INTERNAL</td>
                <td>Security Group с таким названием уже существует</td>
            </tr>
        </table>
        </details>
    ],
    
    sgNotExist: [
        <details>
            <summary>Security Group с таким названием не существует</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>13</td>
                <td>INTERNAL</td>
                <td>Security Group с таким названием не существует</td>
            </tr>
        </table>
        </details>
    ],

    sgTypo: [
        <details>
            <summary>Некорректное значение поля SG</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>3</td>
                <td>INVALID_ARGUMENT</td>
                <td>Некорректное значение поля SG</td>
            </tr>
        </table>
        </details>
    ],

    CIDRIntersection: [
        <details>
            <summary>Пересечение значений CIDR</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>3</td>
                <td>INVALID_ARGUMENT</td>
                <td>Пересечение значений CIDR</td>
            </tr>
        </table>
        </details>
    ],

    defaultActionTypo: [
        <details>
            <summary>Некорректное значение поля defaultAction</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>3</td>
                <td>INVALID_ARGUMENT</td>
                <td>Некорректное значение поля defaultAction</td>
            </tr>
        </table>
        </details>
    ],

    ICMP: [
        <details>
            <summary>Некорректное значение кода ICMP</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>3</td>
                <td>INVALID_ARGUMENT</td>
                <td>Некорректное значение кода ICMP</td>
            </tr>
        </table>
        </details>
    ],

    IPv: [
        <details>
            <summary>Некорректное значение поля IPv</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>3</td>
                <td>INVALID_ARGUMENT</td>
                <td>Некорректное значение поля IPv</td>
            </tr>
        </table>
        </details>
    ],

    sPortsIntersection: [
        <details>
            <summary>Пересечение значений source портов</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>13</td>
                <td>INTERNAL</td>
                <td>Пересечение значений source портов</td>
            </tr>
        </table>
        </details>
    ],

    portsTypo: [
        <details>
            <summary>Некорректное значение порта</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>3</td>
                <td>INVALID_ARGUMENT</td>
                <td>Некорректное значение порта</td>
            </tr>
        </table>
        </details>
    ],

    transportTypo: [
        <details>
            <summary>Некорректное значение протокола</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>3</td>
                <td>INVALID_ARGUMENT</td>
                <td>Некорректное значение протокола</td>
            </tr>
        </table>
        </details>
    ],

    trafficTypo: [
        <details>
            <summary>Некорректное значение поля направления траффика</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>3</td>
                <td>INVALID_ARGUMENT</td>
                <td>Некорректное значение поля направления траффика</td>
            </tr>
        </table>
        </details>
    ],

    syncopTypo: [
        <details>
            <summary>Некорректное значение поля определяющее действие с данными из запроса</summary>
        <table>
            <tr>
                <td>HTTP status code</td>
                <td>gRPC number</td>
                <td>gRPC code</td>
                <td>Описание</td>
            </tr>
            <tr>
                <td>500</td>
                <td>3</td>
                <td>INVALID_ARGUMENT</td>
                <td>Некорректное значение поля определяющее действие с данными из запроса</td>
            </tr>
        </table>
        </details>
    ]

}