import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App.tsx'

import '@/global.css'

import 'mdui'
import { setColorScheme } from 'mdui'
import 'mdui/mdui.css'

setColorScheme('#2b6672')

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
