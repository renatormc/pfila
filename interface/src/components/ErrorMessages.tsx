type Props = {
    errors?: string[],
    className?: string
}

const ErrorMessages = (props: Props) => {
    return <>
        {props.errors && <div className={props.className}>
            {props.errors.map((e, index) => {
                return <p className="text-sm text-red-400 italic" key={index}>{e}</p>
            })}
        </div>}
    </>
}

export default ErrorMessages