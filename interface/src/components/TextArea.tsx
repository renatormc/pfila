import React from 'react'

type Props = {
    onChange?: (v: string) => void,
    value?: string,
    className?: string,
    placeholder?: string,
    autoFocus?: boolean,
    rows?: number
}

const TextArea = ({ value, onChange, className, placeholder, autoFocus, rows }: Props) => {
    return <textarea className={`bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 p-2.5 ${className || ''}`} value={value}
        onChange={(e) => { onChange?.(e.target.value) }} placeholder={placeholder} autoFocus={autoFocus} rows={rows} style={{ whiteSpace: 'pre-line' }}/>
}

export default TextArea