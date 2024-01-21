import { ProcType } from '~/types/types'
import Button from './Button'

type Props = {
    onNew?: (ptype: ProcType) => void,
}


function NavBar1({ onNew }: Props) {

    return <>
        <nav className="border-gray-200 px-10 py-2 rounded fixed top-0 w-screen bg-azul-500">
            <div className="flex flex-wrap justify-between items-center mx-auto ">
                <a href="https://flowbite.com/" className="flex items-center text-dourado-500">
                    <span className="self-center text-xl font-semibold whitespace-nowrap">PFila</span>
                </a>
                <button data-collapse-toggle="navbar-default" type="button" className="inline-flex items-center p-2 ml-3 text-sm text-gray-500 rounded-lg md:hidden hover:bg-gray-100 focus:outline-none focus:ring-2 focus:ring-gray-200 " aria-controls="navbar-default" aria-expanded="false">
                    <span className="sr-only">Open main menu</span>
                    <svg className="w-6 h-6" aria-hidden="true" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fillRule="evenodd" d="M3 5a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 10a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1zM3 15a1 1 0 011-1h12a1 1 0 110 2H4a1 1 0 01-1-1z" clipRule="evenodd"></path></svg>
                </button>
                <div className="hidden w-full md:block md:w-auto">
                    <ul className="flex flex-col px-4 mt-4 rounded-lg border border-gray-100 md:flex-row md:space-x-8 md:mt-0 md:text-sm md:font-medium md:border-0 ">
                        <li>
                            <div className='flex gap-4'>
                                <Button label="Novo IPED" onClick={() => { onNew?.('iped') }} variant='blue' />
                                <Button label="Nova Imagem" onClick={() => { onNew?.('ftkimager') }} variant='blue' />
                                <Button label="Novo comando" onClick={() => { onNew?.('freecmd') }} variant='blue' />
                            </div>

                        </li>
                    </ul>
                </div>
            </div>
        </nav>

    </>

}

export default NavBar1