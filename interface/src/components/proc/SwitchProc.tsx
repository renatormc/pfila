import React from 'react'
import { ProcParams, ProcType } from '~/types/types'
import IpedProc from './IpedProc'
import FtkimagerProc from './FtkimagerProc'

type Props = {
    ptype: ProcType,
    params: ProcParams
}

const SwitchProc = ({ ptype, params }: Props) => {
    switch (ptype) {
        case "iped":
            return <IpedProc params={params}/>
        case "ftkimager":
            return <FtkimagerProc params={params}/>

        default:
            return <p>{ptype}</p>
    }
}

export default SwitchProc