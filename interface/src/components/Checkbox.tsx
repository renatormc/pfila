import React from 'react'

type Props = {
    label?: string,
    onChange?: (v: boolean) => void,
    value?: boolean,
    className?: string
}

const Checkbox = ({ label, onChange, value, className }: Props) => {
    return <div className={`cursor-pointer text-xl flex gap-2 items-center ${className || ''}`} onClick={() => { onChange?.(!value) }}>
        {value ? <i className="fa-regular fa-square-check" /> : <i className="fa-regular fa-square" />}
        <label className='text-lg cursor-pointer'> {label}</label>
    </div>
}

export default Checkbox