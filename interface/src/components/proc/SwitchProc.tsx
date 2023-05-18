import React from 'react'
import { FtkParams, IpedParams, ProcParams, ProcType } from '~/types/types'
import IpedProc from './IpedProc'
import FtkimagerProc from './FtkimagerProc'

type Props = {
    ptype: ProcType,
    params: ProcParams,
    setParams: (pars: ProcParams) => void
}

const SwitchProc = ({ ptype, params, setParams }: Props) => {
    if(!params){
        return <></>
    }
    switch (ptype) {
        case "iped":
            return <IpedProc params={params as IpedParams} setParams={setParams}/>
        case "ftkimager":
            return <FtkimagerProc params={params as FtkParams} setParams={setParams}/>
        default:
            return <p>{ptype}</p>
    }
}

export default SwitchProc