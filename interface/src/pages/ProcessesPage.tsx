import React, { useEffect, useState } from "react";
import NavBar1 from '../components/NavBar'
import { ProcType, Process, getDefaultProcess } from "~/types/types";
import * as api from "~/services/api"
import Modal from "~/components/Modal";
import SwitchProc from "~/components/proc/SwitchProc";
import Button from "~/components/Button";
import Input from "~/components/Input";
import { Dropdown } from "~/components/Dropdown";



function ProcessesPage() {
    const [console, setConsole] = useState("asdfasdf")
    const [procs, setProcs] = useState<Process[]>([])
    const [editingProc, setEditingProc] = useState<Process | null>(null)

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
        setEditingProc(p)
    }

    const save = async () => {
        if (editingProc) {
            const res = api.createProcess(editingProc)
            setEditingProc(null)
        }
    }

    useEffect(() => {
        load()
    }, [])

    return (
        <div className="h-full bg-yellow-200">
            <NavBar1 onNew={onNew} />
            <div className="px-6 h-full grow">
                <p className="text-xl">Processos</p>

                <div className="h-full ">
                    <table className="w-full text-sm text-left text-gray-500 ">
                        <thead className="text-xs text-gray-700 uppercase bg-gray-50">
                            <tr>
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
                                return <tr className="bg-white border-b dark:bg-gray-800 " key={index}>
                                    <th scope="row" className="px-6 py-4 font-medium text-gray-900 whitespace-nowrap ">
                                        {proc.name}
                                    </th>
                                    <td className="px-6 py-4">
                                        {proc.user}
                                    </td>
                                    <td className="px-6 py-4">
                                        {proc.start}
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
                                            <Dropdown.Item text="Item 1"/>
                                            <Dropdown.Item text="Item 2"/>
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
                <div className="w-full h-full bg-gray-600 text-gray-100 p-3  mt-1 " >
                    {console}
                </div>
            </div>
            <Modal className="bg-white w-full max-w-2xl h-fit p-5 rounded-sm pt-8" show={editingProc != null} onToggleShow={() => { setEditingProc(null) }}>
                <div className="flex flex-col">
                    <div className="flex flex-col gap-2">
                        <p className="mb-2 text-blue-600 text-xl">Novo processo</p>
                        <div className="flex flex-col">
                            <label className="m-label">Nome do processo</label>
                            <Input value={editingProc?.name} onChange={(v) => { updateField('name', v) }} />
                        </div>
                        <div className="flex flex-col">
                            <label className="m-label">Usuário</label>
                            <Input value={editingProc?.user} onChange={(v) => { updateField('user', v) }} />
                        </div>

                        {editingProc && <SwitchProc ptype={editingProc.type} params={editingProc.params} setParams={(pars) => { updateField('params', pars) }} />}
                        <Button label="Gravar" variant="blue" onClick={save} />
                    </div>
                </div>
            </Modal>
        </div>
    );
}

export default ProcessesPage