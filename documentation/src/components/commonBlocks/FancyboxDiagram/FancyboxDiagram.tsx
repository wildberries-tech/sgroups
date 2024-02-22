import React, { FC, useRef, useEffect, PropsWithChildren } from 'react'
import { TransformWrapper, TransformComponent } from 'react-zoom-pan-pinch'
import { Fancybox as NativeFancybox } from '@fancyapps/ui'
import '@fancyapps/ui/dist/fancybox/fancybox.css'
import { OptionsType } from '@fancyapps/ui/types/Fancybox/options'
import { Fancybox } from '@site/src/components/commonBlocks/Fancybox'

interface Props {
  options?: Partial<OptionsType>
  delegate?: string
}

export const FancyboxDiagram: FC<PropsWithChildren<Props>> = props => {
  const containerRef = useRef(null)

  useEffect(() => {
    const container = containerRef.current

    const delegate = props.delegate || '[data-fancybox]'
    const options = props.options || {}

    NativeFancybox.bind(container, delegate, options)

    return () => {
      NativeFancybox.unbind(container)
      NativeFancybox.close()
    }
  })

  return (
    <>
      <div id="dialog-content" style={{ display: 'none', minWidth: '80vw', background: '#000' }}>
        <div style={{ display: 'flex', justifyContent: 'center' }}>
          <TransformWrapper>
            <TransformComponent>
              <div style={{ minWidth: '80vw', width: '100%', height: '100%', margin: '0 auto' }}>{props.children}</div>
            </TransformComponent>
          </TransformWrapper>
        </div>
      </div>
      <Fancybox>
        <a data-fancybox="gallery" data-src="#dialog-content">
          {props.children}
        </a>
      </Fancybox>
    </>
  )
}
