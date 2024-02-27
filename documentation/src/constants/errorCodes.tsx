import React from "react";
import { TErrorCodes } from "../customTypes/errorCodes";

export const ERRORCODES: TErrorCodes = {

    apiTypo: {
        summary: "Ошибка в запросе",
        body:
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
    },
    
    networkExist: {
        summary: "Сеть с таким названием уже существует",
        body:
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
    },

    networkNotExist: {
        summary: "Добавление несуществующей сети",
        body:
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
    },

    networkTypo: {
        summary: "Некорректное значение поля Networks",
        body:
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
    },

    sgExist: {
        summary: "Security Group с таким названием уже существует",
        body:
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
    },
    
    sgNotExist: {
        summary: "Security Group с таким названием не существует",
        body:
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
    },

    sgTypo: {
        summary: "Некорректное значение поля SG",
        body:
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
    },

    CIDRIntersection: {
        summary: "Пересечение значений CIDR",
        body:
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
    },

    defaultActionTypo: {
        summary: "Некорректное значение поля defaultAction",
        body:
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
    },

    ICMP: {
        summary: "Некорректное значение кода ICMP",
        body:
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
    },

    IPv: {
        summary: "Некорректное значение поля IPv",
        body:
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
    },

    sPortsIntersection: {
        summary: "Пересечение значений source портов",
        body:
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
    },

    portsTypo: {
        summary: "Некорректное значение порта",
        body:
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
    },

    transportTypo: {
        summary: "Некорректное значение протокола",
        body:
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
    },

    trafficTypo: {
        summary: "Некорректное значение поля направления траффика",
        body:
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
    },

    syncopTypo: {
        summary: "Некорректное значение поля определяющее действие с данными из запроса",
        body:
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
    }

}