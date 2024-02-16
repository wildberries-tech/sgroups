import React, { FC } from 'react'
import { TRestrictionItems } from '@site/src/customTypes/restrictions'

export const Restrictions: FC<{data: TRestrictionItems}> = ({data}) => {
    return (
        <ul>
            {data.map(item => (
                <li className="text-justify">{item}</li>
            ))}
        </ul>
    )
}