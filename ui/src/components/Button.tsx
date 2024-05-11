import { HTMLAttributes, ReactNode } from "react";

export interface ButtonProps extends HTMLAttributes<HTMLButtonElement> {
    children: ReactNode,
}

export default function Button({
    children,
    ...rest
}: ButtonProps) {
    return (
        <button
            {...rest}
        >
            {children}
        </button>
    )
}
