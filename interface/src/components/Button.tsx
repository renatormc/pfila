import React from 'react'

type Props = {
    label: string,
    onClick?: () => void,
    variant: 'blue'
}

const Button = ({ label, onClick, variant }: Props) => {
    switch (variant) {
        case "blue":
            return <button className="block py-2 pr-4 pl-3 text-gray-100 rounded hover:bg-blue-600 bg-blue-400"
                onClick={onClick}>{label}</button>
    }
}

export default Button