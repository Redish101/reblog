import { InputHTMLAttributes } from "react";

export interface TextFieldProps extends InputHTMLAttributes<HTMLElement> {
    label: string;
    placeholder: string;
    variant?: "filled" | "outlined"
}

export default function TextField({
    label,
    placeholder,
    variant = "filled",
    ...rest
}: TextFieldProps) {
    return (
        <div>
            <label
                className="pos-absolute px-4 pt-2 text-sm text-gray-500"
            >{label}</label>
            <input
                placeholder={placeholder}
                className="border-b-2 b-t-none b-x-none pt-7 rounded-t-md text-base flex justify-start items-start px-4 pb-2 bg-gray-200 placeholder-gray-600 placeholder-text-base outline-none focus:b-b-primary ease-in-out transition-all duration-300
                "
                {...rest}
            />
        </div>
    )
}