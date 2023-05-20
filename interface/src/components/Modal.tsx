import React, { useEffect, useRef } from 'react'

type Props = {
    show: boolean,
    onToggleShow?: (value: boolean) => void,
    children: JSX.Element,
    className?: string
}

function Modal(props: Props) {
    const ref = useRef<HTMLDivElement | null>(null)

    const onShow = () => {
        if (props.onToggleShow !== undefined) {
            props.onToggleShow(false)
        }
    }

    const handleClickOutside = (event: MouseEvent) => {
        if ((ref.current) && (!ref.current.contains(event.target as Node)) && (props.onToggleShow !== undefined)) {
            props.onToggleShow(false)
        }
    }

    useEffect(() => {
        if (props.show) {
            document.addEventListener('click', handleClickOutside, true)
        } else {
            document.removeEventListener('click', handleClickOutside, true)
        }
    }, [props.show])


    return (
        <>
            {
                props.show && <div tabIndex={-1} className="z-50 min-h-screen w-screen inset-0 bg-gray-700 fixed bg-opacity-10 flex flex-row justify-center items-center">
                    <div className={`relative ${props.className || ''}`}>
                        {props.onToggleShow && <button type="button" onClick={onShow} className="absolute top-3 right-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center" data-modal-toggle="popup-modal">
                            <svg className="w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z" clipRule="evenodd"></path></svg>
                        </button>}
                        {props.children}

                    </div>
                </div>
            }
        </>
    )
}

export default Modal