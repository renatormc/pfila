import React from 'react'

type Props = {
    errors?: string[],
    children: JSX.Element | JSX.Element[],
    label?: string,
    className?: string
}

const FormField = ({ errors, children, label, className }: Props) => {
    return <div className={`flex flex-col ${className || ''} `}>
        {label && <label className="m-label">{label}</label>}
        <div className={`${errors && 'bg-red-100 p-1 rounded-lg'}`}>
            {children}
        </div>
        {errors?.map((err, i)=>{
            return <span className='text-red-500 text-sm italic'>{err}</span>
        })}
    </div>
}

export default FormField