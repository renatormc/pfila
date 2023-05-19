import React from 'react'
import { ErrorsType, FtkParams, IpedParams, ProcParams, ProcType } from '~/types/types'
import IpedProc from './IpedProc'
import FtkimagerProc from './FtkimagerProc'

type Props = {
    ptype: ProcType,
    params: ProcParams,
    setParams: (pars: ProcParams) => void,
    errors: ErrorsType,
    loadingParams: boolean,
    setLoadingParams: (v: boolean) => void
}

const SwitchProc = ({ ptype, params, setParams, errors, loadingParams, setLoadingParams }: Props) => {
    if (!params) {
        return <></>
    }
    switch (ptype) {
        case "iped":
            return <IpedProc params={params as IpedParams}
                setParams={setParams}
                errors={errors} 
                loadingParams={loadingParams}
                setLoadingParams={setLoadingParams}/>
        case "ftkimager":
            return <FtkimagerProc params={params as FtkParams}
                setParams={setParams}
                errors={errors} 
                loadingParams={loadingParams}
                setLoadingParams={setLoadingParams}/>
        default:
            return <p>{ptype}</p>
    }
}

export default SwitchProc