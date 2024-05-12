import { defineConfig, presetUno } from 'unocss'

export default defineConfig({
    presets: [
        presetUno()
    ],
    theme: {
        colors: {
            primary: '#2b6672',
            secondly: '#3c727d'
        },
        height: {
            9: '2.5rem'
        },
        borderRadius: {
            DEFAULT: '2.5rem'
        },
        breakpoints: {
            5: '1.5rem'
        }
    }
})