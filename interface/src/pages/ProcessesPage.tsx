import React, { useEffect, useState } from "react";
import NavBar1 from '../components/NavBar'

function ProcessesPage() {

    const [console, setConsole] = useState("asdfasdf")

    return (
        <div>
            <NavBar1 />
            <div className="px-5">
                <p className="text-xl">Processos</p>
                <div className="w-full bg-gray-600 text-gray-100 p-3 rounded-md" >
                    {console}
                </div>
            </div>
        </div>
    );
}

export default ProcessesPage