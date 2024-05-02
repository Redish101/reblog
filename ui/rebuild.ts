import { pnpm, copy } from "rebuild"

pnpm("ui", {
    script: "build",
    output: "dist"
})

copy("dist", {
    deps: [
        ":ui"
    ],
    output: "//internal/ui/dist"
})
