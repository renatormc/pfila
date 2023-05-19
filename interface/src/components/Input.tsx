import React from 'react'

type Props = {
    onChange?: (v: string) => void,
    value?: string,
    className?: string,
    placeholder?: string,
    autoFocus?: boolean
}

const Input = ({ value, onChange, className, placeholder, autoFocus }: Props) => {
    return <input className={`bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 p-2.5 ${className || ''}`} value={value}
        onChange={(e) => { onChange?.(e.target.value) }} placeholder={placeholder} autoFocus={autoFocus}/>
}

export default Input