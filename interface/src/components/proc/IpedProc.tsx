import React, { useEffect } from 'react'
import { IpedParams } from '~/types/types'
import Input from '../Input'
import Checkbox from '../Checkbox'

type Props = {
    params: IpedParams,
    setParams: (pars: IpedParams) => void
}

const IpedProc = ({ params, setParams }: Props) => {

    const updateField = <K extends keyof IpedParams>(field: K, value: IpedParams[K]) => {
        const copy = { ...params }
        copy[field] = value
        setParams(copy)
    }

    const updateSource = (index: number, value: string) => {
        const copy = { ...params }
        copy.sources[index] = value
        setParams(copy)
    }

    const addSource = () => {
        const copy = { ...params }
        copy.sources.push("")
        setParams(copy)
    }

    const removeSource = (index: number) => {
        const copy = { ...params }
        copy.sources.splice(index, 1)
        setParams(copy)
    }

    return <>
        <div className="flex flex-col">
            <label className="m-label">Pasta de saída</label>
            <Input value={params?.destination} onChange={(v) => { updateField('destination', v) }} />
        </div>
        <div className="flex flex-col">
            <div className="flex items-center gap-4">
                <label htmlFor="">Fontes</label> <i className="fa-solid fa-plus cursor-pointer hover:text-gray-500" onClick={addSource}></i>
            </div>

            {params.sources.map((src, index) => {
                return <div className="flex gap-3 w-full items-center mb-2" key={index}>
                    <Input className='w-full' value={src} onChange={(v) => { updateSource(index, v) }} />
                    <i className="fa-solid fa-minus text-red-600 cursor-pointer hover:text-gray-500 " onClick={() => { removeSource(index) }}></i>
                </div>
            })}
            
        </div>
    </>
}

export default IpedProc