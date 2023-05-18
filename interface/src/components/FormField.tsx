import React from 'react'

type Props = {
    error?: string,
    children: JSX.Element | JSX.Element[],
    label?: string,
    className?: string
}

const FormField = ({ error, children, label, className }: Props) => {
    return <div className={`flex flex-col ${className || ''} `}>
        {label && <label className="m-label">{label}</label>}
        <div className={`${error && 'border-red-200 border-2'}`}>
            {children}
        </div>

        {error && <span className='text-red-500 text-sm italic'>{error}</span>}
    </div>
}

export default FormField