import React, { useEffect, useRef, useState } from 'react'
import { Process } from '~/types/types'
import * as api from "~/services/api"
import Checkbox from '~/components/Checkbox'

type Props = {
    proc: Process | null,
    className?: string
}

const Console = ({ proc, className }: Props) => {
    const [text, setText] = useState("")
    const ref = useRef<HTMLDivElement>(null)
    const [autoUpdate, setAutoUpdate] = useState(true)
    const load = async () => {
        if (proc) {
            const res = await api.procConsole(proc.id)
            setText(res)
        }
    }

    useEffect(() => {
        if(proc){
            load()
            if (autoUpdate) {
                load()
                const timer = setInterval(async () => {
                    load()
                }, 3000);
                return () => {
                    clearInterval(timer)
                }
            }
        }else{
            setText("")
        }
        
    }, [proc, autoUpdate])

    useEffect(() => {
        if (ref.current) {
            // ref.current.scrollIntoView({ behavior: 'smooth', block: 'end' })
            ref.current.scrollTop = ref.current.scrollHeight
        }

    }, [text])

    return <div className={`${className || ''}`}>
        <div className='flex justify-end cursor-pointer hover:text-gray-500 gap-2 items-center'>
            <span>{proc?.name}</span>
            <div className='grow'></div>
            <Checkbox label='Atualizar automático' value={autoUpdate} onChange={setAutoUpdate}/>
            <i className="fa-solid fa-arrows-rotate" onClick={load}></i>
        </div>
        <div className={`bg-gray-700 text-gray-100 p-3  mt-1 overflow-auto h-full`} ref={ref}>
            <pre>{text}</pre>
        </div>
    </div>
}

export default Console