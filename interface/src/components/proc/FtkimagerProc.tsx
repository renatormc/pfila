import React, { useEffect, useState } from 'react'
import { FtkParams, ProcType } from '~/types/types'
import * as api from '~/services/api'

type Props = {
    params: FtkParams
}

const FtkimagerProc = ({ params }: Props) => {
    const [disks, setDisks] = useState<string[]>([])
    const [selectedDisk, setSelectedDisk] = useState("")

    const loadDisks = async () => {
        const res = await api.getResources<string>("/api/disks")
        setDisks(res)
    }

    useEffect(()=>{
        loadDisks()
    }, [])

    return <>
    <div className='flex flex-col gap-1 border'>
        {disks.map((disk, index)=>{
            return <div key={index} className={`cursor-pointer hover:bg-gray-200 p-2 ${selectedDisk == disk? 'bg-gray-200': ''}`}
            onClick={()=>{setSelectedDisk(disk)}}>{disk}</div>
        })}
    </div>
    </>
}

export default FtkimagerProc