import React from 'react'

type Props = {
    label: string,
    onClick?: () => void,
    variant: 'blue'
}

const Button = ({ label, onClick, variant }: Props) => {
    switch (variant) {
        case "blue":
            return <button className="block py-2 pr-4 pl-3 text-dourado-400 rounded hover:bg-azul-500 bg-azul-400"
                onClick={onClick}>{label}</button>
    }
}

export default Button