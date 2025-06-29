import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'

import './main.css'
import { AuthProvider } from './auth/auth-context'
import { BrowserRouter, Route, Routes } from 'react-router-dom'
import { ProtectedRoutes } from './auth/protected-routes'
import { GuestRoutes } from './auth/guest-routes'
import HomePage from './pages/home-page'
import LoginPage from './pages/login-page'
import { Toaster } from './components/ui/sonner'
import RegisterPage from './pages/register-page'
import Dashboard from './pages/dashboard'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <AuthProvider>
      <BrowserRouter>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route element={<ProtectedRoutes />}>
            <Route path="/dashboard" element={<Dashboard />} />
          </Route>
          <Route element={<GuestRoutes />}>
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegisterPage />} />
          </Route>
          <Route path="*" element={<div>404 Not Found</div>} />
        </Routes>
      </BrowserRouter>
    </AuthProvider>
    <Toaster />
  </StrictMode>,
)
