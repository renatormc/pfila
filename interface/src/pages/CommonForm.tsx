import React from 'react'
import FormField from '~/components/FormField'
import Input from '~/components/Input'
import { ErrorsType, Process } from '~/types/types'

type Props = {
    updateField: <K extends keyof Process>(field: K, value: Process[K]) => void,
    proc: Process,
    errors: ErrorsType
}

const CommonForm = ({ updateField, proc, errors }: Props) => {
    return <>
        <FormField label='Nome do processo' errors={errors.name}>
            <Input value={proc.name} onChange={(v) => { updateField('name', v) }} />
        </FormField>
        <FormField label='Nome do processo' errors={errors.user}>
            <Input value={proc.user} onChange={(v) => { updateField('user', v) }} />
        </FormField>
    </>
}

export default CommonForm