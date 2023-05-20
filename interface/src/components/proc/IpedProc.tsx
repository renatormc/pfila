import React, { useEffect, useState } from 'react'
import { ErrorsType, IpedParams } from '~/types/types'
import Input from '../Input'
import * as api from '~/services/api'
import FormField from '../FormField'
import Select from '../Select'
import Checkbox from '../Checkbox'

type Props = {
    params: IpedParams,
    setParams: (pars: IpedParams) => void,
    errors: ErrorsType,
    loadingParams: boolean,
    setLoadingParams: (v: boolean) => void
}

const IpedProc = ({ params, setParams, errors, setLoadingParams }: Props) => {

    const [profiels, setProfiles] = useState<string[]>([])

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

    const loadIpedProfiles = async () => {
        try {
            setLoadingParams(true)
            const res = await api.getIpdeProfiles()
            setProfiles(res)
        } finally {
            setLoadingParams(false)
        }
    }

    useEffect(() => {
        loadIpedProfiles()
    }, [])

    return <>
        <FormField label='Pasta de saída' errors={errors?.destination}>
            <Input className='w-full' value={params?.destination} onChange={(v) => { updateField('destination', v) }} />
        </FormField>
        <FormField label='Fontes' errors={errors?.sources}
            afterLabel={<i className="fa-solid fa-plus cursor-pointer hover:text-gray-500 ml-2" onClick={addSource}></i>}>
            {params.sources.map((src, index) => {
                return <div className="flex gap-3 w-full items-center mb-2" key={index}>
                    <Input className='w-full' value={src} onChange={(v) => { updateSource(index, v) }} />
                    <i className="fa-solid fa-minus text-red-600 cursor-pointer hover:text-gray-500 " onClick={() => { removeSource(index) }}></i>
                </div>
            })}
        </FormField>

        {/* <div className="flex flex-col">
            <div className="flex items-center gap-4">
                <label htmlFor="">Fontes</label> <i className="fa-solid fa-plus cursor-pointer hover:text-gray-500" onClick={addSource}></i>
            </div>

            {params.sources.map((src, index) => {
                return <div className="flex gap-3 w-full items-center mb-2" key={index}>
                    <Input className='w-full' value={src} onChange={(v) => { updateSource(index, v) }} />
                    <i className="fa-solid fa-minus text-red-600 cursor-pointer hover:text-gray-500 " onClick={() => { removeSource(index) }}></i>
                </div>
            })}

        </div> */}
        <FormField label='Perfil' errors={errors?.profile}>
            <Select options={profiels.map(p => { return { value: p, text: p } })} value={params?.profile} onChange={(v) => { updateField('profile', v) }} />
        </FormField>
        <Checkbox className='mt-3' value={params?.portable} onChange={(v) => { updateField('portable', v) }} label='Portable' />
    </>
}

export default IpedProc