import { HTMLAttributes, ReactNode } from "react";

export interface ButtonProps extends HTMLAttributes<HTMLButtonElement> {
    children: ReactNode,
    variant?: "filled" | "outlined" | "text"
}

export default function Button({
    children,
    variant = "filled",
    ...rest
}: ButtonProps) {
    let varinatStyle = ""
    if (variant === "filled") {
        varinatStyle = "b-none bg-primary text-white hover:bg-secondary"
    }
    if (variant === "outlined") {
        varinatStyle = "b-1 border-gray-400 border-solid bg-bg hover:bg-hover text-primary"
    }
    if (variant === "text") {
        varinatStyle = "b-none bg-transparent text-primary hover:text-secondary hover:bg-hover"
    }

    return (
        <button
            className={"h-9 rounded flex flex-col justify-center items-center px-5 ease-in-out transition-all duration-300 font-semibold " + varinatStyle}
            {...rest}
        >
            {children}
        </button>
    )
}
