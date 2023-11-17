import { env } from "bun"
import { Server } from "http"

 const serve =  Bun.serve({
    fetch() {
        return new Response("")
    },
    port:process.env.PORT
})
console.log(`Listening on port ${serve.port} ...`)