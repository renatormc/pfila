import React, { useEffect, useState } from "react";
import NavBar1 from '../components/NavBar'
import { DEFAULT_PROCESS, Process } from "~/types/types";
import * as api from "~/services/api"
import Modal from "~/components/Modal";

function ProcessesPage() {

    const [console, setConsole] = useState("asdfasdf")
    const [procs, setProcs] = useState<Process[]>([])
    const [editingProc, setEditingProc] = useState<Process | null>(null)

    const load = async () => {
        const res = await api.getResources<Process>("/api/proc")
        setProcs(res)
    }

    useEffect(() => {
        load()
    }, [])

    return (
        <div>
            <NavBar1 onNew={() => { setEditingProc(DEFAULT_PROCESS) }} />
            <div className="px-2">
                <p className="text-xl">Processos</p>

                <div className="relative overflow-x-auto">
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
                                        ${proc.finish}
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
            <Modal className="bg-white w-full h-1/2 p-2 rounded-sm" show={editingProc != null} onToggleShow={()=>{setEditingProc(null)}}>
                <div className="flex flex-col">
                    <div className="flex flex-col">
                        <label htmlFor="">teste</label>
                        <input className="m-input" />
                    </div>

                </div>

            </Modal>
        </div>
    );
}

export default ProcessesPage