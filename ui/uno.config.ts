import { defineConfig, presetUno } from 'unocss'

export default defineConfig({
    presets: [
        presetUno()
    ],
    theme: {
        colors: {
            primary: '#2b6672',
            secondary: '#3c727d',
            bg: '#fbfcfd',
            hover: '#eaf0f2'
        },
        height: {
            9: '2.5rem'
        },
        borderRadius: {
            DEFAULT: '2.5rem'
        },
        fontWeight: {
            semibold: "500",
        },
        breakpoints: {
            5: '1.5rem'
        }
    }
})