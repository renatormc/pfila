import React from 'react'

type Option = {
    value: string,
    text: string
}


type Props = {
    onChange?: (v: string) => void,
    value?: string,
    className?: string,
    options: Option[]
}


const Select = ({ value, onChange, className, options }: Props) => {
    return <select className={`bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 ${className || ''}`}
        onChange={(e) => { onChange?.(e.target.value) }}
        value={value}>
        {options.map((op, index) => {
            return <option key={index} value={op.value}>{op.text}</option>
        })}
    </select>
}

export default Select