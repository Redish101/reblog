import { HTMLAttributes, ReactNode } from "react";

export interface ButtonProps extends HTMLAttributes<HTMLButtonElement> {
    children: ReactNode,
    variant?: "filled" | "outlined" | "text"
}

export default function Button({
    children,
    ...rest
}: ButtonProps) {
    return (
        <button
            className="
                h-9 rounded flex flex-col justify-center items-center px-5 b-none bg-primary text-white
                hover:bg-secondly
                ease-in-out transition-all duration-300
            "
            {...rest}
        >
            {children}
        </button>
    )
}
