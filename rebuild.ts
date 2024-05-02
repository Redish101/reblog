import { go } from "rebuild"

go("server", {
    deps: [
        "//ui:dist"
    ],
    flags: [
        `-ldflags "-w -s" -gcflags "-N -l"`
    ],
    output: [
        "bin/reblog"
    ]
})