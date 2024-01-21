import React, { useEffect, useState } from 'react'
import { ErrorsType, FreecmdParams, FtkParams, ProcType } from '~/types/types'
import * as api from '~/services/api'
import Checkbox from '../Checkbox'
import Select from '../Select'
import FormField from '../FormField'
import Input from '../Input'

type Props = {
    params: FreecmdParams,
    setParams: (pars: FreecmdParams) => void,
    errors: ErrorsType,
    loadingParams: boolean,
    setLoadingParams: (v: boolean) => void
}

const FreecmdProc = ({ params, setParams, errors, loadingParams, setLoadingParams }: Props) => {
    const [disks, setDisks] = useState<string[]>([])
    // const [selectedDisk, setSelectedDisk] = useState("")

    const loadDisks = async () => {
        try {
            const res = await api.getResources<string>("/api/disks")
            setDisks(res)
        } finally {
            setLoadingParams(false)
        }

    }

    const updateField = <K extends keyof FreecmdParams>(field: K, value: FreecmdParams[K]) => {
        const copy = { ...params }
        copy[field] = value
        setParams(copy)
    }


    useEffect(() => {
        setLoadingParams(true)
        loadDisks()
    }, [])

    return <>

        <FormField label='Destino' errors={errors?.destination}>
            <Input className='w-full' onChange={(v) => { updateField('cmd', v) }}
                value={params?.cmd} placeholder='Command to be executed' />
        </FormField>

    </>
}

export default FreecmdProc