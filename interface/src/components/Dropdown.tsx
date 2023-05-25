import { prependOnceListener } from "process"
import React, { useMemo, ReactElement, PropsWithChildren, Children, useState, useRef, useEffect } from "react"

type DropdownItemProps = {
    text: string,
    onClick?: () => void,
    disabled?: boolean
}

export function DropdownlItem(props: DropdownItemProps) {

    // return <>{props.children}</>
    return <></>
};

type Props = {
    children: JSX.Element[],
    label?: string,
    className?: string
}

const DropdownComponent = (props: Props) => {

    const [show, setShow] = useState(false)
    const ref = useRef(null)

    const items = useMemo(
        () => Children.map(props.children as ReactElement<PropsWithChildren<DropdownItemProps>>[], (item) => item.props),
        [props.children],
    )

    const handleClickOutside = (event: MouseEvent) => {
        if (ref.current != event.target) {
            setShow(false)
        }
    }

    const onClickItem = (index: number) => {
        const item = items[index]
        if ((item.onClick) && (!item.disabled)){
            item.onClick()
        }
        setShow(false)
    }

    useEffect(() => {
        if (show) {
            document.addEventListener('click', handleClickOutside, true)
        } else {
            document.removeEventListener('click', handleClickOutside, true)
        }
    }, [show])


    return <div className={`relative shadow-sm  rounded-md ${props.className || ''}`}>
        <button className={`text-dourado-500 bg-azul-600  hover:bg-azul-500 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-4 py-2.5 text-right inline-flex justify-center items-center w-full`}
            type="button"
            onClick={() => { setShow(!show) }}
        >
            {props.label || ''}
            <svg className="ml-2 w-4 h-4" aria-hidden="true" fill="none" stroke="currentColor" viewBox="0 0 24 24" xmlns="http://www.w3.org/2000/svg">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M19 9l-7 7-7-7"></path>
            </svg>
        </button>
        <div ref={ref} className={`${!show ? 'hidden' : ''} z-10  bg-gray-50 rounded divide-y divide-gray-100 shadow  absolute right-0 w-fit min-w-full`}>
            <ul className="py-1 text-sm " aria-labelledby="dropdownDefault">
                {items.map((item, index) => {
                    return (
                        <li key={index} >
                            <span className={`whitespace-nowrap block py-2 px-4 ${item.disabled ? 'text-azul-200' : 'text-azul-500 hover:bg-gray-100 cursor-pointer'}  text-left`} onClick={() => { onClickItem(index) }}>{item.text}</span>
                        </li>
                    )
                })}
            </ul>
        </div>
    </div>
}

DropdownComponent.displayName = 'Dropdown.Group';
DropdownlItem.displayName = 'Dropdown.Item';
export const Dropdown = { Group: DropdownComponent, Item: DropdownlItem };