import React, { FC } from 'react'
import { TTerminology } from '@site/src/customTypes/terminology'

export const Terminology: FC<{data: TTerminology[]}> = ({data}) => {
    return (
        <>
            {data.map((term, index) => (
                <div key={index} className="text-justify">
                    <b>{term.name}</b>
                    {term.comment && <i>({term.comment})</i>}
                    {" - "}
                    {term.definition}
                    {term.link && <a href={term.link}> Подробнее...</a>}
                    <br /><br />
                </div>
            ))}
        </>
    )
}