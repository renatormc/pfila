import React, { useEffect, useState } from "react";
import NavBar1 from '../../components/NavBar'
import { ErrorsType, ProcType, Process, getDefaultProcess } from "~/types/types";
import * as api from "~/services/api"
import Modal from "~/components/Modal";
import SwitchProc from "~/components/proc/SwitchProc";
import Button from "~/components/Button";
import Input from "~/components/Input";
import { Dropdown } from "~/components/Dropdown";
import FormField from "~/components/FormField";
import { AxiosError } from "axios";
import WaitingModal from "~/components/WaitingModal";



function ProcessesPage() {
    const [console, setConsole] = useState("asdfasdf")
    const [procs, setProcs] = useState<Process[]>([])
    const [editingProc, setEditingProc] = useState<Process | null>(null)
    const [selectedProcIndex, setSelectedProcIndex] = useState(-1)
    const [errors, setErrors] = useState<ErrorsType>({})
    const [loadingParams, setLoadingParams] = useState(false)

    const load = async () => {
        const res = await api.getProcessess()
        setProcs(res)
    }

    const updateField = <K extends keyof Process>(field: K, value: Process[K]) => {
        if (editingProc) {
            const copy = { ...editingProc }
            copy[field] = value
            setEditingProc(copy)
        }
    }

    const onNew = (ptype: ProcType) => {

        const p = getDefaultProcess(ptype)
        setErrors({})
        setEditingProc(p)
    }

    const save = async () => {
        if (editingProc) {
            try {
                const res = await api.createProcess(editingProc)
                setEditingProc(null)
                load()
            } catch (error) {
                const err = error as AxiosError
                if (err.response?.status == 422) {
                    setErrors(err.response.data as ErrorsType)
                }
            }
        }
    }

    const queueProcess = async (index: number) => {
        const res = await api.queueProcess(procs[index].id)
        const copy = [...procs]
        copy[index] = res
        setProcs(copy)
    }

    const deleteProcess = async (index: number) => {
        await api.deleteProcess(procs[index].id)
    }

    const editProcess = async (index: number) => {
        setErrors({})
        setEditingProc(procs[index])
    }

    const cancelProcess = async (index: number) => {
        const res = await api.stopProcess(procs[index].id)
        const copy = [...procs]
        copy[index] = res
        setProcs(copy)
    }

    const showConsole = async (index: number) => {
        const res = await api.procConsole(procs[index].id)
        setConsole(res)
    }


    useEffect(() => {
        load()
    }, [])

    useEffect(() => {
        if (selectedProcIndex > -1) {
            showConsole(selectedProcIndex)
        } else {
            setConsole("")
        }

    }, [selectedProcIndex])


    return (
        <div className="h-screen pt-16 ">
            <NavBar1 onNew={onNew} />
            <div className="px-6 h-full ">
                <p className="text-xl">Processos</p>

                <div className="">
                    <table className="w-full text-sm text-left text-gray-500 ">
                        <thead className="text-xs text-gray-700 uppercase bg-gray-50">
                            <tr>
                                <th scope="col" className="px-6 py-3">
                                    ID
                                </th>
                                <th scope="col" className="px-6 py-3">
                                    Nome
                                </th>
                                <th scope="col" className="px-6 py-3">
                                    Usuário
                                </th>
                                <th scope="col" className="px-6 py-3">
                                    Início
                                </th>
                                <th scope="col" className="px-6 py-3">
                                    Entrada na fila
                                </th>
                                <th scope="col" className="px-6 py-3">
                                    Fim
                                </th>
                                <th scope="col" className="px-6 py-3">
                                    Tipo
                                </th>
                                <th scope="col" className="px-6 py-3">
                                    Status
                                </th>
                                <th scope="col" className="px-6 py-3">

                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            {procs.map((proc, index) => {
                                return <tr className={`bg-white border-b hover:bg-gray-100 ${selectedProcIndex == index ? 'bg-gray-300' : 'bg' + proc.status}`}
                                    key={index} >
                                    <td className="px-6 py-4">
                                        {proc.id}
                                    </td>
                                    <th scope="row" className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap " >
                                        <span className="cursor-pointer hover:text-gray-500" onClick={() => { setSelectedProcIndex(index) }}>{proc.name}</span>

                                    </th>
                                    <td className="px-6 py-4">
                                        {proc.user}
                                    </td>
                                    <td className="px-6 py-4">
                                        {proc.start}
                                    </td>
                                    <td className="px-6 py-4">
                                        {proc.start_waiting}
                                    </td>
                                    <td className="px-6 py-4">
                                        {proc.finish}
                                    </td>
                                    <td className="px-6 py-4">
                                        {proc.type}
                                    </td>
                                    <td className="px-6 py-4">
                                        {proc.status}
                                    </td>
                                    <td className="px-6 py-4">
                                        <Dropdown.Group label="OP">
                                            <Dropdown.Item text="Editar" onClick={() => { editProcess(index) }} />
                                            <Dropdown.Item text="Colocar na fila" onClick={() => { queueProcess(index) }} />
                                            <Dropdown.Item text="Deletar" onClick={() => { deleteProcess(index) }} />
                                            <Dropdown.Item text="Cancelar" onClick={() => { cancelProcess(index) }} />
                                        </Dropdown.Group>
                                    </td>
                                </tr>
                            })}
                        </tbody>
                    </table>
                </div>
            </div>
            <div className="fixed bottom-0 h-80 w-full">
                <p className="text-lg mt-6 ml-2">Console</p>
                <div className="w-full h-full bg-gray-600 text-gray-100 p-3  mt-1 " dangerouslySetInnerHTML={{ __html: console }}>

                </div>
            </div>
            <Modal className="bg-white w-full max-w-2xl h-fit p-5 rounded-sm pt-8" show={editingProc != null} onToggleShow={() => { setEditingProc(null) }}>
                <div className="flex flex-col">
                    <div className="flex flex-col gap-2">
                        <p className="mb-2 text-blue-600 text-xl">Novo processo</p>
                        <FormField label='Nome do processo' errors={errors.name}>
                            <Input className="w-full" value={editingProc?.name} onChange={(v) => { updateField('name', v) }} />
                        </FormField>
                        <FormField label='Usuário' errors={errors.user}>
                            <Input className="w-full" value={editingProc?.user} onChange={(v) => { updateField('user', v) }} />
                        </FormField>

                        {editingProc && <SwitchProc ptype={editingProc.type}
                            params={editingProc.params}
                            setParams={(pars) => { updateField('params', pars) }}
                            errors={errors}
                            loadingParams={loadingParams}
                            setLoadingParams={setLoadingParams} />}
                        <Button label="Gravar" variant="blue" onClick={save} />
                    </div>
                </div>
            </Modal>
            <WaitingModal message={loadingParams ? 'carregando': ''}/>
        </div>
    );
}

export default ProcessesPage