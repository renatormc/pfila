import React, { useEffect, useState } from 'react'
import { FtkParams, ProcType } from '~/types/types'
import * as api from '~/services/api'
import Checkbox from '../Checkbox'
import Select from '../Select'
import FormField from '../FormField'
import Input from '../Input'

type Props = {
    params: FtkParams,
    setParams: (pars: FtkParams) => void
}

const FtkimagerProc = ({ params, setParams }: Props) => {
    const [disks, setDisks] = useState<string[]>([])
    const [selectedDisk, setSelectedDisk] = useState("")

    const loadDisks = async () => {
        const res = await api.getResources<string>("/api/disks")
        setDisks(res)
    }

    const updateField = <K extends keyof FtkParams>(field: K, value: FtkParams[K]) => {
        const copy = { ...params }
        copy[field] = value
        setParams(copy)
    }


    useEffect(() => {
        loadDisks()
    }, [])

    return <>
        <FormField label='Disco'>
            <div className='flex flex-col gap-1 border'>
                {disks.map((disk, index) => {
                    return <div key={index} className={`cursor-pointer hover:bg-gray-200 p-2 ${selectedDisk == disk ? 'bg-gray-200' : ''}`}
                        onClick={() => { setSelectedDisk(disk) }}>{disk}</div>
                })}
            </div>
        </FormField>
        <FormField label='Destino'>
            <Input className='w-full' onChange={(v) => { updateField('destination', v) }} 
            value={params?.destination} placeholder='Endereço do arquivo de imagem sem extensão'/>
        </FormField>
        <FormField label='Formato'>
            <Select className='mb-3' onChange={(v) => { updateField('format', v) }} value={params?.format}
                options={[
                    { value: 'e01', text: 'e01' },
                    { value: 'raw', text: 'raw' },
                ]} />
        </FormField>
        <Checkbox className='mb-3' value={params?.verify} onChange={(v) => { updateField('verify', v) }} label='Verificar' />
    </>
}

export default FtkimagerProc