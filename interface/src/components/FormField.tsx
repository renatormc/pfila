import React from 'react'

type Props = {
    errors?: string[],
    children: JSX.Element | JSX.Element[],
    label?: string,
    className?: string,
    afterLabel?: JSX.Element
}

const FormField = ({ errors, children, label, className, afterLabel }: Props) => {
    return <div className={`flex flex-col ${className || ''} `}>
        <div className='flex items-center'>
            {label && <label className="m-label">{label}</label>}
            {afterLabel}
        </div>

        <div className={`${errors && 'bg-red-100 p-1 rounded-lg'}`}>
            {children}
        </div>
        {errors?.map((err, i) => {
            return <span className='text-red-500 text-sm italic'>{err}</span>
        })}
    </div>
}

export default FormField