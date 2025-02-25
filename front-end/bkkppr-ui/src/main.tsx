import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { ThemeProvider } from '@emotion/react'
import createTheme from '@mui/material/styles/createTheme'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Auth from './Components/Login-Auth/login.tsx'
import Welcome from './Components/Welcome/welcome.tsx'
import { CssBaseline } from '@mui/material'

const theme = createTheme({
  components: {
    MuiOutlinedInput: {
      styleOverrides: {
        input: {
          color: '#d1d5de', // input text color override
        },
        root: {
          '& .MuiOutlinedInput-notchedOutline': {
            borderColor: '#d1d5de',
          },
          '&:hover .MuiOutlinedInput-notchedOutline': {
            borderColor: '#d1d5de',
          },
          '&.Mui-focused .MuiOutlinedInput-notchedOutline': {
            borderColor: '#d1d5de',
          },
        },
      },
    },
    MuiInputLabel: {
      styleOverrides: {
        root: {
          color: '#d1d5de', // label text color override
          '&.Mui-focused': {
            color: '#d1d5de',
          },
        },
      },
    },
  },
});

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ThemeProvider theme={theme}>
      <BrowserRouter>
          <Routes>
            <Route path="/" element={<Auth />} />
            <Route path="/welcome" element={<Welcome />} />
          </Routes>
      </BrowserRouter>
    </ThemeProvider>
  </StrictMode>,
)
